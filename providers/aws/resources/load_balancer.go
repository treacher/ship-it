package resources

import (
	"fmt"
	"strings"
)

type LoadBalancer struct{}

type LoadBalancerConfiguration struct {
	Subnets         []string
	IdleTimeout     int32
	CertArn         string
	HealthCheckUrl  string
	Scheme          string
	ApplicationName string
}

func (lb *LoadBalancer) ToCloudformation(config LoadBalancerConfiguration) string {
	cloudformation := lb.cloudformationTemplate()
	subnetString := fmt.Sprintf(`["%s"]`, strings.Join(config.Subnets, `", "`))

	return fmt.Sprintf(cloudformation,
		config.IdleTimeout,
		subnetString,
		config.CertArn,
		config.HealthCheckUrl,
		config.Scheme,
		config.ApplicationName,
	)
}

func (lb *LoadBalancer) cloudformationTemplate() string {
	return `"ElasticLoadBalancer" : {
      "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
      "Properties" : {
        "ConnectionSettings" : {
           "IdleTimeout" : %d
        },
        "Subnets" : %s,
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
            "SSLCertificateId": "%s"
          }
        ],
        "HealthCheck" : {
          "Target" : "%s",
          "HealthyThreshold" : "3",
          "UnhealthyThreshold" : "5",
          "Interval" : "30",
          "Timeout" : "5"
        },
        "ConnectionDrainingPolicy" : {
          "Enabled" : "true",
          "Timeout": "30"
        },
        "Scheme": "%s",
        "Tags" : [
          {
            "Key" : "Name",
            "Value" : "%s"
          }
        ]
      }
    }
`
}
