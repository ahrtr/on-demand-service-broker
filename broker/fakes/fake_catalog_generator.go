// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceCatalogGenerator struct {
	ServicesStub        func(context.Context) ([]brokerapi.Service, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
		arg1 context.Context
	}
	servicesReturns struct {
		result1 []brokerapi.Service
		result2 error
	}
	servicesReturnsOnCall map[int]struct {
		result1 []brokerapi.Service
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceCatalogGenerator) Services(arg1 context.Context) ([]brokerapi.Service, error) {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Services", []interface{}{arg1})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.servicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceCatalogGenerator) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeServiceCatalogGenerator) ServicesCalls(stub func(context.Context) ([]brokerapi.Service, error)) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = stub
}

func (fake *FakeServiceCatalogGenerator) ServicesArgsForCall(i int) context.Context {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	argsForCall := fake.servicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceCatalogGenerator) ServicesReturns(result1 []brokerapi.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []brokerapi.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceCatalogGenerator) ServicesReturnsOnCall(i int, result1 []brokerapi.Service, result2 error) {
	fake.servicesMutex.Lock()
	defer fake.servicesMutex.Unlock()
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []brokerapi.Service
			result2 error
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []brokerapi.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceCatalogGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceCatalogGenerator) recordInvocation(key string, args []interface{}) {
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
