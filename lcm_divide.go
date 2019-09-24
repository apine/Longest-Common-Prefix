package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func preProcess() (*bufio.Scanner, *os.File) {
	// expect input file name from command line parameter
	if len(os.Args) != 2 {
		fmt.Println("No file input indicated!")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(f), f
}

func postProcess(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func fileLineProcess(str string) (string, []string) {
		r := regexp.MustCompile(`[\"\[\]]`)
		str = r.ReplaceAllString(str, "")
		pureStrings := strings.Split(str, ",")
		return pureStrings[len(pureStrings)-1], pureStrings[:len(pureStrings)-1]
}

func main() {
	s, f := preProcess()
	defer postProcess(f)

	for s.Scan() {
		answer, inputs := fileLineProcess(s.Text())

		result := ""
		result = longestCommonPrefixDivideAndConquer(inputs)
		fmt.Println("Answer=", answer)
		fmt.Println("Result=", result)
	}

	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func longestCommonPrefixDivideAndConquer(strs []string) string {
	// basic condition check
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	return divide(strs, 0, len(strs) - 1)
}

func divide(strs []string, left, right int) string {
	if left == right {
		return strs[left]
	} else {
		mid := (left + right) / 2
		lcpLeft := divide(strs, left, mid)
		lcpRight := divide(strs, mid + 1, right)
		return conquer(lcpLeft, lcpRight)
	}
}

func conquer(left, right string) string {
    mx := 0
    for {
		 if mx >= len(left) || mx >= len(right) ||
		 left[mx] != right[mx] {
			 return string(left[:mx])
		 }
        mx++
    }
    return string(left[:mx])
}
