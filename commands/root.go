package commands

import (
	"fmt"

	"github.com/johnthethird/thresher/config"
	"github.com/johnthethird/thresher/configdir"

	"github.com/moby/term"
	"github.com/spf13/cobra"
)

const rootCmdName = "thresher"
var appDirs = configdir.New(rootCmdName)
var appConfig *config.AppConfig = &config.AppConfig{}
const asciiArt = `
 _______ _     _  ______ _______ _______ _     _ _______  ______
    |    |_____| |_____/ |______ |______ |_____| |______ |_____/
    |    |     | |    \_ |______ ______| |     | |______ |    \_
                                                                
`

func NewRootCommand() *cobra.Command {
	
	cmd := &cobra.Command{
		Use:   rootCmdName,
		Short: "short desc",
		Long: `long desc`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// We run this method for its side-effects. On windows, this will enable the windows terminal
			// to understand ANSI escape codes.
			// TODO do we need this? brings in a lot of dependencies
			_, _, _ = term.StdStreams()


			filename, _ := cmd.Flags().GetString("config")
			if config.FileExists(filename) {
				appConfig = config.Load(filename)
			}

			fmt.Print(asciiArt)
		},
	}

	cmd.PersistentFlags().StringP("config", "c", "", "config file which **contains secrets**")
	cmd.PersistentFlags().StringP("log", "l", "thresher.log", "logfile")
	// cmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging")

	cmd.AddCommand(initCommand())
	cmd.AddCommand(versionCommand())
	cmd.AddCommand(walletCommand())
	cmd.AddCommand(testUICommand())
	cmd.AddCommand(debugCommand())
	cmd.AddCommand(bootstrapCommand())

	return cmd
}

func Execute() {
	cobra.CheckErr(NewRootCommand().Execute())
}
