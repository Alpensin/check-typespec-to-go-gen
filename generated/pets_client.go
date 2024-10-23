// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.2, generator: @autorest/go@4.0.0-preview.68)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package generated

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PetsClient contains the methods for the Pets group.
// Don't use this type directly, use a constructor function instead.
type PetsClient struct {
	internal *azcore.Client
}

// List -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 0.0.0
//   - options - PetsClientListOptions contains the optional parameters for the PetsClient.List method.
func (client *PetsClient) List(ctx context.Context, options *PetsClientListOptions) (PetsClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return PetsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PetsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PetsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *PetsClient) listCreateRequest(ctx context.Context, _ *PetsClientListOptions) (*policy.Request, error) {
	urlPath := "/pets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(	host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PetsClient) listHandleResponse(resp *http.Response) (PetsClientListResponse, error) {
	result := PetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PetArray); err != nil {
		return PetsClientListResponse{}, err
	}
	return result, nil
}

