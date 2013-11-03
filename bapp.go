package bapp

import (
	"errors"
	"fmt"
	"github.com/toqueteos/webbrowser"
	"net"
	"net/http"
	"strings"
)

// ErrNoFreePort is returned by NewBapp when bapp could not find a free port.
// This should be extremely unlikely to happen.
// go.bapp searched for a free port between 18000 and 18999.
var ErrNoFreePort = errors.New("could not find a free port")

// Bapp contains the http server and browser information
type Bapp struct {
	port     uint16
	listener net.Listener
	server   *http.Server

	stopChan chan bool
	closed   bool
}

// NewBapp creates a new bapp with server
func NewBapp() (*Bapp, error) {
	var err error

	// new Bapp instance
	b := &Bapp{
		port:     18000,
		stopChan: make(chan bool),
	}

	// setup listener
	for {
		// listen on b.port
		b.listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", b.port))

		// break from for loop when net.Listen(..) was successfull
		if err != nil {
			// check if error is an "address already in use" error
			//  (which is expected when one or more bapp applications are running)
			if strings.Contains(err.Error(), "address already in use") {
				// address in use, thats okay.. we'll try the next port
				// increment port number
				b.port++

				// return error when no free port could be found between 18000 and 18999
				if b.port > 18999 {
					return nil, ErrNoFreePort
				}

				// next loop
				continue
			}

			// other error, return with error
			return nil, err
		}

		// no error
		break
	}

	// new server instance
	b.server = &http.Server{
		Handler: http.DefaultServeMux,
	}

	// start http server on listener in goroutine
	go func() {
		err := b.server.Serve(b.listener)
		if err != nil && !b.closed {
			panic(err)
		}
	}()

	// start bapp lifecycle
	go b.live()

	// all done
	return b, nil
}

func (b *Bapp) live() {
	b.stopChan <- true
	close(b.stopChan)
	b.closed = true
	b.listener.Close() //++ should returned error be checked?
}

// SetHandler set's the handler on the http server
// By default the http.DefaultServeMux is used.
func (b *Bapp) SetHandler(handler http.Handler) {
	// b.server.handler = handler
	b.server.Handler = handler
}

// Close stops the bapp instance.
// false is returned when the bapp was already stopped
func (b *Bapp) Close() bool {
	_, ok := <-b.stopChan
	return ok
}

// Open opens the application in browser
func (b *Bapp) Open() {
	b.OpenPath("/")
}

// OpenPath opens the application in browser with given path
//++ now uses github.com/toqueteos/webbrowser
//++ might use https://bitbucket.org/tebeka/go-wise/src/d8db9bf5c4d1/desktop.go?at=default
//++ preferably write own implementation specific for bapp
//++ (check if certain browsers are installed, specify browser preference on bapp as in: `bapp.PreferBrowser(Browser ... )`  )
func (b *Bapp) OpenPath(path string) {
	webbrowser.Open(fmt.Sprintf("http://localhost:%d%s", b.port, path))
}
