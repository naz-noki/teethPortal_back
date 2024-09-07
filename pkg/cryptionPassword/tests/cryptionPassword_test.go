package cryptionPassword_test

import (
	"MySotre/pkg/cryptionPassword"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Start TestMain", "package: cryptionPassword_test")
	m.Run()
	log.Println("End TestMain", "package: cryptionPassword_test")
}

func TestDecode(t *testing.T) {
	t.Parallel()

	pass1, salt1, secondSalt1 := "123qwerrt", "Hkjsldhgf8w7egfn87n", "JKHbg297832gbnvsS221d3bgjhfgb"
	pass2, salt2, secondSalt2 := "alskjdfh", "JKHABJHSDGAJSdgb*&897taSBDJ", "sJDFgBUGAYSGBf87tb2asdASFTBN2"
	pass3, salt3, secondSalt3 := "3416924hf", "OIASd98yq892nmIUAWNM98", "ADWUKGHNQ13W7dytq87wtdq897wtd"

	t.Run("TestDecode - success - 1", func(t *testing.T) {
		p := cryptionPassword.Encode(pass1, salt1, secondSalt1)
		result, err := cryptionPassword.Decode(p, salt1, secondSalt1)

		if result != pass1 || err != nil {
			t.Errorf("The passwords don't match. Password: %s Result: %s", pass1, result)
		}
	})
	t.Run("TestDecode - error - 2", func(t *testing.T) {
		result, err := cryptionPassword.Decode(pass2, salt2, secondSalt2)

		if result == pass2 || err == nil {
			t.Errorf("The passwords match. Password: %s Result: %s", pass2, result)
		}
	})
	t.Run("TestDecode - success - 3", func(t *testing.T) {
		p := cryptionPassword.Encode(pass3, salt3, secondSalt3)
		result, err := cryptionPassword.Decode(p, salt3, secondSalt3)

		if result != pass3 || err != nil {
			t.Errorf("The passwords match. Password: %s Result: %s", pass3, result)
		}
	})
}
