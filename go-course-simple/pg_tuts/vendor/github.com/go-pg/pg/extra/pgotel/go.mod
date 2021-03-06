module github.com/go-pg/pg/extra/pgotel/v11

go 1.15

replace github.com/go-pg/pg/v11 => ../..

require (
	github.com/go-pg/pg/v11 v11.0.0-alpha.0
	go.opentelemetry.io/otel v0.18.0
	go.opentelemetry.io/otel/trace v0.18.0
)
