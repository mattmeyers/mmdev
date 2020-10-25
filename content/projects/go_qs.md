---
title: "go-qs"
version: "1.1.0"
summary: "An extended query string parser"
link: "https://github.com/mattmeyers/go-qs"
docs: "https://pkg.go.dev/github.com/mattmeyers/go-qs/qs"
date: 2020-10-22T02:25:08-04:00
draft: false
weight: 300
---

Inspired by NodeJS's [`qs` library](https://www.npmjs.com/package/qs), `go-qs` provides the ability to parse and build nested query strings. Transferring complex data structures through a query string can be difficult and Go's standard library parser only allows for basic key-value parsing. This library further parses query strings with nested keys. For example, the query string `key[subkey]=value` will be parsed into the map structure

```go
{
    "key": {
        "subkey": "value"
    }
}
```

This provides more flexibility when defining REST API endpoints as it allows clients to easily provide detailed data with their GET requests. This library has already been used in production to parse query strings containing filter data of the form `filter[operator]=value` e.g. `height[gt]=5.5`.