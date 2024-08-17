package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "It`s Homes")
	fmt.Println(r.URL)
}
