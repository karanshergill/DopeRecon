package cmd

import (
	"fmt"
	"os"
	"log"
	
	"github.com/spf13/cobra"
	
	"doperecon/data"
)



// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "doperecon",
	Short: "Dope Recon is an open-source comprehensive recon and ASM framework for bug bounty hunters",
	Long: `A wrapper around industry standard efficient recon tools used my most bug bounty hunters. doperecon is a tool for automating
	reconnaissance and asset management, helping bug bounty hunters discover and analyze targets.

	Caution!
	This tool uses both active and passive recon techniques.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use the correct command. Use --help for more information.")
	 },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dope-recon.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.AddCommand(data.DataCmd())
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


