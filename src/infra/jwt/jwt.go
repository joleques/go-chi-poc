package jwt

import (
	"github.com/go-chi/jwtauth"
	"github.com/joleques/go-chi-poc/src/application"
	"github.com/pkg/errors"
	renderPkg "github.com/unrolled/render"
	"net/http"
)

var render *renderPkg.Render

type JWT struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

type Claims struct {
	AgentId float64 `json:"agentId"`
}

func (JWT) New() *JWT {
	jwt := &JWT{
		tokenClaim: "agentId",
		tokenAuth:  jwtauth.New("HS256", []byte("my_secret_key"), nil),
	}
	return jwt
}

func (jwt *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt.tokenAuth)
}

func (jwt *JWT) Decode(r *http.Request) float64 {
	val, _ := jwt.Authenticate(r)
	return val
}

func (jwt *JWT) Authenticate(r *http.Request) (float64, error) {
	token, claims, err := jwtauth.FromContext(r.Context())
	if err != nil || token == nil {
		return 0, errors.Wrap(err, "Empty or invalid JWT")
	}
	return claims[jwt.tokenClaim].(float64), nil
}

func (jwt *JWT) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := jwt.Authenticate(r)
			if err != nil {
				render = renderPkg.New()
				result := application.Result{StatusCod: 400, Message: err.Error()}
				render.JSON(w, 400, result)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
