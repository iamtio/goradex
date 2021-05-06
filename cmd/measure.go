package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/iamtio/goradex/radexone"
	"github.com/spf13/cobra"
	"github.com/tarm/serial"
)

// measureCmd represents the measure command
var measureCmd = &cobra.Command{
	Use:   "measure",
	Short: "Get current measures from Radex One",
	Long:  `Return current measueres from USB connected Radex One personal dosimeter`,
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		drr := radexone.NewDataRequest(0)
		encoded := drr.Marshal()

		c := &serial.Config{Name: serialPort, Baud: serialBaud, ReadTimeout: time.Millisecond * 100}
		s, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(">: % X\n", encoded)
		s.Write(encoded)
		buf := make([]byte, 1)
		var result []byte
		for {
			if n, err := s.Read(buf); err != nil || n == 0 {
				break
			}
			result = append(result, buf[0])
		}
		fmt.Printf("<: % X\n", result)
		resp := radexone.DataReadResponse{}
		resp.Unmarshal(result)
		fmt.Printf("CPM: %d, Ambient: %d, Accumulated: %d", resp.CPM, resp.Ambient, resp.Accumulated)

	},
}

func init() {
	rootCmd.AddCommand(measureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// measureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// measureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
