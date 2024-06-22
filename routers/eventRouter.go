package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/model"
	"go.mod/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

func getEvent(id int) (model.Event, *gorm.DB) {

	var event model.Event

	dbResult := dbInstance.Find(&event, "id = ?", id)

	return event, dbResult

}

func CreateEvent(ctx *fiber.Ctx) error {

	var gottenEvent util.Event
	err := ctx.BodyParser(&gottenEvent)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	var id = primitive.NewObjectID()

	event := model.Event{
		ID:                      id,
		NumberOFTicketPrinted:   gottenEvent.NumberOFTicketPrinted,
		NumberOfTicketSold:      gottenEvent.NumberOfTicketSold,
		NumberOfTicketAvailable: gottenEvent.NumberOfTicketAvailable,
		EvenName:                gottenEvent.EvenName,
		EventDescription:        gottenEvent.EventDescription,
		EventCoverImage:         gottenEvent.EventCoverImage,
		EventAddress:            gottenEvent.EventAddress,
		EventDate:               gottenEvent.EventDate,
		TicketStartSalesDate:    gottenEvent.TicketStartSalesDate,
		TicketEndSalesDate:      gottenEvent.TicketEndSalesDate,
		OrganisationId:          gottenEvent.OrganisationId,
	}

	//fmt.Printf("db test type " + strconv.Itoa(int(event.ID)))

	return ctx.Status(200).JSON(&event)
}

//func GetEvent(ctx *fiber.Ctx) error {
//
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of event (int)")
//	}
//
//	event, db := getEvent(id)
//	dbErr := db.Error
//
//	if dbErr != nil {
//		return ctx.Status(500).JSON("Something went wrong")
//	}
//
//	if event.ID == 0 {
//		errMessage := "No event found with id: " + strconv.Itoa(id)
//		return ctx.Status(404).JSON(errMessage)
//	}
//
//	return ctx.Status(200).JSON(&event)
//
//}

//func PutEvent(ctx *fiber.Ctx) error {
//
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of event (int)")
//	}
//
//	dbOrganisation, dbReturnedInstance := getOrganisation(id)
//	dbErr := dbReturnedInstance.Error
//
//	if dbErr != nil {
//		return ctx.Status(500).JSON("Something went wrong")
//	}
//
//	if dbOrganisation.ID == 0 {
//		return ctx.Status(404).JSON("No event found with id: " + strconv.Itoa(id))
//	}
//
//	var gottenOrg model.Organisation
//
//	if err := ctx.BodyParser(&gottenOrg); err != nil {
//		return ctx.Status(500).JSON(err.Error())
//	}
//
//	dbInstance.Save(&dbOrganisation)
//
//	return ctx.Status(200).JSON(&dbOrganisation)
//
//}
//
//func DeleteEvent(ctx *fiber.Ctx) error {
//
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of event (int)")
//	}
//
//	dbOrg, db := getEvent(id)
//	dbErr := db.Error
//
//	if dbErr != nil {
//		return ctx.Status(500).JSON("something went wrong")
//	}
//
//	if dbOrg.ID == 0 {
//		return ctx.Status(404).JSON("No event found wit id: " + strconv.Itoa(id))
//	}
//
//	db.Delete(&dbOrg)
//
//	return ctx.Status(200).JSON("Deleted")
//}

func EventRouter(app *fiber.App) {

	app.Post("/event/", CreateEvent)
	app.Get("/event/:id", GetEvent)
	//app.Put("/event/:id", PutEvent)
	//app.Delete("/event/:id", DeleteEvent)

}
