# study
study directory has many initial steps and explorations, poc which gives you the knowledge and experience about the various microservices communication protocols, patterns and message encodings
It is a playground to try out new things in a sandboxed way before implementing it in a major project

## Protobuff
- generate go code from protobuff definition

```
# Install protoc compiler
# Install golang proto generation
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# generate code
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

- generate code with various import options
```
protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
protoc --proto_path=./study/protobuff/proto --go_out=./study/protobuff/protobuff-golang/device --go_opt=paths=source_relative wifi.proto
```
the compiler will read input files foo.proto and bar/baz.proto from within the src directory, and write output files foo.pb.go and bar/baz.pb.go to the out directory. The compiler automatically creates nested output sub-directories if necessary, but will not create the output directory itself


## Resouorces
- [protobuff go getting started](https://protobuf.dev/getting-started/gotutorial/)


