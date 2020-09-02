package infra

/**
 *@Author tudou
 *@Date 2020/7/27
 **/

//初始化接口
type Initializer interface {
	//用户对象实例化后的初始化操作
	Init()
}

//初始化注册器
type InitializeRegister struct {
	Initializers []Initializer
}

//注册一个初始化对象
func (i *InitializeRegister) Register(ai Initializer) {
	i.Initializers = append(i.Initializers, ai)
}