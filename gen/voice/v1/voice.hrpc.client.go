




package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "bytes"





    

        

    











    
    

    type VoiceServiceClient struct {
        client *http.Client
        serverURL string
    }

    func NewVoiceServiceClient(url string) *VoiceServiceClient {
        return &VoiceServiceClient{
            client:    &http.Client{},
            serverURL: url,
        }
    }

    

        

    



