package interfacecase

type getOperationId struct {
	CreatedByID uint `json:"-"`
	UpdateByID  uint `json:"-"`
	DeleteByID  uint `json:"-"`
}
