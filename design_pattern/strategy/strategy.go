/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/19 19:46
 * @version     v1.0
 * @filename    strategy.go
 * @description
 ***************************************************************************/
package strategy

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (s *strategyA) Execute() {
	fmt.Println("Plan A executed.")
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

type strategyB struct {
}

func (s *strategyB) Execute() {
	fmt.Println("Plan B executed")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Execute() {
	c.strategy.Execute()
}
