package middleware

import (
	"github.com/nikhilsbhat/ingress-traefik-converter/pkg/configs"
	"github.com/nikhilsbhat/ingress-traefik-converter/pkg/converters/models"
)

/* ---------------- UNSUPPORTED/REDUNDANT ANNOTATIONS ---------------- */

// ExtraAnnotations handles the below unsupported annotations.
// Annotations:
//   - "nginx.ingress.kubernetes.io/proxy-buffering"
//   - "nginx.ingress.kubernetes.io/service-upstream"
//   - "nginx.ingress.kubernetes.io/enable-opentracing"
//   - "nginx.ingress.kubernetes.io/enable-opentelemetry"
//   - "nginx.ingress.kubernetes.io/backend-protocol"
//   - "nginx.ingress.kubernetes.io/grpc-backend"
func ExtraAnnotations(ctx configs.Context) {
	ctx.Log.Debug("running converter ExtraAnnotations")

	if ctx.Annotations[string(models.ProxyBuffering)] == "off" {
		ctx.Result.Warnings = append(ctx.Result.Warnings,
			"proxy-buffering=off is default behavior in Traefik",
		)
	}

	if ctx.Annotations[string(models.ServiceUpstream)] == "true" {
		ctx.Result.Warnings = append(ctx.Result.Warnings,
			"service-upstream=true is default behavior in Traefik",
		)
	}

	if ctx.Annotations[string(models.EnableOpentracing)] == "true" {
		ctx.Result.Warnings = append(
			ctx.Result.Warnings,
			"enable-opentracing is global in Traefik and cannot be enabled per Ingress",
		)
	}

	if ctx.Annotations[string(models.EnableOpentelemetry)] == "true" {
		ctx.Result.Warnings = append(
			ctx.Result.Warnings,
			"enable-opentelemetry must be configured globally in Traefik static config"+`tracing:
  otlp:
    grpc:
      endpoint: otel-collector:4317`,
		)
	}

	if v := ctx.Annotations[string(models.BackendProtocol)]; v != "" {
		ctx.Result.Warnings = append(ctx.Result.Warnings,
			"backend-protocol must be applied to IngressRoute service scheme, check for generated ingressroutes.yaml",
		)
	}

	if ctx.Annotations[string(models.GrpcBackend)] == "true" {
		ctx.Result.Warnings = append(ctx.Result.Warnings,
			"grpc-backend requires IngressRoute service scheme h2c or https+h2, check for generated ingressroutes.yaml",
		)
	}
}
