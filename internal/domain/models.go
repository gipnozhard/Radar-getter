package domain

import "time"

type RealTimeData struct {
	EquipmentID string     `json:"equipment_id"`
	LaneData    []LaneData `json:"data"`
	Timestamp   time.Time  `json:"timestamp"`
}

type LaneData struct {
	LaneNumber         int     `json:"lane_number"`
	LaneDirection      int     `json:"lane_direction"`
	LastTime           int64   `json:"last_time"`
	LastTimeRelative   float64 `json:"last_time_relative"`
	LastTimeRegistered float64 `json:"last_time_registered"`
	Occupancy          int     `json:"occupancy"`
}

type StatsData struct {
	EquipmentID string      `json:"equipment_id"`
	Stats       []StatsItem `json:"data"`
	Timestamp   time.Time   `json:"timestamp"`
}

type StatsItem struct {
	LaneNumber        int           `json:"lane_number"`
	LaneDirection     int           `json:"lane_direction"`
	PeriodStart       time.Time     `json:"period_start"`
	PeriodEnd         time.Time     `json:"period_end"`
	Statistics        VehicleStats  `json:"statistics"`
	TrafficFlowParams TrafficParams `json:"traffic_flow_parameters"`
}

type VehicleStats struct {
	Motorbike VehicleStat `json:"motorbike"`
	Car       VehicleStat `json:"car"`
	Truck     VehicleStat `json:"truck"`
	Bus       VehicleStat `json:"bus"`
}

type VehicleStat struct {
	EstimatedAvgSpeed            float64 `json:"estimated_avg_speed"`
	EstimatedSumIntensity        int     `json:"estimated_sum_intensity"`
	EstimatedDefinedSumIntensity int     `json:"estimated_defined_sum_intensity"`
}

type TrafficParams struct {
	AvgSpeed            float64 `json:"avg_speed"`
	SumIntensity        int     `json:"sum_intensity"`
	DefinedSumIntensity int     `json:"defined_sum_intensity"`
	AvgHeadway          float64 `json:"avg_headway"`
}

type Radar struct {
	ID      string
	BaseURL string
}
