package uixt

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"embed"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lingcetech/funplugin/myexec"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/httprunner/code"
	"github.com/test-instructor/yangfan/httprunner/internal/builtin"
	"github.com/test-instructor/yangfan/httprunner/internal/config"
	"github.com/test-instructor/yangfan/httprunner/internal/json"
	"github.com/test-instructor/yangfan/httprunner/pkg/gadb"
	"github.com/test-instructor/yangfan/httprunner/uixt/option"
	"github.com/test-instructor/yangfan/httprunner/uixt/types"
)

const (
	EvalInstallerPackageName   = "sogou.mobile.explorer"
	InstallViaInstallerCommand = "am start -S -n sogou.mobile.explorer/.PackageInstallerActivity -d"
)

//go:embed evalite
var evalite embed.FS

func NewAndroidDevice(opts ...option.AndroidDeviceOption) (device *AndroidDevice, err error) {
	androidOptions := option.NewAndroidDeviceOptions(opts...)

	// get all attached android devices
	adbClient, err := gadb.NewClientWith(
		androidOptions.AdbServerHost, androidOptions.AdbServerPort)
	if err != nil {
		return nil, errors.Wrap(code.DeviceConnectionError, err.Error())
	}
	devices, err := adbClient.DeviceList()
	if err != nil {
		return nil, errors.Wrap(code.DeviceConnectionError, err.Error())
	}
	if len(devices) == 0 {
		return nil, errors.Wrapf(code.DeviceConnectionError,
			"no attached android devices")
	}

	// filter device by serial
	var gadbDevice *gadb.Device
	if androidOptions.SerialNumber == "" {
		if len(devices) > 1 {
			return nil, errors.Wrap(code.DeviceConnectionError,
				"more than one device connected, please specify the serial")
		}
		gadbDevice = devices[0]
		androidOptions.SerialNumber = gadbDevice.Serial()
		log.Warn().Str("serial", androidOptions.SerialNumber).
			Msg("android SerialNumber is not specified, select the attached one")
	} else {
		for _, d := range devices {
			if d.Serial() == androidOptions.SerialNumber {
				gadbDevice = d
				break
			}
		}
		if gadbDevice == nil {
			return nil, errors.Wrapf(code.DeviceConnectionError,
				"android device %s not attached", androidOptions.SerialNumber)
		}
	}

	device = &AndroidDevice{
		Device:  gadbDevice,
		Options: androidOptions,
		Logcat:  NewAdbLogcat(androidOptions.SerialNumber),
	}
	log.Debug().Str("serial", device.Options.SerialNumber).Msg("init android device")

	// setup device
	if err := device.Setup(); err != nil {
		return nil, errors.Wrap(err, "setup android device failed")
	}
	return device, nil
}

type AndroidDevice struct {
	*gadb.Device
	Options *option.AndroidDeviceOptions
	Logcat  *AdbLogcat
}

// RunShellCommand executes adb shell command with a host-side timeout.
//
// Motivation:
// Some shell commands (notably `uiautomator dump`) may hang on certain devices/ROMs.
// Even if upper-layer contexts are cancelled (e.g. case timeout), the underlying gadb
// shell call may keep blocking, which makes the whole step appear "stuck" for minutes.
//
// This method shadows the embedded gadb.Device.RunShellCommand and uses `exec.CommandContext`
// so cancellation/timeout can reliably stop the host-side adb process.
//
// Note:
// We execute via `sh -c` to support pipes/redirection patterns used in this repo
// (e.g. "dumpsys window | grep ...").
func (dev *AndroidDevice) RunShellCommand(cmd string, args ...string) (string, error) {
	return dev.RunShellCommandWithContext(context.Background(), cmd, args...)
}

