package main

import (
  "github.com/mitchellh/mapstructure"
  r "github.com/dancannon/gorethink"
)

func addChannel(client *Client, data interface{}) {
  var channel Channel
  err := mapstructure.Decode(data, &channel)
  if err != nil {
    client.send <- Message{"error", err.Error()}
    return
  }
  go func() {
    err = r.Table("channel").
      Insert(channel).
      Exec(client.session)
    if err != nil {
      client.send <- Message{"error", err.Error()}
    }
  }()
}
