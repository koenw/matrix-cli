package main

import (
	"os"

	"github.com/matrix-org/gomatrix"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("matrix-cli", "Command line client for the [matrix] decentralized communication network.")

	serverURL = app.Flag("server-url", "URL to the matrix server").Default("https://matrix.org").String()
	username  = app.Flag("username", "matrix username").Required().String()
	password  = app.Flag("password", "matrix password").Required().String()

	send         = app.Command("send", "Send a message to a matrix room")
	sendRoom     = send.Arg("room", "matrix room").Required().String()
	sendMessages = send.Arg("message", "Message to send to the matrix room").Required().Strings()
)

func main() {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	client, err := gomatrix.NewClient(*serverURL, "", "")
	if err != nil {
		panic(err)
	}

	login(client, username, password)

	switch cmd {
	case "send":
		sendMsgs(client, *sendRoom, *sendMessages)
	}

}

func login(client *gomatrix.Client, username, password *string) {
	loginResp, err := client.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     *username,
		Password: *password,
	})
	if err != nil {
		panic(err)
	}

	client.SetCredentials(loginResp.UserID, loginResp.AccessToken)
}

func sendMsgs(client *gomatrix.Client, room string, messages []string) {
	resp, err := client.JoinRoom(room, "", nil)
	if err != nil {
		panic(err)
	}

	for _, msg := range messages {
		_, err = client.SendText(resp.RoomID, msg)
		if err != nil {
			panic(err)
		}
	}
}
