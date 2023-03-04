package service

import (
	"IntegrationLab1/configs"
	"IntegrationLab1/internal/domain"
	"fmt"
	"github.com/lukasjarosch/go-docx"
)

type DocumentWriter interface {
	documentWrite(req *domain.UserRequest) error
}

type writeDocService struct {
	cfg *configs.Config
}

func NewWriteDocService(cfg *configs.Config) *writeDocService {
	return &writeDocService{
		cfg: cfg,
	}
}

func (w *writeDocService) documentWrite(req *domain.UserRequest) error {
	clientAndProviderData := docx.PlaceholderMap{
		"client-first-name":   req.ClientFirstName,
		"client-last-name":    req.ClientLastName,
		"client-middle-name":  req.ClientMiddleName,
		"client-company":      req.ClientCompany,
		"client-phone-number": req.ClientPhoneNumber,
		"client-email":        req.ClientEmail,

		"provider-first-name":   req.ProviderFirstName,
		"provider-last-name":    req.ProviderLastName,
		"provider-middle-name":  req.ProviderMiddleName,
		"provider-company":      req.ProviderCompany,
		"provider-phone-number": req.ProviderPhoneNumber,
		"provider-email":        req.ProviderEmail,
	}

	// TODO: add replace map for items in document. Parse array of structs "Item"

	doc, err := docx.Open(w.cfg.Paths.TemplateFile)
	if err != nil {
		return fmt.Errorf("open template file error: %s", err.Error())
	}

	err = doc.ReplaceAll(clientAndProviderData)
	if err != nil {
		return fmt.Errorf("replace all placeholders int template file error: %s", err.Error())
	}

	err = doc.WriteToFile(w.cfg.Paths.OutPath + "test.docx")
	if err != nil {
		return fmt.Errorf("write document to file error: %s", err.Error())
	}

	return nil
}