// RunShellCommandWithContext executes adb shell command with the provided context.
// If ctx has no deadline, a conservative default (10s) will be applied.
func (dev *AndroidDevice) RunShellCommandWithContext(ctx context.Context, cmd string, args ...string) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// conservative default; higher layers can still add their own timeouts
	if _, ok := ctx.Deadline(); !ok {
		timeout := 10 * time.Second
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	serial := ""
	if dev.Options != nil {
		serial = dev.Options.SerialNumber
	}
	if serial == "" && dev.Device != nil {
		serial = dev.Device.Serial()
	}

	// Prefer running as `adb shell <cmd> <args...>` for normal commands to keep behavior
	// consistent with gadb. Only fall back to `sh -c` when we detect shell metacharacters
	// (pipes, redirects, etc.) in the arguments.
	needsShell := false
	for _, a := range append([]string{cmd}, args...) {
		if strings.ContainsAny(a, "|&;<>") || strings.Contains(a, "&&") || strings.Contains(a, "||") {
			needsShell = true
			break
		}
	}

	adbArgs := []string{}
	if serial != "" {
		adbArgs = append(adbArgs, "-s", serial)
	}
	if needsShell {
		fullCmd := strings.TrimSpace(strings.Join(append([]string{cmd}, args...), " "))
		adbArgs = append(adbArgs, "shell", "sh", "-c", fullCmd)
	} else {
		adbArgs = append(adbArgs, append([]string{"shell", cmd}, args...)...)
	}

	c := exec.CommandContext(ctx, "adb", adbArgs...)
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = strings.TrimSpace(stdout.String())
		}
		if msg != "" {
			return "", errors.Wrap(err, msg)
		}
		return "", err
	}
	return stdout.String(), nil
}

func (dev *AndroidDevice) Setup() error {
	dev.Device.RunShellCommand("ime", "enable", option.UnicodeImePackageName)
	dev.Device.RunShellCommand("rm", "-r", config.DeviceActionLogFilePath)

	// setup evalite
	evalToolRaw, err := evalite.ReadFile("evalite")
	if err != nil {
		return errors.Wrap(code.LoadFileError, err.Error())
	}
	err = dev.Device.Push(bytes.NewReader(evalToolRaw), "/data/local/tmp/evalite", time.Now())
	if err != nil {
		return errors.Wrap(code.DeviceShellExecError, err.Error())
	}
	return nil
}

func (dev *AndroidDevice) IsHealthy() (bool, error) {
	state, err := dev.Device.State()
	if err != nil {
		return false, err
	}
	return state == gadb.StateOnline, nil
}

func (dev *AndroidDevice) Teardown() error {
	return nil
}

func (dev *AndroidDevice) UUID() string {
	return dev.Options.SerialNumber
}

func (dev *AndroidDevice) LogEnabled() bool {
	return dev.Options.LogOn
}

func (dev *AndroidDevice) NewDriver() (driver IDriver, err error) {
	// Use CompositeDriver by default to support per-operation driver switching
	if dev.Options.Composite {
		driver, err = NewCompositeDriver(dev)
		if err != nil {
			return nil, errors.Wrap(err, "failed to init Composite driver")
		}
	} else if dev.Options.UIA2 || dev.Options.LogOn {
		driver, err = NewUIA2Driver(dev)
		if err != nil {
			return nil, errors.Wrap(err, "failed to init UIA2 driver")
		}
	} else {
		driver, err = NewADBDriver(dev)
		if err != nil {
			return nil, errors.Wrap(err, "failed to init ADB driver")
		}
	}

	if dev.Options.LogOn {
		err = driver.StartCaptureLog("hrp_adb_log")
		if err != nil {
			return nil, err
		}
	}
	return driver, nil
}

// NewCompositeDriverDirect creates a CompositeDriver directly (convenience method)
func (dev *AndroidDevice) NewCompositeDriverDirect() (*CompositeDriver, error) {
	return NewCompositeDriver(dev)
}

