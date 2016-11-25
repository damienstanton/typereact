package main

import (
	"errors"
	"net"
	"time"
)

// Listener wraps a regular TCPListener
type Listener struct {
	*net.TCPListener
	stop chan int
}

// New creates a new listener
func New(l net.Listener) (*Listener, error) {
	tcpL, ok := l.(*net.TCPListener)

	if !ok {
		return nil, errors.New("Cannot wrap listener")
	}

	retval := &Listener{}
	retval.TCPListener = tcpL
	retval.stop = make(chan int)

	return retval, nil
}

var errStop = errors.New("Listener stopped")

// Accept handles incoming connections
func (l *Listener) Accept() (net.Conn, error) {

	for {
		l.SetDeadline(time.Now().Add(time.Second))

		newConn, err := l.TCPListener.Accept()

		select {
		case <-l.stop:
			if err == nil {
				newConn.Close()
			}
			return nil, errStop
		default:
		}

		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		return newConn, err
	}
}

// Stop closes a given connection
func (l *Listener) Stop() {
	close(l.stop)
}
