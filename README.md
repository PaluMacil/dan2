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