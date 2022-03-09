<div align="center">
<img src="./glugger.png" />
<h1>Glog</h1>
<h4>Glog glugs the log files in a concurrent and thread-safe way.</h4>
<h4>Simple and easy to implement interfaces to log fast and efficiently</h4>

[![Go Report Card](https://goreportcard.com/badge/github.com/ric-v/glog)](https://goreportcard.com/report/github.com/ric-v/glog)
[![CodeFactor](https://www.codefactor.io/repository/github/ric-v/glog/badge)](https://www.codefactor.io/repository/github/ric-v/glog)
[![Maintained](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://img.shields.io/badge/Maintained%3F-yes-green.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ric-v_glog&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ric-v_glog)

</div>

---

## Features

- Thread-safe logger
- Log formatting support
- Logging in unstructured or structured format (JSON)

## Usage

```bash
go get github.com/ric-v/glog
```

simple logger to stdout

```go
package main

import "github.com/ric-v/glog"

func main() {
    defer glog.Cleanup()

    // log the message to the default concurrent logger
    glog.Info("Hello World")
}
```

json logger to file

```go
package main

import "github.com/ric-v/glog"

func main() {
    logger := glog.JSONGlogger("glogger.log")
    defer logger.Cleanup()

    // log the message to custom json logger
    logger.Info("", "Hello", "World")
}
```

## Examples

visit examples [here](https://github.com/ric-v/glog/tree/main/examples)

_Code.Share.Prosper_
