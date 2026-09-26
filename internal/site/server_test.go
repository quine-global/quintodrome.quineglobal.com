package site_test

import (
	"io"
	http2 "net/http"
	"strings"
	"testing"
	"time"

	"maragu.dev/is"

	. "maragu.dev/gomponents"

	"app/internal/component"
	"app/internal/site"
)

func testConfig() site.Config {
	return site.Config{
		Branding: component.Branding{
			Name:      "Test",
			LogoPath:  "/images/test-logo.png",
			GitHubURL: "https://github.com/quine-global/test",
		},
		DownloadConfig: component.DownloadConfig{
			DefaultLinuxOS: "linux",
		},
		RepoOwner:     "quine-global",
		RepoName:      "test",
		ClassifyAsset: func(name string) (string, string, bool) { return "", "", false },
		HomeContent: func(cfg site.Config, props component.PageProps) Node {
			return component.Page(cfg.Branding, props)
		},
	}
}

func TestServer_Start(t *testing.T) {
	t.Run("can start and stop server", func(t *testing.T) {
		s := site.NewServer(site.NewServerOptions{Config: testConfig()})

		go func() {
			is.NotError(t, s.Start())
		}()
		defer func() {
			is.NotError(t, s.Stop())
		}()

		// I know we could check that the server is running here, but it's easier to just wait a bit
		time.Sleep(10 * time.Millisecond)

		res, err := http2.Get("http://localhost:8081/")
		is.NotError(t, err)
		is.Equal(t, http2.StatusOK, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		is.NotError(t, err)
		is.True(t, strings.HasPrefix(string(body), "<!doctype html>"))
	})
}
