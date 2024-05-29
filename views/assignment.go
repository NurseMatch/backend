package views

type Assignment struct {
	DisplayName          string `json:"displayName"`
	Description          string `json:"description"`
	HourlyRate           int    `json:"hourlyRate"`
	EducationRequirement string `json:"educationRequirement"`
	Role                 string `json:"role"`
	Location             string `json:"location"`
}
