package main

import (
    "flag"
    "github.com/coreos/pkg/flagutil"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

// Func to join my two strings
func join(strs ...string) string {
       var sb strings.Builder
       for _, str := range strs {
           sb.WriteString(str)
       }
       return sb.String()
    }

// Validate if requested 'name' flag is defined in cli
func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}

func main() {
    flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
    consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
    consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
    accessToken := flags.String("access-token", "", "Twitter Access Token")
    accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
    tweetMessage := flags.String("message", "", "Tweet Message")
    tweetFile := flags.String("file", "", "File Containing Tweet Message Content")
    dryRun := flags.Bool("dry", false,"Test mode, nothing will be sent to twitter")
    flags.Parse(os.Args[1:]) 
    flagutil.SetFlagsFromEnv(flags, "TWITTER")

    // Validating the credentials are available (unless dryRun)
    if (*consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "") && !*dryRun {
        log.Fatal("Consumer key/secret and Access token/secret required")
    }

    // If the flag file have been given, grab the content of the file after checking if the path exist
    if !isFlagPassed(*tweetFile) {
        _, err := os.Stat(*tweetFile)
        if err != nil {
            if os.IsNotExist(err) {
                log.Fatalf("File %s does not exist.", *tweetFile)
            } else {
                log.Fatal(err)
            }
        }
    }

    // Reading file content
    fileContent, err := ioutil.ReadFile(*tweetFile)

    // Assembling the message then the content of the file for the tweet
    tweetContent := string(join(*tweetMessage, string(fileContent)))

    // Validation a content is available and does not exeed 280 char
    if len(tweetContent) == 0 {
        log.Fatal("Your tweet is empty !")
    } else if len(tweetContent) > 280 {
        log.Fatal("Tweet must be less than 280 char")
    }

    // Posting tweet
    if *dryRun {
        log.Print("Logging in, creating client and updating status.")
    } else {
        // Setup auth
        config := oauth1.NewConfig(*consumerKey, *consumerSecret)
        token := oauth1.NewToken(*accessToken, *accessSecret)

        // http.Client will automatically authorize Requests
        httpClient := config.Client(oauth1.NoContext, token)

        // Twitter client
        client := twitter.NewClient(httpClient)
        _, _, err = client.Statuses.Update(tweetContent, nil)
        // Handling Error
        if err != nil {
            log.Fatal(err)
        }
    }

    log.Printf("Status updated with: " + tweetContent)
}
