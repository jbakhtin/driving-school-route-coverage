package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
	"github.com/jbakhtin/driving-school-route-coverage/internal/domain/repositories"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres"
	"github.com/jbakhtin/driving-school-route-coverage/internal/infrastructure/database/postgres/repository"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
	"log"
)

func main() {

	// Создание линии с использованием точек
	lineString := geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{
		[]float64{1.0, 2.0},
		[]float64{1.0, 2.0},
		[]float64{1.0, 2.0},
	})


	test, _ := wkb.Marshal(lineString, binary.LittleEndian)

	routeCreation := repositories.RouteCreation{
		LineString: test,
	}

	cfg, _ := config.GetConfig()
	client, _ := postgres.New(*cfg)
	repo, _ := repository.NewRouteRepository(client)

	route, err := repo.CreateRoute(context.TODO(), routeCreation)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(route)

	stringJson := string(json)

	fmt.Println(stringJson)
}
