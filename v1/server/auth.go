package server

import (
	// "fmt"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"strconv"
	fiber "github.com/gofiber/fiber/v2"
	// encryption "github.com/0187773933/encryption/v1/encryption"
)

func abs( x int64 ) int64 {
	if x < 0 {
		return -x
	}
	return x
}

const RequestMaxAgeSec int64 = 60

func ( s *Server ) ValidateAPIKey() ( fiber.Handler ) {
	return func( c *fiber.Ctx ) error {
		apiKey := strings.TrimSpace( c.Get( HEADER_API_KEY ) )
		sig := strings.TrimSpace( c.Get( HEADER_SIGNATURE ) )
		ts := strings.TrimSpace( c.Get( HEADER_TIMESTAMP ) )

		if apiKey == "" || sig == "" || ts == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("missing headers")
		}
		if apiKey != s.Config.ServerAPIKey {
			return c.Status(fiber.StatusUnauthorized).SendString("invalid key")
		}

		reqTs, err := strconv.ParseInt(ts, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("bad timestamp")
		}
		if abs( time.Now().Unix()-reqTs) > RequestMaxAgeSec {
			return c.Status(fiber.StatusUnauthorized).SendString("stale request")
		}

		body := c.Body()
		bodyHash := sha256.Sum256(body)
		bodyHex := hex.EncodeToString(bodyHash[:])

		msg := strings.Join([]string{
			c.Method(),
			c.Path(),
			ts,
			bodyHex,
		}, "\n")

		mac := hmac.New(sha256.New, []byte(s.Config.ServerAPISecret))
		mac.Write([]byte(msg))
		expected := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(sig), []byte(expected)) {
			return c.Status(fiber.StatusUnauthorized).SendString("bad signature")
		}
		return c.Next()
	}
}