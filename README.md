# Tiny

A reversible base62 ID obfuscater

## Authors

Originally by Jacob DeHart, with Ruby and Python ports by Kyle Bragger

Now maintained by [Zack Kitzmiller](https://github.com/zackkitzmiller).

## Installation

Install via `go get`

## Usage

```go
tiny := NewTiny("5SX0TEjkR1mLOw8Gvq2VyJxIFhgCAYidrclDWaM3so9bfzZpuUenKtP74QNH6B")

fmt.Println(tiny.To(5))
// E

fmt.Println(tiny.From("E"))
// 5

```

More information and background can be found [here](http://b.z19r.com/post/tiny-php).
