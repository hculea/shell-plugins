package gcloud

import (
	"context"
	"fmt"
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
	"time"
)

func Credentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.Credentials,
		DocsURL:       sdk.URL("https://gcloud.com/docs/credentials"),             // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.gcloud.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Credentials,
				MarkdownDescription: "Credentials used to authenticate to Google Cloud.",
				Secret:              true,
			},
		},
		DefaultProvisioner: provision.EnvVars(map[string]sdk.FieldName{
			"GOOGLE_APPLICATION_CREDENTIALS": fieldname.Credentials,
		}),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(map[string]sdk.FieldName{
				"GOOGLE_APPLICATION_CREDENTIALS": fieldname.Credentials,
			}),
			TryGoogleCloudConfigFile(),
		),
		KeyGenerator: func(ctx context.Context, in sdk.ProvisionInput, out *sdk.ProvisionOutput) (map[sdk.FieldName]string, error) {
			ts, err := impersonate.CredentialsTokenSource(nil, impersonate.CredentialsConfig{
				TargetPrincipal: "test-sa-dynamic-creds@secret-service-dev.iam.gserviceaccount.com",
				Scopes:          []string{"https://www.googleapis.com/auth/cloud-platform"},
			}, option.WithCredentialsJSON([]byte(in.ItemFields[fieldname.Credentials])))
			if err != nil {
				return nil, err
			}

			token, err := ts.Token()
			if err != nil {
				return nil, err
			}

			fmt.Printf("Successfully created the following token:\n%s\n", token.AccessToken)

			err = out.Cache.Put("token", token.AccessToken, time.Now().Add(10*time.Hour))
			if err != nil {
				return nil, err
			}

			newCreds := map[sdk.FieldName]string{
				fieldname.Token: token.AccessToken,
			}

			return newCreds, nil

		},
		KeyRemover: func(ctx context.Context, in sdk.ProvisionInput) error {
			return nil
		},
	}
}

// TODO: Check if the platform stores the Credentials in a local config file, and if so,
// implement the function below to add support for importing it.
func TryGoogleCloudConfigFile() sdk.Importer {
	return importer.TryFile("~/path/to/config/file.yml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		// var config Config
		// if err := contents.ToYAML(&config); err != nil {
		// 	out.AddError(err)
		// 	return
		// }

		// if config. == "" {
		// 	return
		// }

		// out.AddCandidate(sdk.ImportCandidate{
		// 	Fields: map[sdk.FieldName]string{
		// 		fieldname.: config.,
		// 	},
		// })
	})
}

// TODO: Implement the config file schema
// type Config struct {
//	 string
// }
