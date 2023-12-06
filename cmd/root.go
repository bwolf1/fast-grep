/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	matcher "github.com/bwolf1/fast-grep/pkg/matcher"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "fast-grep search_pattern file_path",
	Version: "1.0.0",
	Short:   "fast-grep",
	Long: `fast-grep
	
A faster grep tool!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Not enough arguments")
			cmd.Help()
			return
		}
		if len(args) > 2 {
			fmt.Println("Too many arguments")
			cmd.Help()
			return
		}

		filePath := args[1]

		readFile, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var fileLines []string

		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}

		readFile.Close()

		for _, line := range fileLines {
			matcher := matcher.NewMatcherStruct()
			match := matcher.Match(args[0], line)
			if match != "" {
				fmt.Println(match)
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fast-grep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fast-grep" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fast-grep")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
