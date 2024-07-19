package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"yuanbootctl/generate/projects"
	"yuanbootctl/utils/consolecolors"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of yuanboot",
	Long:  `All software has versions. This is yuanboot's`,
	Run: func(cmd *cobra.Command, args []string) {
		logo := projects.Logo
		fmt.Println(consolecolors.Blue(string(logo)))
		fmt.Println(" ")
		fmt.Printf("%s   (version:  %s)", consolecolors.Green(":: yuanboot ::"), consolecolors.Blue(projects.Version))

		fmt.Print(consolecolors.Blue(`
light and fast , dependency injection based micro-service framework written in Go.
`))

		fmt.Println(" ")
	},
}
