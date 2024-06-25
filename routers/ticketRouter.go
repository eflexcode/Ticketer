package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mod/database"
	"go.mod/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func getTicket(id int) (model.Organisation, *gorm.DB) {

	var organisation model.Organisation

	dbResult := dbInstance.Find(&organisation, "id = ?", id)

	return organisation, dbResult

}

var ticketCollection *mongo.Collection

func InitMongoTicket() {
	ticketCollection = database.GetCollection(mongoClient, "ticket")
}

//func CreateTicket(ctx *fiber.Ctx) error {
//
//	var gottenOrganisation util.Organisation
//	err := ctx.BodyParser(&gottenOrganisation)
//
//	if err != nil {
//		return ctx.Status(400).JSON(err.Error())
//	}
//
//	timeNow := time.Time{}
//	organisation := model.Organisation{
//		CreatedAt:                   timeNow,
//		OrganisationName:            gottenOrganisation.OrganisationName,
//		OrganisationAddress:         gottenOrganisation.OrganisationAddress,
//		OrganisationProfileImageUrl: gottenOrganisation.OrganisationProfileImageUrl,
//		OrganisationOverImageUrl:    gottenOrganisation.OrganisationOverImageUrl,
//		OrganisationDescription:     gottenOrganisation.OrganisationDescription,
//		OrganisationEmail:           gottenOrganisation.OrganisationEmail,
//		OrganisationPassword:        gottenOrganisation.OrganisationPassword,
//	}
//
//	return ctx.Status(200).JSON(&organisation)
//}

func PrintEmptyEventTicket(eventId string) string {

	var id = primitive.NewObjectID()
	var ticketID = GenerateTicketId()
	//realEventId := eventId[11:34]

	ticket := model.Ticket{
		ID:       id,
		EventID:  eventId,
		TicketID: ticketID,
	}

	_, err := ticketCollection.InsertOne(goCtx, &ticket)

	if err != nil {
		log.Info("failed to insert")
	}

	return id.Hex()
}

func GenerateTicketId() string {

	const characters = "!@#$%&*1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var ticketId string = ""

	for i := 0; i <= 10; i++ {

		position := randomIntRang(1, len(characters))
		ticketId = ticketId + getCharacterAtPosition(characters, int(position))
	}

	return ticketId
}

func getCharacterAtPosition(strings string, position int) string {
	return string([]rune(strings)[position])
}

func randomIntRang(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetTicket(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of user (int)")
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	var ticket model.Ticket

	err := ticketCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&ticket)

	if err != nil {
		return ctx.Status(500).JSON("Something went wrong")
	}

	return ctx.Status(200).JSON(&ticket)

}

func PutTicket(id string, boughtBy string, boughtFor string) {

	objId, _ := primitive.ObjectIDFromHex(id)

	var ticket model.Ticket

	_ = ticketCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&ticket)

	now := time.Now()

	update := bson.M{"boughtby": boughtBy, "boughtfor": boughtFor, "buy_date": now.String()}
	_, _ = ticketCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})

}
func DeleteTicket(id string) {

	if id == "" {
		return
	}
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ticketCollection.DeleteOne(goCtx, bson.M{"_id": objID})
	if err != nil {
		return
	}
}

func TicketRouter(app *fiber.App) {
	InitMongoTicket()
	//app.Post("/ticket/", CreateTicket)
	app.Get("/ticket/:id", GetTicket)
	//app.Put("/ticket/:id", PutOrganisation)
	//app.Delete("/ticket/:id", DeleteOrganisation)

}
