package routers

import (
	"encoding/json"
	"github.com/joselimaico/twitt/models"
	"net/http"
	"strconv"

	"github.com/joselimaico/twitt/bd"
)

/*LeoTweets Leo los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	var size int
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("size")) < 1 {
		size = 20
	} else {
		size, err = strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil {
			http.Error(w, "Debe enviar números en el parámetro size", http.StatusBadRequest)
			return
		}
	}

	pag := int64(pagina)
	pageSize := int64(size)
	respuesta, correcto := bd.LeoTweets(ID, pag, pageSize)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	if respuesta == nil {
		respuesta = make([]*models.DevuelvoTweets, 0)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
