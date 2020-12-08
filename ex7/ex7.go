package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)


type Node struct {
	 name string
	 count int
}
type Graph = map[string][]Node

func GetCount(countStr string) int {

	if countStr == "no" {
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Fatal(err)
	}
	return count

}
func GetGraphs(content string) (Graph, Graph) {
	lines := strings.Split(content, "\n")

	reg := regexp.MustCompile(`(.*) bags contain (\d+|no) ([^,.]*) bag`)
	regSecondary := regexp.MustCompile(`(\d+) (.*) bag`)

	graph := make(Graph)
	invGraph := make(Graph)
	parent := ""
	for _, l := range lines {
		if l == "" {
			continue
		}
		lVec := strings.Split(l, ",")
		matchString := reg.FindStringSubmatch(l)
		parent = matchString[1]
		child := matchString[3]

		if child == "other" {
			continue
		}

		count := GetCount(matchString[2])
		graph[child] = append(graph[child], Node{parent, count})
		invGraph[parent] = append(invGraph[parent], Node{child, count})

		for _, ll := range lVec[1:] {
			matchString := regSecondary.FindStringSubmatch(ll)

			count := GetCount(matchString[1])
			child := matchString[2]
			graph[child] = append(graph[child], Node{parent, count})
			invGraph[parent] = append(invGraph[parent], Node{child, count})
		}
	}

	return graph, invGraph

}

func FindTotalCombinations(graph Graph) int {
	startNode := "shiny gold"
	seen := make(map[string]bool)
	pendingQ := make([]Node, 0)
	pendingQ = append(pendingQ, Node{startNode, 0})
	seen[startNode] = true
	count := 0
	for len(pendingQ) > 0 {

		c := pendingQ[len(pendingQ) - 1]


		pendingQ = pendingQ[:len(pendingQ) - 1]
		children, exists := graph[c.name]
		if exists {
			for _, c:= range children {
				if _, ok := seen[c.name]; ok {
					continue
				}
				count++
				seen[c.name] = true
				pendingQ = append(pendingQ, c)
			}
		}
	}
	return count
}

func findTotalCasesNeededImpl(table *map[string]int, graph Graph, node Node) int {

	if table == nil {
		log.Fatalf("Received nil table\n")
	}
	count, found := (*table)[node.name]
	if found {
		return count
	}

	children, found := graph[node.name]

	needed := 0
	if !found {
		return needed
	}

	for _, c := range children {
		needed += c.count
		needed += c.count * findTotalCasesNeededImpl(table, graph, c)
	}
	if *table == nil {
		log.Fatalf("Received nil table\n")
	}
	(*table)[node.name] = needed

	return needed
}

func FindTotalCasesNeeded(graph Graph) int {
	startNode := Node{"shiny gold", 1}
	table := make(map[string]int)
	ans := findTotalCasesNeededImpl(&table, graph, startNode)
	return ans
}

func main() {
	content, err := ioutil.ReadFile("./input7.txt")
	if err != nil {
		log.Fatal(err)
	}

	graph, invGraph := GetGraphs(string(content))
	ans := FindTotalCombinations(graph)
	fmt.Println("Solution for day7 a ", ans)
	ans = FindTotalCasesNeeded(invGraph)
	fmt.Println("Solution for day7 b ", ans)


}
