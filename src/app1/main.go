package main
import (
	"fmt"
	"github.com/newcl/bazel/proto/app1"
	"go.uber.org/fx"
	_ "github.com/rs/zerolog"
    _ "github.com/rs/zerolog/log"
)
func main()  {
	fmt.Printf("Hello from app1 \n")

	p := app1.Person{}
	fmt.Printf("%v\n", p)

	fx.New().Run()
}