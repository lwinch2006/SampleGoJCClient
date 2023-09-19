package customfmt

import "fmt"

func Printfln(template string, args ...any) {
	fmt.Printf(template+"\n", args...)
}
