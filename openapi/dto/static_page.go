package openapidto

type StaticPageResponse struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	PageContent string `json:"page_content"`
}
