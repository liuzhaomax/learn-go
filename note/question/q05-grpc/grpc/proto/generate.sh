# generate codes
# -I .  input from current directory
protoc -I . --go_out=plugins=grpc:. book.proto
