package blockio_webhook_receiver

import (
  "testing"
  "strings"
  "os"
  "io"
  "github.com/shopspring/decimal"
)

var note *Notification

// just read a file into string
func readFile(path string) (string, error) {
	var buf strings.Builder

	fd, fErr := os.Open(path)
	if fErr != nil {
		return "", fErr
	}

	defer fd.Close()

	_, ioErr := io.Copy(&buf, fd)
	if ioErr != nil {
		return "", ioErr
	}

	return buf.String(), nil
}

func strEqual(name string, a string, b string, t *testing.T) {
  if a != b {
    t.Errorf("%s mismatch: expected %s, got %s", name, a, b)
  }
}

func intEqual(name string, a int, b int, t *testing.T) {
  if a != b {
    t.Errorf("%s mismatch: expected %d, got %d", name, a, b)
  }
}

func decEqual(name string, a decimal.Decimal, b decimal.Decimal, t *testing.T) {
  if !a.Equals(b) {
    t.Errorf("%s mismatch: expected %s, got %s", name, a, b)
  }
}

func TestParseNotification(t *testing.T) {
  str, err := readFile("./fixtures/balance_change.json")

  if err != nil {
    t.Errorf("Error reading fixture: %s", err)
  }

  note, err = ParseNotification([]byte(str))

  if err != nil {
    t.Errorf("Parse error: %s", err)
  }

  strEqual("Id", "1", note.Id, t)
  strEqual("Type", "address", note.Type, t)
  intEqual("DeliveryAttempt", 1, note.DeliveryAttempt, t)
  intEqual("CreatedAt", 1386474927, note.CreatedAt, t)

}

func TestAddressData(t *testing.T) {
  data, err := note.AddressData()

  if err != nil {
    t.Errorf("Error getting address data: %s", err)
  }

  bc, _ := decimal.NewFromString("68416.00000000");
  zero := decimal.NewFromInt(0);

  strEqual("Network", "DOGE", data.Network, t)
  strEqual("Address", "DLAznsPDLDRgsVcTFWRMYMG5uH6GddDtv8", data.Address, t)
  decEqual("BalanceChange", bc, data.BalanceChange, t)
  decEqual("AmountSent", zero, data.AmountSent, t)
  decEqual("AmountReceived", bc, data.AmountReceived, t)
  strEqual("TxId", "5f7e779f7600f54e528686e91d5891f3ae226ee907f461692519e549105f521c", data.TxId, t)
  intEqual("Confirmations", 1, data.Confirmations, t)

}
