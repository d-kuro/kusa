package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "kusa",
	Short: "kusa is GitHub contributions create tool",
	Long:  "By using kusa you can create GitHub contributions on any day.",
	Run: func(cmd *cobra.Command, args []string) {
		kusa := `wʕ ◔ϖ◔ʔw
_  _ _  _ ____ ____
|_/  |  | [__  |__|
| \_ |__| ___] |  |`
		fmt.Println(kusa)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
