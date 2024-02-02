package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

// generateV4GetObjectSignedURL generates object signed URL with GET method.
func generateV4GetObjectSignedURL(w io.Writer, bucket, object string) (string, error) {
    // bucket := "bucket-name"
    // object := "object-name"

    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
            return "", fmt.Errorf("storage.NewClient: %w", err)
    }
    defer client.Close()

    // Signing a URL requires credentials authorized to sign a URL. You can pass
    // these in through SignedURLOptions with one of the following options:
    //    a. a Google service account private key, obtainable from the Google Developers Console
    //    b. a Google Access ID with iam.serviceAccounts.signBlob permissions
    //    c. a SignBytes function implementing custom signing.
    // In this example, none of these options are used, which means the SignedURL
    // function attempts to use the same authentication that was used to instantiate
    // the Storage client. This authentication must include a private key or have
    // iam.serviceAccounts.signBlob permissions.
    data, err := os.ReadFile("secret/service_account.json")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        // Handle error
    }

    conf, err := google.JWTConfigFromJSON(data)
    if err != nil {
        fmt.Printf("Error creating JWTConfig: %v\n", err)
    }
    
    opts := &storage.SignedURLOptions{
            Scheme:  storage.SigningSchemeV4,
            Method:  "GET",
            Expires: time.Now().Add(3 * time.Minute),
            GoogleAccessID: conf.Email,
            PrivateKey:     conf.PrivateKey,

    }

    u, err := client.Bucket(bucket).SignedURL(object, opts)
    if err != nil {
            return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
    }

    fmt.Fprintln(w, "Generated GET signed URL:")
    fmt.Fprintf(w, "%q\n", u)
    fmt.Fprintln(w, "You can use this URL with any user agent, for example:")
    fmt.Fprintf(w, "curl %q\n", u)
    return u, nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

func requestSignedURLHandler(w http.ResponseWriter, r *http.Request) {
    bucket := "go-cloud-bucket-playground"
    object := "gattino.jpeg"


    url, err := generateV4GetObjectSignedURL(w, bucket, object)
    if err != nil {
        http.Error(w, fmt.Sprintf("generateV4GetObjectSignedURL: %v", err), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, url)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/requestSignedURL", requestSignedURLHandler)
    http.ListenAndServe(":10000", nil)
}
