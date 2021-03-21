package main

import (
        "context"
        "flag"
        "fmt"
        "log"
        "time"
//        "bytes"
        "github.com/go-ble/ble"
        "github.com/go-ble/ble/examples/lib/dev"
        "github.com/pkg/errors"
		"encoding/json"
)

type Sensor struct {
	Sname       string `json:"Sname"`
	Slocation   string `json:"Slocation"`
	Stype       string `json:"Stype"`
	Smac        string `json:"Smac"`
	valSdatalen int    `json:"valSdatalen"`
	Sdata       Sdata
}
type Sdata struct {
	Mfdstructure     map[string]int  `json:"Sdatalables"`
	MSBLSBSdatabytes []int           `json:"MSBLSBSdatabytes"`
	ValSdatabytes    map[string]int  `json:"valSdatabytes"`
}
type SOutput struct {
	Soname       string `json:"Soname"`
	Solocation   string `json:"Solocation"`
	Sotype       string `json:"Sotype"`
	Sotimedate time.Time  `json:"Sotimedate"`
	Sodatamac    string `json:"Somac"`
	Sodatamid  string  `json:"Sodatamid"`
	SodataPm1  int   `json:"SodataPm1"`
	SodataPm25 int   `json:"SodataPm25"`
	SodataPm10 int   `json:"SodataPm10"`
	Sodatatemp float32 `json:"Sodatatemp"`
	Sodatahumd float32 `json:"Sodatahumd"`
	Sodatavoc  int  `json:"Sodatavoc"`
	Sodataco2  int   `json:"Sodataco02"`
}
var ad[] string
var (
        device = flag.String("device", "default", "implementation of ble")
        du     = flag.Duration("du", 5*time.Second, "scanning duration")
        dup    = flag.Bool("dup", true, "allow duplicate reported")
)

func main() {
        flag.Parse()

        d, err := dev.NewDevice(*device)
        if err != nil {
                log.Fatalf("can't new device : %s", err)
        }
        ble.SetDefaultDevice(d)

        // Scan for specified durantion, or until interrupted by user.
        fmt.Printf("Scanning for %s...\n", *du)
        ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), *du))
        chkErr(ble.Scan(ctx, *dup, advHandler, nil))
}

