package cfg

import (
	"fmt"
	"os"
	"strconv"
)

const (
	FrameDevMode = "FRAME_DEV_MODE"
	FrameDevPort = "FRAME_DEV_PORT"
)

type Values struct {
	// Development config values
	IsDevMode     bool // Is running via the 'frame serve'
	DevServerPort int  // Port number of the 'frame serve' development server
}

func Get() (Values, error) {
	isDevMode, err := strconv.ParseBool(getEnv(FrameDevMode, "false"))
	if err != nil {
		return Values{}, fmt.Errorf("Failed to load env variable 'FRAME_DEV_MODE': %w", err)
	}

	devServerPort, err := strconv.Atoi(getEnv(FrameDevPort, "4000"))
	if err != nil {
		return Values{}, fmt.Errorf("Failed to load env variable 'FRAME_PORT': %w", err)
	}

	return Values{
		IsDevMode:     isDevMode,
		DevServerPort: devServerPort,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
