package initialize

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/test-instructor/yangfan/server/v2/global"
	"go.uber.org/zap"
)

func InstallPip(python3 string) error {
	global.GVA_LOG.Info("check pip installed", zap.String("python3", python3))
	if err := runCommand(python3, "-m", "pip", "--version"); err == nil {
		global.GVA_LOG.Info("pip already installed, upgrade it", zap.String("python3", python3))
		return upgradePip(python3)
	}

	global.GVA_LOG.Info("pip not found, try ensurepip", zap.String("python3", python3))
	if err := runCommand(python3, "-m", "ensurepip", "--upgrade"); err == nil {
		if err := runCommand(python3, "-m", "pip", "--version"); err == nil {
			return upgradePip(python3)
		}
	}

	global.GVA_LOG.Info("ensurepip failed or pip still missing, try get-pip.py", zap.String("python3", python3))
	if err := installPipByGetPip(python3); err != nil {
		return err
	}

	return upgradePip(python3)
}

func upgradePip(python3 string) error {
	global.GVA_LOG.Info("upgrade pip", zap.String("python3", python3))
	if err := runCommand(python3, "-m", "pip", "install", "--upgrade", "pip"); err != nil {
		return fmt.Errorf("upgrade pip failed: %w", err)
	}
	return nil
}

func installPipByGetPip(python3 string) error {
	getPipURL := "https://bootstrap.pypa.io/get-pip.py"
	if customURL := os.Getenv("GET_PIP_URL"); customURL != "" {
		getPipURL = customURL
		global.GVA_LOG.Info("use custom get-pip url", zap.String("url", getPipURL))
	}

	pythonScript := fmt.Sprintf(`
import sys, ssl, urllib.request

url = %q

def exec_bytes(content):
    if isinstance(content, bytes):
        content = content.decode("utf-8", errors="ignore")
    g = {"__name__": "__main__"}
    exec(content, g, g)

try:
    if url.startswith("http://") or url.startswith("https://"):
        ssl_context = ssl.create_default_context()
        ssl_context.check_hostname = False
        ssl_context.verify_mode = ssl.CERT_NONE
        with urllib.request.urlopen(url, context=ssl_context) as response:
            exec_bytes(response.read())
    else:
        path = url
        if url.startswith("file://"):
            path = url[len("file://"):]
        with open(path, "rb") as f:
            exec_bytes(f.read())
    sys.stdout.write("pip bootstrap success\n")
except Exception as e:
    sys.stderr.write("pip bootstrap failed: %%s\n" %% (str(e),))
    sys.exit(1)
`, getPipURL)

	global.GVA_LOG.Info("bootstrap pip by get-pip.py", zap.String("url", getPipURL))
	stdout, stderr, err := runCommandWithOutput(python3, "-c", pythonScript)
	if err != nil {
		global.GVA_LOG.Error(
			"bootstrap pip by get-pip.py failed",
			zap.String("stdout", stdout),
			zap.String("stderr", stderr),
			zap.Error(err),
		)
		return fmt.Errorf("bootstrap pip failed: %w", err)
	}

	if err := runCommand(python3, "-m", "pip", "--version"); err != nil {
		return fmt.Errorf("pip installed but verification failed: %w", err)
	}
	return nil
}

func runCommand(name string, args ...string) error {
	_, _, err := runCommandWithOutput(name, args...)
	return err
}

func runCommandWithOutput(name string, args ...string) (string, string, error) {
	cmd := exec.Command(name, args...)
	cmd.Env = append(os.Environ(), "PIP_DISABLE_PIP_VERSION_CHECK=1")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = strings.TrimSpace(stdout.String())
		}
		if msg == "" {
			return stdout.String(), stderr.String(), err
		}
		return stdout.String(), stderr.String(), fmt.Errorf("%s: %w", msg, err)
	}

	return stdout.String(), stderr.String(), nil
}
