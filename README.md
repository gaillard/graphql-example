For graphql schema validation and example execution:

* Modify `schema.graphql` and `gqlgen.yml`
* Install dependencies:
```bash
dep ensure
```
* Generate files:
```bash
cd ./schema/
go run ./scripts/gqlgen.go
cd -
```
* Run server:
```bash
go run server.go
```
* Open `http://localhost:8080/query` using Altair Chrome extension. Refresh Docs to validate.
* Before next code generation remove previosly generated `resolver_gen.go`.