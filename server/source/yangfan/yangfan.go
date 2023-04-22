package yangfan

func Init() {
	go RegisterApis()
	go RegisterMenus()
	go PyPkg()
}
