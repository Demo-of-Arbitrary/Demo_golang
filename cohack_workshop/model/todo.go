package model

type Todo struct {
  ID int64 `gorm:"PRIMARY"`
  Description string `json:"description"`
  Resolved bool `json:"resolved"`
}

func CreateTodo(des string) (*Todo, error){
  todo := Todo{
    Description: des,
    Resolved: false,
  } 

  err := db.Model(&Todo{}).Save(&todo).Error
}
