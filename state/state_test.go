package state

import "testing"

func TestViewState(t *testing.T) {
    wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
    state := ViewState();
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)

    }

}

func TestPutInBoat(t *testing.T) {
	wanted := "rev"
	state := PutInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "test"
	state := CrossRiver()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}