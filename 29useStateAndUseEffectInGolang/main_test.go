package main

import "testing"

func TestUseState(t *testing.T) {
	stateCh, getCh := useState(0)

	initialState := <-getCh

	if initialState != 0 {
		t.Errorf("Expected initial state to be 0, but got %d", initialState)
	}

	stateCh <- 10
	updatedState := <-getCh

	if initialState != 10 {
		t.Errorf("Expected updated state to be 10, but got %d", updatedState)
	}
}

func TestUseEffect(t *testing.T) {
	sideEffectCh := make(chan string, 1)

	useEffect(func() {
		sideEffectCh <- "Running side effect"
	})

	sideEffect := <-sideEffectCh
	if sideEffect != "Running side effect" {
		t.Errorf("Expected side effect message to be 'Running side effect', but got '%s'", sideEffect)
	}
}

func TestMainFunction(t *testing.T) {
	main()
}
