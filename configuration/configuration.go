package configuration

import (
	"log"

	"github.com/tkanos/gonfig"
)

/*
* used for testing with authorised enpoints, have manually extracted token becuase i didnt want to deal with oauth shit, will leave this struct here to be used to hold this from a proper oauth process later.
*
 */
type TempTokens struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	MembershipId     string `json:"membership_id"`
}

type Configuration struct {
	BasePath string `json:"base-path"`
	API      API    `json:"api"`
}
type GetApplicationAPIUsage struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetBungieApplications struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type App struct {
	GetApplicationAPIUsage GetApplicationAPIUsage `json:"GetApplicationApiUsage"`
	GetBungieApplications  GetBungieApplications  `json:"GetBungieApplications"`
}
type GetBungieNetUserByID struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetCredentialTypesForTargetAccount struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetAvailableThemes struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetMembershipDataByID struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetMembershipDataForCurrentUser struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetMembershipFromHardLinkedCredential struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchByGlobalNamePrefix struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type User struct {
	GetBungieNetUserByID                  GetBungieNetUserByID                  `json:"GetBungieNetUserById"`
	GetCredentialTypesForTargetAccount    GetCredentialTypesForTargetAccount    `json:"GetCredentialTypesForTargetAccount"`
	GetAvailableThemes                    GetAvailableThemes                    `json:"GetAvailableThemes"`
	GetMembershipDataByID                 GetMembershipDataByID                 `json:"GetMembershipDataById"`
	GetMembershipDataForCurrentUser       GetMembershipDataForCurrentUser       `json:"GetMembershipDataForCurrentUser"`
	GetMembershipFromHardLinkedCredential GetMembershipFromHardLinkedCredential `json:"GetMembershipFromHardLinkedCredential"`
	SearchByGlobalNamePrefix              SearchByGlobalNamePrefix              `json:"SearchByGlobalNamePrefix"`
}
type GetContentType struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetContentByID struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetContentByTagAndType struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchContentWithText struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchContentByTagAndType struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchHelpArticles struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Content struct {
	GetContentType            GetContentType            `json:"GetContentType"`
	GetContentByID            GetContentByID            `json:"GetContentById"`
	GetContentByTagAndType    GetContentByTagAndType    `json:"GetContentByTagAndType"`
	SearchContentWithText     SearchContentWithText     `json:"SearchContentWithText"`
	SearchContentByTagAndType SearchContentByTagAndType `json:"SearchContentByTagAndType"`
	SearchHelpArticles        SearchHelpArticles        `json:"SearchHelpArticles"`
}
type GetTopicsPaged struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetCoreTopicsPaged struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPostsThreadedPaged struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPostsThreadedPagedFromChild struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPostAndParent struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPostAndParentAwaitingApproval struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetTopicForContent struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetForumTagSuggestions struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetPoll struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetRecruitmentThreadSummaries struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type Forum struct {
	GetTopicsPaged                   GetTopicsPaged                   `json:"GetTopicsPaged"`
	GetCoreTopicsPaged               GetCoreTopicsPaged               `json:"GetCoreTopicsPaged"`
	GetPostsThreadedPaged            GetPostsThreadedPaged            `json:"GetPostsThreadedPaged"`
	GetPostsThreadedPagedFromChild   GetPostsThreadedPagedFromChild   `json:"GetPostsThreadedPagedFromChild"`
	GetPostAndParent                 GetPostAndParent                 `json:"GetPostAndParent"`
	GetPostAndParentAwaitingApproval GetPostAndParentAwaitingApproval `json:"GetPostAndParentAwaitingApproval"`
	GetTopicForContent               GetTopicForContent               `json:"GetTopicForContent"`
	GetForumTagSuggestions           GetForumTagSuggestions           `json:"GetForumTagSuggestions"`
	GetPoll                          GetPoll                          `json:"GetPoll"`
	GetRecruitmentThreadSummaries    GetRecruitmentThreadSummaries    `json:"GetRecruitmentThreadSummaries"`
}
type GetAvailableAvatars struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetUserClanInviteSetting struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetRecommendedGroups struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GroupSearch struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetGroup struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetGroupByName struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetGroupByNameV2 struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetGroupOptionalConversations struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type EditGroup struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type EditClanBanner struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type EditFounderOptions struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type AddOptionalConversation struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type EditOptionalConversation struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetMembersOfGroup struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetAdminsAndFounderOfGroup struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type EditGroupMembership struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type KickMember struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type BanMember struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type UnbanMember struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetBannedMembersOfGroup struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type AbdicateFoundership struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPendingMemberships struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetInvitedIndividuals struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type ApproveAllPending struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type DenyAllPending struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type ApprovePendingForList struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type ApprovePending struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type DenyPendingForList struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetGroupsForMember struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type RecoverGroupForFounder struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPotentialGroupsForMember struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type IndividualGroupInvite struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type IndividualGroupInviteCancel struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetGroupV2 struct {
	GetAvailableAvatars           GetAvailableAvatars           `json:"GetAvailableAvatars"`
	GetAvailableThemes            GetAvailableThemes            `json:"GetAvailableThemes"`
	GetUserClanInviteSetting      GetUserClanInviteSetting      `json:"GetUserClanInviteSetting"`
	GetRecommendedGroups          GetRecommendedGroups          `json:"GetRecommendedGroups"`
	GroupSearch                   GroupSearch                   `json:"GroupSearch"`
	GetGroup                      GetGroup                      `json:"GetGroup"`
	GetGroupByName                GetGroupByName                `json:"GetGroupByName"`
	GetGroupByNameV2              GetGroupByNameV2              `json:"GetGroupByNameV2"`
	GetGroupOptionalConversations GetGroupOptionalConversations `json:"GetGroupOptionalConversations"`
	EditGroup                     EditGroup                     `json:"EditGroup"`
	EditClanBanner                EditClanBanner                `json:"EditClanBanner"`
	EditFounderOptions            EditFounderOptions            `json:"EditFounderOptions"`
	AddOptionalConversation       AddOptionalConversation       `json:"AddOptionalConversation"`
	EditOptionalConversation      EditOptionalConversation      `json:"EditOptionalConversation"`
	GetMembersOfGroup             GetMembersOfGroup             `json:"GetMembersOfGroup"`
	GetAdminsAndFounderOfGroup    GetAdminsAndFounderOfGroup    `json:"GetAdminsAndFounderOfGroup"`
	EditGroupMembership           EditGroupMembership           `json:"EditGroupMembership"`
	KickMember                    KickMember                    `json:"KickMember"`
	BanMember                     BanMember                     `json:"BanMember"`
	UnbanMember                   UnbanMember                   `json:"UnbanMember"`
	GetBannedMembersOfGroup       GetBannedMembersOfGroup       `json:"GetBannedMembersOfGroup"`
	AbdicateFoundership           AbdicateFoundership           `json:"AbdicateFoundership"`
	GetPendingMemberships         GetPendingMemberships         `json:"GetPendingMemberships"`
	GetInvitedIndividuals         GetInvitedIndividuals         `json:"GetInvitedIndividuals"`
	ApproveAllPending             ApproveAllPending             `json:"ApproveAllPending"`
	DenyAllPending                DenyAllPending                `json:"DenyAllPending"`
	ApprovePendingForList         ApprovePendingForList         `json:"ApprovePendingForList"`
	ApprovePending                ApprovePending                `json:"ApprovePending"`
	DenyPendingForList            DenyPendingForList            `json:"DenyPendingForList"`
	GetGroupsForMember            GetGroupsForMember            `json:"GetGroupsForMember"`
	RecoverGroupForFounder        RecoverGroupForFounder        `json:"RecoverGroupForFounder"`
	GetPotentialGroupsForMember   GetPotentialGroupsForMember   `json:"GetPotentialGroupsForMember"`
	IndividualGroupInvite         IndividualGroupInvite         `json:"IndividualGroupInvite"`
	IndividualGroupInviteCancel   IndividualGroupInviteCancel   `json:"IndividualGroupInviteCancel"`
}
type ClaimPartnerOffer struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type ApplyMissingPartnerOffersWithoutClaim struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPartnerOfferSkuHistory struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Tokens struct {
	ClaimPartnerOffer                     ClaimPartnerOffer                     `json:"ClaimPartnerOffer"`
	ApplyMissingPartnerOffersWithoutClaim ApplyMissingPartnerOffersWithoutClaim `json:"ApplyMissingPartnerOffersWithoutClaim"`
	GetPartnerOfferSkuHistory             GetPartnerOfferSkuHistory             `json:"GetPartnerOfferSkuHistory"`
}
type GetDestinyManifest struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetDestinyEntityDefinition struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchDestinyPlayer struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetLinkedProfiles struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetProfile struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetCharacter struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetClanWeeklyRewardState struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetClanBannerSource struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetItem struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetVendors struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetVendor struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPublicVendors struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetCollectibleNodeDetails struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type TransferItem struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type PullFromPostmaster struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type EquipItem struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type EquipItems struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type SetItemLockState struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type SetQuestTrackedState struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type InsertSocketPlug struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetPostGameCarnageReport struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type ReportOffensivePostGameCarnageReportPlayer struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetHistoricalStatsDefinition struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetClanLeaderboards struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetClanAggregateStats struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetLeaderboards struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetLeaderboardsForCharacter struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchDestinyEntities struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetHistoricalStats struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetHistoricalStatsForAccount struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetActivityHistory struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetUniqueWeaponHistory struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetDestinyAggregateActivityStats struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPublicMilestoneContent struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPublicMilestones struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type AwaInitializeRequest struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type AwaProvideAuthorizationResult struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type AwaGetActionToken struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Destiny2 struct {
	GetDestinyManifest                         GetDestinyManifest                         `json:"GetDestinyManifest"`
	GetDestinyEntityDefinition                 GetDestinyEntityDefinition                 `json:"GetDestinyEntityDefinition"`
	SearchDestinyPlayer                        SearchDestinyPlayer                        `json:"SearchDestinyPlayer"`
	GetLinkedProfiles                          GetLinkedProfiles                          `json:"GetLinkedProfiles"`
	GetProfile                                 GetProfile                                 `json:"GetProfile"`
	GetCharacter                               GetCharacter                               `json:"GetCharacter"`
	GetClanWeeklyRewardState                   GetClanWeeklyRewardState                   `json:"GetClanWeeklyRewardState"`
	GetClanBannerSource                        GetClanBannerSource                        `json:"GetClanBannerSource"`
	GetItem                                    GetItem                                    `json:"GetItem"`
	GetVendors                                 GetVendors                                 `json:"GetVendors"`
	GetVendor                                  GetVendor                                  `json:"GetVendor"`
	GetPublicVendors                           GetPublicVendors                           `json:"GetPublicVendors"`
	GetCollectibleNodeDetails                  GetCollectibleNodeDetails                  `json:"GetCollectibleNodeDetails"`
	TransferItem                               TransferItem                               `json:"TransferItem"`
	PullFromPostmaster                         PullFromPostmaster                         `json:"PullFromPostmaster"`
	EquipItem                                  EquipItem                                  `json:"EquipItem"`
	EquipItems                                 EquipItems                                 `json:"EquipItems"`
	SetItemLockState                           SetItemLockState                           `json:"SetItemLockState"`
	SetQuestTrackedState                       SetQuestTrackedState                       `json:"SetQuestTrackedState"`
	InsertSocketPlug                           InsertSocketPlug                           `json:"InsertSocketPlug"`
	GetPostGameCarnageReport                   GetPostGameCarnageReport                   `json:"GetPostGameCarnageReport"`
	ReportOffensivePostGameCarnageReportPlayer ReportOffensivePostGameCarnageReportPlayer `json:"ReportOffensivePostGameCarnageReportPlayer"`
	GetHistoricalStatsDefinition               GetHistoricalStatsDefinition               `json:"GetHistoricalStatsDefinition"`
	GetClanLeaderboards                        GetClanLeaderboards                        `json:"GetClanLeaderboards"`
	GetClanAggregateStats                      GetClanAggregateStats                      `json:"GetClanAggregateStats"`
	GetLeaderboards                            GetLeaderboards                            `json:"GetLeaderboards"`
	GetLeaderboardsForCharacter                GetLeaderboardsForCharacter                `json:"GetLeaderboardsForCharacter"`
	SearchDestinyEntities                      SearchDestinyEntities                      `json:"SearchDestinyEntities"`
	GetHistoricalStats                         GetHistoricalStats                         `json:"GetHistoricalStats"`
	GetHistoricalStatsForAccount               GetHistoricalStatsForAccount               `json:"GetHistoricalStatsForAccount"`
	GetActivityHistory                         GetActivityHistory                         `json:"GetActivityHistory"`
	GetUniqueWeaponHistory                     GetUniqueWeaponHistory                     `json:"GetUniqueWeaponHistory"`
	GetDestinyAggregateActivityStats           GetDestinyAggregateActivityStats           `json:"GetDestinyAggregateActivityStats"`
	GetPublicMilestoneContent                  GetPublicMilestoneContent                  `json:"GetPublicMilestoneContent"`
	GetPublicMilestones                        GetPublicMilestones                        `json:"GetPublicMilestones"`
	AwaInitializeRequest                       AwaInitializeRequest                       `json:"AwaInitializeRequest"`
	AwaProvideAuthorizationResult              AwaProvideAuthorizationResult              `json:"AwaProvideAuthorizationResult"`
	AwaGetActionToken                          AwaGetActionToken                          `json:"AwaGetActionToken"`
}
type GetCommunityContent struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type CommunityContent struct {
	GetCommunityContent GetCommunityContent `json:"GetCommunityContent"`
}
type GetTrendingCategories struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetTrendingCategory struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetTrendingEntryDetail struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Trending struct {
	GetTrendingCategories  GetTrendingCategories  `json:"GetTrendingCategories"`
	GetTrendingCategory    GetTrendingCategory    `json:"GetTrendingCategory"`
	GetTrendingEntryDetail GetTrendingEntryDetail `json:"GetTrendingEntryDetail"`
}
type GetActivePrivateClanFireteamCount struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetAvailableClanFireteams struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type SearchPublicAvailableClanFireteams struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetMyClanFireteams struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetClanFireteam struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Fireteam struct {
	GetActivePrivateClanFireteamCount  GetActivePrivateClanFireteamCount  `json:"GetActivePrivateClanFireteamCount"`
	GetAvailableClanFireteams          GetAvailableClanFireteams          `json:"GetAvailableClanFireteams"`
	SearchPublicAvailableClanFireteams SearchPublicAvailableClanFireteams `json:"SearchPublicAvailableClanFireteams"`
	GetMyClanFireteams                 GetMyClanFireteams                 `json:"GetMyClanFireteams"`
	GetClanFireteam                    GetClanFireteam                    `json:"GetClanFireteam"`
}
type GetFriendList struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetFriendRequestList struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type IssueFriendRequest struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type AcceptFriendRequest struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type DeclineFriendRequest struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type RemoveFriend struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type RemoveFriendRequest struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type GetPlatformFriendList struct {
	Method     string   `json:"method"`
	Endpoint   string   `json:"endpoint"`
	Parameters []string `json:"parameters"`
}
type Social struct {
	GetFriendList         GetFriendList         `json:"GetFriendList"`
	GetFriendRequestList  GetFriendRequestList  `json:"GetFriendRequestList"`
	IssueFriendRequest    IssueFriendRequest    `json:"IssueFriendRequest"`
	AcceptFriendRequest   AcceptFriendRequest   `json:"AcceptFriendRequest"`
	DeclineFriendRequest  DeclineFriendRequest  `json:"DeclineFriendRequest"`
	RemoveFriend          RemoveFriend          `json:"RemoveFriend"`
	RemoveFriendRequest   RemoveFriendRequest   `json:"RemoveFriendRequest"`
	GetPlatformFriendList GetPlatformFriendList `json:"GetPlatformFriendList"`
}
type GetAvailableLocales struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetCommonSettings struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetUserSystemOverrides struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type GetGlobalAlerts struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}
type Null struct {
	GetAvailableLocales    GetAvailableLocales    `json:"GetAvailableLocales"`
	GetCommonSettings      GetCommonSettings      `json:"GetCommonSettings"`
	GetUserSystemOverrides GetUserSystemOverrides `json:"GetUserSystemOverrides"`
	GetGlobalAlerts        GetGlobalAlerts        `json:"GetGlobalAlerts"`
}
type API struct {
	App              App              `json:"App"`
	User             User             `json:"User"`
	Content          Content          `json:"Content"`
	Forum            Forum            `json:"Forum"`
	GroupV2          GetGroupV2       `json:"GroupV2"`
	Tokens           Tokens           `json:"Tokens"`
	Destiny2         Destiny2         `json:"Destiny2"`
	CommunityContent CommunityContent `json:"CommunityContent"`
	Trending         Trending         `json:"Trending"`
	Fireteam         Fireteam         `json:"Fireteam"`
	Social           Social           `json:"Social"`
	Null             Null             `json:"null"`
}

var BasePath = "https://www.bungie.net/Platform"

func ReadConfig() *Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("configuration/bungieAPI_config.json", &configuration)
	if err != nil {
		log.Fatalln(err)
	}
	return &configuration
}
