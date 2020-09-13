package property

type ResultDataProperty struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`

	UserMessage  string `json:"userMessage"`
	UserMoreInfo string `json:"userMoreInfo"`

	DeveloperMessage  string `json:"developerMessage"`
	DeveloperMoreInfo string `json:"developerMoreInfo"`
}
