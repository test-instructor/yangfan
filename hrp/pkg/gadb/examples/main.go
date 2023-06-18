package main

import (
	"os"
	"strings"

	"github.com/test-instructor/yangfan/hrp/pkg/gadb"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
)

func main() {
	adbClient, err := gadb.NewClient()
	checkErr(err, "fail to connect adb server")

	devices, err := adbClient.DeviceList()
	checkErr(err)

	if len(devices) == 0 {
		global.GVA_LOG.Fatal("list of devices is empty")
	}

	dev := devices[0]

	userHomeDir, _ := os.UserHomeDir()
	apk, err := os.Open(userHomeDir + "/Desktop/xuexi_android_10002068.apk")
	checkErr(err)

	global.GVA_LOG.Info("starting to push apk")

	remotePath := "/data/local/tmp/xuexi_android_10002068.apk"
	err = dev.PushFile(apk, remotePath)
	checkErr(err, "adb push")

	global.GVA_LOG.Info("starting to install apk")

	shellOutput, err := dev.RunShellCommand("pm install", remotePath)
	checkErr(err, "pm install")
	if !strings.Contains(shellOutput, "Success") {
		global.GVA_LOG.Fatal("fail to install", zap.String("output", shellOutput))
	}

	global.GVA_LOG.Info("install completed")

}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	global.GVA_LOG.Fatal(output)
}
