package worker

import (
	"fmt"

	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"github.com/omnis-org/omnis-rest-api/pkg/model"
	"github.com/omnis-org/omnis-server/internal/net"
	log "github.com/sirupsen/logrus"
)

func doLocation(locationName string) (int32, error) {
	location, err := net.GetLocationByName(locationName)
	if err != nil {
		return 0, fmt.Errorf("GetLocationByName failed <- %v", err)
	}

	var idLocation int32
	if !location.Id.Valid {
		log.Info("Create new location : ", locationName)
		var name model.NullString
		name.Scan(locationName)

		idLocation, err = net.InsertLocation(&model.Location{Name: name})

		if err != nil {
			return 0, fmt.Errorf("net.InsertLocation failed <- %v", err)
		}

	} else {
		idLocation = location.Id.Int32
	}

	return idLocation, nil
}

func doPerimeter(perimeterName string) (int32, error) {
	//perimeter
	perimeter, err := net.GetPerimeterByName(perimeterName)
	if err != nil {
		return 0, fmt.Errorf("net.GetPerimeterByName failed <- %v", err)
	}

	var idPerimeter int32
	if !perimeter.Id.Valid {
		log.Info("Create new perimeter : ", perimeterName)
		var name model.NullString
		name.Scan(perimeterName)

		idPerimeter, err = net.InsertPerimeter(&model.Perimeter{Name: name})

		if err != nil {
			return 0, fmt.Errorf("net.InsertLocation failed <- %v", err)
		}

	} else {
		idPerimeter = perimeter.Id.Int32
	}

	return idPerimeter, nil
}

func doOtherInformations(otherInformation *client_informations.OtherInformations) (int32, int32, error) {
	//location

	idLocation, err := doLocation(otherInformation.Location)
	if err != nil {
		return 0, 0, fmt.Errorf("doLocation failed <- %v", err)
	}

	idPerimeter, err := doPerimeter(otherInformation.Perimeter)
	if err != nil {
		return 0, 0, fmt.Errorf("doPerimeter failed <- %v", err)
	}

	return idLocation, idPerimeter, nil
}
