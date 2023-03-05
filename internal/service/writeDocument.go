package service

import (
	"IntegrationLab1/configs"
	"IntegrationLab1/internal/domain"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"time"
)

type DocumentWriter interface {
	DocumentWrite(req *domain.UserRequest) (string, error)
}

type writeDocService struct {
	cfg *configs.Config
}

func NewWriteDocService(cfg *configs.Config) *writeDocService {
	return &writeDocService{
		cfg: cfg,
	}
}

func (w *writeDocService) DocumentWrite(req *domain.UserRequest) (string, error) {
	fileName := req.ClientMiddleName + "_" + time.Now().Format("02-01-2006_15_04_05")

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
		return "", fmt.Errorf("open template file error: %s", err.Error())
	}

	if err = doc.ReplaceAll(clientAndProviderData); err != nil {
		return "", fmt.Errorf("replace all placeholders int template file error: %s", err.Error())
	}

	if err = doc.WriteToFile(w.cfg.Paths.OutPath + fileName + ".docx"); err != nil {
		return "", fmt.Errorf("write document to file error: %s", err.Error())
	}

	return fileName, nil
}
