package coinbase

import (
	"errors"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type MessageType `json:"type"`
}

type Error struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
}

type Subscribe struct {
	Type     MessageType `json:"type"`
	Channels interface{} `json:"channels"`
}

type Unsubscribe struct {
	Type MessageType `json:"type"`
}

type Channel struct {
	Name       MessageType    `json:"name"`
	ProductIDs []CurrencyPair `json:"product_ids"`
}

type Subscriptions struct {
	Type     MessageType `json:"type"`
	Channels []Channel   `json:"message"`
}

type Heartbeat struct {
	Type        MessageType  `json:"type"`
	Sequence    int64        `json:"sequence"`
	LastTradeID int64        `json:"last_trade_id"`
	ProductID   CurrencyPair `json:"product_id"`
	Time        string       `json:"time"`
}

type Status struct {
	Type        MessageType `json:"type"`
	Message     string      `json:"message"`
	Reason      string      `json:"reason"`
	Products    []Product   `json:"products"`
	Currencies  []Currency  `json:"currencies"`
}

type MessageType string
const (
	ERROR         MessageType = "error"

	// Subscription Messages
	SUBSCRIBE     MessageType = "subscribe"
	UNSUBSCRIBE   MessageType = "unsubscribe"
	SUBSCRIPTIONS MessageType = "subscriptions"

	// Channel Messages
	HEARTBEAT     MessageType = "heartbeat"
	STATUS        MessageType = "status"
	TICKER        MessageType = "ticker"

	// Level2 Channel Messages
	SNAPSHOT      MessageType = "snapshot"
	L2UPDATE      MessageType = "l2update"

	// Full Channel Messages
	RECEIVED      MessageType = "received"
	OPEN          MessageType = "open"
	DONE          MessageType = "done"
	MATCH         MessageType = "match"
	CHANGE        MessageType = "change"
	ACTIVATE      MessageType = "activate"
)

// https://docs.pro.coinbase.com/#channels
type channelService struct {
	client *CoinbaseProClient
}

func (channel channelService) wsclient() (*websocket.Conn, error) {
	return websocket.Dial(channel.client.websocketURL, "", channel.client.origin)
}

func (channel channelService) SubscribeToHeartbeat(handler func(heartbeat Heartbeat) error, currencyPairs ...CurrencyPair) error {
	subscribe := Subscribe {
		Type:     SUBSCRIBE,
		Channels: []interface{} {
			Channel {
				Name: HEARTBEAT,
				ProductIDs: currencyPairs,
			},
		},
	}

	wsclient, err := channel.wsclient()
	if err != nil {
		return err
	}
	defer wsclient.Close()

	if err := websocket.JSON.Send(wsclient, subscribe); err != nil {
		return err
	}

	var subscriptions Subscriptions
	if err := websocket.JSON.Receive(wsclient, &subscriptions); err != nil {
		return err
	}

	if subscriptions.Type != SUBSCRIPTIONS {
		return errors.New("failed to subscribe to heartbeat channel")
	}

	for {
		var heartbeat Heartbeat
		if err := websocket.JSON.Receive(wsclient, &heartbeat); err != nil {
			return err
		}

		if heartbeat.Type == HEARTBEAT {
			if err := handler(heartbeat); err != nil {
				return err
			}
			continue
		}
	}
}

func (channel channelService) SubscribeToStatus() {

}

func (channel channelService) SubscribeToTicker() {

}

func (channel channelService) SubscribeToLevelTwo() {

}

func (channel channelService) SubscribeToUser() {

}

func (channel channelService) SubscribeToMatches() {

}

func (channel channelService) SubscribeToFull() {

}

