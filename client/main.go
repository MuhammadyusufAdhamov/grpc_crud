package main

import (
	"context"
	

	crudService "github.com/MuhammadyusufAdhamov/grpc_crud/genproto/crud_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := crudService.NewUserCrudClient(conn)

	// resp, err := client.Create(context.Background(), &crudService.User{FirstName: "Muhammadyusuf", LastName: "Adhamov", Age: 19, PhoneNumber: "93441509"})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(resp)

	// data, err := client.Get(context.Background(), &crudService.IdUser{Id: 8})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data)

	// data, err := client.Update(context.Background(), &crudService.User{
	// 	FirstName: "Muhammadyusuf",
	// 	LastName:  "Adhamov",
	// 	Age:       20,
	// 	Id:        8,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data)

	// data, err := client.GetAll(context.Background(), &crudService.GetAllUsersParams{
	// 	Limit: 10,
	// 	Page:  1,
	// 	Search: "uha",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data)

	_, err = client.Delete(context.Background(), &crudService.IdUser{Id: 8})
	if err != nil {
		panic(err)
	}
}
