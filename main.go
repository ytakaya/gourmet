package main

import (
	"fmt"
	"yamazooooe/gourmet-share/config"
	"yamazooooe/gourmet-share/hotpepper"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	apiClient := hotpepper.New(config.Config.ApiKey)
	res, _ := apiClient.GetGourmet()
	fmt.Println(res[0])
}
