package librmonitor

import (
	"bufio"
	"io"
	"net"
	"sync"
	"time"
)

type Simulator struct {
	sync.Mutex
	clients map[string]chan []byte
	source  bufio.Reader
	stopCh  chan struct{}

	listener *net.TCPListener

	connErrors chan error
	connNotifs chan string
}

func (s *Simulator) ConnErrors() chan error {
	return s.connErrors
}

func (s *Simulator) ConnNotifs() chan string {
	return s.connNotifs
}

func Simulate(addr *net.TCPAddr, source io.Reader, stopCh chan struct{}) (*Simulator, error) {
	buf := bufio.NewReader(source)

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	var svr = &Simulator{
		clients:    make(map[string]chan []byte),
		source:     *buf,
		listener:   listener,
		connNotifs: make(chan string),
		connErrors: make(chan error),
	}

	return svr, nil
}

func (s *Simulator) readAndServe() error {
	for {
		select {
		case <-s.stopCh:
			return nil
		default:
			if len(s.clients) > 0 {
				line, _, err := s.source.ReadLine()
				if err != nil {
					return err
				}

				for _, l := range s.clients {
					l <- line
				}
				time.Sleep(2 * time.Second)
			}
		}
	}
}

func client(conn *net.TCPConn, stopCh chan struct{}, source chan []byte) error {
	for {
		select {
		case <-stopCh:
			return nil
		case line := <-source:
			line = append(line, '\n')
			if _, err := conn.Write(line); err != nil {
				return err
			}
		}
	}
}

func (s *Simulator) Run() {
	go s.readAndServe()

	for {
		c, err := s.listener.AcceptTCP()
		if err != nil {
			s.connErrors <- err
			continue
		}

		s.Lock()

		lch := make(chan []byte)
		s.clients[c.RemoteAddr().String()] = lch

		s.connNotifs <- c.RemoteAddr().String()
		go func() {
			if err := client(c, s.stopCh, lch); err != nil {
				s.connErrors <- err
			} else {
				s.connNotifs <- "client disconnected"
			}
			delete(s.clients, c.RemoteAddr().String())
		}()
		s.Unlock()
	}
}
