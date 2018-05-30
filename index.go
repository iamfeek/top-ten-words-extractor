package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
	"regexp"
)

type kv struct {
	Key   string
	Value int
}

func splitToArray(bodyOfText string) []string {
	arr := strings.Split(bodyOfText, " ")
	return arr
}

func stripAllButLetters(text string) string {
	wordRegexp := regexp.MustCompile("[^a-zA-Z0-9-']+")

	word := wordRegexp.ReplaceAllString(text, "")
	return word
}

func getTopTen(arr []string) []kv {

	var slicable []kv
	m := make(map[string]int)

	for _, text := range arr {
		stripped := stripAllButLetters(text)
		m[stripped] = m[stripped] + 1
	}

	for key, value := range m {
		if len( key ) > 0{
			slicable = append(slicable, kv{key, value})
		}

	}

	sort.Slice(slicable, func(i, j int) bool {
		return slicable[i].Value > slicable[j].Value
	})

	if len(slicable) < 10 {
		return slicable
	} else {
		return slicable[0:10]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	arrayOfStrings := splitToArray(text)
	topTen := getTopTen(arrayOfStrings)

	fmt.Println(topTen)
}
