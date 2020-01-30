// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/task"
)

type FakeBoshClient struct {
	DeployStub        func([]byte, string, *log.Logger, *boshdirector.AsyncTaskReporter) (int, error)
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		arg1 []byte
		arg2 string
		arg3 *log.Logger
		arg4 *boshdirector.AsyncTaskReporter
	}
	deployReturns struct {
		result1 int
		result2 error
	}
	deployReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	GetConfigsStub        func(string, *log.Logger) ([]boshdirector.BoshConfig, error)
	getConfigsMutex       sync.RWMutex
	getConfigsArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getConfigsReturns struct {
		result1 []boshdirector.BoshConfig
		result2 error
	}
	getConfigsReturnsOnCall map[int]struct {
		result1 []boshdirector.BoshConfig
		result2 error
	}
	GetDeploymentStub        func(string, *log.Logger) ([]byte, bool, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getDeploymentReturns struct {
		result1 []byte
		result2 bool
		result3 error
	}
	getDeploymentReturnsOnCall map[int]struct {
		result1 []byte
		result2 bool
		result3 error
	}
	GetEventsStub        func(string, string, *log.Logger) ([]boshdirector.BoshEvent, error)
	getEventsMutex       sync.RWMutex
	getEventsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}
	getEventsReturns struct {
		result1 []boshdirector.BoshEvent
		result2 error
	}
	getEventsReturnsOnCall map[int]struct {
		result1 []boshdirector.BoshEvent
		result2 error
	}
	GetNormalisedTasksByContextStub        func(string, string, *log.Logger) (boshdirector.BoshTasks, error)
	getNormalisedTasksByContextMutex       sync.RWMutex
	getNormalisedTasksByContextArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}
	getNormalisedTasksByContextReturns struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	getNormalisedTasksByContextReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	GetTaskStub        func(int, *log.Logger) (boshdirector.BoshTask, error)
	getTaskMutex       sync.RWMutex
	getTaskArgsForCall []struct {
		arg1 int
		arg2 *log.Logger
	}
	getTaskReturns struct {
		result1 boshdirector.BoshTask
		result2 error
	}
	getTaskReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTask
		result2 error
	}
	GetTasksInProgressStub        func(string, *log.Logger) (boshdirector.BoshTasks, error)
	getTasksInProgressMutex       sync.RWMutex
	getTasksInProgressArgsForCall []struct {
		arg1 string
		arg2 *log.Logger
	}
	getTasksInProgressReturns struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	getTasksInProgressReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	RecreateStub        func(string, string, *log.Logger, *boshdirector.AsyncTaskReporter) (int, error)
	recreateMutex       sync.RWMutex
	recreateArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
		arg4 *boshdirector.AsyncTaskReporter
	}
	recreateReturns struct {
		result1 int
		result2 error
	}
	recreateReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	UpdateConfigStub        func(string, string, []byte, *log.Logger) error
	updateConfigMutex       sync.RWMutex
	updateConfigArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 *log.Logger
	}
	updateConfigReturns struct {
		result1 error
	}
	updateConfigReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshClient) Deploy(arg1 []byte, arg2 string, arg3 *log.Logger, arg4 *boshdirector.AsyncTaskReporter) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		arg1 []byte
		arg2 string
		arg3 *log.Logger
		arg4 *boshdirector.AsyncTaskReporter
	}{arg1Copy, arg2, arg3, arg4})
	fake.recordInvocation("Deploy", []interface{}{arg1Copy, arg2, arg3, arg4})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deployReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeBoshClient) DeployCalls(stub func([]byte, string, *log.Logger, *boshdirector.AsyncTaskReporter) (int, error)) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = stub
}

