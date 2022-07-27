package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-product-category-hierarchy-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetProductCategoryHierarchy(productCategoryInternalID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ProductCategoryHierarchyCollection":
			func() {
				c.ProductCategoryHierarchyCollection(productCategoryInternalID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ProductCategoryHierarchyCollection(productCategoryInternalID string) {
	data, err := c.callProductCategoryHierarchySrvAPIRequirementProductCategoryHierarchyCollection("ProductCategoryHierarchyCollection", productCategoryInternalID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callProductCategoryHierarchySrvAPIRequirementProductCategoryHierarchyCollection(api, productCategoryInternalID string) ([]sap_api_output_formatter.ProductCategoryHierarchyCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithProductCategoryHierarchyCollection(req, productCategoryInternalID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProductCategoryHierarchyCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithProductCategoryHierarchyCollection(req *http.Request, productCategoryInternalID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ProductCategoryInternalID eq '%s'", productCategoryInternalID))
	req.URL.RawQuery = params.Encode()
}
