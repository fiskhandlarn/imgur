package imgur

import (
    "bytes"
    "errors"
    "fmt"
    "net/http"

    "github.com/mattn/go-scan"
    "golang.org/x/oauth2"
)

const (
    endpoint = "https://api.imgur.com/3/image"
)

// oauth configuration
var config = &oauth2.Config{
    ClientID:     "16958ad0bd36ae8",
    Endpoint: oauth2.Endpoint{
        AuthURL:  "https://api.imgur.com/oauth2/authorize",
        TokenURL: "https://api.imgur.com/oauth2/token",
    },
}

//func Upload(imageContents multipart.File, anonymous bool) {
func Upload(imageContents []byte, bearer *string) (string, error) {
    var res *http.Response

    req, err := http.NewRequest("POST", endpoint, bytes.NewReader(imageContents))
    if err != nil {
        fmt.Println("post1:", err.Error());
        //fmt.Fprintln(os.Stderr, "post:", err.Error())
        return "", errors.New("Post error: " + err.Error())
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    if bearer == nil {
        req.Header.Set("Authorization", "Client-ID "+config.ClientID)
    } else {
        req.Header.Set("Authorization", "Bearer " + *bearer)
    }

    res, err = http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("post2:", err.Error());
        //fmt.Fprintln(os.Stderr, "post:", err.Error())
        return "", errors.New("Post error: " + err.Error())
    }

    if res.StatusCode != 200 {
        var message string
        err := scan.ScanJSON(res.Body, "data/error", &message)
        if err != nil {
            message = res.Status
            fmt.Println("post3:", err.Error(), res);
            // fmt.Fprintln(os.Stderr, "post:", message)
            // fmt.Fprintln(os.Stderr, res)
            return "", errors.New("Post error: " + err.Error())
        }
    }
    defer res.Body.Close()

    var link string
    err = scan.ScanJSON(res.Body, "data/link", &link)
    if err != nil {
        fmt.Println("post4:", err.Error());
        //fmt.Fprintln(os.Stderr, "post:", err.Error())
        return "", errors.New("Post error: " + err.Error())
    }

    return link, nil;
}
