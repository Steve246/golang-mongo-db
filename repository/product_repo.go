package repository

import (
	"errors"
	"golang-mongodb/model"
	"golang-mongodb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
	Pagination(limitNumber *model.FindLimit)([]model.Product, error)

	UpdateProduct(id string, updateProduct *model.Product) (*model.Product,error) //update masih blm jalan, tanpa id pun

	DeleteProduct(id string) error

	GetById(id string) (*model.Product, error)

	GetByCategory(category string) ([]model.Product, error)
}

type productRepository struct {
	db *mongo.Database
}

func (p *productRepository) GetByCategory(category string) ([]model.Product, error){
	

	//
	var products []model.Product

	ctx, cancel := utils.InitContext()
	defer cancel() 

	filter := bson.M{"category": category}

	numerLimiter := 5
	
	opts := options.Find().SetLimit(int64(numerLimiter))

	cursor, err := p.db.Collection("products").Find(ctx, filter, opts)

	if err != nil {
		return nil, err 
	}

	for cursor.Next(ctx) {
		var product model.Product

		err = cursor.Decode(&product)

		if err != nil {
			return nil, err 
		}
		products = append(products, product)
	}

	return products, nil 

}

func (p *productRepository) GetById(id string) (*model.Product, error) {

	ctx, cancel := utils.InitContext()
	defer cancel()



	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var product *model.Product

	if err := p.db.Collection("products").FindOne(ctx, query).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return product, nil

}

func (p *productRepository) DeleteProduct(id string) error {
	ctx, cancel := utils.InitContext()
	defer cancel()

	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	res, err := p.db.Collection("products").DeleteOne(ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil


}

func (p *productRepository) UpdateProduct(id string, updateProduct *model.Product) (*model.Product,error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	doc, err := utils.ToDoc(updateProduct)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key:"_id", Value: obId}}
	update := bson.D{{Key:"$set", Value:doc}}

	// result, err := p.db.Collection("product").UpdateOne(ctx, filter, update)

	res := p.db.Collection("products").FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *model.Product

	if err := res.Decode(&updatedPost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatedPost, nil 


}



func (p *productRepository) Pagination(limitNumber *model.FindLimit)([]model.Product, error) {
	var products []model.Product

	ctx, cancel := utils.InitContext()
	defer cancel() 

	filter := bson.D{}

	numerLimiter := limitNumber.Number
	
	opts := options.Find().SetLimit(int64(numerLimiter))

	cursor, err := p.db.Collection("products").Find(ctx, filter, opts)

	if err != nil {
		return nil, err 
	}

	for cursor.Next(ctx) {
		var product model.Product

		err = cursor.Decode(&product)

		if err != nil {
			return nil, err 
		}
		products = append(products, product)
	}

	return products, nil 

}

func (p *productRepository) Add (newProduct *model.Product) error {
	ctx, cancel := utils.InitContext()
	defer cancel()

	newProduct.Id = primitive.NewObjectID()
	_, err := p.db.Collection("products").InsertOne(ctx, newProduct)

	if err != nil {
		return err 
	}

	return nil 
}

func (p *productRepository) Retrieve() ([]model.Product, error) {

	var products []model.Product
	ctx, cancel := utils.InitContext()

	defer cancel()

	cursor, err := p.db.Collection("products").Find(ctx, bson.M{})

	if err != nil {
		return nil, err 
	}

	for cursor.Next(ctx) {
		var product model.Product

		err = cursor.Decode(&product)

		if err != nil {
			return nil, err 
		}
		products = append(products, product)
	}

	return products, nil 

}

func NewProductRepository(db *mongo.Database) ProductRepository{
	repo := new(productRepository)
	repo.db = db
	return repo 
}