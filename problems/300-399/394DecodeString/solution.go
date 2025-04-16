package decodestring

import "strings"

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	currNum := 0
	currStr := ""

	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			currNum = currNum*10 + int(ch-'0')
		case ch == '[':
			numStack = append(numStack, currNum)
			strStack = append(strStack, currStr)
			currNum = 0
			currStr = ""
		case ch == ']':
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prevStr + strings.Repeat(currStr, n)
		default:
			currStr += string(ch)
		}
	}

	return currStr
}
