package aws

import (
	"github.com/1Password/shell-plugins/credentials"
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "aws",
		Platform: schema.PlatformInfo{
			Name:     "AWS",
			Homepage: sdk.URL("https://aws.amazon.com/"),
		},
		Credentials: []schema.CredentialType{
			credentials.AccessKey(),
		},
		Executables: []schema.Executable{
			AWSCLI(),
		},
	}
}
