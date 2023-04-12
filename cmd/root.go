package cmd

import (
	"bytes"
	"embed"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"github.com/theapemachine/wrkspc/brazil"
	"github.com/wrk-grp/errnie"
)

/*
Embed a mini filesystem into the binary to hold the config file,
and some front end templates. This will be compiled into the
binary, so it is easier to manage.
*/
//go:embed cfg/*
var embedded embed.FS

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "am",
		Short: "Agent system",
		Long:  roottxt,
	}
)

var roottxt = `
This is an AI Agent System.
`

/*
Execute configures the CLI and executes the program with the
values that were passed in from the command line.
*/
func Execute() error {
	errnie.Trace()

	// Add the `run` command to the CLI.
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(tradeCmd)

	// Run the program and return any error that may happen.
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Defines the config file that will be loaded, usually just the name of the service.
	// This should be written to the user's home directory as a hidden file.
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		".am.yml",
		"config file (default is $HOME/.am.yml)",
	)

	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

/*
initConfig unpacks the embedded file system and writes the config
file to the home directory of the user if it is not present.
It also automatically generates the CLI documentation for wrkspc.
*/
func initConfig() {
	// Set verbosity level for errnie.
	errnie.Tracing(false)
	errnie.Debugging(true)

	errnie.Trace()

	// Get the config file from the user home path.
	chunks := strings.Split(cfgFile, "/")
	fh, err := embedded.Open("cfg/" + chunks[len(chunks)-1])

	errnie.Handles(err)
	defer fh.Close()

	buf, err := io.ReadAll(fh)
	errnie.Handles(err)

	home := brazil.NewPath("~").Location
	brazil.NewFile(home, cfgFile, bytes.NewBuffer(buf))

	viper.AddConfigPath(home)
	viper.SetConfigType("yml")
	viper.SetConfigName(cfgFile)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	// The method errnie is wrapping here writes the markdown documentation for
	// the command line interface, which is automatically generated.
	brazil.NewPath(brazil.NewPath(".").Location, "docs")
	errnie.Handles(
		doc.GenMarkdownTree(rootCmd, "./docs/"),
	)
}
