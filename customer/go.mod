module github.com/TStuchel/go-service/customer

go 1.14

replace github.com/TStuchel/go-service/common => ../common

require (
	github.com/TStuchel/go-service/common v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.9.0
)
