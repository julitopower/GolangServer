// MIT License
//
// Copyright (c) 2017 Julio Delgado
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// 	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package asynctask

import (
	"fmt"
)

// A Dispacher managers a number of queues for the asynchronous
// processing of tasks. In this implementatino there are 3 such
// queues:
// * Queue for Commands that need to be dispached
// * Requests for Commands of type Request that are ready to be executed
// * Response for Commands of type Response that are ready to be executed
type Dispatcher struct {
	inqueue  chan Command
	Reqqueue chan Command
	Resqueue chan Command
}

// Constructor for dynamic instances of Dispacher. This is the only
// construction method that should be used, otherwise the dispatch
// loop won't get initiated.
func NewDispatcher(buffersize int32) *Dispatcher {
	d := new(Dispatcher)
	d.inqueue = make(chan Command, buffersize)
	d.Reqqueue = make(chan Command, buffersize)
	d.Resqueue = make(chan Command, buffersize)	
	go d.loop()
	return d

}

// Method invoked by clients to send Commands to the
// Dispatcher. Commands are dispatched asynchronously.
func (d *Dispatcher) Dispatch(cmd Command) {
	go func() {
		select {
		case d.inqueue <- cmd:
			fmt.Printf("Enqueued Command: %v\n", cmd)			
		default:
			fmt.Println("Channel is full, can't send Cmd")
		}
	}()
}

// Internal dispatching loop. This is ver simple example that matches
// the Command name, and dispatches to the corresponding queue
func (d *Dispatcher) loop() {
	for {
		cmd := <- d.inqueue
		fmt.Printf("Dequeued Command: %v\n", cmd)
		switch cmdTy := cmd.Type().Name; cmdTy  {
		case "Request":
			d.Reqqueue <- cmd
		case "Response":
			d.Resqueue <- cmd
		default:
			fmt.Println("Unkown command type", cmdTy)
		}
	}
}
