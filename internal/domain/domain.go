package domain

type UserRequest struct {
	ClientFirstName   string `json:"client-first-name" binding:"required" validate:"alpha,max=50"`
	ClientLastName    string `json:"client-last-name" binding:"required" validate:"alpha,max=50"`
	ClientMiddleName  string `json:"client-middle-name" validate:"alpha,max=50"`
	ClientCompany     string `json:"client-company" validate:"max=100"`
	ClientPhoneNumber string `json:"client-phone-number" validate:"e164"`
	ClientEmail       string `json:"client-email" validate:"email"`

	ProviderFirstName   string `json:"provider-first-name" validate:"alpha,max=50"`
	ProviderLastName    string `json:"provider-last-name" validate:"alpha,max=50"`
	ProviderMiddleName  string `json:"provider-middle-name" validate:"alpha,max=50"`
	ProviderCompany     string `json:"provider-company" validate:"max=100"`
	ProviderPhoneNumber string `json:"provider-phone-number" validate:"e164"`
	ProviderEmail       string `json:"provider-email" validate:"email"`

	Items []struct {
		VendorCode string  `json:"vendor-code" validate:"alphanum,max=30"`
		Price      float32 `json:"price" validate:""`
		Name       string  `json:"name" validate:"required"`
		Quantity   uint    `json:"quantity"`
	} `json:"items"`
}

// TODO: add required to all values of struct UserRequest
