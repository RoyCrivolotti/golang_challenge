package authentication

import (
	"context"
	"fmt"
	"golangchallenge/internal/utils"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func FirebaseInit() *auth.Client {
	firebaseKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey.json file")
	}

	opt := option.WithCredentialsFile(firebaseKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		utils.Logger.Panic(fmt.Sprintf("Error initializing Firebase app: %s", err.Error()))
		panic(err.Error())
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		utils.Logger.Panic(fmt.Sprintf("Error instantiating Firebase's Auth client: %s", err.Error()))
		panic(err.Error())
	}

	return auth
}
