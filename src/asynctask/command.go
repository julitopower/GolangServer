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

type Type struct {
	Name string
	Id   int
}

type Command interface {
	Type() Type
	Payload() interface{}
}

type CommandTest struct {
	ty Type
	payload string
}

func (d *CommandTest) Type() Type {
	return d.ty
}

func (d *CommandTest) Payload() interface{} {
	return d.payload
}

func NewCommandTest(name string) *CommandTest {
	cmd := new(CommandTest)
	cmd.payload = name
	cmd.ty.Name = name
	cmd.ty.Id = 1
	return cmd
}
