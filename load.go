package env

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	IsLoaded         bool
	HasCustomProfile bool
)

func Load(path ...string) error {
	if IsLoaded {
		return nil
	}

	var (
		file   = ".env"
		target string
	)

	if len(path) > 0 {
		target = path[0]
	}

	def := filepath.Join(target, file)

	// Loads default .env
	if err := godotenv.Load(def); err != nil {
		return err
	}

	profile := os.Getenv("PROFILE")

	// No custom profile to load
	if profile == "" {
		IsLoaded, HasCustomProfile = true, false

		return nil
	}

	custom := filepath.Join(target, file+"."+profile)

	if err := godotenv.Load(custom); err != nil {
		return err
	}

	IsLoaded, HasCustomProfile = true, true

	return nil
}
