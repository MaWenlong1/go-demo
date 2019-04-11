package user

import (
	"mwl/gravitee/util/routes"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers route handlers for the user service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRoutes(s.GetRoutes(), subRouter)
}

// GetRoutes returns []routes.Route slice for the user service
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "create_user",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: s.createUser,
		},
	}
}
