# Tiny

[![Build Status](https://travis-ci.org/zackkitzmiller/tiny-go.png?branch=master)](https://travis-ci.org/zackkitzmiller/tiny-go)

A reversible base62 ID obfuscater

## Authors

Originally by Jacob DeHart, with Ruby and Python ports by Kyle Bragger

Go port by [Zack Kitzmiller](https://github.com/zackkitzmiller).

PHP port currently maintained by [Zack Kitzmiller](https://github.com/zackkitzmiller).

## Installation

Install via `go get`

## Usage

```go
completelyRandomSeed := "5SX0TEjkR1mLOw8Gvq2VyJxIFhgCAYidrclDWaM3so9bfzZpuUenKtP74QNH6B"
tiny := NewTiny(completelyRandomSeed)

fmt.Println(tiny.To(5))
// E

fmt.Println(tiny.From("E"))
// 5

```

More information and background can be found [here](http://b.z19r.com/post/tiny-php).
