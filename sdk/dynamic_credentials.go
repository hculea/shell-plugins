package sdk

import "context"

type KeyGenerator func(ctx context.Context, in ProvisionInput, out *ProvisionOutput) error

type KeyRemover func(ctx context.Context, in ProvisionInput) error
