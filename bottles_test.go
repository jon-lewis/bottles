package bottles

import "testing"

func TestTheFirstVerse(t *testing.T) {
	expected := "99 bottles of beer on the wall, " +
		"99 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"98 bottles of beer on the wall.\n"

	actual := Bottle{}.verse(99)

	if actual != expected {
		t.Error("Actual value did not equal expected", actual)
	}
}

func TestAnotherVerse(t *testing.T) {
	expected := "3 bottles of beer on the wall, " +
		"3 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"2 bottles of beer on the wall.\n"

	actual := Bottle{}.verse(3)

	if actual != expected {
		t.Error("Actual value did not equal expected", actual)
	}
}

func TestVerseTwo(t *testing.T) {
	expected := "2 bottles of beer on the wall, " +
		"2 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"1 bottle of beer on the wall.\n"

	actual := Bottle{}.verse(2)

	if actual != expected {
		t.Error("Actual value did not equal expected", actual)
	}
}

func TestVerseOne(t *testing.T) {
	expected := "1 bottle of beer on the wall, " +
		"1 bottle of beer.\n" +
		"Take it down and pass it around, " +
		"no more bottles of beer on the wall.\n"

	actual := Bottle{}.verse(1)

	if actual != expected {
		t.Error("Actual value did not equal expected", actual)
	}
}

func TestVerseZero(t *testing.T) {
	expected := "No more bottles of beer on the wall, " +
		"no more bottles of beer.\n" +
		"Go to the store and buy some more, " +
		"99 bottles of beer on the wall.\n"

	actual := Bottle{}.verse(0)

	if actual != expected {
		t.Error("Actual value did not equal expected", actual)
	}
}

func TestACoupleVerses(t *testing.T) {
	expected := "99 bottles of beer on the wall, " +
		"99 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"98 bottles of beer on the wall.\n" +
		"\n" +
		"98 bottles of beer on the wall, " +
		"98 bottles of beer.\n" +
		"Take one down and pass it around, " +
		"97 bottles of beer on the wall.\n"

	actual := Bottle{}.verses(99, 98)

	if actual != expected {
		t.Error("Actual value did not equal expected. Actual:", actual, "\nExpected:", expected)
	}
}
