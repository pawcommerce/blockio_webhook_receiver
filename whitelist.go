package blockio_webhook_receiver

var ServerWhitelist = map[string]bool{
  "45.56.79.5": true,
  "45.56.123.170": true,
  "45.33.20.161": true,
  "45.33.4.167": true,
  "2600:3c00::f03c:91ff:fe33:2e14": true,
  "2600:3c00::f03c:91ff:fe89:bb9b": true,
  "2600:3c00::f03c:91ff:fe33:d082": true,
  "2600:3c00::f03c:92ff:fe5e:4219": true,
}

func IsWhitelisted(addr string) bool {
    _, ok := ServerWhitelist[addr]
    return ok 
}
