package quintodrome

import (
	"testing"

	"maragu.dev/is"
)

func TestClassifyAsset(t *testing.T) {
	cases := []struct {
		name     string
		wantOS   string
		wantArch string
		wantOK   bool
	}{
		// Windows — real asset: Quintodrome_1.0.0-alpha_x64-setup.exe
		{"Quintodrome_1.0.0-alpha_x64-setup.exe", "windows", "x64", true},
		{"Quintodrome.Setup.1.0.0.exe", "windows", "x64", true},

		// Windows MSI — not classified; the .exe installer is preferred
		{"Quintodrome_1.0.0-alpha_x64_en-US.msi", "", "", false},

		// macOS — real assets: Quintodrome_0.1.0_x64.dmg, Quintodrome_0.1.0_aarch64.dmg
		{"Quintodrome_0.1.0_x64.dmg", "mac", "x64", true},
		{"Quintodrome_0.1.0_aarch64.dmg", "mac", "arm64", true},
		{"Quintodrome-4.0.0-arm64.dmg", "mac", "arm64", true},

		// macOS .app.zip — not classified; the .dmg is preferred
		{"Quintodrome-x86_64-apple-darwin.app.zip", "", "", false},
		{"Quintodrome-aarch64-apple-darwin.app.zip", "", "", false},

		// Linux AppImage x64
		{"Quintodrome-4.0.0.AppImage", "linux", "x64", true},
		{"quintodrome-4.0.0.appimage", "linux", "x64", true},

		// Linux AppImage arm64
		{"Quintodrome-4.0.0-arm64.AppImage", "linux", "arm64", true},

		// Linux AppImage armv7l — must NOT be classified as x64
		{"Quintodrome-4.0.0-armv7l.AppImage", "", "", false},
		{"quintodrome-4.0.0-armv7l.appimage", "", "", false},

		// Linux RPM x64 — real asset uses x86_64 suffix
		{"Quintodrome-4.0.0.rpm", "linux-rpm", "x64", true},
		{"Quintodrome-4.0.0-x86_64.rpm", "linux-rpm", "x64", true},

		// Linux RPM arm64 — real asset uses aarch64, not arm64
		{"Quintodrome-4.0.0-arm64.rpm", "linux-rpm", "arm64", true},
		{"Quintodrome-4.0.0-aarch64.rpm", "linux-rpm", "arm64", true},

		// Linux RPM armv7l — must NOT be classified as x64
		{"Quintodrome-4.0.0-armv7l.rpm", "", "", false},

		// Linux DEB x64 — real asset uses amd64 suffix
		{"Quintodrome-4.0.0-amd64.deb", "linux-deb", "x64", true},

		// Linux DEB arm64
		{"Quintodrome-4.0.0-arm64.deb", "linux-deb", "arm64", true},

		// Linux DEB armv7l — must NOT be classified as x64
		{"Quintodrome-4.0.0-armv7l.deb", "", "", false},

		// Linux Snap — real assets: Quintodrome-linux-x86_64.snap, Quintodrome-linux-aarch64.snap
		{"Quintodrome-linux-x86_64.snap", "linux-snap", "x64", true},
		{"Quintodrome-linux-aarch64.snap", "linux-snap", "arm64", true},

		// Linux Snap armv7l — must NOT be classified as x64
		{"Quintodrome-linux-armv7l.snap", "", "", false},

		// Linux Pacman x64
		{"Quintodrome-4.0.0-x64.pacman", "linux-pacman", "x64", true},

		// Linux Pacman arm64 — real asset uses aarch64 suffix
		{"Quintodrome-4.0.0-aarch64.pacman", "linux-pacman", "arm64", true},

		// Linux Pacman armv7l — must NOT be classified as x64
		{"Quintodrome-4.0.0-armv7l.pacman", "", "", false},

		// Linux Flatpak — real assets: Quintodrome-linux-x86_64.flatpak, Quintodrome-linux-aarch64.flatpak
		{"Quintodrome-linux-x86_64.flatpak", "linux-flatpak", "x64", true},
		{"Quintodrome-linux-aarch64.flatpak", "linux-flatpak", "arm64", true},

		// Linux Flatpak armv7l — must NOT be classified as x64
		{"Quintodrome-linux-armv7l.flatpak", "", "", false},

		// Unknown / unrecognised — e.g. navidrome server binaries and checksums
		{"navidrome_1.0.0-alpha_linux_amd64.tar.gz", "", "", false},
		{"navidrome_checksums.txt", "", "", false},
		{"latest.yml", "", "", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotOS, gotArch, gotOK := classifyAsset(c.name)
			is.Equal(t, c.wantOK, gotOK)
			is.Equal(t, c.wantOS, gotOS)
			is.Equal(t, c.wantArch, gotArch)
		})
	}
}
