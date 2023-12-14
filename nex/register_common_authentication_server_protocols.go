package nex

import (
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-go/ticket-granting"
	common_ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/ticket-granting"
	"github.com/PretendoNetwork/mario-kart-7/globals"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticket_granting.NewProtocol(globals.AuthenticationServer)
	commonTicketGrantingProtocol := common_ticket_granting.NewCommonProtocol(ticketGrantingProtocol)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.Scheme = "prudps"
	secureStationURL.Fields.Set("address", os.Getenv("PN_MK7_SECURE_SERVER_HOST"))
	secureStationURL.Fields.Set("port", os.Getenv("PN_MK7_SECURE_SERVER_PORT"))
	secureStationURL.Fields.Set("CID", "1")
	secureStationURL.Fields.Set("PID", "2")
	secureStationURL.Fields.Set("sid", "1")
	secureStationURL.Fields.Set("stream", "10")
	secureStationURL.Fields.Set("type", "2")

	commonTicketGrantingProtocol.SecureStationURL = secureStationURL
	commonTicketGrantingProtocol.BuildName = serverBuildString

	globals.AuthenticationServer.SetPasswordFromPIDFunction(globals.PasswordFromPID)
}
