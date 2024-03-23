package config

type Config struct {
	Author struct {
		AdminUser     string
		AdminPassword string
	}
	WebUI struct {
		Addr string
	}
}
