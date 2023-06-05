module github.com/envoyproxy/go-control-plane/contrib

go 1.20

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4
	github.com/envoyproxy/go-control-plane/envoy v0.0.0-00010101000000-000000000000
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/planetscale/vtprotobuf v0.6.0
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)
