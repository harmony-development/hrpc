package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"
import "google.golang.org/protobuf/types/descriptorpb"

import "github.com/golang/protobuf/ptypes/empty"

var Authᐳv1ᐳauth *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x12auth/v1/auth.proto\x12\x10protocol.auth.v1\x1a\x1bgoogle/protobuf/empty.proto\",\n\x11BeginAuthResponse\x12\x17\n\aauth_id\x18\x01 \x01(\tR\x06authId\"K\n\aSession\x12\x1b\n\auser_id\x18\x01 \x01(\x04B\x020\x01R\x06userId\x12#\n\rsession_token\x18\x02 \x01(\tR\fsessionToken\"\xd4\x04\n\bAuthStep\x12!\n\ffallback_url\x18\x01 \x01(\tR\vfallbackUrl\x12\x1e\n\vcan_go_back\x18\x02 \x01(\bR\tcanGoBack\x12;\n\x06choice\x18\x03 \x01(\v2!.protocol.auth.v1.AuthStep.ChoiceH\x00R\x06choice\x125\n\x04form\x18\x04 \x01(\v2\x1f.protocol.auth.v1.AuthStep.FormH\x00R\x04form\x125\n\asession\x18\x05 \x01(\v2\x19.protocol.auth.v1.SessionH\x00R\asession\x12>\n\awaiting\x18\x06 \x01(\v2\".protocol.auth.v1.AuthStep.WaitingH\x00R\awaiting\x1a8\n\x06Choice\x12\x14\n\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n\aoptions\x18\x02 \x03(\tR\aoptions\x1a\x94\x01\n\x04Form\x12\x14\n\x05title\x18\x01 \x01(\tR\x05title\x12A\n\x06fields\x18\x02 \x03(\v2).protocol.auth.v1.AuthStep.Form.FormFieldR\x06fields\x1a3\n\tFormField\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n\x04type\x18\x02 \x01(\tR\x04type\x1aA\n\aWaiting\x12\x14\n\x05title\x18\x01 \x01(\tR\x05title\x12 \n\vdescription\x18\x02 \x01(\tR\vdescriptionB\x06\n\x04step\"\x87\x03\n\x0fNextStepRequest\x12\x17\n\aauth_id\x18\x01 \x01(\tR\x06authId\x12B\n\x06choice\x18\x02 \x01(\v2(.protocol.auth.v1.NextStepRequest.ChoiceH\x00R\x06choice\x12<\n\x04form\x18\x03 \x01(\v2&.protocol.auth.v1.NextStepRequest.FormH\x00R\x04form\x1a \n\x06Choice\x12\x16\n\x06choice\x18\x01 \x01(\tR\x06choice\x1aa\n\nFormFields\x12\x16\n\x05bytes\x18\x01 \x01(\fH\x00R\x05bytes\x12\x18\n\x06string\x18\x02 \x01(\tH\x00R\x06string\x12\x18\n\x06number\x18\x03 \x01(\x03H\x00R\x06numberB\a\n\x05field\x1aL\n\x04Form\x12D\n\x06fields\x18\x01 \x03(\v2,.protocol.auth.v1.NextStepRequest.FormFieldsR\x06fieldsB\x06\n\x04step\"*\n\x0fStepBackRequest\x12\x17\n\aauth_id\x18\x01 \x01(\tR\x06authId\"-\n\x12StreamStepsRequest\x12\x17\n\aauth_id\x18\x01 \x01(\tR\x06authId\")\n\x0fFederateRequest\x12\x16\n\x06target\x18\x01 \x01(\tR\x06target\";\n\rFederateReply\x12\x14\n\x05token\x18\x01 \x01(\tR\x05token\x12\x14\n\x05nonce\x18\x02 \x01(\tR\x05nonce\"\x1c\n\bKeyReply\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\"N\n\x15LoginFederatedRequest\x12\x1d\n\nauth_token\x18\x01 \x01(\tR\tauthToken\x12\x16\n\x06domain\x18\x02 \x01(\tR\x06domain2\xa1\x04\n\vAuthService\x12N\n\bFederate\x12!.protocol.auth.v1.FederateRequest\x1a\x1f.protocol.auth.v1.FederateReply\x12T\n\x0eLoginFederated\x12'.protocol.auth.v1.LoginFederatedRequest\x1a\x19.protocol.auth.v1.Session\x129\n\x03Key\x12\x16.google.protobuf.Empty\x1a\x1a.protocol.auth.v1.KeyReply\x12H\n\tBeginAuth\x12\x16.google.protobuf.Empty\x1a#.protocol.auth.v1.BeginAuthResponse\x12I\n\bNextStep\x12!.protocol.auth.v1.NextStepRequest\x1a\x1a.protocol.auth.v1.AuthStep\x12I\n\bStepBack\x12!.protocol.auth.v1.StepBackRequest\x1a\x1a.protocol.auth.v1.AuthStep\x12Q\n\vStreamSteps\x12$.protocol.auth.v1.StreamStepsRequest\x1a\x1a.protocol.auth.v1.AuthStep0\x01B3Z1github.com/harmony-development/legato/gen/auth/v1J\xda#\n\a\x12\x05\x00\x00\x82\x01\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n\b\n\x01\x02\x12\x03\x04\x00\x19\n\b\n\x01\b\x12\x03\x06\x00H\n\t\n\x02\b\v\x12\x03\x06\x00H\n\t\n\x02\x04\x00\x12\x03\b\x001\n\n\n\x03\x04\x00\x01\x12\x03\b\b\x19\n\v\n\x04\x04\x00\x02\x00\x12\x03\b\x1c/\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\b\x1c\"\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\b#*\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\b-.\n\n\n\x02\x04\x01\x12\x04\n\x00\r\x01\n\n\n\x03\x04\x01\x01\x12\x03\n\b\x0f\n\v\n\x04\x04\x01\x02\x00\x12\x03\v\x02*\n\f\n\x05\x04\x01\x02\x00\x05\x12\x03\v\x02\b\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\v\t\x10\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\v\x13\x14\n\f\n\x05\x04\x01\x02\x00\b\x12\x03\v\x15)\n\r\n\x06\x04\x01\x02\x00\b\x06\x12\x03\v\x16(\n\v\n\x04\x04\x01\x02\x01\x12\x03\f\x02\x1b\n\f\n\x05\x04\x01\x02\x01\x05\x12\x03\f\x02\b\n\f\n\x05\x04\x01\x02\x01\x01\x12\x03\f\t\x16\n\f\n\x05\x04\x01\x02\x01\x03\x12\x03\f\x19\x1a\n\xdf\x01\n\x02\x04\x02\x12\x04\x14\x00E\x01\x1a\xd2\x01 AuthStep\n A step in the authentication process\n Contains a variety of different types of views\n It is recommended to have a fallback_url specified\n For non-trivial authentication procedures (such as captchas)\n\n\n\n\x03\x04\x02\x01\x12\x03\x14\b\x10\n\x88\x01\n\x04\x04\x02\x03\x00\x12\x04\x18\x02\x1b\x03\x1az Choice\n A step which allows the user to choose from a range of options\n Allows you to show a heading by specifying title\n\n\f\n\x05\x04\x02\x03\x00\x01\x12\x03\x18\n\x10\n\r\n\x06\x04\x02\x03\x00\x02\x00\x12\x03\x19\x04\x15\n\x0e\n\a\x04\x02\x03\x00\x02\x00\x05\x12\x03\x19\x04\n\n\x0e\n\a\x04\x02\x03\x00\x02\x00\x01\x12\x03\x19\v\x10\n\x0e\n\a\x04\x02\x03\x00\x02\x00\x03\x12\x03\x19\x13\x14\n\r\n\x06\x04\x02\x03\x00\x02\x01\x12\x03\x1a\x04 \n\x0e\n\a\x04\x02\x03\x00\x02\x01\x04\x12\x03\x1a\x04\f\n\x0e\n\a\x04\x02\x03\x00\x02\x01\x05\x12\x03\x1a\r\x13\n\x0e\n\a\x04\x02\x03\x00\x02\x01\x01\x12\x03\x1a\x14\x1b\n\x0e\n\a\x04\x02\x03\x00\x02\x01\x03\x12\x03\x1a\x1e\x1f\n{\n\x04\x04\x02\x03\x01\x12\x04 \x021\x03\x1am Form\n A step which requires the user to input information\n Allows you to show a heading by specifying title\n\n\f\n\x05\x04\x02\x03\x01\x01\x12\x03 \n\x0e\n\x9c\x03\n\x06\x04\x02\x03\x01\x03\x00\x12\x04*\x04-\x05\x1a\x8b\x03 FormField\n A field in the form, containing information on how it should\n be rendered\n Here is a list of form types that need to be supported:\n email: a field type that has to contain a valid email\n password: a field type that has to contain a password\n new-password: a field type for new passwords\n text: a field type that has to contain text\n number: a field type that has to contain a number\n\n\x0e\n\a\x04\x02\x03\x01\x03\x00\x01\x12\x03*\f\x15\n\x0f\n\b\x04\x02\x03\x01\x03\x00\x02\x00\x12\x03+\x06\x16\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x00\x05\x12\x03+\x06\f\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x00\x01\x12\x03+\r\x11\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x00\x03\x12\x03+\x14\x15\n\x0f\n\b\x04\x02\x03\x01\x03\x00\x02\x01\x12\x03,\x06\x16\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x01\x05\x12\x03,\x06\f\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x01\x01\x12\x03,\r\x11\n\x10\n\t\x04\x02\x03\x01\x03\x00\x02\x01\x03\x12\x03,\x14\x15\n\r\n\x06\x04\x02\x03\x01\x02\x00\x12\x03/\x04\x15\n\x0e\n\a\x04\x02\x03\x01\x02\x00\x05\x12\x03/\x04\n\n\x0e\n\a\x04\x02\x03\x01\x02\x00\x01\x12\x03/\v\x10\n\x0e\n\a\x04\x02\x03\x01\x02\x00\x03\x12\x03/\x13\x14\n\r\n\x06\x04\x02\x03\x01\x02\x01\x12\x030\x04\"\n\x0e\n\a\x04\x02\x03\x01\x02\x01\x04\x12\x030\x04\f\n\x0e\n\a\x04\x02\x03\x01\x02\x01\x06\x12\x030\r\x16\n\x0e\n\a\x04\x02\x03\x01\x02\x01\x01\x12\x030\x17\x1d\n\x0e\n\a\x04\x02\x03\x01\x02\x01\x03\x12\x030 !\n\xb7\x01\n\x04\x04\x02\x03\x02\x12\x047\x02:\x03\x1a\xa8\x01 Waiting\n A step which requires the user to perform an external action\n The title and description should explain to the user\n what they should do to complete this step\n\n\f\n\x05\x04\x02\x03\x02\x01\x12\x037\n\x11\n\r\n\x06\x04\x02\x03\x02\x02\x00\x12\x038\x04\x15\n\x0e\n\a\x04\x02\x03\x02\x02\x00\x05\x12\x038\x04\n\n\x0e\n\a\x04\x02\x03\x02\x02\x00\x01\x12\x038\v\x10\n\x0e\n\a\x04\x02\x03\x02\x02\x00\x03\x12\x038\x13\x14\n\r\n\x06\x04\x02\x03\x02\x02\x01\x12\x039\x04\x1b\n\x0e\n\a\x04\x02\x03\x02\x02\x01\x05\x12\x039\x04\n\n\x0e\n\a\x04\x02\x03\x02\x02\x01\x01\x12\x039\v\x16\n\x0e\n\a\x04\x02\x03\x02\x02\x01\x03\x12\x039\x19\x1a\n\v\n\x04\x04\x02\x02\x00\x12\x03<\x02\x1a\n\f\n\x05\x04\x02\x02\x00\x05\x12\x03<\x02\b\n\f\n\x05\x04\x02\x02\x00\x01\x12\x03<\t\x15\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03<\x18\x19\n\v\n\x04\x04\x02\x02\x01\x12\x03=\x02\x17\n\f\n\x05\x04\x02\x02\x01\x05\x12\x03=\x02\x06\n\f\n\x05\x04\x02\x02\x01\x01\x12\x03=\a\x12\n\f\n\x05\x04\x02\x02\x01\x03\x12\x03=\x15\x16\n\f\n\x04\x04\x02\b\x00\x12\x04?\x02D\x03\n\f\n\x05\x04\x02\b\x00\x01\x12\x03?\b\f\n\v\n\x04\x04\x02\x02\x02\x12\x03@\x04\x16\n\f\n\x05\x04\x02\x02\x02\x06\x12\x03@\x04\n\n\f\n\x05\x04\x02\x02\x02\x01\x12\x03@\v\x11\n\f\n\x05\x04\x02\x02\x02\x03\x12\x03@\x14\x15\n\v\n\x04\x04\x02\x02\x03\x12\x03A\x04\x12\n\f\n\x05\x04\x02\x02\x03\x06\x12\x03A\x04\b\n\f\n\x05\x04\x02\x02\x03\x01\x12\x03A\t\r\n\f\n\x05\x04\x02\x02\x03\x03\x12\x03A\x10\x11\n\v\n\x04\x04\x02\x02\x04\x12\x03B\x04\x18\n\f\n\x05\x04\x02\x02\x04\x06\x12\x03B\x04\v\n\f\n\x05\x04\x02\x02\x04\x01\x12\x03B\f\x13\n\f\n\x05\x04\x02\x02\x04\x03\x12\x03B\x16\x17\n\v\n\x04\x04\x02\x02\x05\x12\x03C\x04\x18\n\f\n\x05\x04\x02\x02\x05\x06\x12\x03C\x04\v\n\f\n\x05\x04\x02\x02\x05\x01\x12\x03C\f\x13\n\f\n\x05\x04\x02\x02\x05\x03\x12\x03C\x16\x17\n\xa4\x01\n\x02\x04\x03\x12\x04K\x00a\x01\x1a\x97\x01 NextStepRequest\n contains the client's response to the server's challenge\n This needs to be called first with no arguments to\n receive the first step\n\n\n\n\x03\x04\x03\x01\x12\x03K\b\x17\n\v\n\x04\x04\x03\x02\x00\x12\x03L\x02\x15\n\f\n\x05\x04\x03\x02\x00\x05\x12\x03L\x02\b\n\f\n\x05\x04\x03\x02\x00\x01\x12\x03L\t\x10\n\f\n\x05\x04\x03\x02\x00\x03\x12\x03L\x13\x14\nL\n\x04\x04\x03\x03\x00\x12\x03O\x02'\x1a? A simple choice string indicating which option the user chose\n\n\f\n\x05\x04\x03\x03\x00\x01\x12\x03O\n\x10\n\r\n\x06\x04\x03\x03\x00\x02\x00\x12\x03O\x13%\n\x0e\n\a\x04\x03\x03\x00\x02\x00\x05\x12\x03O\x13\x19\n\x0e\n\a\x04\x03\x03\x00\x02\x00\x01\x12\x03O\x1a \n\x0e\n\a\x04\x03\x03\x00\x02\x00\x03\x12\x03O#$\nB\n\x04\x04\x03\x03\x01\x12\x04R\x02X\x03\x1a4 Form fields can either be bytes, string, or int64.\n\n\f\n\x05\x04\x03\x03\x01\x01\x12\x03R\n\x14\n\x0e\n\x06\x04\x03\x03\x01\b\x00\x12\x04S\x04W\x05\n\x0e\n\a\x04\x03\x03\x01\b\x00\x01\x12\x03S\n\x0f\n\r\n\x06\x04\x03\x03\x01\x02\x00\x12\x03T\x06\x16\n\x0e\n\a\x04\x03\x03\x01\x02\x00\x05\x12\x03T\x06\v\n\x0e\n\a\x04\x03\x03\x01\x02\x00\x01\x12\x03T\f\x11\n\x0e\n\a\x04\x03\x03\x01\x02\x00\x03\x12\x03T\x14\x15\n\r\n\x06\x04\x03\x03\x01\x02\x01\x12\x03U\x06\x18\n\x0e\n\a\x04\x03\x03\x01\x02\x01\x05\x12\x03U\x06\f\n\x0e\n\a\x04\x03\x03\x01\x02\x01\x01\x12\x03U\r\x13\n\x0e\n\a\x04\x03\x03\x01\x02\x01\x03\x12\x03U\x16\x17\n\r\n\x06\x04\x03\x03\x01\x02\x02\x12\x03V\x06\x17\n\x0e\n\a\x04\x03\x03\x01\x02\x02\x05\x12\x03V\x06\v\n\x0e\n\a\x04\x03\x03\x01\x02\x02\x01\x12\x03V\f\x12\n\x0e\n\a\x04\x03\x03\x01\x02\x02\x03\x12\x03V\x15\x16\nV\n\x04\x04\x03\x03\x02\x12\x03[\x022\x1aI An array of form fields, in the same order they came in from the server\n\n\f\n\x05\x04\x03\x03\x02\x01\x12\x03[\n\x0e\n\r\n\x06\x04\x03\x03\x02\x02\x00\x12\x03[\x110\n\x0e\n\a\x04\x03\x03\x02\x02\x00\x04\x12\x03[\x11\x19\n\x0e\n\a\x04\x03\x03\x02\x02\x00\x06\x12\x03[\x1a$\n\x0e\n\a\x04\x03\x03\x02\x02\x00\x01\x12\x03[%+\n\x0e\n\a\x04\x03\x03\x02\x02\x00\x03\x12\x03[./\n\f\n\x04\x04\x03\b\x00\x12\x04]\x02`\x03\n\f\n\x05\x04\x03\b\x00\x01\x12\x03]\b\f\n\v\n\x04\x04\x03\x02\x01\x12\x03^\x04\x16\n\f\n\x05\x04\x03\x02\x01\x06\x12\x03^\x04\n\n\f\n\x05\x04\x03\x02\x01\x01\x12\x03^\v\x11\n\f\n\x05\x04\x03\x02\x01\x03\x12\x03^\x14\x15\n\v\n\x04\x04\x03\x02\x02\x12\x03_\x04\x12\n\f\n\x05\x04\x03\x02\x02\x06\x12\x03_\x04\b\n\f\n\x05\x04\x03\x02\x02\x01\x12\x03_\t\r\n\f\n\x05\x04\x03\x02\x02\x03\x12\x03_\x10\x11\n9\n\x02\x04\x04\x12\x03e\x00/\x1a. StepBackRequest\n A request to go back 1 step\n\n\n\n\x03\x04\x04\x01\x12\x03e\b\x17\n\v\n\x04\x04\x04\x02\x00\x12\x03e\x1a-\n\f\n\x05\x04\x04\x02\x00\x05\x12\x03e\x1a \n\f\n\x05\x04\x04\x02\x00\x01\x12\x03e!(\n\f\n\x05\x04\x04\x02\x00\x03\x12\x03e+,\nx\n\x02\x04\x05\x12\x03j\x002\x1am StreamStepsRequest\n Required to be initiated by all authenticating clients\n Allows the server to send steps\n\n\n\n\x03\x04\x05\x01\x12\x03j\b\x1a\n\v\n\x04\x04\x05\x02\x00\x12\x03j\x1d0\n\f\n\x05\x04\x05\x02\x00\x05\x12\x03j\x1d#\n\f\n\x05\x04\x05\x02\x00\x01\x12\x03j$+\n\f\n\x05\x04\x05\x02\x00\x03\x12\x03j./\n\x17\n\x02\x04\x06\x12\x03m\x00.\x1a\f Connection\n\n\n\n\x03\x04\x06\x01\x12\x03m\b\x17\n\v\n\x04\x04\x06\x02\x00\x12\x03m\x1a,\n\f\n\x05\x04\x06\x02\x00\x05\x12\x03m\x1a \n\f\n\x05\x04\x06\x02\x00\x01\x12\x03m!'\n\f\n\x05\x04\x06\x02\x00\x03\x12\x03m*+\n\n\n\x02\x04\a\x12\x04n\x00q\x01\n\n\n\x03\x04\a\x01\x12\x03n\b\x15\n\v\n\x04\x04\a\x02\x00\x12\x03o\x02\x13\n\f\n\x05\x04\a\x02\x00\x05\x12\x03o\x02\b\n\f\n\x05\x04\a\x02\x00\x01\x12\x03o\t\x0e\n\f\n\x05\x04\a\x02\x00\x03\x12\x03o\x11\x12\n\v\n\x04\x04\a\x02\x01\x12\x03p\x02\x13\n\f\n\x05\x04\a\x02\x01\x05\x12\x03p\x02\b\n\f\n\x05\x04\a\x02\x01\x01\x12\x03p\t\x0e\n\f\n\x05\x04\a\x02\x01\x03\x12\x03p\x11\x12\n\t\n\x02\x04\b\x12\x03s\x00$\n\n\n\x03\x04\b\x01\x12\x03s\b\x10\n\v\n\x04\x04\b\x02\x00\x12\x03s\x13\"\n\f\n\x05\x04\b\x02\x00\x05\x12\x03s\x13\x19\n\f\n\x05\x04\b\x02\x00\x01\x12\x03s\x1a\x1d\n\f\n\x05\x04\b\x02\x00\x03\x12\x03s !\n\n\n\x02\x04\t\x12\x04u\x00x\x01\n\n\n\x03\x04\t\x01\x12\x03u\b\x1d\n\v\n\x04\x04\t\x02\x00\x12\x03v\x02\x18\n\f\n\x05\x04\t\x02\x00\x05\x12\x03v\x02\b\n\f\n\x05\x04\t\x02\x00\x01\x12\x03v\t\x13\n\f\n\x05\x04\t\x02\x00\x03\x12\x03v\x16\x17\n\v\n\x04\x04\t\x02\x01\x12\x03w\x02\x14\n\f\n\x05\x04\t\x02\x01\x05\x12\x03w\x02\b\n\f\n\x05\x04\t\x02\x01\x01\x12\x03w\t\x0f\n\f\n\x05\x04\t\x02\x01\x03\x12\x03w\x12\x13\n\v\n\x02\x06\x00\x12\x05z\x00\x82\x01\x01\n\n\n\x03\x06\x00\x01\x12\x03z\b\x13\n\v\n\x04\x06\x00\x02\x00\x12\x03{\x027\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03{\x06\x0e\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03{\x0f\x1e\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03{(5\n\v\n\x04\x06\x00\x02\x01\x12\x03|\x02=\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03|\x06\x14\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03|\x15*\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03|4;\n\v\n\x04\x06\x00\x02\x02\x12\x03}\x023\n\f\n\x05\x06\x00\x02\x02\x01\x12\x03}\x06\t\n\f\n\x05\x06\x00\x02\x02\x02\x12\x03}\n\x1f\n\f\n\x05\x06\x00\x02\x02\x03\x12\x03})1\n\v\n\x04\x06\x00\x02\x03\x12\x03~\x02B\n\f\n\x05\x06\x00\x02\x03\x01\x12\x03~\x06\x0f\n\f\n\x05\x06\x00\x02\x03\x02\x12\x03~\x10%\n\f\n\x05\x06\x00\x02\x03\x03\x12\x03~/@\n\v\n\x04\x06\x00\x02\x04\x12\x03\u007f\x022\n\f\n\x05\x06\x00\x02\x04\x01\x12\x03\u007f\x06\x0e\n\f\n\x05\x06\x00\x02\x04\x02\x12\x03\u007f\x0f\x1e\n\f\n\x05\x06\x00\x02\x04\x03\x12\x03\u007f(0\n\f\n\x04\x06\x00\x02\x05\x12\x04\x80\x01\x022\n\r\n\x05\x06\x00\x02\x05\x01\x12\x04\x80\x01\x06\x0e\n\r\n\x05\x06\x00\x02\x05\x02\x12\x04\x80\x01\x0f\x1e\n\r\n\x05\x06\x00\x02\x05\x03\x12\x04\x80\x01(0\n\f\n\x04\x06\x00\x02\x06\x12\x04\x81\x01\x02?\n\r\n\x05\x06\x00\x02\x06\x01\x12\x04\x81\x01\x06\x11\n\r\n\x05\x06\x00\x02\x06\x02\x12\x04\x81\x01\x12$\n\r\n\x05\x06\x00\x02\x06\x06\x12\x04\x81\x01.4\n\r\n\x05\x06\x00\x02\x06\x03\x12\x04\x81\x015=b\x06proto3")

	err := proto.Unmarshal(data, Authᐳv1ᐳauth)
	if err != nil {
		panic(err)
	}
}

