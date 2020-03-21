package db

import (
	"context"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

func InitInflux() *influxdb.Client {

	token := "Ap_gylHmGd77dD86tHwH1PDlGpeXoz0HqotUWO43-cLNKMAA_K_QtY2-cqvCf6CSviagvPotqiSqalrQ_Ie3EQ=="
	influx, err := influxdb.New("https://eu-central-1-1.aws.cloud2.influxdata.com", token, influxdb.WithHTTPClient(nil))
	if err != nil {
		panic(err) // error handling here; normally we wouldn't use fmt but it works for the example
	}

	return influx

}
func InfluxWrite(temp float32, lightIntensity float32, influx influxdb.Client) {
	data := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"temp": temp, "lightIntensity": lightIntensity},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Date(2018, 3, 4, 5, 6, 7, 8, time.UTC)),
	}
	// The actual write..., this method can be called concurrently.
	if _, err := influx.Write(context.Background(), "sensordata", "170189425a059be1", data...); err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}
}

func InfluxClose(c influxdb.Client) {
	c.Close() // closes the client.  After this the client is useless.
}
