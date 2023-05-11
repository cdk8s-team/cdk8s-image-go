// Build & Push local docker images inside CDK8s applications
package cdk8simage

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-image.BuildArg",
		reflect.TypeOf((*BuildArg)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-image.Image",
		reflect.TypeOf((*Image)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_Image{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-image.ImageProps",
		reflect.TypeOf((*ImageProps)(nil)).Elem(),
	)
}
