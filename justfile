# Run all tests verbose with concurrency of 5
test:
    go test -v -p 5 ./...

# Run tests for a specific package path
test-pkg pkg:
    go test -v -p 5 ./{{pkg}}/...

# Find and run tests for packages matching a name
test-find name:
    go test -v -p 5 $(go list ./... | grep {{name}})

# Run tests with race detection
test-race:
    go test -v -p 5 -race ./...
