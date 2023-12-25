package	cmd

import(
	"fmt"
	"os"
"context"
"net"
"net/http"


	"github.com/spf13/cobra"
		"go.uber.org/fx"
	_ "github.com/rs/zerolog"
    _ "github.com/rs/zerolog/log"
)

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
	  OnStart: func(ctx context.Context) error {
		ln, err := net.Listen("tcp", srv.Addr)
		if err != nil {
		  return err
		}
		fmt.Println("Starting HTTP server at", srv.Addr)
		go srv.Serve(ln)
		return nil
	  },
	  OnStop: func(ctx context.Context) error {
		return srv.Shutdown(ctx)
	  },
	})
	return srv
  }
  


var serveCmd = &cobra.Command{
	Use: "xxxx",
	Short: "short desc",
	Long: "Long desc",
	Run:func(cmd *cobra.Command, args []string){
fx.New(
fx.Provide(NewHTTPServer),
fx.Invoke(func(*http.Server){}),
).Run()
	},
}

func Execute(){
	if err := serveCmd.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}