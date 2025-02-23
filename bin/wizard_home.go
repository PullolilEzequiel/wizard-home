package main

import (
	"github.com/spf13/cobra"

	addCommand "github.com/PullolilEzequiel/wizard-home/internal/command/add_command"
	initCommand "github.com/PullolilEzequiel/wizard-home/internal/command/init_command"
	reverseCommand "github.com/PullolilEzequiel/wizard-home/internal/command/reverse_command"
	saveCommand "github.com/PullolilEzequiel/wizard-home/internal/command/save_command"
	setupCommand "github.com/PullolilEzequiel/wizard-home/internal/command/setup_command"
)

var initC = &cobra.Command{
	Use: "init", Short: "Create initial config for wizard_home",
	Run: initCommand.Execute,
}
var saveC = &cobra.Command{
	Use: "save", Short: "Persist our actual config to our remote repository",
	Run: saveCommand.Execute}
var reverseC = &cobra.Command{
	Use: "reverse", Short: "Reverse the actual system config to the last commit in our remote repository",
	Run: reverseCommand.Execute}
var setupC = &cobra.Command{
	Use: "setup", Short: "Change the system configuration to another Wizard Home repository",
	Args: cobra.ExactArgs(1),
	Run:  setupCommand.Execute,
}

var addF = &cobra.Command{
	Use: "add", Short: "Add a file or directory to the config.json file",
	Args: cobra.ExactArgs(1),
	Run:  addCommand.Execute,
}

func main() {
	rootCmd := &cobra.Command{Use: "wizard_home",
		Short: "Wizard home is a CLI tool that seeks to help preserve the personalized configuration of our system.",
	}

	rootCmd.AddCommand(initC)
	rootCmd.AddCommand(saveC)
	rootCmd.AddCommand(reverseC)
	rootCmd.AddCommand(setupC)
	rootCmd.AddCommand(addF)
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}
