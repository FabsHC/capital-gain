# Capital Gain
![capital-gain](https://img.shields.io/badge/capital--gain-gray?logo=go)
![technology Go 1.20.3](https://img.shields.io/badge/technology-go%201.20.3-blue.svg)

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
The technology used is the Go language in version 1.20.3.

## Project Organization
```
.
├── cmd..................: Contains the main file to run the application.
├── docs.................: Contains the documentation from some examples.
├── internal.............: all core implementation its here.
│   ├── model............: All application structures (DTO).
│   ├── usecase..........: Application core validations.
│   └── util.............: General stuff.

```

## How to run
The application expects data in STDIN format and will return json in STDOUT format. So you can provide a JSON or a file containing several lines with each line having a JSON.
Examples below:

### Using Go Run
You can run the commands below to run the application
```shell
go run cmd/main.go
```
```shell
go run cmd/main.go < resources/case_7 
```

### Using Docker
Just run the docker commands below to create the docker image and run the container.
``` shell
docker run -it --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.20.3 go run cmd/main.go
```
``` shell
docker run -i --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.20.3 go run cmd/main.go < resources/case_1
```

## How to run the tests
Run the command below in the terminal to run the application tests.<br>
### Using docker:
```shell
make -f Makefile test-docker
```
### Using go:
```shell
make -f Makefile test-go
```
