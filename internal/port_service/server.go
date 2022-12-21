import (
	pb "https://github.com/Christopher-Steel/port-service/api"

	"context"

	db_wrapper "internal/db"
)

type portServer struct {
	pb.UnimplementedPortServer
}

func (s *portServer) AddOrUpdate(ctx context.Context, request *pb.AddOrUpdateRequest) {
	db_wrapper.Port.AddOrUpdate(Port{
		Name: *request.Name
		City: *request.City
		Country: *request.Country
		Alias: *request.Alias
		Regions: *request.Regions
		Coordinates: *request.Coordinates
		Province: *request.Province
		Timezone: *request.Timezone
		Unlocs: *request.Unlocs
		Code: *request.Code
	})
}