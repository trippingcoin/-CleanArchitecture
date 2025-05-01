package config

type Config struct {
	GRPCPort string
	NATSURL  string
	Postgres string
}

func Load() Config {
	return Config{
		GRPCPort: ":50053",
		NATSURL:  "nats://localhost:4222",
		Postgres: "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable",
	}
}
