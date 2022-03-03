package service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"todolist/database"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
	Text      string             `bson:"text" json:"text"`
	Completed bool               `bson:"completed" json:"completed"`
}

func (t *Task) handleNewTask(){
	t.Completed = false
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func CreateTask(task *Task) error {
	task.handleNewTask()
	_, err := database.Collection.InsertOne(database.Ctx, task)
	if err != nil{
		return err
	}
	return nil
}

func GetTask() ([]*Task, error) {
	var tasks []*Task
	cursor, err := database.Collection.Find(database.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	//converter cursor result to Task
	for cursor.Next(database.Ctx) {
		var t Task
		err := cursor.Decode(&t)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, &t)
	}

	if len(tasks) == 0{
		tasks = append(tasks, 
			&Task{
				ID: primitive.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Text:"Unamed Task",
				Completed: false,
			},)
		}

	return tasks, nil
}
