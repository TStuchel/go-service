module github.com/TStuchel/go-service/auth

go 1.14

replace github.com/TStuchel/go-service/common => ../common

require (
	github.com/TStuchel/go-service/common v0.0.0-20200508133933-a92c6cf4b188
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.10.0
)
