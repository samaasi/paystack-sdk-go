package paystackapi

// Currency represents a strongly typed currency code.
type Currency string

const (
	CurrencyNGN Currency = "NGN"
	CurrencyGHS Currency = "GHS"
	CurrencyZAR Currency = "ZAR"
	CurrencyUSD Currency = "USD"
	CurrencyKES Currency = "KES"
)

// Channel represents a strongly typed payment channel.
type Channel string

const (
	ChannelCard         Channel = "card"
	ChannelBank         Channel = "bank"
	ChannelUSSD         Channel = "ussd"
	ChannelQR           Channel = "qr"
	ChannelMobileMoney  Channel = "mobile_money"
	ChannelBankTransfer Channel = "bank_transfer"
	ChannelEFT          Channel = "eft"
)

// Status represents a strongly typed transaction status.
type Status string

const (
	StatusSuccess   Status = "success"
	StatusFailed    Status = "failed"
	StatusPending   Status = "pending"
	StatusReversed  Status = "reversed"
	StatusAbandoned Status = "abandoned"
)

// Bearer represents who bears the transaction charges.
type Bearer string

const (
	BearerAccount    Bearer = "account"
	BearerSubaccount Bearer = "subaccount"
)

// RiskAction represents the risk assessment action.
type RiskAction string

const (
	RiskActionDefault RiskAction = "default"
	RiskActionAllow   RiskAction = "allow"
	RiskActionDeny    RiskAction = "deny"
)
