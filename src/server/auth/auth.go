package auth

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func Validate(user, pass, file string) bool {
	f, err := os.Open(file)
	if err != nil {
		return false
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	hashed := fmt.Sprintf("%x", sha256.Sum256([]byte(pass)))

	for scan.Scan() {
		line := scan.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		u := strings.TrimSpace(parts[0])
		p := strings.TrimSpace(parts[1])

		if u == user && p == hashed {
			return true
		}
	}
	return false
}
