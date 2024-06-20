package app

import (
	"encoding/json"
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

	fmt.Println("calling StatsHandler")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Example table name
	table := "fs_stats"

	for {
		select {
		case <-ctx.Done():
			fmt.Println("client disconnected or context canceled")
			return
		default:
			// Get the most recent data from the specified table
			data, err := database.GetRecentStatsData(a.db, table)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to get most recent data: %v", err), http.StatusInternalServerError)
				return
			}

			// convert StatsJSON from []byte to a string
			statsJSONString := string(data.StatsJSON)

			// prepare a map for the SSE message
			sseMessage := map[string]interface{}{
				"id":         data.ID,
				"stats_json": statsJSONString,
				"created_at": data.CreatedAt,
			}

			// encode the SSE message to JSON
			sseJSON, err := json.Marshal(sseMessage)
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}

			// Write SSE message to the response
			fmt.Fprintf(w, "data: %s\n\n", sseJSON)

			// Flush the response to ensure data is sent immediately
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}

			// Sleep or wait for an event to occur before sending the next message
			time.Sleep(2 * time.Second)
		}
	}
}
