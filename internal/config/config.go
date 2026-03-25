package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config はアプリケーション全体の設定を保持する構造体
// `env:"タグ"` を書くことで、環境変数名と紐付けます
type Config struct {
	AppEnv string `env:"APP_ENV" envDefault:"develop"`
	DBName string `env:"DBNAME,required"` // 必須項目の指定も可能
	DBHost string `env:"HOST" envDefault:"localhost"`
	DBUser string `env:"USER"`
	DBPass string `env:"PASSWORD"`
	Port   int    `env:"PORT" envDefault:"8080"` // 自動で int に変換
}

// NewConfig は設定を読み込み、Config構造体を返します（静的ファクトリメソッド）
func NewConfig() (*Config, error) {
	// 1. APP_ENV に応じた .env ファイルを読み込む
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "develop"
	}

	envFile := fmt.Sprintf(".env.%s", appEnv)

	// ファイルが存在する場合のみ読み込む（本番環境等でファイルがないケースを許容）
	_, err := os.Stat(envFile)
	if err == nil {
		if err := godotenv.Load(envFile); err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", envFile, err)
		}
	}

	// 2. 構造体にマッピング（パース）
	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}
