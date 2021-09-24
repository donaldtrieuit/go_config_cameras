package main

import (
	"encoding/json"
	"github.com/donaldtrieuit/go_config_cameras/hkvision/client"
	"github.com/donaldtrieuit/go_config_cameras/hkvision/types"
	"log"
)

/**
 * @author : Donald Trieu
 * @created : 9/25/21, Saturday
**/

func main() {
	newClient, _ := client.NewClient(types.ConstructClient{
		Username: "admin",
		Password: "AI_team123",
		Host: "192.168.1.152",
		Proto: "http",
	})
	channel, err := newClient.GetSingleStreamChannel("101")
	if err != nil {
		log.Print(err)
	}
	bytes, _ := json.Marshal(channel)
	log.Printf("Result : %s\n", string(bytes))
	//channel.Video.ConstantBitRate = 512
	//log.Printf("maxFrameRate: %v\n", channel.Video.ConstantBitRate)
	//err = newClient.UpdateSingleStreamChannel(strconv.Itoa(channel.ID), channel)
	//if err != nil {
	//	log.Print(err)
	//}
}