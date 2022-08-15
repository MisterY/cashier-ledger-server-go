# ledger-server-go
A simple Rest API server for ledger-cli. Used by Cashier PWA.

# Purpose

This is an server application that provides a Rest API endpoint. All it does is forward the
received [ledger-cli](https://ledger-cli.org/) command to the ledger executable and sends back the result. 
It is a simple proxy for HTTP(S) requests and an instance of ledger-cli.

It effectively replaces [CashierSync](https://gitlab.com/alensiljak/cashiersync) and
[CashierSync Go](https://gitlab.com/alensiljak/cashiersync-go). By simplifying the server-side
componenent, all the logic is able to reside in [Cashier](https://github.com/MisterY/cashier).
Any future feature additions will not require changes in the server-side component.

Ledger-cli must be in the path, executable from the current location.
It is advised to have a `.ledgerrc` file in the current directory, to set up the Ledger data sources.

For additional information, see related projects.

# Development

The app is developed in Go(lang).

To debug, simply Run & Debug from vscode. 
Or create a new configuration to attach to process "fresh".

## Hot Reload

To use hot reload during development, simply install Fresh package

```
go get github.com/pilu/fresh
go install github.com/pilu/fresh
```

and then run `fresh` from the project root directory. It will monitor source code files for changes and automatically rebuild the application.

Option B: Use Gin package? Doesn't work.
```
gin
```

## Manual Tests

The API endpoints are listed in the ledger-server.rest file.

# Installation

`go install github.com/MisterY/ledger-server-go@latest`

## Remove

The app gets installed into $GOPATH/bin.
To remove, simple delete the executable file.

# Run

`ledger-server-go [-p <port>]`

The port parameter is optional. By default, the API will serve at port 8080.

# Query

The server supports CORS and can be queried from a PWA JavaScript application.

Simply issue a GET request to the //server?command=...
For example, when running on localhost, the URL would be `http://localhost:8080?command=bal assets and cash`.

where command is the ledger-cli command, just the parameters. 
i.e. `balance assets`
Parameters are also supported, so something like `bal -b 2022-08-10` is possible.

The result, ledger-cli output, is split into lines and return as a JSON array, which can be previewed in a browser or parsed by client apps/tools.

# License

GNU General Public License v3.0

See [license](LICENSE).
