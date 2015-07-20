package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cznic/sortutil"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)

	panicCheck(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	scanner := bufio.NewScanner(reader)

	wordMap := make(map[string][]string)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		baseWord := strings.ToLower(scanner.Text())
		sortedWord := sortWord(baseWord)
		if !contains(baseWord, wordMap[sortedWord]) {
			wordMap[sortedWord] = append(wordMap[sortedWord], baseWord)
		}
	}

	numberOfAnagrams := 0
	for _, v := range wordMap {
		if len(v) > 1 {
			fmt.Println(v)
			numberOfAnagrams += len(v)
		}
	}
	fmt.Printf("total: %d\n", numberOfAnagrams)
}

func contains(s string, m []string) bool {
	for _, a := range m {
		if a == s {
			return true
		}
	}
	return false
}

func sortWord(word string) string {
	var runeSlice sortutil.RuneSlice
	runeSlice = []rune(word)
	runeSlice.Sort()
	return string(runeSlice)
}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}
