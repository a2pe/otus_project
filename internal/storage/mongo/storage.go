package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"otus_project/internal/model"
	"otus_project/internal/model/common"
	"time"
)

type Storage interface {
	SaveItem(item common.Item) error
	GetByID(itemType string, id int) (common.Item, bool)
	GetAllItems(itemType string) ([]common.Item, error)
	UpdateItem(itemType string, item common.Item) bool
	DeleteItem(itemType string, id int) bool
}

type MongoStorage struct {
	client     *mongo.Client
	database   *mongo.Database
	collection map[string]*mongo.Collection
}

func NewMongoStorage(ctx context.Context, uri, dbName string) (*MongoStorage, error) {
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(dbName)

	collections := map[string]*mongo.Collection{
		"user":       db.Collection("users"),
		"project":    db.Collection("projects"),
		"task":       db.Collection("tasks"),
		"reminder":   db.Collection("reminders"),
		"tag":        db.Collection("tags"),
		"time_entry": db.Collection("time_entries"),
	}

	return &MongoStorage{
		client:     client,
		database:   db,
		collection: collections,
	}, nil
}

func (m *MongoStorage) SaveItem(ctx context.Context, item common.Item, itemType string) error {
	coll, ok := m.collection[itemType]
	if !ok {
		return errors.New("unsupported item type")
	}

	item.SetID(uint(time.Now().UnixNano()))
	_, err := coll.InsertOne(ctx, item)
	return err
}

func (m *MongoStorage) GetAllItems(ctx context.Context, itemType string) ([]interface{}, error) {
	coll, ok := m.collection[itemType]
	if !ok {
		return nil, errors.New("unsupported item type")
	}

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []interface{}
	switch itemType {
	case "task":
		var tasks []*model.Task
		if err := cursor.All(ctx, &tasks); err != nil {
			return nil, err
		}
		for _, t := range tasks {
			results = append(results, t)
		}
	case "user":
		var users []*model.User
		if err := cursor.All(ctx, &users); err != nil {
			return nil, err
		}
		for _, u := range users {
			results = append(results, u)
		}
	case "project":
		var projects []*model.Project
		if err := cursor.All(ctx, &projects); err != nil {
			return nil, err
		}
		for _, project := range projects {
			results = append(results, project)
		}
	case "reminder":
		var reminders []*model.Reminder
		if err := cursor.All(ctx, &reminders); err != nil {
			return nil, err
		}
		for _, reminder := range reminders {
			results = append(results, reminder)
		}
	case "tag":
		var tags []*model.Tag
		if err := cursor.All(ctx, &tags); err != nil {
			return nil, err
		}
		for _, tag := range tags {
			results = append(results, tag)
		}
	case "time_entry":
		var timeEntries []*model.TimeEntry
		if err := cursor.All(ctx, &timeEntries); err != nil {
			return nil, err
		}
		for _, timeEntry := range timeEntries {
			results = append(results, timeEntry)
		}
	default:
		return nil, errors.New("unsupported item type")
	}
	return results, nil
}

func (m *MongoStorage) GetByID(ctx context.Context, itemType string, id uint) (interface{}, error) {
	coll, ok := m.collection[itemType]
	if !ok {
		return nil, errors.New("unsupported item type")
	}

	filter := bson.M{"id": id}
	switch itemType {
	case "task":
		var task model.Task
		err := coll.FindOne(ctx, filter).Decode(&task)
		return &task, err
	case "user":
		var user model.User
		err := coll.FindOne(ctx, filter).Decode(&user)
		return &user, err
	case "project":
		var project model.Project
		err := coll.FindOne(ctx, filter).Decode(&project)
		return &project, err
	case "reminder":
		var reminder model.Reminder
		err := coll.FindOne(ctx, filter).Decode(&reminder)
		return &reminder, err
	case "tag":
		var tag model.Tag
		err := coll.FindOne(ctx, filter).Decode(&tag)
		return &tag, err
	case "time_entry":
		var timeEntries []*model.TimeEntry
		err := coll.FindOne(ctx, filter).Decode(&timeEntries)
		return timeEntries, err
	default:
		return nil, fmt.Errorf("unsupported item type")
	}
}

func (m *MongoStorage) DeleteByID(ctx context.Context, itemType string, id uint) (interface{}, error) {
	coll, ok := m.collection[itemType]
	if !ok {
		return nil, errors.New("unsupported item type")
	}

	filter := bson.M{"id": id}

	res, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, fmt.Errorf("no %s found with id %d", itemType, id)
	}
	return fmt.Sprintf("%s with id %d deleted", itemType, id), nil
}

func (m *MongoStorage) UpdateByID(ctx context.Context, itemType string, id uint, updated common.Item) error {
	coll, ok := m.collection[itemType]
	if !ok {
		return fmt.Errorf("unsupported item type: %s", itemType)
	}

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": updated,
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}
