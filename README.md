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
- UI file structure

## P-lan
- [ ] pros and cons discussion of previous interfaces
- [ ] design db interface
  - [ ] two implementations: in memory & postgres
  - [ ] package layout
  - [ ] interface layout (composition)
  - [ ] layers?
    - [ ] powerful base level + queries, wrappers with conveniences
- [ ] where does the object attach?
- [ ] how to access object in methods?
- [ ] implement in-memory db
- [ ] stretch: jumpstart postgres
  - [ ] [database/sql](https://pkg.go.dev/database/sql) seems to be the goto
  - [ ] [squirrel](https://github.com/Masterminds/squirrel) also exists
  - [ ] local postgres
  - [ ] API route with example db interaction
