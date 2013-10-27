package bapp

import (
	"net/http"
	"time"
)

type Bapp struct {
	port uint16

	timeoutDuration time.Duration
	waitChan        chan bool // channel is closed when bapp is closing
	pingChan        chan bool
	stopChan        chan bool
}

func NewBapp() (*Bapp, error) {
	b := &Bapp{
		waitChan: make(chan bool),
	}

	return b, nil
}

func (b *Bapp) live() {
	defer func() {
		close(b.waitChan)
		close(b.pingChan)
		close(b.stopChan)
	}()
	for {
		select {
		case <-time.After(b.timeoutDuration):
			continue
		case b.pingChan <- true:
			continue
		case b.stopChan <- true:
			return
		}
	}
}

//++ TODO: change returned booleans into errors
//++ ErrBappStopped
//++ TODO: consider s/Stop()/Close()/

//++ actually,, maybe should remove timeout stuff al together, and leave that to implementation specific stuff (boilerplate)
//++ this way, the application has good controll over itself, instead of having to manage bapp
//++ it makes bapp really small, boilerplate can be used to

// SetTimeout sets the timeout for the Bapp
// SetTimeout calls Ping before returning
// When SetTimeout returns false, the Bapp has already closed and should not be used
func (b *Bapp) SetTimeout(dur time.Duration) bool {
	b.timeoutDuration = dur
	return b.Ping()
}

func (b *Bapp) SetHandler(handler http.Handler) {
	b.server.handler = handler
}

// Ping returns true when ping was sucessfull
// When ping returns false, the Bapp has closed and should not be used anymore
func (b *Bapp) Ping() bool {
	_, ok := <-b.pingChan
	return ok
}

// Stop returns false when bapp was already stopped
func (b *Bapp) Stop() bool {
	_, ok := <-b.stopChan
	return ok
}

// Wait blocks until the Bapp is stopped
func (b *Bapp) Wait() {
	<-b.waitChan
}

// Open opens the application in browser
func (b *Bapp) Open() {
	b.OpenPath("/")
}

// OpenPath opens the application in browser with given path
func (b *Bapp) OpenPath(path string) {
	//++ xdgOpen
}
