package uixt

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/test-instructor/yangfan/httprunner/uixt/option"
	"github.com/test-instructor/yangfan/httprunner/uixt/types"
)

// Ensure CompositeDriver implements IDriver and SIMSupport interfaces
var (
	_ IDriver    = (*CompositeDriver)(nil)
	_ SIMSupport = (*CompositeDriver)(nil)
)

// CompositeDriver is a driver that can dynamically switch between ADB and UIA2 drivers
// for each operation based on ActionOption.DriverType
type CompositeDriver struct {
	adbDriver   *ADBDriver
	uia2Driver  *UIA2Driver
	device      *AndroidDevice
	defaultType option.DriverType
}

// NewCompositeDriver creates a new CompositeDriver with both ADB and UIA2 drivers initialized
func NewCompositeDriver(device *AndroidDevice) (*CompositeDriver, error) {
	log.Info().Interface("device", device).Msg("init android composite driver")

	// Create ADB driver
	adbDriver, err := NewADBDriver(device)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init ADB driver for composite")
	}

	// Create UIA2 driver (which internally creates its own ADB driver, but we keep separate instances)
	uia2Driver, err := NewUIA2Driver(device)
	if err != nil {
		// UIA2 might fail if server is not installed, fall back to ADB only mode
		log.Warn().Err(err).Msg("failed to init UIA2 driver, composite driver will use ADB only")
		return &CompositeDriver{
			adbDriver:   adbDriver,
			uia2Driver:  nil,
			device:      device,
			defaultType: option.DriverTypeADB,
		}, nil
	}

	// Determine default driver type based on device options
	defaultType := option.DriverTypeUIA2
	if !device.Options.UIA2 {
		defaultType = option.DriverTypeADB
	}

	return &CompositeDriver{
		adbDriver:   adbDriver,
		uia2Driver:  uia2Driver,
		device:      device,
		defaultType: defaultType,
	}, nil
}

// SetDefaultDriverType sets the default driver type for operations without explicit driver type
func (cd *CompositeDriver) SetDefaultDriverType(dt option.DriverType) {
	cd.defaultType = dt
}

// GetDefaultDriverType returns the current default driver type
func (cd *CompositeDriver) GetDefaultDriverType() option.DriverType {
	return cd.defaultType
}

// resolveDriverType determines which driver to use based on options
func (cd *CompositeDriver) resolveDriverType(opts ...option.ActionOption) option.DriverType {
	actionOpts := option.NewActionOptions(opts...)
	if actionOpts.DriverType != "" && actionOpts.DriverType != option.DriverTypeAuto {
		return actionOpts.DriverType
	}
	return cd.defaultType
}

// getADBDriver returns the ADB driver, always available
func (cd *CompositeDriver) getADBDriver() *ADBDriver {
	return cd.adbDriver
}

// getUIA2Driver returns the UIA2 driver, may be nil if not available
func (cd *CompositeDriver) getUIA2Driver() (*UIA2Driver, error) {
	if cd.uia2Driver == nil {
		return nil, errors.New("UIA2 driver is not available")
	}
	return cd.uia2Driver, nil
}

// ============== IDriver Interface Implementation ==============

func (cd *CompositeDriver) GetDevice() IDevice {
	return cd.device
}

func (cd *CompositeDriver) Setup() error {
	if err := cd.adbDriver.Setup(); err != nil {
		return err
	}
	if cd.uia2Driver != nil {
		return cd.uia2Driver.Setup()
	}
	return nil
}

func (cd *CompositeDriver) TearDown() error {
	if cd.uia2Driver != nil {
		if err := cd.uia2Driver.TearDown(); err != nil {
			log.Warn().Err(err).Msg("failed to teardown UIA2 driver")
		}
	}
	return cd.adbDriver.TearDown()
}

func (cd *CompositeDriver) InitSession(capabilities option.Capabilities) error {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.InitSession(capabilities)
	}
	return cd.adbDriver.InitSession(capabilities)
}

