# dbuggen2dbuggen2
Converts dbuggen (issues) to (2) dbuggen 2 (issues)

## What does it actually do?

This program converts issues of [Konglig Datasektionens](https://datasektionen.se/) student paper from the old static markdown files found on [dbuggen](https://github.com/datasektionen/dbuggen) and converting them into postgresql injects which can be loaded onto the new database described on [dbuggen2](https://github.com/datasektionen/dbuggen2). The postgresql schema which is used can be found on that repository.

## Running

The tool requires no real setup except installing the correct version of Go, descibed in `go.mod`. Apart from that you just need to run

```shell
go run .
```

and wait for the program to finish. The result will be written to a `.psql` file. The tool downloads the latest version of main of the [dbuggen repository](https://github.com/datasektionen/dbuggen), and works from there, which is why an internet connection is required for running this.