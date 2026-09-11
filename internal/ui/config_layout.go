// Package ui - config_layout.go
// The appearance.layout setting: the default explorer layout for new tabs.
package ui

import (
	"strings"

	"github.com/janosmiko/lfk/internal/logger"
)

// applyExplorerLayout validates and applies the appearance.layout config value.
// Empty keeps the compiled default. Unknown values warn and keep it too.
func applyExplorerLayout(raw string) {
	v := strings.ToLower(strings.TrimSpace(raw))
	if v == "" {
		return
	}
	switch v {
	case LayoutNormal, LayoutSidebarHidden, LayoutFullscreen:
		ConfigExplorerLayout = v
	default:
		logger.Warn("Invalid appearance.layout; using default",
			"accepted", []string{LayoutNormal, LayoutSidebarHidden, LayoutFullscreen},
			"default", LayoutNormal)
	}
}
