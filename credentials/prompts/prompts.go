package prompts

import (
	"fmt"

	"github.com/charmbracelet/huh"

	"github.com/engmtcdrm/minno/app"
	"github.com/engmtcdrm/minno/credentials"
	pp "github.com/engmtcdrm/minno/utils/prettyprint"
)

// GetCredOptions returns a slice of huh.Options for all available credentials
func GetCredOptions() ([]huh.Option[credentials.Credential], error) {
	credFiles, err := credentials.GetCredFiles()
	if err != nil {
		return nil, err
	}

	if len(credFiles) == 0 {
		return nil, fmt.Errorf("no credentials found to update\n\nPlease run command %s to create a credential", pp.Greenf("%s create", app.Name))
	}

	options := []huh.Option[credentials.Credential]{}

	for _, cred := range credFiles {
		options = append(options, huh.NewOption(cred.Name, cred))
	}

	return options, nil
}
