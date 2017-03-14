package resources_test

import (
	. "github.com/treacher/ship-it/providers/aws/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoadBalancer", func() {
	var (
		loadBalancerResource      LoadBalancer
		loadBalancerConfiguration LoadBalancerConfiguration
	)

	BeforeEach(func() {
		loadBalancerConfiguration = LoadBalancerConfiguration{
			Subnets:         []string{"sb-12345", "sb-54321", "sb-543"},
			IdleTimeout:     36000,
			CertArn:         "arn:aws:acm:foo:123131313",
			HealthCheckUrl:  "HTTPS:443/health_check",
			Scheme:          "internet-facing",
			ApplicationName: "FooBar",
		}
	})

	Describe(".generate", func() {
		It("Generates an AWS AutoScalingGroup resources with the correct attributes", func() {
			expectedOutput := `"ElasticLoadBalancer" : {
      "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
      "Properties" : {
        "ConnectionSettings" : {
           "IdleTimeout" : 36000
        },
        "Subnets" : ["sb-12345", "sb-54321", "sb-543"],
        "CrossZone" : "true",
        "SecurityGroups" : [
          { "Ref" : "LoadBalancerSecurityGroup"}
        ],
        "Listeners" : [
          {
            "LoadBalancerPort" : "443",
            "InstancePort" : "443",
            "Protocol" : "HTTPS",
            "InstanceProtocol" : "HTTPS",
            "SSLCertificateId": "arn:aws:acm:foo:123131313"
          }
        ],
        "HealthCheck" : {
          "Target" : "HTTPS:443/health_check",
          "HealthyThreshold" : "3",
          "UnhealthyThreshold" : "5",
          "Interval" : "30",
          "Timeout" : "5"
        },
        "ConnectionDrainingPolicy" : {
          "Enabled" : "true",
          "Timeout": "30"
        },
        "Scheme": "internet-facing",
        "Tags" : [
          {
            "Key" : "Name",
            "Value" : "FooBar"
          }
        ]
      }
    }
`
			Expect(loadBalancerResource.ToCloudformation(loadBalancerConfiguration)).To(Equal(expectedOutput))
		})
	})
})
