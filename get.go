package env

import (
	"log"
	"os"
	"sync"
)

func Get(key string) string {
	var (
		once sync.Once
	)

	if !IsLoaded {
		if err := Load(); err != nil {
			once.Do(func() {
				log.Print("[warning] there is no .env file to load")
			})
		}
	}

	return os.Getenv(key)
}
