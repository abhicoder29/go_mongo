package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// guuid "github.com/google/uuid"
)

type ProduceAsset struct {
	ID         string `json:"_id"`
	LASTNUMBER string `json:"LASTNUMBER"`
	PRODUCEID  string `json:"PRODUCEID"`
	PRODUCE    string `json:"PRODUCE" binding:"required"`
	VARIETY    []struct {
		Name string `json:"name"`
		Qty  int    `json:"qty"`
	} `json:"VARIETY"`
	FARMLOCATION string    `json:"FARMLOCATION"`
	EXTENT       int       `json:"EXTENT"`
	PLANTINGDATE time.Time `json:"PLANTINGDATE"`
	GAP          struct {
		TYPE    string `json:"TYPE"`
		DETAILS struct {
			STATUS    string `json:"STATUS"`
			INTERSTED bool   `json:"INTERSTED"`
		} `json:"DETAILS"`
	} `json:"GAP"`
	MAINID   string `json:"MAINID" binding:"required"`
	TIMELINE struct {
		REGISTRATION struct {
			STATUS string `json:"STATUS"`
			DATE   string `json:"DATE"`
		} `json:"REGISTRATION"`
		CACAS struct {
			STATUS string `json:"STATUS"`
		} `json:"CACAS"`
		HARVEST struct {
			STATUS string `json:"STATUS"`
		} `json:"HARVEST"`
		INSPECTION struct {
			STATUS string `json:"STATUS"`
		} `json:"INSPECTION"`
		CERTIFY struct {
			STATUS string `json:"STATUS"`
		} `json:"CERTIFY"`
		BINS struct {
			STATUS string `json:"STATUS"`
		} `json:"BINS"`
		SIGNOFF struct {
			STATUS string `json:"STATUS"`
		} `json:"SIGNOFF"`
	} `json:"TIMELINE"`
	SELECTEDUNIT struct {
		NAME   string `json:"NAME"`
		WEIGHT string `json:"WEIGHT"`
	} `json:"SELECTED_UNIT"`
	TOTALAVAILABILITY []struct {
		Name string `json:"name"`
		Qty  int    `json:"qty"`
	} `json:"TOTAL_AVAILABILITY"`
	BASEUNIT         string    `json:"BASE_UNIT"`
	ICON             string    `json:"ICON"`
	COUNTRY          string    `json:"COUNTRY"`
	STATE            string    `json:"STATE"`
	STATUS           string    `json:"STATUS"`
	ENDDATE          string    `json:"END_DATE"`
	CREATEDATE       time.Time `json:"CREATE_DATE"`
	CREATEBY         string    `json:"CREATE_BY"`
	MODIFIEDDATE     time.Time `json:"MODIFIED_DATE"`
	MODIFIEDBY       string    `json:"MODIFIED_BY"`
	LASTLOGGEDIN     string    `json:"LAST_LOGGED_IN"`
	APPROXHARVESTING time.Time `json:"APPROXHARVESTING"`
	NOTES            []struct {
		ItemID          int    `json:"itemId"`
		NoteTitle       string `json:"noteTitle"`
		NoteDescription string `json:"noteDescription"`
		NoteImages      []struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"noteImages"`
	} `json:"NOTES"`
}


// DATABASE INSTANCE
var collection *mongo.Collection

func ProduceCollection(c *mongo.Database) {
	collection = c.Collection("Produce")
}

func GetAllProduces(c *gin.Context) {
	ProduceAssets := []ProduceAsset{}
	cursor, err := collection.Find(context.PRODUCE(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all ProduceAssets, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.PRODUCE()) {
		var produceAsset ProduceAsset
		cursor.Decode(&produceAsset)
		ProduceAssets = append(ProduceAssets, produceAsset)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Produces",
		"data":    ProduceAssets,
	})
	return
}

func CreateProduce(c *gin.Context) {
	var produceAsset ProduceAsset
	c.BindJSON(&produceAsset)
	LASTNUMBER := produceAsset.LASTNUMBER
	PRODUCEID := produceAsset.PRODUCEID
	PRODUCE := produceAsset.PRODUCE
	// id := guuid.New().String()
	newProduce := ProduceAsset{
		LASTNUMBER:     LASTNUMBER,
		PRODUCEID:      PRODUCEID,
		PRODUCE: PRODUCE,

	}
	_, err := collection.InsertOne(context.PRODUCE(), newProduce)
	if err != nil {
	log.Printf("Error while inserting new produceAsset into db, Reason: %v\n", err)
	c.JSON(http.StatusInternalServerError, gin.H{
	"status":  http.StatusInternalServerError,
	"message": "Something went wrong",
	})
	return
	}
	c.JSON(http.StatusCreated, gin.H{
	"status":  http.StatusCreated,
	"message": "ProduceAsset created Successfully",
	})
	return
	}


func GetSingleProduce(c *gin.Context) {
	PRODUCEID := c.Param("PRODUCEID")

	produceAsset := ProduceAsset{}
	err := collection.FindOne(context.PRODUCE(), bson.M{"PRODUCEID": PRODUCEID}).Decode(&produceAsset)
	if err != nil {
			log.Printf("Error while getting a single produceAsset, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "ProduceAsset not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single ProduceAsset",
		"data": produceAsset,
	})
	return
}

func UpdateProduce(c *gin.Context) {
	PRODUCEID := c.Param("PRODUCEID")
	var produceAsset ProduceAsset
	c.BindJSON(&produceAsset)
	PRODUCE := produceAsset.PRODUCE

	newData := bson.M{
            "$set": bson.M{
            "PRODUCE": PRODUCE,
            },
        }

	_, err := collection.UpdateOne(context.PRODUCE(), bson.M{"PRODUCEID": PRODUCEID}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message":  "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "ProduceAsset Edited Successfully",
	})
	return
}

func DeleteProduce(c *gin.Context) {
PRODUCEID := c.Param("PRODUCEID")

	_, err := collection.DeleteOne(context.PRODUCE(), bson.M{"PRODUCEID": PRODUCEID})
	if err != nil {
		log.Printf("Error while deleting a single produceAsset, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ProduceAsset deleted successfully",
	})
	return
}
