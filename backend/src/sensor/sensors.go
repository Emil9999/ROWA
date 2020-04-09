package sensor

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MarcelCode/ROWA/src/influx"
	"github.com/kuzemkon/aws-iot-device-sdk-go/device"
	"github.com/tarm/serial"
)

/*Serial Port Configs
/dev/ttyACM0 - Rasp
/dev/cu.usbmodem1434301 Macbook
COM5 windows
*/

func initAws() /**iotdataplane.IoTDataPlane*/ {
	/*sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"), Endpoint: aws.String("arig5vtggzkh.iot.eu-central-1.amazonaws.com")},
	)
	if err != nil {
		log.Fatal(err)
	}
	//iotSvc := iot.New(sess)
	iotDataSvc := iotdataplane.New(sess)
	//var listInput iot.ListThingsInput
	//listInput.MaxResults = *int64(10)*/
	thing, err := device.NewThing(
		device.KeyPair{
			PrivateKeyPath:    "../../keys/9a7699c575-private.pem.key",
			CertificatePath:   "../../keys/9a7699c575-certificate.pem.crt",
			CACertificatePath: "../../keys/rootCA1.pem.crt",
		},
		"arig5vtggzkh.iot.eu-central-1.amazonaws.com", // AWS IoT endpoint
		device.ThingName("farm_1"),
	)

	if err != nil {
		panic(err)
	}

	s, err := thing.GetThingShadow()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	payload := []byte(`{
		'state': {
		 'desired':{
				'humidity':10,
				'temp':10            
		   }
		}
	  }`)
	err = thing.UpdateThingShadow(payload)
	if err != nil {
		panic(err)
	}
	//things, err := iotSvc.ListThings()
	//return iotDataSvc

}

func setupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "/dev/cu.usbmodem1434301", Baud: 9600}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ActivateModuleLight(moduleNumber int) {
	// Open Serial Connection
	s, err := setupSerialConnection()
	defer s.Close()

	// Create String from Module Number and send to connection
	moduleString := strconv.Itoa(moduleNumber)
	_, err = s.Write([]byte(moduleString))
	if err != nil {
		log.Fatal(err)
	}

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}

func DeactivateModuleLight() {
	// Open Serial Connection
	s, err := setupSerialConnection()
	defer s.Close()

	// Create String from Module Number and send to connection
	_, err = s.Write([]byte("99"))
	if err != nil {
		log.Fatal(err)
	}

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}

func LightSwitch(state bool) {

	s, err := setupSerialConnection()
	defer s.Close()

	//send turn off or on to arduino
	if state {
		fmt.Println("d")
		_, err = s.Write([]byte("80"))

		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 1")
		statement.Exec()
		database.Close()
	} else {
		fmt.Println("c")
		_, err = s.Write([]byte("81"))

		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 1")
		statement.Exec()
		database.Close()
	}

	if err != nil {
		log.Fatal(err)
	}

	// Give Connection time to send Data
	time.Sleep(2 * time.Second)
}
func ReadFakeSensorData() {
	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()
	client := influx.InitInflux()

	for {
		var s []float32
		datetime := time.Now()
		temp := rand.Float32()*2 + 20
		lightIntensity := rand.Float32()*100 + 400
		humidity := rand.Float32()*10 + 30
		waterLevel := rand.Float32()*4 + 19
		waterTemp := rand.Float32()*5 + 18
		waterpH := rand.Float32()*2 + 6
		s = append(s, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
		influx.InfluxWrite(s, datetime, client)
		initAws()
		//AWS
		/*svc := initAws()

		input := &iotdataplane.PublishInput{
			Payload: []byte(`{
				'state': {
				 'desired':{
						'humidity':10,
						'temp':10
				   }
				}
			  }`),
			Topic: aws.String("/update"),
			Qos:   aws.Int64(0),
		}
		_ = input
		_ = svc
		//resp, err := svc.Publish(input)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)*/

		datetimeStr := datetime.UTC().Format(time.RFC3339)
		fmt.Println(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
		statement.Exec(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)

		//Cloud part

		//str := "mem,host=host1 used_percent=23.43234543 1556896326"
		//url := "https://eu-central-1-1.aws.cloud2.influxdata.com/api/v2/write?org=170189425a059be1&bucket=sensordata&precision=s"
		//req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
		time.Sleep(60 * time.Second)
	}
	//TODO on system shutdown influx.InfluxClose(client)

}
func ReadSensorData() {
	var serialString string

	s, _ := setupSerialConnection()
	defer s.Close()

	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		serialIncome := string(buf[:n])
		serialString += serialIncome

		if strings.HasSuffix(serialString, "\n") {
			datetime := time.Now().UTC().Format(time.RFC3339)
			raw_string := strings.TrimSuffix(serialString, "\r\n")
			data_array := strings.Split(raw_string, ",")
			if len(data_array) >= 6 {
				temp, err1 := strconv.ParseFloat(data_array[0], 32)
				lightIntensity, err2 := strconv.ParseFloat(data_array[1], 32)
				humidity, err3 := strconv.ParseFloat(data_array[2], 32)
				waterLevel, err4 := strconv.ParseFloat(data_array[3], 32)
				waterTemp, err5 := strconv.ParseFloat(data_array[4], 32)
				waterpH, err6 := strconv.ParseFloat(data_array[5], 32)

				if err1 == nil && err2 == nil && err4 == nil && err3 == nil && err5 == nil && err6 == nil {
					fmt.Println(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
					statement.Exec(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
				}
			}
			serialString = ""
		}

	}
}
