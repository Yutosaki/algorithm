package main

import (
	"fmt"
	"strconv"
)

func isPalindromicInBase(base int, num int64) bool {
	if num == 0 {
		return true
	}

	str := make([]int64, 0)
	for num > 0 {
		str = append(str, num%int64(base))
		num /= int64(base)
	}

	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var a int
	var n int64
	fmt.Scan(&a, &n)

	var sum int64

	for i := int64(1); ; i++ {
		s := strconv.FormatInt(i, 10)

		rev := ""
		for j := len(s) - 1; j >= 0; j-- {
			rev += string(s[j])
		}

		evenPalindromic := string(s) + rev
		oddPalindromic := string(s) + rev[1:]

		evenPalindromicInt, _ := strconv.Atoi(evenPalindromic)
		oddPalindromicInt, _ := strconv.Atoi(oddPalindromic)

		if int64(evenPalindromicInt) > int64(n) && int64(oddPalindromicInt) > int64(n) {
			break
		}

		if int64(evenPalindromicInt) <= int64(n) {
			if isPalindromicInBase(a, int64(evenPalindromicInt)) {
				sum += int64(evenPalindromicInt)
			}
		}

		if int64(oddPalindromicInt) <= int64(n) {
			if isPalindromicInBase(a, int64(oddPalindromicInt)) {
				sum += int64(oddPalindromicInt)
			}
		}
	}
	fmt.Println(sum)
}
