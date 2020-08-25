package proxy

import "bytes"

type App struct {
}

func (a *App) Do() string {
	return "App is running"
}

// We need to abstract out method first

type Iapp interface {
	Do() string
}

type AppProxy struct {
	app Iapp
}

func NewAppProxy(app Iapp) *AppProxy {
	return &AppProxy{
		app: app,
	}
}

func (a *AppProxy) Do() string {
	var buf bytes.Buffer
	buf.WriteString("Before:")
	// call here
	buf.WriteString(a.app.Do())
	buf.WriteString("#After")
	return buf.String()
}
