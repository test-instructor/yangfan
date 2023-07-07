//go:build windows

package dial

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"

	"github.com/test-instructor/yangfan/hrp/internal/builtin"
	"github.com/test-instructor/yangfan/hrp/internal/myexec"
)

var (
	regexTracertPass    = regexp.MustCompile(`(\d+)[\s*<]+(\d+)\s+ms`)
	regexTracertFailure = regexp.MustCompile(`(\d+)[\s*]+Request timed out`)
)

func DoTraceRoute(traceRouteOptions *TraceRouteOptions, args []string) (err error) {
	if len(args) != 1 {
		return errors.New("there should be one argument")
	}
	var traceRouteResult TraceRouteResult
	defer func() {
		if traceRouteOptions.SaveTests {
			dir, _ := os.Getwd()
			traceRouteResultName := fmt.Sprintf("traceroute_result_%v.json", time.Now().Format("20060102150405"))
			traceRouteResultPath := filepath.Join(dir, traceRouteResultName)
			err = builtin.Dump2JSON(traceRouteResult, traceRouteResultPath)
			if err != nil {
				global.GVA_LOG.Error("save traceroute result failed", zap.Error(err))
			}
		}
	}()

	traceRouteTarget := args[0]
	parsedURL, err := url.Parse(traceRouteTarget)
	if err == nil && parsedURL.Host != "" {
		global.GVA_LOG.Info("parse input url", zap.String("url", traceRouteTarget), zap.String("host", parsedURL.Host))
		traceRouteTarget = strings.Split(parsedURL.Host, ":")[0]
	}

	cmd := myexec.Command("tracert", "-h", strconv.Itoa(traceRouteOptions.MaxTTL), traceRouteTarget)
	global.GVA_LOG.Info("execute traceroute", zap.String("command", cmd.String()))
	stdout, _ := cmd.StdoutPipe()

	startT := time.Now()
	defer func() {
		global.GVA_LOG.Info("traceroute costs", zap.String("target", traceRouteTarget), zap.Duration("cost", time.Since(startT)))
	}()

	global.GVA_LOG.Info("start to traceroute", zap.String("target", traceRouteTarget))
	err = cmd.Start()
	if err != nil {
		traceRouteResult.Suc = false
		traceRouteResult.ErrMsg = "execute traceroute failed"
		global.GVA_LOG.Error("start command failed", zap.Error(err))
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		hopLine := scanner.Text()
		fmt.Println(hopLine)
		failureLine := regexTracertFailure.FindStringSubmatch(hopLine)
		if len(failureLine) == 2 {
			hopID, _ := strconv.Atoi(failureLine[1])
			traceRouteResult.Details = append(traceRouteResult.Details, TraceRouteResultNode{
				Id: hopID,
			})
			continue
		}
		passLine := regexTracertPass.FindStringSubmatch(hopLine)
		if len(passLine) == 3 {
			hopID, _ := strconv.Atoi(passLine[1])
			fields := strings.Fields(hopLine)
			hopIP := strings.Trim(fields[len(fields)-1], "[]")
			traceRouteResult.Details = append(traceRouteResult.Details, TraceRouteResultNode{
				Id:   hopID,
				Ip:   hopIP,
				Time: passLine[2],
			})
			traceRouteResult.Suc = true
		}
	}
	hopCount := len(traceRouteResult.Details)
	traceRouteResult.IP = traceRouteResult.Details[hopCount-1].Ip
	err = cmd.Wait()
	if err != nil {
		traceRouteResult.Suc = false
		traceRouteResult.ErrMsg = "wait traceroute finish failed"
		global.GVA_LOG.Error("wait command failed", zap.Error(err))
		return
	}
	return
}
