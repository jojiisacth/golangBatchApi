export GOPATH=$PWD


echo loading dependancies 
go get github.com/gorilla/mux
echo running app  , please browse http://localhost:8081/products/
go run main.go
