package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"time"
)

func StartSimulationSession(e echo.Context) error {

	//L1; 172.25.0.101; 100; L1|SetCodeStatus|26N00001CGUMZYCB00001SKAKI64P2GJ8NYWOQ3UZVOCRQM3M4QIH5TP2I0IA69P4RAO55PNITZSMJE7CGXWFJ7LBXO77HNTGPJVQ4HZCASFYBQC8FPNRILFUBM962GGR81L5TSPGGVKMGCHF00001|(TID)1|(SI)NA|(MRG)2631600001|@; ; ; 0; ; 172.25.0.100; UDP
	//Device;Address;Delay;Cmd;LinkedDevice;LinkedAddress;LinkedDelay;LinkedCmd;CLMAddress;Protocol
	//Device[0] Address[1] Delay[2] Cmd[3] LinkedDevice[4] LinkedAddress[5] LinkedDelay[6] LinkedCmd[7] CLMAddress[8] Protocol[9]

	csvReader := csv.NewReader(bufio.NewReader(e.Request().Body))
	csvReader.Comma = ','
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1 //un numero negativo sta a significare che la lunghezza delle righe del csv è variabile

	rows, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Cannot read the csv file: " + err.Error())
	}

	for i, row := range rows {
		if len(row) < 5 {
			continue // skippa i records troppo corti
		}
		if i == 0 {
			continue
		}
		device := row[0]
		addr := row[1]

		delay, err := strconv.Atoi(row[2])
		if err != nil {
			_, _ = color.Set(color.Red)
			fmt.Println("Error parsing: " + err.Error())
			_, _ = color.Reset()
		}

		cmd := row[3]
		linkedDevice := row[4]
		linkedAddress := row[5]

		linkedDelay, err := strconv.Atoi(row[6])
		if err != nil {
			if linkedDevice != "" {
				_, _ = color.Set(color.Red)
				fmt.Println("Error parsing: " + err.Error())
				_, _ = color.Reset()
			}
			linkedDelay = 0
		}

		linkedCmd := row[7]

		if device == "L1" || device == "L8" || device == "L9" || device == "L10" || device == "MTS" {
			go func(c string, a string) {
				WriteUDPMessage(c, a)
			}(cmd, addr)
		}

		if strings.ToUpper(device) == "MODBUS" {
			go func(c string, a string) {
				SendCommandModbus(c, a)
			}(cmd, addr)
		}

		if device == "L6" {
			go func(c string, a string) {
				WriteTCPMessage(c, a)
			}(cmd, addr)

			if linkedDevice == "L7" {
				go func(c string, a string) {
					time.Sleep(time.Duration(linkedDelay) * time.Millisecond)
					WriteUDPMessage(c, a)
				}(linkedCmd, linkedAddress)
			}
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	return e.JSON(200, "ok")
}

//func ReadCSVFromHttpRequest(w http.ResponseWriter, r *http.Request) {
//	// parse POST body as csv
//	reader := csv.NewReader(r.Body)
//	r.Header.Get("Content-Type")
//	var results [][]string
//	for {
//		record, err := reader.Read()
//		if err == io.EOF {
//			break
//		}
//
//		// add record to result set
//		fmt.Println(record)
//		results = append(results, record)
//	}
//}
