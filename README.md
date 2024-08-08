# Tenbounce

## TODOs

- [ ] README
- [ ] TODOs
- [X] design and implement db interface
- [ ] API vs db models
- [X] config
- [X] secrets
- [ ] deploy to GCP
- [X] auth concerns, now and future
- [ ] CRITICAL: non-success API responses
- [ ] non-API tests
- [ ] API tests
- [ ] UI file structure -- static directory and handlers; see http.FileServer, also see [echo static files](https://echo.labstack.com/docs/static-files)
- [ ] ZAP

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
- [X] stretch: jumpstart postgres
  - [X] [database/sql](https://pkg.go.dev/database/sql) seems to be the goto
  - [ ] [squirrel](https://github.com/Masterminds/squirrel) also exists
  - [X] local postgres
  - [X] API route with example db interaction
