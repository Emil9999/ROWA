package influx

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type SensorData struct {
	FarmId         string  `json:"farm_id"`
	Temp           float32 `json:"temp"`
	LightIntensity float32 `json:"lightIntensity"`
	Humid          float32 `json:"humidity"`
	WaterLevel     float32 `json:"waterLevel"`
	WaterTemp      float32 `json:"waterTemp"`
	WaterpH        float32 `json:"waterpH"`
	Timestamp      int64   `json:"timestamp"`
}

/*func InitInflux() *influxdb.Client {
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
}*/

func awsInit() *iotdataplane.IoTDataPlane {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"), Endpoint: aws.String("arig5vtggzkh-ats.iot.eu-central-1.amazonaws.com")},
	)
	if err != nil {
		log.Fatal(err)
	}
	iotDataSvc := iotdataplane.New(sess)
	return iotDataSvc
}
func awsPublishInput(iotDataSvc *iotdataplane.IoTDataPlane, s []float32, timestamp int64) {
	data := &SensorData{
		FarmId:         "1",
		Temp:           s[0],
		LightIntensity: s[1],
		Humid:          s[2],
		WaterLevel:     s[3],
		WaterTemp:      s[4],
		WaterpH:        s[5],
		Timestamp:      timestamp,
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("could not marshal JSON: %s", err)
	}
	input := &iotdataplane.PublishInput{
		Payload: jsonStr,
		Topic:   aws.String("$aws/things/farm_1/shadow/update"),
		Qos:     aws.Int64(0),
	}
	//_ = input
	//_ = svc
	resp, err := iotDataSvc.Publish(input)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
