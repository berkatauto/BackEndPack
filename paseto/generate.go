package paseto

import (
	"fmt"

	"github.com/o1egl/paseto"
)

func generateToken() {
	privateKey, publicKey, err := paseto.generate(paseto.V2)
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return
	}
	payload := map[string]interface{}{
		"user_id":   0012345,
		"user_name": "admin",
	}
	// Creating a PASETO token
	token, err := paseto.NewV2().Sign(privateKey, payload, nil)
	if err != nil {
		fmt.Println("Error creating PASETO token:", err)
		return
	}

	fmt.Println("PASETO Token:", token)

	// Verifying the token
	var verifiedPayload map[string]interface{}
	err = paseto.NewV2().Verify(token, publicKey, &verifiedPayload, nil)
	if err != nil {
		fmt.Println("Token verification failed:", err)
		return
	}

	fmt.Println("Verified Payload:", verifiedPayload)
}
