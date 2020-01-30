package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isNum(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

func pars(s string) int {
	r := []rune(s)

	if len(r) == 0 {
		fmt.Printf("Where?\n")
		return -1
	}

	if isNum(r[0]) {
		fmt.Printf("Try harder\n")
		return -2
	}
	var sym rune
	var res strings.Builder
	n2 := 0
	for i, r1 := range r {
		//sym = r[i]
		//sym = r1
		if i > n2 || i == 0 {
			asci := rune('\\')
			if r1 == asci && sym != asci {
				sym = r1
				continue
			}

			if isNum(r1) && r1 != asci {
				n2 = i
				strNum := string(r1)
				for j := i + 1; j < len(r); j++ {
					if isNum(r[j]) {
						strNum += string(r[j])
						n2 = j
					} else {
						break
					}
				}
				n, _ := strconv.Atoi(strNum)
				for n > 1 {
					res.WriteRune(sym)
					n--
				}
			} else {
				res.WriteRune(r1)
				sym = r1
			}
		}

	}
	fmt.Print(res.String())
	return 0
}

func main() {
	//fmt.Printf("world \u9333")
	//r := []rune("123world \u9333")
	//fmt.Printf("%v", r)
	pars("a2s3de\\\\5a13")
}
