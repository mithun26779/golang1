package golang1

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("MSG from golang1: , %s", name)
}

