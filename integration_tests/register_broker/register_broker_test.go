package register_broker_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/pivotal-cf/on-demand-service-broker/config"
	"gopkg.in/yaml.v2"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-cf/on-demand-service-broker/integration_tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("RegisterBroker", func() {
	var (
		cfServer                                              *ghttp.Server
		serviceBrokersHandler                                 *helpers.FakeHandler
		createBrokerHandler                                   *helpers.FakeHandler
		updateBrokerHandler                                   *helpers.FakeHandler
		servicesHandler                                       *helpers.FakeHandler
		servicePlansHandler                                   *helpers.FakeHandler
		servicePlansOfferingHandler                           *helpers.FakeHandler
		servicePlansVisibilityHandler                         *helpers.FakeHandler
		brokerName, brokerUsername, brokerPassword, brokerURL string
		serviceID, serviceGUID, planName                      string
		errandConfig                                          config.RegisterBrokerErrandConfig
	)

	BeforeEach(func() {
		brokerName = "some-service-broker"
		brokerURL = "http://example.broker.com"
		brokerUsername = "username"
		brokerPassword = "password"
		serviceGUID = "service-guid"

		serviceID = "a-service"
		planName = "a-plan"

		cfServer = ghttp.NewServer()

		serviceBrokersHandler = new(helpers.FakeHandler)
		createBrokerHandler = new(helpers.FakeHandler)
		updateBrokerHandler = new(helpers.FakeHandler)
		servicesHandler = new(helpers.FakeHandler)
		servicePlansHandler = new(helpers.FakeHandler)
		servicePlansOfferingHandler = new(helpers.FakeHandler)
		servicePlansVisibilityHandler = new(helpers.FakeHandler)

		cfServer.RouteToHandler(http.MethodPost, "/oauth/token", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte(`{"access_token":"authtoken"}`))
		})
		cfServer.RouteToHandler(http.MethodGet, "/v2/service_brokers", serviceBrokersHandler.Handle)
		cfServer.RouteToHandler(http.MethodPost, "/v2/service_brokers", createBrokerHandler.Handle)
		cfServer.RouteToHandler(http.MethodPut, regexp.MustCompile(`/v2/service_brokers/.*`), updateBrokerHandler.Handle)
		cfServer.RouteToHandler(http.MethodGet, "/v2/services", servicesHandler.Handle)
		cfServer.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/services/.*/service_plans`), servicePlansOfferingHandler.Handle)
		cfServer.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/service_plan_visibilities?q=service_plan_guid:.*`), servicePlansVisibilityHandler.Handle)
		cfServer.RouteToHandler(http.MethodPut, regexp.MustCompile(`/v2/service_plans/.*`), servicePlansHandler.Handle)

		errandConfig = config.RegisterBrokerErrandConfig{
			BrokerName:     brokerName,
			BrokerUsername: brokerUsername,
			BrokerPassword: brokerPassword,
			BrokerURL:      brokerURL,
			CF: config.CF{
				URL: cfServer.URL(),
				Authentication: config.Authentication{
					UAA: config.UAAAuthentication{
						URL:             cfServer.URL(),
						UserCredentials: config.UserCredentials{Username: "foo", Password: "bar"},
					},
				},
			},
			DisableSSLCertVerification: true,
			ServicePlans: config.ServicePlans{
				ServiceID: serviceID,
				CFPlanAccesses: []config.CFPlanAccess{
					{
						Name:       planName,
						AccessType: config.Enable,
					},
				},
			},
		}
	})

	It("creates a service broker when the broker is not yet registered", func() {
		serviceBrokersHandler.RespondsWith(http.StatusOK, `{"resources":[]}`)
		createBrokerHandler.RespondsWith(http.StatusCreated, "")
		servicePlansHandler.RespondsWith(http.StatusCreated, ``)

		executeBinarySuccessfully(errandConfig, GinkgoWriter, GinkgoWriter)

		Expect(updateBrokerHandler.RequestsReceived()).To(BeZero(), "update broker was called")
		Expect(createBrokerHandler.RequestsReceived()).To(BeNumerically(">", 0), "no request was made to create broker")
		createRequest := createBrokerHandler.GetRequestForCall(0)
		Expect(createRequest.Body).To(MatchJSON(fmt.Sprintf(`{
				"name": "%s", 
				"broker_url": "%s",
				"auth_username": "%s",
				"auth_password": "%s"
			}`, brokerName, brokerURL, brokerUsername, brokerPassword)))
	})

	It("updates the existing broker when the broker is already registered", func() {
		cfBrokerResponse := fmt.Sprintf(`{
			"resources": [{
				"entity": {"name": "%s"},
				"metadata": {"guid": "%s"}
			}]
		}`, brokerName, serviceGUID)
		serviceBrokersHandler.RespondsWith(http.StatusOK, cfBrokerResponse)
		updateBrokerHandler.RespondsWith(http.StatusOK, "")
		servicePlansHandler.RespondsWith(http.StatusCreated, ``)

		executeBinarySuccessfully(errandConfig, GinkgoWriter, GinkgoWriter)

		Expect(createBrokerHandler.RequestsReceived()).To(BeZero(), "create broker was called")
		Expect(updateBrokerHandler.RequestsReceived()).To(BeNumerically(">", 0), "no request was made to update broker")

		updateRequest := updateBrokerHandler.GetRequestForCall(0)
		Expect(updateRequest.Body).To(MatchJSON(fmt.Sprintf(`{
				"name": "%s", 
				"broker_url": "%s",
				"auth_username": "%s",
				"auth_password": "%s"
			}`, brokerName, brokerURL, brokerUsername, brokerPassword)))
		Expect(updateRequest.URL).To(Equal("/v2/service_brokers/" + serviceGUID))
	})

	It("enables access to all plans accordingly to the specification", func() {
		serviceBrokersHandler.RespondsWith(http.StatusOK, `{"resources":[]}`)
		createBrokerHandler.RespondsWith(http.StatusCreated, "")

		servicesHandler.RespondsWith(http.StatusOK, fmt.Sprintf(`{
		  "resources": [
			{
			  "entity": {
				"unique_id": %q,
				"service_plans_url": "/v2/services/%s/service_plans"
			  }
			}
		  ]
		}`, serviceID, serviceGUID))

		planGUID := "unique-uid"
		servicePlansOfferingHandler.RespondsWith(http.StatusOK, fmt.Sprintf(`{
		  "resources" : [
			{
			  "entity": {
				"name": %q
			  },
			  "metadata": {
				"guid": %q
			  }
			}
		  ]
		}`, planName, planGUID))

		servicePlansVisibilityHandler.RespondsWith(http.StatusOK, `{"resources": null}`) //TODO test not null in this test
		servicePlansHandler.RespondsWith(http.StatusCreated, ``)
		executeBinarySuccessfully(errandConfig, GinkgoWriter, GinkgoWriter)

		Expect(servicePlansHandler.RequestsReceived()).To(BeNumerically(">", 0), "no request was made to service access")
		servicePlansRequest := servicePlansHandler.GetRequestForCall(0)
		Expect(servicePlansRequest.URL).To(Equal("/v2/service_plans/" + planGUID))
		Expect(servicePlansRequest.Body).To(MatchJSON(`{
				"public": true
			}`))
	})

	Describe("error handling", func() {
		It("fails when config path is not specified", func() {
			cmd := exec.Command(binaryPath)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1), "succeeded unexpectedly")
			Expect(session).To(gbytes.Say("-configPath must be given as argument"))
		})

		It("fails when config path is not a file", func() {
			cmd := exec.Command(binaryPath, "-configPath", "not-a-file")

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1), "succeeded unexpectedly")
			Expect(session).To(gbytes.Say("error reading file -configPath"))
		})

		It("fails when running the errand fails", func() {
			serviceBrokersHandler.RespondsWith(http.StatusInternalServerError, "")

			session := executeBinary(errandConfig, GinkgoWriter, GinkgoWriter)
			Expect(session).To(gexec.Exit(1))
		})
	})

})

