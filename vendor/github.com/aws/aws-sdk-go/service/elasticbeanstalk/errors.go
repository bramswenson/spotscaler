// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

const (

	// ErrCodeCodeBuildNotInServiceRegionException for service response error code
	// "CodeBuildNotInServiceRegionException".
	//
	// AWS CodeBuild is not available in the specified region.
	ErrCodeCodeBuildNotInServiceRegionException = "CodeBuildNotInServiceRegionException"

	// ErrCodeInsufficientPrivilegesException for service response error code
	// "InsufficientPrivilegesException".
	//
	// The specified account does not have sufficient privileges for one of more
	// AWS services.
	ErrCodeInsufficientPrivilegesException = "InsufficientPrivilegesException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// One or more input parameters is not valid. Please correct the input parameters
	// and try the operation again.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeManagedActionInvalidStateException for service response error code
	// "ManagedActionInvalidStateException".
	//
	// Cannot modify the managed action in its current state.
	ErrCodeManagedActionInvalidStateException = "ManagedActionInvalidStateException"

	// ErrCodeOperationInProgressException for service response error code
	// "OperationInProgressFailure".
	//
	// Unable to perform the specified operation because another operation that
	// effects an element in this activity is already in progress.
	ErrCodeOperationInProgressException = "OperationInProgressFailure"

	// ErrCodePlatformVersionStillReferencedException for service response error code
	// "PlatformVersionStillReferencedException".
	//
	// You cannot delete the platform version because there are still environments
	// running on it.
	ErrCodePlatformVersionStillReferencedException = "PlatformVersionStillReferencedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A resource doesn't exist for the specified Amazon Resource Name (ARN).
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceTypeNotSupportedException for service response error code
	// "ResourceTypeNotSupportedException".
	//
	// The type of the specified Amazon Resource Name (ARN) isn't supported for
	// this operation.
	ErrCodeResourceTypeNotSupportedException = "ResourceTypeNotSupportedException"

	// ErrCodeS3LocationNotInServiceRegionException for service response error code
	// "S3LocationNotInServiceRegionException".
	//
	// The specified S3 bucket does not belong to the S3 region in which the service
	// is running. The following regions are supported:
	//
	//    * IAD/us-east-1
	//
	//    * PDX/us-west-2
	//
	//    * DUB/eu-west-1
	ErrCodeS3LocationNotInServiceRegionException = "S3LocationNotInServiceRegionException"

	// ErrCodeS3SubscriptionRequiredException for service response error code
	// "S3SubscriptionRequiredException".
	//
	// The specified account does not have a subscription to Amazon S3.
	ErrCodeS3SubscriptionRequiredException = "S3SubscriptionRequiredException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// A generic service exception has occurred.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeSourceBundleDeletionException for service response error code
	// "SourceBundleDeletionFailure".
	//
	// Unable to delete the Amazon S3 source bundle associated with the application
	// version. The application version was deleted successfully.
	ErrCodeSourceBundleDeletionException = "SourceBundleDeletionFailure"

	// ErrCodeTooManyApplicationVersionsException for service response error code
	// "TooManyApplicationVersionsException".
	//
	// The specified account has reached its limit of application versions.
	ErrCodeTooManyApplicationVersionsException = "TooManyApplicationVersionsException"

	// ErrCodeTooManyApplicationsException for service response error code
	// "TooManyApplicationsException".
	//
	// The specified account has reached its limit of applications.
	ErrCodeTooManyApplicationsException = "TooManyApplicationsException"

	// ErrCodeTooManyBucketsException for service response error code
	// "TooManyBucketsException".
	//
	// The specified account has reached its limit of Amazon S3 buckets.
	ErrCodeTooManyBucketsException = "TooManyBucketsException"

	// ErrCodeTooManyConfigurationTemplatesException for service response error code
	// "TooManyConfigurationTemplatesException".
	//
	// The specified account has reached its limit of configuration templates.
	ErrCodeTooManyConfigurationTemplatesException = "TooManyConfigurationTemplatesException"

	// ErrCodeTooManyEnvironmentsException for service response error code
	// "TooManyEnvironmentsException".
	//
	// The specified account has reached its limit of environments.
	ErrCodeTooManyEnvironmentsException = "TooManyEnvironmentsException"

	// ErrCodeTooManyPlatformsException for service response error code
	// "TooManyPlatformsException".
	//
	// You have exceeded the maximum number of allowed platforms associated with
	// the account.
	ErrCodeTooManyPlatformsException = "TooManyPlatformsException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The number of tags in the resource would exceed the number of tags that each
	// resource can have.
	//
	// To calculate this, the operation considers both the number of tags the resource
	// already has and the tags this operation would add if it succeeded.
	ErrCodeTooManyTagsException = "TooManyTagsException"
)
