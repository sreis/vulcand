package proxy

import (
	"crypto/tls"

	"github.com/mailgun/vulcand/engine"
)

type ProtocolHandler interface {
	GetFile() (*FileDescriptor, error)
	String() string
	isTLS() bool
	updateListener(l engine.Listener) error
	isServing() bool
	hasListeners() bool
	takeFile(f *FileDescriptor) error
	// newHTTPServer() *httpServer
	reload() error
	shutdown()
	newTLSConfig() (*tls.Config, error)
	start() error
	// serve(srv *manners.GracefulServer)

	Listener() engine.Listener
}

type ProtocolHandlerFactory func(m *mux, l engine.Listener) (ProtocolHandler, error)

var protocolHandlers = map[string]ProtocolHandlerFactory{}

//
func RegisterFactory(protocol string, factory ProtocolHandlerFactory) error {
	// TODO: make protocol a type?
	// TODO: error handling
	protocolHandlers[protocol] = factory
	return nil
}

//
func GetFactory(protocol string) ProtocolHandlerFactory {
	// TODO: error handling
	return protocolHandlers[protocol]
}
