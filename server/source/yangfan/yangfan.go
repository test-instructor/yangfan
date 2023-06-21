package yangfan

func Init() {
	go RegisterApis()
	go RegisterMenus()
	go InitPythonPackage(true)
}
