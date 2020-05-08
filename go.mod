module github.com/TStuchel/go-service

go 1.14

replace github.com/TStuchel/go-service/common => ./common

replace github.com/TStuchel/go-service/app => ./app

replace github.com/TStuchel/go-service/customer => ./customer

require (
	github.com/TStuchel/go-service/app v0.0.0-00010101000000-000000000000
	github.com/TStuchel/go-service/customer v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.7.4
)
