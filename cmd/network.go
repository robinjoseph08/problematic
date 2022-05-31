package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(networkCmd)
}

var networkCmd = &cobra.Command{
	Use:   "network <url>",
	Short: "Make an HTTP request to the given URL",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("url is required")
		}
		if _, err := url.Parse(args[0]); err != nil {
			return fmt.Errorf("url is malformed: %w", err)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Making a request to %s\n", args[0])

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := client.Get(args[0])
		if err != nil {
			return fmt.Errorf("http error: %w", err)
		}

		fmt.Printf("Got status code %d\n", resp.StatusCode)
		return nil
	},
}
