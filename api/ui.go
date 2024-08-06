package api

import _ "embed"

//go:embed ui/homepage.html
var homepageHTML string

//go:embed ui/unauthorized.html
var unauthorizedHTML string
