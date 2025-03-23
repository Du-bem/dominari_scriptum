package types

import (
	"strings"
	"time"

	"github.com/bitcoin-sv/spv-wallet/models/response"
)

const timeFormat string = "2006-01-02 15:04:05"

type satelliteTime time.Time

func (c *satelliteTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(timeFormat, value) //parse time
	if err != nil {
		return err
	}
	*c = satelliteTime(t) //set result using the pointer
	return nil
}

func (c satelliteTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format(timeFormat) + `"`), nil
}

// SatelliteRequestData defines the raw data received from the satelite.
type SatelliteRequestData struct {
	Name       string        `json:"name"`
	RecordedOn satelliteTime `json:"time"`
	Position   []string      `json:"position"`
	Velocity   []string      `json:"velocity"`
}

// SatelliteUIData describes the data processed for the UI display
type SatelliteUIData struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PositionX  string    `json:"position_x"`
	PositionY  string    `json:"position_y"`
	PositionZ  string    `json:"position_z"`
	VelocityX  string    `json:"velocity_x"`
	VelocityY  string    `json:"velocity_y"`
	VelocityZ  string    `json:"velocity_z"`
	RecordedOn time.Time `json:"time"`
	CreatedOn  time.Time `json:"created_on"`
	UpdatedOn  time.Time `json:"updated_on"`
	CheckSum   string    `json:"checksum"`
}

type ServerAPI struct {
	DBInfo
	AccountWalletInfo
}

type AccountWalletInfo interface {
	GetBalance() uint64
	GetTransactions() ([]*response.Transaction, error)
	GetTransaction(txID string) (*response.Transaction, error)
	PublishCheckSum(checksum string) (string, error)
}

type DBInfo interface {
	Insert(checksum string, data *SatelliteRequestData) (int, error)
	List(offset int) ([]*SatelliteUIData, error)
	SearchByID(recordID int) (*SatelliteUIData, error)
}
