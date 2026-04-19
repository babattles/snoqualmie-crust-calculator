package snowobs

type Response struct {
	Units     Units      `json:"UNITS"`
	Variables []Variable `json:"VARIABLES"`
	Stations  []Station  `json:"STATION"`
}

type Units struct {
	AirTemp            string `json:"air_temp"`
	BatteryVoltage     string `json:"battery_voltage"`
	IntermittentSnow   string `json:"intermittent_snow"`
	PrecipAccumOneHour string `json:"precip_accum_one_hour"`
	RelativeHumidity   string `json:"relative_humidity"`
	SnowDepth          string `json:"snow_depth"`
	SnowDepth24h       string `json:"snow_depth_24h"`
	SolarRadiation     string `json:"solar_radiation"`
	WindDirection      string `json:"wind_direction"`
	WindGust           string `json:"wind_gust"`
	WindSpeed          string `json:"wind_speed"`
	WindSpeedMin       string `json:"wind_speed_min"`
}

type Variable struct {
	Variable string `json:"variable"`
	LongName string `json:"long_name"`
}

type Station struct {
	ID           string        `json:"id"`
	Stid         string        `json:"stid"`
	Name         string        `json:"name"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`
	Elevation    float64       `json:"elevation"`
	Timezone     string        `json:"timezone"`
	Source       string        `json:"source"`
	Meta         StationMeta   `json:"meta"`
	StationNote  []StationNote `json:"station_note"`
	Observations Observations  `json:"observations"`
}

type StationMeta struct {
	State                 string `json:"state"`
	DataloggerNumID       int    `json:"datalogger_num_id"`
	DataloggerCharID      string `json:"datalogger_char_id"`
	WeatherStationPartner string `json:"weather_station_partner"`
}

type StationNote struct {
	ID          string  `json:"id"`
	ClientID    int     `json:"client_id"`
	Stid        string  `json:"stid"`
	DateCreated string  `json:"date_created"`
	DateUpdated string  `json:"date_updated"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Status      string  `json:"status"`
	Note        string  `json:"note"`
	History     *string `json:"history"`
}

type Observations struct {
	DateTime           []string  `json:"date_time"`
	AirTemp            []float64 `json:"air_temp,omitempty"`
	BatteryVoltage     []float64 `json:"battery_voltage,omitempty"`
	IntermittentSnow   []float64 `json:"intermittent_snow,omitempty"`
	PrecipAccumOneHour []float64 `json:"precip_accum_one_hour,omitempty"`
	RelativeHumidity   []float64 `json:"relative_humidity,omitempty"`
	SnowDepth          []float64 `json:"snow_depth,omitempty"`
	SnowDepth24h       []float64 `json:"snow_depth_24h,omitempty"`
	SolarRadiation     []float64 `json:"solar_radiation,omitempty"`
	WindDirection      []float64 `json:"wind_direction,omitempty"`
	WindGust           []float64 `json:"wind_gust,omitempty"`
	WindSpeed          []float64 `json:"wind_speed,omitempty"`
	WindSpeedMin       []float64 `json:"wind_speed_min,omitempty"`
}
