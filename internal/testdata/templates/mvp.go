// Package var1 provides ... for Ambient apps.
package var1

import (
	"context"

	"github.com/ambientkit/ambient"
)

// Plugin represents an Ambient plugin.
type Plugin struct {
	*ambient.PluginBase
}

// New returns a new var1 plugin.
func New() *Plugin {
	return &Plugin{
		PluginBase: &ambient.PluginBase{},
	}
}

// PluginName returns the plugin name.
func (p *Plugin) PluginName(context.Context) string {
	return "var1"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion(context.Context) string {
	return "1.0.0"
}
