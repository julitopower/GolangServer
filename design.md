<!---
MIT License

Copyright (c) 2017 Julio Delgado

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
--->
# Introduction

The idea is to develop a web server that offers a synchronous interface to external callers, but that handles communication with N backend worker processes as asynchronous tasks.

# Concepts

* HTTPServer
* HTTPHandler
* CommandQueue
* WorkerConfig
* WorkerProcess
* Runner

# Live of a request

## Simplest design

The size of the queues provides the traffic throttling mechanism. The approach doesn't require sweeper threads to take care of waiting queues.

* HTTP connection is stablished between HTTPServer and the client
* Request is passed to HTTPHandler which transform the request into a command
* Commads are routed to one in a number of CommandQueues. Each queue containes Commands destined to particular, non-overlapping, set of workder processes
* A number of Runners pick up tasks from each queue and send them to the worker process asynchronously.

Note: By having the number of CommandQueues equal to the number of worker processes, and having only one Runner associated with each queue, we can treat the worker code as single threaded.

* A response arrive. A Handler asigns it to one of N response queues. M Runners per queue pick up tasks and send results back to callers.

Cons: This approach uses different machinery for requests and responses

## Independent command execution machinery

* HTTPC connection is stablished between HTTPSever and client
* Request is passed to HTTPHandler

