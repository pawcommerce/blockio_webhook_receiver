# blockio_webhook_receiver

High-performance webhook endpoint for block.io, to process your transactions rapidly.
See https://block.io/docs/notifications for instructions how to set up the notifications.

### Usage

Getting started is easy. Just import the receiver, create a callback function,
set a port to listen to, and start the server.

```go
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
  log.Fatal(recv.New(":8083", "/", PrintNotification).Start())
}
```

### Deployment

The server expects calls to be routed through a front-end load balancer,
such as HAProxy or Nginx, or a hosted solution like CloudFlare or CloudFront.

The allow list operates on the `X-Forwarded-For` header, so make sure that
this is properly set-up on the front-end server.

### API

##### func New (port, path, callback)

```go
func New(listen string, path string, handler NotificationHandler) *server {}
```

Returns a server object that will listen on ```port``` and ```path```, parse
messages and forward them to ```handler```

##### func server.SetFilter(type)

```go
func (s *server) SetFilter(noteType string) *server {}
```

Set a filter to only process specific notification types. Non-matching types will be voided.
Supported values:

- "address"

##### func server.DisableAllowlist()

```go
func (s *server) DisableAllowlist() *server {}
```

Disable matching to allow lists (enabled by default) for development environments.

**NOTE: DO NOT USE THIS IN PRODUCTION!**

##### func server.EnableAllowlist()

```go
func (s *server) EnableAllowlist() *server {}
```

Enable matching to allow lists.

##### func server.Start()

```go
func (s *server) Start() error {}
```

Starts the server or returns an error if this is not possible.
