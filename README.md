# Whatlanggo

[![Build Status](https://travis-ci.org/abadojack/whatlanggo.svg?branch=master)](https://travis-ci.org/abadojack/whatlanggo)  [![Go Report Card](https://goreportcard.com/badge/github.com/abadojack/whatlanggo)](https://goreportcard.com/report/github.com/abadojack/whatlanggo)  [![GoDoc](https://godoc.org/github.com/abadojack/whatlanggo?status.png)](https://godoc.org/github.com/abadojack/whatlanggo)

Natural language detection for Go.
## Features
* Supports [84 languages](https://github.com/abadojack/whatlanggo/blob/master/SUPPORTED_LANGUAGES.md)
* 100% written in Go
* No external dependencies
* Fast
* Recognizes not only a language, but also a script (Latin, Cyrillic, etc)

## Getting started
Installation:
```sh
    go get -u github.com/abadojack/whatlanggo
```

Simple usage example:
```go
package main

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
)

func main() {
	info := whatlanggo.Detect("Foje funkcias kaj foje ne funkcias")
	fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
}
```

## Blacklisting and whitelisting
```go
import "github.com/abadojack/whatlanggo"

//Blacklist
options := whatlanggo.Options{
	Blacklist: map[whatlanggo.Lang]bool{
		whatlanggo.Ydd: true,
	},
}

info := whatlanggo.DetectWithOptions("האקדמיה ללשון העברית", options)

fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])

//Whitelist
options1 := whatlanggo.Options{
	Whitelist: map[whatlanggo.Lang]bool{
		whatlanggo.Epo: true,
		whatlanggo.Ukr: true,
	},
}

info = whatlanggo.DetectWithOptions("Mi ne scias", options1)
fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
```
For more details, please check the [documentation](https://godoc.org/github.com/abadojack/whatlanggo)

##TODO
Add reliabilty metrics in the _[Info](https://godoc.org/github.com/abadojack/whatlanggo#Info)_ struct

## License
[MIT](https://github.com/abadojack/whatlanggo/blob/master/LICENSE)

## Derivation
whatlanggo is a derivative [Franc](https://github.com/wooorm/franc) (JavaScript, MIT) by [Titus Wormer](https://github.com/wooorm)

## Acknowledgements
Thanks to [greyblake](https://github.com/greyblake) Potapov Sergey for creating [whatlang-rs](https://github.com/greyblake/whatlang-rs) from where I got the idea and logic.
