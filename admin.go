package admin

import (
	"context"

	gmcore_bundle "github.com/gmcorenet/sdk/gmcore-bundle"
	gmcore_crud "github.com/gmcorenet/sdk/gmcore-crud"
)

func init() {
}

type MenuItem struct {
	Label    string     `json:"label" yaml:"label"`
	Route    string     `json:"route" yaml:"route"`
	Icon     string     `json:"icon" yaml:"icon"`
	Order    int        `json:"order" yaml:"order"`
	Role     string     `json:"role" yaml:"role"`
	Children []MenuItem `json:"children" yaml:"children"`
	URL      string     `json:"url" yaml:"url"`
}

type Dashboard struct {
	Path     string   `json:"path" yaml:"path"`
	Title    string   `json:"title" yaml:"title"`
	Widgets  []string `json:"widgets" yaml:"widgets"`
	Theme    string   `json:"theme" yaml:"theme"`
}

type Branding struct {
	Name    string `json:"name" yaml:"name"`
	Logo    string `json:"logo" yaml:"logo"`
	Favicon string `json:"favicon" yaml:"favicon"`
}

type AdminCRUDResource struct {
	Name       string
	Entity     interface{}
	Config     gmcore_crud.Config
	AutoFields bool
}

type Bundle struct {
	gmcore_bundle.BaseBundle
	Dashboard Dashboard
	Menu      []MenuItem
	Branding  Branding
	Resources []AdminCRUDResource
}

func NewBundle() *Bundle {
	return &Bundle{
		Menu:      make([]MenuItem, 0),
		Resources: make([]AdminCRUDResource, 0),
	}
}

func (b *Bundle) Name() string { return "admin-bundle" }

func (b *Bundle) Boot(ctx context.Context) error {
	return nil
}

func (b *Bundle) Shutdown() error {
	return nil
}

func (b *Bundle) AddMenuItem(item MenuItem) {
	b.Menu = append(b.Menu, item)
}

func (b *Bundle) RegisterResource(res AdminCRUDResource) {
	b.Resources = append(b.Resources, res)
}

func (b *Bundle) BuildMenu() []MenuItem {
	return b.Menu
}
