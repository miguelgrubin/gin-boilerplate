# Dev Tools

## GoSec

Security scanner. Reads your code and checks it agains 40 security rules.

Install:

`go install github.com/securego/gosec/v2/cmd/gosec@latest`

Run:
`gosec ./...`

## Air

Hot realoding. Useful when you are developing.

Install:

`go install github.com/cosmtrek/air@latest`

Just run: `air`

## Delve

Debugger.

Install:
`go install github.com/go-delve/delve/cmd/dlv@latest`

Debug test:
`dlv test ./pkg/server_test.go`

Debug main:
`dlv debug ./main.go`

Set breakpoint:
`break pkg.SetupRouter`
or
`b pkg.SetupRouter`
or
`break 48`

Remove breakpoint:
`clear pkg.SetupRouter`

Print vars:
`locals`
`p var`
`args`

Show current lines:
`ls`

You can also use VSCode with launch.json given.
