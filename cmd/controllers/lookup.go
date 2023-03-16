package controllers

import (
	"net"
	"strings"

	"github.com/OrbKonSec/mailintel/cmd/utils"
)

func Lookup(email string) {
	defer utils.ProgressBar.Add(10)
	splt := strings.Split(email, "@")
	iprecords, _ := net.LookupIP(splt[1])
	for _, ip := range iprecords {
		row := []string{"IP", ip.String()}
		utils.Lookup_temp_result = append(utils.Lookup_temp_result, row)
	}
	nameserver, _ := net.LookupNS(splt[1])
	for _, ns := range nameserver {
		row := []string{"NS", ns.Host}
		utils.Lookup_temp_result = append(utils.Lookup_temp_result, row)
	}
	mxrecords, _ := net.LookupMX(splt[1])
	for _, mx := range mxrecords {
		row := []string{"MX", mx.Host}
		utils.Lookup_temp_result = append(utils.Lookup_temp_result, row)
	}
	txtrecords, _ := net.LookupTXT(splt[1])
	for _, txt := range txtrecords {
		row := []string{"TXT", txt}
		utils.Lookup_temp_result = append(utils.Lookup_temp_result, row)
	}
}