package packageone

import "fmt"

var PackageVar = "package var in packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
