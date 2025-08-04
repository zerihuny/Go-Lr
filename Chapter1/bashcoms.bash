# Initialize a new Go module (project)
go mod init myproject
# Creates a go.mod file to manage your project's dependencies

# Add a dependency or update it
go get github.com/gin-gonic/gin
# Downloads and adds the Gin web framework to your project

# Tidy up the go.mod and go.sum files
go mod tidy
# Removes unused modules and adds missing ones

# Update all dependencies to latest compatible versions
go get -u ./...
# Updates all project dependencies recursively

# Run a single Go file
go run main.go
# Compiles and runs the specified Go file

# Build the project into an executable
go build
# Compiles all Go files in the current directory

# Build a specific file
go build myapp.go
# Compiles just one file into a binary

# Build and name the binary
go build -o server
# Creates an executable named 'server'

# Run all tests in the current directory
go test
# Executes all test functions in *_test.go files

# Run tests with verbose output
go test -v
# Shows detailed test execution info

# Run tests in a specific file
go test -v handlers_test.go
# Runs only tests from the given file

# Run a specific test function
go test -run TestLogin
# Runs only the TestLogin test function

# Install the current package
go install
# Builds and installs the binary to GOPATH/bin

# Format all Go code in current folder and subfolders
go fmt ./...
# Formats Go source code using standard style

# Analyze code for potential issues
go vet ./...
# Finds suspicious constructs, possible bugs

# Show documentation for a package
go doc github.com/gin-gonic/gin
# Shows docs for the imported Gin package

# List all used modules
go list -m all
# Shows all modules used and their versions

# Show Go installation environment variables
go env
# Prints current Go environment setup

# Show where compiled binaries are installed
go env GOBIN
# Displays the path to Go binary install location

# Print your current GOPATH (workspace)
go env GOPATH
# Shows where Go workspace is located
