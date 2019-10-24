package models

type Target struct {
	Model
	TargetTitle       string  `json:"target_title"`
	TargetDesc        string  `json:"target_desc"`
	TargetCount       float64 `json:"target_count"`
	TargetUnit        string  `json:"target_unit"`
	TargetType        string  `json:"target_type"`
	StartMonth        MyTime  `json:"start_month"`
	EndMonth          MyTime  `json:"end_month"`
	ExecuteDepartment uint    `json:"execute_department"`
	CompleteCount     float64 `json:"complete_count"`
	ExecuteStatus     byte    `json:"execute_status"`
	DispathType       byte    `json:"dispatch_type"`
	CompanyId         byte    `json:"company_id"`
	CreatedBy         string  `json:"created_by"`
	ExecuteBy         string  `json:"execute_by"`
}

func GetTargets(pageNum int, pageSize int, maps interface{}) (targets []Target) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&targets)
	return
}

func GetTargetTotal(maps interface{}) (count int) {
	db.Model(&Target{}).Where(maps).Count(&count)
	return
}
