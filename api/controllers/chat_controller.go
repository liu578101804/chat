package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"fmt"
	"sync"
	"strconv"
)

//映射表
var clientMap map[int64]*Node = make(map[int64]*Node,0)
//读写锁
var rwLocker sync.RWMutex

type Node struct {
	Conn *websocket.Conn
	DataQueue chan []byte
}

type ChatController struct {
	BaseController
}

func (this *ChatController)WxChat(c *gin.Context)  {

	id := c.Query("id")
	userId ,_ := strconv.ParseInt(id,10,64)

	//升级get请求为webSocket协议
	conn,err := (&websocket.Upgrader{
		CheckOrigin: func (r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print(err.Error())
		return
	}

	//defer conn.Close()
	node := &Node{
		Conn: conn,
		DataQueue: make(chan []byte, 50),
	}

	rwLocker.Lock()
	clientMap[userId]=node
	rwLocker.Unlock()

	sendMsg(node, []byte(fmt.Sprintf("%v 加入聊天室",userId)))

	go sendProcess(node)
	go receiveProcess(node)

	//设置关闭时回调
	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("连接关闭了", code, text)
		sendMsg(node, []byte(fmt.Sprintf("%v 离开了聊天室",userId)))
		delete(clientMap, userId)
		return nil
	})
}

func dispatch(data []byte)  {
	fmt.Printf("收到的信息: %s \n",data)
}

func sendMsg(node *Node, data []byte)  {
	node.DataQueue <- data
}

func sendProcess(node *Node)  {
	for {
		select {
		case data := <-node.DataQueue:

			for _,v := range clientMap {
				err := v.Conn.WriteMessage(websocket.TextMessage, data)
				if err != nil{
					log.Println(err.Error())
				}
			}
		}
	}
}

func receiveProcess(node *Node)  {
	for {
		_,data,err := node.Conn.ReadMessage()
		if err != nil{
			log.Println(err.Error())
			return
		}
		dispatch(data)

		//回写数据
		sendMsg(node, data)
	}
}

func NewChatController() *ChatController {
	return &ChatController{}
}

