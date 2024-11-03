package api

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"net/http"
	"strconv"
	"strings"

	"github.com/FabioSM46/placeholder-image-api/imagegen"
)

func parseHexColor(s string) (color.RGBA, error) {
	c := color.RGBA{0, 0, 0, 255} // Default to black if parsing fails
	if len(s) == 6 {
		r, err1 := strconv.ParseUint(s[0:2], 16, 8)
		g, err2 := strconv.ParseUint(s[2:4], 16, 8)
		b, err3 := strconv.ParseUint(s[4:6], 16, 8)
		if err1 != nil || err2 != nil || err3 != nil {
			return c, fmt.Errorf("invalid hex color")
		}
		c = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	}
	return c, nil
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the path and split it into parts
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	// Variables for color and dimensions
	var hexColor string
	var dimensions string

	// Determine if color is provided or only dimensions
	if len(parts) == 2 {
		hexColor = parts[0]
		dimensions = parts[1]
	} else if len(parts) == 1 {
		hexColor = "" // Use default color
		dimensions = parts[0]
	} else {
		http.Error(w, "Invalid format. Use /<hexcolor>/widthxheight or /widthxheight", http.StatusBadRequest)
		return
	}

	// Parse dimensions
	dimParts := strings.Split(dimensions, "x")
	if len(dimParts) != 2 {
		http.Error(w, "Invalid dimensions", http.StatusBadRequest)
		return
	}

	width, err1 := strconv.Atoi(dimParts[0])
	height, err2 := strconv.Atoi(dimParts[1])
	if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
		http.Error(w, "Invalid dimensions", http.StatusBadRequest)
		return
	}

	// Set the default background color to gray
	bgColor := color.RGBA{128, 128, 128, 255} // Default gray

	// Parse the hex color if provided
	if hexColor != "" {
		var err error
		bgColor, err = parseHexColor(hexColor)
		if err != nil {
			http.Error(w, "Invalid hex color format", http.StatusBadRequest)
			return
		}
	}

	// Generate the image with the specified background color
	img := imagegen.GenerateImage(width, height, bgColor)

	// Encode and write the image to the response
	w.Header().Set("Content-Type", "image/jpeg")
	if err := jpeg.Encode(w, img, nil); err != nil {
		http.Error(w, "Could not encode image", http.StatusInternalServerError)
	}
}