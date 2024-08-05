# Tenbounce

## TODOs

- README
- TODOs
- design and implement db interface
- API vs db models
- config
- secrets
- deploy to GCP
- auth concerns, now and future
- non-success API responses
- non-API tests
- API tests
- UI file structure -- static directory and handlers; see http.FileServer

## P-lan
- [X] pros and cons discussion of previous interfaces
- [X] design db interface
  - [X] two implementations: in memory & postgres
  - [X] package layout
  - [X] interface layout (composition)
  - [ ] layers?
    - [ ] powerful base level + queries, wrappers with conveniences
- [X] where does the object attach?
- [X] how to access object in methods?
- [X] implement in-memory db
- [ ] SQLite db
- [ ] stretch: jumpstart postgres
  - [ ] [database/sql](https://pkg.go.dev/database/sql) seems to be the goto
  - [ ] [squirrel](https://github.com/Masterminds/squirrel) also exists
  - [ ] local postgres
  - [ ] API route with example db interaction
