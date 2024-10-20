# Capital Gain
![capital-gain](https://img.shields.io/badge/capital--gain-gray?logo=go)
![technology Go 1.22](https://img.shields.io/badge/technology-go%201.22-blue.svg)
![Build & Test](https://github.com/FabsHC/capital-gain/actions/workflows/go.yml/badge.svg)
[![Go Coverage](https://github.com/FabsHC/capital-gain/wiki/coverage.svg)](https://raw.githack.com/wiki/FabsHC/capital-gain/coverage.html)

This project simulates the purchase and sale of shares. It does not use any database, all data is stored in memory during the execution of a list of operations.

The application will receive a list of operations that it can execute purchase and sale, returning for each operation how much tax was paid.
The operations will be in the order in which they occurred, that is, the second operation in the list happened after the first and so on.
Each line is an independent simulation, the program will not maintain the state obtained in one line for the others.

Below we have the application input and output data.<br>

Input:

| Name              | Detail                                                           |
|:------------------|:-----------------------------------------------------------------| 
| operation         | Whether the operation is a purchase(buy) or sale(sell) operation |
| unit-cost         | Share unit price in a currency with two decimal places           |
| quantity          | Number of shares traded                                          |

Output:

| Name | Detail                                         |
|:-----|:-----------------------------------------------| 
| tax  | Amount of tax paid for the operation performed |

In the [resources](resources) folder we have several examples to run the application.
In the [CASES.md](docs/CASES.md) we are detailing these examples.

## Technology
The technology used is the Go language in version 1.22.

## Project Organization
```
├── /cmd
    ├── /app....................: Contains main file to run the application
    └── /handlers...............: Contains the entry point for application integration
├── /docs.......................: Contains the documentation from some examples
├── /internal...................: All core implementation its here
    ├── /models.................: Contains Data structure
    ├── /services...............: Contains all business validation
    └── /utils..................: General stuff
```

## How to run
The application expects data in STDIN format and will return json in STDOUT format. So you can provide a JSON or a file containing several lines with each line having a JSON.
Examples below:

### Using Go Run
You can run the commands below to run the application
```shell
go run cmd/app/main.go
```
```shell
go run cmd/app/main.go < resources/case_7 
```

### Using Docker
Just run the docker commands below to create the docker image and run the container.
``` shell
docker run -it --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22 go run cmd/app/main.go
```
``` shell
docker run -i --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22 go run cmd/app/main.go < resources/case_1
```

## How to run the tests
Run the command below in the terminal to run the application tests.<br>
### Using docker:
```shell
make -f Makefile test-docker
```
### Using go:
```shell
make -f Makefile test
```

## Lint and Coverage
The commands bellow need `gotestsum` and `golangci-lint` installed
```shell
make -f Makefile lint
```
```shell
make -f Makefile cover
```