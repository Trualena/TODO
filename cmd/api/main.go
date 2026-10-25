package main

import (
	"context"
	"log"
	"os"
	"time"
	"todo/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	//загрузка env
	if error := godotenv.Load(); error != nil { //чтение env и загрузка переменной / ошбка
		log.Println("env не найдн, переменные из окружения")
	}
	dns := os.Getenv("DATABASE_URL") //чтение переменной
	if dns == "" {
		log.Fatalf("DATABASE_URL не найден")
	}
	//ctx записывает резт работы функции
	var ctx context.Context
	var cancel context.CancelFunc //сигнал отмены
	//в cacel функия для отмены
	//подключение к бд
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var pool *pgxpool.Pool
	var error error
	pool, error = database.Connect(ctx, dns)
	if error != nil {
		log.Fatalf("не удалось подключиться к бизе %v", error)
	}
	defer pool.Close()
	router := gin.Default() //логирует запрос и возвр 500 при ошибке
	router.GET("/health", func(cont *gin.Context) {
		error := pool.Ping(cont.Request.Context())
		if error == nil { //если нет ошибки
			cont.JSON(200, gin.H{"status_database": "OK"})
			return
		}
		cont.JSON(503, gin.H{"status_database": "error"})
	})
	log.Println("сервер запущен на localhost^8080")
	if error := router.Run(":8080"); error != nil {
		log.Fatalf("сервер не запустился %v", error)
	}

}
