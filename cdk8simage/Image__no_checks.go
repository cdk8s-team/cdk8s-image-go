//go:build no_runtime_type_checking

// Build & Push local docker images inside CDK8s applications
package cdk8simage

// Building without runtime type checking enabled, so all the below just return nil

func validateImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewImageParameters(scope constructs.Construct, id *string, props *ImageProps) error {
	return nil
}

