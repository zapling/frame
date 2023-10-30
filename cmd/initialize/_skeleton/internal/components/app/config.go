package app

type Config struct {
	IsRunningInDevMode bool // Is running via 'frame serve'
	DevServerPort      int  // Development server port via 'frame serve'
}
