// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"

	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/service"
	"github.com/pivotal-cf/on-demand-service-broker/upgrader"
)

type FakeListener struct {
	FailedToRefreshInstanceInfoStub        func(instance string)
	failedToRefreshInstanceInfoMutex       sync.RWMutex
	failedToRefreshInstanceInfoArgsForCall []struct {
		instance string
	}
	StartingStub        func(maxInFlight int)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		maxInFlight int
	}
	RetryAttemptStub        func(num, limit int)
	retryAttemptMutex       sync.RWMutex
	retryAttemptArgsForCall []struct {
		num   int
		limit int
	}
	RetryCanariesAttemptStub        func(num, limit, remainingCanaries int)
	retryCanariesAttemptMutex       sync.RWMutex
	retryCanariesAttemptArgsForCall []struct {
		num               int
		limit             int
		remainingCanaries int
	}
	InstancesToUpgradeStub        func(instances []service.Instance)
	instancesToUpgradeMutex       sync.RWMutex
	instancesToUpgradeArgsForCall []struct {
		instances []service.Instance
	}
	InstanceUpgradeStartingStub        func(instance string, index int, totalInstances int, isCanary bool)
	instanceUpgradeStartingMutex       sync.RWMutex
	instanceUpgradeStartingArgsForCall []struct {
		instance       string
		index          int
		totalInstances int
		isCanary       bool
	}
	InstanceUpgradeStartResultStub        func(instance string, status services.UpgradeOperationType)
	instanceUpgradeStartResultMutex       sync.RWMutex
	instanceUpgradeStartResultArgsForCall []struct {
		instance string
		status   services.UpgradeOperationType
	}
	InstanceUpgradedStub        func(instance string, result string)
	instanceUpgradedMutex       sync.RWMutex
	instanceUpgradedArgsForCall []struct {
		instance string
		result   string
	}
	WaitingForStub        func(instance string, boshTaskId int)
	waitingForMutex       sync.RWMutex
	waitingForArgsForCall []struct {
		instance   string
		boshTaskId int
	}
	ProgressStub        func(pollingInterval time.Duration, orphanCount, upgradedCount, upgradesLeftCount, deletedCount int)
	progressMutex       sync.RWMutex
	progressArgsForCall []struct {
		pollingInterval   time.Duration
		orphanCount       int
		upgradedCount     int
		upgradesLeftCount int
		deletedCount      int
	}
	FinishedStub        func(orphanCount, upgradedCount, deletedCount, couldNotStartCount int, failedInstances ...string)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		orphanCount        int
		upgradedCount      int
		deletedCount       int
		couldNotStartCount int
		failedInstances    []string
	}
	CanariesStartingStub        func(canaries int)
	canariesStartingMutex       sync.RWMutex
	canariesStartingArgsForCall []struct {
		canaries int
	}
	CanariesFinishedStub        func()
	canariesFinishedMutex       sync.RWMutex
	canariesFinishedArgsForCall []struct{}
	invocations                 map[string][][]interface{}
	invocationsMutex            sync.RWMutex
}

func (fake *FakeListener) FailedToRefreshInstanceInfo(instance string) {
	fake.failedToRefreshInstanceInfoMutex.Lock()
	fake.failedToRefreshInstanceInfoArgsForCall = append(fake.failedToRefreshInstanceInfoArgsForCall, struct {
		instance string
	}{instance})
	fake.recordInvocation("FailedToRefreshInstanceInfo", []interface{}{instance})
	fake.failedToRefreshInstanceInfoMutex.Unlock()
	if fake.FailedToRefreshInstanceInfoStub != nil {
		fake.FailedToRefreshInstanceInfoStub(instance)
	}
}

func (fake *FakeListener) FailedToRefreshInstanceInfoCallCount() int {
	fake.failedToRefreshInstanceInfoMutex.RLock()
	defer fake.failedToRefreshInstanceInfoMutex.RUnlock()
	return len(fake.failedToRefreshInstanceInfoArgsForCall)
}

func (fake *FakeListener) FailedToRefreshInstanceInfoArgsForCall(i int) string {
	fake.failedToRefreshInstanceInfoMutex.RLock()
	defer fake.failedToRefreshInstanceInfoMutex.RUnlock()
	return fake.failedToRefreshInstanceInfoArgsForCall[i].instance
}

