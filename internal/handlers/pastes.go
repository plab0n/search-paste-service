package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/plab0n/search-paste/internal/model"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) GetPastesHandler(w http.ResponseWriter, r *http.Request) {
}
func (h *Handlers) GetPasteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	paste, err := h.Storage.GetPaste(ctx, id)

	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err)
		return
	}
	h.Sender.JSON(w, http.StatusOK, paste)
	return
}
func (h *Handlers) AddPasteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var paste model.Paste
	err := json.NewDecoder(r.Body).Decode(&paste)

	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = Validate.Struct(paste)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err.Field()+" "+err.Tag())
		}
		h.Sender.JSON(w, http.StatusBadRequest, strings.Join(errs, ", "))
		return
	}
	id, err := h.Storage.AddPaste(ctx, model.AddPasteRequest{Title: paste.Title, Text: paste.Text})

	if id == 0 {
		h.Sender.JSON(w, http.StatusInternalServerError, http.NoBody)
		return
	}
	response := model.IDResponse{ID: id}
	h.Sender.JSON(w, http.StatusCreated, response)
	return
}

func (h *Handlers) UpdatePasteHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeletePasteHandler(w http.ResponseWriter, r *http.Request) {

}
