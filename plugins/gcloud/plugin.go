package gcloud

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "gcloud",
		Platform: schema.PlatformInfo{
			Name:     "Google Cloud",
			Homepage: sdk.URL("https://gcloud.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			Credentials(),
		},
		Executables: []schema.Executable{
			GoogleCloudCLI(),
		},
	}
}
