package resources

import (
	"fmt"
	"strings"
)

type AutoScalingGroup struct{}

type AutoScalingGroupConfiguration struct {
	Subnets                []string
	LaunchConfigurationRef string
	LoadBalancerRef        string
	DesiredInstances       uint16
	MaxInstances           uint16
	MinInstances           uint16
	ApplicationName        string
	DockerImage            string
}

func (asg *AutoScalingGroup) ToCloudformation(config AutoScalingGroupConfiguration) string {
	cloudformation := asg.cloudformationTemplate()
	subnetString := fmt.Sprintf(`["%s"]`, strings.Join(config.Subnets, `", "`))

	return fmt.Sprintf(cloudformation,
		subnetString,
		config.LaunchConfigurationRef,
		config.LoadBalancerRef,
		config.DesiredInstances,
		config.MaxInstances,
		config.MinInstances,
		config.ApplicationName,
		config.DockerImage,
		config.DesiredInstances,
		config.DesiredInstances,
		config.DesiredInstances,
	)
}

func (asg *AutoScalingGroup) cloudformationTemplate() string {
	return `"AutoScalingGroup" : {
  "Type": "AWS::AutoScaling::AutoScalingGroup",
  "Properties": {
    "HealthCheckPeriod": "600",
    "HealthCheckType": "ELB",
    "VPCZoneIdentifier": %s,
    "LaunchConfigurationName" : { "Ref" : "%s" },
    "LoadBalancerNames" : [ { "Ref" : "%s" } ],
    "DesiredCapacity" : %d,
    "MaxSize" : %d,
    "MinSize" : %d,
    "Tags" : [
      {
        "Key" : "Name",
        "Value" : "%s",
        "PropagateAtLaunch" : "true"
      },
      {
        "Key" : "DockerImage",
        "Value" : "%s",
        "PropagateAtLaunch" : "true"
      }
    ]
  }
  "CreationPolicy": {
    "ResourceSignal": {
      "Count": %d,
      "Timeout": "PT10M"
    }
  },
  "UpdatePolicy": {
    "AutoScalingRollingUpdate": {
      "MaxBatchSize": %d,
      "MinInstancesInService": %d,
      "PauseTime": "PT10M",
      "WaitOnResourceSignals": true
    }
  }
}
`
}
