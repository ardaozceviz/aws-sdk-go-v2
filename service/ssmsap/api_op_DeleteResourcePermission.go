// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmsap

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmsap/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes permissions associated with the target database.
func (c *Client) DeleteResourcePermission(ctx context.Context, params *DeleteResourcePermissionInput, optFns ...func(*Options)) (*DeleteResourcePermissionOutput, error) {
	if params == nil {
		params = &DeleteResourcePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResourcePermission", params, optFns, c.addOperationDeleteResourcePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResourcePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourcePermissionInput struct {

	// The Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	ResourceArn *string

	// Delete or restore the permissions on the target database.
	ActionType types.PermissionActionType

	// The Amazon Resource Name (ARN) of the source resource.
	SourceResourceArn *string

	noSmithyDocumentSerde
}

type DeleteResourcePermissionOutput struct {

	// The policy that removes permissions on the target database.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResourcePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteResourcePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteResourcePermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteResourcePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourcePermission(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteResourcePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-sap",
		OperationName: "DeleteResourcePermission",
	}
}
