/*
Package qingstor provided support for qingstor object storage (https://www.qingcloud.com/products/qingstor/)
*/
package qingstor

//go:generate go run github.com/golang/mock/mockgen -package qingstor -destination mock_test.go github.com/qingstor/qingstor-sdk-go/v4/interface Service,Bucket
//go:generate go run -tags tools github.com/aos-dev/go-storage/v3/cmd/definitions service.toml
