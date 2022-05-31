package cmd

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/cobra"
)

var (
	goroutines int64
)

func init() {
	cpuCmd.Flags().Int64VarP(&goroutines, "goroutines", "g", 100, "Number of goroutines to start")
	rootCmd.AddCommand(cpuCmd)
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Take up a significant amount of CPU",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting up %d goroutines\n", goroutines)

		for i := 0; i < int(goroutines); i++ {
			go stall()
		}

		fmt.Println("Waiting to be terminated...")
		<-time.After(time.Duration(math.MaxInt64))
	},
}

func stall() {
	for {
	}
}
