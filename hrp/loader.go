package hrp

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
)

func LoadTestCases(iTestCases ...ITestCase) ([]*TestCase, error) {
	testCases := make([]*TestCase, 0)

	for _, iTestCase := range iTestCases {
		if _, ok := iTestCase.(*TestCase); ok {
			testcase, err := iTestCase.ToTestCase()
			if err != nil {
				global.GVA_LOG.Error("failed to convert ITestCase interface to TestCase struct", zap.Error(err))
				return nil, err
			}
			testCases = append(testCases, testcase)
			continue
		}

		// iTestCase should be a TestCasePath, file path or folder path
		tcPath, ok := iTestCase.(*TestCasePath)
		if !ok {
			return nil, errors.New("invalid iTestCase type")
		}

		casePath := tcPath.GetPath()
		err := fs.WalkDir(os.DirFS(casePath), ".", func(path string, dir fs.DirEntry, e error) error {
			if dir == nil {
				// casePath is a file other than a dir
				path = casePath
			} else if dir.IsDir() && path != "." && strings.HasPrefix(path, ".") {
				// skip hidden folders
				return fs.SkipDir
			} else {
				// casePath is a dir
				path = filepath.Join(casePath, path)
			}

			// ignore non-testcase files
			ext := filepath.Ext(path)
			if ext != ".yml" && ext != ".yaml" && ext != ".json" {
				return nil
			}

			// filtered testcases
			testCasePath := TestCasePath(path)
			tc, err := testCasePath.ToTestCase()
			if err != nil {
				return nil
			}
			testCases = append(testCases, tc)
			return nil
		})
		if err != nil {
			return nil, errors.Wrap(err, "read dir failed")
		}
	}

	global.GVA_LOG.Info("load testcases successfully", zap.Int("count", len(testCases)))
	return testCases, nil
}
