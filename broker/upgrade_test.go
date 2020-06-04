// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package broker_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	sdk "github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi/v7/domain"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	brokerfakes "github.com/pivotal-cf/on-demand-service-broker/broker/fakes"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-service-broker/serviceadapter"
)

var _ = Describe("Upgrade", func() {
	var (
		upgradeOperationData broker.OperationData
		details              domain.UpdateDetails
		instanceID           string
		logger               *log.Logger
		boshTaskID           int
		redeployErr          error
		fakeUAAClient        *brokerfakes.FakeUAAClient
		expectedClient       map[string]string
	)

	BeforeEach(func() {
		instanceID = "some-instance"
		boshTaskID = 876
		arbContext := map[string]interface{}{
			"instance_name": "some-instance-name",
		}
		serialisedArbitraryContext, _ := json.Marshal(arbContext)
		details = domain.UpdateDetails{
			PlanID:     existingPlanID,
			RawContext: serialisedArbitraryContext,
		}
		logger = loggerFactory.NewWithRequestID()
		b = createDefaultBroker()
		fakeDeployer.UpgradeReturns(boshTaskID, []byte("new-manifest-fetched-from-adapter"), nil)

		fakeUAAClient = new(brokerfakes.FakeUAAClient)

		expectedClient = map[string]string{
			"client_secret": "some-secret",
			"client_id":     "some-id",
			"foo":           "bar",
		}
		fakeUAAClient.UpdateClientReturns(expectedClient, nil)
		b.SetUAAClient(fakeUAAClient)
	})

	It("when the deployment goes well deploys with the new planID", func() {
		upgradeOperationData, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

		Expect(redeployErr).NotTo(HaveOccurred())
		Expect(fakeDeployer.CreateCallCount()).To(Equal(0))
		Expect(fakeDeployer.UpgradeCallCount()).To(Equal(1))
		Expect(fakeDeployer.UpdateCallCount()).To(Equal(0))
		actualDeploymentName, actualPlan, actualBoshContextID, _, _ := fakeDeployer.UpgradeArgsForCall(0)
		Expect(actualPlan).To(Equal(existingPlan))
		Expect(actualDeploymentName).To(Equal(broker.InstancePrefix + instanceID))
		Expect(actualBoshContextID).To(BeEmpty())
	})

	Context("when instance is already up to date", func() {
		It("should return error", func() {
			expectedError := broker.NewOperationAlreadyCompletedError(errors.New("instance is already up to date"))
			fakeDeployer.UpgradeReturns(0, nil, expectedError)

			_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

			Expect(redeployErr).To(Equal(expectedError))
		})
	})

	Context("when there is a previous deployment for the service instance", func() {
		It("responds with the correct upgradeOperationData", func() {
			upgradeOperationData, _ = b.Upgrade(context.Background(), instanceID, details, logger)

			Expect(upgradeOperationData).To(Equal(
				broker.OperationData{
					BoshTaskID:    boshTaskID,
					OperationType: broker.OperationTypeUpgrade,
				},
			))

			Expect(logBuffer.String()).To(MatchRegexp(`\[[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}\] \d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2} upgrading instance`))
		})

		It("and post-deploy errand is configured deploys with a context id", func() {
			details = domain.UpdateDetails{
				PlanID: postDeployErrandPlanID,
			}

			upgradeOperationData, _ = b.Upgrade(context.Background(), instanceID, details, logger)

			_, _, contextID, _, _ := fakeDeployer.UpgradeArgsForCall(0)
			Expect(contextID).NotTo(BeEmpty())
			Expect(upgradeOperationData.BoshContextID).NotTo(BeEmpty())
			Expect(upgradeOperationData).To(Equal(
				broker.OperationData{
					BoshTaskID:    boshTaskID,
					OperationType: broker.OperationTypeUpgrade,
					BoshContextID: upgradeOperationData.BoshContextID,
					Errands: []config.Errand{{
						Name:      "health-check",
						Instances: []string{"redis-server/0"},
					}},
				},
			))
		})

		It("and the service adapter returns a UnknownFailureError with a user message returns the error for the user", func() {
			err := serviceadapter.NewUnknownFailureError("error for cf user")
			fakeDeployer.UpgradeReturns(boshTaskID, nil, err)
			_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

			Expect(redeployErr).To(Equal(err))
		})

		It("and the service adapter returns a UnknownFailureError with no message returns a generic error", func() {
			fakeDeployer.UpgradeReturns(boshTaskID, nil, serviceadapter.NewUnknownFailureError(""))
			_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

			Expect(redeployErr).To(MatchError(ContainSubstring("There was a problem completing your request. Please contact your operations team providing the following information")))
		})
	})

	It("when no update details are provided returns an error", func() {
		details = domain.UpdateDetails{}
		_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

		Expect(redeployErr).To(MatchError(ContainSubstring("no plan ID provided in upgrade request body")))
	})

	It("when the plan cannot be found upgrade fails and does not redeploy", func() {
		planID := "plan-id-doesnt-exist"

		details = domain.UpdateDetails{
			PlanID: planID,
		}
		_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

		Expect(redeployErr).To(MatchError(ContainSubstring(fmt.Sprintf("plan %s not found", planID))))
		Expect(logBuffer.String()).To(ContainSubstring(fmt.Sprintf("error: finding plan ID %s", planID)))
		Expect(fakeDeployer.UpgradeCallCount()).To(BeZero())
	})

	It("when there is a task in progress on the instance upgrade returns an OperationInProgressError", func() {
		fakeDeployer.UpgradeReturns(0, nil, broker.TaskInProgressError{})
		_, redeployErr = b.Upgrade(context.Background(), instanceID, details, logger)

		Expect(redeployErr).To(BeAssignableToTypeOf(broker.OperationInProgressError{}))
	})

	It("should not request the json schemas from the service adapter", func() {
		fakeAdapter := new(brokerfakes.FakeServiceAdapterClient)
		fakeAdapter.GeneratePlanSchemaReturns(domain.ServiceSchemas{}, fmt.Errorf("derp!"))
		broker := createBrokerWithAdapter(fakeAdapter)

		_, upgradeErr := broker.Upgrade(context.Background(), instanceID, details, logger)

		Expect(fakeAdapter.GeneratePlanSchemaCallCount()).To(Equal(0))
		Expect(upgradeErr).NotTo(HaveOccurred())
	})

	Context("the service instance UAA client", func() {
		var newlyGeneratedManifest []byte
		var existingClient map[string]string

		BeforeEach(func() {
			dashboardURL := "http://example.com/dashboard"
			serviceAdapter.GenerateDashboardUrlReturns(dashboardURL, nil)
			newlyGeneratedManifest = []byte("name: new-name")
			fakeDeployer.UpgradeReturns(boshTaskID, newlyGeneratedManifest, nil)
			existingClient = map[string]string{
				"client_id":    "some-id",
				"redirect_uri": "http://uri.com/example",
			}
			fakeUAAClient.GetClientReturns(existingClient, nil)
			fakeUAAClient.HasClientDefinitionReturns(true)
		})

		It("passes the client to the deployer", func() {
			_, upgradeError := b.Upgrade(context.Background(), instanceID, details, logger)
			Expect(upgradeError).NotTo(HaveOccurred())

			Expect(fakeDeployer.UpgradeCallCount()).To(Equal(1))
			_, _, _, actualClient, _ := fakeDeployer.UpgradeArgsForCall(0)
			Expect(actualClient).To(Equal(existingClient))
		})

		It("updates the service instance client", func() {
			_, upgradeError := b.Upgrade(context.Background(), instanceID, details, logger)
			Expect(upgradeError).NotTo(HaveOccurred())

			By("regenerating the dashboard url", func() {
				Expect(serviceAdapter.GenerateDashboardUrlCallCount()).To(Equal(1))
				instanceID, plan, boshManifest, _ := serviceAdapter.GenerateDashboardUrlArgsForCall(0)
				Expect(instanceID).To(Equal(instanceID))
				expectedProperties := sdk.Properties{
					"some_other_global_property": "other_global_value",
					"a_global_property":          "global_value",
					"super":                      "no",
				}
				Expect(plan).To(Equal(sdk.Plan{
					Properties:     expectedProperties,
					InstanceGroups: existingPlan.InstanceGroups,
					Update:         existingPlan.Update,
				}))
				Expect(boshManifest).To(Equal(newlyGeneratedManifest))
			})

			Expect(fakeUAAClient.UpdateClientCallCount()).To(Equal(1))
			actualClientID, actualRedirectURI := fakeUAAClient.UpdateClientArgsForCall(0)

			Expect(actualClientID).To(Equal(instanceID))
			Expect(actualRedirectURI).To(Equal("http://example.com/dashboard"))
		})

		When("updating the uaa client fails", func() {
			It("returns a generic error message", func() {
				fakeUAAClient.GetClientReturns(map[string]string{"client_id": "1"}, nil)
				fakeUAAClient.UpdateClientReturns(nil, errors.New("oh no"))

				_, upgradeError := b.Upgrade(context.Background(), instanceID, details, logger)

				Expect(upgradeError).To(HaveOccurred())
				Expect(upgradeError).To(MatchError(ContainSubstring(
					"There was a problem completing your request. Please contact your operations team providing the following information:",
				)))
			})
		})

		When("getting the uaa client fails", func() {
			It("returns a generic error message", func() {
				fakeUAAClient.GetClientReturns(nil, errors.New("oh no"))

				_, upgradeError := b.Upgrade(context.Background(), instanceID, details, logger)

				Expect(upgradeError).To(HaveOccurred())
				Expect(upgradeError).To(MatchError(ContainSubstring(
					"There was a problem completing your request. Please contact your operations team providing the following information:",
				)))
			})
		})

		When("generating the dashboard fails", func() {
			It("returns a generic error", func() {
				fakeUAAClient.GetClientReturns(map[string]string{"a": "b"}, nil)
				serviceAdapter.GenerateDashboardUrlReturns("", errors.New("fooo"))
				_, upgradeError := b.Upgrade(context.Background(), instanceID, details, logger)

				Expect(upgradeError).To(MatchError(
					ContainSubstring("There was a problem completing your request"),
				))
			})
		})
	})
})
