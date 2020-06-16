package Excel

type WxSkillMain struct {
	SkillId   int            `json:"skillId"`
	SkillType int            `json:"skillType"`
	SkillName string         `json:"skillName"`
	MaxLevel  int            `json:"maxLevel"`
	MainDes   string         `json:"mainDes"`
	Levels    []WxSkillLevel `json:"levels"`
}

type WxSkillLevel struct {
	Xiuwei   int      `json:"xiuwei"`
	Banggong int      `json:"banggong"`
	Suiyin   int      `json:"suiyin"`
	Des      string   `json:"des"`
	Props    []string `json:"props"`
}
