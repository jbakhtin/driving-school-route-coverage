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

	point1 := geom.NewPoint(geom.XY).MustSetCoords([]float64{1.0, 2.0})
	point2 := geom.NewPoint(geom.XY).MustSetCoords([]float64{3.0, 4.0})
	point3 := geom.NewPoint(geom.XY).MustSetCoords([]float64{5.0, 6.0})
	// Создание линии с использованием точек
	lineString := geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{
		point1.Coords(),
		point2.Coords(),
		point3.Coords(),
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
