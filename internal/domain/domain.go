package domain

type UserRequest struct {
	ClientFirstName   string `json:"clientFirstName" binding:"required" validate:"alphaunicode,max=50"`
	ClientLastName    string `json:"clientLastName" binding:"required" validate:"alphaunicode,max=50"`
	ClientMiddleName  string `json:"clientMiddleName" binding:"required" validate:"alphaunicode,max=50"`
	ClientCompany     string `json:"clientCompany" binding:"required" validate:"max=100"`
	ClientPhoneNumber string `json:"clientPhoneNumber" binding:"required" validate:"e164"`
	ClientEmail       string `json:"clientEmail" binding:"required" validate:"email"`

	ProviderFirstName   string `json:"providerFirstName" binding:"required" validate:"alphaunicode,max=50"`
	ProviderLastName    string `json:"providerLastName" binding:"required" validate:"alphaunicode,max=50"`
	ProviderMiddleName  string `json:"providerMiddleName" binding:"required" validate:"alphaunicode,max=50"`
	ProviderCompany     string `json:"providerCompany" binding:"required" validate:"max=100"`
	ProviderPhoneNumber string `json:"providerPhoneNumber" binding:"required" validate:"e164"`
	ProviderEmail       string `json:"providerEmail" binding:"required" validate:"email"`

	Items []struct {
		VendorCode string  `json:"vendorCode" binding:"required" validate:"alphanum,max=30"`
		Price      float32 `json:"price" binding:"required"`
		Name       string  `json:"name" binding:"required" validate:"required"`
		Quantity   uint    `json:"quantity" binding:"required"`
	} `json:"items"`
}

type Response struct {
	FileName string `json:"fileName"`
}
