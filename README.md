# goparser

[![go report card](https://goreportcard.com/badge/github.com/jjkoh95/goparser "go report card")](https://goreportcard.com/report/github.com/jjkoh95/goparser)
[![GoDoc](https://godoc.org/github.com/jjkoh95/goparser?status.svg)](https://godoc.org/github.com/jjkoh95/goparser)
[![Actions Status](https://github.com/jjkoh95/goparser/workflows/Go/badge.svg)](https://github.com/jjkoh95/goparser/actions)

## Introduction
A smart parser in go inspired by Python Pandas package.

This is mainly tested and used in moving files between places (spreadsheets and buckets) and between csv format and json format.

While opinionated language is awesome in many ways, this package is initiated to provide a little more flexibility.

## Getting started
- go get github.com/jjkoh95/goparser

```go
    // var rows [][]string
    var typeParser goparser.TypeParser
    typeParser.GetSmartParser(rows) // infer type of each column
    // typeParser.GetParserFromYAML(YAMLParserFile)
    result := goparser.ParseRowToList(rows, &typeParser)
    // result is []map[string]interface{}

    // essentially turn structure from
    // [{"col1", "col2"},
    //  {   1  ,    2  },
    //  {  10  ,   20  }]
    // to
    // [{"col1":  1, "col2":  2},
    //  {"col1": 10, "col2": 20}]
```

## License
[MIT License](https://github.com/jjkoh95/goparser/blob/master/LICENSE)
