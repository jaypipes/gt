# `gt` - Go Terminal UI library

[![Go Reference](https://pkg.go.dev/badge/github.com/jaypipes/gt.svg)](https://pkg.go.dev/github.com/jaypipes/ghw)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaypipes/gt)](https://goreportcard.com/report/github.com/jaypipes/ghw)
[![Build Status](https://github.com/jaypipes/gt/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/jaypipes/ghw/actions)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

`gt` is a Go library providing a simple framework for constructing Terminal UI
(TUI) applications.

## Design Principles

* Intuitive interface that does NOT use the Elm Architecture

    `gt` was borne out of frustration using the Bubbletea library, which uses
    the Elm Architecture. Its use in UI frameworks like Bubbletea and React is
    lauded as an approach to cleanly separate and manage state in client-side
    applications, but I personally find it overly complicated and obtuse, difficult
    to debug and impossible to explain.

* Well-documented code and plenty of example code

    The code itself should be well-documented with lots of usage examples.

* Interfaces should be consistent across modules

    Each module in the library should be structured in a consistent fashion,
    and the structs returned by various library functions should have
    consistent attribute and method names.
