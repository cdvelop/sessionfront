module github.com/cdvelop/sessionfrontend

go 1.20

require (
	github.com/cdvelop/model v0.0.107
	github.com/cdvelop/sessionhandler v0.0.19
)

require (
	github.com/cdvelop/object v0.0.66 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/structs v0.0.1 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/object => ../object

replace github.com/cdvelop/sessionhandler => ../sessionhandler
