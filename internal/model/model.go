package model

import "gorm.io/gorm"

func CreateTask(db *gorm.DB, task *Task) error {
	return db.Create(task).Error
}

func GetTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTask(db *gorm.DB, id int) (*Task, error) {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(db *gorm.DB, id int, task *Task) error {
	if err := db.First(task, id).Error; err != nil {
		return err
	}
	return db.Save(task).Error
}

func DeleteTask(db *gorm.DB, id int) error {
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}
	return db.Delete(&task).Error
}
