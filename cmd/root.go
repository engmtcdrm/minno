package cmd

import (
	"context"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/engmtcdrm/minno/app"
	"github.com/engmtcdrm/minno/env"
)

var (
	rootCmd = &cobra.Command{
		Use:     app.Name,
		Short:   app.ShortDesc,
		Long:    app.LongDesc,
		Example: app.Name,
		Version: getSemVer(app.Version),
	}

	credName string
	output   string
)

// Execute executes the root command.
func Execute() error {
	rootCmd.SilenceUsage = true
	return rootCmd.ExecuteContext(context.Background())
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	cobra.OnInitialize(configInit)
}

func configInit() {
	envVars, err := env.GetEnv()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(envVars.AppHomeDir); os.IsNotExist(err) {
		// Directory does not exist, create it
		err = os.MkdirAll(envVars.AppHomeDir, 0700)
		if err != nil {
			panic(err)
		}

		// Change permission again to get rid of any sticky bits
		err = os.Chmod(envVars.AppHomeDir, 0700)
		if err != nil {
			panic(err)
		}
	} else {
		// Directory exists, make sure directories and files are secure
		err = filepath.Walk(envVars.AppHomeDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return os.Chmod(path, 0700)
			}
			return os.Chmod(path, 0600)
		})
		if err != nil {
			panic(err)
		}
	}
}

// getSemVer returns the semantic version of the input string if it
// matches the pattern `vX.Y.Z`. Otherwise, it returns the input string.
func getSemVer(input string) string {
	// Define the regular expression for semantic versioning
	re := regexp.MustCompile(`^v?(\d+\.\d+\.\d+)$`)

	match := re.FindStringSubmatch(input)

	// If there's a match return the semantic version
	if len(match) > 1 {
		return match[1]
	}

	// If no match, return the original input
	return input
}
