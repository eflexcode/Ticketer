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

// func PutTicket(ctx *fiber.Ctx) error {
//
//	id, err := ctx.ParamsInt("id")
//
//	if err != nil {
//		return ctx.Status(400).JSON("Please insert valid id of organisation (int)")
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
//		return ctx.Status(404).JSON("No org found with id: " + strconv.Itoa(id))
//	}
//
//	var gottenOrg model.Organisation
//
//	if err := ctx.BodyParser(&gottenOrg); err != nil {
//		return ctx.Status(500).JSON(err.Error())
//	}
//
//	if gottenOrg.OrganisationName != "" {
//		dbOrganisation.OrganisationName = gottenOrg.OrganisationName
//	}
//	if gottenOrg.OrganisationAddress != "" {
//		dbOrganisation.OrganisationAddress = gottenOrg.OrganisationAddress
//	}
//	if gottenOrg.OrganisationProfileImageUrl != "" {
//		dbOrganisation.OrganisationProfileImageUrl = gottenOrg.OrganisationProfileImageUrl
//	}
//	if gottenOrg.OrganisationOverImageUrl != "" {
//		dbOrganisation.OrganisationOverImageUrl = gottenOrg.OrganisationOverImageUrl
//	}
//	if gottenOrg.OrganisationDescription != "" {
//		dbOrganisation.OrganisationDescription = gottenOrg.OrganisationDescription
//	}
//
//	dbInstance.Save(&dbOrganisation)
//
//	return ctx.Status(200).JSON(&dbOrganisation)
//
// }
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
