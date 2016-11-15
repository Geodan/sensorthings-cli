package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunLogin(cmd, args)
		if err != nil {
			exitWithError(err)
		}

	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}

// RunLogin set the login
func RunLogin(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}
	gitOauth := args[0]
	configPath := os.Getenv("HOME")
	configName := ".sti"
	configType := "yaml"
	var configYaml []byte
	var madeConfigFile = false
	configFile := path.Join(configPath,
		(configName + "." + configType))
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)

	if _, err := os.Stat(configFile); err != nil {
		var file, err = os.Create(configFile)
		defer file.Close()
		if err != nil {
			exitWithError(fmt.Errorf("Could not create config: %s", configFile))
		}
		configYaml = []byte("st_server: " + gitOauth + "\n")

		defer file.Close()
		madeConfigFile = true
	} else {
		input, err := ioutil.ReadFile(configFile)
		if err != nil {
			exitWithError(err)
		}

		lines := strings.Split(string(input), "\n")

		isUpdate := false
		gitLine := "st_server: " + gitOauth
		for i, line := range lines {
			if strings.Contains(line, "st_server: ") {
				isUpdate = true
				lines[i] = gitLine
			}
		}
		output := strings.Join(lines, "\n")
		if !isUpdate {
			output = output + "\n" + gitLine + "\n"
		}
		configYaml = []byte(output)
		if madeConfigFile {
			err_del := os.Remove(configFile)
			if err_del != nil {
				exitWithError(fmt.Errorf("Could not delete config: %s", configFile))
			}
		}

	}
	err := ioutil.WriteFile(configFile, configYaml, 0644)
	if err != nil {
		exitWithError(fmt.Errorf("Could not write config to %s", configFile))
	}

	if madeConfigFile {
		fmt.Printf("Login Succeeded")
	} else {
		fmt.Printf("Login Succeeded and Updated")
	}
	return nil
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
