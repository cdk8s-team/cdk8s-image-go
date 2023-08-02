package cdk8simage


// Props for `Image`.
type ImageProps struct {
	// The docker build context directory (where `Dockerfile` is).
	Dir *string `field:"required" json:"dir" yaml:"dir"`
	// List of build args to pass to the build action.
	BuildArgs *[]*BuildArg `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to Dockerfile.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Name for the image.
	//
	// Docker convention is {registry_name}/{name}:{tag}
	// Visit https://docs.docker.com/engine/reference/commandline/tag/ for more information
	// Default: - auto-generated name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Set to specify the target platform for the build output, (for example, linux/amd64, linux/arm64, or darwin/amd64).
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// The registry URL to use.
	//
	// This will be used as the prefix for the image name.
	//
	// For example, if you have a local registry listening on port 500, you can set this to `localhost:5000`.
	// Default: "docker.io/library"
	//
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// Tag for the image.
	//
	// Docker convention is {registry_name}/{name}:{tag}
	// Visit https://docs.docker.com/engine/reference/commandline/tag/ for more information
	// Default: "latest".
	//
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

