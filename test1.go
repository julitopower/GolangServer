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
	net "net/http"
)

type handler struct {
	x, y int;
}

func (h handler) ServeHTTP(w net.ResponseWriter, r *net.Request) {
	// Serve data with a method of http.ResponseWriter.
	w.Write([]byte("You learned Go in Y minutes!"))
}

func main() {
	fmt.Println("Hi there");
	var x int;
	x,_ = print_param(34);
	fmt.Println(x);

	// Let's start a web server
	net.ListenAndServe(":8192", handler{});
}

func print_param(x int) (int, int) {
	fmt.Println(x);
	return x, x;
}

