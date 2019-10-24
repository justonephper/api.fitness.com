package models

type Department struct {
	Model
	Pid              int          `json:"pid"`
	DepartmentName   string       `json:"department_name"`
	DepartmentStatus int          `json:"department_status"`
	CompanyId        int          `json:"company_id"`
	CreatedBy        int          `json:"created_by"`
	UpdatedBy        int          `json:"updated_by"`
	Child            []Department `json:"child"`
}

//获取部门列表
func GetDepartments(pageNum int, pageSize int, where interface{}) (departments []Department) {
	db.Where(where).Offset(pageNum).Limit(pageSize).Find(&departments)
	return
}

//获取记录条数
func GetDepartmentTotal(where interface{}) (count int) {
	db.Model(&Department{}).Where(where).Count(&count)
	return
}

//部门分组
func DepartmentGroup(where interface{}) []Department {
	var department []Department
	var group []Department
	db.Where(where).Find(&department)
	if len(department) <= 0 {
		return department
	}
	for _, val := range department {
		if val.Pid == 0 {
			val.Child = getChild(val.ID, department)
			group = append(group, val)
		}
	}
	return group
}

func getChild(pid int, data []Department) []Department {
	returnData := []Department{}
	for _, val := range data {
		if pid == val.Pid {
			val.Child = getChild(val.ID, data)
			returnData = append(returnData, val)
		}
	}
	return returnData
}

//获取部门信息
func GetDepartmentById(id int) (department Department) {
	db.Where("id = ?", id).First(&department)
	return
}

//检测部门是否存在
func DepartmentExist(id int) bool {
	var depart Department
	db.Where("id = ?", id).First(&depart)
	if depart.ID > 0 {
		return true
	}
	return false
}

//添加部门
func AddDepartment(data map[string]interface{}) bool {
	db.Create(&Department{
		Pid:              data["pid"].(int),
		DepartmentName:   data["department_name"].(string),
		DepartmentStatus: data["department_status"].(int),
		CompanyId:        data["company_id"].(int),
		CreatedBy:        data["created_by"].(int),
		UpdatedBy:        data["updated_by"].(int),
	})
	return true
}

//更新部门信息
func UpdateDepartment(id int, data map[string]interface{}) bool {
	db.Model(&Department{}).Where("id = ?", id).Updates(data)
	return true
}

//删除部门信息
func DeleteDepartment(id int) bool {
	db.Where("id = ?", id).Delete(&Department{})
	return true
}
