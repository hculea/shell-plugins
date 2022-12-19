package terraform

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func terraformCLI() schema.Executable {
	return schema.Executable{
		Name:      "terraform CLI", // TODO: Check if this is correct
		Runs:      []string{"terraform"},
		DocsURL:   sdk.URL("https://terraform.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.AccessKey,
			},
		},
	}
}
