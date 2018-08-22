package httpserver

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/op/go-logging"
)

// HTTPServer application server HTTP server
// uses labstack/echo
type HTTPServer struct {
	e               *echo.Echo
	log             *logging.Logger
	cfg             Config
	onStartCallback func(*echo.Echo)
}

// New creates the HTTP server
func New(l *logging.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		e:   echo.New(),
		log: l,
		cfg: cfg,
	}
}

// SetOnStartCallback register a startup callback just before listening to the configured port
func (h *HTTPServer) SetOnStartCallback(cb func(*echo.Echo)) {
	h.onStartCallback = cb
}

// Start the HTTP server & run onStartCallback callback if defined
func (h *HTTPServer) Start() {
	if h.cfg.EnableAccessLogs {
		h.e.Use(middleware.Logger())
	}

	if h.onStartCallback != nil {
		h.onStartCallback(h.e)
	}

	h.log.Errorf("HTTP server error: %s",
		h.e.Start(fmt.Sprintf(":%d", h.cfg.Port)),
	)
}
