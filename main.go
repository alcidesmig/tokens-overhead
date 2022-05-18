package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
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
		routinesNumber := 50
		numRoles := 2500
		numRolesPerRoutine := numRoles / routinesNumber
		var wg sync.WaitGroup
		for routine := 0; routine < routinesNumber; routine++ {
			wg.Add(1)
			go func(routineNumber int, svc service.TokenService, reqTimes int, wg *sync.WaitGroup) {
				defer wg.Done()
				for i := 0; i < numRolesPerRoutine; i++ {
					numRoles := routineNumber*numRolesPerRoutine + i
					for requestTimes := 0; requestTimes < reqTimes; requestTimes++ {
						err = svc.Execute(numRoles, targetAddress)
						for err != nil {
							time.Sleep(2 * time.Second)
							err = svc.Execute(numRoles, targetAddress)
						}
					}
				}
			}(routine, svc, reqTimes, &wg)
		}
		wg.Wait()
		log.Print("end")
	}
}
