package main

func main() {
	var p Publisher
	p = newPublisher()
	p.broadcast("welcome")
}

type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subId string)
	broadcast(msg string)
}

type Subscriber interface {
	id() string
	react(msg string)
}

type publisher struct {
	subscribers map[string]Subscriber
}

func newPublisher() publisher {
	return publisher{subscribers: make(map[string]Subscriber)}
}

func (p publisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[subscriber.id()] = subscriber
}
func (p publisher) removeSubscriber(subId string) {
	delete(p.subscribers, subId)
}

func (p publisher) broadcast(msg string) {
	for _, subscriber := range p.subscribers {
		subscriber.react(msg)
	}
}

type subscriber struct {
	subId string
}

func newSubscriber(subId string) subscriber {
	return subscriber{subId: subId}
}

func (s subscriber) id() string {

}

func (s subscriber) react(msg string) {

}
