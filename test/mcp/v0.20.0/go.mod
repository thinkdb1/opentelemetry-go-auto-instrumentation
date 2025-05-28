module mcp

go 1.23.3

replace github.com/alibaba/opentelemetry-go-auto-instrumentation => ../../../../opentelemetry-go-auto-instrumentation

replace github.com/alibaba/opentelemetry-go-auto-instrumentation/test/verifier => ../../../../opentelemetry-go-auto-instrumentation/test/verifier

require (
	github.com/alibaba/opentelemetry-go-auto-instrumentation/test/verifier v0.0.0-20250409012242-ef76c1556ebc
	github.com/mark3labs/mcp-go v0.20.0
	go.opentelemetry.io/otel/sdk v1.35.0
	go.opentelemetry.io/otel/trace v1.35.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
