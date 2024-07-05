module github.com/bluenviron/gomavlib/v3

go 1.20

require (
	bou.ke/monkey v1.0.2
	github.com/alecthomas/kong v0.9.0
	github.com/pion/transport/v2 v2.2.5
	github.com/stretchr/testify v1.9.0
	go.bug.st/serial v1.6.2
)

require (
	github.com/creack/goselect v0.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.bug.st/serial => github.com/zyxkad/go-serial v1.6.3-0.20240705033605-bc847fb5fdc0