func (cd *CompositeDriver) GetSession() *DriverSession {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.GetSession()
	}
	return cd.adbDriver.GetSession()
}

func (cd *CompositeDriver) DeleteSession() error {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.DeleteSession()
	}
	return cd.adbDriver.DeleteSession()
}

// ============== Device Info Methods ==============

func (cd *CompositeDriver) Status() (types.DeviceStatus, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.Status()
	}
	return cd.adbDriver.Status()
}

func (cd *CompositeDriver) DeviceInfo() (types.DeviceInfo, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.DeviceInfo()
	}
	return cd.adbDriver.DeviceInfo()
}

func (cd *CompositeDriver) BatteryInfo() (types.BatteryInfo, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.BatteryInfo()
	}
	return cd.adbDriver.BatteryInfo()
}

func (cd *CompositeDriver) ForegroundInfo() (types.AppInfo, error) {
	// NOTE: ForegroundInfo usage has been disabled per requirement, so the
	// CompositeDriver no longer forwards this call to underlying drivers.
	// The original implementation is kept below for reference:
	//
	// if cd.uia2Driver != nil {
	// 	return cd.uia2Driver.ForegroundInfo()
	// }
	// return cd.adbDriver.ForegroundInfo()
	return types.AppInfo{}, nil
}

func (cd *CompositeDriver) WindowSize() (types.Size, error) {
	// WindowSize can use either driver, prefer UIA2 for accuracy
	if cd.uia2Driver != nil {
		return cd.uia2Driver.WindowSize()
	}
	return cd.adbDriver.WindowSize()
}

func (cd *CompositeDriver) ScreenShot(opts ...option.ActionOption) (*bytes.Buffer, error) {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.ScreenShot(opts...)
	}
	return cd.adbDriver.ScreenShot(opts...)
}

func (cd *CompositeDriver) ScreenRecord(opts ...option.ActionOption) (string, error) {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.ScreenRecord(opts...)
	}
	return cd.adbDriver.ScreenRecord(opts...)
}

func (cd *CompositeDriver) Source(srcOpt ...option.SourceOption) (string, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.Source(srcOpt...)
	}
	return cd.adbDriver.Source(srcOpt...)
}

func (cd *CompositeDriver) Orientation() (types.Orientation, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.Orientation()
	}
	return cd.adbDriver.Orientation()
}

func (cd *CompositeDriver) Rotation() (types.Rotation, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.Rotation()
	}
	return cd.adbDriver.Rotation()
}

// ============== Config Methods ==============

func (cd *CompositeDriver) SetRotation(rotation types.Rotation) error {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.SetRotation(rotation)
	}
	return cd.adbDriver.SetRotation(rotation)
}

func (cd *CompositeDriver) SetIme(ime string) error {
	// SetIme uses ADB command, so always use ADB driver
	return cd.adbDriver.SetIme(ime)
}

// ============== Basic Actions ==============

func (cd *CompositeDriver) Home() error {
	// Home button can use either, prefer ADB for reliability
	return cd.adbDriver.Home()
}

func (cd *CompositeDriver) Unlock() error {
	return cd.adbDriver.Unlock()
}

func (cd *CompositeDriver) Back() error {
	return cd.adbDriver.Back()
}

func (cd *CompositeDriver) PressButton(button types.DeviceButton) error {
	return cd.adbDriver.PressButton(button)
}

// ============== Tap Actions (Support Driver Switching) ==============

func (cd *CompositeDriver) TapXY(x, y float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	log.Debug().Str("driver", string(driverType)).Float64("x", x).Float64("y", y).Msg("CompositeDriver.TapXY")
	if driverType == option.DriverTypeUIA2 {
		if cd.uia2Driver != nil {
			return cd.uia2Driver.TapXY(x, y, opts...)
		}
		log.Warn().Float64("x", x).Float64("y", y).Msg("UIA2 driver requested but not available; fallback to ADB")
	}
	return cd.adbDriver.TapXY(x, y, opts...)
}

