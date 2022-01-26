/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/26 23:16
 * @version     v1.0
 * @filename    composition_test.go
 * @description 调用container的0位置的composition群体的遍历方法，所有container的个体都会执行遍历方法
 ***************************************************************************/
package composition

import (
	"fmt"
	"testing"
)

func TestComposition_Traverse(t *testing.T) {
	container := make([]Composition, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			container[i].Add(NewLeaf(i*3 + j))
		}
	}
	for i := 1; i < 4; i++ {
		container[0].Add(&container[i]) //把其他组合，组合到container[0]中
	}
	for i := 0; i < 4; i++ {
		container[i].Traverse()
		fmt.Println("Finished")
	}
}
