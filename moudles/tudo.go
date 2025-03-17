package moudles

import "bubbon/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}

// CreateATable 添加表记录

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func GetAllTodo() (todo []Todo, err error) {
	if err = dao.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
