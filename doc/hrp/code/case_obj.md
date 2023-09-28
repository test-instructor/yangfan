
执行测试用例的过程涉及两层循环。在外层循环中，系统会遍历加载的所有测试用例，依次对它们进行执行。在内层循环中，针对每个测试用例，系统会遍历其中的接口（或步骤），逐一执行这些接口。

这种双层循环的设计能够确保每个测试用例中的所有接口都得到执行，从而实现全面的自动化测试。外层循环保证了所有测试用例都被执行，而内层循环保证了每个测试用例中的每个接口都得到了执行，从而覆盖了不同的测试场景和情况。

这样的设计模式使得自动化测试框架能够高效、全面地执行测试，对于确保软件质量和功能稳定性具有重要意义。
## CaseRunner对象初始化
```go
for _, testcase := range testCases {
		// 每个测试用例都有自己的用例运行器，把用例转换成可迭代的用例
		caseRunner, err := r.NewCaseRunner(testcase)
		// 使用参数迭代器执行用例多次，每次执行有不同的参数
		for it := caseRunner.parametersIterator; it.HasNext(); {
		}
	}
```

获取测试用例后，我们通过调用`NewCaseRunner`函数创建一个执行用例对象。这一步的主要操作包括：
##3 CaseRunner 主要内容
1. 创建用例运行对象：生成一个用例运行对象，用于执行测试步骤并生成报告。
2. 初始化函数插件：初始化可能在测试用例执行过程中需要使用的函数插件，以便在需要时进行调用。
3. 解析配置：解析测试用例中的配置信息，确保在执行过程中能够正确应用配置。
4. 设置超时时间：根据配置中设置的超时时间，为测试用例设置适当的时间限制，防止执行超时。
5. 配置中设置函数插件：将函数插件的配置应用到当前测试用例的执行环境中，以确保测试步骤能够正确调用。

这个阶段可以被看作是每个测试用例的初始化部分，它确保了每个测试用例都在一个干净、适当的环境下执行。如果有多个测试用例需要执行，并且它们的函数插件路径相同，那么会共享同一个函数插件，从而节省资源开销。

在这个过程中，每个测试用例都会拥有独立的配置，通过值传递的方式进行设置。此外，通过`Export`参数，可以在不同的测试用例之间共享数据，实现信息的传递和共享。这个步骤的目标是确保每个测试用例都处于一个可控、独立的执行环境中，以获得准确的测试结果。

```go
// NewCaseRunner 创建一个新的用例运行器（CaseRunner）用于指定的测试用例（testcase）。
// 每个测试用例都有自己的用例运行器。
func (r *HRPRunner) NewCaseRunner(testcase *TestCase) (*CaseRunner, error) {
	// 创建一个新的用例运行器（CaseRunner）对象
	caseRunner := &CaseRunner{
		testCase:  testcase,    // 设置用例运行器的测试用例字段
		hrpRunner: r,           // 设置用例运行器的HRPRunner字段
		parser:    newParser(), // 创建并初始化一个新的解析器（parser）对象
	}

	// 初始化函数插件
	plugin, err := initPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)

	// 解析测试用例的配置
	if err := caseRunner.parseConfig(); err != nil {
		return nil, errors.Wrap(err, "parse testcase config failed")
	}

	// 设置测试用例的超时时间（timeout）（单位：秒）
	if testcase.Config.Timeout != 0 {
		timeout := time.Duration(testcase.Config.Timeout*1000) * time.Millisecond
		r.SetTimeout(timeout)
	}

	// 将插件信息加载到测试用例的配置中
	if plugin != nil {
	}

	return caseRunner, nil 
}

```
## CaseRunner 字段解析
`CaseRunner`是一个结构体，它包含了多个关键成员，用于管理一个测试用例的执行过程。以下是`CaseRunner`中各成员的作用：

1. `testCase`: 这是一个指向测试用例的引用，它包含了测试用例的各种信息，如配置、名称等。
2. `hrpRunner`: 这是一个指向`HRPRunner`对象的引用，它用于管理整个测试执行的运行时环境，包括HTTP请求、插件管理等。
3. `parser`: 这是一个指向字符串解析器（`Parser`）的引用，用于解析测试步骤中的字符串，包括变量和函数等。
4. `parsedConfig`: 这是一个指向测试用例配置（`TConfig`）的引用，它存储了解析后的测试用例配置信息，包括请求信息、验证规则等。
5. `parametersIterator`: 这是一个用于迭代测试用例参数的迭代器（`ParametersIterator`），它管理了测试用例执行时不同参数组合的迭代过程。
6. `rootDir`: 这是一个字符串，表示项目的根目录路径。它用于确定插件所在的目录等。

## 小结
总体而言，`CaseRunner`在测试用例的执行过程中，协调了不同成员的功能，包括管理测试用例配置、字符串解析、参数迭代、环境配置等。这有助于保持代码的组织性和可读性，从而实现有效的测试用例执行和管理。

```go
// CaseRunner 结构体用于执行单个测试用例的运行器
type CaseRunner struct {
	testCase           *TestCase
	hrpRunner          *HRPRunner
	parser             *Parser

	parsedConfig       *TConfig
	parametersIterator *ParametersIterator
	rootDir            string // 项目根目录
}

// parseConfig 方法解析测试用例的配置，将解析结果存储到 parsedConfig 中
func (r *CaseRunner) parseConfig() error {
}

// NewSession 方法用于创建一个新的会话运行器（SessionRunner）
func (r *CaseRunner) NewSession() *SessionRunner {
	return sessionRunner
}
```

