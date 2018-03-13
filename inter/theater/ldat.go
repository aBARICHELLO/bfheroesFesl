package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

// Lobbies Data
type ansLDAT struct {
	TID             string `fesl:"TID"`
	FavoriteGames   string `fesl:"FAVORITE-GAMES"`
	FavoritePlayers string `fesl:"FAVORITE-PLAYERS"`
	LobbyID         string `fesl:"LID"`
	Locale          string `fesl:"LOCALE"`
	MaxGames        string `fesl:"MAX-GAMES"`
	Name            string `fesl:"NAME"`
	NumGames        string `fesl:"NUM-GAMES"`
	Passing         string `fesl:"PASSING"`
}

func (tm *Theater) LobbyData(event network.EventClientProcess) {
	logrus.Println("===LDAT===")
	event.Client.Answer(&codec.Pkt{
		Type: thtrLDAT,
		Content: ansLDAT{
			TID:             "1",
			FavoriteGames:   "0",
			FavoritePlayers: "0",
			LobbyID:         "1",
			Locale:          "en_US",
			MaxGames:        "10000",
			Name:            "bfwestPC02",
			NumGames:        "1",
			Passing:         "0",
		},
	})
}