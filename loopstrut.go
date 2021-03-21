package main

import (
	"fmt"
	//"reflect"
	"time"
)

type Sensor struct {
	Sname       string `json:"Sname"`
	Slocation   string `json:"Slocation"`
	Stype       string `json:"Stype"`
	valSdatalen int    `json:"valSdatalen"`
	Smac        string `json:"Smac"`
	Sdata       Sdata
}
type Sdata struct {
	Mfdstructure     map[string]int  `json:"Sdatalables"`
	MSBLSBSdatabytes []int           `json:"MSBLSBSdatabytes"`
	valSdatabytes    map[string]int `json:"valSdatabytes"`
}
type SOutput struct {
	Sotimedate   string  `json:"Sotimedate"`
	Soname       string `json:"Soname"`
	Solocation   string `json:"Solocation"`
	Sotype       string `json:"Sotype"`
	Sodatamid    string  `json:"Sodatamid"`
	Sodatamac    string `json:"Somac"`
	SodataPm1    int32   `json:"SodataPm1"`
	SodataPm25   int32   `json:"SodataPm25"`
	SodataPm10   int32   `json:"SodataPm10"`
	Sodatatemp   float32 `json:"Sodatatemp"`
	Sodatahumd   float32 `json:"Sodatahumd"`
	Sodatavoc    int32   `json:"Sodatavoc"`
	Sodataco2    int32   `json:"Sodataco02"`
}

//var  ManufacturerData = []byte {195, 213, 0, 20, 0, 33, 0, 42, 193, 8, 86, 18, 144, 3, 199, 3, 0, 0, 0, 0, 0, 0, 0}

func main() {
        SvData := []*Sensor{}
	    Datasval := new(Sensor)
		Datasval.Sname =   "aW1"
		Datasval.Slocation= "Garage"
		Datasval.Stype =   "Full"
		Datasval.Smac = "DF:5A:AD:BD:53:33"
		Datasval.Sdata.Mfdstructure = map[string]int{"MSBmid": 0, "LSBmid": 1, "MSBPm01": 2, "LSBPm01": 3, "MSBPm25": 4, "LSBPm25": 5, "MSBPm10": 6, "LSBPm10": 7, "MSBTemp": 8, "LSBTemp": 9, "MSBHund": 10, "LSBHund": 11, "MSBVoc": 12, "LSBVoc": 13, "MSBCo2": 14, "LSBCo2": 15}
		//Mfdstructure: {"MSBmid", "LSBmid", "MSBPm01", "LSBPm01", "MSBPm25", "LSBPm25", "MSBPm10", "LSBPm10", "MSBTemp", "LSBTemp", "MSBHund", "LSBHund", "MSBVoc", "LSBVoc", "MSBCo2", "LSBCo2"},
		//MSBLSBSdatabytes: {0, 1, 2, 3,  4,  5,  6,  7,  8, 9, 10, 11, 12, 13, 14, 15},
		
	
	
	//byte pmdata[23] = { 0xC3,0xD5,hPM01,lPM01,hPM2_5,lPM2_5,hPM10,lPM10,tH,tL,hH,hL,vH,vL,cH,cL}

	var ManufacturerData = []byte{195, 213, 0, 20, 0, 33, 0, 42, 193, 8, 86, 18, 144, 3, 199, 3, 0, 0, 0, 0, 0, 0, 0
	
	var sd = Datasval.Sdata.Mfdstructure 
	//var sdval = s.Sdata.valSdatabytes
	Datasval.Sdata.valSdatabytes = make(map[string]int)
	for key, element := range sd {
        fmt.Println("Key:", key, "=>", "Element:", element , ManufacturerData[element])
        //s.Sdata.valSdatabytes := append(s.Sdata.valSdatabytes, key, (int(ManufacturerData[element])))
        Datasval.Sdata.valSdatabytes[key] = (int(ManufacturerData[element]))
        }
        SvData = append(SvData, Datasval)
        //for key, element := range sd {
        //fmt.Println("Key:", key, "=>", "Element:", element , ManufacturerData[element])
        //sdval = add(sdval, key, (int(ManufacturerData[element])))
        //s.Sdata.valSdatabytes[key] = (int(ManufacturerData[element]))
        //}
for i := range(SvData) {
        SvData := SvData[i]
        fmt.Println("Location:", SvData)
    }
	SvData := []*SOutput {}
	SoDatasval := new(SOutput)


	//fmt.Printf("interesting %v\n",SvData)
	//tagsList := make([Sensor], 0)
	
	//for i, p := range s1.Sdata.Sdatalables {
	//fmt.Printf("%v\n", s1.Sdata.Sdatalables[i])
	//if s1.Sdata.Sdatalables[i] == "MSBmid" {
	//vdatamid := (MDdata[i] | MDdata[i+1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBPm1" {
	//	vdataPm1 := (MDdata[i] + MDdata[i+1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBPm25" {
	//vdataPm25 := (MDdata[i] + MDdata[i+1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBPm10" {
	//vdataPm10 := (MDdata[i] + MDdata[i+1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBTemp" {
	//MDBdatatemp := ((MDdata[i]))
	//vdatatemo := (MDBdatatemp | MDdata[i + 1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBHund" {
	//	vdatahumd := (MDdata[i] | MDdata[i + 1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBVoc" {
	//	vdatavoc := (((MDdata[i]) ) | MDdata[i+1])
	//}
	//if s1.Sdata.Sdatalables[i] == "MSBCo2" {
	//	vdataco2 := (((MDdata[i])) | MDdata[i+1])
	//}
	i++
	//fmt.Printf("%v\n", i)
	//}

	//fmt.Printf("%v\n", vdatamid)

	//p++
}
