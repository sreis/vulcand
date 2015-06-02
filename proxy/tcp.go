package proxy

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/mailgun/vulcand/engine"

	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/mailgun/log"
)

func init() {
	RegisterFactory("tcp", newTcpHandler)
}

type tcpSrv struct {
	mux      *mux
	listener engine.Listener
}

func newTcpHandler(m *mux, l engine.Listener) (ProtocolHandler, error) {
	return &tcpSrv{
		mux:      m,
		listener: l,
	}, nil
}

func (s *tcpSrv) Listener() engine.Listener {
	return s.listener
}

func (s *tcpSrv) GetFile() (*FileDescriptor, error) {
	log.Infof("[sreis]@tcp.GetFile")
	// TODO
	return nil, nil
}

func (s *tcpSrv) String() string {
	return fmt.Sprintf("%s->(%v)", s.mux, &s.listener)
}

func (s *tcpSrv) isTLS() bool {
	log.Infof("[sreis]@tcp.isTLS")
	return false
}

func (s *tcpSrv) updateListener(l engine.Listener) error {
	log.Infof("[sreis]@tcp.updateListener")
	if s.listener.Protocol != l.Protocol {
		return fmt.Errorf("conflicting protocol %s and %s", s.listener.Protocol, l.Protocol)
	}

	log.Infof("%v update %v", s, &l)
	s.listener = l
	return s.reload()
}

func (s *tcpSrv) isServing() bool {
	log.Infof("[sreis]@tcp.isServing")
	// TODO
	return true
}
func (s *tcpSrv) hasListeners() bool {
	log.Infof("[sreis]@tcp.hasListeners")
	return false
}
func (s *tcpSrv) takeFile(f *FileDescriptor) error {
	log.Infof("[sreis]@tcp.takeFile")
	// TODO
	return nil
}

func (s *tcpSrv) reload() error {
	log.Infof("[sreis]@tcp.reload")
	// TODO
	return nil
}

func (s *tcpSrv) shutdown() {
	log.Infof("[sreis]@tcp.shutdown")
	// TODO
}

func (s *tcpSrv) newTLSConfig() (*tls.Config, error) {
	log.Infof("[sreis]@tcp.newTLSConfig")
	// TODO
	return nil, nil
}

func (s *tcpSrv) start() error {
	log.Infof("%s start", s)

	listener, err := net.Listen(s.listener.Address.Network, s.listener.Address.Address)
	if err != nil {
		return err
	}

	go acceptRawTCP(listener)

	return nil
}

func handleRawTCP(conn net.Conn) {
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	log.Warningf("Received message: %s", msg)
}

func acceptRawTCP(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Errorf("Error accepting connection.")
		} else {
			log.Infof("Received Message from %s", conn.RemoteAddr())
			go handleRawTCP(conn)
		}
	}
}
