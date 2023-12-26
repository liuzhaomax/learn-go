package uuid

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"testing"
)

func TestUUID(t *testing.T) {
	//返回基于当前时间戳和MAC地址的UUID。
	u1 := uuid.NewV1()
	u11 := uuid.NewV1()
	fmt.Println("u1 >>>", u1)
	fmt.Println("u11 >>>", u11)

	//返回基于POSIX UID/GID的DCE安全UUID。
	u2 := uuid.NewV2(0)
	fmt.Println("u2 >>>", u2)

	//返回基于命名空间UUID和名称的MD5哈希的UUID。
	u3 := uuid.NewV3(u2, "abc")
	fmt.Println("u3 >>>", u3)

	//返回随机生成的UUID。
	u4 := uuid.NewV4()
	fmt.Println("u4 >>>", u4)

	//返回基于命名空间UUID和名称的SHA-1哈希的UUID。
	u5 := uuid.NewV5(u2, "abc")
	fmt.Println("u5 >>>", u5)

	//本质上  没有差距 不过更新版本之前 某些方法 是有两个返回值的，Must是基于前者 进行的一个包装
	u6 := uuid.Must(u5, nil)
	fmt.Println("u6 >>>", u6)

	//如果 u5 和 u6 相等，则返回true，否则返回false。
	u7 := uuid.Equal(u5, u6)
	fmt.Println("u7 >>>", u7)

	// 将字符串转换成UUID
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	uuidObject, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	fmt.Println("successfully parsed UUID Version", uuidObject)
}
