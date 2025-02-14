// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform configuration templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform configuration templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeploymentTemplate(params *CreateDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeploymentTemplateOK, *CreateDeploymentTemplateCreated, error)

	DeleteDeploymentTemplate(params *DeleteDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeploymentTemplateOK, error)

	GetDeploymentTemplate(params *GetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentTemplateOK, error)

	GetDeploymentTemplates(params *GetDeploymentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentTemplatesOK, error)

	GetGlobalDeploymentTemplates(params *GetGlobalDeploymentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalDeploymentTemplatesOK, error)

	SetDeploymentTemplate(params *SetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetDeploymentTemplateOK, *SetDeploymentTemplateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDeploymentTemplate creates deployment template

  Creates a deployment template.
*/
func (a *Client) CreateDeploymentTemplate(params *CreateDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeploymentTemplateOK, *CreateDeploymentTemplateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-deployment-template",
		Method:             "POST",
		PathPattern:        "/platform/configuration/templates/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateDeploymentTemplateOK:
		return value, nil, nil
	case *CreateDeploymentTemplateCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for platform_configuration_templates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDeploymentTemplate deletes deployment template

  Deletes a deployment template by id.
*/
func (a *Client) DeleteDeploymentTemplate(params *DeleteDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDeploymentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-deployment-template",
		Method:             "DELETE",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDeploymentTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-deployment-template: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTemplate gets deployment template

  Retrieves a deployment template by id.
*/
func (a *Client) GetDeploymentTemplate(params *GetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-deployment-template",
		Method:             "GET",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-deployment-template: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploymentTemplates gets deployment templates

  Retrieves all deployment templates.
*/
func (a *Client) GetDeploymentTemplates(params *GetDeploymentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-deployment-templates",
		Method:             "GET",
		PathPattern:        "/platform/configuration/templates/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentTemplatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-deployment-templates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGlobalDeploymentTemplates gets all templates cross region

  Global deployment template endpoint which fetches the deployment templates across all region services.
*/
func (a *Client) GetGlobalDeploymentTemplates(params *GetGlobalDeploymentTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGlobalDeploymentTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalDeploymentTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-global-deployment-templates",
		Method:             "GET",
		PathPattern:        "/platform/configuration/templates/deployments/global",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalDeploymentTemplatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalDeploymentTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-global-deployment-templates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetDeploymentTemplate sets deployment template

  Creates or updates a deployment template.
*/
func (a *Client) SetDeploymentTemplate(params *SetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetDeploymentTemplateOK, *SetDeploymentTemplateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDeploymentTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set-deployment-template",
		Method:             "PUT",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SetDeploymentTemplateOK:
		return value, nil, nil
	case *SetDeploymentTemplateCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for platform_configuration_templates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