func (cd *CompositeDriver) TapAbsXY(x, y float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	log.Debug().Str("driver", string(driverType)).Float64("x", x).Float64("y", y).Msg("CompositeDriver.TapAbsXY")
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.TapAbsXY(x, y, opts...)
	}
	return cd.adbDriver.TapAbsXY(x, y, opts...)
}

func (cd *CompositeDriver) TapBySelector(selector string, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	log.Debug().Str("driver", string(driverType)).Str("selector", selector).Msg("CompositeDriver.TapBySelector")
	if driverType == option.DriverTypeUIA2 {
		if cd.uia2Driver != nil {
			return cd.uia2Driver.TapBySelector(selector, opts...)
		}
		log.Warn().Str("selector", selector).Msg("UIA2 driver requested but not available; fallback to ADB")
	}
	return cd.adbDriver.TapBySelector(selector, opts...)
}

func (cd *CompositeDriver) DoubleTap(x, y float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.DoubleTap(x, y, opts...)
	}
	return cd.adbDriver.DoubleTap(x, y, opts...)
}

func (cd *CompositeDriver) TouchAndHold(x, y float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.TouchAndHold(x, y, opts...)
	}
	return cd.adbDriver.TouchAndHold(x, y, opts...)
}

func (cd *CompositeDriver) HoverBySelector(selector string, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.HoverBySelector(selector, opts...)
	}
	return cd.adbDriver.HoverBySelector(selector, opts...)
}

// ============== Secondary Click ==============

func (cd *CompositeDriver) SecondaryClick(x, y float64) error {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.SecondaryClick(x, y)
	}
	return cd.adbDriver.SecondaryClick(x, y)
}

func (cd *CompositeDriver) SecondaryClickBySelector(selector string, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.SecondaryClickBySelector(selector, opts...)
	}
	return cd.adbDriver.SecondaryClickBySelector(selector, opts...)
}

// ============== Swipe Actions (Support Driver Switching) ==============

func (cd *CompositeDriver) Swipe(fromX, fromY, toX, toY float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	log.Debug().Str("driver", string(driverType)).Msg("CompositeDriver.Swipe")
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.Swipe(fromX, fromY, toX, toY, opts...)
	}
	return cd.adbDriver.Swipe(fromX, fromY, toX, toY, opts...)
}

func (cd *CompositeDriver) Drag(fromX, fromY, toX, toY float64, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.Drag(fromX, fromY, toX, toY, opts...)
	}
	return cd.adbDriver.Drag(fromX, fromY, toX, toY, opts...)
}

// ============== Input Actions (Support Driver Switching) ==============

func (cd *CompositeDriver) Input(text string, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	log.Debug().Str("driver", string(driverType)).Str("text", text).Msg("CompositeDriver.Input")
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.Input(text, opts...)
	}
	return cd.adbDriver.Input(text, opts...)
}

func (cd *CompositeDriver) Backspace(count int, opts ...option.ActionOption) error {
	driverType := cd.resolveDriverType(opts...)
	if driverType == option.DriverTypeUIA2 && cd.uia2Driver != nil {
		return cd.uia2Driver.Backspace(count, opts...)
	}
	return cd.adbDriver.Backspace(count, opts...)
}

// ============== App Related ==============

func (cd *CompositeDriver) AppLaunch(packageName string) error {
	// App launch uses ADB command
	return cd.adbDriver.AppLaunch(packageName)
}

func (cd *CompositeDriver) AppTerminate(packageName string) (bool, error) {
	return cd.adbDriver.AppTerminate(packageName)
}

func (cd *CompositeDriver) AppClear(packageName string) error {
	return cd.adbDriver.AppClear(packageName)
}

// ============== Image/File Related ==============

func (cd *CompositeDriver) PushImage(localPath string) error {
	return cd.adbDriver.PushImage(localPath)
}

func (cd *CompositeDriver) PullImages(localDir string) error {
	return cd.adbDriver.PullImages(localDir)
}

