package middleware

import (
	"github.com/nikhilsbhat/ingress-traefik-converter/pkg/configs"
	"github.com/nikhilsbhat/ingress-traefik-converter/pkg/converters/models"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefik "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/* ---------------- REDIRECT ---------------- */

// SSLRedirect handles the below annotations.
// Annotations:
//   - "nginx.ingress.kubernetes.io/ssl-redirect"
//   - "nginx.ingress.kubernetes.io/force-ssl-redirect"
func SSLRedirect(ctx configs.Context) {
	ctx.Log.Debug("running converter SSLRedirect")

	ssl := ctx.Annotations[string(models.SslRedirect)]
	force := ctx.Annotations[string(models.ForceSslRedirect)]

	if ssl != "true" && force != "true" {
		return
	}

	ctx.Result.Middlewares = append(ctx.Result.Middlewares, &traefik.Middleware{
		TypeMeta: metav1.TypeMeta{
			APIVersion: traefik.SchemeGroupVersion.String(),
			Kind:       "Middleware",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      mwName(ctx, "https-redirect"),
			Namespace: ctx.Namespace,
		},
		Spec: traefik.MiddlewareSpec{
			RedirectScheme: &dynamic.RedirectScheme{
				Scheme:    "https",
				Permanent: true,
			},
		},
	})
}
