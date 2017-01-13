# stemmer
[![Build Status](https://travis-ci.org/abadojack/stemmer.svg?branch=master)](https://travis-ci.org/abadojack/stemmer) [![GoDoc](https://godoc.org/github.com/abadojack/stemmer?status.png)](http://godoc.org/github.com/abadojack/stemmer)

Stemmer is a light and aggressive stemmer for Esperanto. It is still under active development
and a lot of things are likely to change.

## Install
```sh
go get -u github.com/abadojack/stemmer
```

## Usage
```go
import "github.com/abadojack/stemmer"
```

### Light stemming
```go
s := stemmer.Stem("ludas")
fmt.Println(s)  //ludi
```

### Aggressive stemming
```go
s := stemmer.StemAggressive("ludas")
fmt.Println(s)  //lud
```
