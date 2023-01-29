// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all memberships resources within the caller's account.
func (c *Client) ListMemberships(ctx context.Context, params *ListMembershipsInput, optFns ...func(*Options)) (*ListMembershipsOutput, error) {
	if params == nil {
		params = &ListMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMemberships", params, optFns, c.addOperationListMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMembershipsInput struct {

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// A filter which will return only memberships in the specified status.
	Status types.MembershipStatus

	noSmithyDocumentSerde
}

type ListMembershipsOutput struct {

	// The list of memberships returned from the ListMemberships operation.
	//
	// This member is required.
	MembershipSummaries []types.MembershipSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMemberships{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMemberships(options.Region), middleware.Before); err != nil {
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

// ListMembershipsAPIClient is a client that implements the ListMemberships
// operation.
type ListMembershipsAPIClient interface {
	ListMemberships(context.Context, *ListMembershipsInput, ...func(*Options)) (*ListMembershipsOutput, error)
}

var _ ListMembershipsAPIClient = (*Client)(nil)

// ListMembershipsPaginatorOptions is the paginator options for ListMemberships
type ListMembershipsPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMembershipsPaginator is a paginator for ListMemberships
type ListMembershipsPaginator struct {
	options   ListMembershipsPaginatorOptions
	client    ListMembershipsAPIClient
	params    *ListMembershipsInput
	nextToken *string
	firstPage bool
}

// NewListMembershipsPaginator returns a new ListMembershipsPaginator
func NewListMembershipsPaginator(client ListMembershipsAPIClient, params *ListMembershipsInput, optFns ...func(*ListMembershipsPaginatorOptions)) *ListMembershipsPaginator {
	if params == nil {
		params = &ListMembershipsInput{}
	}

	options := ListMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMemberships page.
func (p *ListMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMembershipsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListMemberships(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "ListMemberships",
	}
}
