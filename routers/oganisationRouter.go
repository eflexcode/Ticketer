package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/model"
	"go.mod/util"
	"time"
)

func CreateOrganisation(ctx *fiber.Ctx) error {

	var gottenOrganisation util.Organisation
	err := ctx.BodyParser(&gottenOrganisation)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	timeNow := time.Time{}
	organisation := model.Organisation{
		CreatedAt:                   timeNow,
		OrganisationName:            gottenOrganisation.OrganisationName,
		OrganisationAddress:         gottenOrganisation.OrganisationAddress,
		OrganisationProfileImageUrl: gottenOrganisation.OrganisationProfileImageUrl,
		OrganisationOverImageUrl:    gottenOrganisation.OrganisationOverImageUrl,
		OrganisationDescription:     gottenOrganisation.OrganisationDescription,
	}

	dbInstance.Create(&organisation)

	return ctx.Status(200).JSON(&organisation)
}

func OrganisationRouter(app *fiber.App) {

	app.Post("/org/", CreateOrganisation)

}
