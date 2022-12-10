package chapter1

import (
	"fmt"
	"strings"
)

func ExampleStringReplace() {
	broken := "G# r#coks!"
	replace := strings.NewReplacer("#", "o")
	fixed := replace.Replace(broken)
	fmt.Println(fixed)
}
