package models

type GeneralModel interface {
	ToModel() App
}

type App struct {
	Bundle            string            `json:"bundle" db:"bundle"`
	DeveloperId       string            `json:"developerId" db:"developer_id"`
	Developer         string            `json:"developer" db:"developer"`
	Title             string            `json:"title" db:"title"`
	Categories        string            `json:"categories" db:"categories"`
	Price             string            `json:"price" db:"price"`
	Picture           string            `json:"picture" db:"picture"`
	Screenshots       []Screenshots     `json:"screenshots,omitempty"`
	Rating            string            `json:"rating" db:"rating"`
	ReviewCount       string            `json:"reviewCount" db:"review_count"`
	RatingHistogram   []string          `json:"ratingHistogram" db:"rating_histogram"`
	Description       string            `json:"description" db:"description"`
	ShortDescription  string            `json:"shortDescription" db:"short_description"`
	RecentChanges     string            `json:"recentChanges" db:"recent_changes"`
	ReleaseDate       string            `json:"releaseDate" db:"release_date"`
	LastUpdateDate    string            `json:"lastUpdateDate" db:"last_update_date"`
	AppSize           string            `json:"appSize" db:"app_size"`
	Installs          string            `json:"installs" db:"installs"`
	Version           string            `json:"version" db:"version"`
	AndroidVersion    string            `json:"androidVersion" db:"android_version"`
	ContentRating     string            `json:"contentRating" db:"content_rating"`
	DeveloperContacts DeveloperContacts `json:"developerContacts" db:"developer_contacts"`
	PrivacyPolicy     string            `json:"privacyPolicy,omitempty"`
}

type GoogleApp struct {
	App
	Screenshots []string `json:"screenshots,omitempty"`
}

func (ga *GoogleApp) ToModel() App {
	return App{
		Bundle:      ga.Bundle,
		DeveloperId: ga.DeveloperId,
		Developer:   ga.Developer,
		Title:       ga.Title,
		Categories:  ga.Categories,
		Price:       ga.Price,
		Picture:     ga.Picture,
		Screenshots: []Screenshots{
			{Device: "android", Screens: ga.Screenshots},
		},
		Rating:            ga.Rating,
		ReviewCount:       ga.ReviewCount,
		RatingHistogram:   ga.RatingHistogram,
		Description:       ga.Description,
		ShortDescription:  ga.ShortDescription,
		RecentChanges:     ga.RecentChanges,
		ReleaseDate:       ga.ReleaseDate,
		LastUpdateDate:    ga.LastUpdateDate,
		AppSize:           ga.AppSize,
		Installs:          ga.Installs,
		Version:           ga.Version,
		AndroidVersion:    ga.AndroidVersion,
		ContentRating:     ga.ContentRating,
		DeveloperContacts: ga.DeveloperContacts,
		PrivacyPolicy:     ga.PrivacyPolicy,
	}
}

type DeveloperContacts struct {
	Email    string `json:"email,omitempty"`
	Contacts string `json:"contacts,omitempty"`
}

type Screenshots struct {
	Device  string   `json:"device,omitempty"`
	Screens []string `json:"screenshots,omitempty"`
}