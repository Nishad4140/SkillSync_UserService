package usecase

import "github.com/Nishad4140/SkillSync_ProtoFiles/pb"

type UsecaseInterface interface {
	UploadImage(req *pb.ImageRequest, profileId string) (string, error)
}
