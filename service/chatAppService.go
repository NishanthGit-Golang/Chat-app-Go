package service

import (
	"fmt"
	"sync"
)

type Client struct {
	ID      string
	Message chan string
}

type ChatRoom struct {
	clients  map[string]*Client
	messages chan string
	mutex    sync.Mutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:  make(map[string]*Client),
		messages: make(chan string),
	}
}

func (c *ChatRoom) Join(id string) (*Client, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.clients[id]; exists {
		return nil, fmt.Errorf("client already exists")
	}

	client := &Client{ID: id, Message: make(chan string, 10)}
	c.clients[id] = client
	return client, nil
}

func (c *ChatRoom) Leave(id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if client, exists := c.clients[id]; exists {
		close(client.Message)
		delete(c.clients, id)
		return nil
	}

	return fmt.Errorf("client not found")
}

func (c *ChatRoom) SendMessage(id, msg string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, exists := c.clients[id]; !exists {
		return fmt.Errorf("client not found")
	}

	c.messages <- fmt.Sprintf("%s: %s", id, msg)
	return nil
}

func (c *ChatRoom) BroadcastMessages() {
	for msg := range c.messages {
		c.mutex.Lock()
		for _, client := range c.clients {
			select {
			case client.Message <- msg:
			default:
			}
		}
		c.mutex.Unlock()
	}
}

func (c *ChatRoom) GetMessages(id string) (string, error) {
	c.mutex.Lock()
	client, exists := c.clients[id]
	c.mutex.Unlock()

	if !exists {
		return "", fmt.Errorf("client not found")
	}

	select {
	case msg := <-client.Message:
		return msg, nil
	default:
		return "No new messages", nil
	}
}
