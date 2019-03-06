# Whatlanggo

[![Build Status](https://travis-ci.org/abadojack/whatlanggo.svg?branch=master)](https://travis-ci.org/abadojack/whatlanggo)  [![Go Report Card](https://goreportcard.com/badge/github.com/abadojack/whatlanggo)](https://goreportcard.com/report/github.com/abadojack/whatlanggo)  [![GoDoc](https://godoc.org/github.com/abadojack/whatlanggo?status.png)](https://godoc.org/github.com/abadojack/whatlanggo) [![Coverage Status](https://coveralls.io/repos/github/abadojack/whatlanggo/badge.svg)](https://coveralls.io/github/abadojack/whatlanggo)

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
	fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)
}
```

## Blacklisting and whitelisting
```go
package main

import (
	"fmt"

	"github.com/abadojack/whatlanggo"
)

func main() {
	//Blacklist
	options := whatlanggo.Options{
		Blacklist: map[whatlanggo.Lang]bool{
			whatlanggo.Ydd: true,
		},
	}

	info := whatlanggo.DetectWithOptions("האקדמיה ללשון העברית", options)

	fmt.Println("Language:", info.Lang.String(), "Script:", whatlanggo.Scripts[info.Script])

	//Whitelist
	options1 := whatlanggo.Options{
		Whitelist: map[whatlanggo.Lang]bool{
			whatlanggo.Epo: true,
			whatlanggo.Ukr: true,
		},
	}

	info = whatlanggo.DetectWithOptions("Mi ne scias", options1)
	fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script])
}
```
For more details, please check the [documentation](https://godoc.org/github.com/abadojack/whatlanggo).

## Requirements
Go 1.8 or higher

## How does it work?

### How does the language recognition work?

The algorithm is based on the trigram language models, which is a particular case of n-grams.
To understand the idea, please check the original whitepaper [Cavnar and Trenkle '94: N-Gram-Based Text Categorization'](https://www.researchgate.net/publication/2375544_N-Gram-Based_Text_Categorization).

### How _IsReliable_ calculated?

It is based on the following factors:
* How many unique trigrams are in the given text
* How big is the difference between the first and the second(not returned) detected languages? This metric is called `rate` in the code base.

Therefore, it can be presented as 2d space with threshold functions, that splits it into "Reliable" and "Not reliable" areas.
This function is a hyperbola and it looks like the following one:

<img alt="Language recognition whatlang rust" src="https://raw.githubusercontent.com/abadojack/whatlanggo/master/images/whatlang_is_reliable.png" width="450" height="300" />

For more details, please check a blog article [Introduction to Rust Whatlang Library and Natural Language Identification Algorithms](https://www.greyblake.com/blog/2017-07-30-introduction-to-rust-whatlang-library-and-natural-language-identification-algorithms/).

## License
[MIT](https://github.com/abadojack/whatlanggo/blob/master/LICENSE)

## Derivation
whatlanggo is a derivative of [Franc](https://github.com/wooorm/franc) (JavaScript, MIT) by [Titus Wormer](https://github.com/wooorm).

## Acknowledgements
Thanks to [greyblake](https://github.com/greyblake) (Potapov Sergey) for creating [whatlang-rs](https://github.com/greyblake/whatlang-rs) from where I got the idea and algorithms.
