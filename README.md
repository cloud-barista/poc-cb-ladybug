# CB-Ladybug (POC) :beetle:
> Multi-Cloud Application Management Framework

[![Go Report Card](https://goreportcard.com/badge/github.com/cloud-barista/poc-cb-ladybug)](https://goreportcard.com/report/github.com/cloud-barista/poc-cb-ladybug)
[![Build](https://img.shields.io/github/workflow/status/cloud-barista/poc-cb-ladybug/Build%20amd64%20container%20image)](https://github.com/cloud-barista/poc-cb-ladybug/actions?query=workflow%3A%22Build+amd64+container+image%22)
[![Top Language](https://img.shields.io/github/languages/top/cloud-barista/poc-cb-ladybug)](https://github.com/cloud-barista/poc-cb-ladybug/search?l=go)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cloud-barista/poc-cb-ladybug?label=go.mod)](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/go.mod)
[![Repo Size](https://img.shields.io/github/repo-size/cloud-barista/poc-cb-ladybug)](#)
[![GoDoc](https://godoc.org/github.com/cloud-barista/poc-cb-ladybug?status.svg)](https://pkg.go.dev/github.com/cloud-barista/poc-cb-ladybug@master)
[![Release Version](https://img.shields.io/github/v/release/cloud-barista/poc-cb-ladybug?color=blue)](https://github.com/cloud-barista/poc-cb-ladybug/releases/latest)
[![License](https://img.shields.io/github/license/cloud-barista/poc-cb-ladybug?color=blue)](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/LICENSE)

```
[NOTE]
CB-Ladybug (POC) is currently under development. (the latest release is none) 
So, we do not recommend using the current release in production.
Please note that the functionalities of CB-Ladybug are not stable and secure yet.
If you have any difficulties in using CB-Ladybug, please let us know.
(Open an issue or Join the cloud-barista Slack)
```

## Getting started

### Preparation

* Golang 1.16.+ ([Download and install](https://golang.org/doc/install))

### Dependencies

* CB-MCKS [v0.4.3](https://github.com/cloud-barista/cb-mcks/releases/tag/v0.4.3)
* CB-Tumblebug [v0.4.7](https://github.com/cloud-barista/cb-tumblebug/releases/tag/v0.4.7)
* CB-Spider [v0.4.10](https://github.com/cloud-barista/cb-spider/releases/tag/v0.4.10)


### Clone

```
$ git clone https://github.com/cloud-barista/poc-cb-ladybug.git
$ cd poc-cb-ladybug
$ go get -v all
```

### Run 

```
$ export CBLOG_ROOT="$(pwd)"
$ export CBSTORE_ROOT="$(pwd)"
$ go run cmd/cb-ladybug/main.go
```

### Build and Execute

```
$ go build -o cb-ladybug cmd/cb-ladybug/main.go
```

```
$ export CBLOG_ROOT="$(pwd)"
$ export CBSTORE_ROOT="$(pwd)"
$ nohup ./cb-ladybug & > /dev/null
```

### Test

```
$ ./scripts/get-health.sh

[INFO]
- Ladybug URL is 'http://localhost:1592/ladybug'

------------------------------------------------------------------------------
cloud-barista cb-ladybug is alived
```

### API documentation

* Under construction

## Documents

* Under construction


## Contribution
Learn how to start contribution on the [Contributing Guideline](https://github.com/cloud-barista/docs/tree/master/contributing) and [Style Guideline](https://github.com/cloud-barista/poc-cb-ladybug/blob/master/STYLE_GUIDE.md)