func (fake *FakeListener) Starting(maxInFlight int) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		maxInFlight int
	}{maxInFlight})
	fake.recordInvocation("Starting", []interface{}{maxInFlight})
	fake.startingMutex.Unlock()
	if fake.StartingStub != nil {
		fake.StartingStub(maxInFlight)
	}
}

func (fake *FakeListener) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeListener) StartingArgsForCall(i int) int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return fake.startingArgsForCall[i].maxInFlight
}

func (fake *FakeListener) RetryAttempt(num int, limit int) {
	fake.retryAttemptMutex.Lock()
	fake.retryAttemptArgsForCall = append(fake.retryAttemptArgsForCall, struct {
		num   int
		limit int
	}{num, limit})
	fake.recordInvocation("RetryAttempt", []interface{}{num, limit})
	fake.retryAttemptMutex.Unlock()
	if fake.RetryAttemptStub != nil {
		fake.RetryAttemptStub(num, limit)
	}
}

func (fake *FakeListener) RetryAttemptCallCount() int {
	fake.retryAttemptMutex.RLock()
	defer fake.retryAttemptMutex.RUnlock()
	return len(fake.retryAttemptArgsForCall)
}

func (fake *FakeListener) RetryAttemptArgsForCall(i int) (int, int) {
	fake.retryAttemptMutex.RLock()
	defer fake.retryAttemptMutex.RUnlock()
	return fake.retryAttemptArgsForCall[i].num, fake.retryAttemptArgsForCall[i].limit
}

func (fake *FakeListener) RetryCanariesAttempt(num int, limit int, remainingCanaries int) {
	fake.retryCanariesAttemptMutex.Lock()
	fake.retryCanariesAttemptArgsForCall = append(fake.retryCanariesAttemptArgsForCall, struct {
		num               int
		limit             int
		remainingCanaries int
	}{num, limit, remainingCanaries})
	fake.recordInvocation("RetryCanariesAttempt", []interface{}{num, limit, remainingCanaries})
	fake.retryCanariesAttemptMutex.Unlock()
	if fake.RetryCanariesAttemptStub != nil {
		fake.RetryCanariesAttemptStub(num, limit, remainingCanaries)
	}
}

func (fake *FakeListener) RetryCanariesAttemptCallCount() int {
	fake.retryCanariesAttemptMutex.RLock()
	defer fake.retryCanariesAttemptMutex.RUnlock()
	return len(fake.retryCanariesAttemptArgsForCall)
}

func (fake *FakeListener) RetryCanariesAttemptArgsForCall(i int) (int, int, int) {
	fake.retryCanariesAttemptMutex.RLock()
	defer fake.retryCanariesAttemptMutex.RUnlock()
	return fake.retryCanariesAttemptArgsForCall[i].num, fake.retryCanariesAttemptArgsForCall[i].limit, fake.retryCanariesAttemptArgsForCall[i].remainingCanaries
}

func (fake *FakeListener) InstancesToUpgrade(instances []service.Instance) {
	var instancesCopy []service.Instance
	if instances != nil {
		instancesCopy = make([]service.Instance, len(instances))
		copy(instancesCopy, instances)
	}
	fake.instancesToUpgradeMutex.Lock()
	fake.instancesToUpgradeArgsForCall = append(fake.instancesToUpgradeArgsForCall, struct {
		instances []service.Instance
	}{instancesCopy})
	fake.recordInvocation("InstancesToUpgrade", []interface{}{instancesCopy})
	fake.instancesToUpgradeMutex.Unlock()
	if fake.InstancesToUpgradeStub != nil {
		fake.InstancesToUpgradeStub(instances)
	}
}

func (fake *FakeListener) InstancesToUpgradeCallCount() int {
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	return len(fake.instancesToUpgradeArgsForCall)
}

func (fake *FakeListener) InstancesToUpgradeArgsForCall(i int) []service.Instance {
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	return fake.instancesToUpgradeArgsForCall[i].instances
}

func (fake *FakeListener) InstanceUpgradeStarting(instance string, index int, totalInstances int, isCanary bool) {
	fake.instanceUpgradeStartingMutex.Lock()
	fake.instanceUpgradeStartingArgsForCall = append(fake.instanceUpgradeStartingArgsForCall, struct {
		instance       string
		index          int
		totalInstances int
		isCanary       bool
	}{instance, index, totalInstances, isCanary})
	fake.recordInvocation("InstanceUpgradeStarting", []interface{}{instance, index, totalInstances, isCanary})
	fake.instanceUpgradeStartingMutex.Unlock()
	if fake.InstanceUpgradeStartingStub != nil {
		fake.InstanceUpgradeStartingStub(instance, index, totalInstances, isCanary)
	}
}

