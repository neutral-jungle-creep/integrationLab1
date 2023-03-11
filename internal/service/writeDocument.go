package service

import (
	"IntegrationLab1/configs"
	"IntegrationLab1/internal/domain"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"strings"
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
	var templateFile = req.TemplateFile

	fileName := getFileNameWithExtension(req.ClientFullName)

	clientAndProviderData := docx.PlaceholderMap{
		"client-full-name":         req.ClientFullName,
		"client-phone-number":      req.ClientPhoneNumber,
		"client-email":             req.ClientEmail,
		"client-company":           req.ClientCompany,
		"client-company-full-name": req.ClientCompanyFullName,
		"client-company-inn-kpp":   req.ClientCompanyInnKpp,
		"client-company-address":   req.ClientCompanyAddress,

		"provider-full-name":         req.ProviderFullName,
		"provider-phone-number":      req.ProviderPhoneNumber,
		"provider-email":             req.ProviderEmail,
		"provider-company":           req.ProviderCompany,
		"provider-company-full-name": req.ProviderCompanyFullName,
		"provider-company-inn-kpp":   req.ProviderCompanyInnKpp,
		"provider-company-address":   req.ProviderCompanyAddress,

		//"delivery-address": req.DeliveryAddress,
	}

	// TODO: add replace map for items in document. Parse array of structs "Item"

	doc, err := docx.Open(w.cfg.Paths.TemplatePath + templateFile)
	if err != nil {
		return "", fmt.Errorf("open template file error: %s", err.Error())
	}
	defer doc.Close()

	if err = doc.ReplaceAll(clientAndProviderData); err != nil {
		return "", fmt.Errorf("replace all placeholders int template file error: %s", err.Error())
	}

	if err = doc.WriteToFile(w.cfg.Paths.OutPath + fileName); err != nil {
		return "", fmt.Errorf("write document to file error: %s", err.Error())
	}

	return fileName, nil
}

func getFileNameWithExtension(clientFullName string) string {
	var formatName = strings.Join(strings.Split(clientFullName, " "), "_")
	var formatData = time.Now().Format("02_01_2006_15_04_05")
	return formatName + "_" + formatData + ".docx"
}
