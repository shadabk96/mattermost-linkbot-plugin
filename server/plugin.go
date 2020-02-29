package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/v5/plugin"

	"database/sql"
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

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {

	// get Mattermost-User-Id
	userID := r.Header.Get("Mattermost-User-Id")
	log.Println("userID" + userID)
	// get channels for user
	team := "8r4yr7x4nfdytndqn1sidruuxw"
	channels, err4 := p.API.GetChannelsForTeamForUser(team, userID, false)
	if err4 != nil {
		log.Fatal(err4, "failed to get channels for team: %s for user: %s", team, userID)
	}
	log.Println(channels)
	// get links for channel from db
	channelsListStr := "'"
	for _, channel := range channels {
		channelsListStr += channel.Name + "','"
	}
	channelsListStr = channelsListStr[:len(channelsListStr)-2]
	log.Println(channelsListStr)

	db, err := sql.Open("sqlite3", "/home/shadab/.local/lib/python2.7/site-packages/mmpy_bot/sqlite.db")
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
	rows, err2 := db.Query("select l.id, l.author, l.message, l.link, t.tag from links l join tags t on l.id = t.message_id where l.channel in (" + channelsListStr + ")")
	if err != nil {
		log.Fatal(err2)
	}
	defer rows.Close()
	tagmap := make(map[int]string)
	var ok bool
	for rows.Next() {
		err3 := rows.Scan(&id, &author, &message, &link, &tag)
		if err3 != nil {
			log.Fatal(err)
		}
		log.Println(id, author, message, link, tag)
		_, ok = tagmap[id]
		if ok {
			tagmap[id] += ", " + tag
		} else {
			tagmap[id] = tag
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tagmap)

	// convert to json
	// return
	fmt.Fprint(w, "Hello, world!")
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
