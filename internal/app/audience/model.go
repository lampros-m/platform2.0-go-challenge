package audience

import "fmt"

// Audience : Descibes info of an audience.
type Audience struct {
	AssetID      uint32 `json:"-"`
	DateProduced string `json:"date_produced"`
	Gender       string `json:"gender"`
	Country      string `json:"country"`
	AgeFrom      uint32 `json:"age_from"`
	AgeTo        uint32 `json:"age_to"`
	HoursSpent   uint32 `json:"hours_spent"`
}

// AudienceSocialMedia : Describes info about audience in social media.
type AudienceSocialMedia struct {
	Audience
	SocialMedia string `json:"social_media"`
}

// AudienceSocialMediaMultiple : Describes multiple audiences for social media.
type AudienceSocialMediaMultiple []AudienceSocialMedia

// AudienceProducts : Describes info about shopping audience.
type AudienceProducts struct {
	Audience
	Product string `json:"string"`
}

// AudienceProductsMultiple : Describes multiple audiences for produects.
type AudienceProductsMultiple []AudienceProducts

// ProduceMessage : Produces message based on audience products info.
func (o *AudienceProducts) ProduceMessage() string {
	/*
		message example :
		On
		(1)2021-11-10 in
		(2)Greece
		(3)men from age
		(4)25 to
		(5)35 years old spent
		(6)3 hours looking for
		(7)shoes
	*/
	const messagePattern = "On %s in %s %s from age %d to %d years old spent %d hours looking for %s"

	var gender string
	switch o.Gender {
	case "man":
		gender = "men"
	case "woman":
		gender = "women"
	default:
		gender = "people"
	}

	return fmt.Sprintf(messagePattern, o.DateProduced, o.Country, gender, o.AgeFrom, o.AgeTo, o.HoursSpent, o.Product)
}

// AudienceSocialMedia : Produces message based on audience social media info.
func (o *AudienceSocialMedia) ProduceMessage() string {
	/*
		message example :
		On
		(1)2021-11-10 in
		(2)Greece
		(3)men from age
		(4)25 to
		(5)35 years old spent
		(6)3 hours on average using
		(7)Instagram
	*/
	const messagePattern = "On %s in %s %s from age %d to %d years old spent %d hours on avera geusing %s"

	var gender string
	switch o.Gender {
	case "man":
		gender = "men"
	case "woman":
		gender = "women"
	default:
		gender = "people"
	}

	return fmt.Sprintf(messagePattern, o.DateProduced, o.Country, gender, o.AgeFrom, o.AgeTo, o.HoursSpent, o.SocialMedia)
}
