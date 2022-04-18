package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	renderChi "github.com/go-chi/render"
	"github.com/joleques/go-chi-poc/src/application"
	"github.com/joleques/go-chi-poc/src/infra/base"
	"github.com/joleques/go-chi-poc/src/infra/jwt"
	renderPkg "github.com/unrolled/render"
	"net/http"
	"time"
)

var render *renderPkg.Render

var dataBase = base.NewMemoryDataBase()

func main() {
	render = renderPkg.New()
	dataBase.Status()
	auth := jwt.JWT{}.New()
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(renderChi.SetContentType(renderChi.ContentTypeJSON))
	r.Use(auth.Verifier())
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Status Ok"))
	})

	r.Route("/product", func(r chi.Router) {
		r.Use(auth.Authenticator())
		r.Post("/", SaveProduct)
		r.Route("/{productId}", func(r chi.Router) {
			r.Get("/", GetProduto)
		})
	})

	http.ListenAndServe(":3000", r)
}

func SaveProduct(writer http.ResponseWriter, request *http.Request) {
	data := &application.Product{}
	if err := renderChi.Bind(request, data); err != nil {
		render.JSON(writer, 400, "invalid request")
		return
	}
	result := application.SaveProductUseCase(*data, dataBase)
	render.JSON(writer, result.StatusCod, result)
}

func GetProduto(writer http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "productId")

	product := application.GetProductUseCase(productId, dataBase)

	render.JSON(writer, 200, product)
}