type AuthServiceServer interface {
	Federate(ctx context.Context, r *FederateRequest, headers http.Header) (resp *FederateReply, err error)

	LoginFederated(ctx context.Context, r *LoginFederatedRequest, headers http.Header) (resp *Session, err error)

	Key(ctx context.Context, r *empty.Empty, headers http.Header) (resp *KeyReply, err error)

	BeginAuth(ctx context.Context, r *empty.Empty, headers http.Header) (resp *BeginAuthResponse, err error)

	NextStep(ctx context.Context, r *NextStepRequest, headers http.Header) (resp *AuthStep, err error)

	StepBack(ctx context.Context, r *StepBackRequest, headers http.Header) (resp *AuthStep, err error)

	StreamSteps(ctx context.Context, r *StreamStepsRequest, out chan *AuthStep, headers http.Header)
}

var AuthServiceServerFederateData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bFederate\x12!.protocol.auth.v1.FederateRequest\x1a\x1f.protocol.auth.v1.FederateReply")

	err := proto.Unmarshal(data, AuthServiceServerFederateData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerLoginFederatedData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0eLoginFederated\x12'.protocol.auth.v1.LoginFederatedRequest\x1a\x19.protocol.auth.v1.Session")

	err := proto.Unmarshal(data, AuthServiceServerLoginFederatedData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerKeyData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x03Key\x12\x16.google.protobuf.Empty\x1a\x1a.protocol.auth.v1.KeyReply")

	err := proto.Unmarshal(data, AuthServiceServerKeyData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerBeginAuthData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\tBeginAuth\x12\x16.google.protobuf.Empty\x1a#.protocol.auth.v1.BeginAuthResponse")

	err := proto.Unmarshal(data, AuthServiceServerBeginAuthData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerNextStepData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bNextStep\x12!.protocol.auth.v1.NextStepRequest\x1a\x1a.protocol.auth.v1.AuthStep")

	err := proto.Unmarshal(data, AuthServiceServerNextStepData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerStepBackData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\bStepBack\x12!.protocol.auth.v1.StepBackRequest\x1a\x1a.protocol.auth.v1.AuthStep")

	err := proto.Unmarshal(data, AuthServiceServerStepBackData)
	if err != nil {
		panic(err)
	}
}

var AuthServiceServerStreamStepsData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vStreamSteps\x12$.protocol.auth.v1.StreamStepsRequest\x1a\x1a.protocol.auth.v1.AuthStep0\x01")

	err := proto.Unmarshal(data, AuthServiceServerStreamStepsData)
	if err != nil {
		panic(err)
	}
}

type AuthServiceHandler struct {
	Server       AuthServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     func(meth *descriptorpb.MethodDescriptorProto, d *descriptorpb.FileDescriptorProto, f func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)) func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)
	upgrader     websocket.Upgrader
}

func NewAuthServiceHandler(s AuthServiceServer, errHandler func(err error, w http.ResponseWriter)) *AuthServiceHandler {
	return &AuthServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *AuthServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.auth.v1.AuthService/Federate":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(FederateRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.Federate(c, req.(*FederateRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerFederateData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/LoginFederated":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(LoginFederatedRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.LoginFederated(c, req.(*LoginFederatedRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerLoginFederatedData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/Key":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(empty.Empty)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.Key(c, req.(*empty.Empty), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerKeyData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/BeginAuth":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(empty.Empty)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.BeginAuth(c, req.(*empty.Empty), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerBeginAuthData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/NextStep":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(NextStepRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.NextStep(c, req.(*NextStepRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerNextStepData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/StepBack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(StepBackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.StepBack(c, req.(*StepBackRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(AuthServiceServerStepBackData, Authᐳv1ᐳauth, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.auth.v1.AuthService/StreamSteps":
		{
			var err error

			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			in := new(StreamStepsRequest)
			err = proto.Unmarshal(body, in)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *AuthStep)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(out)
							return
						}
					}
				}
			}()

			h.Server.StreamSteps(req.Context(), in, out, req.Header)
		}

	}
}
