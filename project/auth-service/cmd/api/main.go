package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gorilla/mux"
)

var (
	userPoolClient = "11bmhphtdahuft4pdmes9qgkk0"
	awsRegion      = "ap-southeast-2"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth service"))
	}).Methods("GET")
	router.HandleFunc("/api/signup", signupHandler).Methods("POST")
	router.HandleFunc("/api/signin", signinHandler).Methods("POST")
	router.HandleFunc("/api/authenticate", authenticateHandler).Methods("POST")
	router.HandleFunc("/api/verify-email", verifyEmailHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyEmailRequest struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

type AuthenticationRequest struct {
	Token string `json:"token"`
}

// function to handle signup request
// this function will create a new user in the user pool
// and return the response from AWS Cognito
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	}))

	svc := cognitoidentityprovider.New(sess)

	params := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(userPoolClient),
		Password: aws.String(req.Password),
		Username: aws.String(req.Email),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(req.Email),
			},
		},
	}

	resp, err := svc.SignUp(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// function to handle signin request
// this function will authenticate the user and return the response from AWS Cognito
func signinHandler(w http.ResponseWriter, r *http.Request) {
	var req SigninRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	}))

	svc := cognitoidentityprovider.New(sess)

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		ClientId: aws.String(userPoolClient),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(req.Email),
			"PASSWORD": aws.String(req.Password),
		},
	}

	resp, err := svc.InitiateAuth(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// function to handle authentication request
// this function will authenticate the user using the access token
// and return the response from AWS Cognito
func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthenticationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	}))

	svc := cognitoidentityprovider.New(sess)

	params := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(req.Token),
	}

	resp, err := svc.GetUser(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// function to handle email verification request
// this function will verify the email of the user using the verification code
// and return the response from AWS Cognito
func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	}))

	svc := cognitoidentityprovider.New(sess)

	params := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(userPoolClient),
		ConfirmationCode: aws.String(req.VerificationCode),
		Username:         aws.String(req.Email),
	}

	resp, err := svc.ConfirmSignUp(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
