package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/model"
	"go.mod/util"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func getTicket(id int) (model.Organisation, *gorm.DB) {

	var organisation model.Organisation

	dbResult := dbInstance.Find(&organisation, "id = ?", id)

	return organisation, dbResult

}

func CreateTicket(ctx *fiber.Ctx) error {

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
		OrganisationEmail:           gottenOrganisation.OrganisationEmail,
		OrganisationPassword:        gottenOrganisation.OrganisationPassword,
	}

	dbInstance.Create(&organisation)

	return ctx.Status(200).JSON(&organisation)
}

func GetTicket(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("Please insert valid id of user (int)")
	}

	organisation, db := getOrganisation(id)
	dbErr := db.Error

	if dbErr != nil {
		return ctx.Status(500).JSON("Something went wrong")
	}

	if organisation.ID == 0 {
		errMessage := "No user found with id: " + strconv.Itoa(id)
		return ctx.Status(404).JSON(errMessage)
	}

	return ctx.Status(200).JSON(&organisation)

}

func PutTicket(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("Please insert valid id of organisation (int)")
	}

	dbOrganisation, dbReturnedInstance := getOrganisation(id)
	dbErr := dbReturnedInstance.Error

	if dbErr != nil {
		return ctx.Status(500).JSON("Something went wrong")
	}

	if dbOrganisation.ID == 0 {
		return ctx.Status(404).JSON("No org found with id: " + strconv.Itoa(id))
	}

	var gottenOrg model.Organisation

	if err := ctx.BodyParser(&gottenOrg); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if gottenOrg.OrganisationName != "" {
		dbOrganisation.OrganisationName = gottenOrg.OrganisationName
	}
	if gottenOrg.OrganisationAddress != "" {
		dbOrganisation.OrganisationAddress = gottenOrg.OrganisationAddress
	}
	if gottenOrg.OrganisationProfileImageUrl != "" {
		dbOrganisation.OrganisationProfileImageUrl = gottenOrg.OrganisationProfileImageUrl
	}
	if gottenOrg.OrganisationOverImageUrl != "" {
		dbOrganisation.OrganisationOverImageUrl = gottenOrg.OrganisationOverImageUrl
	}
	if gottenOrg.OrganisationDescription != "" {
		dbOrganisation.OrganisationDescription = gottenOrg.OrganisationDescription
	}

	dbInstance.Save(&dbOrganisation)

	return ctx.Status(200).JSON(&dbOrganisation)

}

func DeleteTicket(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(400).JSON("Please insert valid id of organisation (int)")
	}

	dbOrg, db := getOrganisation(id)
	dbErr := db.Error

	if dbErr != nil {
		return ctx.Status(500).JSON("something went wrong")
	}

	if dbOrg.ID == 0 {
		return ctx.Status(404).JSON("No Org found wit id: " + strconv.Itoa(id))
	}

	db.Delete(&dbOrg)

	return ctx.Status(200).JSON("Deleted")
}

func TicketRouter(app *fiber.App) {

	app.Post("/ticket/", CreateTicket)
	app.Get("/ticket/:id", GetOrganisation)
	app.Put("/ticket/:id", PutOrganisation)
	app.Delete("/ticket/:id", DeleteOrganisation)

}
