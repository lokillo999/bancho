package pid

import (
	"strconv"
)

var tostr = [...]string{
	"OsuSendUserState",              
	"OsuSendIRCMessage",
	"OsuExit",
	"OsuRequestStatusUpdate",
	"OsuPong",
	"BanchoLoginReply",
	"BanchoCommandError",
	"BanchoSendMessage",
	"BanchoPing",
	"BanchoHandleIRCUsernameChange",
	"BanchoHandleIRCQuit",
	"BanchoHandleUserUpdate",
	"BanchoHandleUserQuit",
	"BanchoSpectatorJoined",
	"BanchoSpectatorLeft",
	"BanchoSpectateFrames",
	"OsuStartSpectating",
	"OsuStopSpectating",
	"OsuSpectateFrames",
	"BanchoVersionUpdate",
	"OsuErrorReport",
	"OsuCantSpectate",
	"BanchoSpectatorCantSpectate",
	"BanchoGetAttention",
	"BanchoAnnounce",
	"OsuSendIRCMessagePrivate",
	"BanchoMatchUpdate",
	"BanchoMatchNew",
	"BanchoMatchDisband",
	"OsuLobbyPart",
	"OsuLobbyJoin",
	"OsuMatchCreate",
	"OsuMatchJoin",
	"OsuLobbySomething",
	"BanchoLobbyJoinOBSOLETE",
	"BanchoLobbyPartOBSOLETE",
	"BanchoMatchJoinSuccess",
	"BanchoMatchJoinFail",
	"OsuMatchChangeSlot",
	"OsuMatchReady",
	"OsuMatchLock",
	"OsuMatchChangeSettings",
	"BanchoFellowSpectatorJoined",
	"BanchoFellowSpectatorLeft",
	"OsuMatchStart",
	"AllPlayersLoaded",
	"BanchoMatchStart",
	"OsuMatchScoreUpdate",
	"BanchoMatchScoreUpdate",
	"OsuMatchComplete",
	"BanchoMatchTransferHost",
	"OsuMatchChangeMods",
	"OsuMatchLoadComplete",
	"BanchoMatchAllPlayersLoaded",
	"OsuMatchNoBeatmap",
	"OsuMatchNotReady",
	"OsuMatchFailed",
	"BanchoMatchPlayerFailed",
	"BanchoMatchComplete",
	"OsuMatchHasBeatmap",
	"OsuMatchSkipRequest",
	"BanchoMatchSkip",
	"BanchoUnauthorised",
	"OsuChannelJoin",
	"BanchoChannelJoinSuccess",
	"BanchoChannelAvailable",
	"BanchoChannelRevoked",
	"BanchoChannelAvailableAutojoin",
	"OsuBeatmapInfoRequest",
	"BanchoBeatmapInfoReply",
	"OsuMatchTransferHost",
	"BanchoLoginPermissions",
	"BanchoFriendList",
	"OsuFriendAdd",
	"OsuFriendRemove",
	"BanchoProtocolVersion",
	"BanchoTitleUpdate",
	"OsuMatchChangeTeam",
	"OsuChannelLeave",
	"OsuReceiveUpdates",
	"BanchoMonitor",
	"BanchoMatchPlayerSkipped",
	"OsuSetIrcAwayMessage",
	"BanchoUserPresence",
	"IRCOnly",
	"OsuUserStatsRequest",
	"BanchoRestart",
	"OsuInvite",
	"BanchoInvite",
	"BanchoChannelListingComplete",
	"OsuMatchChangePassword",
	"BanchoMatchChangePassword",
	"BanchoBanInfo",
	"OsuSpecialMatchInfoRequest",
	"BanchoUserSilenced",
	"BanchoUserPresenceSingle",
	"BanchoUserPresenceBundle",
	"OsuUserPresenceRequest",
	"OsuUserPresenceRequestAll",
	"OsuUserToggleBlockNonFriendPM",
	"BanchoUserPMBlocked",
	"BanchoTargetIsSilenced",
	"BanchoVersionUpdateForced",
	"BanchoSwitchServer",
	"BanchoAccountRestricted",
	"BanchoRTX",
	"OsuMatchAbort",
	"BanchoSwitchTourneyServer",
	"OsuSpecialJoinMatchChannel",
	"OsuSpecialLeaveMatchChannel",
}

func String(packetID int) string {
	if packetID > (len(tostr) - 1) {
		return strconv.Itoa(packetID)
	} 
	return tostr[packetID]
}
