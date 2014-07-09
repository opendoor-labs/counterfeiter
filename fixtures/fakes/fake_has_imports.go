// This file was generated by counterfeiter
package fakes

import (
	"io"
	"net/http"
	some_alias "os"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeHasImports struct {
	DoThingsStub        func(io.Writer, *some_alias.File) *http.Client
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		arg1 io.Writer
		arg2 *some_alias.File
	}
	doThingsReturns struct {
		result1 *http.Client
	}
}

func (fake *FakeHasImports) DoThings(arg1 io.Writer, arg2 *some_alias.File) *http.Client {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		arg1 io.Writer
		arg2 *some_alias.File
	}{arg1, arg2})
	if fake.DoThingsStub != nil {
		return fake.DoThingsStub(arg1, arg2)
	} else {
		return fake.doThingsReturns.result1
	}
}

func (fake *FakeHasImports) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeHasImports) DoThingsArgsForCall(i int) (io.Writer, *some_alias.File) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return fake.doThingsArgsForCall[i].arg1, fake.doThingsArgsForCall[i].arg2
}

func (fake *FakeHasImports) DoThingsReturns(result1 *http.Client) {
	fake.doThingsReturns = struct {
		result1 *http.Client
	}{result1}
}

var _ fixtures.HasImports = new(FakeHasImports)