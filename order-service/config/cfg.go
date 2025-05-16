package config

import (
	"time"

	"github.com/Neroframe/ecommerce-platform/order-service/pkg/mongo"
	"github.com/caarlos0/env/v10"
)

type (
	Config struct {
		Version string `env:"VERSION" envDefault:"1.0.0"`

		Mongo   mongo.Config
		Server  Server
		Nats    Nats
	}

	Server struct {
		GRPCServer GRPCServer
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
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
