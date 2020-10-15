package main

import (
	"fmt"
	"os"
)

// type Process
// func FindProcess(pid int) (*Process, error)
// func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
// func (p *Process) Kill() error
// func (p *Process) Release() error
// func (p *Process) Signal(sig Signal) error
// func (p *Process) Wait() (*ProcessState, error)
func main() {
	// 先运行test获取到测试进程pid
	// find process 25520
	p, e := os.FindProcess(16356)
	fmt.Println(p, e)

	if p != nil {
		// 释放资源
		// e = p.Release()
		// fmt.Println(e)

		// 杀死进程
		// e = p.Kill()
		// fmt.Println(e)

		// 等待进程退出
		w, e := p.Wait()
		fmt.Printf("%+v\n", w)
		fmt.Printf("%+v\n", e)
	}

	// 底层接口  上层使用os/exec包
	// pa := new(os.ProcAttr)
	// tp, e := os.StartProcess("test123", nil, pa)
	// fmt.Println(tp)
	// fmt.Print(e)

}
