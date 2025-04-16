package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"Radar-getter/internal/config"
	"Radar-getter/internal/delivery"
	"Radar-getter/internal/repository"
	"Radar-getter/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig("data.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация репозиториев
	postgresRepo, err := repository.NewPostgresRepository(&cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize Postgres repository: %v", err)
	}
	defer postgresRepo.Close()

	mongoRepo, err := repository.NewMongoRepository(&cfg.MongoDB)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB repository: %v", err)
	}
	defer mongoRepo.Close()

	// Создание коллекторов для каждого радара
	var collectors []*usecase.Collector
	for _, radar := range cfg.Radars {
		client := delivery.NewRadarClient(radar.BaseURL)
		collector := usecase.NewCollector(client, postgresRepo, mongoRepo, radar.ID)
		collectors = append(collectors, collector)
	}

	// Контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для обработки сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Таймер для сбора данных каждую минуту
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	log.Println("Service started")

	var wg sync.WaitGroup

	for {
		select {
		case <-ticker.C:
			for _, collector := range collectors {
				wg.Add(1)
				go func(c *usecase.Collector) {
					defer wg.Done()
					if err := c.CollectData(ctx); err != nil {
						log.Printf("Failed to collect data: %v", err)
					}
				}(collector)
			}
			wg.Wait()
			log.Println("Data collected at", time.Now().Format(time.RFC3339))
		case <-sigChan:
			log.Println("Shutting down...")
			return
		}
	}
}
