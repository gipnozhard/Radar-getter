package usecase

import (
	"Radar-getter/internal/delivery"
	"context"
	"sync"
	_ "time"

	_ "Radar-getter/internal/domain"
	"Radar-getter/internal/repository"
)

type Collector struct {
	radarClient  *delivery.RadarClient
	postgresRepo *repository.PostgresRepository
	mongoRepo    *repository.MongoRepository
	radarID      string
}

func NewCollector(
	radarClient *delivery.RadarClient,
	postgresRepo *repository.PostgresRepository,
	mongoRepo *repository.MongoRepository,
	radarID string,
) *Collector {
	return &Collector{
		radarClient:  radarClient,
		postgresRepo: postgresRepo,
		mongoRepo:    mongoRepo,
		radarID:      radarID,
	}
}

func (c *Collector) CollectData(ctx context.Context) error {
	var wg sync.WaitGroup
	var realTimeErr, statsErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		realTimeData, err := c.radarClient.GetRealTimeData(ctx)
		if err != nil {
			realTimeErr = err
			return
		}

		if err := c.postgresRepo.SaveRealTimeData(ctx, realTimeData); err != nil {
			realTimeErr = err
			return
		}

		if err := c.mongoRepo.SaveRealTimeData(ctx, realTimeData); err != nil {
			realTimeErr = err
		}
	}()

	go func() {
		defer wg.Done()
		statsData, err := c.radarClient.GetStatsData(ctx)
		if err != nil {
			statsErr = err
			return
		}

		if err := c.postgresRepo.SaveStatsData(ctx, statsData); err != nil {
			statsErr = err
			return
		}

		if err := c.mongoRepo.SaveStatsData(ctx, statsData); err != nil {
			statsErr = err
		}
	}()

	wg.Wait()

	if realTimeErr != nil {
		return realTimeErr
	}
	if statsErr != nil {
		return statsErr
	}

	return nil
}
