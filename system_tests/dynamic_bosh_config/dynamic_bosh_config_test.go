package dynamic_bosh_config_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"

	cf "github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/cf_helpers"
)

var _ = Describe("DynamicBoshConfig", func() {

	var (
		serviceId string
	)

	BeforeEach(func() {
		serviceId = ""
	})

	It("handles bosh configs during the lifecycle of a service instance", func() {

		serviceInstanceName = "service" + brokerInfo.TestSuffix
		boshConfig := fmt.Sprintf(`{"vm_extensions_config": "vm_extensions: [{name: vm-ext%s}]"}`, brokerInfo.TestSuffix)

		cf.CreateService(brokerInfo.ServiceOffering, "redis-with-post-deploy", serviceInstanceName, boshConfig)

		serviceId = "service-instance_" + cf.GetServiceInstanceGUID(serviceInstanceName)
		configDetails, err := bosh_helpers.GetBOSHConfig("cloud", serviceId)
		Expect(err).NotTo(HaveOccurred())

		Expect(configDetails).To(ContainSubstring("name: vm-ext" + brokerInfo.TestSuffix))

		// update and check bosh configs

		cf.DeleteService(serviceInstanceName)

		_, err = bosh_helpers.GetBOSHConfig("cloud", serviceId)
		Expect(err).To(HaveOccurred(), "cloud config wasn't deleted during DeleteService")
	})

	AfterEach(func() {
		if serviceId != "" {
			bosh_helpers.DeleteBOSHConfig("cloud", serviceId)
		}
	})
})
