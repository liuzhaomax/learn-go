/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 3:42
 * @version     v1.0
 * @filename    singleton.go
 * @description
 ***************************************************************************/
package singleton

import "sync"

var ctx *Context
var once sync.Once

func init() {
	once.Do(func() {
		ctx = &Context{}
	})
}

func GetInstanceOfContext() *Context {
	return ctx
}

type Context struct {
	Name string
}

func (c *Context) Work() {

}
