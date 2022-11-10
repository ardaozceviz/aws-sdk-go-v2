// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the protection status of a task. You can set protectionEnabled to true
// to protect your task from termination during scale-in events from Service
// Autoscaling
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-auto-scaling.html)
// or deployments
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html).
// Task-protection, by default, expires after 2 hours at which point Amazon ECS
// unsets the protectionEnabled property making the task eligible for termination
// by a subsequent scale-in event. You can specify a custom expiration period for
// task protection from 1 minute to up to 2,880 minutes (48 hours). To specify the
// custom expiration period, set the expiresInMinutes property. The
// expiresInMinutes property is always reset when you invoke this operation for a
// task that already has protectionEnabled set to true. You can keep extending the
// protection expiration period of a task by invoking this operation repeatedly. To
// learn more about Amazon ECS task protection, see Task scale-in protection
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-scale-in-protection.html)
// in the Amazon Elastic Container Service Developer Guide. This operation is only
// supported for tasks belonging to an Amazon ECS service. Invoking this operation
// for a standalone task will result in an TASK_NOT_VALID failure. For more
// information, see API failure reasons
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/api_failures_messages.html.html).
// If you prefer to set task protection from within the container, we recommend
// using the Amazon ECS container agent endpoint
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-endpoint.html).
func (c *Client) UpdateTaskProtection(ctx context.Context, params *UpdateTaskProtectionInput, optFns ...func(*Options)) (*UpdateTaskProtectionOutput, error) {
	if params == nil {
		params = &UpdateTaskProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTaskProtection", params, optFns, c.addOperationUpdateTaskProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskProtectionInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task sets exist in.
	//
	// This member is required.
	Cluster *string

	// Specify true to mark a task for protection and false to unset protection, making
	// it eligible for termination.
	//
	// This member is required.
	ProtectionEnabled bool

	// A list of up to 10 task IDs or full ARN entries.
	//
	// This member is required.
	Tasks []string

	// If you set protectionEnabled to true, you can specify the duration for task
	// protection in minutes. You can specify a value from 1 minute to up to 2,880
	// minutes (48 hours). During this time, your task will not be terminated by
	// scale-in events from Service Auto Scaling or deployments. After this time period
	// lapses, protectionEnabled will be reset to false. If you don’t specify the time,
	// then the task is automatically protected for 120 minutes (2 hours).
	ExpiresInMinutes *int32

	noSmithyDocumentSerde
}

type UpdateTaskProtectionOutput struct {

	// Any failures associated with the call.
	Failures []types.Failure

	// A list of tasks with the following information.
	//
	// * taskArn: The task ARN.
	//
	// *
	// protectionEnabled: The protection status of the task. If scale-in protection is
	// enabled for a task, the value is true. Otherwise, it is false.
	//
	// *
	// expirationDate: The epoch time when protection for the task will expire.
	ProtectedTasks []types.ProtectedTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTaskProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTaskProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTaskProtection{}, middleware.After)
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
	if err = addOpUpdateTaskProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTaskProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTaskProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateTaskProtection",
	}
}
