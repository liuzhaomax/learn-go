/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/19 19:56
 * @version     v1.0
 * @filename    strategy_test.go
 * @description
 ***************************************************************************/
package strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategyA := NewStrategyA()
	c := NewContext()
	c.SetStrategy(strategyA)
	c.Execute()

	strategyB := NewStrategyB()
	c.SetStrategy(strategyB)
	c.Execute()
}
