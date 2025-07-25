/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"likenft-indexer/cmd/worker/cmd/task"
	verbosityutil "likenft-indexer/cmd/worker/cmd/util/verbosity"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "likenft-indexer-worker",
	Short: "Worker to query EVM data",
	Long: `Worker to query EVM data and either print it on
	screen or put it into DB base on config. A long running
	event loop worker or one off command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.likenft-indexer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("rpc", "r", "http://localhost:8545", "EMV RPC URL")

	// verbosity flag. Only change the log level of application logger. Won't affect loggers in library e.g. asynq
	rootCmd.PersistentFlags().StringP("verbose", "v", string(verbosityutil.VerbosityDefault), "Verbosity of logger")

	rootCmd.AddCommand(task.TaskCmd)
}
