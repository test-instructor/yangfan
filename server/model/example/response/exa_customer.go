package response

import "github.com/test-instructor/yangfan/server/v2/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
