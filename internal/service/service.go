package service

import (
	"context"
	"fmt"

	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"github.com/Nishad4140/SkillSync_UserService/entities"
	"github.com/Nishad4140/SkillSync_UserService/internal/adapters"
	"github.com/Nishad4140/SkillSync_UserService/internal/helper"
	"github.com/Nishad4140/SkillSync_UserService/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	adapters adapters.AdapterInterface
	usecase  usecase.UsecaseInterface
	pb.UnimplementedUserServiceServer
}

func NewUserService(adapters adapters.AdapterInterface, usecase usecase.UsecaseInterface) *UserService {
	return &UserService{
		adapters: adapters,
		usecase:  usecase,
	}
}

func (user *UserService) ClientSignup(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	if req.Email == "" {
		return nil, fmt.Errorf("email can't be empty")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}
	if req.Phone == "" {
		return nil, fmt.Errorf("phone can't be empty")
	}
	if req.Password == "" {
		return nil, fmt.Errorf("password can't be empty")
	}
	emailCheck, err := user.adapters.GetClientByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if emailCheck.Name != "" {
		return nil, fmt.Errorf("this email is already used")
	}
	phoneCheck, err := user.adapters.GetClientByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if phoneCheck.Name != "" {
		return nil, fmt.Errorf("this phone number is already used")
	}
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	reqEntity := entities.Client{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	}
	res, err := user.adapters.ClientSignup(reqEntity)
	if err != nil {
		return nil, err
	}
	return &pb.UserSignUpResponse{
		Id:    res.ID.String(),
		Name:  res.Name,
		Email: res.Email,
		Phone: res.Phone,
	}, nil
}

func (user *UserService) CreateProfile(ctx context.Context, req *pb.GetUserById) (*emptypb.Empty, error) {
	if err := user.adapters.CreateProfile(req.Id); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (user *UserService) FreelancerSignup(ctx context.Context, req *pb.FreelancerSignUpRequest) (*pb.FreelancerSignUpResponse, error) {
	if req.Name == "" {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("enter a valid name")
	}
	emailCheck, err := user.adapters.GetFreelancerByEmail(req.Email)
	if err != nil {
		return &pb.FreelancerSignUpResponse{}, err
	}
	if emailCheck.Name != "" {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("this email is already in use")
	}
	phoneCheck, err := user.adapters.GetFreelancerByPhone(req.Phone)
	if err != nil {
		return &pb.FreelancerSignUpResponse{}, err
	}
	if phoneCheck.Name != "" {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("this phone number is already in use")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return &pb.FreelancerSignUpResponse{}, err
	}
	reqEntity := entities.Freelancer{
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		CategoryId: int32(req.CategoryId),
		Password:   hashedPassword,
	}
	res, err := user.adapters.FreelancerSignup(reqEntity)
	return &pb.FreelancerSignUpResponse{
		Id:         res.ID.String(),
		Name:       res.Name,
		Email:      res.Email,
		Phone:      res.Phone,
		CategoryId: res.CategoryId,
	}, err
}
