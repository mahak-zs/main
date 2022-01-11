package main

import (
	"fmt"
	"strings"
)

func bob(s string) string {
	strings.TrimSpace(s)
	if s == "" { // without addressing him
		return "Fine. Be that way!"
	} else if isUpper(s) { //yell at him
		if s[len(s)-1:] == "?" { //yell at him with question
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out?"
	} else if s[len(s)-1:] == "?" { // question ask
		return "Sure"

	} else {
		return "Whatever" // for anything else
	}
}
func isUpper(s string) bool { //to check whether string is capital or not
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 97 && int(s[i]) <= 122 { // if it finds any lowercase then return false
			return false
		}

	}
	return true
}

func main() {
	fmt.Println(bob("HELLO?"))
}