func (fake *FakeListener) InstanceUpgradeStartingCallCount() int {
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	return len(fake.instanceUpgradeStartingArgsForCall)
}

func (fake *FakeListener) InstanceUpgradeStartingArgsForCall(i int) (string, int, int, bool) {
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	return fake.instanceUpgradeStartingArgsForCall[i].instance, fake.instanceUpgradeStartingArgsForCall[i].index, fake.instanceUpgradeStartingArgsForCall[i].totalInstances, fake.instanceUpgradeStartingArgsForCall[i].isCanary
}

func (fake *FakeListener) InstanceUpgradeStartResult(instance string, status services.UpgradeOperationType) {
	fake.instanceUpgradeStartResultMutex.Lock()
	fake.instanceUpgradeStartResultArgsForCall = append(fake.instanceUpgradeStartResultArgsForCall, struct {
		instance string
		status   services.UpgradeOperationType
	}{instance, status})
	fake.recordInvocation("InstanceUpgradeStartResult", []interface{}{instance, status})
	fake.instanceUpgradeStartResultMutex.Unlock()
	if fake.InstanceUpgradeStartResultStub != nil {
		fake.InstanceUpgradeStartResultStub(instance, status)
	}
}

func (fake *FakeListener) InstanceUpgradeStartResultCallCount() int {
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	return len(fake.instanceUpgradeStartResultArgsForCall)
}

func (fake *FakeListener) InstanceUpgradeStartResultArgsForCall(i int) (string, services.UpgradeOperationType) {
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	return fake.instanceUpgradeStartResultArgsForCall[i].instance, fake.instanceUpgradeStartResultArgsForCall[i].status
}

func (fake *FakeListener) InstanceUpgraded(instance string, result string) {
	fake.instanceUpgradedMutex.Lock()
	fake.instanceUpgradedArgsForCall = append(fake.instanceUpgradedArgsForCall, struct {
		instance string
		result   string
	}{instance, result})
	fake.recordInvocation("InstanceUpgraded", []interface{}{instance, result})
	fake.instanceUpgradedMutex.Unlock()
	if fake.InstanceUpgradedStub != nil {
		fake.InstanceUpgradedStub(instance, result)
	}
}

func (fake *FakeListener) InstanceUpgradedCallCount() int {
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	return len(fake.instanceUpgradedArgsForCall)
}

func (fake *FakeListener) InstanceUpgradedArgsForCall(i int) (string, string) {
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	return fake.instanceUpgradedArgsForCall[i].instance, fake.instanceUpgradedArgsForCall[i].result
}

func (fake *FakeListener) WaitingFor(instance string, boshTaskId int) {
	fake.waitingForMutex.Lock()
	fake.waitingForArgsForCall = append(fake.waitingForArgsForCall, struct {
		instance   string
		boshTaskId int
	}{instance, boshTaskId})
	fake.recordInvocation("WaitingFor", []interface{}{instance, boshTaskId})
	fake.waitingForMutex.Unlock()
	if fake.WaitingForStub != nil {
		fake.WaitingForStub(instance, boshTaskId)
	}
}

func (fake *FakeListener) WaitingForCallCount() int {
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	return len(fake.waitingForArgsForCall)
}

func (fake *FakeListener) WaitingForArgsForCall(i int) (string, int) {
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	return fake.waitingForArgsForCall[i].instance, fake.waitingForArgsForCall[i].boshTaskId
}

func (fake *FakeListener) Progress(pollingInterval time.Duration, orphanCount int, upgradedCount int, upgradesLeftCount int, deletedCount int) {
	fake.progressMutex.Lock()
	fake.progressArgsForCall = append(fake.progressArgsForCall, struct {
		pollingInterval   time.Duration
		orphanCount       int
		upgradedCount     int
		upgradesLeftCount int
		deletedCount      int
	}{pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount})
	fake.recordInvocation("Progress", []interface{}{pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount})
	fake.progressMutex.Unlock()
	if fake.ProgressStub != nil {
		fake.ProgressStub(pollingInterval, orphanCount, upgradedCount, upgradesLeftCount, deletedCount)
	}
}

func (fake *FakeListener) ProgressCallCount() int {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	return len(fake.progressArgsForCall)
}

