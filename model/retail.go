package model

// MasterItem for Master Material ...
type MasterItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Grouping    string `json:"grouping"`
	CreatedBy   string `json:"create_by"`
	Updatedby   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
}

// GroupingMaterial for grouping material ...
type GroupingMaterial struct {
	GroupingID  string `json:"grouping_id"`
	GroupName   string `json:"group_name"`
	CreatedBy   string `json:"created_by"`
	CreatedDate string `json:"created_date"`
}
