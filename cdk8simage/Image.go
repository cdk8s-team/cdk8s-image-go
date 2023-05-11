package cdk8simage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-image-go/cdk8simage/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-image-go/cdk8simage/internal"
)

// Represents a docker image built during synthesis from a context directory (`dir`) with a `Dockerfile`.
//
// The image will be built using `docker build` and then pushed through `docker
// push`. The URL of the pushed image can be accessed through `image.url`.
//
// If you push to a registry other than docker hub, you can specify the registry
// URL through the `registry` option.
type Image interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The image URL to use in order to pull this instance of the image.
	Url() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Image
type jsiiProxy_Image struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Image) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewImage(scope constructs.Construct, id *string, props *ImageProps) Image {
	_init_.Initialize()

	if err := validateNewImageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Image{}

	_jsii_.Create(
		"cdk8s-image.Image",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewImage_Override(i Image, scope constructs.Construct, id *string, props *ImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-image.Image",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Image_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-image.Image",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Image) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

