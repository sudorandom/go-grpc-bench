## go-grpc-bench
This is a repo I created to host some benchmarks I was running with grpc-go, grpc-go (using ServeHTTP) and go-connect. I made this repo to make it easier for me to add profiling and tweak more server settings when testing.

## Terminology

- **Profiles**: a combination of the gRPC server and server settings. For this project, it's a small shell script that starts the server.
- **Scenarios**: the settings used to run the benchmark. For this project, it's a JSON file with the relevant [ghz](https://ghz.sh/) configuration.
- **Results**: For this project, they live in `results/{YYYY-MM-DD}/{SCENARIO}/{PROFILE}.json`.

## Running the project
Ensure that you have [ghz](https://ghz.sh/) installed, and just run make to run all the scenarios:

```
make
```

To run a specific scenario, just use the name like these commands (there are only two scenarios right now):
```
make empty
make complex
```
