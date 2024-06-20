package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/enkdr/monitor/database"
)

type pageData struct {
	Title string
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling IndexHandler")
	pageData := pageData{
		Title: "M O N I T O R",
	}

	err := a.templates.ExecuteTemplate(w, "index.html", pageData)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (a *App) StatsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("calling data handler")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("client disconnected or context canceled")
			return
		default:

			data, err := database.GetRecentStatsData(a.db, "fs_stats")

			// // Marshal the Data object to JSON
			// jsonData, err := json.Marshal(data)

			if err != nil {
				fmt.Println("error selecting data:", err)
				return
			}

			// convert StatsJSON from []byte to string
			statsJSONString := string(data.StatsJSON)

			// Prepare a map for JSON encoding
			responseData := map[string]interface{}{
				"id":         data.ID,
				"stats_json": statsJSONString,
				"created_at": data.CreatedAt,
			}

			// if err := json.NewEncoder(w).Encode(responseData); err != nil {
			// 	http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			// }

			fmt.Fprintf(w, "data: %s\n\n", responseData)

			w.(http.Flusher).Flush()

			i++
			time.Sleep(2 * time.Second)
		}
	}
}
