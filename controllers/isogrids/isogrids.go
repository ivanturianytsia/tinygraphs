package isogrids

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/draw/isogrids"
	"github.com/ivanturianytsia/tinygraphs/extract"
	"github.com/ivanturianytsia/tinygraphs/write"
)

// Isogrids is the handler for /isogrids/:key
// builds a 10x10 grid with alternate colors based on the string passed in the url.
func Isogrids(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	colors := extract.Colors(r)
	size := extract.Size(r)
	lines := extract.Lines(r)

	write.ImageSVG(w)
	isogrids.Isogrids(w, key, colors, size, lines)
}
