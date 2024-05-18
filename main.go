package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{
		"alice", "alimony", "albatross", "ali",
	}

	fmt.Println(longestCommonPrefix(s))

}

func lettersSame(s []byte) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	pre := ""
	sl := []byte{}
	shortestLen := 0

	s := slices.Clone(strs)

	for i := 0; i < len(s)-1; i++ {
		if len(s[i]) < len(s[i+1]) {
			shortestLen = len(s[i])
			s[i+1] = s[i]
		}
	}

	lastEl := s[len(s)-1]

	if len(lastEl) < len(s[len(s)-2]) {
		shortestLen = len(lastEl)
	}
	for j := 0; j < shortestLen; j++ {
		for i := 0; i < len(strs); i++ {
			sl = append(sl, strs[i][j])
			if lettersSame(sl) {
				pre += string(sl[0])
			} else {
				break
			}
		}
	}
	return pre
}

//shortestLen := 0
//s := slices.Clone(strs)
//
//for i := 0; i < len(s)-1; i++ {
//	if len(s[i]) < len(s[i+1]) {
//		shortestLen = len(s[i])
//		s[i+1] = s[i]
//	}
//}
//
//lastEl := s[len(s)-1]
//
//if len(lastEl) < len(s[len(s)-2]) {
//	shortestLen = len(lastEl)
//}
//pre := ""
//
//strs[0][0] == strs[1][0] && strs[1][0] == strs[2][0]
//
//for i := 0; i < strs; i++ {
//	for j := 0; j < shortestLen; j++ {
//		if strs[i][j] == strs []
//	}
//}

//firstL := string(strs[0][0])
//for i := 0; i < len(strs) - 1; i++ {
//	if strings.HasPrefix(strs[i+1],firstL) {
//		pre += firstL
//	}
//
//}

//fmt.Printf("copied slice: %v", s)
//fmt.Printf("original slice: %v", strs)
//return strconv.Itoa(shortestLen)

//firstEl := strs[0]
//
//for _, str := range strs {
//	i := 0
//	for ; i < len(str) && i < len(firstEl) && firstEl[i] == str[i]; i++ {
//	}
//	firstEl = firstEl[:i]
//}
//return firstEl
