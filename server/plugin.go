package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/v5/plugin"

	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// Link struct
type Link struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Link    string `json:"link"`
	Tags    string `json:"tags"`
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {

	// user configs:
	team := "8r4yr7x4nfdytndqn1sidruuxw"
	dbLocation := "/home/shadab/.local/lib/python2.7/site-packages/linkbot/sqlite.db"

	userID := r.Header.Get("Mattermost-User-Id")
	channels, err4 := p.API.GetChannelsForTeamForUser(team, userID, false)
	if err4 != nil {
		log.Fatal(err4, "failed to get channels for team: %s for user: %s", team, userID)
	}
	channelsListStr := "'"
	for _, channel := range channels {
		channelsListStr += channel.Name + "','"
	}
	channelsListStr = channelsListStr[:len(channelsListStr)-2]

	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id      int
		author  string
		message string
		link    string
		tag     string
	)
	rows, err2 := db.Query("select l.id, t.tag from links l join tags t on l.id = t.message_id where l.channel in (" + channelsListStr + ")")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer rows.Close()
	rowsLink, err21 := db.Query("select l.id, l.author, l.message, l.link from links l where l.channel in (" + channelsListStr + ")")
	if err21 != nil {
		log.Fatal(err21)
	}
	defer rowsLink.Close()
	tagmap := make(map[int]string)
	var ok bool
	for rows.Next() {
		err3 := rows.Scan(&id, &tag)
		if err3 != nil {
			log.Fatal(err)
		}
		_, ok = tagmap[id]
		if ok {
			tagmap[id] += ", " + tag
		} else {
			tagmap[id] = tag
		}
	}
	var l Link
	var result string
	for rowsLink.Next() {
		err3 := rowsLink.Scan(&id, &author, &message, &link)
		if err3 != nil {
			log.Fatal(err)
		}
		l = Link{id, author, message, link, tagmap[id]}
		s, _ := json.Marshal(l)
		result += string(s) + ","
	}
	result = "[" + result[:len(result)-1] + "]"
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, result)
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
