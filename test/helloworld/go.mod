module helloworld

go 1.23.0

replace github.com/alibaba/opentelemetry-go-auto-instrumentation => ../../../opentelemetry-go-auto-instrumentation

replace github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg => ../../../opentelemetry-go-auto-instrumentation/pkg

replace github.com/alibaba/opentelemetry-go-auto-instrumentation/test/verifier => ../../../opentelemetry-go-auto-instrumentation/test/verifier

require (
	go.opentelemetry.io/otel v1.35.0
	golang.org/x/text v0.25.0
	golang.org/x/time v0.11.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
)
