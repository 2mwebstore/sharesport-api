package routes

import (
	"wwb99/controllers"
	"wwb99/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// -----------------
	// Public routes
	// -----------------
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/refresh", controllers.RefreshToken).Methods("POST")

	// Client routes (public)
	r.HandleFunc("/api/news_home", controllers.GetNewsHome).Methods("GET")
	r.HandleFunc("/api/highlights_home", controllers.GetHighlightsHome).Methods("GET")
	r.HandleFunc("/api/footers_home", controllers.GetFootersHome).Methods("GET")
	r.HandleFunc("/api/sponsors_home", controllers.GetSponsorsHome).Methods("GET")
	r.HandleFunc("/api/news/getbyid", controllers.GetNewsByID).Methods("GET")

	// -----------------
	// Admin routes (secured)
	// -----------------
	admin := r.PathPrefix("/api").Subrouter()
	admin.Use(middleware.AuthMiddleware)

	// News
	admin.HandleFunc("/news", controllers.GetNews).Methods("GET")
	admin.HandleFunc("/news/create", controllers.CreateNews).Methods("POST")
	admin.HandleFunc("/news/update/{id}", controllers.UpdateNews).Methods("PUT")
	admin.HandleFunc("/news/delete", controllers.DeleteNews)

	// Highlights
	admin.HandleFunc("/highlights", controllers.GetHighlights).Methods("GET")
	admin.HandleFunc("/highlights/create", controllers.CreateHighlights).Methods("POST")
	admin.HandleFunc("/highlights/update", controllers.UpdateHighlights).Methods("PUT")
	admin.HandleFunc("/highlights/delete", controllers.DeleteHighlights)
	admin.HandleFunc("/highlights/getbyid", controllers.GetHighlightsByID)

	// Footers
	admin.HandleFunc("/footers", controllers.GetFooters).Methods("GET")
	admin.HandleFunc("/footers/create", controllers.CreateFooter).Methods("POST")
	admin.HandleFunc("/footers/update", controllers.UpdateFooter).Methods("PUT")
	admin.HandleFunc("/footers/delete", controllers.DeleteFooter)
	admin.HandleFunc("/footers/getbyid", controllers.GetFooterByID)

	// Sponsors
	admin.HandleFunc("/sponsors", controllers.GetSponsors).Methods("GET")
	admin.HandleFunc("/sponsors/create", controllers.CreateSponsor).Methods("POST")
	admin.HandleFunc("/sponsors/update", controllers.UpdateSponsor).Methods("PUT")
	admin.HandleFunc("/sponsors/delete", controllers.DeleteSponsor)
	admin.HandleFunc("/sponsors/getbyid", controllers.GetSponsorByID)

	// Permissions
	admin.HandleFunc("/permissions", controllers.GetPermissions).Methods("GET")
	admin.HandleFunc("/permissions/create", controllers.CreatePermission).Methods("POST")
	admin.HandleFunc("/permissions/update", controllers.UpdatePermission).Methods("PUT")

	// Roles
	admin.HandleFunc("/roles", controllers.GetRoles).Methods("GET")
	admin.HandleFunc("/roles", controllers.CreateRole).Methods("POST")
	admin.HandleFunc("/roles", controllers.UpdateRole).Methods("PUT")
	admin.HandleFunc("/roles", controllers.DeleteRole).Methods("DELETE")
	admin.HandleFunc("/roles/getbyid", controllers.GetRoleByID).Methods("GET")
	admin.HandleFunc("/roles/permissions", controllers.GetPermissionRoles).Methods("GET")
	admin.HandleFunc("/roles/assign", controllers.AssignPermissions).Methods("PUT")

	// Example secured profile route
	admin.HandleFunc("/profile", controllers.Profile).Methods("GET")

	return r
}
