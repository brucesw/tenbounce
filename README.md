# Tenbounce

## Introduction

Tenbounce is web app that helps competitive trampoline athletes and coaches track their progress. After they do _something_ at practice, be it a routine or a certain combination of skills or their ten highest bounces in a row, they can log their time of flight in the app, as timed with a stopwatch. Higher is better. After logging enough data, the coach and the athlete can reflect on the data to gain various insights.

This app in particular is a rebuild of the app I built when I was in college, before I had any professional experience as a software engineer. The main focus of the app was the functionality because that's really all I was capable of. The app itself was really poorly written, and I paid for it in development and maintenance costs. I decided to rebuild portions of the app to leverage my updated skillset and to refresh my engineering skills.

N.B. The backend is 100% me and the frontend is 100% ChatGPT.  Vibe coded.

## TODOs

- [x] README
- [x] TODOs
- [x] design and implement db interface
- [ ] API vs db models
- [x] config
- [x] secrets
- [x] deploy to GCP
- [x] auth concerns, now and future
- [x] immediate redirect for login
- [ ] CRITICAL: non-success API responses
- [x] non-API tests
- [ ] API tests
- [ ] UI file structure -- static directory and handlers; see http.FileServer, also see [echo static files](https://echo.labstack.com/docs/static-files)
- [ ] ZAP logging
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
- [ ] App communicates if no db

## Features

- [ ] gyms/teams/groups/classes
- [ ] permissioning system
- [x] delete entities
- [x] responsive UI

## Deploy & Infrastructure

The app is deployed to GCP Cloud Run via `make deploy`. The deploy script builds a Docker image, pushes it to Artifact Registry and points a new Cloud Run revision to the fresh image.

Secrets, in the format shown in [secrets/tenbounce-example.yaml](secrets/tenbounce-example.yaml), exist in GCP Secrets Manager and are attached to the Cloud Run container via volume mount.

The data lives in a Cloud SQL Postgres instance accessed via private IP.

## Auth

Each user is provided a unique link, `/set_user/<secret_string_here>`, that allows them to log in. It places a cookie on their browser, then redirects them to the app. The cookie contains their user ID and a hash of their (user ID + secret signing key). The user ID, of course, lets the API know who is using it and the hash allows an auth middleware to verify the user.

Future plans involve leveraging Auth0 for login and proper jwts.
