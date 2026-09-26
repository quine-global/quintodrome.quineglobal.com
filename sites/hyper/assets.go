package hyper

import "strings"

// classifyAsset maps an asset filename to an OS and architecture.
func classifyAsset(name string) (os, arch string, ok bool) {
	l := strings.ToLower(name)
	switch {
	case strings.HasSuffix(l, ".exe"):
		return "windows", "x64", true
	case strings.Contains(l, "arm64") && strings.HasSuffix(l, ".dmg"):
		return "mac", "arm64", true
	case strings.HasSuffix(l, ".dmg"):
		return "mac", "x64", true
	case strings.Contains(l, "arm64") && strings.HasSuffix(l, ".appimage"):
		return "linux", "arm64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".appimage"):
		return "", "", false
	case strings.HasSuffix(l, ".appimage"):
		return "linux", "x64", true
	case (strings.Contains(l, "arm64") || strings.Contains(l, "aarch64")) && strings.HasSuffix(l, ".rpm"):
		return "linux-rpm", "arm64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".rpm"):
		return "", "", false
	case strings.HasSuffix(l, ".rpm"):
		return "linux-rpm", "x64", true
	case strings.Contains(l, "arm64") && strings.HasSuffix(l, ".deb"):
		return "linux-deb", "arm64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".deb"):
		return "", "", false
	case strings.HasSuffix(l, ".deb"):
		return "linux-deb", "x64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".snap"):
		return "", "", false
	case strings.HasSuffix(l, ".snap"):
		return "linux-snap", "x64", true
	case (strings.Contains(l, "arm64") || strings.Contains(l, "aarch64")) && strings.HasSuffix(l, ".pacman"):
		return "linux-pacman", "arm64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".pacman"):
		return "", "", false
	case strings.HasSuffix(l, ".pacman"):
		return "linux-pacman", "x64", true
	case (strings.Contains(l, "arm64") || strings.Contains(l, "aarch64")) && strings.HasSuffix(l, ".flatpak"):
		return "linux-flatpak", "arm64", true
	case strings.Contains(l, "armv7l") && strings.HasSuffix(l, ".flatpak"):
		return "", "", false
	case strings.HasSuffix(l, ".flatpak"):
		return "linux-flatpak", "x64", true
	default:
		return "", "", false
	}
}
