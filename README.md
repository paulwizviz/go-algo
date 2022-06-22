# Overview

A collection of executable data models and algorithms implemented in Go.

Please install Go version 1.18 or later to work with the executable examples.

## Project structure

The Go codes are organised in these folders:

* `build` -- scripts used to create apps and containers
* `cmd` -- source codes for building executables
* `internal` - share packages to for `cmd`
• `scripts` - bash scripts to help you build and execute apps.

## Topics

The topics covered in this projects are:

### Applications

* [BIFID cipher](./docs/bifid.md)
* [Palindrom detector](./cmd/palindrome/) is an application that detects a word or number that reads the same forward or backward. The application includes two ways of determining if a word is a palindrome. The two approaches are benchmarked.

### Data models

* [Table models](./docs/table.md)
* [Tree structures](./docs/treemdl.md)
* [Roman numerals](./docs/romans.md)

## Disclaimers

* The packages in this project are intended for educational purpose only.
* This package is constantly updated and items may be removed and modified without warning.

## Copyright

Unless specificed copyright in this project are assigned as follows.

Copyright 2022 Paul Sitoh

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.