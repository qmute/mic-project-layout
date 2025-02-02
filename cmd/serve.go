package cmd

import (
	"easyslip.cc/mic-project-layout/internal/app/serve"
	"github.com/quexer/utee"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "startup serve",
	Long:  `启动api服务`,
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := serve.New()
		utee.Chk(err)
		defer cleanup()

		err = app.Init().Run()
		utee.Chk(err)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
