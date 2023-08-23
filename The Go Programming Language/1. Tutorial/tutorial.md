## 1.1 Hello World
Go code is organized into packages, which are similar to libraries or modules in other languages.

Package main is special. It defines a standalone executable program, not a library. Within package main the function main is also special—it’s where execution of the program begins

Newlines following certain tokens are converted into semicolons, so where newlines are placed matters to proper parsing of Go code. 

## 1.2 Command Line Arguments
Command-line arguments are available to a program in a variable named Args that is part of the os package

The first element of os.Args, os.Args[0], is the name of the command itself; the other elements are the arguments that were presented to the program when it started execution

A slice expression of the form s[m:n] yields a slice that refers to elements m through n-1

In each iteration of the loop, range produces a pair of values: the index and the value of the element at that index. 
```
for _, arg := range os.Args[1:] {

    }
```
