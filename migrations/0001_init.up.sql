-- Создание таблицы для реальных данных
CREATE TABLE realtime_data (
                               id SERIAL PRIMARY KEY,
                               equipment_id UUID NOT NULL,
                               lane_number INT NOT NULL,
                               lane_direction INT NOT NULL,
                               last_time BIGINT NOT NULL,
                               last_time_relative FLOAT NOT NULL,
                               last_time_registered FLOAT NOT NULL,
                               occupancy INT NOT NULL,
                               timestamp TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Создание таблицы для статистики
CREATE TABLE stats_data (
                            id SERIAL PRIMARY KEY,
                            equipment_id UUID NOT NULL,
                            lane_number INT NOT NULL,
                            lane_direction INT NOT NULL,
                            period_start TIMESTAMP WITH TIME ZONE NOT NULL,
                            period_end TIMESTAMP WITH TIME ZONE NOT NULL,

    -- Статистика по мотоциклам
                            motorbike_avg_speed FLOAT NOT NULL,
                            motorbike_sum_intensity INT NOT NULL,
                            motorbike_defined_sum_intensity INT NOT NULL,

    -- Статистика по автомобилям
                            car_avg_speed FLOAT NOT NULL,
                            car_sum_intensity INT NOT NULL,
                            car_defined_sum_intensity INT NOT NULL,

    -- Статистика по грузовикам
                            truck_avg_speed FLOAT NOT NULL,
                            truck_sum_intensity INT NOT NULL,
                            truck_defined_sum_intensity INT NOT NULL,

    -- Статистика по автобусам
                            bus_avg_speed FLOAT NOT NULL,
                            bus_sum_intensity INT NOT NULL,
                            bus_defined_sum_intensity INT NOT NULL,

    -- Параметры трафика
                            avg_speed FLOAT NOT NULL,
                            sum_intensity INT NOT NULL,
                            defined_sum_intensity INT NOT NULL,
                            avg_headway FLOAT NOT NULL,

                            timestamp TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Индексы для ускорения запросов
CREATE INDEX idx_realtime_data_equipment_id ON realtime_data (equipment_id);
CREATE INDEX idx_realtime_data_timestamp ON realtime_data (timestamp);
CREATE INDEX idx_stats_data_equipment_id ON stats_data (equipment_id);
CREATE INDEX idx_stats_data_timestamp ON stats_data (timestamp);