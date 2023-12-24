package main
import (
	"fmt"
	"github.com/newcl/bazel/proto/app1"
)
func main()  {
	fmt.Printf("Hello from app1 \n")

	p := app1.Person{}
	fmt.Printf("%v", p)
}