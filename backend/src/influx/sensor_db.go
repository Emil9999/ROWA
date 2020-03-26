package influx

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

func InitInflux() *influxdb.Client {
	c := &http.Client{}
	token := "Ap_gylHmGd77dD86tHwH1PDlGpeXoz0HqotUWO43-cLNKMAA_K_QtY2-cqvCf6CSviagvPotqiSqalrQ_Ie3EQ=="

	influx, err := influxdb.New("https://eu-central-1-1.aws.cloud2.influxdata.com", token, influxdb.WithHTTPClient(c))
	log.Println(err)
	if err != nil {
		panic(err) // error handling here; normally we wouldn't use fmt but it works for the example
	}

	return influx

}
func InfluxWrite(s []float32, datetime time.Time, influx *influxdb.Client) {
	//now := time.Now()
	data := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{
				"temp":           s[0],
				"lightIntensity": s[1],
				"humidity":       s[2],
				"waterLevel":     s[3],
				"waterTemp":      s[4],
				"waterpH":        s[5],
			},
			"sensor-data",
			map[string]string{"farm-id": "1"},
			datetime),
	}
	// The actual write..., this method can be called concurrently.
	if _, err := influx.Write(context.Background(), "sensordata", "170189425a059be1", data...); err != nil {
		log.Fatal(err) // as above use your own error handling here.
	}
}

func InfluxClose(c *influxdb.Client) {
	c.Close() // closes the client.  After this the client is useless.
}
