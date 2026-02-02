package eventloop

import (
	"context"
	"sync"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// New returns a new initialized EventLoop.
func New(ctx context.Context) *EventLoop {
	e := &EventLoop{}
	e.ch = make(chan types.Event)
	e.stop = make(chan struct{})
	e.ctx = ctx
	return e
}

// EventLoop describes a self-contained event loop that can be embedded in
// another struct that wants to communicate with other structures via a simple
// pub-sub manner.
type EventLoop struct {
	// ch is the channel where Events are processed.
	ch chan types.Event
	// stop is the channel used to stop the EventLoop.
	stop chan struct{}
	// ctx is the context created for the EventLoop.
	ctx context.Context
	// wg is the wait group we use to wait on the loop to complete.
	wg sync.WaitGroup
}

// Start starts the EventLoop.
func (e *EventLoop) Start() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		for {
			select {
			case <-e.ctx.Done():
				// context that was passed to us was canceled or completed.
				return
			case <-e.stop:
				// we were force-stopped.
				return
			case ev := <-e.ch:
				// an event occurred, process it.
				gtlog.Debug(e.ctx, "received event %v", ev)
			}
		}
	}()
}

// Stop stops the EventLoop.
func (e *EventLoop) Stop() {
	close(e.stop)
	e.wg.Wait()
	close(e.ch)
}
