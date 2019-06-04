// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeFirstInterface struct {
	DoThingsStub        func()
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFirstInterface) DoThings() {
	fake.doThingsMutex.Lock()
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
	}{})
	fake.recordInvocation("DoThings", []interface{}{})
	fake.doThingsMutex.Unlock()
	if fake.DoThingsStub != nil {
		fake.DoThingsStub()
	}
}

func (fake *FakeFirstInterface) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeFirstInterface) DoThingsCalls(stub func()) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = stub
}

func (fake *FakeFirstInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFirstInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fixtures.FirstInterface = new(FakeFirstInterface)
