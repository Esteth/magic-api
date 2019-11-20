package proto

//go:generate docker run --user 1000:1000 -v $PWD:/defs namely/protoc-all -f ./api.proto -l go -o proto
