package interpreter

import "github.com/indeedeng/iwf/gen/iwfidl"

type InternalChannel struct {
	// key is channel name
	receivedData map[string][]*iwfidl.EncodedObject
}

func NewInterStateChannel() *InternalChannel {
	return &InternalChannel{
		receivedData: map[string][]*iwfidl.EncodedObject{},
	}
}

func RebuildInterStateChannel(refill map[string][]*iwfidl.EncodedObject) *InternalChannel {
	return &InternalChannel{
		receivedData: refill,
	}
}

func (i *InternalChannel) DumpReceived() map[string][]*iwfidl.EncodedObject {
	return i.receivedData
}

func (i *InternalChannel) HasData(channelName string) bool {
	l := i.receivedData[channelName]
	return len(l) > 0
}

func (i *InternalChannel) ProcessPublishing(publishes []iwfidl.InterStateChannelPublishing) {
	for _, pub := range publishes {
		i.receive(pub.ChannelName, pub.Value)
	}
}

func (i *InternalChannel) receive(channelName string, data *iwfidl.EncodedObject) {
	l := i.receivedData[channelName]
	l = append(l, data)
	i.receivedData[channelName] = l
}

func (i *InternalChannel) Retrieve(channelName string) *iwfidl.EncodedObject {
	l := i.receivedData[channelName]
	if len(l) <= 0 {
		panic("critical bug, this shouldn't happen")
	}
	data := l[0]
	l = l[1:]
	if len(l) == 0 {
		delete(i.receivedData, channelName)
	} else {
		i.receivedData[channelName] = l
	}

	return data
}
