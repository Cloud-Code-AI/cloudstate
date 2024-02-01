## Steps to Build release binaries

```
# For Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o cloudstate-windows.exe ./cmd/main.go

# For macOS 64-bit
GOOS=darwin GOARCH=amd64 go build -o cloudstate-macos ./cmd/main.go

# For Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o cloudstate-linux ./cmd/main.go

```