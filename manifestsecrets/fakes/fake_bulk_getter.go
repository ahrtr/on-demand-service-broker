// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/manifestsecrets"
)

type FakeBulkGetter struct {
	BulkGetStub        func([][]byte) (map[string]string, error)
	bulkGetMutex       sync.RWMutex
	bulkGetArgsForCall []struct {
		arg1 [][]byte
	}
	bulkGetReturns struct {
		result1 map[string]string
		result2 error
	}
	bulkGetReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBulkGetter) BulkGet(arg1 [][]byte) (map[string]string, error) {
	var arg1Copy [][]byte
	if arg1 != nil {
		arg1Copy = make([][]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.bulkGetMutex.Lock()
	ret, specificReturn := fake.bulkGetReturnsOnCall[len(fake.bulkGetArgsForCall)]
	fake.bulkGetArgsForCall = append(fake.bulkGetArgsForCall, struct {
		arg1 [][]byte
	}{arg1Copy})
	fake.recordInvocation("BulkGet", []interface{}{arg1Copy})
	fake.bulkGetMutex.Unlock()
	if fake.BulkGetStub != nil {
		return fake.BulkGetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bulkGetReturns.result1, fake.bulkGetReturns.result2
}

func (fake *FakeBulkGetter) BulkGetCallCount() int {
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	return len(fake.bulkGetArgsForCall)
}

func (fake *FakeBulkGetter) BulkGetArgsForCall(i int) [][]byte {
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	return fake.bulkGetArgsForCall[i].arg1
}

func (fake *FakeBulkGetter) BulkGetReturns(result1 map[string]string, result2 error) {
	fake.BulkGetStub = nil
	fake.bulkGetReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeBulkGetter) BulkGetReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.BulkGetStub = nil
	if fake.bulkGetReturnsOnCall == nil {
		fake.bulkGetReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.bulkGetReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeBulkGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bulkGetMutex.RLock()
	defer fake.bulkGetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBulkGetter) recordInvocation(key string, args []interface{}) {
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

var _ manifestsecrets.BulkGetter = new(FakeBulkGetter)
