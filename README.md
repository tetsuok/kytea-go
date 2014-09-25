kytea-go [![Build Status](https://travis-ci.org/tetsuok/kytea-go.svg?branch=master)](https://travis-ci.org/tetsuok/kytea-go)
========

Go package which provides access to KyTea.

### Requirement

- A fork of KyTea: https://github.com/tetsuok/kytea

This Go package uses cgo, which means you need to install KyTea's headers and library `libkytea.so`
under the "correct" path in the sense that the compiler can find the headers and library.
