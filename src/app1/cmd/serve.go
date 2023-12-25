package	cmd

import(
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "xxxx",
	Short: "short desc",
	Long: "Long desc",
	Run:func(cmd *cobra.Command, args []string){
panic("DON'T RUN ME")
	},
}

func Execute(){
	if err := serveCmd.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}