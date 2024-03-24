package usecase

import "github.com/Nishad4140/SkillSync_ProtoFiles/pb"

type UsecaseInterface interface {
	UploadClientImage(req *pb.ImageRequest, profileId string) (string, error)
	UploadFreelancerImage(req *pb.ImageRequest, profileId string) (string, error)
}
