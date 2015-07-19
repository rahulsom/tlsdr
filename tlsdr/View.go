package main

import (
	"bytes"
	"container/list"
	htmltemplate "html/template"
	"strconv"
	"text/template"
	//"fmt"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
)

func Visualize(data list.List, format string) []byte {
	//TODO replace with real data
	groups := groupConnectionsDataModel(data)
	//var result []byte
	output := new(bytes.Buffer)
	switch format {
	case "txt":
		{
			tmpl, err := template.ParseFiles("template/txt/HandshakeProtocolDetails.txt")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(output, groups)
			if err != nil {
				panic(err)
			}
		}
	case "html":
		{
			tmpl, err := htmltemplate.ParseFiles("template/html/HandshakeProtocolDetails.html")
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(output, groups)
			if err != nil {
				panic(err)
			}
		}
	case "json":
		{
			b, err := json.Marshal(groups)
			if err != nil {
				panic(err)
			}
			output.Write(b)
		}
	}
	return output.Bytes()
}

func groupConnectionsDataModel(connections list.List) [][]Connection {

	groupMap := make(map[string][]Connection)
	for e := connections.Front(); e != nil; e = e.Next() {
		conn := e.Value.(Connection)

		//do grouping
		var key string
		if conn.Success {
			key = conn.SrcHost + "-" + conn.DestHost + strconv.Itoa(conn.Events.Len()) + "-success"
		} else {
			key = conn.SrcHost + "-" + conn.DestHost + strconv.Itoa(conn.Events.Len()) + "-false-" + conn.FailedReason
		}
		existingGroup, found := groupMap[key]
		if !found {
			existingGroup = make([]Connection, 0)
		}
		groupMap[key] = append(existingGroup, conn)
	}

	groups := make([][]Connection, len(groupMap))
	var i int = 0
	for _, value := range groupMap {
		groups[i] = value
		i++
	}
	return groups
}

func ColorizeOutput(bytes []byte) []byte {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	//blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	var s string = string(bytes[:len(bytes)])
	s = strings.Replace(s, "Failure", red("Failure"), -1)
	s = strings.Replace(s, "-Recommendations", yellow("-Recommendations"), -1)
	s = strings.Replace(s, "Success", green("Success"), -1)
	s = strings.Replace(s, "All success", green("All success"), -1)
	s = strings.Replace(s, "src:", cyan("src:"), -1)
	s = strings.Replace(s, "dest:", cyan("dest:"), -1)

	//a little hack to get rid of empty failure reason
	s = strings.Replace(s, "<>", " ", -1)

	return []byte(s)
}
