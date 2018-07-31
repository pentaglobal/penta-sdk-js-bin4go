// Package sdk4js contains the console JavaScript dependencies Go embedded.
package sdk4js

//go:generate go-bindata -nometadata -pkg sdk4js_bin -o ./sdk4js_bin/bindata.go penta.js
//go:generate gofmt -w -s ./sdk4js_bin/bindata.go
