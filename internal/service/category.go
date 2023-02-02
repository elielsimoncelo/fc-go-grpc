package service

import (
	"context"
	"io"

	"github.com/elielsimoncelo/fc-go-grpc/internal/database"
	"github.com/elielsimoncelo/fc-go-grpc/internal/pb"
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

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	createdCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	response := &pb.CategoryResponse{
		Category: createdCategory,
	}

	return response, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.EmptyRequest) (*pb.CategoryListResponse, error) {
	categories, err := c.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.CategoryResponse

	for _, category := range categories {
		foundCategory := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoryResponse := &pb.CategoryResponse{
			Category: foundCategory,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	response := &pb.CategoryListResponse{
		Categories: categoriesResponse,
	}

	return response, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Find(in.Id)

	if err != nil {
		return nil, err
	}

	foundCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	response := &pb.CategoryResponse{
		Category: foundCategory,
	}

	return response, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryListResponse{}

	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.CategoryResponse{
			Category: &pb.Category{
				Id:          result.ID,
				Name:        category.Name,
				Description: category.Description,
			},
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		response := &pb.CategoryResponse{
			Category: &pb.Category{
				Id:          result.ID,
				Name:        category.Name,
				Description: category.Description,
			},
		}

		err = stream.Send(response)

		if err != nil {
			return err
		}
	}
}
