package helper

import (
	"encoding/json"
	"fmt"
	"italking.tomotoes.com/m/v1/config"
	"net/http"
)

type RoomData struct {
	ChannelExist bool `json:"channel_exist"`
	Mode         int
	Total        int
	Users        []int
}

type RoomResponse struct {
	Success bool
	Data    RoomData
}

func IsRoomExist(roomName string) bool {
	url := fmt.Sprintf("https://api.agora.io/dev/v1/channel/user/%s/%s", config.GetAgoraAppId(), roomName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+config.GetAgoraAuthToken())
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false
	}

	var response RoomResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if !response.Success ||
		!response.Data.ChannelExist ||
		response.Data.Total == 0 ||
		len(response.Data.Users) == 0 {
		return false
	}

	return true
}
