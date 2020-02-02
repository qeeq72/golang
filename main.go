package main

import (
	"fmt"
	"os"
	"regexp"
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

/*func getNumValues(s string) []uint32 {
	re := regexp.MustCompile(`\d{1,}`)
	sArr := re.FindAllString(s, -1)
	for _, s1 := range sArr {
		fmt.Println(s1)
	}
}*/

func checkInput(s string) int {
	if len(s) == 0 {
		fmt.Println("Empty string!")
		return -1
	}
	if isNum(rune(s[0])) {
		fmt.Println("Wrong string!")
		return -2
	}
	return 0
}

func replaceNumToLiteralValues(s string) string {
	if checkInput(s) != 0 {
		return "Try again!"
	}
	re := regexp.MustCompile(`[a-z]\d{1,}`)
	numLitArr := re.FindAllString(s, -1)
	res := s
	for _, s1 := range numLitArr {
		n, _ := strconv.Atoi(s1[1:])
		res = strings.Replace(res, s1, getMultiLiteral(rune(s1[0]), n), 1)
	}
	re = regexp.MustCompile(`\\\d\\\d\d{0,}`)
	dnumDslashArr := re.FindAllString(res, -1)
	for _, s1 := range dnumDslashArr {
		n1 := string(s1[1])
		n2 := s1[3]
		n, _ := strconv.Atoi(s1[4:])
		res = strings.Replace(res, s1, n1+getMultiLiteral(rune(n2), n), 1)
	}
	re = regexp.MustCompile(`\\\\\d{1,}`)
	numDslash := re.FindAllString(res, -1)
	for _, s1 := range numDslash {
		n, _ := strconv.Atoi(string(s1[2:]))
		res = strings.Replace(res, s1, getMultiLiteral('\\', n), 1)
	}
	re = regexp.MustCompile(`\\\d\d{1,}`)
	dnumSlash := re.FindAllString(res, -1)
	for _, s1 := range dnumSlash {
		n1 := s1[1]
		n2, _ := strconv.Atoi(string(s1[2:]))
		res = strings.Replace(res, s1, getMultiLiteral(rune(n1), n2), 1)
	}
	return res
}

func getMultiLiteral(r rune, n int) string {
	if n == 0 {
		return string(r)
	}
	res := ""
	for n > 0 {
		res += string(r)
		n--
	}
	return res
}

func main() {
	//fmt.Printf("world \u9333")
	//r := []rune("123world \u9333")
	//fmt.Printf("%v", r)
	//pars("a2s3de\\\\5a13")
	var s string
	fmt.Println("Type correct string, please!")
	fmt.Fscan(os.Stdin, &s)
	fmt.Println("Your result:")
	fmt.Print(replaceNumToLiteralValues(s))
}
