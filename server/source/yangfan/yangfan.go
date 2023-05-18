package yangfan

func Init() {
	go RegisterApis()
	go RegisterMenus()
	go InitPythonPackage(false)
}
