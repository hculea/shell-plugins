package credentials

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestAccessKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, AccessKey().Importer, map[string]plugintest.ImportCase{
		"AWS CLI default config file location": {
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
				"~/.aws/config":      plugintest.LoadFixture(t, "config"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "eu-central-1",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-east-1",
					},
				},
			},
		},
		"AWS CLI custom config file in home dir": {
			Environment: map[string]string{
				"AWS_CONFIG_FILE": "~/.config-custom",
			},
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
				"~/.config-custom":   plugintest.LoadFixture(t, "custom-config"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
			},
		},
		"AWS CLI custom config file in root dir": {
			Environment: map[string]string{
				"AWS_CONFIG_FILE": "/.config-custom",
			},
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
				"/.config-custom":    plugintest.LoadFixture(t, "custom-config"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
			},
		},
		"AWS CLI custom config file in root dir 2": {
			Environment: map[string]string{
				"AWS_CONFIG_FILE": ".config-custom",
			},
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
				".config-custom":     plugintest.LoadFixture(t, "custom-config"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
						fieldname.DefaultRegion:   "us-west-1",
					},
				},
			},
		},
		"AWS CLI empty config file": {
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
				"~/.aws/config":      plugintest.LoadFixture(t, "empty-config"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					},
				},
			},
		},
		"AWS CLI NO config file": {
			Files: map[string]string{
				"~/.aws/credentials": plugintest.LoadFixture(t, "credentials"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIADEFFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "DEFlrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					},
				},
				{
					NameHint: "user1",
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					},
				},
			},
		},
		"default env vars": {
			Environment: map[string]string{
				"AWS_ACCESS_KEY_ID":     "AKIAHPIZFMD5EEXAMPLE",
				"AWS_SECRET_ACCESS_KEY": "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
				"AWS_DEFAULT_REGION":    "us-central-1",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAHPIZFMD5EEXAMPLE",
						fieldname.SecretAccessKey: "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
						fieldname.DefaultRegion:   "us-central-1",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.DefaultRegion: "us-central-1",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.DefaultRegion: "us-central-1",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.DefaultRegion: "us-central-1",
					},
				},
			},
		},
		"env vars with AMAZON_ prefix": {
			Environment: map[string]string{
				"AMAZON_ACCESS_KEY_ID":     "AKIAHPIZFMD5EEXAMPLE",
				"AMAZON_SECRET_ACCESS_KEY": "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAHPIZFMD5EEXAMPLE",
						fieldname.SecretAccessKey: "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
					},
				},
			},
		},
		"AWS_SECRET_KEY": {
			Environment: map[string]string{
				"AWS_ACCESS_KEY": "AKIAHPIZFMD5EEXAMPLE",
				"AWS_SECRET_KEY": "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAHPIZFMD5EEXAMPLE",
						fieldname.SecretAccessKey: "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID: "AKIAHPIZFMD5EEXAMPLE",
					},
				},
			},
		},
		"AWS_ACCESS_SECRET": {
			Environment: map[string]string{
				"AWS_ACCESS_KEY":    "AKIAHPIZFMD5EEXAMPLE",
				"AWS_ACCESS_SECRET": "RnnHD6qgcZ0OpYB3chaK73TcobH1YY7yEEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID: "AKIAHPIZFMD5EEXAMPLE",
					},
				},
				{
					Fields: map[sdk.FieldName]string{
						fieldname.AccessKeyID:     "AKIAHPIZFMD5EEXAMPLE",
						fieldname.SecretAccessKey: "RnnHD6qgcZ0OpYB3chaK73TcobH1YY7yEEXAMPLE",
					},
				},
			},
		},
	})
}

func TestAccessKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, AccessKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.AccessKeyID:     "AKIAHPIZFMD5EEXAMPLE",
				fieldname.SecretAccessKey: "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
				fieldname.DefaultRegion:   "us-central-1",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"AWS_ACCESS_KEY_ID":     "AKIAHPIZFMD5EEXAMPLE",
					"AWS_SECRET_ACCESS_KEY": "lBfKB7P5ScmpxDeRoFLZvhJbqNGPoV0vIEXAMPLE",
					"AWS_DEFAULT_REGION":    "us-central-1",
				},
			},
		},
	})
}
