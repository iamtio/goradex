package cmd

import (
	"github.com/spf13/cobra"
)

var serialPort string
var serialBaud int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goradex",
	Short: "Radex One CLI",
	Long:  `Radex One CLI`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serialPort, "serial", "s", "", "virtual serial port of radex one. example: COM9")
	rootCmd.MarkPersistentFlagRequired("serial")
	rootCmd.PersistentFlags().IntVarP(&serialBaud, "baud", "b", 9600, "virtual serial port baud rate")
}
