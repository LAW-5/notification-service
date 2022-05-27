package shared

import (
	"context"
	"fmt"
	"log"
	"net"
	"notification/database"
	pb "notification/protobuf"
	"notification/utils"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
	pb.UnimplementedNotificationServiceServer
}

func (s *Server) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	var notifications []database.Notification
	s.db.Where("user_id = ?", in.UserId).Find(&notifications)

	log.Printf("user_id:%d - requesting to list notification", in.UserId)

	var datas = []*pb.Data{}
	for _, v := range notifications {
		data := &pb.Data{
			Id: int32(v.ID),
			Header: v.Header,
			Message: v.Message,
		}
		datas = append(datas, data)
	}

	response := &pb.ListResponse{
		Data: datas,
	}

	return response, nil
}

func NewNotificationGRPCServer(db *gorm.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", utils.ApiConfig.PORT))
	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &Server{db: db})

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis) 
	if err != nil {
		log.Fatal(err.Error())
	}
}