package terraform

import (
	"github.com/1Password/shell-plugins/credentials"
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "terraform",
		Platform: schema.PlatformInfo{
			Name:     "terraform",
			Homepage: sdk.URL("https://terraform.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			credentials.AccessKey(),
		},
		Executables: []schema.Executable{
			terraformCLI(),
		},
	}
}
