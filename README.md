# ForSAM Logger Formatters

[![GoDoc](https://godoc.org/github.com/forsam-education/loggerformatters?status.svg)](https://godoc.org/github.com/forsam-education/loggerformatters)
[![Go Report Card](https://goreportcard.com/badge/github.com/forsam-education/loggerformatters)](https://goreportcard.com/report/github.com/forsam-education/loggerformatters)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

These logger formatters are meant to be used in the ForSAM softwares, but are leaved open source to allow users of [simplelogger](https://github.com/forsam-education/simplelogger) to have some examples, and why not inspire from it.

## Quality assurance

To fix the basics of code format, you can run `go fmt`.

For a bit more advanced code style checks, you can run `golint $(go list ./... | grep -v /vendor/)`. You'll have to run `go get -u golang.org/x/lint/golint` before.

## Available Formatters

### JSON Formatter

**File**: `json_formatter.go`

This formatter is used in every services that needs to log in ElasticSearch, through CloudWatch or LogStash.
