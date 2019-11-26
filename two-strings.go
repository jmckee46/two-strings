package main

import "strings"

func twoStrings(s1 string, s2 string) string {
	present := strings.ContainsAny(s1, s2)

	if present {
		return "YES"
	}

	return "NO"
}

func main() {}
