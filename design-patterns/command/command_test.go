package command

import "testing"

func TestCommand_Execute(t *testing.T) {
	wang := NewPerson("wang", NewCommand(nil, nil))
	zhang := NewPerson("zhang", NewCommand(wang, wang.Listen))
	feng := NewPerson("feng", NewCommand(zhang, zhang.Buy))
	ding := NewPerson("ding", NewCommand(feng, feng.Cook))
	li := NewPerson("li", NewCommand(ding, ding.Wash))

	li.Talk()

}
