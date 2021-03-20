package main

/*
const api_root_url = "https://api.us-east-1.mbedcloud.com/v2/"
const callback_path = "notification/callback/"

type CallbackURl struct {
	Url string	`json:"url"`
}
*/

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	// You can generate a Token from the "Tokens Tab" in the UI
	const token = "XWAa4nUekIJmxTD5DsyDwJk7FUn4kBfwckTFugN8wRKpbVoiKpRGo_n4LARJYEqzl32voZL0xDfNg6WSqSWkqg=="
	const bucket = "davidbourke94's Bucket"
	const org = "davidbourke94@gmail.com"

	client := influxdb2.NewClient("https://us-central1-1.gcp.cloud2.influxdata.com", token)
	// always close client at the end
	defer client.Close()

	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)

	for i := 0; i < 60; i++ {
		// write line protocol
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 45.0))
		// Flush writes
		writeAPI.Flush()

		time.Sleep(100 * time.Millisecond)
	}

}
