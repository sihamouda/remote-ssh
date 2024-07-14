package cmd

import (
	"github.com/sihamouda/remote-ssh/config"
	"github.com/spf13/cobra"
)

func Main() {
	cobra.OnInitialize(config.Main)

	var rootCmd = &cobra.Command{}

	rootCmd.AddCommand(InitPrint)
	rootCmd.Execute()
}
