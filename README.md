# Overview

The purpose of this module is to demonstrate implementation of data structure using Go.

NOTE: The implementation in this module are mostly based on Go 1.18

## Generics and constraints

In the package `model` you will find an example, demonstrating the use of Go generics.

The interface named [NumericType](./internal/model/model.go) is an example of a constraints.

## Binary tree structure

A binary tree is a tree data structure in which each node has at most two children. Please refer to [wikipedia](https://en.wikipedia.org/wiki/Binary_tree) for more information.

An implementation of a binary is in this package [bintree](./internal/bintree).

## Trie data structure

A Trie is a data structure for storing strings. To learn more about the inner workings of this data structure, I recommend this video titled [Trie Data Structure](https://www.youtube.com/watch?v=-urNrIAQnNo).

You will find a Go implementation in this package [trie](./internal/trie).

## Disclaimers

The packages in this module are intended for educational purpose only.