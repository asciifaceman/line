/*
Copyright © 2023 Charles <Asciifaceman> Corbett
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/asciifaceman/line/lineutil"
	"github.com/spf13/cobra"
)

var (
	lineRange []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "line -l n -l n-N filename",
	DisableFlagsInUseLine: true,
	Short:                 "Read only specific lines or ranges of lines",
	Long: `Read only specific lines or ranges of lines.
Singleton and ranges are accepted by the lines flag. This can be in 
the format of either [ -l 5 ] or a range such as [ -l 5-10 ]

If the given range is not present at all in the file, it will return an EOF
however if a partial range is present, it will print what it found and present
an EOF warning to stderr

	`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}

		stat, err := os.Stat(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "[prerun] error stating file [%s]: %v", args[0], err)
			os.Exit(1)
		}

		if stat.IsDir() {
			fmt.Fprintf(os.Stderr, "[prerun] given file [%s] is a directory\n", args[0])
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(lineRange) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		filename := args[0]

		lineRanges, err := lineutil.ParseLineRanges(lineRange)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, r := range lineRanges {
			lines, err := lineutil.ReadLineRangeFromFile(filename, r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: failed to read line range %v from file [%s]: %v\n", r, args[0], err)
				return
			}

			for _, line := range lines {
				fmt.Println(line)
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.line.yaml)")
	rootCmd.PersistentFlags().StringSliceVarP(&lineRange, "lines", "l", lineRange, "A line number or range of lines in the target file to print defined as N-N (ex. 12-15)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
