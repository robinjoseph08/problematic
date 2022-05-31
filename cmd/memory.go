package cmd

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

const megabyte = 1 << 20

var (
	memory int64
)

func init() {
	memoryCmd.Flags().Int64VarP(&memory, "memory", "m", 10, "Amount of memory to take up")
	rootCmd.AddCommand(memoryCmd)
}

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Take up a significant portion of memory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Taking up %dMB of memory\n", memory)

		// Create a slice and fill it up so it takes up physical memory,
		// not just virtual memory.
		ballast := make([]byte, memory*megabyte)
		for i := 0; i < len(ballast); i++ {
			ballast[i] = byte('A')
		}

		fmt.Println("Waiting to be terminated...")
		<-time.After(time.Duration(math.MaxInt64))
	},
}
