




package v1

import "context"
import "net/http"





    

        

    











    
    

    type VoiceServiceServer interface {
    

        

    
    }

    type VoiceServiceHandler struct {
        Server VoiceServiceServer
    }

    func (h *VoiceServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        switch (req.URL.Path) {
        
        case "/protocol.voice.v1.VoiceService/Connect": {
            panic("unimplemented")
        }
        
        }
    }