func (cd *CompositeDriver) ClearImages() error {
	return cd.adbDriver.ClearImages()
}

func (cd *CompositeDriver) PushFile(localPath string, remoteDir string) error {
	return cd.adbDriver.PushFile(localPath, remoteDir)
}

func (cd *CompositeDriver) PullFiles(localDir string, remoteDirs ...string) error {
	return cd.adbDriver.PullFiles(localDir, remoteDirs...)
}

func (cd *CompositeDriver) ClearFiles(paths ...string) error {
	return cd.adbDriver.ClearFiles(paths...)
}

// ============== Log Capture ==============

func (cd *CompositeDriver) StartCaptureLog(identifier ...string) error {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.StartCaptureLog(identifier...)
	}
	return cd.adbDriver.StartCaptureLog(identifier...)
}

func (cd *CompositeDriver) StopCaptureLog() (interface{}, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.StopCaptureLog()
	}
	return cd.adbDriver.StopCaptureLog()
}

// ============== Clipboard ==============

func (cd *CompositeDriver) GetPasteboard() (string, error) {
	if cd.uia2Driver != nil {
		return cd.uia2Driver.GetPasteboard()
	}
	return cd.adbDriver.GetPasteboard()
}

// ============== SIMSupport Interface Implementation ==============

func (cd *CompositeDriver) SIMClickAtPoint(x, y float64, opts ...option.ActionOption) error {
	if cd.uia2Driver == nil {
		return errors.New("SIMClickAtPoint requires UIA2 driver")
	}
	return cd.uia2Driver.SIMClickAtPoint(x, y, opts...)
}

func (cd *CompositeDriver) SIMSwipeWithDirection(direction string, fromX, fromY, simMinDistance, simMaxDistance float64, opts ...option.ActionOption) error {
	if cd.uia2Driver == nil {
		return errors.New("SIMSwipeWithDirection requires UIA2 driver")
	}
	return cd.uia2Driver.SIMSwipeWithDirection(direction, fromX, fromY, simMinDistance, simMaxDistance, opts...)
}

func (cd *CompositeDriver) SIMSwipeInArea(direction string, simAreaStartX, simAreaStartY, simAreaEndX, simAreaEndY, simMinDistance, simMaxDistance float64, opts ...option.ActionOption) error {
	if cd.uia2Driver == nil {
		return errors.New("SIMSwipeInArea requires UIA2 driver")
	}
	return cd.uia2Driver.SIMSwipeInArea(direction, simAreaStartX, simAreaStartY, simAreaEndX, simAreaEndY, simMinDistance, simMaxDistance, opts...)
}

func (cd *CompositeDriver) SIMSwipeFromPointToPoint(fromX, fromY, toX, toY float64, opts ...option.ActionOption) error {
	if cd.uia2Driver == nil {
		return errors.New("SIMSwipeFromPointToPoint requires UIA2 driver")
	}
	return cd.uia2Driver.SIMSwipeFromPointToPoint(fromX, fromY, toX, toY, opts...)
}

func (cd *CompositeDriver) SIMInput(text string, opts ...option.ActionOption) error {
	if cd.uia2Driver == nil {
		return errors.New("SIMInput requires UIA2 driver")
	}
	return cd.uia2Driver.SIMInput(text, opts...)
}

// ============== Direct Driver Access (for advanced use cases) ==============

// GetADBDriver returns the underlying ADB driver for direct access
func (cd *CompositeDriver) GetADBDriver() *ADBDriver {
	return cd.adbDriver
}

// GetUIA2Driver returns the underlying UIA2 driver for direct access
// Returns nil if UIA2 is not available
func (cd *CompositeDriver) GetUIA2Driver() *UIA2Driver {
	return cd.uia2Driver
}

// IsUIA2Available returns true if UIA2 driver is available
func (cd *CompositeDriver) IsUIA2Available() bool {
	return cd.uia2Driver != nil
}
