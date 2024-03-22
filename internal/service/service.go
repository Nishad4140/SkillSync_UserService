package service

import (
	"context"
	"fmt"

	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"github.com/Nishad4140/SkillSync_UserService/entities"
	"github.com/Nishad4140/SkillSync_UserService/internal/adapters"
	"github.com/Nishad4140/SkillSync_UserService/internal/helper"
	"github.com/Nishad4140/SkillSync_UserService/internal/usecase"
	"github.com/google/uuid"
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

func (user *UserService) ClientSignup(ctx context.Context, req *pb.ClientSignUpRequest) (*pb.ClientSignUpResponse, error) {
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
	return &pb.ClientSignUpResponse{
		Id:    res.ID.String(),
		Name:  res.Name,
		Email: res.Email,
		Phone: res.Phone,
	}, nil
}

func (user *UserService) CreateProfile(ctx context.Context, req *pb.GetUserById) (*emptypb.Empty, error) {
	if err := user.adapters.CreateClientProfile(req.Id); err != nil {
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

func (user *UserService) ClientLogin(ctx context.Context, req *pb.LoginRequest) (*pb.ClientSignUpResponse, error) {
	if req.Emial == "" {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("please enter the email")
	}
	clientData, err := user.adapters.GetClientByEmail(req.Emial)
	if err != nil {
		return &pb.ClientSignUpResponse{}, err
	}
	if clientData.Email == "" {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("invalid credentials")
	}
	if !helper.CompareHashedPassword(clientData.Password, req.Password) {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("invalid password")
	}
	return &pb.ClientSignUpResponse{
		Id:    clientData.ID.String(),
		Name:  clientData.Name,
		Email: clientData.Email,
		Phone: clientData.Phone,
	}, nil
}

func (user *UserService) FreelancerLogin(ctx context.Context, req *pb.LoginRequest) (*pb.FreelancerSignUpResponse, error) {
	if req.Emial == "" {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("enter a valid email")
	}
	freelancerData, err := user.adapters.GetFreelancerByEmail(req.Emial)
	if err != nil {
		return &pb.FreelancerSignUpResponse{}, err
	}
	if freelancerData.Email == "" {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("invalid credentials")
	}
	if !helper.CompareHashedPassword(freelancerData.Password, req.Password) {
		return &pb.FreelancerSignUpResponse{}, fmt.Errorf("invalid password")
	}
	return &pb.FreelancerSignUpResponse{
		Id:         freelancerData.ID.String(),
		Name:       freelancerData.Name,
		Email:      freelancerData.Email,
		Phone:      freelancerData.Phone,
		CategoryId: freelancerData.CategoryId,
	}, nil
}

func (user *UserService) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.ClientSignUpResponse, error) {
	if req.Emial == "" {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("enter a valid mail")
	}
	adminData, err := user.adapters.GetAdminByEmail(req.Emial)
	if err != nil {
		return &pb.ClientSignUpResponse{}, err
	}
	if adminData.Email == "" {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("invalid credentials")
	}
	if !helper.CompareHashedPassword(adminData.Password, req.Password) {
		return &pb.ClientSignUpResponse{}, fmt.Errorf("invalid password")
	}
	return &pb.ClientSignUpResponse{
		Id:    adminData.ID.String(),
		Name:  adminData.Name,
		Email: adminData.Email,
		Phone: adminData.Phone,
	}, nil
}

func (user *UserService) AddCategory(ctx context.Context, req *pb.AddCategoryRequest) (*emptypb.Empty, error) {
	reqEntity := entities.Category{
		Name: req.Category,
	}
	nameCheck, err := user.adapters.GetCategoryByName(req.Category)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if nameCheck.Name != "" {
		return &emptypb.Empty{}, fmt.Errorf("category already exists")
	}
	err = user.adapters.AdminAddCategory(reqEntity)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (user *UserService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*emptypb.Empty, error) {
	reqEntity := entities.Category{
		ID:   int(req.Id),
		Name: req.Category,
	}
	nameCheck, err := user.adapters.GetCategoryByName(req.Category)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if nameCheck.Name != "" {
		return &emptypb.Empty{}, fmt.Errorf("category already exists")
	}
	err = user.adapters.AdminUpdateCategory(reqEntity)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (user *UserService) GetAllCategory(req *emptypb.Empty, srv pb.UserService_GetAllCategoryServer) error {
	categories, err := user.adapters.GetAllCategories()
	if err != nil {
		return err
	}
	for _, category := range categories {
		res := &pb.UpdateCategoryRequest{
			Id:       int32(category.ID),
			Category: category.Name,
		}
		if err := srv.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (user *UserService) ClientAddAddress(ctx context.Context, req *pb.AddAddressRequest) (*emptypb.Empty, error) {
	address, err := user.adapters.GetAddressByUserId(req.UserId)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if address.Country != "" {
		return &emptypb.Empty{}, fmt.Errorf("you have added a address already, please edit on that")
	}
	reqEntity := entities.Address{
		Country:  req.Country,
		State:    req.State,
		District: req.District,
		City:     req.City,
	}
	if err := user.adapters.ClientAddAddress(reqEntity, req.UserId); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (user *UserService) ClientUpdateAddress(ctx context.Context, req *pb.AddressResponse) (*emptypb.Empty, error) {
	addressId, err := uuid.Parse(req.Id)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	reqEntity := entities.Address{
		Id:       addressId,
		Country:  req.Country,
		State:    req.State,
		District: req.District,
		City:     req.City,
	}
	if err := user.adapters.ClientUpdateAddress(reqEntity); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (user *UserService) ClientGetAddress(ctx context.Context, req *pb.GetUserById) (*pb.AddressResponse, error) {
	address, err := user.adapters.GetAddressByUserId(req.Id)
	if err != nil {
		return &pb.AddressResponse{}, err
	}
	res := &pb.AddressResponse{
		Id:       address.Id.String(),
		Country:  address.Country,
		State:    address.State,
		District: address.District,
		City:     address.City,
	}
	return res, nil
}
