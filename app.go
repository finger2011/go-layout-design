package main

import "context"

//App app object
type App struct {
	name string
	version string
	ctx context.Context
	cancel func () 
	servers []interface{} // server实现因为时间关系来不及写了
	metadata map[string]string
}

//NewApp new app
func NewApp() *App {
	//初始化
	//然后写入App对象
	return &App{}
}

//Run run app
func (app *App) Run() error {
	//var w waitgroup 
	//w.Go起各个server
	//w.Go起signal.Notify
	//大概思路类似https://github.com/finger2011/go-demo/blob/master/third_week/main.go 里的main函数
	return nil
}

//Stop 停止 app
func (app *App) Stop() error {
	if app.cancel != nil {
		app.cancel()
	}
	return nil
}
