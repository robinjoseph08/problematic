package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func init() {
	fileCmd.AddCommand(readCmd)
	fileCmd.AddCommand(writeCmd)
	fileCmd.AddCommand(readwriteCmd)
	rootCmd.AddCommand(fileCmd)
}

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Make file system operations",
}

var readCmd = &cobra.Command{
	Use:   "read <filepath>",
	Short: "Attempt to read from the file at the given path",
	Args:  fileArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileRun("read", cmd, args)
	},
}

var writeCmd = &cobra.Command{
	Use:   "write <filepath>",
	Short: "Attempt to write to the file at the given path",
	Args:  fileArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileRun("write", cmd, args)
	},
}

var readwriteCmd = &cobra.Command{
	Use:   "readwrite <filepath>",
	Short: "Attempt to write to and then read from the file at the given path",
	Args:  fileArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fileRun("readwrite", cmd, args)
	},
}

func fileArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("filepath is required")
	}
	return nil
}

func fileRun(action string, cmd *cobra.Command, args []string) {
	filepath := args[0]

	if action == "write" || action == "readwrite" {
		fmt.Printf("Attempting to write to %s\n", filepath)
		err := ioutil.WriteFile(filepath, []byte("test"), 0644)
		if err != nil {
			fmt.Printf("Encountered an error when writing: %s\n", err.Error())
		} else {
			fmt.Printf("Successfully wrote 4 bytes\n")
		}
	}

	if action == "read" || action == "readwrite" {
		fmt.Printf("Attempting to read from %s\n", filepath)
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf("Encountered an error when reading: %s\n", err.Error())
		} else {
			fmt.Printf("Successfully read %d byte(s)\n", len(buf))
		}
	}
}
