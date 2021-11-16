package pagination

// PageInfoRequest : Pangeinfo from a request.
type PageInfoRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"per_page"`
}

// PageInfoResponse : Pageinfo for a response.
type PageInfoResponse struct {
	TotalPages uint32 `json:"total_pages,omitempty"`
	TotalItems uint32 `json:"total_items,omitempty"`
}

// PageInfo : A combination of pagination request and response.
type PageInfo struct {
	PageInfoRequest
	PageInfoResponse
}

// Offset : Returns the offset based on page and pagesize.
func (p *PageInfoRequest) Offset() uint32 {
	return (p.Page - 1) * p.PageSize
}

// GetOrDefaultPageInfoRequest : Checks for zero page or pagesize and replaces with default.
func (p *PageInfoRequest) GetOrDefaultPageInfoRequest(defaultPage, defaultSize uint32) {
	if p.Page == 0 {
		p.Page = defaultPage
	}

	if p.PageSize == 0 {
		p.PageSize = defaultSize
	}
}
