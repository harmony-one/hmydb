package cmd

import (
	"fmt"
	"os"
	"path"

	color "github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	verbose        bool
	noLatest       bool
	noPrettyOutput bool
	node           string
	keyStoreDir    string
	givenFilePath  string
	// RootCmd is single entry point of the HmyDB
	RootCmd = &cobra.Command{
		Use:          "hmydb",
		Short:        "Harmony blockchain",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		Long: fmt.Sprintf(`
CLI interface to the Harmony blockchain database

%s`, "Invoke 'hmydb cookbook' for examples of the most common, important usages"),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
)

func init() {
	vS := "dump out debug information, same as env var HMY_ALL_DEBUG=true"
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, vS)
	RootCmd.PersistentFlags().StringVarP(&node, "node", "n", "", "<host>")
	RootCmd.PersistentFlags().BoolVar(
		&noLatest, "no-latest", false, "Do not add 'latest' to RPC params",
	)
	RootCmd.PersistentFlags().BoolVar(
		&noPrettyOutput, "no-pretty", false, "Disable pretty print JSON outputs",
	)
	RootCmd.AddCommand(&cobra.Command{
		Use:   "cookbook",
		Short: "Example usages of the most important, frequently used commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("cookbook")
			return nil
		},
	})
	RootCmd.PersistentFlags().StringVar(&givenFilePath, "file", "", "Path to file for given command when applicable")
	RootCmd.AddCommand(&cobra.Command{
		Use:   "docs",
		Short: fmt.Sprintf("Generate docs to a local %s directory", ""),
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, _ := os.Getwd()
			docDir := path.Join(cwd, "doc")
			os.Mkdir(docDir, 0700)
			doc.GenMarkdownTree(RootCmd, docDir)
			return nil
		},
	})
}

var (
	// VersionWrapDump meant to be set from main.go
	VersionWrapDump = ""
	cookbook        = color.GreenString("hmy cookbook")
)

// Execute kicks off the hmydb CLI
func Execute() {
	RootCmd.SilenceErrors = true
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(errors.Wrapf(err, "commit: %s, error", VersionWrapDump).Error())
		fmt.Println("check " + cookbook + " for valid examples or try adding a `--help` flag")
		os.Exit(1)
	}
}
