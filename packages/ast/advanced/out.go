package advanced

import "log"

type WelcomeDefaultImpl struct {
}

func (w *WelcomeDefaultImpl) Hello(name string) interface{} { log.Println("name = ", name); return nil }
