package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"weather/models"
)

const weatherUrl = "https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", getWeather)

	err := http.ListenAndServe(":1234", mux)
	if err != nil {
		fmt.Printf("[error] failed to create server: %s\n", err)
	}
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	hasAPIKey := r.URL.Query().Has("appid")
	if !hasAPIKey {
		http.Error(w, "error: missing `appid` param", http.StatusBadRequest)
		return
	}

	hasLatitude := r.URL.Query().Has("lat")
	if !hasLatitude {
		http.Error(w, "error: missing `lat` param", http.StatusBadRequest)
		return
	}

	hasLongitude := r.URL.Query().Has("lon")
	if !hasLongitude {
		http.Error(w, "error: missing `lon` param", http.StatusBadRequest)
		return
	}

	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	apiKey := r.URL.Query().Get("appid")

	requestURL := fmt.Sprintf(weatherUrl, lat, lon, apiKey)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("[error] failed creating request: %s\n", err)
		http.Error(w, "something blew up internally ", http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("[error] failed making http request: %s\n", err)
		http.Error(w, "something blew up internally ", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[error] failed reading response body: %s\n", err)
		http.Error(w, "something blew up internally ", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[error] failed call to OpenWeather: %s\n", string(resBody))
		http.Error(w, "error: failed to complete request to OpenWeather", resp.StatusCode)
		return
	}

	var weather models.Weather
	err = json.Unmarshal(resBody, &weather)
	if err != nil {
		fmt.Printf("[error] failed to unmarshal resp: %s\n", err)
		http.Error(w, "something blew up internally ", http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Has("units") {
		weather.Units = strings.ToLower(r.URL.Query().Get("units"))
	}

	io.WriteString(w, weather.Explanation())
	fmt.Printf("> successfully completed /weather call\n")
}
