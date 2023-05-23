# goenvloader
goenvloader is a golang package to load any *.env* file

# Description
You can either load our *.env* into a map[string]string or into a struct.

There are two functions :

- *LoadToMap(filename string) map[string]string* which loads your env file into a map[string]string 
- *Load(filename string, config interface{})* which loads your env file into a well defined structure

You can use the *mapstructure* to remap value between your struct and the env file.

NB: Do not use composed struct

# How to get ?
go get github.com/JustinCassart/goenvloader
