package routers

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/database"
	"go.mod/model"
	"go.mod/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func getEvent(id int) (model.Event, *gorm.DB) {

	var event model.Event

	dbResult := dbInstance.Find(&event, "id = ?", id)

	return event, dbResult

}

var eventCollection *mongo.Collection

func InitMongoEventCollection() {
	eventCollection = database.GetCollection(mongoClient, "event")
}

func CreateEvent(ctx *fiber.Ctx) error {

	var gottenEvent util.Event
	err := ctx.BodyParser(&gottenEvent)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	var id = primitive.NewObjectID()

	ticketIds := []string{}

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

	for i := 1; i <= gottenEvent.NumberOFTicketPrinted; i++ {
		ticketIds = append(ticketIds, PrintEmptyEventTicket(id.Hex()))
	}
	event.TicketIds = ticketIds

	_, err = eventCollection.InsertOne(goCtx, &event)

	if err != nil {
		ctx.Status(500).JSON("Something went wrong")
	}

	return ctx.Status(200).JSON(&event)
}

func GetEvent(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of event (int)")
	}

	var event model.Event

	objId, _ := primitive.ObjectIDFromHex(id)

	err := eventCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&event)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(&event)

}
func GetEventByName(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of event (int)")
	}

	var event model.Event

	objId, _ := primitive.ObjectIDFromHex(id)

	err := eventCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&event)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(&event)

}

func PutEvent(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var dbEvent model.Event

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of event (int)")
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	err := eventCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&dbEvent)

	if err != nil {
		ctx.Status(500).JSON(err.Error())
	}
	var gottenEvent util.Event

	if err := ctx.BodyParser(&gottenEvent); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if gottenEvent.NumberOFTicketPrinted != 0 {
		dbEvent.NumberOFTicketPrinted = gottenEvent.NumberOFTicketPrinted
	}
	if gottenEvent.NumberOfTicketSold != 0 {
		dbEvent.NumberOfTicketSold = gottenEvent.NumberOfTicketSold
	}
	if gottenEvent.NumberOfTicketAvailable != 0 {
		dbEvent.NumberOfTicketAvailable = gottenEvent.NumberOfTicketAvailable
	}
	if gottenEvent.EvenName != "" {
		dbEvent.EvenName = gottenEvent.EvenName
	}
	if gottenEvent.EventDescription != "" {
		dbEvent.EventDescription = gottenEvent.EventDescription
	}
	if gottenEvent.EventCoverImage != "" {
		dbEvent.EventCoverImage = gottenEvent.EventCoverImage
	}
	if gottenEvent.EventAddress != "" {
		dbEvent.EventAddress = gottenEvent.EventAddress
	}
	if gottenEvent.EventDate != "" {
		dbEvent.EventDate = gottenEvent.EventDate
	}
	if gottenEvent.TicketStartSalesDate != "" {
		dbEvent.TicketStartSalesDate = gottenEvent.TicketStartSalesDate
	}
	if gottenEvent.TicketEndSalesDate != "" {
		dbEvent.TicketEndSalesDate = gottenEvent.TicketEndSalesDate
	}
	if gottenEvent.OrganisationId != "" {
		dbEvent.OrganisationId = gottenEvent.OrganisationId
	}

	update := bson.M{"numberofticketprinted": dbEvent.NumberOFTicketPrinted,
		"numberofticketsold":      dbEvent.NumberOfTicketSold,
		"numberofticketavailable": dbEvent.NumberOfTicketAvailable,
		"evenname":                dbEvent.EvenName,
		"eventdescription":        dbEvent.EventDescription,
		"eventcoverimage":         dbEvent.EventCoverImage,
		"eventaddress":            dbEvent.EventAddress,
		"eventdate":               dbEvent.EventDate,
		"ticketstartsalesdate":    dbEvent.TicketStartSalesDate,
		"ticketendsalesdate":      dbEvent.TicketEndSalesDate,
		"organisationid":          dbEvent.OrganisationId}

	updateResult, err := eventCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if updateResult.MatchedCount != 1 {
		return ctx.Status(500).JSON("Something went wrong")
	}

	return ctx.Status(200).JSON(&dbEvent)

}
func DeleteEvent(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of event (int)")
	}

	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := eventCollection.DeleteOne(goCtx, bson.M{"_id": objID})

	if err != nil {
		ctx.Status(500).JSON(err.Error())
	}

	if result.DeletedCount < 1 {
		ctx.Status(404).JSON("No user found with id")
	}

	return ctx.Status(200).JSON("Deleted")
}

func EventRouter(app *fiber.App) {

	InitMongoEventCollection()
	app.Post("/event/", CreateEvent)
	app.Get("/event/:id", GetEvent)
	app.Get("/event/:name", GetEventByName)
	app.Put("/event/:id", PutEvent)
	app.Delete("/event/:id", DeleteEvent)

}
