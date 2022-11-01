package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const target = "capture.pcap"
const flag = "HTB{"

func main() {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("--------------Solve HTBOO Trick or Breach---------------")
	fmt.Println("--------------by: Ir0nM4n-------------------------------")
	fmt.Println("--------------@: https://github.com/brennenwright-------")
	fmt.Println("\n-----Step One: parse the capture.pcap file")

	// Open file instead of device
	handle, err := pcap.OpenOffline(target)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	fmt.Println("--Filter PCAP for DNS outbound")
	// Set filter
	var filter string = "udp and dst port 53"
	fmt.Println("    Filter: ", filter)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("--Create a flag.xlsx file")
	file, err := os.Create("flag.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("--Decode the DNS qname as hex and write it to the xlsx")
	for packet := range packetSource.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer == nil {
			continue
		}
		dns := dnsLayer.(*layers.DNS)
		qst := dns.Questions[0]
		name := qst.Name[0 : len(dns.Questions[0].Name)-16]

		decoded := make([]byte, hex.DecodedLen(len(name)))
		n, err := hex.Decode(decoded, name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(decoded[:n])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("\n-----Step Two: Read the flagfile")
	fmt.Println("--Open the completed file for reading as an excel file")

	f, err := excelize.OpenFile("flag.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	flagfile := ""
	//loop the sheets
	for sheet := range f.GetSheetList() {
		// Get all the rows in the Sheet
		rows, _ := f.GetRows(f.GetSheetList()[sheet])
		for _, row := range rows {
			for _, colCell := range row {
				flagfile = flagfile + colCell
			}
		}
	}
	fmt.Println("--Parse for text starting with: ", flag)
	//parse for the flag
	fmt.Println("\n-----Flag: ")
	i := strings.Index(string(flagfile), flag)
	fmt.Println(string(flagfile)[i : strings.Index(string(flagfile)[i:], "}")+i+1])

}
