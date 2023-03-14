package domain

type UserRequest struct {
	TemplateFile string `json:"templateFile" binding:"required"`

	ClientFullName        string `json:"clientFullName" binding:"required" validate:"max=250"`
	ClientPhoneNumber     string `json:"clientPhoneNumber" binding:"required" validate:"e164"`
	ClientEmail           string `json:"clientEmail" binding:"required" validate:"email"`
	ClientCompany         string `json:"clientCompany" binding:"required" validate:"max=250"`
	ClientCompanyFullName string `json:"clientCompanyFullName" binding:"required" validate:"max=250"`
	ClientCompanyInnKpp   string `json:"clientCompanyInnKpp" binding:"required" validate:"min=22,max=24"`
	ClientCompanyAddress  string `json:"clientCompanyAddress" binding:"required" validate:"max=250"`

	ProviderFullName        string `json:"providerFullName" binding:"required" validate:"max=50"`
	ProviderPhoneNumber     string `json:"providerPhoneNumber" binding:"required" validate:"e164"`
	ProviderEmail           string `json:"providerEmail" binding:"required" validate:"email"`
	ProviderCompany         string `json:"providerCompany" binding:"required" validate:"max=250"`
	ProviderCompanyFullName string `json:"providerCompanyFullName" binding:"required" validate:"max=250"`
	ProviderCompanyInnKpp   string `json:"providerCompanyInnKpp" binding:"required" validate:"min=22,max=24"`
	ProviderCompanyAddress  string `json:"providerCompanyAddress" binding:"required" validate:"max=250"`

	DeliveryAddress string `json:"deliveryAddress" binding:"required" validate:"max=250"`

	Items []Items `json:"items"`
}

type Items struct {
	VendorCode string  `json:"vendorCode" binding:"required" validate:"alphanum,max=30"`
	Name       string  `json:"itemName" binding:"required" validate:"required"`
	Quantity   uint    `json:"quantity" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}

type Response struct {
	FileName string `json:"fileName"`
}
