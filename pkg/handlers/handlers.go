package handlers

import (
	"net/http"

	"github.com/arc2501/bnb/pkg/config"
	"github.com/arc2501/bnb/pkg/models"
	"github.com/arc2501/bnb/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Giving Handlers a reciever so that all of them can have access to app config
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Extracting IP address of the caller
	remoteIP := r.RemoteAddr
	// Putting/Storing that remote IP as key value pair
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Heyyy Kaimchho Majama"

	//Fetching the remoteIP from the session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
