#include "voice/v1/voice.hrpc.client.h"
auto VoiceServiceServiceClient::Connect() -> Receive__protocol_voice_v1_Signal__Send__protocol_voice_v1_ClientSignal__Stream*
{
	auto sock = new Receive__protocol_voice_v1_Signal__Send__protocol_voice_v1_ClientSignal__Stream();
	sock->open(QUrl(wsProtocol()+host));
	return sock;
}