func (dev *AndroidDevice) Install(apkPath string, opts ...option.InstallOption) error {
	installOpts := option.NewInstallOptions(opts...)
	brand, err := dev.Device.Brand()
	if err != nil {
		return err
	}
	args := []string{}
	if installOpts.Reinstall {
		args = append(args, "-r")
	}
	if installOpts.GrantPermission {
		args = append(args, "-g")
	}
	if installOpts.Downgrade {
		args = append(args, "-d")
	}
	switch strings.ToLower(brand) {
	case "vivo":
		return dev.installVivoSilent(apkPath, args...)
	case "oppo", "realme", "oneplus":
		if dev.Device.IsPackageInstalled(EvalInstallerPackageName) {
			return dev.installViaInstaller(apkPath, args...)
		}
		log.Warn().Msg("oppo not install eval installer")
		return dev.installCommon(apkPath, args...)
	default:
		return dev.installCommon(apkPath, args...)
	}
}

func (dev *AndroidDevice) installVivoSilent(apkPath string, args ...string) error {
	currentTime := builtin.GetCurrentDay()
	md5HashInBytes := md5.Sum([]byte(currentTime))
	verifyCode := hex.EncodeToString(md5HashInBytes[:])
	verifyCode = base64.StdEncoding.EncodeToString([]byte(verifyCode))
	verifyCode = verifyCode[:8]
	verifyCode = "-V" + verifyCode
	args = append([]string{verifyCode}, args...)
	_, err := dev.Device.InstallAPK(apkPath, args...)
	return err
}

func (dev *AndroidDevice) installViaInstaller(apkPath string, args ...string) error {
	appRemotePath := "/data/local/tmp/" + strconv.FormatInt(time.Now().UnixMilli(), 10) + ".apk"
	err := dev.Device.PushFile(apkPath, appRemotePath, time.Now())
	if err != nil {
		return err
	}
	done := make(chan error)
	defer func() {
		close(done)
	}()
	logcat := NewAdbLogcatWithCallback(dev.Device.Serial(), func(line string) {
		re := regexp.MustCompile(`\{.*?}`)
		match := re.FindString(line)
		if match == "" {
			return
		}
		var result InstallResult
		err := json.Unmarshal([]byte(match), &result)
		if err != nil {
			log.Warn().Msg("parse Install msg line error: " + match)
			return
		}
		if result.Result == 0 {
			// 安装成功
			done <- nil
		} else {
			done <- errors.New(match)
		}
	})
	err = logcat.CatchLogcat("PackageInstallerCallback")
	if err != nil {
		return err
	}
	defer func() {
		_ = logcat.Stop()
	}()

	// 需要监听是否完成安装
	command := strings.Split(InstallViaInstallerCommand, " ")
	args = append(command, appRemotePath)
	_, err = dev.Device.RunShellCommand("am", args[1:]...)
	if err != nil {
		return err
	}
	// 等待安装完成或超时
	timeout := 8 * time.Minute
	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("install via installer timed out after %v", timeout)
	}
}

