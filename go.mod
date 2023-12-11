module github.com/cdvelop/sessionfrontend

go 1.20

require github.com/cdvelop/model v0.0.78

require github.com/cdvelop/sessionhandler v0.0.6

require (
	github.com/cdvelop/object v0.0.42 // indirect
	github.com/cdvelop/strings v0.0.8 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/sessionhandler => ../sessionhandler
