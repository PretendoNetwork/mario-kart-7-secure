package nex

import (
	"github.com/PretendoNetwork/mario-kart-7/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/matchmake-extension"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/match-making-ext"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/match-making-ext"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/nat-traversal"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"

	nex_matchmake_extension_common "github.com/PretendoNetwork/mario-kart-7/nex/matchmake-extension/common"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol(globals.SecureServer)
	common_secure.NewCommonProtocol(secureProtocol)

	natTraversalProtocol := nat_traversal.NewProtocol(globals.SecureServer)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol(globals.SecureServer)
	common_match_making.NewCommonProtocol(matchMakingProtocol)

	matchMakingExtProtocol := match_making_ext.NewProtocol(globals.SecureServer)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol(globals.SecureServer)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = nex_matchmake_extension_common.CleanupSearchMatchmakeSession
}
