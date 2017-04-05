export GOPATH=$PWD


echo loading dependancies 
go get github.com/labstack/echo
echo building app  
go run server.go
