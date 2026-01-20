package terminal

import (
	"encoding/json"
	"github.com/samaasi/paystack-sdk-go/paystackapi"
)

type Terminal struct {
	ID          int    `json:"id"`
	Serial      string `json:"serial"`
	DeviceType  string `json:"device_type"`
	TerminalID  string `json:"terminal_id"`
	Integration int    `json:"integration"`
	Domain      string `json:"domain"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Status      string `json:"status"`
}

type TerminalEvent struct {
	ID         string          `json:"id"`
	TerminalID string          `json:"terminal_id"`
	Type       string          `json:"type"`
	Action     string          `json:"action"`
	Data       json.RawMessage `json:"data"`
	Status     string          `json:"status"`
}

type SendEventRequest struct {
	Type   string                 `json:"type"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type UpdateTerminalRequest struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type TerminalResponse struct {
	paystackapi.Response[Terminal]
}

type ListTerminalsResponse struct {
	paystackapi.Response[[]Terminal]
}

type TerminalEventResponse struct {
	paystackapi.Response[TerminalEvent]
}

type TerminalPresenceResponse struct {
	paystackapi.Response[struct {
		Online    bool `json:"online"`
		Available bool `json:"available"`
	}]
}
