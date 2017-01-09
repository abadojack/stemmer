# stemmer

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


#NOTE: As of Tue Jan 10 04:34:15 IST 2017, not all test cases pass.
