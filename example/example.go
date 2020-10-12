package main

import (
  "fmt"
  "log"
  recv "github.com/pawcommerce/blockio_webhook_receiver"
)

func PrintNotification(n *recv.Notification) bool {
  fmt.Println("=========================")
  fmt.Printf("Id:        %s\n", n.Id)
  fmt.Printf("Timestamp: %d\n", n.CreatedAt)
  fmt.Printf("Attempt:   %d\n", n.DeliveryAttempt)
  addr, _ :=  n.AddressData()
  fmt.Printf("%+v\n", addr)

  return true
}

func main() {
  log.Fatal(recv.New(":8083", "/", PrintNotification).SetFilter("address").DisableAllowlist().Start())
}
