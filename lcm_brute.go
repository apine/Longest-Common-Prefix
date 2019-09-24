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
		result = longestCommonPrefixBruteForce(inputs)
		fmt.Println("Answer=", answer)
		fmt.Println("Result=", result)
	}

	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func longestCommonPrefixBruteForce(strs []string) string {
	// basic condition check
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := []byte(strs[0])
	for i := 1; i < len(strs); i++ {
		//fmt.Println("Checking ", strs[i], " with current prefix=", string(prefix))
		if len(prefix) > len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		}
		// Prefix is always <= len(strs[i])
		for j := range strs[i] {
			if j >= len(prefix) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return string(prefix)
}
