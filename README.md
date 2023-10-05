# go-commandline-tool

## How to use various commands

### cvimg

Convert images to ones with a specified extention.

```
go install cmd/cvimg/cvimg.go
cvimg -e png -f images/sample.jpeg -d images/sample.png
```

### gcat

Display the contents of files.

```
go install cmd/gcat/gcat.go
gcat -n README.md go.mod
```
