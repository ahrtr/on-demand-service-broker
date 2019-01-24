// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package upgrade_deployment_tests

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pborman/uuid"
	cf "github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/cf_helpers"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Upgrading deployment", func() {
	var (
		serviceInstanceName string
		testAppName         string
		testAppURL          string
	)

	BeforeEach(func() {
		By("creating a service instance")
		serviceInstanceName = fmt.Sprintf("my-service-%s", uuid.New()[:7])
		cf.CreateService(serviceOffering, plan, serviceInstanceName, "")

		By("pushing an app and binding to it")
		testAppName = uuid.New()[:7]
		testAppURL = cf.PushAndBindApp(testAppName, serviceInstanceName, path.Join(ciRootPath, exampleAppDirName))

		By("exercising the service instance")
		cf.PutToTestApp(testAppURL, "foo", "bar")
	})

	AfterEach(func() {
		By("deleting the app")
		Eventually(cf.Cf("delete", testAppName, "-f", "-r"), cf.CfTimeout).Should(gexec.Exit(0))

		By("ensuring the service instance is deleted")
		Eventually(cf.Cf("delete-service", serviceInstanceName, "-f"), cf.CfTimeout).Should(gexec.Exit())
		cf.AwaitServiceDeletion(serviceInstanceName)
	})

	It("Upgrades from an older release", func() {
		By("redeploying ODB with latest version")
		manifest := manifestForUpgrade()
		changeReleaseVersion("on-demand-service-broker", manifest, latestODBVersion)
		changeReleaseVersion("redis-example-service-adapter", manifest, latestServiceAdapterVersion)
		boshClient.DeployODB(*manifest)

		By("exercising the service instance")
		cf.PutToTestApp(testAppURL, "foo", "bar")

		By("running the upgrade errand")
		taskOutput := boshClient.RunErrand(brokerBoshDeploymentName, "upgrade-all-service-instances", []string{}, "")
		Expect(taskOutput.ExitCode).To(Equal(0))
		Expect(taskOutput.StdOut).To(ContainSubstring("Number of successful operations: 1"))

		By("updating the service instance")
		session := cf.Cf("update-service", serviceInstanceName, "-c", `{"maxclients": 60}`)
		Eventually(session, cf.CfTimeout).Should(gexec.Exit(0))
		cf.AwaitServiceUpdate(serviceInstanceName)

		By("running the delete all errand")
		taskOutput = boshClient.RunErrand(brokerBoshDeploymentName, "delete-all-service-instances", []string{}, "")
		Expect(taskOutput.ExitCode).To(Equal(0))
		cf.AwaitServiceDeletion(serviceInstanceName)
	})
})

func manifestForUpgrade() *bosh.BoshManifest {
	manifestContent, err := ioutil.ReadFile(path.Join(ciRootPath, manifestForUpgradePath))
	Expect(err).NotTo(HaveOccurred())
	manifest := new(bosh.BoshManifest)
	err = yaml.Unmarshal(manifestContent, manifest)
	Expect(err).NotTo(HaveOccurred())
	return manifest
}

func changeReleaseVersion(releaseName string, manifest *bosh.BoshManifest, releaseVersion string) {
	for index, release := range manifest.Releases {
		if strings.Contains(release.Name, releaseName) {
			manifest.Releases[index].Version = releaseVersion
			return
		}
	}
	Fail("No release found for " + releaseName)
}
