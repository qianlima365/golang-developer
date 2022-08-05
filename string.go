package go_example

import (
	"fmt"
	"strings"
)

func main()  {
	var a []string
	b := "a,b,c"
	a = strings.Split(b, ",")
	fmt.Println(len(a))
}
