// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new email campaigns API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for email campaigns API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateEmailCampaign creates an email campaign
*/
func (a *Client) CreateEmailCampaign(params *CreateEmailCampaignParams, authInfo runtime.ClientAuthInfoWriter) (*CreateEmailCampaignCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEmailCampaignParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEmailCampaign",
		Method:             "POST",
		PathPattern:        "/emailCampaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEmailCampaignReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEmailCampaignCreated), nil

}

/*
DeleteEmailCampaigns deletes an email campaign
*/
func (a *Client) DeleteEmailCampaigns(params *DeleteEmailCampaignsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEmailCampaignsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEmailCampaignsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEmailCampaigns",
		Method:             "DELETE",
		PathPattern:        "/emailCampaigns/{campaignId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEmailCampaignsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEmailCampaignsNoContent), nil

}

/*
EmailExportRecipients exports the recipients of a campaign
*/
func (a *Client) EmailExportRecipients(params *EmailExportRecipientsParams, authInfo runtime.ClientAuthInfoWriter) (*EmailExportRecipientsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailExportRecipientsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "emailExportRecipients",
		Method:             "POST",
		PathPattern:        "/emailCampaigns/{campaignId}/exportRecipients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailExportRecipientsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EmailExportRecipientsAccepted), nil

}

/*
GetEmailCampaign gets campaign informations
*/
func (a *Client) GetEmailCampaign(params *GetEmailCampaignParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmailCampaignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmailCampaignParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEmailCampaign",
		Method:             "GET",
		PathPattern:        "/emailCampaigns/{campaignId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEmailCampaignReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmailCampaignOK), nil

}

/*
GetEmailCampaigns returns all your created campaigns
*/
func (a *Client) GetEmailCampaigns(params *GetEmailCampaignsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEmailCampaignsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmailCampaignsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEmailCampaigns",
		Method:             "GET",
		PathPattern:        "/emailCampaigns",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEmailCampaignsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmailCampaignsOK), nil

}

/*
SendEmailCampaignNow sends an email campaign id of the campaign immediately
*/
func (a *Client) SendEmailCampaignNow(params *SendEmailCampaignNowParams, authInfo runtime.ClientAuthInfoWriter) (*SendEmailCampaignNowNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendEmailCampaignNowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendEmailCampaignNow",
		Method:             "POST",
		PathPattern:        "/emailCampaigns/{campaignId}/sendNow",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendEmailCampaignNowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendEmailCampaignNowNoContent), nil

}

/*
SendReport sends the report of a campaigns

A PDF will be sent to the specified email addresses
*/
func (a *Client) SendReport(params *SendReportParams, authInfo runtime.ClientAuthInfoWriter) (*SendReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendReport",
		Method:             "POST",
		PathPattern:        "/emailCampaigns/{campaignId}/sendReport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendReportNoContent), nil

}

/*
SendTestEmail sends an email campaign to your test list
*/
func (a *Client) SendTestEmail(params *SendTestEmailParams, authInfo runtime.ClientAuthInfoWriter) (*SendTestEmailNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendTestEmailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendTestEmail",
		Method:             "POST",
		PathPattern:        "/emailCampaigns/{campaignId}/sendTest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendTestEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SendTestEmailNoContent), nil

}

/*
UpdateCampaignStatus updates a campaign status
*/
func (a *Client) UpdateCampaignStatus(params *UpdateCampaignStatusParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCampaignStatusNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCampaignStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCampaignStatus",
		Method:             "PUT",
		PathPattern:        "/emailCampaigns/{campaignId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCampaignStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateCampaignStatusNoContent), nil

}

/*
UpdateEmailCampaigns updates a campaign
*/
func (a *Client) UpdateEmailCampaigns(params *UpdateEmailCampaignsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateEmailCampaignsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEmailCampaignsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEmailCampaigns",
		Method:             "PUT",
		PathPattern:        "/emailCampaigns/{campaignId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEmailCampaignsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateEmailCampaignsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
