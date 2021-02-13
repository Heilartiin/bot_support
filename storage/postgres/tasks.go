package postgres

import (
	"github.com/Heilartin/bot_support/models"
	"github.com/pkg/errors"
)

func (db *DB) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	q := `SELECT * FROM mrp_tasks WHERE active = true;`
	err := db.DB.Select(&tasks, q)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tasks, nil
}

func (db *DB) GetTaskByID(taskID int) (*models.Task, error)  {
	var task models.Task
	q := `SELECT * FROM mrp_tasks WHERE id = $1;`
	err := db.DB.Get(&task, q, taskID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &task, nil
}
