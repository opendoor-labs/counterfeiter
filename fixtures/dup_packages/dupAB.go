package dup_packages // import "github.com/maxbrunsfeld/counterfeiter/v6/fixtures/dup_packages"

//go:generate counterfeiter . DupAB
type DupAB interface {
	DupA
	DupB
}
