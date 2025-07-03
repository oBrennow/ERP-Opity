package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config representa a configuração completa do sistema
type Config struct {
	Server        ServerConfig       `mapstructure:"server"`
	Database      DatabaseConfig     `mapstructure:"database"`
	Security      SecurityConfig     `mapstructure:"security"`
	Logging       LoggingConfig      `mapstructure:"logging"`
	Modules       []ModuleConfig     `mapstructure:"modules"`
	Cache         CacheConfig        `mapstructure:"cache"`
	Email         EmailConfig        `mapstructure:"email"`
	Upload        UploadConfig       `mapstructure:"upload"`
	Notifications NotificationConfig `mapstructure:"notifications"`
	Backup        BackupConfig       `mapstructure:"backup"`
	Monitoring    MonitoringConfig   `mapstructure:"monitoring"`
	Development   DevelopmentConfig  `mapstructure:"development"`
	Retail        RetailConfig       `mapstructure:"retail"`
	Integrations  IntegrationConfig  `mapstructure:"integrations"`
	Reports       ReportConfig       `mapstructure:"reports"`
}

// ServerConfig configurações do servidor HTTP
type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	Host         string        `mapstructure:"host"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// DatabaseConfig configurações do banco de dados
type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// SecurityConfig configurações de segurança
type SecurityConfig struct {
	JWTSecret        string        `mapstructure:"jwt_secret"`
	JWTExpiration    time.Duration `mapstructure:"jwt_expiration"`
	BcryptCost       int           `mapstructure:"bcrypt_cost"`
	SessionTimeout   time.Duration `mapstructure:"session_timeout"`
	MaxLoginAttempts int           `mapstructure:"max_login_attempts"`
	LockoutDuration  time.Duration `mapstructure:"lockout_duration"`
}

// LoggingConfig configurações de log
type LoggingConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// ModuleConfig configuração de um módulo
type ModuleConfig struct {
	Name    string `mapstructure:"name"`
	Enabled bool   `mapstructure:"enabled"`
	Version string `mapstructure:"version"`
}

// CacheConfig configurações de cache
type CacheConfig struct {
	Driver string        `mapstructure:"driver"`
	Redis  RedisConfig   `mapstructure:"redis"`
	TTL    time.Duration `mapstructure:"ttl"`
}

// RedisConfig configurações do Redis
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// EmailConfig configurações de email
type EmailConfig struct {
	Driver   string         `mapstructure:"driver"`
	SMTP     SMTPConfig     `mapstructure:"smtp"`
	SendGrid SendGridConfig `mapstructure:"sendgrid"`
}

// SMTPConfig configurações SMTP
type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}

// SendGridConfig configurações SendGrid
type SendGridConfig struct {
	APIKey string `mapstructure:"api_key"`
	From   string `mapstructure:"from"`
}

// UploadConfig configurações de upload
type UploadConfig struct {
	Driver       string      `mapstructure:"driver"`
	MaxSize      int         `mapstructure:"max_size"`
	AllowedTypes []string    `mapstructure:"allowed_types"`
	Local        LocalConfig `mapstructure:"local"`
	S3           S3Config    `mapstructure:"s3"`
}

// LocalConfig configurações de armazenamento local
type LocalConfig struct {
	Path string `mapstructure:"path"`
}

// S3Config configurações do S3
type S3Config struct {
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

// NotificationConfig configurações de notificação
type NotificationConfig struct {
	WebSocket WebSocketConfig `mapstructure:"websocket"`
	Push      PushConfig      `mapstructure:"push"`
}

// WebSocketConfig configurações do WebSocket
type WebSocketConfig struct {
	Enabled bool `mapstructure:"enabled"`
	Port    int  `mapstructure:"port"`
}

// PushConfig configurações de push notification
type PushConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	FCMKey  string `mapstructure:"fcm_key"`
}

// BackupConfig configurações de backup
type BackupConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Schedule    string `mapstructure:"schedule"`
	Retention   int    `mapstructure:"retention"`
	Path        string `mapstructure:"path"`
	Compression bool   `mapstructure:"compression"`
}

// MonitoringConfig configurações de monitoramento
type MonitoringConfig struct {
	Metrics     MetricsConfig     `mapstructure:"metrics"`
	HealthCheck HealthCheckConfig `mapstructure:"health_check"`
}

// MetricsConfig configurações de métricas
type MetricsConfig struct {
	Enabled bool `mapstructure:"enabled"`
	Port    int  `mapstructure:"port"`
}

// HealthCheckConfig configurações de health check
type HealthCheckConfig struct {
	Enabled  bool          `mapstructure:"enabled"`
	Interval time.Duration `mapstructure:"interval"`
}

// DevelopmentConfig configurações de desenvolvimento
type DevelopmentConfig struct {
	Debug     bool       `mapstructure:"debug"`
	HotReload bool       `mapstructure:"hot_reload"`
	CORS      CORSConfig `mapstructure:"cors"`
}

// CORSConfig configurações de CORS
type CORSConfig struct {
	Enabled bool     `mapstructure:"enabled"`
	Origins []string `mapstructure:"origins"`
	Methods []string `mapstructure:"methods"`
	Headers []string `mapstructure:"headers"`
}

// RetailConfig configurações específicas do varejo
type RetailConfig struct {
	POS       POSConfig       `mapstructure:"pos"`
	Inventory InventoryConfig `mapstructure:"inventory"`
	Pricing   PricingConfig   `mapstructure:"pricing"`
}

// POSConfig configurações do PDV
type POSConfig struct {
	Timeout         time.Duration `mapstructure:"timeout"`
	ReceiptTemplate string        `mapstructure:"receipt_template"`
}

// InventoryConfig configurações de inventário
type InventoryConfig struct {
	LowStockThreshold int  `mapstructure:"low_stock_threshold"`
	AutoReorder       bool `mapstructure:"auto_reorder"`
}

// PricingConfig configurações de preços
type PricingConfig struct {
	MarkupPercentage int `mapstructure:"markup_percentage"`
	DiscountLimit    int `mapstructure:"discount_limit"`
}

// IntegrationConfig configurações de integração
type IntegrationConfig struct {
	PaymentGateways []PaymentGatewayConfig `mapstructure:"payment_gateways"`
	Shipping        []ShippingConfig       `mapstructure:"shipping"`
}

// PaymentGatewayConfig configuração de gateway de pagamento
type PaymentGatewayConfig struct {
	Name     string `mapstructure:"name"`
	Enabled  bool   `mapstructure:"enabled"`
	APIKey   string `mapstructure:"api_key"`
	ClientID string `mapstructure:"client_id"`
	Secret   string `mapstructure:"secret"`
}

// ShippingConfig configuração de frete
type ShippingConfig struct {
	Name          string `mapstructure:"name"`
	Enabled       bool   `mapstructure:"enabled"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	APIKey        string `mapstructure:"api_key"`
	AccountNumber string `mapstructure:"account_number"`
}

