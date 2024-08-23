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
- [x] non-API tests
- [ ] API tests
- [ ] UI file structure -- static directory and handlers; see http.FileServer, also see [echo static files](https://echo.labstack.com/docs/static-files)
- [ ] ZAP
- [x] open source
- [x] Secrets Manager
- [ ] Github Action
- [x] Rotate postgres pw
- [x] New signing secret
- [x] mount config file in Cloud Run
- [x] mount user_secrets.json in Cloud Run
- [x] deploy script
- [x] version endpoint
- [x] Makefile

## Features

- [ ] gyms/teams/groups/classes
- [ ] permissioning system
- [ ] delete entities
- [x] responsive UI
- [x] immediate redirect for login

### Deploy

```sh
docker build -t tenbounce-image .
docker tag tenbounce-image us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce:release9
docker push us-central1-docker.pkg.dev/tenbounce-prod/tenbounce/tenbounce
<create new revision>
```
