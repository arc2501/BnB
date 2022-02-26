package handlers

import (
	"net/http"

	"github.com/arc2501/bnb/pkg/config"
	"github.com/arc2501/bnb/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")

}
