package isogrids

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/cache"
	"github.com/ivanturianytsia/tinygraphs/draw/isogrids"
	"github.com/ivanturianytsia/tinygraphs/extract"
	"github.com/ivanturianytsia/tinygraphs/write"
)

// Hexa is the handler for /isogrids/hexa/:key
// builds an hexagon with alternate colors.
func Hexa(w http.ResponseWriter, r *http.Request) {
	size := extract.Size(r)

	key := mux.Vars(r)["key"]
	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	lines := extract.Hexalines(r)
	colors := extract.Colors(r)

	if Cache.IsCached(&w, r, key, colors, size) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	write.ImageSVG(w)
	isogrids.Hexa(w, key, colors, size, lines)
}
