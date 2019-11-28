package proto

//go:generate docker run -v $PWD:/defs namely/protoc-all -f ./api.proto -l go -o proto
