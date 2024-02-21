
# Weather Service

Creates an http server that exposes an endpoint to get weather data from OpenWeather and gives description of the current weather an conditions.

Notes:
- Must provide apiKey as param or set in main.go.
- Left api call just like OpenWeather's api. Didn't include all the optional params except for `unit`





## API Reference

#### Get weather explanation

```http
  GET /weather
```

| Parameter | Description                |
| :-------- |  :------------------------- |
| `lat`   |  **Required**. Latitude coordinate of city  |
| `lon`   |  **Required**. Longitude coordinate of city |
| `appid`   |  **Required**. Your OpenWeather API key. Overwrites anything set in main.go  |
| `unit`   |  Optional. Default is Kelvin. `imperial` = °F, `metric` = °C |





## Usage/Examples

Build or run main then execute curl

```
curl 'http://localhost:1234/weather?lat=33.4255&lon=-111.9400&appid={appid}&units=imperial'
```

Response:
```
In Tempe Junction, the temperature today is a high of 68 and a low of 56. It is currently cold at 62°F. Current weather conditions include: broken clouds.
```
