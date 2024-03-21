package main

import (
	"log"
	"net"
	"os"

	"github.com/Nishad4140/SkillSync_ProtoFiles/pb"
	"github.com/Nishad4140/SkillSync_UserService/db"
	"github.com/Nishad4140/SkillSync_UserService/initializer"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err.Error())
	}
	addr := os.Getenv("DB_KEY")
	db, err := db.InitDB(addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	services := initializer.Initializer(db)
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, services)
	lis, err := net.Listen("tcp", ":4001")
	if err != nil {
		log.Fatalf("failed to run on the port 4001 : %v", err)
	}
	log.Printf("user service listening on the port 4001")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to run on the port 4001 : %v", err)
	}
}
