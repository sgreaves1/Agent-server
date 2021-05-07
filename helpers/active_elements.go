package helpers

import "github.com/sgreaves1/Agent-server/structs"

func GetActiveInfo() []*structs.Element {
	return []*structs.Element{
		{
			Name: "Agent 1",
			Type: "Agent",
			Location: "",
		},
		{
			Name: "CCTV 1",
			Type: "CCTV",
			Location: "",
		},
		{
			Name: "Car 1",
			Type: "Car",
			Location: "",
		},
		{
			Name: "Bike 1",
			Type: "Bike",
			Location: "",
		},
	}

}