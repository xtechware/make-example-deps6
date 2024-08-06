package helloDeps6

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps6.PrintPhrase")
	return phrase
}
