package store

import (
	"github.com/JulianDuniec/stockgobot/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	session *mgo.Session
)

func Init() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	ensureIndexes()
}

func SaveSymbol(symbol *models.Symbol) {
	if session == nil {
		Init()
	}
	session.DB("stock").C("symbols").Upsert(bson.M{"Symbol": symbol.Symbol}, symbol)
}

func SaveHistory(history *models.HistoricalDataPoint) {
	if session == nil {
		Init()
	}
	session.DB("stock").C("history").Upsert(bson.M{
		"Symbol": history.Symbol, "Date": history.Date},
		bson.M{"$set": bson.M{
			"Open":   history.Open,
			"Close":  history.Close,
			"High":   history.High,
			"Low":    history.Low,
			"Volume": history.Volume}})
}

func GetSymbols() []models.Symbol {
	if session == nil {
		Init()
	}
	result := []models.Symbol{}
	session.DB("stock").C("symbols").Find(bson.M{}).All(&result)
	return result
}

func ensureIndexes() {
	historyIndex := mgo.Index{
		Key:        []string{"Date", "Symbol"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     false}

	session.DB("stock").C("history").EnsureIndex(historyIndex)

	symbolsIndex := mgo.Index{
		Key:        []string{"Symbol"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     false}

	session.DB("stock").C("history").EnsureIndex(symbolsIndex)
}
