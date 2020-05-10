package golang1

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	fmt.Println("golang version !!!!")
   return fmt.Sprintf("MSG from golangv1.0.4:  %s", name)
}

