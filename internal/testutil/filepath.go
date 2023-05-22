package testutil

import (
	"path/filepath"
	"testing"
)

// AbsFilePath 使用这个函数比直接使用filepath.Abs(path)多了错误的处理，
// 通过创建一个包含错误处理的函数，可以避免每次调用这个函数的地方都进行错处处理
// `testing.T` 是内置的测试类型，通常记录测试状态和属性
func AbsFilePath(t *testing.T, path string) string {

	// t.Helper() 表示打印日志的时候不会打印一大堆堆栈信息，
	// 他告诉测试框架，当前函数不是测试逻辑，而是被其他函数调用共享一些共同的逻辑
	t.Helper()

	s, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}

	return s
}
