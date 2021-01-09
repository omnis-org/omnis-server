package client

import (
	"fmt"

	"github.com/omnis-org/omnis-client/pkg/client_informations"

	"github.com/omnis-org/omnis-server/internal/db"
	"github.com/omnis-org/omnis-server/internal/model"
	log "github.com/sirupsen/logrus"
)

func doLocation(locationName string) (int32, error) {
	location, err := db.GetLocationByName(locationName, true)
	if err != nil {
		return 0, fmt.Errorf("GetLocationByName failed <- %v", err)
	}

	var idLocation int32
	if !location.ID.Valid {
		var name model.NullString
		name.Scan(locationName)
		idLocation, err = db.InsertLocation(&model.Location{Name: &name}, true)
		if err != nil {
			return 0, fmt.Errorf("net.InsertLocation failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("new location : %s", locationName))

	} else {
		idLocation = location.ID.Int32
	}

	return idLocation, nil
}

func doPerimeter(perimeterName string) (int32, error) {
	//perimeter
	perimeter, err := db.GetPerimeterByName(perimeterName, true)
	if err != nil {
		return 0, fmt.Errorf("net.GetPerimeterByName failed <- %v", err)
	}

	var idPerimeter int32
	if !perimeter.ID.Valid {
		var name model.NullString
		name.Scan(perimeterName)

		idPerimeter, err = db.InsertPerimeter(&model.Perimeter{Name: &name}, true)
		if err != nil {
			return 0, fmt.Errorf("net.InsertLocation failed <- %v", err)
		}

		log.Debug(fmt.Sprintf("new perimeter : %s", perimeterName))
	} else {
		idPerimeter = perimeter.ID.Int32
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
