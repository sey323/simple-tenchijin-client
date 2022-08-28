package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const endpoint = "https://hackathon-api.compass.tenchijin.co.jp/v1"
const endpoint2 = "https://hackathon-api2.compass.tenchijin.co.jp/v1"

type TenchijinClient struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Token string `json:"token"`
}

func NewTenchijinClient(username string, password string, token string) *TenchijinClient {
	return &TenchijinClient{username, password, token}
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (tc *TenchijinClient) Login() {
	endpoint := fmt.Sprintf("%s/access-token", endpoint)

	// ユーザ名とパスワードからトークンを発行。
	jsonString, err := json.Marshal(tc)

	tknReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	tknReq.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	tcjResp, err := client.Do(tknReq)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	defer tknReq.Body.Close()

	if tcjResp.StatusCode == 200 {
		decoder := json.NewDecoder(tcjResp.Body)
		var token TokenResponse
		err := decoder.Decode(&token)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		tc.Token = token.Token
	}
}

type WeatherRainFallRequest struct {
	Lat                float32 `json:"lat"`
	Lng                float32 `json:"lng"`
	Since              string  `json:"since"`
	Until              string  `json:"until"`
	TemporalResolution string  `json:"temporal_resolution"`
	Timezone           string  `json:"timezone"`
}

func NewWeatherRainFallRequest(lat float32, lng float32, since string, until string, temporal_resolution string, timezone string) *WeatherRainFallRequest {
	return &WeatherRainFallRequest{lat, lng, since, until, temporal_resolution, timezone}
}

func (tc *TenchijinClient) WeatherRainFall(
	lat float32,
	lng float32,
	since string,
	until string,
	temporal_resolution string,
	timezone string,
) string {
	endpoint := fmt.Sprintf("%s/weather/rainfall", endpoint)

	nwrr := NewWeatherRainFallRequest(
		lat,
		lng,
		since,
		until,
		temporal_resolution,
		timezone,
	)
	// ユーザ名とパスワードからトークンを発行。
	jsonString, err := json.Marshal(nwrr)

	tknReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	tknReq.Header.Add("Content-Type", "application/json")
	tknReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.Token))

	client := &http.Client{}
	tcjResp, err := client.Do(tknReq)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	defer tknReq.Body.Close()

	if tcjResp.StatusCode == 200 {
		respBody, _ := io.ReadAll(tcjResp.Body)
		return string(respBody)
	}
	return "error"
}

type WeatherLstRequest struct {
	Lat                float32 `json:"lat"`
	Lng                float32 `json:"lng"`
	Since              string  `json:"since"`
	Until              string  `json:"until"`
	TemporalResolution string  `json:"temporal_resolution"`
}

func NeWeatherLstRequest(lat float32, lng float32, since string, until string, temporal_resolution string) *WeatherLstRequest {
	return &WeatherLstRequest{lat, lng, since, until, temporal_resolution}
}

func (tc *TenchijinClient) WeatherLst(
	lat float32,
	lng float32,
	since string,
	until string,
	temporal_resolution string,
) string {
	endpoint := fmt.Sprintf("%s/weather/lst", endpoint2)

	nwrr := NeWeatherLstRequest(
		lat,
		lng,
		since,
		until,
		temporal_resolution,
	)
	// ユーザ名とパスワードからトークンを発行。
	jsonString, err := json.Marshal(nwrr)

	tknReq, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	tknReq.Header.Add("Content-Type", "application/json")
	tknReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.Token))

	client := &http.Client{}
	tcjResp, err := client.Do(tknReq)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	defer tknReq.Body.Close()

	if tcjResp.StatusCode == 200 {
		respBody, _ := io.ReadAll(tcjResp.Body)
		return string(respBody)
	}
	return "error"
}
