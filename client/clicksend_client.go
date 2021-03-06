// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/danikarik/clicksend-go/client/account"
	"github.com/danikarik/clicksend-go/client/account_recharge"
	"github.com/danikarik/clicksend-go/client/contact"
	"github.com/danikarik/clicksend-go/client/contact_list"
	"github.com/danikarik/clicksend-go/client/countries"
	"github.com/danikarik/clicksend-go/client/delivery_issues"
	"github.com/danikarik/clicksend-go/client/detect_address"
	"github.com/danikarik/clicksend-go/client/email_delivery_receipt_rules"
	"github.com/danikarik/clicksend-go/client/email_marketing"
	"github.com/danikarik/clicksend-go/client/email_to_sms"
	"github.com/danikarik/clicksend-go/client/fax"
	"github.com/danikarik/clicksend-go/client/fax_delivery_receipt_rules"
	"github.com/danikarik/clicksend-go/client/inbound_fax_rules"
	"github.com/danikarik/clicksend-go/client/inbound_sms_rules"
	"github.com/danikarik/clicksend-go/client/master_email_templates"
	"github.com/danikarik/clicksend-go/client/mms"
	"github.com/danikarik/clicksend-go/client/mms_campaign"
	"github.com/danikarik/clicksend-go/client/number"
	"github.com/danikarik/clicksend-go/client/post_letter"
	"github.com/danikarik/clicksend-go/client/post_postcard"
	"github.com/danikarik/clicksend-go/client/post_return_address"
	"github.com/danikarik/clicksend-go/client/referral_account"
	"github.com/danikarik/clicksend-go/client/reseller_account"
	"github.com/danikarik/clicksend-go/client/search"
	"github.com/danikarik/clicksend-go/client/sms"
	"github.com/danikarik/clicksend-go/client/sms_campaign"
	"github.com/danikarik/clicksend-go/client/sms_delivery_receipt_rules"
	"github.com/danikarik/clicksend-go/client/statistics"
	"github.com/danikarik/clicksend-go/client/subaccount"
	"github.com/danikarik/clicksend-go/client/timezones"
	"github.com/danikarik/clicksend-go/client/transactional_email"
	"github.com/danikarik/clicksend-go/client/transfer_credit"
	"github.com/danikarik/clicksend-go/client/upload"
	"github.com/danikarik/clicksend-go/client/user_email_templates"
	"github.com/danikarik/clicksend-go/client/voice"
	"github.com/danikarik/clicksend-go/client/voice_delivery_receipt_rules"
)

// Default clicksend HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "rest.clicksend.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v3"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new clicksend HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Clicksend {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new clicksend HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Clicksend {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new clicksend client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Clicksend {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Clicksend)
	cli.Transport = transport

	cli.Account = account.New(transport, formats)

	cli.AccountRecharge = account_recharge.New(transport, formats)

	cli.Contact = contact.New(transport, formats)

	cli.ContactList = contact_list.New(transport, formats)

	cli.Countries = countries.New(transport, formats)

	cli.DeliveryIssues = delivery_issues.New(transport, formats)

	cli.DetectAddress = detect_address.New(transport, formats)

	cli.EmailDeliveryReceiptRules = email_delivery_receipt_rules.New(transport, formats)

	cli.EmailMarketing = email_marketing.New(transport, formats)

	cli.EmailToSMS = email_to_sms.New(transport, formats)

	cli.FAX = fax.New(transport, formats)

	cli.FAXDeliveryReceiptRules = fax_delivery_receipt_rules.New(transport, formats)

	cli.InboundFAXRules = inbound_fax_rules.New(transport, formats)

	cli.InboundSMSRules = inbound_sms_rules.New(transport, formats)

	cli.MasterEmailTemplates = master_email_templates.New(transport, formats)

	cli.MMS = mms.New(transport, formats)

	cli.MMSCampaign = mms_campaign.New(transport, formats)

	cli.Number = number.New(transport, formats)

	cli.PostLetter = post_letter.New(transport, formats)

	cli.PostPostcard = post_postcard.New(transport, formats)

	cli.PostReturnAddress = post_return_address.New(transport, formats)

	cli.ReferralAccount = referral_account.New(transport, formats)

	cli.ResellerAccount = reseller_account.New(transport, formats)

	cli.Search = search.New(transport, formats)

	cli.SMS = sms.New(transport, formats)

	cli.SMSCampaign = sms_campaign.New(transport, formats)

	cli.SMSDeliveryReceiptRules = sms_delivery_receipt_rules.New(transport, formats)

	cli.Statistics = statistics.New(transport, formats)

	cli.Subaccount = subaccount.New(transport, formats)

	cli.Timezones = timezones.New(transport, formats)

	cli.TransactionalEmail = transactional_email.New(transport, formats)

	cli.TransferCredit = transfer_credit.New(transport, formats)

	cli.Upload = upload.New(transport, formats)

	cli.UserEmailTemplates = user_email_templates.New(transport, formats)

	cli.Voice = voice.New(transport, formats)

	cli.VoiceDeliveryReceiptRules = voice_delivery_receipt_rules.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Clicksend is a client for clicksend
