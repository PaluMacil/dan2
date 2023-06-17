package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
// TODO: check if fixed (when grpc is fixed in ent): go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
// patched in  go run -mod=mod github.com/rdeusser/contrib/entproto/cmd/entproto -path ./schema
//go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
