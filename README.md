gcd (Google Cloud DataStore)
-

[![Go Report Card](https://goreportcard.com/badge/github.com/thepkg/gcd)](https://goreportcard.com/report/github.com/thepkg/gcd)
[![godoc](https://godoc.org/github.com/thepkg/gcd?status.svg)](https://godoc.org/github.com/thepkg/gcd)

Package `gcd` provides helpful functions to work with Google Cloud DataStore.

## Installation

`go get -u github.com/thepkg/gcd`

## Usage

````
import "github.com/thepkg/gcd"

ctx := appengine.NewContext(r)
keys := gcd.GetKeys(ctx, "MyKind")
gcd.DropKind(ctx, "MyKind")
````
