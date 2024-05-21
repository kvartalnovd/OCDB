package app

type AppOpts struct {
	Version string
}

type App struct {
	version string
}

func NewApp(opts AppOpts) *App {
	return &App{
		version: opts.Version,
	}
}
