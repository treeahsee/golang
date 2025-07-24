package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("best")
		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defintion := "this is just a test"

		dictionary.Add(word, defintion)

		assertDefinition(t, dictionary, word, defintion)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		dictionary := Dictionary{word: defintion}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, defintion)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		dictionary := Dictionary{word: defintion}
		newDefinition := "new defintion"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "this is just a test"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete word", func(t *testing.T) {
		word := "test"
		defintion := "this is just a test"
		dictionary := Dictionary{word: defintion}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)

	})

	t.Run("unkown word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q given %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}
