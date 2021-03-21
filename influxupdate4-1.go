package main

import (
    "context"
    "fmt"
    "time"
    "flag"
	"log"
	"runtime"
    
	"github.com/influxdata/influxdb-client-go"
    //"github.com/influxdata/influxdb-client-go"
  //"github.com/influxdata/influxdb-client-go"
  nats "github.com/nats-io/nats.go"
)
type sensorData struct {
	SensorID    string     `json:"sensorId"`
	LocationTag string     `json:"locationTag"`
	TemP    float32 `json:"temP"`
	HumD    float32 `json:"humD"`
	VoC     int32 `json:"voC"`
	TvoCppB int32 `json:"tvoCppB"`
	Eco2PpM int32 `json:"eco2ppM"`
	pM25    int32  `json:"pM2.5"`
	pM5     int32 `json:"pM5"`
    pM10    int32 `json:"pM10"`
    aQi     int32 `json:"aQi"`
}
func main() {
    var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var authUser = flag.String("u", "nats", "The nats server authentication user for clients")
	var authPassword = flag.String("p", "", "The nats server authentication password for clients")
	var command = flag.String("c", "", "Whether to produce or consume a message")
	log.SetFlags(0)
	flag.Parse()
    args := flag.Args()
    nc, err := nats.Connect(*urls, nats.UserInfo(*authUser, *authPassword))
	if err != nil {
		log.Fatal(err)
    }
    wg := sync.WaitGroup{}
    wg.Add(1)
    userName := "user"
    password := "user1234"
     // Create a new client using an InfluxDB server base URL and an authentication token
    // For authentication token supply a string in the form: "username:password" as a token. Set empty value for an unauthenticated server
    client := influxdb2.NewClient("http://localhost:8086", fmt.Sprintf("%s:%s",userName, password))
    // Get the blocking write client
    // Supply a string in the form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
    // Org name is not used
    if _, err := ec.Subscribe("sData", func(s *sensorData) {
        log.Printf("Stock: %s - Price: %v", s.SensorID, s.LocationTag, s.TemP, s.HumD, s.VoC, s.TvoCppB, s.Eco2PpM, s.pM25, s.pM5, s.pM10, s.aQi)
        wg.Done()
    }); err != nil {
        log.Fatal(err)
    }
    writeAPI := client.WriteAPIBlocking("sensor_data", "sensor_data")
    // create point using full params constructor
    p = influxdb2.NewPointWithMeasurement("stat").
        AddTag("SensorID", s.SensorID).
        AddTag("LocationTag", s.Location).
        AddField("TemP", s.TemP).
        AddField("HumD", s.HumD).
        AddField("VoC", s.VoC).
        AddField("TvoCppB", S.TvoCppB).
        AddField("Eco2PpM", s.Eco2PpM).
        AddField("pM25", s.pM25).
        AddField("pM5", s,pM5).
        AddField("pM10", s,pM10).
        AddField("aQi", s,aQi).
        SetTime(time.Now())
    // Write data
    err := writeAPI.WritePoint(context.Background(), p)
    if err != nil {
        fmt.Printf("Write error: %s\n", err.Error())
    }
    client.Close()
    wg.Wait()
}
