# stack2struct
[![Build Status](https://travis-ci.org/kaeuferportal/stack2struct.svg?branch=master)](https://travis-ci.org/kaeuferportal/stack2struct)
[![Coverage](http://gocover.io/_badge/github.com/kaeuferportal/stack2struct)](http://gocover.io/github.com/kaeuferportal/stack2struct)
[![GoDoc](https://godoc.org/github.com/kaeuferportal/stack2struct?status.svg)](http://godoc.org/github.com/kaeuferportal/stack2struct)

stack2struct parses raw golang stack traces ([]byte) to a slice of well formated structs.

## Getting Started

### Installation
```
  $ go get github.com/kaeuferportal/stack2struct
```
### Usage
```
package main

import (
  "encoding/json"
  "fmt"
  "runtime"

  "github.com/kaeuferportal/stack2struct"
)

type stackTraceElement struct {
  LineNumber int    `json:"lineNumber"`
  ClassName  string `json:"className"`
  FileName   string `json:"fileName"`
  MethodName string `json:"methodName"`
}

type stackTrace []stackTraceElement

func (s *stackTrace) AddEntry(lineNumber int, packageName, fileName, methodName string) {
  *s = append(*s, stackTraceElement{lineNumber, packageName, fileName, methodName})
}

func main() {

  rawStackTrace := make([]byte, 1<<16)
  rawStackTrace = rawStackTrace[:runtime.Stack(rawStackTrace, false)]

  stack := make(stackTrace, 0, 0)
  stack2struct.Parse(rawStackTrace, &stack)

  enc, _ := json.MarshalIndent(stack, "", "\t")
  fmt.Println(string(enc))

}

```

should print
```
[
  {
    "lineNumber": 27,
    "className": "main",
    "fileName": "test.go",
    "methodName": "main()"
  }
]
```


## Bugs and feature requests

Have a bug or a feature request? Please first check the list of
[issues](https://github.com/kaeuferportal/stack2struct/issues).

If your problem or idea is not addressed yet, [please open a new
issue](https://github.com/kaeuferportal/stack2struct/issues/new), or contact us at
[oss@kaeuferportal.de](mailto:oss@kaeuferportal.de).
