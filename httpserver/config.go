package httpserver

// Config http server configuration
type Config struct {
	Port                 int  `yaml:"port"`
	EnableAccessLogs     bool `yaml:"enable-access-logs"`
	EnableHealthEndpoint bool `yaml:"enable-health-endpoint"`
}
