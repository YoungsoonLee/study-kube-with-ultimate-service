package auth_test

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"testing"

	"github.com/YoungsoonLee/study-kube-with-ultimate-service/business/auth"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestAuth(t *testing.T) {
	t.Log("Given the need to be able to authenticate and authorize access.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen handling a single user", testID)
		{
			privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				log.Fatalln(err)
			}

			const KeyID = "1234567890-0987654321"
			lookup := func(kid string) (*rsa.PublicKey, error) {
				switch kid {
				case KeyID:
					return &privateKey.PublicKey, nil
				}
				return nil, fmt.Errorf("no public key found for the specified kid: %s", kid)
			}

			a, err := auth.New("RS256", lookup, auth.Keys{KeyID: privateKey})
			if err != nil {
				//t.Fatalf()
			}
			t.Logf("%v", a)
		}
	}
}
