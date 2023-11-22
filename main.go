package main





type apiConfigData struct{


	OpenWeatherApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct{

	Name string `json:"name"`

	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}

func loadApiConfig(filename string) (apiConfigData){


	
}