func advHandler(a ble.Advertisement) {
//      if a.Connectable() {
        //      fmt.Printf("[%s] C %3d:", a.Addr(), a.RSSI())
//      } else {
        //      fmt.Printf("[%s] N %3d:", a.Addr(), a.RSSI())
        //      }                   
        comma := ""
        if a.LocalName() == "aW1" {
                fmt.Printf(" Name: %s", a.LocalName())
                comma = ","
        if len(a.ManufacturerData()) > 0 {
                fmt.Printf("%s MD: %d", comma, a.ManufacturerData())
                ManufacturerData := a.ManufacturerData()
				vData := []*Sensor{}
				Datasval := new(Sensor)
					Datasval.Sname =   "aW1"
					Datasval.Slocation= "Garage"
					Datasval.Stype =   "Full"
					Datasval.Smac = "DF:5A:AD:BD:53:33"
					Datasval.Sdata.Mfdstructure = map[string]int{"MSBmid": 0, "LSBmid": 1, "MSBPm01": 2, "LSBPm01": 3, "MSBPm25": 4, "LSBPm25": 5, "MSBPm10": 6, "LSBPm10": 7, "MSBTemp": 8, "LSBTemp": 9, "MSBHund": 10, "LSBHund": 11, "MSBVoc": 12, "LSBVoc": 13, "MSBCo2": 14, "LSBCo2": 15}
					//Mfdstructure: {"MSBmid", "LSBmid", "MSBPm01", "LSBPm01", "MSBPm25", "LSBPm25", "MSBPm10", "LSBPm10", "MSBTemp", "LSBTemp", "MSBHund", "LSBHund", "MSBVoc", "LSBVoc", "MSBCo2", "LSBCo2"},
					//MSBLSBSdatabytes: {0, 1, 2, 3,  4,  5,  6,  7,  8, 9, 10, 11, 12, 13, 14, 15},
					
				
				
				//byte pmdata[23] = { 0xC3,0xD5,hPM01,lPM01,hPM2_5,lPM2_5,hPM10,lPM10,tH,tL,hH,hL,vH,vL,cH,cL}
			
				//var ManufacturerData = []byte{195, 213, 0, 20, 0, 33, 0, 42, 193, 8, 86, 18, 144, 3, 199, 3, 0, 0, 0, 0, 0, 0, 0}
				
				var sd = Datasval.Sdata.Mfdstructure 
				Datasval.Sdata.ValSdatabytes = make(map[string]int) 
				for key, element := range sd {
					fmt.Println("Key:", key, "=>", "Element:", element , ManufacturerData[element])
					Datasval.Sdata.ValSdatabytes[key] = (int(ManufacturerData[element]))
					}
					fmt.Println("datavals her we go",Datasval.Sdata.ValSdatabytes)
					SvData = append(SvData, Datasval)
					fmt.Println("datavals her we go",SvData)
					
			for i := range(SvData) {
					SvData := SvData[i]
					fmt.Println("data ,%tv\n:", SvData)
				}
			for key, value := range SvData { 
			fmt.Println("Key:", key, "Value:", value)
			}
			
			res2B, _ := json.Marshal(SvData)
			fmt.Printf("json %v\n",res2B)
			fmt.Println("her the json recod",string(res2B))
			fmt.Printf("datapacket %v\n",Datasval.Sdata.ValSdatabytes["LSBCo2"])
			fmt.Printf("datapacket %v\n",Datasval)
			
			
			///defining sensor output json file
			SoData := []*SOutput{}
			SoDatasval := new(SOutput)
				 SoDatasval.Soname = Datasval.Sname
				 SoDatasval.Solocation = Datasval.Slocation
				 SoDatasval.Sotype = Datasval.Stype
				 SoDatasval.Sotimedate = time.Now()
				 SoDatasval.Sodatamac = Datasval.Smac
				 SoDatasval.Sodatamid  = string(((Datasval.Sdata.ValSdatabytes["MSBmid"]) | (Datasval.Sdata.ValSdatabytes["LSBmid"])))
				 SoDatasval.SodataPm1  = ((Datasval.Sdata.ValSdatabytes["MSBPm01"]) + (Datasval.Sdata.ValSdatabytes["LSBPm01"]))
				 SoDatasval.SodataPm25 = ((Datasval.Sdata.ValSdatabytes["MSBPm25"]) + (Datasval.Sdata.ValSdatabytes["LSBPm25"]))
				 SoDatasval.SodataPm10 = ((Datasval.Sdata.ValSdatabytes["MSBPm10"]) + (Datasval.Sdata.ValSdatabytes["LSBPm10"]))
				 SoDatasval.Sodatatemp = (float32(((Datasval.Sdata.ValSdatabytes["MSBTemp"]) | (Datasval.Sdata.ValSdatabytes["LSBTemp"] *256)))/100)
				 SoDatasval.Sodatahumd = (float32(((Datasval.Sdata.ValSdatabytes["MSBHund"]) | (Datasval.Sdata.ValSdatabytes["LSBHund"] *256)))/100)
				 SoDatasval.Sodatavoc  = ((Datasval.Sdata.ValSdatabytes["MSBVoc"]) | (Datasval.Sdata.ValSdatabytes["LSBVoc"] *256))
				 SoDatasval.Sodataco2  = ((Datasval.Sdata.ValSdatabytes["MSBCo2"]) | (Datasval.Sdata.ValSdatabytes["LSBCo2"] *256))
				fmt.Printf("Sensor outout %v\n",SoDatasval)
				SoData = append(SoData, SoDatasval)
				sores2B, _ := json.Marshal(SoData)
					fmt.Printf("json %v\n",sores2B)
					fmt.Println("her the json recod output",string(sores2B))
			
					}              
                }
         }                                
        


func chkErr(err error) {
        switch errors.Cause(err) {
        case nil:
        case context.DeadlineExceeded:
                fmt.Printf("done\n")
        case context.Canceled:
                fmt.Printf("canceled\n")
        default:
                log.Fatalf(err.Error())
        }     		
}
