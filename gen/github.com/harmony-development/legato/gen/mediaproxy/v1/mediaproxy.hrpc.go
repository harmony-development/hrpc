




package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"









    
    

    type MediaProxyServiceClient struct {
        client *http.Client
        serverURL string
    }

    func NewMediaProxyServiceClient(url string) *MediaProxyServiceClient {
        return &MediaProxyServiceClient{
            client:    &http.Client{},
            serverURL: url,
        }
    }

    

    
        
        

        
        

        // Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

        func (client *MediaProxyServiceClient) FetchLinkMetadata(r *v1.FetchLinkMetadataRequest) (*v1.FetchLinkMetadataRequest, error) {
            input, err := proto.Marshal(r)
            if err != nil {
                return nil, fmt.Errorf("could not martial request: %w", err)
            }
            resp, err := client.client.Post(fmt.Sprintf("%s/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata", client.serverURL), "application/octet-stream", bytes.NewReader(input))
            if err != nil {
                return nil, fmt.Errorf("error posting request: %w", err)
            }
            defer resp.Body.Close()
            data, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                return nil, fmt.Errorf("error reading response: %w", err)
            }
            output := &v1.FetchLinkMetadataRequest{}
            err = proto.Unmarshal(data, output)
            if err != nil {
                return nil, fmt.Errorf("error unmarshalling response: %w", err)
            }
            return output, nil
        }
    

    

    
        
        

        
        

        // Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

        func (client *MediaProxyServiceClient) InstantView(r *v1.InstantViewRequest) (*v1.InstantViewRequest, error) {
            input, err := proto.Marshal(r)
            if err != nil {
                return nil, fmt.Errorf("could not martial request: %w", err)
            }
            resp, err := client.client.Post(fmt.Sprintf("%s/protocol.mediaproxy.v1.MediaProxyService/InstantView", client.serverURL), "application/octet-stream", bytes.NewReader(input))
            if err != nil {
                return nil, fmt.Errorf("error posting request: %w", err)
            }
            defer resp.Body.Close()
            data, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                return nil, fmt.Errorf("error reading response: %w", err)
            }
            output := &v1.InstantViewRequest{}
            err = proto.Unmarshal(data, output)
            if err != nil {
                return nil, fmt.Errorf("error unmarshalling response: %w", err)
            }
            return output, nil
        }
    

    

    
        
        

        
        

        // Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

        func (client *MediaProxyServiceClient) CanInstantView(r *v1.InstantViewRequest) (*v1.InstantViewRequest, error) {
            input, err := proto.Marshal(r)
            if err != nil {
                return nil, fmt.Errorf("could not martial request: %w", err)
            }
            resp, err := client.client.Post(fmt.Sprintf("%s/protocol.mediaproxy.v1.MediaProxyService/CanInstantView", client.serverURL), "application/octet-stream", bytes.NewReader(input))
            if err != nil {
                return nil, fmt.Errorf("error posting request: %w", err)
            }
            defer resp.Body.Close()
            data, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                return nil, fmt.Errorf("error reading response: %w", err)
            }
            output := &v1.InstantViewRequest{}
            err = proto.Unmarshal(data, output)
            if err != nil {
                return nil, fmt.Errorf("error unmarshalling response: %w", err)
            }
            return output, nil
        }
    

    





import "github.com/harmony-development/legato/gen/mediaproxy/v1"