func executeBinarySuccessfully(errandConfig config.RegisterBrokerErrandConfig, stdout, stderr io.Writer) *gexec.Session {
	session := executeBinary(errandConfig, stdout, stderr)
	Expect(session).To(gexec.Exit(0))

	return session
}

func executeBinary(errandConfig config.RegisterBrokerErrandConfig, stdout, stderr io.Writer) *gexec.Session {
	errandConfigPath, err := ioutil.TempFile("/tmp", "")
	Expect(err).ToNot(HaveOccurred())

	b, err := yaml.Marshal(errandConfig)
	Expect(err).ToNot(HaveOccurred())

	_, err = errandConfigPath.Write(b)
	Expect(err).ToNot(HaveOccurred())

	cmd := exec.Command(binaryPath, "-configPath", errandConfigPath.Name())

	session, err := gexec.Start(cmd, stdout, stderr)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session, 5*time.Second).Should(gexec.Exit())
	Expect(os.RemoveAll(errandConfigPath.Name())).To(Succeed())

	return session
}

/*
* system tests:  test the release
* integration:   mock the http endpoints, compile and execute the binary
* collaboration: create client with fake structs
* contract:      tests the real clients (CF, Bosh...)
* unit tests:    test functions
* */

/*
* verifies:
*		- when no argument is passed, it fails
*		- when configPath is passed
*			- but it's not a file, it fails
*			- but it's not a yaml, it fails
*			- but it doesn't contain required fields, it fails (MAYBE release level only?)
*			- mocked client would receive method call with the correct config
*
* System level:
*		- all the pieces work together on the happy case (existing test)
*			- guarantee? that we call the right cmd
*
* Integration:
*		- when no argument is passed, it fails
*		- when configPath is not passed
*		- happy case for register broker
*			- when {
				Broker doesn't exist
					Calls the Create Broker endpoint
				Broker exists
					Calls the update broker endpoint
*
* Unit for registrar:
*   - Test the runner. Not compiling the binary
*			- but it's not a file, it fails
* */
