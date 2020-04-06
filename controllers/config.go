package controllers

import "go.mongodb.org/mongo-driver/mongo"

// activity collection for the db
var activityCollection *mongo.Collection

// calorie collection for the db
var calorieCollection *mongo.Collection

// collection for the db
var heartCollection *mongo.Collection

// collection for the db
var powerCollection *mongo.Collection

// collection for the db
var stepCollection *mongo.Collection
