package cal

import (
	"testing"
)

/**
 golang 单元测试框架，使用方式 方法名称以Test开头，后续不能为a-z的其他字符命名
执行方式：
 1.在编译器执行Test*方法
 2.在控制台进入对应目录，执行命令go test 此方法会自动调用所有的测试用例
  go test -v 调用测试用例并打印日志
 指定运行某个测试用例
go test -v -test.run TestGetSum

*/
func TestGetSum(t *testing.T) {
	sum := getSum(1, 3)

	//用于输出日志
	t.Logf("调用getSum方法单元测试执行结果=%v", sum)
	t.Fatalf("调用getSum方法单元测试执行结果=%v,期望值为=%v", sum, 4)

}

func TestAddUpper(t *testing.T) {
	sum := addUpper(10)

	t.Logf("调用addUpper方法单元测试执行结果=%v", sum)
}

func TestGetSub(t *testing.T) {
	sub := getSub(109, 5)

	t.Logf("调用getSub方法单元测试执行结果=%v", sub)
}
