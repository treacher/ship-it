package resources

type Resource interface {
	ToCloudformation(AutoScalingGroupConfiguration)
}
