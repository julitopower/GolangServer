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
	"time"
	"io/ioutil"
	"net/http"
)

// A Runner is an entity able to execute commands asynchornously
type Runner interface {
	Run()
}

// Test RequestRunner
type RequestRunner struct {
	Reqqueue chan Command
}

// Test ResponseRunner
type ResponseRunner struct {
	Resqueue chan Command
}

// Initiates the execution loop in a separate go rutine
func (r *RequestRunner) Run() {
	go func() {
		for {
			// Grab command from channel
			cmd := <-r.Reqqueue
			fmt.Println("Running Request Cmd: ", cmd)
			resp, _ := http.Get("http://localhost:9090")
			payload, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(len(payload), string(payload))
			resp.Body.Close()
		}
	}()
}

// Initiates the execution loop in a separate go rutine
func (r *ResponseRunner) Run() {
	go func() {
		for {
			// Grab command from channel
			cmd := <-r.Resqueue
			fmt.Println("Running Response Cmd: ", cmd)
			time.Sleep(time.Millisecond * 40)
		}
	}()
}
