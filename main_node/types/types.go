package types

import (
	"time"

	"github.com/bitcoin-sv/spv-wallet/models/response"
)

// SatelliteRequestData defines the raw data received from the satelite.
type SatelliteRequestData struct {
	Name       string    `json:"name"`
	RecordedOn time.Time `json:"time"`
	Position   []float64 `json:"position"`
	Velocity   []float64 `json:"velocity"`
}

// SatelliteUIData describes the data processed for the UI display
type SatelliteUIData struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PositionX  float64   `json:"position_x"`
	PositionY  float64   `json:"position_y"`
	PositionZ  float64   `json:"position_z"`
	VelocityX  float64   `json:"velocity_x"`
	VelocityY  float64   `json:"velocity_y"`
	VelocityZ  float64   `json:"velocity_z"`
	RecordedOn time.Time `json:"time"`
	CreatedOn  time.Time `json:"created_on"`
	UpdatedOn  time.Time `json:"updated_on"`
	CheckSum   string    `json:"checksum"`
}

type AccountWalletInfo interface {
	GetBalance() uint64
	GetTransactions() ([]*response.Transaction, error)
	GetTransaction(txID string) (*response.Transaction, error)
}
