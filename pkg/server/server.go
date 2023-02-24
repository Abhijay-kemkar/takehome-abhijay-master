package server

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
	"slai.io/takehome/pkg/common"

)

type Client struct {
	host string
	ws   *websocket.Conn
}

const addr = "localhost:5555"

var upgrader = websocket.Upgrader{}
var wg sync.WaitGroup

func handleMessage(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	client := Client{
		ws:   c,
		host: r.Host,
	}

	log.Println("Connected to host: ", client.host)

	defer c.Close()
	defer wg.Wait()

	for {
		_, msg, err := c.ReadMessage()

		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var request common.BaseRequest
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Println("Invalid request:", err)
			break
		}

		switch request.RequestType {
		case string(common.Echo):
			go HandleEcho(msg, &client)
		case string(common.Sync):
			go HandleSync(msg, &client)
		}
	}
}

func HandleSync(msg []byte, client *Client) {
	var syncReq common.SyncRequest
	err := json.Unmarshal(msg, &syncReq)
	if err != nil {
		log.Println("Invalid sync request:", err)
		return
	}

	filePath := filepath.Join(syncReq.Directory, syncReq.Filename)
	err = ioutil.WriteFile(filePath, []byte(syncReq.Contents), os.ModePerm)
	if err != nil {
		log.Println("Error writing file:", err)
		SendResponse(client.ws, common.BaseResponse{Success: false})
		return
	}

	log.Println("File written successfully: ", filePath)
	SendResponse(client.ws, common.BaseResponse{Success: true})
}

func SendResponse(ws *websocket.Conn, response common.BaseResponse) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println("Error marshalling response:", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, responseJson)
	if err != nil {
		log.Println("Error sending response:", err)
		return
	}
}

func StartServer() {
	flag.Parse()
	http.HandleFunc("/", handleMessage)
	log.Println("Starting server @", addr)
	log.Fatal((http.ListenAndServe(addr, nil)))
}
