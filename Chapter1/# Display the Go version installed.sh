# Display the Go version installed
go version
# Shows the currently installed Go version (e.g., go1.21.0)

# Print the Go installation path and settings
go env
# Outputs all Go environment variables like GOPATH, GOROOT, GOOS, GOARCH, etc.

# Create a new workspace/project folder structure (manual)
mkdir myproject && cd myproject
# Start a new project by creating and navigating into a folder

# Run multiple Go files together
go run main.go routes.go db.go
# Runs all listed files as a single program

# Install a Go tool globally
go install golang.org/x/tools/cmd/stringer@latest
# Installs and adds a Go CLI tool to your system

# Download modules but don’t install/build yet
go mod download
# Downloads all required modules into the module cache

# Check if go.mod is tidy without modifying it
go mod tidy -v
# Shows what was added or removed

# Check whether all imports are correctly resolved
go list ./...
# Lists all packages in your project

# View the import path of a package
go list -f '{{.ImportPath}}'
# Useful for modules and nested packages

# Check for updates to modules
go list -u -m all
# Lists current module versions and if updates are available

# Display the call graph for a Go program (advanced)
go tool callgraph -format digraph main.go
# Shows the call graph for functions (requires setup)
// Display the Go version installed
go version
# Shows the currently installed Go version (e.g., go1.21.0)
# Display the Go installation path and settings
go env
# Outputs all Go environment variables like GOPATH, GOROOT, GOOS, GOARCH, etc.
# Create a new workspace/project folder structure (manual)
mkdir myproject && cd myproject
# Start a new project by creating and navigating into a folder


# Clean compiled binaries and cached files
go clean
# Removes object files, binaries, and cached data

# See what executable will be built
go build -x
# Shows detailed steps of the build process

# Benchmark performance of code (if benchmark tests are written)
go test -bench=.
# Runs performance benchmarks
