package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"sync"
	"tokens-overhead/api"
	"tokens-overhead/repository/froles"
	"tokens-overhead/repository/gorm"
	"tokens-overhead/repository/http"
	"tokens-overhead/repository/jwt"
	"tokens-overhead/service"
)

func main() {
	mode := os.Getenv("MODE")
	if mode == "api" {
		jwtTokenRepo, err := jwt.NewJWTTokenImpl("keys/rsa.pub", "keys/rsa")
		if err != nil {
			panic(err)
		}
		svc := service.NewTokenService(
			nil, jwtTokenRepo, nil, nil, os.Getenv("MACHINE_NAME"), os.Getenv("CRYPT_METHOD"),
		)
		api.InitAPI(svc, os.Getenv("API_HOST"), os.Getenv("API_PORT"))
	} else if mode == "requester" {
		rolesRepo, err := froles.NewFileRolesImpl("roles/10x_gcloud_roles")
		if err != nil {
			panic(err)
		}
		jwtTokenRepo, err := jwt.NewJWTTokenImpl("keys/rsa.pub", "keys/rsa")
		if err != nil {
			panic(err)
		}
		// postgresURI := fmt.Sprintf("user=%s dbname=%s host=%s port=%v sslmode=%v password=%v")
		databaseRepo, err := gorm.NewGormDatabase(os.Getenv("POSTGRES_URI"))
		if err != nil {
			panic(err)
		}
		requestRepo := http.NewHTTPRequester()

		svc := service.NewTokenService(
			rolesRepo, jwtTokenRepo, requestRepo, databaseRepo, os.Getenv("MACHINE_NAME"), os.Getenv("CRYPT_METHOD"),
		)
		targetAddress := os.Getenv("TARGET_ADDRESS")
		reqTimes, _ := strconv.Atoi(os.Getenv("REQUEST_TIMES"))
		var wg sync.WaitGroup
		for routine := 0; routine < 50; routine++ {
			wg.Add(1)
			go func(routineNumber int, svc service.TokenService, reqTimes int, wg *sync.WaitGroup) {
				defer wg.Done()
				for i := 0; i < 1000/50; i++ {
					numRoles := 50*routineNumber + i
					for requestTimes := 0; requestTimes < reqTimes; requestTimes++ {
						for err = errors.New("start"); err != nil; err = svc.Execute(numRoles, targetAddress) {
						}
					}
				}
			}(routine, svc, reqTimes, &wg)
		}
		wg.Wait()
		log.Print("end")
	}
}