func (fake *FakeBoshClient) DeployArgsForCall(i int) ([]byte, string, *log.Logger, *boshdirector.AsyncTaskReporter) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	argsForCall := fake.deployArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBoshClient) DeployReturns(result1 int, result2 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeployReturnsOnCall(i int, result1 int, result2 error) {
	fake.deployMutex.Lock()
	defer fake.deployMutex.Unlock()
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetConfigs(arg1 string, arg2 *log.Logger) ([]boshdirector.BoshConfig, error) {
	fake.getConfigsMutex.Lock()
	ret, specificReturn := fake.getConfigsReturnsOnCall[len(fake.getConfigsArgsForCall)]
	fake.getConfigsArgsForCall = append(fake.getConfigsArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetConfigs", []interface{}{arg1, arg2})
	fake.getConfigsMutex.Unlock()
	if fake.GetConfigsStub != nil {
		return fake.GetConfigsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getConfigsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetConfigsCallCount() int {
	fake.getConfigsMutex.RLock()
	defer fake.getConfigsMutex.RUnlock()
	return len(fake.getConfigsArgsForCall)
}

func (fake *FakeBoshClient) GetConfigsCalls(stub func(string, *log.Logger) ([]boshdirector.BoshConfig, error)) {
	fake.getConfigsMutex.Lock()
	defer fake.getConfigsMutex.Unlock()
	fake.GetConfigsStub = stub
}

func (fake *FakeBoshClient) GetConfigsArgsForCall(i int) (string, *log.Logger) {
	fake.getConfigsMutex.RLock()
	defer fake.getConfigsMutex.RUnlock()
	argsForCall := fake.getConfigsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBoshClient) GetConfigsReturns(result1 []boshdirector.BoshConfig, result2 error) {
	fake.getConfigsMutex.Lock()
	defer fake.getConfigsMutex.Unlock()
	fake.GetConfigsStub = nil
	fake.getConfigsReturns = struct {
		result1 []boshdirector.BoshConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetConfigsReturnsOnCall(i int, result1 []boshdirector.BoshConfig, result2 error) {
	fake.getConfigsMutex.Lock()
	defer fake.getConfigsMutex.Unlock()
	fake.GetConfigsStub = nil
	if fake.getConfigsReturnsOnCall == nil {
		fake.getConfigsReturnsOnCall = make(map[int]struct {
			result1 []boshdirector.BoshConfig
			result2 error
		})
	}
	fake.getConfigsReturnsOnCall[i] = struct {
		result1 []boshdirector.BoshConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDeployment(arg1 string, arg2 *log.Logger) ([]byte, bool, error) {
	fake.getDeploymentMutex.Lock()
	ret, specificReturn := fake.getDeploymentReturnsOnCall[len(fake.getDeploymentArgsForCall)]
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetDeployment", []interface{}{arg1, arg2})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDeploymentReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBoshClient) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeBoshClient) GetDeploymentCalls(stub func(string, *log.Logger) ([]byte, bool, error)) {
	fake.getDeploymentMutex.Lock()
	defer fake.getDeploymentMutex.Unlock()
	fake.GetDeploymentStub = stub
}

func (fake *FakeBoshClient) GetDeploymentArgsForCall(i int) (string, *log.Logger) {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	argsForCall := fake.getDeploymentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBoshClient) GetDeploymentReturns(result1 []byte, result2 bool, result3 error) {
	fake.getDeploymentMutex.Lock()
	defer fake.getDeploymentMutex.Unlock()
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetDeploymentReturnsOnCall(i int, result1 []byte, result2 bool, result3 error) {
	fake.getDeploymentMutex.Lock()
	defer fake.getDeploymentMutex.Unlock()
	fake.GetDeploymentStub = nil
	if fake.getDeploymentReturnsOnCall == nil {
		fake.getDeploymentReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 bool
			result3 error
		})
	}
	fake.getDeploymentReturnsOnCall[i] = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetEvents(arg1 string, arg2 string, arg3 *log.Logger) ([]boshdirector.BoshEvent, error) {
	fake.getEventsMutex.Lock()
	ret, specificReturn := fake.getEventsReturnsOnCall[len(fake.getEventsArgsForCall)]
	fake.getEventsArgsForCall = append(fake.getEventsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetEvents", []interface{}{arg1, arg2, arg3})
	fake.getEventsMutex.Unlock()
	if fake.GetEventsStub != nil {
		return fake.GetEventsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getEventsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetEventsCallCount() int {
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	return len(fake.getEventsArgsForCall)
}

func (fake *FakeBoshClient) GetEventsCalls(stub func(string, string, *log.Logger) ([]boshdirector.BoshEvent, error)) {
	fake.getEventsMutex.Lock()
	defer fake.getEventsMutex.Unlock()
	fake.GetEventsStub = stub
}

func (fake *FakeBoshClient) GetEventsArgsForCall(i int) (string, string, *log.Logger) {
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	argsForCall := fake.getEventsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBoshClient) GetEventsReturns(result1 []boshdirector.BoshEvent, result2 error) {
	fake.getEventsMutex.Lock()
	defer fake.getEventsMutex.Unlock()
	fake.GetEventsStub = nil
	fake.getEventsReturns = struct {
		result1 []boshdirector.BoshEvent
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetEventsReturnsOnCall(i int, result1 []boshdirector.BoshEvent, result2 error) {
	fake.getEventsMutex.Lock()
	defer fake.getEventsMutex.Unlock()
	fake.GetEventsStub = nil
	if fake.getEventsReturnsOnCall == nil {
		fake.getEventsReturnsOnCall = make(map[int]struct {
			result1 []boshdirector.BoshEvent
			result2 error
		})
	}
	fake.getEventsReturnsOnCall[i] = struct {
		result1 []boshdirector.BoshEvent
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContext(arg1 string, arg2 string, arg3 *log.Logger) (boshdirector.BoshTasks, error) {
	fake.getNormalisedTasksByContextMutex.Lock()
	ret, specificReturn := fake.getNormalisedTasksByContextReturnsOnCall[len(fake.getNormalisedTasksByContextArgsForCall)]
	fake.getNormalisedTasksByContextArgsForCall = append(fake.getNormalisedTasksByContextArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetNormalisedTasksByContext", []interface{}{arg1, arg2, arg3})
	fake.getNormalisedTasksByContextMutex.Unlock()
	if fake.GetNormalisedTasksByContextStub != nil {
		return fake.GetNormalisedTasksByContextStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getNormalisedTasksByContextReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextCallCount() int {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	return len(fake.getNormalisedTasksByContextArgsForCall)
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextCalls(stub func(string, string, *log.Logger) (boshdirector.BoshTasks, error)) {
	fake.getNormalisedTasksByContextMutex.Lock()
	defer fake.getNormalisedTasksByContextMutex.Unlock()
	fake.GetNormalisedTasksByContextStub = stub
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextArgsForCall(i int) (string, string, *log.Logger) {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	argsForCall := fake.getNormalisedTasksByContextArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturns(result1 boshdirector.BoshTasks, result2 error) {
	fake.getNormalisedTasksByContextMutex.Lock()
	defer fake.getNormalisedTasksByContextMutex.Unlock()
	fake.GetNormalisedTasksByContextStub = nil
	fake.getNormalisedTasksByContextReturns = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturnsOnCall(i int, result1 boshdirector.BoshTasks, result2 error) {
	fake.getNormalisedTasksByContextMutex.Lock()
	defer fake.getNormalisedTasksByContextMutex.Unlock()
	fake.GetNormalisedTasksByContextStub = nil
	if fake.getNormalisedTasksByContextReturnsOnCall == nil {
		fake.getNormalisedTasksByContextReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTasks
			result2 error
		})
	}
	fake.getNormalisedTasksByContextReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTask(arg1 int, arg2 *log.Logger) (boshdirector.BoshTask, error) {
	fake.getTaskMutex.Lock()
	ret, specificReturn := fake.getTaskReturnsOnCall[len(fake.getTaskArgsForCall)]
	fake.getTaskArgsForCall = append(fake.getTaskArgsForCall, struct {
		arg1 int
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetTask", []interface{}{arg1, arg2})
	fake.getTaskMutex.Unlock()
	if fake.GetTaskStub != nil {
		return fake.GetTaskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTaskReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetTaskCallCount() int {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	return len(fake.getTaskArgsForCall)
}

func (fake *FakeBoshClient) GetTaskCalls(stub func(int, *log.Logger) (boshdirector.BoshTask, error)) {
	fake.getTaskMutex.Lock()
	defer fake.getTaskMutex.Unlock()
	fake.GetTaskStub = stub
}

func (fake *FakeBoshClient) GetTaskArgsForCall(i int) (int, *log.Logger) {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	argsForCall := fake.getTaskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBoshClient) GetTaskReturns(result1 boshdirector.BoshTask, result2 error) {
	fake.getTaskMutex.Lock()
	defer fake.getTaskMutex.Unlock()
	fake.GetTaskStub = nil
	fake.getTaskReturns = struct {
		result1 boshdirector.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTaskReturnsOnCall(i int, result1 boshdirector.BoshTask, result2 error) {
	fake.getTaskMutex.Lock()
	defer fake.getTaskMutex.Unlock()
	fake.GetTaskStub = nil
	if fake.getTaskReturnsOnCall == nil {
		fake.getTaskReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTask
			result2 error
		})
	}
	fake.getTaskReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasksInProgress(arg1 string, arg2 *log.Logger) (boshdirector.BoshTasks, error) {
	fake.getTasksInProgressMutex.Lock()
	ret, specificReturn := fake.getTasksInProgressReturnsOnCall[len(fake.getTasksInProgressArgsForCall)]
	fake.getTasksInProgressArgsForCall = append(fake.getTasksInProgressArgsForCall, struct {
		arg1 string
		arg2 *log.Logger
	}{arg1, arg2})
	fake.recordInvocation("GetTasksInProgress", []interface{}{arg1, arg2})
	fake.getTasksInProgressMutex.Unlock()
	if fake.GetTasksInProgressStub != nil {
		return fake.GetTasksInProgressStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTasksInProgressReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetTasksInProgressCallCount() int {
	fake.getTasksInProgressMutex.RLock()
	defer fake.getTasksInProgressMutex.RUnlock()
	return len(fake.getTasksInProgressArgsForCall)
}

func (fake *FakeBoshClient) GetTasksInProgressCalls(stub func(string, *log.Logger) (boshdirector.BoshTasks, error)) {
	fake.getTasksInProgressMutex.Lock()
	defer fake.getTasksInProgressMutex.Unlock()
	fake.GetTasksInProgressStub = stub
}

func (fake *FakeBoshClient) GetTasksInProgressArgsForCall(i int) (string, *log.Logger) {
	fake.getTasksInProgressMutex.RLock()
	defer fake.getTasksInProgressMutex.RUnlock()
	argsForCall := fake.getTasksInProgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBoshClient) GetTasksInProgressReturns(result1 boshdirector.BoshTasks, result2 error) {
	fake.getTasksInProgressMutex.Lock()
	defer fake.getTasksInProgressMutex.Unlock()
	fake.GetTasksInProgressStub = nil
	fake.getTasksInProgressReturns = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasksInProgressReturnsOnCall(i int, result1 boshdirector.BoshTasks, result2 error) {
	fake.getTasksInProgressMutex.Lock()
	defer fake.getTasksInProgressMutex.Unlock()
	fake.GetTasksInProgressStub = nil
	if fake.getTasksInProgressReturnsOnCall == nil {
		fake.getTasksInProgressReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTasks
			result2 error
		})
	}
	fake.getTasksInProgressReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Recreate(arg1 string, arg2 string, arg3 *log.Logger, arg4 *boshdirector.AsyncTaskReporter) (int, error) {
	fake.recreateMutex.Lock()
	ret, specificReturn := fake.recreateReturnsOnCall[len(fake.recreateArgsForCall)]
	fake.recreateArgsForCall = append(fake.recreateArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *log.Logger
		arg4 *boshdirector.AsyncTaskReporter
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Recreate", []interface{}{arg1, arg2, arg3, arg4})
	fake.recreateMutex.Unlock()
	if fake.RecreateStub != nil {
		return fake.RecreateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.recreateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) RecreateCallCount() int {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	return len(fake.recreateArgsForCall)
}

func (fake *FakeBoshClient) RecreateCalls(stub func(string, string, *log.Logger, *boshdirector.AsyncTaskReporter) (int, error)) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = stub
}

func (fake *FakeBoshClient) RecreateArgsForCall(i int) (string, string, *log.Logger, *boshdirector.AsyncTaskReporter) {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	argsForCall := fake.recreateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBoshClient) RecreateReturns(result1 int, result2 error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = nil
	fake.recreateReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) RecreateReturnsOnCall(i int, result1 int, result2 error) {
	fake.recreateMutex.Lock()
	defer fake.recreateMutex.Unlock()
	fake.RecreateStub = nil
	if fake.recreateReturnsOnCall == nil {
		fake.recreateReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.recreateReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) UpdateConfig(arg1 string, arg2 string, arg3 []byte, arg4 *log.Logger) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.updateConfigMutex.Lock()
	ret, specificReturn := fake.updateConfigReturnsOnCall[len(fake.updateConfigArgsForCall)]
	fake.updateConfigArgsForCall = append(fake.updateConfigArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 *log.Logger
	}{arg1, arg2, arg3Copy, arg4})
	fake.recordInvocation("UpdateConfig", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.updateConfigMutex.Unlock()
	if fake.UpdateConfigStub != nil {
		return fake.UpdateConfigStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateConfigReturns
	return fakeReturns.result1
}

func (fake *FakeBoshClient) UpdateConfigCallCount() int {
	fake.updateConfigMutex.RLock()
	defer fake.updateConfigMutex.RUnlock()
	return len(fake.updateConfigArgsForCall)
}

func (fake *FakeBoshClient) UpdateConfigCalls(stub func(string, string, []byte, *log.Logger) error) {
	fake.updateConfigMutex.Lock()
	defer fake.updateConfigMutex.Unlock()
	fake.UpdateConfigStub = stub
}

func (fake *FakeBoshClient) UpdateConfigArgsForCall(i int) (string, string, []byte, *log.Logger) {
	fake.updateConfigMutex.RLock()
	defer fake.updateConfigMutex.RUnlock()
	argsForCall := fake.updateConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBoshClient) UpdateConfigReturns(result1 error) {
	fake.updateConfigMutex.Lock()
	defer fake.updateConfigMutex.Unlock()
	fake.UpdateConfigStub = nil
	fake.updateConfigReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshClient) UpdateConfigReturnsOnCall(i int, result1 error) {
	fake.updateConfigMutex.Lock()
	defer fake.updateConfigMutex.Unlock()
	fake.UpdateConfigStub = nil
	if fake.updateConfigReturnsOnCall == nil {
		fake.updateConfigReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateConfigReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.getConfigsMutex.RLock()
	defer fake.getConfigsMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	fake.getTasksInProgressMutex.RLock()
	defer fake.getTasksInProgressMutex.RUnlock()
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	fake.updateConfigMutex.RLock()
	defer fake.updateConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBoshClient) recordInvocation(key string, args []interface{}) {
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

var _ task.BoshClient = new(FakeBoshClient)
