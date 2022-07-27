package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-category-hierarchy-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToProductCategoryHierarchyCollection(raw []byte, l *logger.Logger) ([]ProductCategoryHierarchyCollection, error) {
	pm := &responses.ProductCategoryHierarchyCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductCategoryHierarchyCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productCategoryHierarchyCollection := make([]ProductCategoryHierarchyCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productCategoryHierarchyCollection = append(productCategoryHierarchyCollection, ProductCategoryHierarchyCollection{
			ObjectID:                        data.ObjectID,
			ETag:                            data.ETag,
			RootProductCategoryInternalID:   data.RootProductCategoryInternalID,
			ProductCategoryInternalID:       data.ProductCategoryInternalID,
			ParentProductCategoryInternalID: data.ParentProductCategoryInternalID,
			LanguageCode:                    data.LanguageCode,
			LanguageCodeText:                data.LanguageCodeText,
			ProductCategory:                 data.ProductCategory,
			ProductAssignmentAllowed:        data.ProductAssignmentAllowed,
		})
	}

	return productCategoryHierarchyCollection, nil
}
