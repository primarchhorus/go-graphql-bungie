package resolvers

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/primarchhorus/go-graphql-bungie/configuration"
	"github.com/primarchhorus/go-graphql-bungie/data"
)

var config = configuration.ReadConfig()

func GetApplicationApiUsage(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.App.GetApplicationAPIUsage.Endpoint, p.Args["applicationId"].(string))
	var return_val data.ApiUsage
	ret, err := base_resolver(p, url, config.API.App.GetApplicationAPIUsage.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetBungieApplications(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.App.GetBungieApplications.Endpoint)
	var return_val data.Application
	ret, err := base_resolver(p, url, config.API.App.GetBungieApplications.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetBungieNetUserById(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetBungieNetUserByID.Endpoint, p.Args["id"].(string))
	var return_val data.GeneralUser
	ret, err := base_resolver(p, url, config.API.User.GetBungieNetUserByID.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCredentialTypesForTargetAccount(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetCredentialTypesForTargetAccount.Endpoint, p.Args["membershipId"].(string))
	var return_val data.GetCredentialTypesForAccountResponse
	ret, err := base_resolver(p, url, config.API.User.GetCredentialTypesForTargetAccount.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetAvailableThemes(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetAvailableThemes.Endpoint)
	var return_val data.UserTheme
	ret, err := base_resolver(p, url, config.API.User.GetAvailableThemes.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetMembershipDataById(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetMembershipDataByID.Endpoint, p.Args["membershipId"].(string), p.Args["membershipType"].(string))
	var return_val data.UserMembershipData
	ret, err := base_resolver(p, url, config.API.User.GetMembershipDataByID.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetMembershipDataForCurrentUser(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetMembershipDataForCurrentUser.Endpoint)
	var return_val data.UserMembershipData
	ret, err := base_resolver(p, url, config.API.User.GetMembershipDataForCurrentUser.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetMembershipFromHardLinkedCredential(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.GetMembershipFromHardLinkedCredential.Endpoint, p.Args["crType"].(string), p.Args["credential"].(string))
	var return_val data.HardLinkedUserMembership
	ret, err := base_resolver(p, url, config.API.User.GetMembershipFromHardLinkedCredential.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchByGlobalNamePrefix(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.User.SearchByGlobalNamePrefix.Endpoint, p.Args["displayNamePrefix"].(string), p.Args["page"].(string))
	var return_val data.UserSearchResponse
	ret, err := base_resolver(p, url, config.API.User.SearchByGlobalNamePrefix.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetContentType(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.GetContentType.Endpoint, p.Args["type"].(string))
	var return_val data.ContentTypeDescription
	ret, err := base_resolver(p, url, config.API.Content.GetContentType.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetContentById(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.GetContentByID.Endpoint, p.Args["id"].(string), p.Args["locale"].(string))
	var return_val data.ContentItemPublicContract
	ret, err := base_resolver(p, url, config.API.Content.GetContentByID.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetContentByTagAndType(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.GetContentByTagAndType.Endpoint, p.Args["tag"].(string), p.Args["type"].(string), p.Args["locale"].(string))
	var return_val data.ContentItemPublicContract
	ret, err := base_resolver(p, url, config.API.Content.GetContentByTagAndType.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchContentWithText(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.SearchContentWithText.Endpoint, p.Args["locale"].(string))
	var return_val data.SearchResultOfContentItemPublicContract
	ret, err := base_resolver(p, url, config.API.Content.SearchContentWithText.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchContentByTagAndType(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.SearchContentByTagAndType.Endpoint, p.Args["tag"].(string), p.Args["type"].(string), p.Args["locale"].(string))
	var return_val data.SearchResultOfContentItemPublicContract
	ret, err := base_resolver(p, url, config.API.Content.SearchContentByTagAndType.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchHelpArticles(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Content.SearchHelpArticles.Endpoint, p.Args["searchtext"].(string), p.Args["size"].(string))
	var return_val interface{}
	ret, err := base_resolver(p, url, config.API.Content.SearchHelpArticles.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetTopicsPaged(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetTopicsPaged.Endpoint, p.Args["page"].(string), p.Args["pageSize"].(string), p.Args["group"].(string), p.Args["sort"].(string), p.Args["quickDate"].(string), p.Args["categoryFilter"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetTopicsPaged.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCoreTopicsPaged(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetCoreTopicsPaged.Endpoint, p.Args["page"].(string), p.Args["sort"].(string), p.Args["quickDate"].(string), p.Args["categoryFilter"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetCoreTopicsPaged.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPostsThreadedPaged(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetPostsThreadedPaged.Endpoint, p.Args["parentPostId"].(string), p.Args["page"].(string), p.Args["pageSize"].(string), p.Args["replySize"].(string), p.Args["getParentPost"].(string), p.Args["rootThreadMode"].(string), p.Args["sortMode"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetPostsThreadedPaged.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPostsThreadedPagedFromChild(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetPostsThreadedPagedFromChild.Endpoint, p.Args["childPostId"].(string), p.Args["page"].(string), p.Args["pageSize"].(string), p.Args["replySize"].(string), p.Args["rootThreadMode"].(string), p.Args["sortMode"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetPostsThreadedPagedFromChild.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPostAndParent(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetPostAndParent.Endpoint, p.Args["childPostId"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetPostAndParent.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPostAndParentAwaitingApproval(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetPostAndParentAwaitingApproval.Endpoint, p.Args["childPostId"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetPostAndParentAwaitingApproval.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetTopicForContent(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetTopicForContent.Endpoint, p.Args["contentId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.Forum.GetTopicForContent.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetForumTagSuggestions(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetForumTagSuggestions.Endpoint)
	var return_val data.TagResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetForumTagSuggestions.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPoll(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetPoll.Endpoint, p.Args["topicId"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.Forum.GetPoll.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetRecruitmentThreadSummaries(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Forum.GetRecruitmentThreadSummaries.Endpoint)
	var return_val data.ForumRecruitmentDetail
	ret, err := base_resolver(p, url, config.API.Forum.GetRecruitmentThreadSummaries.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetAvailableAvatars(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetAvailableAvatars.Endpoint)
	var return_val string
	ret, err := base_resolver(p, url, config.API.GroupV2.GetAvailableAvatars.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}

// func GetAvailableThemes(p graphql.ResolveParams) (interface{}, error) {
//     url := fmt.Sprintf(config.API.GroupV2.GetAvailableThemes.Endpoint)
//     var return_val data.GroupTheme
//     ret, err := base_resolver(p, url, config.API.GroupV2.GetAvailableThemes.Method, &return_val)
//     if err != nil {
//         log.Println(err)
//     }
//     return ret, nil
// }
func GetUserClanInviteSetting(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetUserClanInviteSetting.Endpoint, p.Args["mType"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.GroupV2.GetUserClanInviteSetting.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetRecommendedGroups(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetRecommendedGroups.Endpoint, p.Args["groupType"].(string), p.Args["createDateRange"].(string))
	var return_val data.GroupV2Card
	ret, err := base_resolver(p, url, config.API.GroupV2.GetRecommendedGroups.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GroupSearch(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GroupSearch.Endpoint)
	var return_val data.GroupSearchResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GroupSearch.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGroup(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetGroup.Endpoint, p.Args["groupId"].(string))
	var return_val data.GroupResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GetGroup.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGroupByName(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetGroupByName.Endpoint, p.Args["groupName"].(string), p.Args["groupType"].(string))
	var return_val data.GroupResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GetGroupByName.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGroupByNameV2(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetGroupByNameV2.Endpoint)
	var return_val data.GroupResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GetGroupByNameV2.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGroupOptionalConversations(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetGroupOptionalConversations.Endpoint, p.Args["groupId"].(string))
	var return_val data.GroupOptionalConversation
	ret, err := base_resolver(p, url, config.API.GroupV2.GetGroupOptionalConversations.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EditGroup(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.EditGroup.Endpoint, p.Args["groupId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.EditGroup.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EditClanBanner(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.EditClanBanner.Endpoint, p.Args["groupId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.EditClanBanner.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EditFounderOptions(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.EditFounderOptions.Endpoint, p.Args["groupId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.EditFounderOptions.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AddOptionalConversation(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.AddOptionalConversation.Endpoint, p.Args["groupId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.AddOptionalConversation.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EditOptionalConversation(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.EditOptionalConversation.Endpoint, p.Args["groupId"].(string), p.Args["conversationId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.EditOptionalConversation.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetMembersOfGroup(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetMembersOfGroup.Endpoint, p.Args["groupId"].(string))
	var return_val data.SearchResultOfGroupMember
	ret, err := base_resolver(p, url, config.API.GroupV2.GetMembersOfGroup.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetAdminsAndFounderOfGroup(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetAdminsAndFounderOfGroup.Endpoint, p.Args["groupId"].(string))
	var return_val data.SearchResultOfGroupMember
	ret, err := base_resolver(p, url, config.API.GroupV2.GetAdminsAndFounderOfGroup.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EditGroupMembership(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.EditGroupMembership.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string), p.Args["memberType"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.EditGroupMembership.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func KickMember(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.KickMember.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val data.GroupMemberLeaveResult
	ret, err := base_resolver(p, url, config.API.GroupV2.KickMember.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func BanMember(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.BanMember.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.BanMember.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func UnbanMember(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.UnbanMember.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.GroupV2.UnbanMember.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetBannedMembersOfGroup(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetBannedMembersOfGroup.Endpoint, p.Args["groupId"].(string))
	var return_val data.SearchResultOfGroupBan
	ret, err := base_resolver(p, url, config.API.GroupV2.GetBannedMembersOfGroup.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AbdicateFoundership(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.AbdicateFoundership.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["founderIdNew"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.GroupV2.AbdicateFoundership.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPendingMemberships(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetPendingMemberships.Endpoint, p.Args["groupId"].(string))
	var return_val data.SearchResultOfGroupMemberApplication
	ret, err := base_resolver(p, url, config.API.GroupV2.GetPendingMemberships.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetInvitedIndividuals(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetInvitedIndividuals.Endpoint, p.Args["groupId"].(string))
	var return_val data.SearchResultOfGroupMemberApplication
	ret, err := base_resolver(p, url, config.API.GroupV2.GetInvitedIndividuals.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ApproveAllPending(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.ApproveAllPending.Endpoint, p.Args["groupId"].(string))
	var return_val data.EntityActionResult
	ret, err := base_resolver(p, url, config.API.GroupV2.ApproveAllPending.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func DenyAllPending(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.DenyAllPending.Endpoint, p.Args["groupId"].(string))
	var return_val data.EntityActionResult
	ret, err := base_resolver(p, url, config.API.GroupV2.DenyAllPending.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ApprovePendingForList(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.ApprovePendingForList.Endpoint, p.Args["groupId"].(string))
	var return_val data.EntityActionResult
	ret, err := base_resolver(p, url, config.API.GroupV2.ApprovePendingForList.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ApprovePending(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.ApprovePending.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.GroupV2.ApprovePending.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func DenyPendingForList(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.DenyPendingForList.Endpoint, p.Args["groupId"].(string))
	var return_val data.EntityActionResult
	ret, err := base_resolver(p, url, config.API.GroupV2.DenyPendingForList.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGroupsForMember(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetGroupsForMember.Endpoint, p.Args["membershipType"].(string), p.Args["membershipId"].(string), p.Args["filter"].(string), p.Args["groupType"].(string))
	var return_val data.GetGroupsForMemberResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GetGroupsForMember.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func RecoverGroupForFounder(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.RecoverGroupForFounder.Endpoint, p.Args["membershipType"].(string), p.Args["membershipId"].(string), p.Args["groupType"].(string))
	var return_val data.GroupMembershipSearchResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.RecoverGroupForFounder.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPotentialGroupsForMember(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.GetPotentialGroupsForMember.Endpoint, p.Args["membershipType"].(string), p.Args["membershipId"].(string), p.Args["filter"].(string), p.Args["groupType"].(string))
	var return_val data.GroupPotentialMembershipSearchResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.GetPotentialGroupsForMember.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func IndividualGroupInvite(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.IndividualGroupInvite.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val data.GroupApplicationResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.IndividualGroupInvite.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func IndividualGroupInviteCancel(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.GroupV2.IndividualGroupInviteCancel.Endpoint, p.Args["groupId"].(string), p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val data.GroupApplicationResponse
	ret, err := base_resolver(p, url, config.API.GroupV2.IndividualGroupInviteCancel.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ClaimPartnerOffer(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Tokens.ClaimPartnerOffer.Endpoint)
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Tokens.ClaimPartnerOffer.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ApplyMissingPartnerOffersWithoutClaim(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Tokens.ApplyMissingPartnerOffersWithoutClaim.Endpoint, p.Args["partnerApplicationId"].(string), p.Args["targetBnetMembershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Tokens.ApplyMissingPartnerOffersWithoutClaim.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPartnerOfferSkuHistory(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Tokens.GetPartnerOfferSkuHistory.Endpoint, p.Args["partnerApplicationId"].(string), p.Args["targetBnetMembershipId"].(string))
	var return_val data.PartnerOfferSkuHistoryResponse
	ret, err := base_resolver(p, url, config.API.Tokens.GetPartnerOfferSkuHistory.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetDestinyManifest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetDestinyManifest.Endpoint)
	var return_val data.DestinyManifest
	ret, err := base_resolver(p, url, config.API.Destiny2.GetDestinyManifest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetDestinyEntityDefinition(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetDestinyEntityDefinition.Endpoint, p.Args["entityType"].(string), p.Args["hashIdentifier"].(string))
	var return_val data.DestinyDefinition
	ret, err := base_resolver(p, url, config.API.Destiny2.GetDestinyEntityDefinition.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchDestinyPlayer(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.SearchDestinyPlayer.Endpoint, p.Args["membershipType"].(string), p.Args["displayName"].(string))
	var return_val data.UserInfoCard
	ret, err := base_resolver(p, url, config.API.Destiny2.SearchDestinyPlayer.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetLinkedProfiles(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetLinkedProfiles.Endpoint, p.Args["membershipType"].(string), p.Args["membershipId"].(string))
	var return_val data.DestinyLinkedProfilesResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetLinkedProfiles.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetProfile(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetProfile.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string))
	var return_val data.DestinyProfileResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetProfile.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCharacter(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetCharacter.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyCharacterResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetCharacter.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetClanWeeklyRewardState(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetClanWeeklyRewardState.Endpoint, p.Args["groupId"].(string))
	var return_val data.DestinyMilestone
	ret, err := base_resolver(p, url, config.API.Destiny2.GetClanWeeklyRewardState.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetClanBannerSource(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetClanBannerSource.Endpoint)
	var return_val data.ClanBannerSource
	ret, err := base_resolver(p, url, config.API.Destiny2.GetClanBannerSource.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetItem(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetItem.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["itemInstanceId"].(string))
	var return_val data.DestinyItemResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetItem.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetVendors(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetVendors.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyVendorsResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetVendors.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetVendor(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetVendor.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string), p.Args["vendorHash"].(string))
	var return_val data.DestinyVendorResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetVendor.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPublicVendors(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetPublicVendors.Endpoint)
	var return_val data.DestinyPublicVendorsResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetPublicVendors.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCollectibleNodeDetails(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetCollectibleNodeDetails.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string), p.Args["collectiblePresentationNodeHash"].(string))
	var return_val data.DestinyCollectibleNodeDetailResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.GetCollectibleNodeDetails.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func TransferItem(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.TransferItem.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.TransferItem.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func PullFromPostmaster(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.PullFromPostmaster.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.PullFromPostmaster.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EquipItem(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.EquipItem.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.EquipItem.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func EquipItems(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.EquipItems.Endpoint)
	var return_val data.DestinyEquipItemResults
	ret, err := base_resolver(p, url, config.API.Destiny2.EquipItems.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SetItemLockState(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.SetItemLockState.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.SetItemLockState.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SetQuestTrackedState(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.SetQuestTrackedState.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.SetQuestTrackedState.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func InsertSocketPlug(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.InsertSocketPlug.Endpoint)
	var return_val data.DestinyItemChangeResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.InsertSocketPlug.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPostGameCarnageReport(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetPostGameCarnageReport.Endpoint, p.Args["activityId"].(string))
	var return_val data.DestinyPostGameCarnageReportData
	ret, err := base_resolver(p, url, config.API.Destiny2.GetPostGameCarnageReport.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func ReportOffensivePostGameCarnageReportPlayer(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.ReportOffensivePostGameCarnageReportPlayer.Endpoint, p.Args["activityId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.ReportOffensivePostGameCarnageReportPlayer.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetHistoricalStatsDefinition(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetHistoricalStatsDefinition.Endpoint)
	var return_val data.DestinyHistoricalStatsDefinition
	ret, err := base_resolver(p, url, config.API.Destiny2.GetHistoricalStatsDefinition.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetClanLeaderboards(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetClanLeaderboards.Endpoint, p.Args["groupId"].(string))
	var return_val data.DestinyLeaderboard
	ret, err := base_resolver(p, url, config.API.Destiny2.GetClanLeaderboards.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetClanAggregateStats(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetClanAggregateStats.Endpoint, p.Args["groupId"].(string))
	var return_val data.DestinyClanAggregateStat
	ret, err := base_resolver(p, url, config.API.Destiny2.GetClanAggregateStats.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetLeaderboards(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetLeaderboards.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string))
	var return_val data.DestinyLeaderboard
	ret, err := base_resolver(p, url, config.API.Destiny2.GetLeaderboards.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetLeaderboardsForCharacter(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetLeaderboardsForCharacter.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyLeaderboard
	ret, err := base_resolver(p, url, config.API.Destiny2.GetLeaderboardsForCharacter.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchDestinyEntities(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.SearchDestinyEntities.Endpoint, p.Args["type"].(string), p.Args["searchTerm"].(string))
	var return_val data.DestinyEntitySearchResult
	ret, err := base_resolver(p, url, config.API.Destiny2.SearchDestinyEntities.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetHistoricalStats(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetHistoricalStats.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyHistoricalStatsByPeriod
	ret, err := base_resolver(p, url, config.API.Destiny2.GetHistoricalStats.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetHistoricalStatsForAccount(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetHistoricalStatsForAccount.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string))
	var return_val data.DestinyHistoricalStatsAccountResult
	ret, err := base_resolver(p, url, config.API.Destiny2.GetHistoricalStatsForAccount.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetActivityHistory(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetActivityHistory.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyActivityHistoryResults
	ret, err := base_resolver(p, url, config.API.Destiny2.GetActivityHistory.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetUniqueWeaponHistory(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetUniqueWeaponHistory.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyHistoricalWeaponStatsData
	ret, err := base_resolver(p, url, config.API.Destiny2.GetUniqueWeaponHistory.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetDestinyAggregateActivityStats(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetDestinyAggregateActivityStats.Endpoint, p.Args["membershipType"].(string), p.Args["destinyMembershipId"].(string), p.Args["characterId"].(string))
	var return_val data.DestinyAggregateActivityResults
	ret, err := base_resolver(p, url, config.API.Destiny2.GetDestinyAggregateActivityStats.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPublicMilestoneContent(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetPublicMilestoneContent.Endpoint, p.Args["milestoneHash"].(string))
	var return_val data.DestinyMilestoneContent
	ret, err := base_resolver(p, url, config.API.Destiny2.GetPublicMilestoneContent.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPublicMilestones(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.GetPublicMilestones.Endpoint)
	var return_val data.DestinyPublicMilestone
	ret, err := base_resolver(p, url, config.API.Destiny2.GetPublicMilestones.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AwaInitializeRequest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.AwaInitializeRequest.Endpoint)
	var return_val data.AwaInitializeResponse
	ret, err := base_resolver(p, url, config.API.Destiny2.AwaInitializeRequest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AwaProvideAuthorizationResult(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.AwaProvideAuthorizationResult.Endpoint)
	var return_val int
	ret, err := base_resolver(p, url, config.API.Destiny2.AwaProvideAuthorizationResult.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AwaGetActionToken(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Destiny2.AwaGetActionToken.Endpoint, p.Args["correlationId"].(string))
	var return_val data.AwaAuthorizationResult
	ret, err := base_resolver(p, url, config.API.Destiny2.AwaGetActionToken.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCommunityContent(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.CommunityContent.GetCommunityContent.Endpoint, p.Args["sort"].(string), p.Args["mediaFilter"].(string), p.Args["page"].(string))
	var return_val data.PostSearchResponse
	ret, err := base_resolver(p, url, config.API.CommunityContent.GetCommunityContent.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetTrendingCategories(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Trending.GetTrendingCategories.Endpoint)
	var return_val data.TrendingCategories
	ret, err := base_resolver(p, url, config.API.Trending.GetTrendingCategories.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetTrendingCategory(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Trending.GetTrendingCategory.Endpoint, p.Args["categoryId"].(string), p.Args["pageNumber"].(string))
	var return_val data.SearchResultOfTrendingEntry
	ret, err := base_resolver(p, url, config.API.Trending.GetTrendingCategory.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetTrendingEntryDetail(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Trending.GetTrendingEntryDetail.Endpoint, p.Args["trendingEntryType"].(string), p.Args["identifier"].(string))
	var return_val data.TrendingDetail
	ret, err := base_resolver(p, url, config.API.Trending.GetTrendingEntryDetail.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetActivePrivateClanFireteamCount(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Fireteam.GetActivePrivateClanFireteamCount.Endpoint, p.Args["groupId"].(string))
	var return_val int
	ret, err := base_resolver(p, url, config.API.Fireteam.GetActivePrivateClanFireteamCount.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetAvailableClanFireteams(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Fireteam.GetAvailableClanFireteams.Endpoint, p.Args["groupId"].(string), p.Args["platform"].(string), p.Args["activityType"].(string), p.Args["dateRange"].(string), p.Args["slotFilter"].(string), p.Args["publicOnly"].(string), p.Args["page"].(string))
	var return_val data.SearchResultOfFireteamSummary
	ret, err := base_resolver(p, url, config.API.Fireteam.GetAvailableClanFireteams.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func SearchPublicAvailableClanFireteams(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Fireteam.SearchPublicAvailableClanFireteams.Endpoint, p.Args["platform"].(string), p.Args["activityType"].(string), p.Args["dateRange"].(string), p.Args["slotFilter"].(string), p.Args["page"].(string))
	var return_val data.SearchResultOfFireteamSummary
	ret, err := base_resolver(p, url, config.API.Fireteam.SearchPublicAvailableClanFireteams.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetMyClanFireteams(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Fireteam.GetMyClanFireteams.Endpoint, p.Args["groupId"].(string), p.Args["platform"].(string), p.Args["includeClosed"].(string), p.Args["page"].(string))
	var return_val data.SearchResultOfFireteamResponse
	ret, err := base_resolver(p, url, config.API.Fireteam.GetMyClanFireteams.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetClanFireteam(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Fireteam.GetClanFireteam.Endpoint, p.Args["groupId"].(string), p.Args["fireteamId"].(string))
	var return_val data.FireteamResponse
	ret, err := base_resolver(p, url, config.API.Fireteam.GetClanFireteam.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetFriendList(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.GetFriendList.Endpoint)
	var return_val data.BungieFriendListResponse
	ret, err := base_resolver(p, url, config.API.Social.GetFriendList.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetFriendRequestList(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.GetFriendRequestList.Endpoint)
	var return_val data.BungieFriendRequestListResponse
	ret, err := base_resolver(p, url, config.API.Social.GetFriendRequestList.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func IssueFriendRequest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.IssueFriendRequest.Endpoint, p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Social.IssueFriendRequest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func AcceptFriendRequest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.AcceptFriendRequest.Endpoint, p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Social.AcceptFriendRequest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func DeclineFriendRequest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.DeclineFriendRequest.Endpoint, p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Social.DeclineFriendRequest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func RemoveFriend(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.RemoveFriend.Endpoint, p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Social.RemoveFriend.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func RemoveFriendRequest(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.RemoveFriendRequest.Endpoint, p.Args["membershipId"].(string))
	var return_val bool
	ret, err := base_resolver(p, url, config.API.Social.RemoveFriendRequest.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetPlatformFriendList(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Social.GetPlatformFriendList.Endpoint, p.Args["friendPlatform"].(string), p.Args["page"].(string))
	var return_val data.PlatformFriendResponse
	ret, err := base_resolver(p, url, config.API.Social.GetPlatformFriendList.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetAvailableLocales(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Null.GetAvailableLocales.Endpoint)
	var return_val string
	ret, err := base_resolver(p, url, config.API.Null.GetAvailableLocales.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetCommonSettings(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Null.GetCommonSettings.Endpoint)
	var return_val data.CoreSettingsConfiguration
	ret, err := base_resolver(p, url, config.API.Null.GetCommonSettings.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetUserSystemOverrides(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Null.GetUserSystemOverrides.Endpoint)
	var return_val data.CoreSystem
	ret, err := base_resolver(p, url, config.API.Null.GetUserSystemOverrides.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
func GetGlobalAlerts(p graphql.ResolveParams) (interface{}, error) {
	url := fmt.Sprintf(config.API.Null.GetGlobalAlerts.Endpoint)
	var return_val data.GlobalAlert
	ret, err := base_resolver(p, url, config.API.Null.GetGlobalAlerts.Method, &return_val)
	if err != nil {
		log.Println(err)
	}
	return ret, nil
}
