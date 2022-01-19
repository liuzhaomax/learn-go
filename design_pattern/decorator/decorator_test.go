/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/19 19:26
 * @version     v1.0
 * @filename    decorator_test.go
 * @description
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
