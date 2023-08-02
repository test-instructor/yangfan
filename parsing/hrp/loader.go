package hrp

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func LoadTestCases(iTestCases ...ITestCase) ([]*TestCase, error) {
	testCases := make([]*TestCase, 0)
	// 遍历传入的ITestCase接口切片，处理每个测试用例。
	for _, iTestCase := range iTestCases {
		// 如果当前ITestCase是*TestCase类型，则直接转换为TestCase并添加到testCases切片中
		if _, ok := iTestCase.(*TestCase); ok {
			testcase, err := iTestCase.ToTestCase()
			if err != nil {
				log.Error().Err(err).Msg("failed to convert ITestCase interface to TestCase struct")
				return nil, err
			}
			testCases = append(testCases, testcase)
			continue
		}

		// 否则，iTestCase应该是一个TestCasePath，表示文件路径或文件夹路径
		tcPath, ok := iTestCase.(*TestCasePath)
		if !ok {
			return nil, errors.New("invalid iTestCase type")
		}
		// 获取测试用例路径
		casePath := tcPath.GetPath()
		// 使用fs.WalkDir函数遍历目录结构，处理每个测试用例文件
		err := fs.WalkDir(os.DirFS(casePath), ".", func(path string, dir fs.DirEntry, e error) error {
			if dir == nil {
				// casePath是文件而不是目录
				path = casePath
			} else if dir.IsDir() && path != "." && strings.HasPrefix(path, ".") {
				// 跳过隐藏文件夹
				return fs.SkipDir
			} else {
				// casePath是目录
				path = filepath.Join(casePath, path)
			}

			// 忽略非测试用例文件
			ext := filepath.Ext(path)
			if ext != ".yml" && ext != ".yaml" && ext != ".json" {
				return nil
			}

			// 获取TestCasePath并转换为TestCase结构，然后添加到testCases切片中
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

	log.Info().Int("count", len(testCases)).Msg("load testcases successfully")
	return testCases, nil
}
