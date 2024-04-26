.PHONY: all

all:
	./benchmark.sh connectrpc complex
	./benchmark.sh connectrpc complex -withvt
	./benchmark.sh grpc-go complex
	./benchmark.sh grpc-go complex -withvt
	./benchmark.sh grpc-go-serve-http complex
	./benchmark.sh grpc-go-serve-http complex -withvt
