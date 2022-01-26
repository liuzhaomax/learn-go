/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/19 19:26
 * @version     v1.0
 * @filename    decorator_test.go
 * @description 执行Pi的时候会在外面包裹一个函数，返回Pi，保留Pi的参数，然后在包裹里加点别的东西，比如log
 ***************************************************************************/
package decorator

import (
	"log"
	"os"
	"testing"
)

func TestWrapLogger(t *testing.T) {
	foo := WrapLogger(Pi, log.New(os.Stdout, "test", 1))
	foo(10000)
}
