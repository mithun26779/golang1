package golang1

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	fmt.Println("hello world !!!!")
   return fmt.Sprintf("MSG from golangv1.0.3:  %s", name)
}

func init() {
    fmt.Println("go3 init 2")
}

func init() {
    fmt.Println("go3 init 1")
}
