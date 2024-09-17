package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func checkLuhn(purportedCC string) bool {
	var sum = 0
	var nDigits = len(purportedCC)
	var parity = nDigits % 2
	for i := 0; i < nDigits; i++ {
		var digit = int(purportedCC[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

func leftPad(s string, p string, c int) string {
	var t bytes.Buffer
	if c <= 0 {
		fmt.Println("Invalid Length of Padding")
		return ""
	}
	if len(p) < 1 {
		fmt.Println("Invalid Pad string")
		return ""
	}
	for i := 0; i < (c - len(s)); i++ {
		t.WriteString(p)
	}
	t.WriteString(s)
	return t.String()
}

func main() {
	mxNum := 999999999
	mxLen := len(strconv.Itoa(mxNum))
	var ival string
	for i := 0; i < mxNum; i++ {
		ival = leftPad(strconv.Itoa(i), "0", mxLen)
		if checkLuhn(ival) {
			fmt.Printf("%s\n", ival)
		}
	}
	return
}
