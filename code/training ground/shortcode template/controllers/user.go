package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the controller"))
	if r.URL.Path == "/sctypes" {
		switch r.Method {
		case http.MethodPost:

		}
	}
}

/*func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//w.Write([]bytes("Could not parse ShortCode object"))
		w.Write([]byte("Could not parse"))
		return
	}

	u, err = models.NewSC(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	//encodeRequestAsJSON(u, w)
}

func (uc *userController) parseRequest(r *http.Request) (models.ShortCodeRequest, error) {
	dec := json.NewDecoder(r.Body)
	var sc models.ShortCodeRequest
	err := dec.Decode(&sc)
	if err != nil {
		return models.ShortCodeRequest{}, err
	}
	return sc, nil
}*/

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/sctypes/(\d+)/?`),
	}
}
