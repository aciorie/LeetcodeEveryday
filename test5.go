package main

import "fmt"

// 了解过选项模式吗？
// 能否写一段代码实现一个函数选项模式？

// 选项设计模式
// 问题：有一个结构体，定义一个函数，
// 给结构体初始化

// 结构体
type Options struct {
	str1 string
	int1 int
}

// 声明一个函数类型的变量，用于传参
type Option func(opts *Options)

func InitOptions(opts ...Option) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	fmt.Printf("options:%#v\n", options)
}

func WithString1(str string) Option {
	return func(opts *Options) {
		opts.str1 = str
	}
}

func WithInt1(int1 int) Option {
	return func(opts *Options) {
		opts.int1 = int1
	}
}

func main5() {
	InitOptions(WithString1("ace"), WithInt1(5))
}
