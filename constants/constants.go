package constants

import (
	"fmt"
	"os"
)

var (
	mongoUser = os.Getenv("MONGO_USER")
	mongoPass = os.Getenv("MONGO_PASS")
	MongoURI  = fmt.Sprintf("mongodb+srv://%s:%s@cluster-px3sbnfz.g9x7l.mongodb.net/?retryWrites=true&w=majority", mongoUser, mongoPass)
)
