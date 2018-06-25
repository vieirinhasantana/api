package app

import (
	"log"
	"net/http"

	"developer.patronus.io/mobile_backend/app/bundle"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouter()
}

func (a *App) setRouter() {
	a.Post("/auth", a.PostAuthenticationUser)

	a.Get("/categories", a.GetCategories)
	a.Post("/categories", a.PostCategories)
	a.Put("/categories/{id}", a.PutCategories)
	a.Delete("/categories/{id}", a.DeleteCategories)

	a.Get("/products", a.GetProducts)
	a.Post("/products", a.PostProducts)
	a.Put("/products/{id}", a.PutProducts)
	a.Delete("/products/{id}", a.DeleteProducts)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) PostAuthenticationUser(w http.ResponseWriter, r *http.Request) {
	bundle.PostLoginUser(w, r)
}

/************** CATEGORIA *******************/
func (a *App) GetCategories(w http.ResponseWriter, r *http.Request) {
	bundle.GetAllCategories(w, r)
}

func (a *App) PostCategories(w http.ResponseWriter, r *http.Request) {
	bundle.PostCategory(w, r)
}

func (a *App) PutCategories(w http.ResponseWriter, r *http.Request) {
	bundle.PutCategory(w, r)
}

func (a *App) DeleteCategories(w http.ResponseWriter, r *http.Request) {
	bundle.DeleteCategory(w, r)
}

/************** PRODUTOS *******************/
func (a *App) GetProducts(w http.ResponseWriter, r *http.Request) {
	bundle.GetAllProducts(w, r)
}

func (a *App) PostProducts(w http.ResponseWriter, r *http.Request) {
	bundle.PostProduct(w, r)
}

func (a *App) PutProducts(w http.ResponseWriter, r *http.Request) {
	bundle.PutProduct(w, r)
}

func (a *App) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	bundle.DeleteProduct(w, r)
}

func (a *App) Run(host string) {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATH", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Accept"},
	})
	handler := cors.Handler(a.Router)
	log.Fatal(http.ListenAndServe(host, handler))
}
