// Package var1 provides ... for Ambient apps.
package var1

import (
	"html/template"
	"io"
	"net/http"

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
func (p *Plugin) PluginName() string {
	return "var1"
}

// PluginVersion returns the plugin version.
func (p *Plugin) PluginVersion() string {
	return "1.0.0"
}

// Logger returns a logger.
func (p *Plugin) Logger(appName string, appVersion string, optionalWriter io.Writer) (ambient.AppLogger, error) {
	return nil, nil
}

// Storage returns data and session storage.
func (p *Plugin) Storage(logger ambient.Logger) (ambient.DataStorer, ambient.SessionStorer, error) {
	return nil, nil, nil
}

// Router returns a router.
func (p *Plugin) Router(logger ambient.Logger, te ambient.Renderer) (ambient.AppRouter, error) {
	return nil, nil
}

// SessionManager returns a session manager.
func (p *Plugin) SessionManager(logger ambient.Logger, ss ambient.SessionStorer) (ambient.AppSession, error) {
	return nil, nil
}

// Routes sets routes for the plugin.
func (p *Plugin) Routes() {}

// Middleware returns router middleware.
func (p *Plugin) Middleware() []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{}
}

// TemplateEngine returns a template engine.
func (p *Plugin) TemplateEngine(logger ambient.Logger, injector ambient.AssetInjector) (ambient.Renderer, error) {
	return nil, nil
}

// GrantRequests returns a list of grants requested by the plugin.
func (p *Plugin) GrantRequests() []ambient.GrantRequest {
	return []ambient.GrantRequest{}
}

// Settings returns a list of plugin settings.
func (p *Plugin) Settings() []ambient.Setting {
	return []ambient.Setting{}
}

// Assets returns a list of assets and an embedded filesystem.
func (p *Plugin) Assets() ([]ambient.Asset, ambient.FileSystemReader) {
	return []ambient.Asset{}, nil
}

// FuncMap returns a callable function that accepts a request.
func (p *Plugin) FuncMap() func(r *http.Request) template.FuncMap {
	return func(r *http.Request) template.FuncMap {
		fm := make(template.FuncMap)
		return fm
	}
}
