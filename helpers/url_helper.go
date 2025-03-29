package helpers

import (
	"fmt"
	"strings"

	"github.com/fauzan264/evermos-rakamin/config"
)

func GetImageURL(path string) string {
	cfg := config.LoadConfig()

	if !strings.HasPrefix(cfg.AppHost, "http://") && !strings.HasPrefix(cfg.AppHost, "https://") {
		cfg.AppHost = "http://" + cfg.AppHost
	}

	baseURL := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)

	// Jika path sudah full URL, langsung return
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}

	return fmt.Sprintf("%s/%s", baseURL, path)
}