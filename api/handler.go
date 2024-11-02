package api

import (
	"image/jpeg"
	"net/http"
	"strconv"
	"strings"

	"github.com/FabioSM46/placeholder-image-api/imagegen"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract width and height from URL
	path := strings.TrimPrefix(r.URL.Path, "/")
	dimensions := strings.Split(path, "x")
	if len(dimensions) != 2 {
		http.Error(w, "Invalid dimensions", http.StatusBadRequest)
		return
	}

	width, err1 := strconv.Atoi(dimensions[0])
	height, err2 := strconv.Atoi(dimensions[1])
	if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
		http.Error(w, "Invalid dimensions", http.StatusBadRequest)
		return
	}

	img := imagegen.GenerateImage(width, height)

	w.Header().Set("Content-Type", "image/jpeg")
	if err := jpeg.Encode(w, img, nil); err != nil {
		http.Error(w, "Could not encode image", http.StatusInternalServerError)
	}
}