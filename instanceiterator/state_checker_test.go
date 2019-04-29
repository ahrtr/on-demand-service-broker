// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package instanceiterator_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi/domain"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/instanceiterator"
	"github.com/pivotal-cf/on-demand-service-broker/instanceiterator/fakes"
)

var _ = Describe("State checker", func() {
	var (
		guid                  string
		expectedOperationData broker.OperationData
		fakeBrokerService     *fakes.FakeBrokerServices
		stateChecker          instanceiterator.StateChecker
	)

	BeforeEach(func() {
		guid = "some-guid"
		expectedOperationData = broker.OperationData{BoshTaskID: 123}
		fakeBrokerService = new(fakes.FakeBrokerServices)
		stateChecker = instanceiterator.NewStateChecker(fakeBrokerService)
	})

	It("returns OperationSucceeded when last operation reports success", func() {
		fakeBrokerService.LastOperationReturns(domain.LastOperation{State: domain.Succeeded}, nil)

		state, err := stateChecker.Check(guid, expectedOperationData)
		Expect(err).NotTo(HaveOccurred())

		By("pulling the last operation with the right arguments")
		Expect(fakeBrokerService.LastOperationCallCount()).To(Equal(1))
		guidArg, operationData := fakeBrokerService.LastOperationArgsForCall(0)
		Expect(guidArg).To(Equal(guid))
		Expect(operationData).To(Equal(expectedOperationData))

		Expect(state).To(Equal(services.BOSHOperation{Type: services.OperationSucceeded, Data: expectedOperationData}))
	})

	It("returns an error if it fails to pull last operation", func() {
		fakeBrokerService.LastOperationReturns(domain.LastOperation{}, errors.New("oops"))

		_, err := stateChecker.Check(guid, expectedOperationData)
		Expect(err).To(MatchError("error getting last operation: oops"))
	})

	It("returns OperationFailed when last operation reports failure", func() {
		fakeBrokerService.LastOperationReturns(domain.LastOperation{State: domain.Failed}, nil)

		state, err := stateChecker.Check(guid, expectedOperationData)
		Expect(err).NotTo(HaveOccurred())

		Expect(state).To(Equal(services.BOSHOperation{Type: services.OperationFailed, Data: expectedOperationData}))
	})

	It("returns OperationAccepted when last operation reports the operation is in progress", func() {
		fakeBrokerService.LastOperationReturns(domain.LastOperation{State: domain.InProgress}, nil)

		state, err := stateChecker.Check(guid, expectedOperationData)
		Expect(err).NotTo(HaveOccurred())

		Expect(state).To(Equal(services.BOSHOperation{Type: services.OperationAccepted, Data: expectedOperationData}))
	})

	It("returns an error if last operation returns an unknown state", func() {
		fakeBrokerService.LastOperationReturns(domain.LastOperation{State: "not-a-state"}, nil)

		_, err := stateChecker.Check(guid, expectedOperationData)
		Expect(err).To(MatchError("unknown state from last operation: not-a-state"))
	})
})
