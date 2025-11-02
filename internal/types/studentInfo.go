package types

type StudentInfo struct {
	GroupName  string `json:"groupName"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
