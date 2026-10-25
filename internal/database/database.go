package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// соединение с бд
// constring string содержит текст с адресом бд, логином, паролем
// фия вернет (*pgxpool.Pool, error) набор соединение с бд и ошибку
func Connect(ctx context.Context, constring string) (*pgxpool.Pool, error) {
	//New- имя функции
	pool, error := pgxpool.New(ctx, constring)
	if error != nil {
		return nil, fmt.Errorf("не удалось создать пул соединений %w", error)
	}
	if error := pool.Ping(ctx); error != nil {
		pool.Close()
		return nil, fmt.Errorf("бд не отвечает на пинг %w", error)
	}
	return pool, nil
}
