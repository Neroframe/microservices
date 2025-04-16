package grpcserver

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *InventoryGRPCServer) CreateCategory(ctx context.Context, req *inventorypb.CreateCategoryRequest) (*inventorypb.CategoryResponse, error) {
	c := &domain.Category{
		Name: req.Name,
	}

	if err := s.categoryUsecase.Create(ctx, c); err != nil {
		utils.Log.Error("CreateCategory failed", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &inventorypb.CategoryResponse{
		Id:   c.ID,
		Name: c.Name,
	}, nil
}

func (s *InventoryGRPCServer) GetCategoryByID(ctx context.Context, req *inventorypb.GetCategoryRequest) (*inventorypb.CategoryResponse, error) {
	cat, err := s.categoryUsecase.GetByID(ctx, req.Id)
	if err != nil {
		utils.Log.Error("GetCategoryByID failed", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	if cat == nil {
		return nil, status.Error(codes.NotFound, "category not found")
	}

	return &inventorypb.CategoryResponse{
		Id:   cat.ID,
		Name: cat.Name,
	}, nil
}

func (s *InventoryGRPCServer) UpdateCategory(ctx context.Context, req *inventorypb.UpdateCategoryRequest) (*inventorypb.CategoryResponse, error) {
	c := &domain.Category{
		ID:   req.Id,
		Name: req.Name,
	}

	if err := s.categoryUsecase.Update(ctx, c); err != nil {
		utils.Log.Error("UpdateCategory failed", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &inventorypb.CategoryResponse{
		Id:   c.ID,
		Name: c.Name,
	}, nil
}

func (s *InventoryGRPCServer) DeleteCategory(ctx context.Context, req *inventorypb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if err := s.categoryUsecase.Delete(ctx, req.Id); err != nil {
		utils.Log.Error("DeleteCategory failed", "err", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *InventoryGRPCServer) ListCategories(ctx context.Context, _ *inventorypb.ListCategoriesRequest) (*inventorypb.ListCategoriesResponse, error) {
	categories, err := s.categoryUsecase.List(ctx)
	if err != nil {
		utils.Log.Error("ListCategories failed", "err", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var res []*inventorypb.CategoryResponse
	for _, c := range categories {
		res = append(res, &inventorypb.CategoryResponse{
			Id:   c.ID,
			Name: c.Name,
		})
	}

	return &inventorypb.ListCategoriesResponse{Categories: res}, nil
}
