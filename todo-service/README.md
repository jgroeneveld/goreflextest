# Todo Service

Uses [ent](https://entgo.io/docs/getting-started/) for persistance.

create new model
`go run -mod=mod entgo.io/ent/cmd/ent new MODELNAME`

adjust fields in ent/schema and run generator
`go generate ./ent`
