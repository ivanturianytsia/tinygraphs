package spaceinvaders

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/cache"
	"github.com/ivanturianytsia/tinygraphs/draw/spaceinvaders"
	"github.com/ivanturianytsia/tinygraphs/extract"
	"github.com/ivanturianytsia/tinygraphs/write"
)

// SpaceInvaders handler for /spaceinvaders/:key
func SpaceInvaders(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colors := extract.Colors(r)
	size := extract.Size(r)

	if Cache.IsCached(&w, r, key, colors, size) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	write.ImageSVG(w)
	spaceinvaders.SpaceInvaders(w, key, colors, size)
}
