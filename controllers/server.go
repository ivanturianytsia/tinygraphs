package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/controllers/isogrids"
	"github.com/ivanturianytsia/tinygraphs/controllers/spaceinvaders"
	"github.com/ivanturianytsia/tinygraphs/controllers/squares"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// router.Handle("/{type}/{key}", Handler)
// types = [squares, isogrids, spaceinvaders, hexa, hexa16]
func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	graphType := mux.Vars(r)["type"]
	switch graphType {
	case "squares":
		squares.Square(w, r)
		break
	case "isogrids":
		isogrids.Isogrids(w, r)
		break
	case "spaceinvaders":
		spaceinvaders.SpaceInvaders(w, r)
		break
	case "hexa":
		isogrids.Hexa(w, r)
		break
	case "hexa16":
		isogrids.Hexa16(w, r)
		break
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func Route(router *mux.Router) {
	router.HandleFunc("/squares/{key}", squares.Square)
	router.HandleFunc("/isogrids/{key}", isogrids.Isogrids)
	router.HandleFunc("/spaceinvaders/{key}", spaceinvaders.SpaceInvaders)
	router.HandleFunc("/labs/isogrids/hexa/{key}", isogrids.Hexa)
	router.HandleFunc("/labs/isogrids/hexa16/{key}", isogrids.Hexa16)
}
