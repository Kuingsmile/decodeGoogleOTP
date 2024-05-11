package cmd

import (
	"fmt"
	"os"

	"github.com/kuingsmile/decodeGoogleOTP/core"
	"github.com/kuingsmile/decodeGoogleOTP/utils"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "decodeGoogleOTP",
	Short: "A command line tool to decode Google OTP QR codes",
	Long:  `decodeGoogleOTP is a command line tool to decode Google OTP QR codes. Output can be json, csv, qrcode or plain text.`,
	Run:   runFunc,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print version information")
	RootCmd.Flags().StringVarP(&inputFlag, "input", "i", "", "Input file path")
	RootCmd.Flags().StringVarP(&csvFlag, "csv", "c", "", "Output in CSV format and specify the output file")
	RootCmd.Flags().StringVarP(&jsonFlag, "json", "j", "", "Output in JSON format and specify the output file")
	RootCmd.Flags().StringVarP(&qrDirFlag, "qrcode", "q", "", "Output in QR code image format and specify the output directory")
	RootCmd.Flags().StringVarP(&textFlag, "text", "t", "", "Output url list in plain text format and specify the output file")
	RootCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "Output in URL format and specify the output file")
	RootCmd.Flags().BoolVarP(&printQRFlag, "print-qr", "p", false, "Print QR code to terminal")
	RootCmd.Flags().BoolVarP(&debugFlag, "debug", "d", false, "Enable debug mode")
	RootCmd.Flags().BoolVarP(&silentFlag, "silent", "s", false, "Enable silent mode")
}

func runFunc(cmd *cobra.Command, args []string) {
	if versionFlag {
		printVersion(cmd, args)
		return
	}

	if inputFlag == "" {
		cmd.Help()
		return
	}

	var flags = utils.Flags{
		Input:   inputFlag,
		Csv:     csvFlag,
		Json:    jsonFlag,
		Text:    textFlag,
		QrDir:   qrDirFlag,
		Url:     urlFlag,
		Version: versionFlag,
		PrintQR: printQRFlag,
		Debug:   debugFlag,
		Silent:  silentFlag,
	}

	err := core.Flow(flags)

	if err != nil {
		cmd.Help()
	}
}
