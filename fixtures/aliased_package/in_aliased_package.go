package the_aliased_package // import "github.com/maxbrunsfeld/counterfeiter/fixtures/aliased_package"

//go:generate go run github.com/maxbrunsfeld/counterfeiter . InAliasedPackage
type InAliasedPackage interface {
	Stuff(int) string
}
