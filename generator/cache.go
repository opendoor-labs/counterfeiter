package generator

import "go/types"

type Cache struct {
	packageMap map[string][]*types.Package
}

type FakeCache struct{}

func (c *FakeCache) Load(packagePath string) ([]*types.Package, bool)    { return nil, false }
func (c *FakeCache) Store(packagePath string, packages []*types.Package) {}

type Cacher interface {
	Load(packagePath string) ([]*types.Package, bool)
	Store(packagePath string, packages []*types.Package)
}

func (c *Cache) Load(packagePath string) ([]*types.Package, bool) {
	p, ok := c.packageMap[packagePath]
	if !ok {
		return nil, false
	}
	return p, ok
}

func (c *Cache) Store(packagePath string, packages []*types.Package) {
	if c.packageMap == nil {
		c.packageMap = map[string][]*types.Package{}
	}
	c.packageMap[packagePath] = packages
}
