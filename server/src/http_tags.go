package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func httpTags(c *gin.Context) {
	// Local variables
	w := c.Writer

	// Parse the player from the URL
	player := c.Param("player")
	if player == "" {
		http.Error(w, "Error: You must specify a player.", http.StatusNotFound)
		return
	}
	normalizedUsername := normalizeString(player)

	// Check if the player exists
	var user User
	if exists, v, err := models.Users.GetUserFromNormalizedUsername(
		normalizedUsername,
	); err != nil {
		logger.Error("Failed to check to see if player \""+player+"\" exists:", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	} else if exists {
		user = v
	} else {
		http.Error(w, "Error: That player does not exist in the database.", http.StatusNotFound)
		return
	}

	// Search through the database for tags matching this user ID
	var gamesMap map[int][]string
	if v, err := models.GameTags.SearchByUserID(user.ID); err != nil {
		logger.Error("Failed to search for games matching a user ID of "+
			strconv.Itoa(user.ID)+":", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	} else {
		gamesMap = v
	}

	// Get the keys of the map
	// https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
	gameIDs := make([]int, len(gamesMap))
	i := 0
	for k := range gamesMap {
		gameIDs[i] = k
		i++
	}

	// Get the games corresponding to these IDs
	var gameHistoryList []*GameHistory
	if v, err := models.Games.GetHistory(gameIDs); err != nil {
		logger.Error("Failed to get the games from the database:", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	} else {
		gameHistoryList = v
	}

	// Attach the tags to each GameHistory object
	// (they do not normally come from the database with the tags on them)
	for _, gameHistory := range gameHistoryList {
		if tags, ok := gamesMap[gameHistory.ID]; ok {
			sort.Strings(tags)
			gameHistory.Tags = strings.Join(tags, ", ")
		}
	}

	data := HistoryData{
		Title:   "Tagged Games",
		Name:    user.Username,
		History: gameHistoryList,
		Tags:    gamesMap,
	}
	httpServeTemplate(w, data, "profile", "history")
}
