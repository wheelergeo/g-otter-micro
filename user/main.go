// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	usercmd "github.com/wheelergeo/g-otter-micro-user/cmd"
)

func main() {
	err := usercmd.Command().Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
