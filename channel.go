package coinbase

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"sync"
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
	sync.Mutex
	polling       bool
	client        *CoinbaseProClient
	subscriptions chan Subscriptions
	heartbeat     chan Heartbeat
}

func (channel *channelService) poll() {
	for channel.polling {
		var payload map[string]interface{}
		if err := websocket.JSON.Receive(channel.client.socket, &payload); err != nil {
			continue
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			continue
		}

		mtype, ok := payload["type"]
		if !ok {
			continue
		}

		switch messageType(mtype.(string)) {
		case subscriptions_message:
			var subscriptions Subscriptions
			if err := json.Unmarshal(marshalled, &subscriptions); err == nil && channel.subscriptions != nil {
				channel.subscriptions <- subscriptions
			}
		case heartbeat_message:
			var heartbeat Heartbeat
			if err := json.Unmarshal(marshalled, &heartbeat); err == nil && channel.heartbeat != nil {
				channel.heartbeat <- heartbeat
			}
		}
	}
}

func (channel *channelService) beginPolling() {
	if !channel.polling {
		channel.Lock()
		if !channel.polling {
			channel.polling = true
			go channel.poll()
		}
		channel.Unlock()
	}
}

func (channel *channelService) SubscribeToSubscriptions() <-chan Subscriptions {
	if channel.subscriptions == nil {
		channel.subscriptions = make(chan Subscriptions)
		channel.beginPolling()
	}
	return channel.subscriptions
}

func (channel *channelService) SubscribeToHeartbeat(currencyPairs ...CurrencyPair) <-chan Heartbeat {
	if channel.heartbeat == nil {
		channel.heartbeat = make(chan Heartbeat)
		channel.beginPolling()
	}

	subscribe := Subscribe {
		Type: subscribe_message,
		Channels: []interface{} {
			Channel {
				Name: heartbeat_message,
				ProductIDs: currencyPairs,
			},
		},
	}

	_ = websocket.JSON.Send(channel.client.socket, subscribe)

	return channel.heartbeat
}

func (channel *channelService) SubscribeToStatus() {

}

func (channel *channelService) SubscribeToTicker() {

}

func (channel *channelService) SubscribeToLevelTwo() {

}

func (channel *channelService) SubscribeToUser() {

}

func (channel *channelService) SubscribeToMatches() {

}

func (channel *channelService) SubscribeToFull() {

}

