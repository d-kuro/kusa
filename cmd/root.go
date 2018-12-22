package cmd

import (
	"fmt"
	"os"

	"github.com/d-kuro/kusa/log"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use: "kusa",
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
		log.Error("failed execute command", zap.Error(err))
		os.Exit(1)
	}
}
