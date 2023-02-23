package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hmrbcnto.com/gin-api/entities"
)

type UserRepo interface {
	CreateUser(*entities.CreateUserRequest) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserByEmail(string) (*entities.User, error)
}

type userRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Client) UserRepo {
	return &userRepo{
		db: db.Database("kanban").Collection("users"),
	}
}

func (ur *userRepo) CreateUser(user *entities.CreateUserRequest) (*entities.User, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating user
	insertionResult, err := ur.db.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	// Getting inserted data
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Re querying the data
	createdRecord := ur.db.FindOne(ctx, filter)

	// Decode value to user entity
	createdUser := &entities.User{}
	createdRecord.Decode(createdUser)

	// returning value
	return createdUser, nil
}

func (ur *userRepo) GetAllUsers() ([]entities.User, error) {
	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	/// Getting all users
	// Creating query
	query := bson.D{{}}

	cursor, err := ur.db.Find(ctx, query)

	if err != nil {
		return nil, err
	}

	var users []entities.User
	// Iterate and decode
	err = cursor.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepo) GetUserByEmail(email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	user := &entities.User{}

	query := bson.D{primitive.E{Key: "email", Value: email}}

	err := ur.db.FindOne(ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
