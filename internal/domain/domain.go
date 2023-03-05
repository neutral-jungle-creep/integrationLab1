package domain

type UserRequest struct {
	ClientFirstName   string `json:"clientFirstName" binding:"required" validate:"alpha,max=50"`
	ClientLastName    string `json:"clientLastName" binding:"required" validate:"alpha,max=50"`
	ClientMiddleName  string `json:"clientMiddleName" validate:"alpha,max=50"`
	ClientCompany     string `json:"clientCompany" validate:"max=100"`
	ClientPhoneNumber string `json:"clientPhoneNumber" validate:"e164"`
	ClientEmail       string `json:"clientEmail" validate:"email"`

	ProviderFirstName   string `json:"providerFirstName" validate:"alpha,max=50"`
	ProviderLastName    string `json:"providerLastName" validate:"alpha,max=50"`
	ProviderMiddleName  string `json:"providerMiddleName" validate:"alpha,max=50"`
	ProviderCompany     string `json:"providerCompany" validate:"max=100"`
	ProviderPhoneNumber string `json:"providerPhoneNumber" validate:"e164"`
	ProviderEmail       string `json:"providerEmail" validate:"email"`

	Items []struct {
		VendorCode string  `json:"vendorCode" validate:"alphanum,max=30"`
		Price      float32 `json:"price" validate:""`
		Name       string  `json:"name" validate:"required"`
		Quantity   uint    `json:"quantity"`
	} `json:"items"`
}

// TODO: add required to all values of struct UserRequest

type Response struct {
	FileName string `json:"fileName"`
}
