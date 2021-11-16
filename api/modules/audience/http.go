package audience

import (
	"gwi/platform2.0-go-challenge/pkg/daterange"
	"gwi/platform2.0-go-challenge/pkg/pagination"
	"gwi/platform2.0-go-challenge/pkg/sorting"
)

// AudienceRequest : Describes a general audience request.
type AudienceRequest struct {
	daterange.DateRange
}

// GetAudienceSocialMediaRequest : Describes a request for social media audiences.
type GetAudienceSocialMediaRequest struct {
	AudienceRequest
	pagination.PageInfoRequest
	sorting.Sorting
}

// GetAudienceProductsRequest : Describes a request for product audiences.
type GetAudienceProductsRequest struct {
	AudienceRequest
	pagination.PageInfoRequest
	sorting.Sorting
}
