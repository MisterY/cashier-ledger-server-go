# ledger-server-go
A simple Rest API server for ledger-cli. Used by Cashier PWA.

# Purpose

This is an server application that provides a Rest API endpoint. All it does is forward the
received [ledger-cli](https://ledger-cli.org/) command to the ledger executable and send back the result. It is a simple
proxy for HTTP(S) requests and an instance of ledger-cli.

It effectively replaces [CashierSync](https://gitlab.com/alensiljak/cashiersync) and
[CashierSync Go](https://gitlab.com/alensiljak/cashiersync-go). By simplifying the server-side
componenent, all the logic is able to reside in [Cashier](https://github.com/MisterY/cashier).
Any future feature additions will not require changes in the server-side component.

For additional information, see related projects.

# Development

The app is developed in Go(lang) v1.19.

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

todo: .rest file

# Installation

todo: go install

# Run

`ledger-server-go`

# Query

The server supports CORS and can be queried from a PWA JavaScript application.

Simply issue a GET request to the //server?command=...

where command is the ledger-cli command

# License

GNU General Public License v3.0

See [license](LICENSE).
