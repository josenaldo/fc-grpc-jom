package service

import (
	"context"

	"github.com/josenaldo/fc-grpc-jom/internal/database"
	"github.com/josenaldo/fc-grpc-jom/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {

	category, err := c.CategoryDB.Create(req.Name, req.Description)

	if err != nil {
		return nil, err
	}

	category := &pb.Category {
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
	},


	categoryResponse := &pb.CategoryResponse{
		Category: &category,
	}

	return categoryResponse, nil

}
