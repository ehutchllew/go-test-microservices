module ehutchllew-broker

go 1.21.3

toolchain go1.22.5

require (
	ehutchllew/go-utils v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
)

replace ehutchllew/go-utils => ../utils
