package api

import "net/http"

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
    svg := `
    <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 32 32">
        <circle cx="16" cy="16" r="16" fill="#3498db"/>
        <text x="16" y="21" font-size="16" text-anchor="middle" fill="#ffffff" font-family="Arial" font-weight="bold">API</text>
    </svg>
    `
    w.Header().Set("Content-Type", "image/svg+xml")
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte(svg))
}