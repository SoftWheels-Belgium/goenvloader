# goenvloader
goenvloader is a golang package to load any env file

# Description
This package contains only one fonction *Load*.

To use it you have to pass the path of the env file and any structure pointer to get environment data.

You can use the *mapstructure* to remap value between your struct and env file.

NB: Do not use composed struct

# How to get ?
export GOPRIVATE=github.com/JustinCassart/goenvloader
go get github.com/JustinCassart/goenvloader
