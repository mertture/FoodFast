package controllers

import "github.com/mertture/FoodFast/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/restaurants", middlewares.SetMiddlewareJSON(s.CreateRestaurant)).Methods("POST")
	s.Router.HandleFunc("/restaurants", middlewares.SetMiddlewareJSON(s.GetRestaurants)).Methods("GET")
	s.Router.HandleFunc("/restaurants/{id}", middlewares.SetMiddlewareJSON(s.GetRestaurant)).Methods("GET")
	s.Router.HandleFunc("/restaurants/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateRestaurant))).Methods("PUT")
	s.Router.HandleFunc("/restaurants/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteRestaurant)).Methods("DELETE")
}