// ReportConfig configurações de relatórios
type ReportConfig struct {
	DefaultFormat string           `mapstructure:"default_format"`
	StoragePath   string           `mapstructure:"storage_path"`
	Retention     int              `mapstructure:"retention"`
	Scheduling    SchedulingConfig `mapstructure:"scheduling"`
}

// SchedulingConfig configurações de agendamento
type SchedulingConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Timezone string `mapstructure:"timezone"`
}

// Load carrega a configuração do arquivo especificado
func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Configurar valores padrão
	setDefaults()

	// Ler arquivo de configuração
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
	}

	// Deserializar para struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erro ao deserializar configuração: %w", err)
	}

	// Validar configuração
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("erro na validação da configuração: %w", err)
	}

	return &config, nil
}

// setDefaults define valores padrão para a configuração
func setDefaults() {
	// Server
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.read_timeout", "15s")
	viper.SetDefault("server.write_timeout", "15s")
	viper.SetDefault("server.idle_timeout", "60s")

	// Database
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 5)
	viper.SetDefault("database.conn_max_lifetime", "5m")

	// Security
	viper.SetDefault("security.jwt_expiration", "24h")
	viper.SetDefault("security.bcrypt_cost", 12)
	viper.SetDefault("security.session_timeout", "8h")
	viper.SetDefault("security.max_login_attempts", 5)
	viper.SetDefault("security.lockout_duration", "30m")

	// Logging
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "json")
	viper.SetDefault("logging.output", "stdout")
	viper.SetDefault("logging.max_size", 100)
	viper.SetDefault("logging.max_age", 30)
	viper.SetDefault("logging.max_backups", 10)

	// Cache
	viper.SetDefault("cache.driver", "memory")
	viper.SetDefault("cache.ttl", "1h")

	// Development
	viper.SetDefault("development.debug", false)
	viper.SetDefault("development.hot_reload", false)
	viper.SetDefault("development.cors.enabled", true)

	// Retail
	viper.SetDefault("retail.pos.timeout", "30s")
	viper.SetDefault("retail.pos.receipt_template", "default")
	viper.SetDefault("retail.inventory.low_stock_threshold", 10)
	viper.SetDefault("retail.inventory.auto_reorder", false)
	viper.SetDefault("retail.pricing.markup_percentage", 30)
	viper.SetDefault("retail.pricing.discount_limit", 20)

	// Reports
	viper.SetDefault("reports.default_format", "pdf")
	viper.SetDefault("reports.retention", 90)
	viper.SetDefault("reports.scheduling.enabled", true)
	viper.SetDefault("reports.scheduling.timezone", "America/Sao_Paulo")
}

// validateConfig valida a configuração carregada
func validateConfig(config *Config) error {
	// Validar configurações obrigatórias
	if config.Database.Name == "" {
		return fmt.Errorf("nome do banco de dados é obrigatório")
	}

	if config.Database.User == "" {
		return fmt.Errorf("usuário do banco de dados é obrigatório")
	}

	if config.Security.JWTSecret == "" {
		return fmt.Errorf("chave JWT é obrigatória")
	}

	if config.Server.Port <= 0 {
		return fmt.Errorf("porta do servidor deve ser maior que 0")
	}

	return nil
}

// GetDSN retorna a string de conexão do banco de dados
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode)
}

// IsModuleEnabled verifica se um módulo está habilitado
func (c *Config) IsModuleEnabled(moduleName string) bool {
	for _, module := range c.Modules {
		if module.Name == moduleName {
			return module.Enabled
		}
	}
	return false
}
