package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/sihamouda/remote-ssh/config"
	"github.com/spf13/cobra"
)

func initFunction(cmd *cobra.Command, args []string) {
	currentPemFolder, err := config.GetPemFolder()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	fmt.Printf("Current .pem directory: %s\n", currentPemFolder)
	fmt.Printf("Changing .pem directory to: %s\n", args[0])

	err = config.SetConfigPemFolder(args[0])
	if err != nil {
		fmt.Printf("Failed! %s\n", err)
	} else {
		fmt.Println("Succeed!")
	}
}

var InitPrint = &cobra.Command{
	Use:   "init [path of directory where .pem files are stored]",
	Short: "Set up path inital configuration",
	Args:  cobra.ExactArgs(1),
	Run:   initFunction,
}
