package main

import "fmt"

func main() {
	fmt.Printf("%s\n", superReducedString("aaabccddd"))
	fmt.Printf("%s\n", superReducedString("aa"))
	fmt.Printf("%s\n", superReducedString("baab"))
}

func superReducedString(s string) string {
	p := 1

	for {
		if len(s) <= p {
			if s == "" {
				return "Empty String"
			}
			return s
		}

		if s[p-1:p] == s[p:p+1] {
			s = s[0:p-1] + s[p+1:]
			if p > 1 {
				p--
			}
		} else {	
			p++
		}
	}
}