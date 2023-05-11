package cdk8simage


// Build arg to pass to the docker build.
type BuildArg struct {
	// the name of the build arg.
	Name *string `field:"required" json:"name" yaml:"name"`
	// the value of the build arg.
	Value *string `field:"required" json:"value" yaml:"value"`
}

