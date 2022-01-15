/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 3:44
 * @version     v1.0
 * @filename    singleton_test.go
 * @description
 ***************************************************************************/
package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstanceOfContext(t *testing.T) {
	context := GetInstanceOfContext()
	context.Work()
	context.Name = "123"

	ctx := new(Context)
	ctx.Work()

	fmt.Println(&ctx, &context)
	fmt.Println(GetInstanceOfContext().Name)
}
