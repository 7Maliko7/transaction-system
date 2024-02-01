package oc

import (
	"github.com/go-kit/kit/endpoint"
	kitoc "github.com/go-kit/kit/tracing/opencensus"

	"go.opencensus.io/trace"
)

func ServerEndpoint(operationName string, attrs ...trace.Attribute) endpoint.Middleware {
	attrs = append(
		attrs, trace.StringAttribute("gokit.endpoint.type", "server"),
	)
	return kitoc.TraceEndpoint(
		"gokit/endpoint "+operationName,
		kitoc.WithEndpointAttributes(attrs...),
	)
}
