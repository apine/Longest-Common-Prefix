# Longest-Common-Prefix
Implementation of two algorithms solving longest common prefix problem
in Go.
1. Brute Force
2. Divide and conquer

The program takes 1 parameter as the input filename and output the:
1. Correct answer from the input
2. Longest common prefix of the input based on the algorithms 

# Installation
Please follow [Go’s official document](https://golang.org/doc/install) to install official binary distribution. This will install Go Tools for running/building Go code.


# Usage
There are two source code files, one for each algorithm. Both programs take
input from the command line parameter and output the result to **STDOUT**.
The input file formart please refer to the sample input, *50.txt*.

For brute force algorithm:
```
#go run lcm_brute.go 50.txt
Answer= ivpdhxwxiasdksjmkgyv
Result= ivpdhxwxiasdksjmkgyv
```
For divide and conquer algorithm:
```
#go run lcm_divide.go 50.txt
Answer= ivpdhxwxiasdksjmkgyv
Result= ivpdhxwxiasdksjmkgyv
```
