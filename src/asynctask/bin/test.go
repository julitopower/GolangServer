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

package main

import (
	"fmt"
	at "julio/asynctask"
	"time"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting test application")
	// Create a dispatcher
	d := at.NewDispatcher(100)
	r1 := at.RequestRunner{d.Reqqueue}
	r2 := at.ResponseRunner{d.Resqueue}
	r1.Run()
	r2.Run()
	// Set the number of iterations
	args := os.Args
	var iter int
	if len(args) == 1 {
		iter = 1000
	} else {
		iter, _ =  strconv.Atoi(args[1])
	}

	// Use the dispatcher to dispatch messages
	for i := 0 ; i < iter ; i++ {
		var cmd at.Command
		if i % 2 == 0 {
			cmd = at.NewCommandTest("Request")
		} else {
			cmd = at.NewCommandTest("Response")
		}
		d.Dispatch(cmd)
	}
	time.Sleep(time.Millisecond * 20000)
}
