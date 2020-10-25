---
title: "qb"
version: "WIP"
summary: "A simple SQL query builder"
link: "https://github.com/mattmeyers/qb"
date: 2020-10-22T21:34:16-04:00
draft: false
weight: 500
---

`qb` is a query builder aimed primarily at PostgreSQL. It is able to build complicated queries by recursively constructing queries from objects that satisfy the `Builder` interface. 

```go
type Builder interface {
    Build() (string, []interface{}, error)
}
```

This library focuses on providing a straightforward DSL that builds correct and safe queries. Calling the `Build()` method returns a parameterized query that can run without fear of SQL injection. `qb` is designed to easily integrate into existing projects. It functions without any knowledge of the database being used, meaning `qb` functions can simply replace current query definitions.