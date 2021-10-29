package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func mustAtoi(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func createRegexForRules(rules []string, override bool) string {
	dict := make(map[string][]string)

	for _, r := range rules {
		vec := strings.Split(r, ":")
		//id := mustAtoi(vec[0])
		id := vec[0]
		subs := vec[1]
		subs = strings.Replace(subs, "\"", "", -1)
		subs = strings.TrimSpace(subs)

		dict[id] = strings.Split(subs, " ")
	}

	if override {
		repeat := 20

		dict["8"] = []string{"42", "+"}

		ans := []string{"42", "31"}
		for i := 2; i < repeat; i++ {
			ans = append(ans, "|")
			ans = append(ans, "42")
			ans = append(ans, fmt.Sprintf("{%d}", i))
			ans = append(ans, "31")
			ans = append(ans, fmt.Sprintf("{%d}", i))
		}
		dict["11"] = ans
		fmt.Println(dict["8"])
		fmt.Println(dict["11"])
	}
	index := "0"
	currentString := []string{"^"}
	for _, c := range dict[index] {
		currentString = append(currentString, c)
	}
	replaced := true
	for replaced {
		replaced = false
		for i := 0; ; i++ {
			if i >= len(currentString) {
				break
			}
			repl := currentString[i]
			v, exists := dict[repl]
			if exists {
				newString := make([]string, len(currentString[:i]))
				copy(newString, currentString[:i])

				newString = append(newString, "(")
				for _, vj := range v {
					newString = append(newString, vj)
				}
				newString = append(newString, ")")
				for _, vj := range currentString[i+1:] {
					newString = append(newString, vj)
				}

				i += len(v) + 1
				currentString = newString
				replaced = true
			} else {
			}
		}
	}

	currentString = append(currentString, "$")
	reg := strings.Join(currentString, "")

	fmt.Println(reg)
	return reg
}

func IsValid(content string, override bool) int {

	rules := make([]string, 0)
	words := make([]string, 0)

	vec := strings.Split(content, "\n")
	state := 0
	for _, l := range vec {
		if l == "" {
			state = 1
			continue
		}

		if state == 0 {
			rules = append(rules, l)
		}
		if state == 1 {
			words = append(words, l)
		}
	}


	reg := createRegexForRules(rules, override)

	re := regexp.MustCompile(reg)
	sum := 0
	for _, w := range words {
		match := re.MatchString(w)
		if match {
			sum++
		}

	}

	return sum
}


func main() {

	content, err := ioutil.ReadFile("./input19.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := IsValid(string(content), false)

	fmt.Println("Part a", ans)

	ans = IsValid(string(content), true)

	fmt.Println("Part b", ans)

}
