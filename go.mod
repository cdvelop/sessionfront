module github.com/cdvelop/sessionfrontend

go 1.20

require github.com/cdvelop/model v0.0.75

require github.com/cdvelop/sessionhandler v0.0.2

require (
	github.com/cdvelop/input v0.0.59 // indirect
	github.com/cdvelop/object v0.0.39 // indirect
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/timetools v0.0.24 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
