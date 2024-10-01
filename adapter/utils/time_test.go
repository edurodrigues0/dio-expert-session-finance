package utils

import "testing"

func TestStringToTime(t *testing.T) {
	convertedTime := StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("Espera que o ano seja 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Espera que o mes seja 02")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espera que a hora seja 10h")
	}
}
