package handler

import (
	"fmt"
	"net/http"

	"github.com/bihe/commons-go/security"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

// --------------------------------------------------------------------------
// Objects
// --------------------------------------------------------------------------

// Meta specifies application metadata
// swagger:model
type Meta struct {
	Version string `json:"version"`
	UserInfo
}

// UserInfo provides information about the currently logged-in user
type UserInfo struct {
	Email       string   `json:"email"`
	DisplayName string   `json:"displayName"`
	Roles       []string `json:"roles"`
}

// --------------------------------------------------------------------------
// Request and Response objects using go-chi render
// --------------------------------------------------------------------------

// Response wraps the data struct into a framework response
type AppInfoResponse struct {
	*Meta
}

// Render the specific response
func (a AppInfoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// --------------------------------------------------------------------------
// Handler implementation
// --------------------------------------------------------------------------

// AppInfoHandler is responsible to return meta-information about the application
type AppInfoHandler struct {
	Handler
	// Version of the application using https://semver.org/
	Version string
	// Build identifies the specific build, e.g. git-hash
	Build string
}

// swagger:operation GET /appinfo appinfo HandleAppInfo
//
// provides information about the application
//
// meta-data of the application including authenticated user and version
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Meta
//     schema:
//       "$ref": "#/definitions/Meta"
//   '401':
//     description: ProblemDetail
//     schema:
//       "$ref": "#/definitions/ProblemDetail"
//   '403':
//     description: ProblemDetail
//     schema:
//       "$ref": "#/definitions/ProblemDetail"
func (a *AppInfoHandler) HandleAppInfo(user security.User, w http.ResponseWriter, r *http.Request) error {
	log.WithField("func", "api.HandleAppInfo").Debugf("return the application metadata info")
	info := Meta{
		Version: fmt.Sprintf("%s-%s", a.Version, a.Build),
		UserInfo: UserInfo{
			Email:       user.Email,
			DisplayName: user.DisplayName,
			Roles:       user.Roles,
		},
	}
	render.Render(w, r, AppInfoResponse{Meta: &info})
	return nil
}
