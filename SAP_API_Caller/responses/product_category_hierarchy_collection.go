package responses

type ProductCategoryHierarchyCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                        string `json:"ObjectID"`
			ETag                            string `json:"ETag"`
			RootProductCategoryInternalID   string `json:"RootProductCategoryInternalID"`
			ProductCategoryInternalID       string `json:"ProductCategoryInternalID"`
			ParentProductCategoryInternalID string `json:"ParentProductCategoryInternalID"`
			LanguageCode                    string `json:"LanguageCode"`
			LanguageCodeText                string `json:"LanguageCodeText"`
			ProductCategory                 string `json:"ProductCategory"`
			ProductAssignmentAllowed        bool   `json:"ProductAssignmentAllowed"`
		} `json:"results"`
	} `json:"d"`
}
