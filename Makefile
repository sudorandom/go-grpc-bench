.PHONY: all complex empty

all: empty complex

# Note that we don't run these tests with extra protobuf plugins because empty messages don't benefit from these optimizations at all
empty:
	@./benchmark.sh empty connectrpc
	@./benchmark.sh empty grpc-go
	@./benchmark.sh empty grpc-go-servehttp

complex:
	@./benchmark.sh complex connectrpc
	@./benchmark.sh complex connectrpc-withvt
	@./benchmark.sh complex grpc-go
	@./benchmark.sh complex grpc-go-withvt
	@./benchmark.sh complex grpc-go-servehttp
	@./benchmark.sh complex grpc-go-servehttp-withvt
