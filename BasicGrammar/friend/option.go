// @program: 不定参数传递
// @author: aslanwang
// @create: 2021-11-6
package friend

import (
	"sync"
)

var (
	cache = &sync.Pool{
		New : func() interface{}{
			return &option{}
		},
	}
)

// 规定Option的类型是函数func(*option) 无返回值
type Option func(*option)

type option struct{
	sex    int
	age    int
	height int
	weight int
	hobby  string
}

// reset 实现option的reset方法
func (o *option) reset() {
	o.sex = 0
	o.age = 0
	o.height = 0
	o.weight = 0
	o.hobby = ""
}

func getOption() *option {
	// 一种类型断言 仅在Get()返回的对象可以被强制转为*option时 正常运行
	return cache.Get().(*option)
}

func releaseOption(opt *option) {
	opt.reset()
	cache.Put(opt)
}

// WithSex setup sex, 1=female 2=male
func WithSex(sex int) Option {
	return func(opt *option) {
		opt.sex = sex
	}
}

// withAge 写法给定参数 返回函数变量
func WithAge(v int) Option{
	return func(o *option){
		o.age = v
	}
}

// WithHeight set up height
func WithHeight(height int) Option {
	return func(opt *option) {
		opt.height = height
	}
}

// WithWeight set up weight
func WithWeight(weight int) Option {
	return func(opt *option) {
		opt.weight = weight
	}
}

// WithHobby set up Hobby
func WithHobby(hobby string) Option {
	return func(opt *option) {
		opt.hobby = hobby
	}
}
