package resources_test

import (
	. "github.com/treacher/ship-it/providers/aws/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AutoScalingGroup", func() {
	var (
		autoScalingGroupResource      AutoScalingGroup
		autoscalingGroupConfiguration AutoScalingGroupConfiguration
	)

	BeforeEach(func() {
		autoscalingGroupConfiguration = AutoScalingGroupConfiguration{
			Subnets:                []string{"sb-12345", "sb-54321"},
			LaunchConfigurationRef: "LaunchConfig1234",
			LoadBalancerRef:        "LB12345",
			DesiredInstances:       2,
			MaxInstances:           3,
			MinInstances:           1,
			ApplicationName:        "FooBar",
			DockerImage:            "docker.io/foo:1.23.4",
		}
	})

	Describe(".generate", func() {
		It("Generates an AWS AutoScalingGroup resources with the correct attributes", func() {
			expectedOutput := `"AutoScalingGroup" : {
  "Type": "AWS::AutoScaling::AutoScalingGroup",
  "Properties": {
    "HealthCheckPeriod": "600",
    "HealthCheckType": "ELB",
    "VPCZoneIdentifier": ["sb-12345", "sb-54321"],
    "LaunchConfigurationName" : { "Ref" : "LaunchConfig1234" },
    "LoadBalancerNames" : [ { "Ref" : "LB12345" } ],
    "DesiredCapacity" : 2,
    "MaxSize" : 3,
    "MinSize" : 1,
    "Tags" : [
      {
        "Key" : "Name",
        "Value" : "FooBar",
        "PropagateAtLaunch" : "true"
      },
      {
        "Key" : "DockerImage",
        "Value" : "docker.io/foo:1.23.4",
        "PropagateAtLaunch" : "true"
      }
    ]
  }
  "CreationPolicy": {
    "ResourceSignal": {
      "Count": 2,
      "Timeout": "PT10M"
    }
  },
  "UpdatePolicy": {
    "AutoScalingRollingUpdate": {
      "MaxBatchSize": 2,
      "MinInstancesInService": 2,
      "PauseTime": "PT10M",
      "WaitOnResourceSignals": true
    }
  }
}
`
			Expect(autoScalingGroupResource.ToCloudformation(autoscalingGroupConfiguration)).To(Equal(expectedOutput))
		})
	})
})
