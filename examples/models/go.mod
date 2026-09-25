module github.com/inference-gateway/examples/models

go 1.26.7

replace github.com/inference-gateway/sdk => ../..

require github.com/inference-gateway/sdk v0.0.0-00010101000000-000000000000

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/oapi-codegen/runtime v1.7.0 // indirect
)
