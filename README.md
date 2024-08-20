# Tenbounce

## TODOs

- [ ] README
- [x] TODOs
- [x] design and implement db interface
- [ ] API vs db models
- [x] config
- [x] secrets
- [x] deploy to GCP
- [x] auth concerns, now and future
- [ ] CRITICAL: non-success API responses
- [ ] non-API tests
- [ ] API tests
- [ ] UI file structure -- static directory and handlers; see http.FileServer, also see [echo static files](https://echo.labstack.com/docs/static-files)
- [ ] ZAP
- [ ] open source
- [ ] Secrets Manager
- [ ] Github Action
- [ ] New Postgres user
- [ ] New signing secret
- [ ] mount config file in Cloud Run
- [ ] deploy script

## Features

- [ ] gyms/teams/groups/classes
- [ ] permissioning system
- [ ] delete entities
- [ ] responsive UI
- [ ] immediate redirect for login

### Deploy

```sh
docker build -t tenbounce-image .
docker tag tenbounce-image us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce:release3
docker push us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce
<create new revision>
```
