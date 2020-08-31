module github.com/dnahurnyi/uploader/clientAPI

go 1.14

replace portDomain => ../portDomain

require (
	github.com/rs/zerolog v1.19.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.31.1
	portDomain v0.0.0-00010101000000-000000000000
)
