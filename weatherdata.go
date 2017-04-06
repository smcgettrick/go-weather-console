package main

import "encoding/json"

// WeatherData is the JSON Response Object
type WeatherData struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Geolookup  int `json:"geolookup"`
			Conditions int `json:"conditions"`
		} `json:"features"`
	} `json:"response"`
	Location struct {
		Type                  string `json:"type"`
		Country               string `json:"country"`
		CountryIso3166        string `json:"country_iso3166"`
		CountryName           string `json:"country_name"`
		State                 string `json:"state"`
		City                  string `json:"city"`
		TzShort               string `json:"tz_short"`
		TzLong                string `json:"tz_long"`
		Lat                   string `json:"lat"`
		Lon                   string `json:"lon"`
		Zip                   string `json:"zip"`
		Magic                 string `json:"magic"`
		Wmo                   string `json:"wmo"`
		L                     string `json:"l"`
		Requesturl            string `json:"requesturl"`
		Wuiurl                string `json:"wuiurl"`
		NearbyWeatherStations struct {
			Airport struct {
				Station []struct {
					City    string `json:"city"`
					State   string `json:"state"`
					Country string `json:"country"`
					Icao    string `json:"icao"`
					Lat     string `json:"lat"`
					Lon     string `json:"lon"`
				} `json:"station"`
			} `json:"airport"`
			Pws struct {
				Station []struct {
					Neighborhood string  `json:"neighborhood"`
					City         string  `json:"city"`
					State        string  `json:"state"`
					Country      string  `json:"country"`
					ID           string  `json:"id"`
					Lat          float64 `json:"lat"`
					Lon          float64 `json:"lon"`
					DistanceKm   int     `json:"distance_km"`
					DistanceMi   int     `json:"distance_mi"`
				} `json:"station"`
			} `json:"pws"`
		} `json:"nearby_weather_stations"`
	} `json:"location"`
	CurrentObservation struct {
		Image struct {
			URL   string `json:"url"`
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"image"`
		DisplayLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			StateName      string `json:"state_name"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Zip            string `json:"zip"`
			Magic          string `json:"magic"`
			Wmo            string `json:"wmo"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"display_location"`
		ObservationLocation struct {
			Full           string `json:"full"`
			City           string `json:"city"`
			State          string `json:"state"`
			Country        string `json:"country"`
			CountryIso3166 string `json:"country_iso3166"`
			Latitude       string `json:"latitude"`
			Longitude      string `json:"longitude"`
			Elevation      string `json:"elevation"`
		} `json:"observation_location"`
		Estimated struct {
		} `json:"estimated"`
		StationID             string      `json:"station_id"`
		ObservationTime       string      `json:"observation_time"`
		ObservationTimeRfc822 string      `json:"observation_time_rfc822"`
		ObservationEpoch      string      `json:"observation_epoch"`
		LocalTimeRfc822       string      `json:"local_time_rfc822"`
		LocalEpoch            string      `json:"local_epoch"`
		LocalTzShort          string      `json:"local_tz_short"`
		LocalTzLong           string      `json:"local_tz_long"`
		LocalTzOffset         string      `json:"local_tz_offset"`
		Weather               string      `json:"weather"`
		TemperatureString     string      `json:"temperature_string"`
		TempF                 float64     `json:"temp_f"`
		TempC                 float64     `json:"temp_c"`
		RelativeHumidity      string      `json:"relative_humidity"`
		WindString            string      `json:"wind_string"`
		WindDir               string      `json:"wind_dir"`
		WindDegrees           int         `json:"wind_degrees"`
		WindMph               float64     `json:"wind_mph"`
		WindGustMph           json.Number `json:"wind_gust_mph, string"`
		WindKph               json.Number `json:"wind_kph, stirng"`
		WindGustKph           json.Number `json:"wind_gust_kph, string"`
		PressureMb            string      `json:"pressure_mb"`
		PressureIn            string      `json:"pressure_in"`
		PressureTrend         string      `json:"pressure_trend"`
		DewpointString        string      `json:"dewpoint_string"`
		DewpointF             int         `json:"dewpoint_f"`
		DewpointC             int         `json:"dewpoint_c"`
		HeatIndexString       string      `json:"heat_index_string"`
		HeatIndexF            string      `json:"heat_index_f"`
		HeatIndexC            string      `json:"heat_index_c"`
		WindchillString       string      `json:"windchill_string"`
		WindchillF            string      `json:"windchill_f"`
		WindchillC            string      `json:"windchill_c"`
		FeelslikeString       string      `json:"feelslike_string"`
		FeelslikeF            string      `json:"feelslike_f"`
		FeelslikeC            string      `json:"feelslike_c"`
		VisibilityMi          string      `json:"visibility_mi"`
		VisibilityKm          string      `json:"visibility_km"`
		Solarradiation        string      `json:"solarradiation"`
		UV                    string      `json:"UV"`
		Precip1HrString       string      `json:"precip_1hr_string"`
		Precip1HrIn           string      `json:"precip_1hr_in"`
		Precip1HrMetric       string      `json:"precip_1hr_metric"`
		PrecipTodayString     string      `json:"precip_today_string"`
		PrecipTodayIn         string      `json:"precip_today_in"`
		PrecipTodayMetric     string      `json:"precip_today_metric"`
		Icon                  string      `json:"icon"`
		IconURL               string      `json:"icon_url"`
		ForecastURL           string      `json:"forecast_url"`
		HistoryURL            string      `json:"history_url"`
		ObURL                 string      `json:"ob_url"`
		Nowcast               string      `json:"nowcast"`
	} `json:"current_observation"`
}
