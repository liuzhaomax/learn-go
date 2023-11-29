package responsibility_chain

import (
	"fmt"
	"strconv"
)

type IHandler interface {
	Handle(handlerID int) string
}

type Handler struct {
	name      string
	next      IHandler
	handlerID int
}

func NewHandler(name string, next IHandler, handlerID int) *Handler {
	return &Handler{
		name:      name,
		next:      next,
		handlerID: handlerID,
	}
}

func (h *Handler) Handle(handlerID int) string {
	if h.handlerID == handlerID {
		return h.name + " handled " + strconv.Itoa(handlerID)
	}
	if h.next == nil {
		return ""
	}
	fmt.Println(h.name + " pass to next")
	return h.next.Handle(handlerID)
}
