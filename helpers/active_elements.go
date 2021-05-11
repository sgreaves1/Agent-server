package helpers

import "github.com/sgreaves1/Agent-server/structs"

func GetActiveInfo() []*structs.Element {
	return []*structs.Element{
		{
			Id: "cae85d23-8c2d-4bea-999a-216ccae0996f",
			Name: "Agent 1",
			Type: "Agent",
			Location: "",
		},
		{
			Id: "fd47f095-9e42-4d55-99b6-63183411f804",
			Name: "CCTV 1",
			Type: "CCTV",
			Location: "",
		},
		{
			Id: "24b5c69e-329f-475a-a007-1c64576af373",
			Name: "Car 1",
			Type: "Car",
			Location: "",
		},
		{
			Id: "9c6ec2d9-44e3-4b7a-bae6-51539e47bea8",
			Name: "Bike 1",
			Type: "Bike",
			Location: "",
		},
	}
}

func GetAssetById(Id string) *structs.Element {
	activeAssets := GetActiveInfo()

	if activeAssets != nil {

		for i := range activeAssets {
			if activeAssets[i].Id == Id {
				return activeAssets[i]
			}
		}
	}
	return nil
}