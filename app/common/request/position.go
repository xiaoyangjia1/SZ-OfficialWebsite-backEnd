package request

type Position struct {
	Title string `form:"title" json:"title" binding:"required"`
	Batch string `form:"batch" json:"batch" binding:"required"`
	Deadline int64 `form:"deadline" json:"deadline" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Requirements string `form:"requirements" json:"requirements" binding:"required"`
	Test int `form:"test" json:"test" binding:"required"`
	Interview int `form:"interview" json:"interview" binding:"required"`
	Check1 int `form:"check1" json:"check1" binding:"required"`
	Check2 int `form:"check2" json:"check2" binding:"required"`
	Head_count int `form:"head_count" json:"head_count" binding:"required"`
}

func (position Position) GetMessages() ValidatorMessages {
    return ValidatorMessages{
		"Title.required": "标题不能为空",
		"Batch.required": "批次不能为空",
		"Deadline.required": "截止日期不能为空",
		"Description.required": "描述不能为空",
		"Requirements.required": "要求不能为空",
		"Test.required": "笔试不能为空",
		"Interview.required": "面试不能为空",
		"Check1.required": "考核1不能为空",
		"Check2.required": "考核2不能为空",
		"Head_count.required": "录取名额不能为空",
    }
}