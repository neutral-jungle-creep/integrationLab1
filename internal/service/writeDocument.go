package service

import (
	"IntegrationLab1/configs"
	"IntegrationLab1/internal/domain"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"math"
	"math/rand"
	"strconv"
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
	fileName := getFileNameWithExtension(req.ClientFullName)
	clientAndProviderData := makeReplaceMap(req)

	doc, err := docx.Open(w.cfg.Paths.TemplatePath + req.TemplateFile)
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

func makeReplaceMap(req *domain.UserRequest) docx.PlaceholderMap {
	var now = time.Now()

	replMap := docx.PlaceholderMap{
		"doc-number": rand.Intn(999999) + 1,
		"day":        now.Format("02"),
		"month":      now.Format("January"),

		"client-full-name":         strings.TrimSpace(req.ClientFullName),
		"client-phone-number":      strings.TrimSpace(req.ClientPhoneNumber),
		"client-email":             strings.TrimSpace(req.ClientEmail),
		"client-company":           strings.TrimSpace(req.ClientCompany),
		"client-company-full-name": strings.TrimSpace(req.ClientCompanyFullName),
		"client-company-inn-kpp":   strings.TrimSpace(req.ClientCompanyInnKpp),
		"client-company-address":   strings.TrimSpace(req.ClientCompanyAddress),

		"provider-full-name":         strings.TrimSpace(req.ProviderFullName),
		"provider-phone-number":      strings.TrimSpace(req.ProviderPhoneNumber),
		"provider-email":             strings.TrimSpace(req.ProviderEmail),
		"provider-company":           strings.TrimSpace(req.ProviderCompany),
		"provider-company-full-name": strings.TrimSpace(req.ProviderCompanyFullName),
		"provider-company-inn-kpp":   strings.TrimSpace(req.ProviderCompanyInnKpp),
		"provider-company-address":   strings.TrimSpace(req.ProviderCompanyAddress),

		"delivery-address": strings.TrimSpace(req.DeliveryAddress),
		"total-price":      "",
	}

	addEmptyItemsToReplaceMap(replMap)

	if len(req.Items) > 0 {
		addItemsToMap(replMap, req.Items)
	}

	return replMap
}

func addEmptyItemsToReplaceMap(replMap docx.PlaceholderMap) {
	for i := 0; i < 15; i++ {
		index := strconv.Itoa(i)
		replMap["vendor-code-"+index] = ""
		replMap["item-name-"+index] = ""
		replMap["quantity-"+index] = ""
		replMap["price-"+index] = ""
		replMap["total-"+index] = ""
	}
}

func addItemsToMap(replMap docx.PlaceholderMap, items []domain.Items) {
	var totalPrice float64

	for i := 0; i < len(items); i++ {
		index := strconv.Itoa(i)
		total := math.Round(items[i].Price * float64(items[i].Quantity))
		totalPrice += total

		replMap["vendor-code-"+index] = items[i].VendorCode
		replMap["item-name-"+index] = items[i].Name
		replMap["quantity-"+index] = items[i].Quantity
		replMap["price-"+index] = fmt.Sprintf("%.2f %s", items[i].Price, "р.")
		replMap["total-"+index] = fmt.Sprintf("%.2f %s", total, "р.")
	}
}