type InstallResult struct {
	Result    int    `json:"result"`
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func (dev *AndroidDevice) installCommon(apkPath string, args ...string) error {
	_, err := dev.Device.InstallAPK(apkPath, args...)
	return err
}

func (dev *AndroidDevice) Uninstall(packageName string) error {
	_, err := dev.Device.Uninstall(packageName)
	return err
}

func (dev *AndroidDevice) GetCurrentWindow() (windowInfo types.WindowInfo, err error) {
	// adb shell dumpsys window | grep -E 'mCurrentFocus|mFocusedApp'
	output, err := dev.Device.RunShellCommand("dumpsys", "window", "|", "grep", "-E", "'mCurrentFocus|mFocusedApp'")
	if err != nil {
		return types.WindowInfo{}, errors.Wrap(err, "get current window failed")
	}
	// mCurrentFocus=Window{a33bc55 u0 com.miui.home/com.miui.home.launcher.Launcher}
	reFocus := regexp.MustCompile(`mCurrentFocus=Window{.*? (\S+)/(\S+)}`)
	matches := reFocus.FindStringSubmatch(output)
	if len(matches) == 3 {
		windowInfo = types.WindowInfo{
			PackageName: matches[1],
			Activity:    matches[2],
		}
		return windowInfo, nil
	}
	// mFocusedApp=ActivityRecord{2db504f u0 com.miui.home/.launcher.Launcher t2}
	reApp := regexp.MustCompile(`mFocusedApp=ActivityRecord{.*? (\S+)/(\S+?)\s`)
	matches = reApp.FindStringSubmatch(output)
	if len(matches) == 3 {
		windowInfo = types.WindowInfo{
			PackageName: matches[1],
			Activity:    matches[2],
		}
		return windowInfo, nil
	}

	// adb shell dumpsys activity activities | grep mResumedActivity
	output, err = dev.Device.RunShellCommand("dumpsys", "activity", "activities", "|", "grep", "mResumedActivity")
	if err != nil {
		return types.WindowInfo{}, errors.Wrap(err, "get current activity failed")
	}
	// mResumedActivity: ActivityRecord{2db504f u0 com.miui.home/.launcher.Launcher t2}
	reActivity := regexp.MustCompile(`mResumedActivity: ActivityRecord{.*? (\S+)/(\S+?)\s`)
	matches = reActivity.FindStringSubmatch(output)
	if len(matches) == 3 {
		windowInfo = types.WindowInfo{
			PackageName: matches[1],
			Activity:    matches[2],
		}
		return windowInfo, nil
	}

	return types.WindowInfo{}, errors.New("failed to extract current window")
}

func (dev *AndroidDevice) ListPackages() ([]string, error) {
	return dev.Device.ListPackages()
}

func (dev *AndroidDevice) GetPackageInfo(packageName string) (types.AppInfo, error) {
	appInfo := types.AppInfo{
		Name: packageName,
	}
	// get package version
	appVersion, err := dev.getPackageVersion(packageName)
	if err == nil {
		appInfo.AppBaseInfo.VersionName = appVersion
	} else {
		log.Warn().Msg("failed to get package version")
		return appInfo, errors.Wrap(code.DeviceAppNotInstalled, err.Error())
	}

	// get package path
	packagePath, err := dev.getPackagePath(packageName)
	if err == nil {
		appInfo.AppBaseInfo.AppPath = packagePath
	} else {
		log.Warn().Msg("failed to get package path")
		return appInfo, errors.Wrap(code.DeviceAppNotInstalled, err.Error())
	}

	// get package md5
	packageMD5, err := dev.getPackageMD5(packagePath)
	if err == nil {
		appInfo.AppBaseInfo.AppMD5 = packageMD5
	} else {
		log.Warn().Msg("failed to get package md5")
		return appInfo, errors.Wrap(code.DeviceAppNotInstalled, err.Error())
	}

	log.Info().Interface("appInfo", appInfo).Msg("get package info")
	return appInfo, nil
}

func (dev *AndroidDevice) ScreenShot() (*bytes.Buffer, error) {
	raw, err := dev.Device.ScreenCap()
	if err != nil {
		return nil, errors.Wrapf(code.DeviceScreenShotError,
			"adb screencap failed %v", err)
	}
	return bytes.NewBuffer(raw), nil
}

func (dev *AndroidDevice) GetAppInfo(packageName string) (app types.AppInfo, err error) {
	packageInfo, err := dev.RunShellCommand(
		"CLASSPATH=/data/local/tmp/evalite", "app_process", "/",
		"com.bytedance.iesqa.eval_process.PackageService", packageName, "2>/dev/null")
	if packageInfo == "" {
		return app, nil
	}
	if err != nil {
		return app, err
	}
	err = json.Unmarshal([]byte(strings.TrimSpace(packageInfo)), &app)
	if err != nil {
		log.Error().Err(err).Str("packageInfo", packageInfo)
	}
	return
}

func (dev *AndroidDevice) getPackageVersion(packageName string) (string, error) {
	output, err := dev.Device.RunShellCommand("dumpsys", "package", packageName, "|", "grep", "versionName")
	if err != nil {
		return "", errors.Wrap(err, "get package version failed")
	}
	appVersion := ""
	re := regexp.MustCompile(`versionName=(.+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		appVersion = matches[1]
		return appVersion, nil
	}
	return "", errors.New("failed to get package version")
}

func (dev *AndroidDevice) getPackagePath(packageName string) (string, error) {
	if packageName == "" {
		return "", errors.Wrap(code.InvalidParamError, "packageName is empty")
	}
	output, err := dev.Device.RunShellCommand("pm", "path", packageName)
	if err != nil {
		return "", errors.Wrap(err, "get package path failed")
	}
	re := regexp.MustCompile(`package:(.+)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("failed to get package path")
}

func (dev *AndroidDevice) getPackageMD5(packagePath string) (string, error) {
	output, err := dev.Device.RunShellCommand("md5sum", packagePath)
	if err != nil {
		return "", errors.Wrap(err, "get package md5 failed")
	}
	matches := strings.Split(output, " ")
	if len(matches) > 1 {
		return matches[0], nil
	}
	return "", errors.New("failed to get package md5")
}

type LineCallback func(string)

type AdbLogcat struct {
	serial string
	// logBuffer *bytes.Buffer
	errs     []error
	stopping chan struct{}
	done     chan struct{}
	cmd      *exec.Cmd
	callback LineCallback
	logs     []string
}

func NewAdbLogcatWithCallback(serial string, callback LineCallback) *AdbLogcat {
	return &AdbLogcat{
		serial: serial,
		// logBuffer: new(bytes.Buffer),
		stopping: make(chan struct{}),
		done:     make(chan struct{}),
		callback: callback,
		logs:     make([]string, 0),
	}
}

func NewAdbLogcat(serial string) *AdbLogcat {
	return &AdbLogcat{
		serial: serial,
		// logBuffer: new(bytes.Buffer),
		stopping: make(chan struct{}),
		done:     make(chan struct{}),
		logs:     make([]string, 0),
	}
}

// CatchLogcatContext starts logcat with timeout context
func (l *AdbLogcat) CatchLogcatContext(timeoutCtx context.Context) (err error) {
	if err = l.CatchLogcat(""); err != nil {
		return
	}
	go func() {
		select {
		case <-timeoutCtx.Done():
			_ = l.Stop()
		case <-l.stopping:
		}
	}()
	return
}

func (l *AdbLogcat) Stop() error {
	select {
	case <-l.stopping:
	default:
		close(l.stopping)
		<-l.done
		close(l.done)
	}
	return l.Errors()
}

func (l *AdbLogcat) Errors() (err error) {
	for _, e := range l.errs {
		if err != nil {
			err = fmt.Errorf("%v |[DeviceLogcatErr] %v", err, e)
		} else {
			err = fmt.Errorf("[DeviceLogcatErr] %v", e)
		}
	}
	return
}

func (l *AdbLogcat) CatchLogcat(filter string) (err error) {
	if l.cmd != nil {
		log.Warn().Msg("logcat already start")
		return nil
	}

	// FIXME: replace with gadb shell command
	// clear logcat
	if err = myexec.RunCommand("adb", "-s", l.serial, "shell", "logcat", "-c"); err != nil {
		return
	}
	args := []string{"-s", l.serial, "logcat", "--format", "time"}
	if filter != "" {
		args = append(args, "-s", filter)
	}
	// start logcat
	l.cmd = myexec.Command("adb", args...)
	// l.cmd.Stderr = l.logBuffer
	// l.cmd.Stdout = l.logBuffer
	reader, err := l.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = l.cmd.Start(); err != nil {
		return
	}
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			if l.callback != nil {
				l.callback(line) // Process each line with callback
			} else {
				l.logs = append(l.logs, line) // Store line if no callback
			}
		}
	}()
	go func() {
		<-l.stopping
		if e := reader.Close(); e != nil {
			log.Error().Err(e).Msg("close logcat reader failed")
		}
		if e := myexec.KillProcessesByGpid(l.cmd); e != nil {
			log.Error().Err(e).Msg("kill logcat process failed")
		}
		l.done <- struct{}{}
	}()

	return
}
