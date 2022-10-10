package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hexagonal_boilerplate/core/entities"
	"hexagonal_boilerplate/infrastructure/repository/mongo/mapper"
	"hexagonal_boilerplate/infrastructure/repository/mongo/models"
)

type (
	RepositoryOutlet struct {
		mongo  *mongo.Client
		DBName string
	}
)

var collectionOutletName = "outlet"

func NewOutletRepo(db *mongo.Client, dbname string) *RepositoryOutlet {
	return &RepositoryOutlet{db, dbname}
}
func (r *RepositoryOutlet) Get(ID string) (*entities.Outlet, error) {
	var outletModel models.OutletModel
	objID, _ := primitive.ObjectIDFromHex(ID)
	var selector = bson.M{"_id": objID}
	coll := r.mongo.Database(r.DBName).Collection(collectionOutletName)
	err := coll.FindOne(context.TODO(), selector).Decode(&outletModel)
	if err != nil {
		return nil, err
	}
	entity, err := mapper.MapToEntities(&outletModel)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (r *RepositoryOutlet) Create(c context.Context, outlet *entities.Outlet) error {
	objModel := mapper.MapToWriteModels(outlet)
	coll := r.mongo.Database(r.DBName).Collection(collectionOutletName)
	insert, err := coll.InsertOne(c, objModel)
	if err != nil {
		return err
	}

	var insertID string
	if oid, ok := insert.InsertedID.(primitive.ObjectID); ok {
		insertID = oid.Hex()
	} else {
		// Not objectid.ObjectID, do what you want
	}
	outlet.SetID(&insertID, &insertID)
	return nil
}
func (r *RepositoryOutlet) Update(c context.Context, outlet *entities.Outlet) error {
	objModel := mapper.MapToWriteModels(outlet)
	var selector = bson.M{"_id": *outlet.ID}
	var changes = objModel
	coll := r.mongo.Database(r.DBName).Collection(collectionOutletName)
	_, err := coll.UpdateOne(c, selector, bson.M{"$set": changes})
	if err != nil {
		return err
	}
	return nil
}
