# Technical Test VocaGames

## Introduction

This repository contains the code for the technical test of the Requirements Backend Role in VocaGames.

## Using

- Go 1.23.4

## Project Structure

```
.
├── dist                        // Distribution files (binary)
├── parking                     // Parking lot implementation package
│   ├── command.go              // Command parser
│   ├── parking_lot.go          // Parking lot implementation
│   ├── parking_lot_test.go     // Parking lot tests
│   └── run.go                  // parking package run API
├── command_case1.txt           // Test case for the command
├── go.mod                      // Go module file
├── main_test.go                // Main program tests
├── main.go                     // Main program
├── README.md                   // This file
└── Technical Test Back End.pdf // Technical test document
```

## Run from Source

To run the program from source, use the following command:

```bash
go run main.go <command_file>
```

Replace `<command_file>` with the path to the command file you want to run.

For example:

```bash
go run main.go command_case1.txt
```

This will run the program with the command file `command_case1.txt` and print the output to the console.

## Run from Binary

You can use distribution files (binary) to run the program, or you can build the program from source.

### Build from Source

To build the program from source, use the following command:

```bash
go build -o ./parkinglot .

# or

go build -o ./parkinglot.exe .
```

This will build the program and create a binary file named `parkinglot` in the current directory.

To run the program from binary, use the following command:

```bash
./parkinglot <command_file>
```

Replace `<command_file>` with the path to the command file you want to run.

For example:

```bash
./parkinglot command_case1.txt
```

This will run the program with the command file `command_case1.txt` and print the output to the console.

## Test

To run the tests, use the following command:

```bash
go test -v ./...
```

This will run all the tests in the project and print the output to the console.

## Benchmark

To run the benchmarks, use the following command:

```bash
go test -bench=. -benchtime=1000x -benchmem
```

This will run all the benchmarks in the project and print the output to the console.

### Benchmark Results

The benchmark results are as follows:

```
goos: darwin
goarch: amd64
pkg: github.com/ifanfairuz/technical-test-vocagames
cpu: Intel(R) Core(TM) i3-1000NG4 CPU @ 1.10GHz
BenchmarkParking-4          1000             53345 ns/op          162265 B/op        130 allocs/op
PASS
ok      github.com/ifanfairuz/technical-test-vocagames  0.472s
```
