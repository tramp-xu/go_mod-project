package service

import (
	"fmt"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router  {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		fmt.Println(route)
		router.Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return router
}
