package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/sctypes", *uc)
	http.Handle("/sctypes/", *uc)
}

/*func encodeRequestAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}*/
