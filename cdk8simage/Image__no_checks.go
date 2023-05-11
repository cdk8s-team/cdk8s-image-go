//go:build no_runtime_type_checking

package cdk8simage

// Building without runtime type checking enabled, so all the below just return nil

func validateImage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewImageParameters(scope constructs.Construct, id *string, props *ImageProps) error {
	return nil
}

