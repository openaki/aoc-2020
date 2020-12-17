package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func atoiMust(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func GetScanningRate(content string) int {

	lines := strings.Split(content, "\n")
	state := 0 // "clases"
	classRe := regexp.MustCompile(`.*: (\d+)-(\d+) or (\d+)-(\d+)`)

	type Pair struct {
		min, max int
	}

	ranges := make([]Pair, 0)
	ans := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		if l == "your ticket:" {
			state = 1 // "your"
			continue
		}

		if l == "nearby tickets:" {
			state = 2 // "nearby"
			continue
		}

		if state == 0 {
			rg := classRe.FindStringSubmatch(l)
			ranges = append(ranges, Pair{atoiMust(rg[1]), atoiMust(rg[2])})
			ranges = append(ranges, Pair{atoiMust(rg[3]), atoiMust(rg[4])})
		}

		if state == 2 {

			nums := strings.Split(l, ",")

			for _, n := range nums {
				if n == "" {
					continue
				}

				no := atoiMust(n)

				found := false

				for _, r := range ranges {
					if no >= r.min && no <= r.max {
						found = true
						break
					}
				}
				if !found {
					ans += no
				}
			}
		}
	}

	return ans
}

type Rules struct {
	name string
	min1, max1 int
	min2, max2 int
}


func isValueValidForRule(v int, rule Rules) bool {
	return (v >= rule.min1 && v <= rule.max1) || (v >= rule.min2 && v <= rule.max2)

}
func MapLocations(content string) int {

	lines := strings.Split(content, "\n")
	state := 0 // "clases"
	classRe := regexp.MustCompile(`(.*): (\d+)-(\d+) or (\d+)-(\d+)`)

	tickets := make ([][]int, 0)
	myTicket := make([]int, 0)

	rules := make([]Rules, 0)
	for _, l := range lines {
		if l == "" {
			continue
		}
		if l == "your ticket:" {
			state = 1 // "your"
			continue
		}

		if l == "nearby tickets:" {
			state = 2 // "nearby"
			continue
		}

		if state == 0 {
			rg := classRe.FindStringSubmatch(l)
			rules = append(rules, Rules{
				rg[1],
				atoiMust(rg[2]),
				atoiMust(rg[3]),
				atoiMust(rg[4]),
				atoiMust(rg[5]),
			})
		}

		if state == 1 || state == 2 {

			ticket := make([]int, 0)
			nums := strings.Split(l, ",")

			acceptable := 0
			for _, n := range nums {
				if n == "" {
					continue
				}
				no := atoiMust(n)
				ticket = append(ticket, no)

				for _, r := range rules {
					if isValueValidForRule(no, r) {
						acceptable++
						break
					}
				}
			}

			if state == 1 {
				myTicket = ticket
			} else {
				if acceptable == len(nums) {
					tickets = append(tickets, ticket)
				}
			}
		}
	}


	// Lets create a list of possible identifications for each column
	colCount := len(tickets[0])

	fieldsAllowed := make(map[int][]string)

	for c := 0; c < colCount; c++ {
		for _, rule := range rules {
			acceptable := true
			for r:= 0; r < len(tickets); r++ {
				v := tickets[r][c]
				if !((v >= rule.min1 && v <= rule.max1) || (v >= rule.min2 && v <= rule.max2)) {
					acceptable = false
					break
				}
			}

			if acceptable {
				fieldsAllowed[c] = append(fieldsAllowed[c], rule.name)
			}
		}
	}

	// now find uniques and remove them from other list

	changed := true
	for changed {
		changed = false
		for ko,v := range fieldsAllowed {
			if len(v) == 1 {
				for k,existing := range fieldsAllowed {
					if k == ko {
						continue
					}

					fieldsAllowed[k] = make([]string, 0)
					for _, ve := range existing {
						if ve == v[0] {
							changed = true
							continue
						}
						fieldsAllowed[k] = append(fieldsAllowed[k], ve)
					}
				}

			}
		}
	}

	ans := 1
	for k, v:= range fieldsAllowed {

		if strings.Contains(v[0], "departure") {
			fmt.Println(v, myTicket[k])
			ans *= myTicket[k]


		}

	}

	return ans
}

func main() {

	conent, err := ioutil.ReadFile("./input16.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := GetScanningRate(string(conent))
	fmt.Println(ans)

	ans = MapLocations(string(conent))
	fmt.Println(ans)



	
}
