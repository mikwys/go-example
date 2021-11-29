# go-example
This is example go code used in tech talk:
> Should I use go as my next microservice?

This is not prod-ready code, acting as a showcase for go great possibilities, especially:

- fast compilation time
- live reload using `modd`
- html templating
- simple http with html templates

using almost none dependencies and external libraries (only m3o client which uses basic go stdlib mostly).


 Required env variables:
- M3O_API_TOKEN   - you can obtain free account at https://m3o.com

Usage (build): ``make all`` and run binary for specific OS.
Usage (build temp & run): go run cmd/tws/main.go 


Thanks to:
- https://github.com/cortesi/modd
- https://github.com/m3o/m3o-go
