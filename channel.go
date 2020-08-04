package coinbase

import (
	"errors"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type messageType `json:"type"`
}

type Error struct {
	Type    messageType `json:"type"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
}

type Subscribe struct {
	Type     messageType `json:"type"`
	Channels interface{} `json:"channels"`
}

type Unsubscribe struct {
	Type messageType `json:"type"`
}

type Channel struct {
	Name       messageType    `json:"name"`
	ProductIDs []CurrencyPair `json:"product_ids"`
}

type Subscriptions struct {
	Type     messageType `json:"type"`
	Channels []Channel   `json:"message"`
}

type Heartbeat struct {
	Type        messageType  `json:"type"`
	Sequence    int64        `json:"sequence"`
	LastTradeID int64        `json:"last_trade_id"`
	ProductID   CurrencyPair `json:"product_id"`
	Time        string       `json:"time"`
}

type Status struct {
	Type       messageType `json:"type"`
	Message    string      `json:"message"`
	Reason     string      `json:"reason"`
	Products   []Product   `json:"products"`
	Currencies []Currency  `json:"currencies"`
}

type messageType string
const (
	activate_message      messageType = "activate"
	change_message        messageType = "change"
	done_message          messageType = "done"
	error_message         messageType = "error"
	heartbeat_message     messageType = "heartbeat"
	l2update_message      messageType = "l2update"
	match_message         messageType = "match"
	open_message          messageType = "open"
	received_message      messageType = "received"
	snapshot_message      messageType = "snapshot"
	status_message        messageType = "status"
	subscribe_message     messageType = "subscribe"
	subscriptions_message messageType = "subscriptions"
	ticker_message        messageType = "ticker"
	unsubscribe_message   messageType = "unsubscribe"
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
		Type:     subscribe_message,
		Channels: []interface{} {
			Channel {
				Name: heartbeat_message,
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

	if subscriptions.Type != subscriptions_message {
		return errors.New("failed to subscribe to heartbeat channel")
	}

	for {
		var heartbeat Heartbeat
		if err := websocket.JSON.Receive(wsclient, &heartbeat); err != nil {
			return err
		}

		if heartbeat.Type == heartbeat_message {
			if err := handler(heartbeat); err != nil {
				return err
			}
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

