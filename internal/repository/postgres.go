package repository

import (
	"Radar-getter/internal/config"
	"context"
	"fmt"
	_ "time"

	"Radar-getter/internal/domain"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(cfg *config.PostgresConfig) (*PostgresRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) SaveRealTimeData(ctx context.Context, data *domain.RealTimeData) error {
	query := `
		INSERT INTO realtime_data (
			equipment_id, lane_number, lane_direction, last_time, 
			last_time_relative, last_time_registered, occupancy, timestamp
		) VALUES (
			:equipment_id, :lane_number, :lane_direction, :last_time, 
			:last_time_relative, :last_time_registered, :occupancy, :timestamp
		)`

	for _, lane := range data.LaneData {
		_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
			"equipment_id":         data.EquipmentID,
			"lane_number":          lane.LaneNumber,
			"lane_direction":       lane.LaneDirection,
			"last_time":            lane.LastTime,
			"last_time_relative":   lane.LastTimeRelative,
			"last_time_registered": lane.LastTimeRegistered,
			"occupancy":            lane.Occupancy,
			"timestamp":            data.Timestamp,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PostgresRepository) SaveStatsData(ctx context.Context, data *domain.StatsData) error {
	query := `
		INSERT INTO stats_data (
			equipment_id, lane_number, lane_direction, period_start, period_end,
			motorbike_avg_speed, motorbike_sum_intensity, motorbike_defined_sum_intensity,
			car_avg_speed, car_sum_intensity, car_defined_sum_intensity,
			truck_avg_speed, truck_sum_intensity, truck_defined_sum_intensity,
			bus_avg_speed, bus_sum_intensity, bus_defined_sum_intensity,
			avg_speed, sum_intensity, defined_sum_intensity, avg_headway, timestamp
		) VALUES (
			:equipment_id, :lane_number, :lane_direction, :period_start, :period_end,
			:motorbike_avg_speed, :motorbike_sum_intensity, :motorbike_defined_sum_intensity,
			:car_avg_speed, :car_sum_intensity, :car_defined_sum_intensity,
			:truck_avg_speed, :truck_sum_intensity, :truck_defined_sum_intensity,
			:bus_avg_speed, :bus_sum_intensity, :bus_defined_sum_intensity,
			:avg_speed, :sum_intensity, :defined_sum_intensity, :avg_headway, :timestamp
		)`

	for _, stat := range data.Stats {
		_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
			"equipment_id":                    data.EquipmentID,
			"lane_number":                     stat.LaneNumber,
			"lane_direction":                  stat.LaneDirection,
			"period_start":                    stat.PeriodStart,
			"period_end":                      stat.PeriodEnd,
			"motorbike_avg_speed":             stat.Statistics.Motorbike.EstimatedAvgSpeed,
			"motorbike_sum_intensity":         stat.Statistics.Motorbike.EstimatedSumIntensity,
			"motorbike_defined_sum_intensity": stat.Statistics.Motorbike.EstimatedDefinedSumIntensity,
			"car_avg_speed":                   stat.Statistics.Car.EstimatedAvgSpeed,
			"car_sum_intensity":               stat.Statistics.Car.EstimatedSumIntensity,
			"car_defined_sum_intensity":       stat.Statistics.Car.EstimatedDefinedSumIntensity,
			"truck_avg_speed":                 stat.Statistics.Truck.EstimatedAvgSpeed,
			"truck_sum_intensity":             stat.Statistics.Truck.EstimatedSumIntensity,
			"truck_defined_sum_intensity":     stat.Statistics.Truck.EstimatedDefinedSumIntensity,
			"bus_avg_speed":                   stat.Statistics.Bus.EstimatedAvgSpeed,
			"bus_sum_intensity":               stat.Statistics.Bus.EstimatedSumIntensity,
			"bus_defined_sum_intensity":       stat.Statistics.Bus.EstimatedDefinedSumIntensity,
			"avg_speed":                       stat.TrafficFlowParams.AvgSpeed,
			"sum_intensity":                   stat.TrafficFlowParams.SumIntensity,
			"defined_sum_intensity":           stat.TrafficFlowParams.DefinedSumIntensity,
			"avg_headway":                     stat.TrafficFlowParams.AvgHeadway,
			"timestamp":                       data.Timestamp,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
