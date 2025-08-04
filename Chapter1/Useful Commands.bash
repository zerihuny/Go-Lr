# Initialize a new Go module (project)
go mod init myproject
# Creates a go.mod file to manage dependencies for the project named 'myproject'

# Add a dependency or update it to latest version
go get github.com/gin-gonic/gin
# Downloads and adds the Gin web framework to your project

# Tidy up the go.mod and go.sum files
go mod tidy
# Removes unused packages and adds missing ones based on imports

# Run a Go file directly
go run main.go
# Compiles and runs the 'main.go' file without creating a binary

# Build the application into a binary
go build
# Compiles all Go files in the current folder into an executable

# Build a specific file
go build myapp.go
# Compiles 'myapp.go' into an executable

# Run tests in the current directory
go test
# Runs all test functions (starting with 'Test') in *_test.go files

# Run tests with verbose output
go test -v
# Shows detailed output about each test being run

# Install the current package
go install
# Compiles and installs the binary to the Go bin directory

# Format all Go code in the current directory
go fmt
# Fixes spacing, indentation, etc., to follow Go style rules

# Show documentation for a package
go doc fmt
# Displays documentation for the 'fmt' package

# List all installed Go modules
go list -m all
# Shows all modules used in your project and their versions

# Check for common issues in code
go vet
# Reports potential errors and suspicious constructs in Go code

# Print your current GOPATH (workspace)
go env GOPATH
# Outputs the current GOPATH setting used by Go
