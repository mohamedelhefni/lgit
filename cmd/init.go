package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const GIT_DIR = ".lgit"

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init git repo",
	Run: func(cmd *cobra.Command, args []string) {
		currentPath, _ := os.Getwd()
		fmt.Printf("Initing empty repo in %s/%s \n", currentPath, GIT_DIR)
		err := os.Mkdir(GIT_DIR, os.ModePerm)
		if err != nil {
			panic(err)
		}
		err = os.Mkdir(GIT_DIR+"/objects", os.ModePerm)
		if err != nil {
			panic(err)
		}

	},
}
