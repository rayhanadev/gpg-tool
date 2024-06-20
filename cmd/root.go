package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "gpgtool",
    Short: "A simple GPG CLI tool",
    Long:  `A simple GPG CLI tool to encrypt and decrypt messages using GPG keys.`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
