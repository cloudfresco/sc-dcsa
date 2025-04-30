package common

import (
	"context"
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func HandleCacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set cache-control headers
		headers := w.Header()
		headers.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		headers.Set("Pragma", "no-cache")
		headers.Set("Expires", "0")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
		}

		// Handle preflight OPTIONS requests
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// ChainMiddlewares - chains multiple middlewares together
func ChainMiddlewares(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(finalHandler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
		return finalHandler
	}
}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Permissions []string `json:"permissions"`
	Email       string   `json:"email"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (c CustomClaims) HasPermissions(expectedClaims []string) bool {
	if len(expectedClaims) == 0 {
		return false
	}
	for _, scope := range expectedClaims {
		if !Contains(c.Permissions, scope) {
			return false
		}
	}
	return true
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ValidateToken(audience string, domain string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := getToken(r)
			if err != nil {
				http.Error(w, "Error parsing token", http.StatusUnauthorized)
				return
			}
			middleware, claims, err := getClaims(audience, domain, tokenString, w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			v := ContextStruct{}

			v.Email = claims.Email
			v.TokenString = tokenString
			ctx := context.WithValue(r.Context(), KeyEmailToken, v)

			middleware.CheckJWT(next).ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getToken(r *http.Request) (string, error) {
	authHeaderParts := strings.Fields(r.Header.Get("Authorization"))
	if len(authHeaderParts) > 0 && strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Error parsing token")
	}
	return authHeaderParts[1], nil
}

func NewProxyHandler(backendURL string) http.Handler {
	backend, _ := url.Parse(backendURL)
	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			// Copy the original request's context to the new request
			ctx := r.In.Context()
			r.Out = r.Out.WithContext(r.In.Context())
			// Extract the context data (e.g., JWT claims)
			if claims, ok := ctx.Value(KeyEmailToken).(ContextStruct); ok {
				// Add the context data to the request headers
				r.Out.Header.Set("X-User-Email", claims.Email)
				r.Out.Header.Set("X-Auth-Token", claims.TokenString)
			}

			// Set the backend URL and other headers
			r.SetURL(backend)
			r.SetXForwarded()
			r.Out.Header.Set("X-Forwarded-Host", r.In.Header.Get("Host"))
			r.Out.Host = backend.Host
		},
	}
	return proxy
}
