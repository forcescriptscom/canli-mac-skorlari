package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Match struct {
	HomeTeam string `json:"home_team"`
	AwayTeam string `json:"away_team"`
	Score    struct {
		HomeTeam struct {
			Current int `json:"current"`
		} `json:"homeTeam"`
		AwayTeam struct {
			Current int `json:"current"`
		} `json:"awayTeam"`
		Minute int `json:"minute"`
	} `json:"score"`
	Date int64 `json:"date"`
}

type ApiResponse struct {
	Data []Match `json:"data"`
}

func main() {
	apiUrl := "https://apiv.forcescripts.com/DEMO"

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer resp.Body.Close()

	var result ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("JSON çözme hatası:", err)
		return
	}

	for _, match := range result.Data {
		fmt.Printf("Takımlar: %s vs %s\n", match.HomeTeam, match.AwayTeam)
		fmt.Printf("Skor: %d - %d\n", match.Score.HomeTeam.Current, match.Score.AwayTeam.Current)
		fmt.Printf("Dakika: %d\n", match.Score.Minute)
		fmt.Printf("Başlama Saati: %s\n", time.Unix(match.Date, 0).Format("2006-01-02 15:04:05"))
		fmt.Println()
	}
}
