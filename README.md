go-apns-server
==============

A quick and dirty script to mock the new HTTP/2 interface to Apple's push gateway.
Most requests to this server will return a successful response, however, some of
them will return a "BadDeviceToken" or an "Unregistered" from time to time.

# Usage
```
  -cert string
    	The path to the certificate (default "cert.pem")
  -key string
    	The path to the certificate key (default "key.pem")
  -port int
    	The port for the mock server to listen to (default 8443)
```

# Getting started
## Install the Go programming language
See [here](https://golang.org/doc/install).

## Install the HTTP/2 library
### Set your GOPATH
By default, Go installs itself to /usr/local/go on a Mac:
```
$ export GOPATH=/usr/local/go/
```

### Install the library
After exporting the GOPATH variable, install the HTTP/2 library:
```
$ go get golang.org/x/net/http2
```

## Generate your self signed certificate
Create a self signed certificate using one of Go's scripts:
```
$ go run /usr/local/go/src/crypto/tls/generate_cert.go --host clevertap.com
```

## Start the server
If you created the certificate elsewhere, you'll need to use the
options -cert and -key to specify where the certificate and it's key are located
```
$ go run go-apns-server.go
```

The default port is 8443, which can be changed by passing the -port argument

# License
Licensed under the [New 3-Clause BSD License](http://opensource.org/licenses/BSD-3-Clause).

```
Copyright (c) 2016, CleverTap
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of CleverTap nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
