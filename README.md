gcd (Google Cloud DataStore)
-

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
