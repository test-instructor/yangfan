package system

type RouterGroup struct {
	ApiRouter
	JwtRouter
	SysRouter
	BaseRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AutoCodeRouter
	AuthorityRouter
	DictionaryRouter
	OperationRecordRouter
	DictionaryDetailRouter
	AuthorityBtnRouter
	ProjectRouter
}
