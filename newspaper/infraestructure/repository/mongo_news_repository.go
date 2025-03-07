package repository

import (
	"context"
	"errors"
	"main/newspaper/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoAlbumRepository struct {
	db *mongo.Database
}

func NewMongoAlbumRepository(conn *mongo.Database) *MongoAlbumRepository {
	return &MongoAlbumRepository{db: conn}
}

func (r *MongoAlbumRepository) Save(ctx context.Context, news domain.News) error {
	collection := r.db.Collection("news")

	_, err := collection.InsertOne(ctx, bson.M{
		"Title":           news.Title,
		"Author":          news.Author,
		"Description":     news.Description,
		"Content":         news.Content,
		"Topic":           news.Topic,
		"PublicationDate": news.PublicationDate,
	})
	if err != nil {
		return errors.New("Error al guardar la noticia: " + err.Error())
	}
	return nil
}

func (r *MongoAlbumRepository) GetNewsById(ctx context.Context, id primitive.ObjectID) (domain.News, error) {
	collection := r.db.Collection("news")

	filter := bson.M{"_id": id}

	var news domain.News

	err := collection.FindOne(ctx, filter).Decode(&news)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.News{}, domain.ErrNewsNotFound
		}
		return domain.News{}, errors.New("error al buscar la noticia: " + err.Error())
	}

	return news, nil
}

func (r *MongoAlbumRepository) GetAllNews(ctx context.Context) ([]domain.News, error) {
	collection := r.db.Collection("news")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Error al obtener los álbumes: " + err.Error())
	}
	defer cursor.Close(ctx)

	var news []domain.News
	for cursor.Next(ctx) {
		var new domain.News
		if err := cursor.Decode(&new); err != nil {
			return nil, errors.New("Error al decodificar un álbum: " + err.Error())
		}
		news = append(news, new)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.New("Error iterando sobre las noticias: " + err.Error())
	}

	return news, nil
}

func (r *MongoAlbumRepository) UpdateNews(ctx context.Context, news domain.News) (domain.News, error) {
	collection := r.db.Collection("news")

	objectId, err := primitive.ObjectIDFromHex(news.Id.Hex())
	if err != nil {
		return domain.News{}, errors.New("ID invalido: " + err.Error())
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"Title":            news.Title,
			"Author":           news.Author,
			"Description":      news.Description,
			"Content":          news.Content,
			"Topic":            news.Topic,
			"Publication Date": news.PublicationDate,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.News{}, errors.New("Error al actualizar la noticia: " + err.Error())
	}

	if result.MatchedCount == 0 {
		return domain.News{}, errors.New("Ningun noticia coincide con el filtro")
	}
	return news, nil
}

func (r *MongoAlbumRepository) DeleteNews(ctx context.Context, id primitive.ObjectID) error {
	collection := r.db.Collection("news")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return domain.ErrNewsNotFound
	}
	return nil
}
