package api

import (
	"fmt"
	"github.com/plab0n/search-paste/config"
	"github.com/plab0n/search-paste/internal/bus"
	"github.com/plab0n/search-paste/internal/handlers"
	"github.com/plab0n/search-paste/internal/middlewares"
	"github.com/plab0n/search-paste/internal/storage"
	"github.com/plab0n/search-paste/pkg/httputils"
	"github.com/plab0n/search-paste/pkg/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/unrolled/render"
	"github.com/unrolled/secure"

	"github.com/urfave/negroni"
)

type AppServer struct {
	Env     string
	Port    string
	Version string
	handlers.Handlers
}

func (app *AppServer) Run(appConfig config.ApiEnvConfig) {
	app.Env = appConfig.Env
	app.Port = appConfig.Port
	app.Version = appConfig.Version
	app.Sender = &httputils.Sender{
		Render: render.New(render.Options{
			IndentJSON: true,
		}),
	}

	// can change DB implementation from here
	store, err := storage.NewPostgresDB()
	if err != nil {
		logger.Log.Error(err)
		panic(err.Error())
	}
	// Migrations which will update the DB or create it if it doesn't exist.
	//if err := storage.MigratePostgres("file://migrations"); err != nil {
	//	logger.Log.Fatal(err)
	//}
	app.Storage = store
	app.Bus = bus.New()
	router := mux.NewRouter().StrictSlash(true)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.NotAllowedHandler)
	router.NotFoundHandler = http.HandlerFunc(app.NotFoundHandler)
	router.Methods("GET").Path("/api/pastes").HandlerFunc(app.GetPastesHandler)
	router.Methods("GET").Path("/api/paste/{id:[0-9]+}").HandlerFunc(app.GetPasteHandler)
	router.Methods("POST").Path("/api/paste/add").HandlerFunc(app.AddPasteHandler)
	router.Methods("PATCH").Path("/api/paste/update").HandlerFunc(app.UpdatePasteHandler)
	router.Methods("DELETE").Path("/api/paste/delete/{id:[0-9]+}").HandlerFunc(app.DeletePasteHandler)
	// other handlers

	if app.Env != config.PROD_ENV {
		router.Methods("GET").PathPrefix("/api/docs/").Handler(httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprint("http://localhost:", app.Port, "/api/docs/doc.json")),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		))
	}

	// Security Middlewares
	secureMiddleware := secure.New(secure.Options{
		IsDevelopment:      app.Env == "DEV",
		ContentTypeNosniff: true,
		SSLRedirect:        true,
		// If the app is behind a proxy
		// SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})

	// Usual Middlewares
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.Use(negroni.HandlerFunc(middlewares.TrackRequestMiddleware))
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           86400,
	})
	// router with cors middleware
	wrappedRouter := corsMiddleware.Handler(router)
	n.UseHandler(wrappedRouter)

	startupMessage := "Starting API server (v" + app.Version + ")"
	startupMessage = startupMessage + " on port " + app.Port
	startupMessage = startupMessage + " in " + app.Env + " mode."
	logger.Log.Info(startupMessage)

	addr := ":" + app.Port
	if app.Env == "DEV" {
		addr = "0.0.0.0:" + app.Port
	}

	server := http.Server{
		Addr:         addr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      n,
	}

	logger.Log.Info("Listening...")

	server.ListenAndServe()
}

// OnShutdown is called when the server has a panic.
func (app *AppServer) OnShutdown() {
	// Do cleanup or logging
	logger.Log.Error("Executed OnShutdown")
}

// Special server handlers, outside of specific routes we have
func (app *AppServer) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := app.Sender.JSON(w, http.StatusNotFound, fmt.Sprint("Not Found ", r.URL))
	if err != nil {
		panic(err)
	}
}

func (app *AppServer) NotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	err := app.Sender.JSON(w, http.StatusMethodNotAllowed, fmt.Sprint(r.Method, " method not allowed"))
	if err != nil {
		panic(err)
	}
}

// cSpell:ignore negroni httputils Nosniff urfave sirupsen logrus
