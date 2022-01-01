package main

import (
	"fmt"
	"strings"
)

func main() {
	// *** Strings are immutable!!!!

	// name := "Hello World!"
	// biteRepresentation(name)
	// unicodeRepresentation(name)

	// concatination()

	// stringsPackage()

	stringsManipulation()

}

func biteRepresentation(name string) {
	// %x => bite representation
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
}

func unicodeRepresentation(name string) {
	// Rune is 32bit integer
	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		// y is unicode formatted
		fmt.Println(x, "\t", y, "\t", string(y))
	}
}

func concatination() {
	fmt.Println("Three ways to concatenate strings")
	h := "hello, "
	w := "world."

	// *** easy but not efficient
	myEasyString := h + w
	fmt.Println(myEasyString)

	// *** more efficient
	myString := fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// *** string builders are more efficient!!
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())
}

func stringsPackage() {
	courses := []string{
		"Learn Go for beginners crash course",
		"Learn Java for beginners crash course",
		"Learn Python for beginners crash course",
		"Learn C for beginners crash course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "at index of", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language, Go for it."
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasSuffix(newString, "?"))
	fmt.Println("number of \"Go\" acuurance is", strings.Count(newString, "Go"))
	// if index not found it returns -1
	fmt.Println("\"Go\" starts at position", strings.Index(newString, "Go"))
	fmt.Println("\"Python\" starts at position", strings.Index(newString, "Python"))
	fmt.Println("\"Go\" latest accurance starts at position", strings.LastIndex(newString, "Go"))
}

func stringsManipulation() {
	newString := "Go is a great programming language, Go for it."
	if strings.Contains(newString, "Go") {
		newString = strings.Replace(newString, "Go", "GoLang", 1) // -1 means every where
	}

	fmt.Println(newString)

	badEmail := " me@here.com "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=\n", badEmail)

	str := "alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)
}

func replaceNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}

		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)
	}

	return s
}
