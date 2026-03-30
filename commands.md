# go run hello_world.go
# go build hello_world.go  then ./hello

# 

# Common Go Commands

- `go run file.go` – Compile and run a Go program in one step
- `go build file.go` – Compile the program and create an executable
- `./program_name` – Run the compiled executable
- `go install` – Compile and install the binary into `$GOBIN`
- `go env` – Show Go environment variables
- `go mod init main` – Like git init in go this stores version and all (dependencies)
- `go build`- -- Compile to binary
- `go fmt ./...`- -- Format all code
- `go vet ./...` --Static analysis
- `go test ./...` -- Run tests
- `go mod tidy`	-- Clean dependencies
- `go get pkg` --	Download package
- `go install` --	Install binary
- `go doc fmt.Println` --	View docs