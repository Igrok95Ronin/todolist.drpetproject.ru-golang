package routes

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Note struct {
	Title string `json:"title"` // заголовок заметки
}

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var note Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		h.logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(note.Title)
	json.NewEncoder(w).Encode(map[string]string{"message": "Okay"})
}
