package controllers

import (
	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/controllers/isogrids"
	"github.com/ivanturianytsia/tinygraphs/controllers/spaceinvaders"
	"github.com/ivanturianytsia/tinygraphs/controllers/squares"
)

func Route(router *mux.Router) {
	router.HandleFunc("/squares/{key}", squares.Square) //cached
	router.HandleFunc("/isogrids/{key}", isogrids.Isogrids)
	router.HandleFunc("/spaceinvaders/{key}", spaceinvaders.SpaceInvaders)
	router.HandleFunc("/labs/isogrids/hexa/{key}", isogrids.Hexa)
	router.HandleFunc("/labs/isogrids/hexa16/{key}", isogrids.Hexa16)
}