func (fake *FakeListener) ProgressArgsForCall(i int) (time.Duration, int, int, int, int) {
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	return fake.progressArgsForCall[i].pollingInterval, fake.progressArgsForCall[i].orphanCount, fake.progressArgsForCall[i].upgradedCount, fake.progressArgsForCall[i].upgradesLeftCount, fake.progressArgsForCall[i].deletedCount
}

func (fake *FakeListener) Finished(orphanCount int, upgradedCount int, deletedCount int, couldNotStartCount int, failedInstances ...string) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		orphanCount        int
		upgradedCount      int
		deletedCount       int
		couldNotStartCount int
		failedInstances    []string
	}{orphanCount, upgradedCount, deletedCount, couldNotStartCount, failedInstances})
	fake.recordInvocation("Finished", []interface{}{orphanCount, upgradedCount, deletedCount, couldNotStartCount, failedInstances})
	fake.finishedMutex.Unlock()
	if fake.FinishedStub != nil {
		fake.FinishedStub(orphanCount, upgradedCount, deletedCount, couldNotStartCount, failedInstances...)
	}
}

func (fake *FakeListener) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeListener) FinishedArgsForCall(i int) (int, int, int, int, []string) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return fake.finishedArgsForCall[i].orphanCount, fake.finishedArgsForCall[i].upgradedCount, fake.finishedArgsForCall[i].deletedCount, fake.finishedArgsForCall[i].couldNotStartCount, fake.finishedArgsForCall[i].failedInstances
}

func (fake *FakeListener) CanariesStarting(canaries int) {
	fake.canariesStartingMutex.Lock()
	fake.canariesStartingArgsForCall = append(fake.canariesStartingArgsForCall, struct {
		canaries int
	}{canaries})
	fake.recordInvocation("CanariesStarting", []interface{}{canaries})
	fake.canariesStartingMutex.Unlock()
	if fake.CanariesStartingStub != nil {
		fake.CanariesStartingStub(canaries)
	}
}

func (fake *FakeListener) CanariesStartingCallCount() int {
	fake.canariesStartingMutex.RLock()
	defer fake.canariesStartingMutex.RUnlock()
	return len(fake.canariesStartingArgsForCall)
}

func (fake *FakeListener) CanariesStartingArgsForCall(i int) int {
	fake.canariesStartingMutex.RLock()
	defer fake.canariesStartingMutex.RUnlock()
	return fake.canariesStartingArgsForCall[i].canaries
}

func (fake *FakeListener) CanariesFinished() {
	fake.canariesFinishedMutex.Lock()
	fake.canariesFinishedArgsForCall = append(fake.canariesFinishedArgsForCall, struct{}{})
	fake.recordInvocation("CanariesFinished", []interface{}{})
	fake.canariesFinishedMutex.Unlock()
	if fake.CanariesFinishedStub != nil {
		fake.CanariesFinishedStub()
	}
}

func (fake *FakeListener) CanariesFinishedCallCount() int {
	fake.canariesFinishedMutex.RLock()
	defer fake.canariesFinishedMutex.RUnlock()
	return len(fake.canariesFinishedArgsForCall)
}

func (fake *FakeListener) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.failedToRefreshInstanceInfoMutex.RLock()
	defer fake.failedToRefreshInstanceInfoMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.retryAttemptMutex.RLock()
	defer fake.retryAttemptMutex.RUnlock()
	fake.retryCanariesAttemptMutex.RLock()
	defer fake.retryCanariesAttemptMutex.RUnlock()
	fake.instancesToUpgradeMutex.RLock()
	defer fake.instancesToUpgradeMutex.RUnlock()
	fake.instanceUpgradeStartingMutex.RLock()
	defer fake.instanceUpgradeStartingMutex.RUnlock()
	fake.instanceUpgradeStartResultMutex.RLock()
	defer fake.instanceUpgradeStartResultMutex.RUnlock()
	fake.instanceUpgradedMutex.RLock()
	defer fake.instanceUpgradedMutex.RUnlock()
	fake.waitingForMutex.RLock()
	defer fake.waitingForMutex.RUnlock()
	fake.progressMutex.RLock()
	defer fake.progressMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.canariesStartingMutex.RLock()
	defer fake.canariesStartingMutex.RUnlock()
	fake.canariesFinishedMutex.RLock()
	defer fake.canariesFinishedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeListener) recordInvocation(key string, args []interface{}) {
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

var _ upgrader.Listener = new(FakeListener)
