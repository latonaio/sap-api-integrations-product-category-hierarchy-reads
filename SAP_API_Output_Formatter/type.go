package sap_api_output_formatter

type ProductCategoryHierarchy struct {
	ConnectionKey                string `json:"connection_key"`
	Result                       bool   `json:"result"`
	RedisKey                     string `json:"redis_key"`
	Filepath                     string `json:"filepath"`
	APISchema                    string `json:"api_schema"`
	ProductCategoryHierarchyCode string `json:"product_category_hierarchy_code"`
	Deleted                      bool   `json:"deleted"`
}

type ProductCategoryHierarchyCollection struct {
	ObjectID                        string `json:"ObjectID"`
	ETag                            string `json:"ETag"`
	RootProductCategoryInternalID   string `json:"RootProductCategoryInternalID"`
	ProductCategoryInternalID       string `json:"ProductCategoryInternalID"`
	ParentProductCategoryInternalID string `json:"ParentProductCategoryInternalID"`
	LanguageCode                    string `json:"LanguageCode"`
	LanguageCodeText                string `json:"LanguageCodeText"`
	ProductCategory                 string `json:"ProductCategory"`
	ProductAssignmentAllowed        bool   `json:"ProductAssignmentAllowed"`
}