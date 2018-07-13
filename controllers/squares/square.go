package squares

import (
	"crypto/md5"
	"fmt"
	"image"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ivanturianytsia/tinygraphs/draw/squares"
	"github.com/ivanturianytsia/tinygraphs/extract"
	"github.com/ivanturianytsia/tinygraphs/format"
	"github.com/ivanturianytsia/tinygraphs/write"
)

// Square handler for /squares/:key
// builds a 6 by 6 grid with alternate colors based the key passed in the url.
func Square(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]

	theme := extract.Theme(r)

	h := md5.New()
	io.WriteString(h, key)
	key = fmt.Sprintf("%x", h.Sum(nil)[:])

	e := `"` + theme + key + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	colors := extract.Colors(r)
	size := extract.Size(r)

	if f := extract.Format(r); f == format.JPEG {
		m := image.NewRGBA(image.Rect(0, 0, size, size))
		squares.Image(m, key, colors)
		var img image.Image = m
		write.ImageJPEG(w, &img)
	} else if f == format.SVG {
		write.ImageSVG(w)
		squares.SVG(w, key, colors, size)
	}
}
