// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeSecondInterface struct {
	EmbeddedMethodStub        func() string
	embeddedMethodMutex       sync.RWMutex
	embeddedMethodArgsForCall []struct {
	}
	embeddedMethodReturns struct {
		result1 string
	}
	embeddedMethodReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecondInterface) EmbeddedMethod() string {
	fake.embeddedMethodMutex.Lock()
	ret, specificReturn := fake.embeddedMethodReturnsOnCall[len(fake.embeddedMethodArgsForCall)]
	fake.embeddedMethodArgsForCall = append(fake.embeddedMethodArgsForCall, struct {
	}{})
	fake.recordInvocation("EmbeddedMethod", []interface{}{})
	fake.embeddedMethodMutex.Unlock()
	if fake.EmbeddedMethodStub != nil {
		return fake.EmbeddedMethodStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.embeddedMethodReturns
	return fakeReturns.result1
}

func (fake *FakeSecondInterface) EmbeddedMethodCallCount() int {
	fake.embeddedMethodMutex.RLock()
	defer fake.embeddedMethodMutex.RUnlock()
	return len(fake.embeddedMethodArgsForCall)
}

func (fake *FakeSecondInterface) EmbeddedMethodCalls(stub func() string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = stub
}

func (fake *FakeSecondInterface) EmbeddedMethodReturns(result1 string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = nil
	fake.embeddedMethodReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSecondInterface) EmbeddedMethodReturnsOnCall(i int, result1 string) {
	fake.embeddedMethodMutex.Lock()
	defer fake.embeddedMethodMutex.Unlock()
	fake.EmbeddedMethodStub = nil
	if fake.embeddedMethodReturnsOnCall == nil {
		fake.embeddedMethodReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.embeddedMethodReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSecondInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.embeddedMethodMutex.RLock()
	defer fake.embeddedMethodMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecondInterface) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.SecondInterface = new(FakeSecondInterface)
