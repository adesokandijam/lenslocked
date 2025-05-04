package main

import (
	"net/http"

	"github.com/adesokandijam/lenslocked/controllers"
	"github.com/adesokandijam/lenslocked/models"
	"github.com/adesokandijam/lenslocked/templates"
	"github.com/adesokandijam/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//parse template
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))
	usersC := controllers.Users{
		UserService: &models.UserService{
			DB: db,
		},
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.html", "tailwind.html"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.Get)
	http.ListenAndServe(":3000", r)
}
