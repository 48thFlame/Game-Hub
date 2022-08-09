package games

import (
	"reflect"
	"testing"
)

func TestConnect4(t *testing.T) {
	game1 := Connect4Game{
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
	}
	game2 := Connect4Game{
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
	}
	game3 := Connect4Game{
		{Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
	}
	game4 := Connect4Game{
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, ConnectEmpty, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, ConnectEmpty, Connect1, Connect2, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{ConnectEmpty, Connect1, Connect2, Connect1, ConnectEmpty, ConnectEmpty, ConnectEmpty},
		{Connect1, Connect2, Connect2, Connect2, Connect1, ConnectEmpty, ConnectEmpty},
	}
	games := []Connect4Game{game1, game2, game3, game4}
	want := []bool{false, true, true, true}
	got := []bool{}

	for gameI, game := range games {
		gameNum := gameI + 1

		won := game.won(Connect1)
		if won {
			t.Logf("Game %v is a win", gameNum)
		} else {
			t.Logf("Game %v is a loss", gameNum)
		}

		got = append(got, won)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
