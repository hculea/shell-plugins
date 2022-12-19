package gcloud

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func GoogleCloudCLI() schema.Executable {
	return schema.Executable{
		Name:      "Google Cloud CLI", // TODO: Check if this is correct
		Runs:      []string{"gcloud"},
		DocsURL:   sdk.URL("https://gcloud.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.Credentials,
			},
		},
	}
}
