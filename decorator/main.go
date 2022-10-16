package decorator

import (
	"fmt"
	"strings"
)

type Decorated = string
type Decorator = func(string) string

func Decorate(c Decorated, ds ...Decorator) Decorated {
	decorated := c
	for _, decorator := range ds {
		decorated = decorator(decorated)
	}
	return decorated
}

func main() {

	var toLower Decorator = strings.ToLower
	var toUpper Decorator = strings.ToUpper
	input := "Decorate"
	inputUppercase := Decorate(input, []Decorator{toUpper}...)
	fmt.Println(inputUppercase)

	allDecorators := []Decorator{toUpper, toLower}
	output := Decorate(input, allDecorators...)
	fmt.Println(output)
}
