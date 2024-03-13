package config

import (
	"fmt"
	"os"
)

func PrintError(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func PrintInfo(v ...interface{}) {
	fmt.Fprintln(os.Stdin, v...)
}

func ReverseRelationshipMap(rel string) string {
	relationshipMap := map[string]string{
		"father":   "son",
		"mother":   "son",
		"son":      "father",
		"daughter": "father",
	}

	return relationshipMap[rel]
}
