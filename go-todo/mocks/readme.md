use mockery to generate mocks
`mockery --dir=./services/ --name=TodoServiceInterface --output=mocks/`
`mockery --dir=./services/ --name=DBInterface --output=mocks/`

install with:
`go install github.com/vektra/mockery/v2@latest && export PATH=$PATH:$(go env GOPATH)/bin`
