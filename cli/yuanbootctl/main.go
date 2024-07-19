package main

import (
	"github.com/liangboceo/yuanboot/cli/yuanbootctl/cmds"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yuanbootctl",
	Short: "A generator for Cobra based Applications",
	Long: `yuanbootctl is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a yuanboot application.`,
}

func main() {
	//templates.GetProjectByName("console").List()
	rootCmd.AddCommand(cmds.VersionCmd)
	rootCmd.AddCommand(cmds.RunCmd)
	rootCmd.AddCommand(cmds.BuildCmd)
	rootCmd.AddCommand(cmds.NewCmd)
	_ = rootCmd.Execute()

}
