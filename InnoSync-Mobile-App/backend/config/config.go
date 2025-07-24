package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Upload   UploadConfig   `yaml:"upload"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Mode    string `yaml:"mode"` // gin mode: debug, release, test
	Timeout int    `yaml:"timeout"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey       string        `yaml:"secret_key"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
}

// UploadConfig holds file upload configuration
type UploadConfig struct {
	MaxFileSize int64  `yaml:"max_file_size"`
	UploadDir   string `yaml:"upload_dir"`
	AllowedExts []string `yaml:"allowed_extensions"`
}

// LoadConfig loads configuration from environment variables and config file
func LoadConfig() (*Config, error) {
	config := &Config{}

	// Load from config file if exists
	if _, err := os.Stat("config.yaml"); err == nil {
		if err := loadFromFile(config, "config.yaml"); err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	// Override with environment variables
	loadFromEnv(config)

	// Set defaults if not provided
	setDefaults(config)

	return config, nil
}

// loadFromFile loads configuration from YAML file
func loadFromFile(config *Config, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(config *Config) {
	// Server config
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Server.Port = p
		}
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		config.Server.Mode = mode
	}

	// Database config
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Database.Port = p
		}
	}
	if user := os.Getenv("DB_USER"); user != "" {
		config.Database.User = user
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		config.Database.DBName = dbname
	}
	if sslmode := os.Getenv("DB_SSLMODE"); sslmode != "" {
		config.Database.SSLMode = sslmode
	}

	// JWT config
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		config.JWT.SecretKey = secret
	}
	if accessTTL := os.Getenv("JWT_ACCESS_TTL"); accessTTL != "" {
		if ttl, err := time.ParseDuration(accessTTL); err == nil {
			config.JWT.AccessTokenTTL = ttl
		}
	}
	if refreshTTL := os.Getenv("JWT_REFRESH_TTL"); refreshTTL != "" {
		if ttl, err := time.ParseDuration(refreshTTL); err == nil {
			config.JWT.RefreshTokenTTL = ttl
		}
	}

	// Upload config
	if maxSize := os.Getenv("UPLOAD_MAX_SIZE"); maxSize != "" {
		if size, err := strconv.ParseInt(maxSize, 10, 64); err == nil {
			config.Upload.MaxFileSize = size
		}
	}
	if uploadDir := os.Getenv("UPLOAD_DIR"); uploadDir != "" {
		config.Upload.UploadDir = uploadDir
	}
}

// setDefaults sets default values for configuration
func setDefaults(config *Config) {
	// Server defaults
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}
	if config.Server.Mode == "" {
		config.Server.Mode = "debug"
	}
	if config.Server.Timeout == 0 {
		config.Server.Timeout = 30
	}

	// Database defaults
	if config.Database.Host == "" {
		config.Database.Host = "localhost"
	}
	if config.Database.Port == 0 {
		config.Database.Port = 5432
	}
	if config.Database.User == "" {
		config.Database.User = "postgres"
	}
	if config.Database.Password == "" {
		config.Database.Password = "postgres"
	}
	if config.Database.DBName == "" {
		config.Database.DBName = "innosync"
	}
	if config.Database.SSLMode == "" {
		config.Database.SSLMode = "disable"
	}

	// JWT defaults
	if config.JWT.SecretKey == "" {
		config.JWT.SecretKey = "your-secret-key-change-this-in-production"
	}
	if config.JWT.AccessTokenTTL == 0 {
		config.JWT.AccessTokenTTL = 15 * time.Minute
	}
	if config.JWT.RefreshTokenTTL == 0 {
		config.JWT.RefreshTokenTTL = 7 * 24 * time.Hour
	}

	// Upload defaults
	if config.Upload.MaxFileSize == 0 {
		config.Upload.MaxFileSize = 10 * 1024 * 1024 // 10MB
	}
	if config.Upload.UploadDir == "" {
		config.Upload.UploadDir = "uploads"
	}
	if len(config.Upload.AllowedExts) == 0 {
		config.Upload.AllowedExts = []string{".jpg", ".jpeg", ".png", ".pdf", ".doc", ".docx"}
	}
}

// GetDatabaseURL returns the database connection URL
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

// GetServerAddress returns the server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}