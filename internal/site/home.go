package site

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"

	"app/internal/component"
)

// Home registers the index and download handlers.
func (s *Server) Home(r chi.Router) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return s.config.HomeContent(s.config, component.PageProps{}), nil
	}))

	r.Get("/download", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		// On-demand refresh, debounced to once per 2 minutes across all requests.
		go s.releases.TryRefresh(context.Background(), 2*time.Minute)

		detectedOS, detectedArch := s.detectPlatform(r.Header.Get("User-Agent"))

		selectedOS := detectedOS
		selectedArch := detectedArch

		if p := r.URL.Query().Get("os"); p == "mac" || p == "windows" || p == "linux" || p == "linux-rpm" || p == "linux-deb" || p == "linux-snap" || p == "linux-pacman" || p == "linux-flatpak" {
			selectedOS = p
		}
		archExplicit := false
		if a := r.URL.Query().Get("arch"); a == "arm64" || a == "x64" {
			selectedArch = a
			archExplicit = true
		}

		return component.DownloadPage(s.config.Branding, s.config.DownloadConfig, component.PageProps{}, component.DownloadProps{
			DetectedOS:   detectedOS,
			DetectedArch: detectedArch,
			SelectedOS:   selectedOS,
			SelectedArch: selectedArch,
			ArchExplicit: archExplicit,
			Releases:     s.releases.Get(),
		}), nil
	}))
}

func (s *Server) detectPlatform(ua string) (os, arch string) {
	u := strings.ToLower(ua)
	switch {
	case strings.Contains(u, "windows"):
		return "windows", "x64"
	case strings.Contains(u, "macintosh"), strings.Contains(u, "mac os"):
		return "mac", "arm64" // default Apple Silicon; user can override
	case strings.Contains(u, "linux"):
		return s.config.DefaultLinuxOS, "x64"
	default:
		return "mac", "arm64"
	}
}
