
@REM This command runs the "skaffold-prebuild" target in the Makefile.
make skaffold-prebuild

@REM This command runs the "kind" profile in Skaffold.
skaffold run -p kind

@REM This command builds the fission.exe binary using the main.go file.
go build -o fission.exe cmd/fission-cli/main.go
