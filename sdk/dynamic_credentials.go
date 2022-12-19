package sdk

import "context"

type KeyGenerator func(ctx context.Context, in ProvisionInput, out *ProvisionOutput) (map[FieldName]string, error)

type KeyRemover func(ctx context.Context, in ProvisionInput) error
