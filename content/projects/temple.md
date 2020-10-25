---
title: "temple"
version: "0.1.0"
summary: "A Go template CLI compiler and function library"
link: "https://github.com/mattmeyers/temple"
date: 2020-10-22T02:13:05-04:00
draft: false
weight: 200
---

`temple` focuses on simplifying the template development workflow by providing a simple CLI tool for parsing and executing Go templates. With the file watching feature, a developer can edit their template or data files and immediately see their changes in the built template. This tool was originally created to help coworkers work on Go templates without requiring a full development environment.

In addition to the CLI tool, `temple` is a growing library of template functions designed to improve the templating experience. This library is designed to complement the standard library's template functions and can easily integrate with the user's custom functions. The CLI tool ships with the `temple` library built in. However, the CLI functionality is an exported package. This package can be used to easily recompile `temple` with whatever template functions the user wants.