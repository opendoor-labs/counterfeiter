package generator

import (
	"fmt"
	"go/build"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/tools/go/gcexportdata"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
)

func (f *Fake) loadPackages(c Cacher, workingDir, exportDataFile string) error {
	log.Println("loading packages...")
	p, ok := c.Load(f.TargetPackage)
	if ok {
		f.Packages = p
		log.Printf("loaded %v packages from cache\n", len(f.Packages))
		return nil
	}

	importPath := f.TargetPackage
	if !filepath.IsAbs(importPath) {
		bp, err := build.Import(f.TargetPackage, workingDir, build.FindOnly)
		if err != nil {
			return err
		}
		importPath = bp.ImportPath
	}

	var err error
	if exportDataFile != "" {
		p, err = f.loadWithGCExportData(exportDataFile, importPath)
	} else {
		p, err = f.loadWithPackages(workingDir, importPath)
	}
	if err != nil {
		return err
	}

	f.Packages = p
	c.Store(f.TargetPackage, p)
	log.Printf("loaded %v packages\n", len(f.Packages))
	return nil
}

func (f *Fake) loadWithGCExportData(exportDataFile, importPath string) ([]*types.Package, error) {
	exportFile, err := os.Open(exportDataFile)
	if err != nil {
		return nil, err
	}
	defer exportFile.Close()

	r, err := gcexportdata.NewReader(exportFile)
	if err != nil {
		return nil, err
	}

	// Decode the export data.
	fset := token.NewFileSet()
	imports := make(map[string]*types.Package)
	pkg, err := gcexportdata.Read(r, fset, imports, importPath)
	if err != nil {
		return nil, err
	}

	return []*types.Package{pkg}, nil
}

func (f *Fake) loadWithPackages(workingDir, importPath string) ([]*types.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes,
		Dir:   workingDir,
		Tests: true,
	}, importPath)
	if err != nil {
		return nil, err
	}

	for i := range pkgs {
		if len(pkgs[i].Errors) > 0 {
			if i == 0 {
				err = pkgs[i].Errors[0]
			}
			for j := range pkgs[i].Errors {
				log.Printf("error loading packages: %v", strings.TrimPrefix(fmt.Sprintf("%v", pkgs[i].Errors[j]), "-: "))
			}
		}
	}
	if err != nil {
		return nil, err
	}

	var p []*types.Package
	for _, pkg := range pkgs {
		p = append(p, pkg.Types)
	}
	return p, nil
}

func (f *Fake) findPackage() error {
	var target *types.TypeName
	var pkg *types.Package
	for i := range f.Packages {
		if f.Packages[i] == nil || f.Packages[i].Scope() == nil {
			continue
		}
		pkg = f.Packages[i]
		if f.Mode == Package {
			break
		}

		raw := pkg.Scope().Lookup(f.TargetName)
		if raw != nil {
			if typeName, ok := raw.(*types.TypeName); ok {
				target = typeName
				break
			}
		}
		pkg = nil
	}
	if pkg == nil {
		switch f.Mode {
		case Package:
			return fmt.Errorf("cannot find package with name: %s", f.TargetPackage)
		case InterfaceOrFunction:
			return fmt.Errorf("cannot find package with target: %s", f.TargetName)
		}
	}
	f.Target = target
	f.Package = pkg
	f.TargetPackage = imports.VendorlessPath(pkg.Path())
	t := f.Imports.Add(pkg.Name(), f.TargetPackage)
	f.TargetAlias = t.Alias
	if f.Mode != Package {
		f.TargetName = target.Name()
	}

	if f.Mode == InterfaceOrFunction {
		if !f.IsInterface() && !f.IsFunction() {
			return fmt.Errorf("cannot generate an fake for %s because it is not an interface or function", f.TargetName)
		}
	}

	if f.IsInterface() {
		log.Printf("Found interface with name: [%s]\n", f.TargetName)
	}
	if f.IsFunction() {
		log.Printf("Found function with name: [%s]\n", f.TargetName)
	}
	if f.Mode == Package {
		log.Printf("Found package with name: [%s]\n", f.TargetPackage)
	}
	return nil
}

// addImportsFor inspects the given type and adds imports to the fake if importable
// types are found.
func (f *Fake) addImportsFor(typ types.Type) {
	if typ == nil {
		return
	}

	switch t := typ.(type) {
	case *types.Basic:
		return
	case *types.Pointer:
		f.addImportsFor(t.Elem())
	case *types.Map:
		f.addImportsFor(t.Key())
		f.addImportsFor(t.Elem())
	case *types.Chan:
		f.addImportsFor(t.Elem())
	case *types.Named:
		if t.Obj() != nil && t.Obj().Pkg() != nil {
			f.Imports.Add(t.Obj().Pkg().Name(), t.Obj().Pkg().Path())
		}
	case *types.Slice:
		f.addImportsFor(t.Elem())
	case *types.Array:
		f.addImportsFor(t.Elem())
	case *types.Interface:
		return
	case *types.Signature:
		f.addTypesForMethod(t)
	case *types.Struct:
		for i := 0; i < t.NumFields(); i++ {
			f.addImportsFor(t.Field(i).Type())
		}
	default:
		log.Printf("!!! WARNING: Missing case for type %s\n", reflect.TypeOf(typ).String())
	}
}
