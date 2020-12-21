package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Expr struct  {
	t int // 0 -> num, 1 -> add, 2->sub, 3->another expr
	val int
	e1 *Expr
	e2 *Expr
}

type Parser struct {
	line string
	index int
	advanced bool
}


func mustAtoi(ns string) int {
	num, err := strconv.Atoi(ns)

	if err != nil {
		log.Fatal(err)
	}
	return num
}

func (p *Parser) skipSpace() {
	index := p.index
	if index >= len(p.line) {
		return
	}
	for _, c := range p.line[p.index:] {
		if c != ' ' {
			break
		}
		index++
	}

	p.index = index
}

func (p *Parser) getSubExpr() *Expr {
	// ignore all the spaces first
	p.skipSpace()
	numStr := make([]int32, 0)
	for _, c := range p.line[p.index:]  {

		if c == ')' {
			break
		}

		p.index++
		if c == ' '{
			break
		}
		if c == '(' {
			pe := p.parseExpresion(true)
			return pe
		}
		numStr = append(numStr, c)
	}

	e1 := new(Expr)
	e1.t = 0
	e1.val = mustAtoi(string(numStr))


	return e1
}

func (p *Parser) parseExpresion(consumeClosing bool)  *Expr {

	// "2 * 3 + (4 * 5)"
	e1 := p.getSubExpr()
	p.skipSpace()
	var ans *Expr

	p.skipSpace()

	for {
		if p.index >= len(p.line) || p.line[p.index] == ')' {
			// Just a number right now
			if consumeClosing {
				p.index++
			}
			if ans == nil {
				return e1
			} else {
				return ans

			}
		}

		if ans != nil{
			ans2 := new(Expr)
			ans2.e1 = ans
			ans = ans2
		} else {
			ans = new(Expr)
			ans.e1 = e1

		}

		op := p.line[p.index]
		p.index++

		p.skipSpace()
		var e2  *Expr
		if op == '+' {
			ans.t = 1
			e2 = p.getSubExpr()
			ans.e2 = e2

		} else if op == '*' {
			ans.t = 2
			if p.advanced {
				e2 = p.parseExpresion(false)
			} else {
				e2 = p.getSubExpr()
			}
			ans.e2 = e2
		} else {
			fmt.Errorf("Failed to recognize op: %c\n", op)
		}

		p.skipSpace()

	}

	return ans
}

func Solve(expr *Expr) int {

	ans := 0;
	switch (expr.t) {
	case 0: {ans = expr.val}
	case 1: {ans = Solve(expr.e1) + Solve(expr.e2)}
	case 2: {ans = Solve(expr.e1) * Solve(expr.e2)}
	default: {log.Fatalf("Unknown type")}
	}
	return ans
}

func Eval(content string) int {

	p := Parser{content, 0, false}
	e := p.parseExpresion(true)

	ans := Solve(e)
	return ans
}

func ppp(e *Expr, spaceCount int) {
	if e == nil {
		return
	}
	fmt.Println(e.t, e.val)
	for i:=0; i < spaceCount; i++ {
		fmt.Print(" ")
	}
	ppp(e.e1, spaceCount+1)
	ppp(e.e2, spaceCount+1)
}
func EvalAdvanced(content string) int {

	p := Parser{content, 0, true}
	e := p.parseExpresion(true)

	ans := Solve(e)
	return ans
}

func main() {
	content, err :=ioutil.ReadFile("./input18.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	ans := 0
	for _, l := range lines {
		ans += Eval(l)
	}

	fmt.Println("Part a ", ans)
	ans = 0
	for _, l := range lines {
		ans += EvalAdvanced(l)
	}

	fmt.Println("Part b ", ans)

}
