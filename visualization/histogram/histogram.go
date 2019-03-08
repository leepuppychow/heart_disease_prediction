package histogram

import (
	"net/http"

	"github.com/ajstarks/svgo"
)

func Histogram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Rect(0, 0, 100, 20, "fill:red")
	s.Rect(0, 20, 200, 20, "fill:blue")
	s.Rect(0, 40, 150, 20, "fill:green")
	s.End()
}
