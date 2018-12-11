# mdreminder

[![Build Status](https://travis-ci.org/AaronWharton/mdreminder.svg?branch=master)](https://travis-ci.org/AaronWharton/mdreminder)
[![Go Report Card](https://goreportcard.com/badge/github.com/AaronWharton/mdreminder)](https://goreportcard.com/report/github.com/AaronWharton/mdreminder)
[![codecov](https://codecov.io/gh/AaronWharton/mdreminder/branch/master/graph/badge.svg)](https://codecov.io/gh/AaronWharton/mdreminder)
![DUB](https://img.shields.io/dub/l/vibe-d.svg)
[![Coverage Status](https://coveralls.io/repos/github/AaronWharton/mdreminder/badge.svg)](https://coveralls.io/github/AaronWharton/mdreminder)

## Introduction

Have you ever been bothered by plenty of invalid url links in the porject especially when you change the project structure, which contains several `*.md` files that contain many links? mdreminder is exactly the project for detecting invalid links in `*.md` files at one time, which is lightweight and user-friendly.

## Usage
**First** you should `go get` it:
  ```go
  go get -u github.com/AaronWharton/mdreminder
  ```

**Then** just run it by coding like this:
  ```go
  package main
  
  import "github.com/AaronWharton/mdreminder"
  
  func main() {
	  mdreminder.ExecuteByPath("/Users")  // you can choose the path you want to detect, note some directories may need access permission
  }
  ```

**Finally** see the result in the ouput, you can see the example [here](https://github.com/AaronWharton/mdreminder/tree/master/_example).

## Contribute
Project now is under developping, any [pull request](https://github.com/AaronWharton/mdreminder/pulls) is welcome! If you have any question you can new [issue](https://github.com/AaronWharton/mdreminder/issues).

## License
mdreminder is open-source software licensed under the [MIT License](https://github.com/AaronWharton/mdreminder/blob/master/LICENSE).
