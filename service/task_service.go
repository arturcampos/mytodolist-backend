package service

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"todolist/database"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
	Text      string             `bson:"text" json:"text"`
	Completed bool               `bson:"completed" json:"completed"`
}

func (t *Task) handleNewTask() {
	t.ID = primitive.NewObjectID()
	t.Completed = false
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func CreateTask(task *Task) error {
	task.handleNewTask()
	_, err := database.Collection.InsertOne(database.Ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func GetTask() ([]*Task, error) {

	cursor, err := database.Collection.Find(database.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	return cursorConverter(cursor), nil

}

func cursorConverter(cursor *mongo.Cursor) []*Task{
	var tasks []*Task
		//converter cursor result to Task
		for cursor.Next(database.Ctx) {
			var t Task
			err := cursor.Decode(&t)
			if err != nil {
				panic(err)
			}
	
			tasks = append(tasks, &t)
		}
	
		if len(tasks) == 0 {
			tasks = append(tasks,
				&Task{
					ID:        primitive.NewObjectID(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					Text:      "Unamed Task",
					Completed: false,
				})
		}
	
		return tasks
}

func GetTaskById(id string) *Task {
	var task *Task

	objectId, _ := primitive.ObjectIDFromHex(id)
	database.Collection.FindOne(database.Ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&task)
	return task

}

func UpdateTaskText(id string, text string) (*Task, error) {

	task := GetTaskById(id)
	if task == nil {
		return nil, errors.New("Entity not found: " + id)
	}

	task.Text = text
	task.UpdatedAt = time.Now()
	_, err := database.Collection.ReplaceOne(database.Ctx, bson.D{{Key: "_id", Value: &task.ID}}, &task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func CompleteTask(id string) (*Task, error) {
	task := GetTaskById(id)
	if task == nil {
		return nil, errors.New("Entity not found: " + id)
	}
	task.changeStatus(true)
	return task, nil
}

func UncompleteTask(id string) (*Task, error) {
	task := GetTaskById(id)
	if task == nil {
		return nil, errors.New("Entity not found: " + id)
	}
	task.changeStatus(false)
	return task, nil
}

func (t *Task) changeStatus(status bool) {
	t.Completed = status
	t.UpdatedAt = time.Now()

	database.Collection.ReplaceOne(database.Ctx, bson.D{{Key: "_id", Value: t.ID}}, &t)
}

func DeleteTask(id string) error {

	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := database.Collection.DeleteOne(database.Ctx, bson.D{{Key: "_id", Value: objectId}})
	return err

}
