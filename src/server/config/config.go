package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	AllowedIP   string
	Port        string
	MaxAttempts int
	UsersFile   string
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])

		switch k {
		case "allowed_ip":
			cfg.AllowedIP = v
		case "port":
			cfg.Port = v
		case "max_attempts":
			cfg.MaxAttempts = atoi(v)
		case "users_file":
			cfg.UsersFile = v
		}
	}

	return cfg, nil
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}
