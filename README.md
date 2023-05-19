# dan2

An experiment in htmx and templ.

## Generate

### templ

Html templating defined in `.templ` files of the `www` directory.

```
templ generate
```

### ent

Database models, added with and managed by the [ent ORM](https://entgo.io/).

Add models to the schema with `go run -mod=mod entgo.io/ent/cmd/ent new ModelTypeName`.

After updating the schema (fields, edges), run `go generate ./ent`.

## Run

```
./dan2 serve
```

## TODO

### GRPC

#### Auth

- https://pkg.go.dev/google.golang.org/grpc@v1.55.0/credentials#AuthInfo
- https://github.com/johanbrandhorst/grpc-auth-example

#### Protobuf

- https://github.com/search?q=repo%3Aargoproj%2Fargo-workflows+proto&type=code
- https://github.com/argoproj/argo-workflows/blob/e35f4912daa0b1ea9ecb6884c07eae862a61c278/pkg/apiclient/workflowtemplate/workflow-template.proto
