package a // import "github.com/maxbrunsfeld/counterfeiter/fixtures/dup_packages/a"

import "github.com/maxbrunsfeld/counterfeiter/fixtures/dup_packages/a/foo"

//go:generate go run github.com/maxbrunsfeld/counterfeiter . A
type A interface {
	V1() foo.I
}
