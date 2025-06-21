package reverse_go

import (
	"fmt"
  "golang.org/x/example/hello/reverse"
)

func String(input string) string {
	return fmt.Sprintf("==> %s <==", reverse.String(input))
}
