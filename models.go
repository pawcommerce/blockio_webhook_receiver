package blockio_webhook_receiver

import (
  "encoding/json"
  "errors"
  "github.com/shopspring/decimal"
)

type AddressData struct {
  Network          string             `json:"network"`
  Address          string             `json:"address"`
  BalanceChange    decimal.Decimal    `json:"balance_change"`
  AmountSent       decimal.Decimal    `json:"amount_sent"`
  AmountReceived   decimal.Decimal    `json:"amount_received"`
  TxId             string             `json:"txid"`
  Confirmations    int                `json:"confirmations"`
  IsGreen          bool               `json:"is_green"`
}

type Notification struct {
  Id               string           `json:"notification_id"`
  DeliveryAttempt  int              `json:"delivery_attempt"`
  CreatedAt        int              `json:"created_at"`
  Type             string           `json:"type"`
  RawData          json.RawMessage  `json:"data"`
}

func (n *Notification) AddressData() (*AddressData, error) {
  var addrData AddressData

  if n.Type != "address" {
    return nil, errors.New("Can only get Address data from address typed notifications")
  }

  err := json.Unmarshal(n.RawData, &addrData)
  return &addrData, err

}

func ParseNotification(jsonstr []byte) (*Notification, error) {
  var n Notification
  err := json.Unmarshal(jsonstr, &n)
  return &n, err
}
