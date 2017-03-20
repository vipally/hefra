package event

import (
	"github.com/vipally/hefra/regable"
)

var (
	eventIdMgr = regable.NewIdNameMapping()
)

type EventId uint32

//define an event id
func DefineEvent(id uint32, name string) (event EventId) {
	if _id, err := eventIdMgr.Insert(id, name); err == nil {
		event = EventId(_id)
	} else {
		panic(name)
	}
	return
}
