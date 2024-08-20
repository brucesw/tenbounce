# Tenbounce

## TODOs

- [ ] README
- [ ] TODOs
- [x] design and implement db interface
- [ ] API vs db models
- [x] config
- [x] secrets
- [ ] deploy to GCP
- [x] auth concerns, now and future
- [ ] CRITICAL: non-success API responses
- [ ] non-API tests
- [ ] API tests
- [ ] UI file structure -- static directory and handlers; see http.FileServer, also see [echo static files](https://echo.labstack.com/docs/static-files)
- [ ] ZAP
- [ ] open source

## Features

- [ ] gyms/teams/groups/classes
- [ ] permissioning system
- [ ] delete entities

## P-lan

### Prework

- [x] create GCP account w free trial + credit card: `tenbounce.official@gmail.com`
- [x] create project: `tenbounce-prod`
- [x] share access
- [x] update config file (project?, db creds, other?, signing secret)
- [x] stretch: Postgres db set up
  - [x] `tenbounce-db-prod` smallest possible db, ~private IP, auto ip range, default network~, also updated to have public IP creds in tenbounce.yaml
- [x] stretch: Postgres db tables created

### Work

- [x] pick destination
  - [ ] App Engine: https://cloud.google.com/sql/docs/postgres/connect-app-engine-standard#go
  - [x] Cloud Run: https://cloud.google.com/sql/docs/postgres/connect-run
- [x] deploy
- [x] stretch: Postgres db connected
- [ ] stretch: Cloud SQL proxy?
  - [ ]: https://cloud.google.com/sql/docs/postgres/sql-proxy
  - [ ]: https://cloud.google.com/sql/docs/postgres/sql-proxy
- [ ]: super stretch: GitHub Action
- [ ]: super duper stretch: Secrets Manager

### Postwork

- [x] revoke access
- [ ] $$ PROFIT $$

### Deploy

```sh
docker build -t tenbounce-image .
docker tag tenbounce-image us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce:release3
docker push us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce
<create new revision>
```
