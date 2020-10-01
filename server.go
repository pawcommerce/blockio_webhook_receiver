package blockio_webhook_receiver

import (
  "github.com/valyala/fasthttp"
  "strings"
)

type NotificationHandler func(*Notification) bool

type server struct {
  listen string
  path string
  handler NotificationHandler
  filter string
  enforceWhitelist bool
}

func New(listen string, path string, handler NotificationHandler) *server {
  return &server{listen, path, handler, "", true}
}

func (s *server) SetFilter(noteType string) *server {
  s.filter = noteType
  return s
}

func (s *server) DisableWhitelist() *server {
  s.enforceWhitelist = false
  return s
}

func (s *server) EnableWhitelist() *server {
  s.enforceWhitelist = true
  return s
}

func (s *server) Start() error {
  return fasthttp.ListenAndServe(s.listen, s.muxer)
}

func (s *server) authorize(ctx *fasthttp.RequestCtx) bool {
  if !s.enforceWhitelist {
    return true
  }

  allowed := false
  ctx.Request.Header.VisitAll(func(key, value []byte) {
    if strings.ToLower(string(key)) == "x-forwarded-for" {
      allowed = IsWhitelisted(string(value))
    }
  })

  return allowed
}

func (s *server) muxer(ctx *fasthttp.RequestCtx) {
  if (string(ctx.Path()) == s.path && string(ctx.Method()) == "POST") {

    if s.authorize(ctx) {
      HTTPNotificationHandler(s.handler, s.filter)(ctx)
      return
    }

    ctx.Error("Forbidden", 403)
  } else {
    ctx.NotFound()
  }
}

func HTTPNotificationHandler(h NotificationHandler, f string) fasthttp.RequestHandler {
  return func (ctx *fasthttp.RequestCtx) {
    body := ctx.PostBody()

    if len(body) == 0 {
      ctx.Error("No content", 400)
      return
    }

    n, err := ParseNotification(body)

    if err != nil {
      ctx.Error("Error parsing json", 400)
      return
    }

    if (f != "" && n.Type != f) {
      ctx.SuccessString("text/plain", "OK")
    }

    if h(n) {
      ctx.SuccessString("text/plain", "OK")
    } else {
      ctx.Error("Error processing notification", 500)
    }

  }
}
