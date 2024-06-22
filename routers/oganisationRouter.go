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
	"time"
)

func getOrganisation(id int) (model.Organisation, *gorm.DB) {

	var organisation model.Organisation

	dbResult := dbInstance.Find(&organisation, "id = ?", id)

	return organisation, dbResult

}

var orgCollection *mongo.Collection

func InitMongoOrgCollection() {
	orgCollection = database.GetCollection(mongoClient, "organisation")
}

func CreateOrganisation(ctx *fiber.Ctx) error {

	var gottenOrganisation util.Organisation
	err := ctx.BodyParser(&gottenOrganisation)

	if err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var id = primitive.NewObjectID()

	timeNow := time.Time{}
	organisation := model.Organisation{
		ID:                          id,
		CreatedAt:                   timeNow,
		OrganisationName:            gottenOrganisation.OrganisationName,
		OrganisationAddress:         gottenOrganisation.OrganisationAddress,
		OrganisationProfileImageUrl: gottenOrganisation.OrganisationProfileImageUrl,
		OrganisationOverImageUrl:    gottenOrganisation.OrganisationOverImageUrl,
		OrganisationDescription:     gottenOrganisation.OrganisationDescription,
		OrganisationEmail:           gottenOrganisation.OrganisationEmail,
		OrganisationPassword:        gottenOrganisation.OrganisationPassword,
	}

	_, err = orgCollection.InsertOne(goCtx, &organisation)

	if err != nil {
		//log.Fatal(err.Error())
		return ctx.Status(500).JSON("Something went wrong try again")

	}

	return ctx.Status(200).JSON(&organisation)
}

func GetOrganisation(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of organisation")
	}

	var organisation model.Organisation

	objID, _ := primitive.ObjectIDFromHex(id)

	err := orgCollection.FindOne(goCtx, bson.M{"_id": objID}).Decode(&organisation)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	return ctx.Status(200).JSON(&organisation)

}

func PutOrganisation(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var organisation model.Organisation

	objId, _ := primitive.ObjectIDFromHex(id)

	err := orgCollection.FindOne(goCtx, bson.M{"_id": objId}).Decode(&organisation)

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	var gottenOrganisation model.Organisation

	if err := ctx.BodyParser(&gottenOrganisation); err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if gottenOrganisation.OrganisationName != "" {
		organisation.OrganisationName = gottenOrganisation.OrganisationName
	}
	if gottenOrganisation.OrganisationEmail != "" {
		organisation.OrganisationName = gottenOrganisation.OrganisationEmail
	}
	if gottenOrganisation.OrganisationPassword != "" {
		organisation.OrganisationName = gottenOrganisation.OrganisationPassword
	}
	if gottenOrganisation.OrganisationAddress != "" {
		organisation.OrganisationAddress = gottenOrganisation.OrganisationAddress
	}
	if gottenOrganisation.OrganisationProfileImageUrl != "" {
		organisation.OrganisationProfileImageUrl = gottenOrganisation.OrganisationProfileImageUrl
	}
	if gottenOrganisation.OrganisationOverImageUrl != "" {
		organisation.OrganisationOverImageUrl = gottenOrganisation.OrganisationOverImageUrl
	}
	if gottenOrganisation.OrganisationDescription != "" {
		organisation.OrganisationDescription = gottenOrganisation.OrganisationDescription
	}

	update := bson.M{"organisationname": organisation.OrganisationName,
		"organisationemail":           organisation.OrganisationEmail,
		"organisationpassword":        organisation.OrganisationPassword,
		"organisationaddress":         organisation.OrganisationAddress,
		"organisationprofileimageurl": organisation.OrganisationProfileImageUrl,
		"organisationoverimageurl":    organisation.OrganisationOverImageUrl,
		"organisationdescription":     organisation.OrganisationDescription}

	updateResult, err := orgCollection.UpdateOne(goCtx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return ctx.Status(500).JSON(err.Error())
	}

	if updateResult.MatchedCount != 1 {
		return ctx.Status(500).JSON("Something went wrong")
	}
	return ctx.Status(200).JSON(&organisation)

}
func DeleteOrganisation(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(400).JSON("Please insert valid id of organisation (int)")
	}

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := orgCollection.DeleteOne(goCtx, bson.M{"_id": objId})

	if err != nil {
		ctx.Status(500).JSON(err.Error())
	}

	if result.DeletedCount < 1 {
		ctx.Status(404).JSON("No user found with id")
	}

	return ctx.Status(200).JSON("Deleted")
}

func OrganisationRouter(app *fiber.App) {

	InitMongoOrgCollection()
	app.Post("/org/", CreateOrganisation)
	app.Get("/org/:id", GetOrganisation)
	app.Put("/org/:id", PutOrganisation)
	app.Delete("/org/:id", DeleteOrganisation)

}
