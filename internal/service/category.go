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

func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {

	category, err := c.CategoryDB.Create(req.Name, req.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil

}

func (c *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {

	categories, err := c.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{
		Categories: categoriesResponse,
	}, nil
}
