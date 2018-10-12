// Package sdk4js contains the console JavaScript dependencies Go embedded.
package sdk4js

//go:generate go-bindata -nometadata -pkg sdk4js -o bindata.go penta.js
//go:generate gofmt -w -s bindata.go
