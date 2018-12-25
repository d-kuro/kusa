package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	cobra.OnInitialize()
}

var (
	logger, _ = zap.NewDevelopment()
	rootCmd   = &cobra.Command{
		Use: "kusa",
		Run: func(cmd *cobra.Command, args []string) {
			kusa := `_  _ _  _ ____ ____
|_/  |  | [__  |__|
| \_ |__| ___] |  |`
			fmt.Println(kusa)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("failed execute command", zap.Error(err))
		os.Exit(1)
	}
}
