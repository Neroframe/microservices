package config

import (
	"time"

	"github.com/Neroframe/ecommerce-platform/inventory-service/pkg/mongo"
	"github.com/caarlos0/env/v10"
)

type (
	Config struct {
		Mongo  mongo.Config
		Server Server
		Nats   Nats
		Redis  Redis
		Cache  Cache
	}

	Server struct {
		GRPCServer GRPCServer
		// HTTPServer
		// Metrics
	}

	GRPCServer struct {
		Port                  int           `env:"GRPC_PORT,notEmpty"`
		MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
		MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
		MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
	}

	Nats struct {
		Hosts        []string `env:"NATS_HOSTS,notEmpty" envSeparator:","`
		NKey         string   `env:"NATS_NKEY" envDefault:"SUACSSL3UAHUDXKFSNVUZRF5UHPMWZ6BFDTJ7M6USDXIEDNPPQYYYCU3VY"`
		IsTest       bool     `env:"NATS_IS_TEST,notEmpty" envDefault:"true"`
		NatsSubjects NatsSubjects
	}

	NatsSubjects struct {
		OrderCreatedSubject string `env:"NATS_ORDER_CREATED_SUBJECT,notEmpty"`
		OrderUpdatedSubject string `env:"NATS_ORDER_UPDATED_SUBJECT,notEmpty"`
		OrderDeletedSubject string `env:"NATS_ORDER_DELETED_SUBJECT,notEmpty"`
	}

	Redis struct {
		Host         string        `env:"REDIS_HOSTS,notEmpty" envSeparator:","`
		Password     string        `env:"REDIS_PASSWORD"`
		TLSEnable    bool          `env:"REDIS_TLS_ENABLE" envDefault:"true"`
		DialTimeout  time.Duration `env:"REDIS_DIAL_TIMEOUT" envDefault:"60s"`
		WriteTimeout time.Duration `env:"REDIS_WRITE_TIMEOUT" envDefault:"60s"`
		ReadTimeout  time.Duration `env:"REDIS_READ_TIMEOUT" envDefault:"30s"`
	}

	Cache struct {
		ClientTTL              time.Duration `env:"REDIS_CACHE_CLIENT_TTL" envDefault:"24h"`
		CMSVariableRefreshTime time.Duration `env:"CLIENT_REFRESH_TIME" envDefault:"1m"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
