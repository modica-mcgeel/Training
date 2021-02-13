package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/sctypes", *uc)
	//http.Handle("/sctypes/, *uc")
}
