package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"), Endpoint: aws.String("arig5vtggzkh-ats.iot.eu-central-1.amazonaws.com")},
	)
	if err != nil {
		log.Fatal(err)
	}
	//iotSvc := iot.New(sess)
	iotDataSvc := iotdataplane.New(sess)
	//iotSvc.ListCertificates()
	input := &iotdataplane.PublishInput{
		Payload: []byte(`{
			"state": {
			 "reported":{
				"farm_id": "1",
				"humidity": 50.1,
				"temp": 10.5,
				"lightIntensity": 150.5,
				"waterLevel": 40.5,
				"waterTemp": 25.5,
				"waterpH": 7.1,
				"timestamp": 1586704334
			  }
			}
		  }`),
		Topic: aws.String("$aws/things/farm_1/shadow/update"),
		Qos:   aws.Int64(0),
	}
	//_ = input
	//_ = svc
	resp, err := iotDataSvc.Publish(input)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
