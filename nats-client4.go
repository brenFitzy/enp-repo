package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"
	"context"
	//"strconv"
	
	"github.com/influxdata/influxdb-client-go"
	nats "github.com/nats-io/nats.go"
)
type sensordata struct {
	Sensorid    string   
	Locationtag string     
	TemP    float32 
	HumD    float32 
	VoC     int `json:"VoC"`
	TvoCppB int `json:"TvoCppB"`
	Eco2PpM int `json:"Eco2PpM"`
	PM25    int `json:"PM25"`
	PM5     int `json:"PM5"`
	PM10    int `json:"PM10"`
	AQi     int `json:"AQi"`
}
func usage_prod() {
	log.Fatalf("Usage: nats-pub [-s server (%s)] [-u user (%s)] [-p password (%s)] -c produce <subject> <msg> \n", nats.DefaultURL, "nats", "S3Cr3TP@5w0rD")
}

func usage_con() {
	log.Fatalf("Usage: nats-pub [-s server (%s)] [-u user (%s)] [-p password (%s)] -c consume <subject> \n", nats.DefaultURL, "nats", "S3Cr3TP@5w0rD")
}

func db_update(d *sensordata) {
	userName := "user"
    password := "user1234"
     // Create a new client using an InfluxDB server base URL and an authentication token
    // For authentication token supply a string in the form: "username:password" as a token. Set empty value for an unauthenticated server
    client := influxdb2.NewClient("http://localhost:8086", fmt.Sprintf("%s:%s",userName, password))
    // Get the blocking write client
    // Supply a string in the form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
    // Org name is not used
	writeAPI := client.WriteAPIBlocking("sensor_data", "sensor_data")
    // create point using full params constructor
    p := influxdb2.NewPointWithMeasurement("data").
        AddField("SensorID", d.Sensorid).
        AddField("LocationTag", d.Locationtag).
        AddField("TemP", d.TemP).
        AddField("HumD", d.HumD).
        AddField("VoC", d.VoC).
        AddField("TvoCppB", d.TvoCppB).
        AddField("Eco2PpM", d.Eco2PpM).
        AddField("PM25", d.PM25).
        AddField("PM5", d.PM5).
        AddField("PM10", d.PM10).
        AddField("AQi", d.AQi).
        SetTime(time.Now())
    // Write data
    err := writeAPI.WritePoint(context.Background(), p)
    if err != nil {
        fmt.Printf("Write error: %s\n", err.Error())
    }
    client.Close()
    log.Printf("Received message from function'%+v,v,v,v,v,v,v\n", string(d.Sensorid), string(d.Locationtag), d.TemP, d.HumD, d.VoC, d.TvoCppB, d.Eco2PpM, d.PM25, d.PM5, d.PM10, d.AQi)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var authUser = flag.String("u", "nats", "The nats server authentication user for clients")
	var authPassword = flag.String("p", "", "The nats server authentication password for clients")
	var command = flag.String("c", "", "Whether to produce or consume a message")
	log.SetFlags(0)
	flag.Parse()
	args := flag.Args()
	if *command == "" {
		log.Fatalf("Error: Indicate the command using '-command' flag")
	}
	if *command != "produce" && *command != "consume" {
		log.Fatalf("Error: Supported commands are: consume & produce")
	}
	nc, err := nats.Connect(*urls, nats.UserInfo(*authUser, *authPassword))
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to NATS server: " + *urls)
	if *command == "produce" { 
		if len(args) < 2 {
			usage_prod()
		}
		//subj, msg := args[0], []byte(args[1])
		subj := args[0]
		data := &sensordata{Sensorid: "1234", Locationtag: "Room", TemP: 22.1, HumD: 45.2, VoC: 15000, TvoCppB: 500, Eco2PpM: 1000, PM25: 12, PM5: 26, PM10: 27, AQi: 7}
		ec.Publish(subj, data)
		ec.Flush()
		if err := ec.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subj, data)
		} 
	}
	if *command == "consume" {
		if len(args) < 1 {
			usage_con()
		}
		subj := args[0]
		ec.Subscribe(subj, func(s *sensordata) {
			log.Printf("Received message %+v\n", s)
		log.Printf("Received message '%+v\n", string(s.Sensorid))
		log.Printf("Received message '%+v\n", s.AQi)
		cp := s
		log.Printf("Received message cp '%+v\n", cp.AQi)
		db_update(cp)
		})
		ec.Flush()
		if err := ec.LastError(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Listening on [%s]\n", subj)
		runtime.Goexit()
	}
}