type Clicksend struct {
	Account *account.Client

	AccountRecharge *account_recharge.Client

	Contact *contact.Client

	ContactList *contact_list.Client

	Countries *countries.Client

	DeliveryIssues *delivery_issues.Client

	DetectAddress *detect_address.Client

	EmailDeliveryReceiptRules *email_delivery_receipt_rules.Client

	EmailMarketing *email_marketing.Client

	EmailToSMS *email_to_sms.Client

	FAX *fax.Client

	FAXDeliveryReceiptRules *fax_delivery_receipt_rules.Client

	InboundFAXRules *inbound_fax_rules.Client

	InboundSMSRules *inbound_sms_rules.Client

	MasterEmailTemplates *master_email_templates.Client

	MMS *mms.Client

	MMSCampaign *mms_campaign.Client

	Number *number.Client

	PostLetter *post_letter.Client

	PostPostcard *post_postcard.Client

	PostReturnAddress *post_return_address.Client

	ReferralAccount *referral_account.Client

	ResellerAccount *reseller_account.Client

	Search *search.Client

	SMS *sms.Client

	SMSCampaign *sms_campaign.Client

	SMSDeliveryReceiptRules *sms_delivery_receipt_rules.Client

	Statistics *statistics.Client

	Subaccount *subaccount.Client

	Timezones *timezones.Client

	TransactionalEmail *transactional_email.Client

	TransferCredit *transfer_credit.Client

	Upload *upload.Client

	UserEmailTemplates *user_email_templates.Client

	Voice *voice.Client

	VoiceDeliveryReceiptRules *voice_delivery_receipt_rules.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Clicksend) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Account.SetTransport(transport)

	c.AccountRecharge.SetTransport(transport)

	c.Contact.SetTransport(transport)

	c.ContactList.SetTransport(transport)

	c.Countries.SetTransport(transport)

	c.DeliveryIssues.SetTransport(transport)

	c.DetectAddress.SetTransport(transport)

	c.EmailDeliveryReceiptRules.SetTransport(transport)

	c.EmailMarketing.SetTransport(transport)

	c.EmailToSMS.SetTransport(transport)

	c.FAX.SetTransport(transport)

	c.FAXDeliveryReceiptRules.SetTransport(transport)

	c.InboundFAXRules.SetTransport(transport)

	c.InboundSMSRules.SetTransport(transport)

	c.MasterEmailTemplates.SetTransport(transport)

	c.MMS.SetTransport(transport)

	c.MMSCampaign.SetTransport(transport)

	c.Number.SetTransport(transport)

	c.PostLetter.SetTransport(transport)

	c.PostPostcard.SetTransport(transport)

	c.PostReturnAddress.SetTransport(transport)

	c.ReferralAccount.SetTransport(transport)

	c.ResellerAccount.SetTransport(transport)

	c.Search.SetTransport(transport)

	c.SMS.SetTransport(transport)

	c.SMSCampaign.SetTransport(transport)

	c.SMSDeliveryReceiptRules.SetTransport(transport)

	c.Statistics.SetTransport(transport)

	c.Subaccount.SetTransport(transport)

	c.Timezones.SetTransport(transport)

	c.TransactionalEmail.SetTransport(transport)

	c.TransferCredit.SetTransport(transport)

	c.Upload.SetTransport(transport)

	c.UserEmailTemplates.SetTransport(transport)

	c.Voice.SetTransport(transport)

	c.VoiceDeliveryReceiptRules.SetTransport(transport)

}
