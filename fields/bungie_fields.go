package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/primarchhorus/go-graphql-bungie/resolvers"
	"github.com/primarchhorus/go-graphql-bungie/schema"
)

var (
	Fields = graphql.Fields{
		"getapplicationapiusage": &graphql.Field{
			Type: schema.ApiUsage,
			Args: graphql.FieldConfigArgument{
				"applicationId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetApplicationApiUsage(p)
			},
		},

		"getbungieapplications": &graphql.Field{
			Type: graphql.NewList(schema.Application),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetBungieApplications(p)
			},
		},

		"getbungienetuserbyid": &graphql.Field{
			Type: schema.GeneralUser,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetBungieNetUserById(p)
			},
		},

		"getcredentialtypesfortargetaccount": &graphql.Field{
			Type: graphql.NewList(schema.GetCredentialTypesForAccountResponse),
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCredentialTypesForTargetAccount(p)
			},
		},

		"getavailablethemesuser": &graphql.Field{
			Type: graphql.NewList(schema.UserTheme),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetAvailableThemesUser(p)
			},
		},

		"getmembershipdatabyid": &graphql.Field{
			Type: schema.UserMembershipData,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetMembershipDataById(p)
			},
		},

		"getmembershipdataforcurrentuser": &graphql.Field{
			Type: schema.UserMembershipData,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetMembershipDataForCurrentUser(p)
			},
		},

		"getmembershipfromhardlinkedcredential": &graphql.Field{
			Type: schema.HardLinkedUserMembership,
			Args: graphql.FieldConfigArgument{
				"crType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"credential": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetMembershipFromHardLinkedCredential(p)
			},
		},

		"searchbyglobalnameprefix": &graphql.Field{
			Type: schema.UserSearchResponse,
			Args: graphql.FieldConfigArgument{
				"displayNamePrefix": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchByGlobalNamePrefix(p)
			},
		},

		"getcontenttype": &graphql.Field{
			Type: schema.ContentTypeDescription,
			Args: graphql.FieldConfigArgument{
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetContentType(p)
			},
		},

		"getcontentbyid": &graphql.Field{
			Type: schema.ContentItemPublicContract,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"locale": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetContentById(p)
			},
		},

		"getcontentbytagandtype": &graphql.Field{
			Type: schema.ContentItemPublicContract,
			Args: graphql.FieldConfigArgument{
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"locale": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetContentByTagAndType(p)
			},
		},

		"searchcontentwithtext": &graphql.Field{
			Type: schema.SearchResultOfContentItemPublicContract,
			Args: graphql.FieldConfigArgument{
				"locale": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchContentWithText(p)
			},
		},

		"searchcontentbytagandtype": &graphql.Field{
			Type: schema.SearchResultOfContentItemPublicContract,
			Args: graphql.FieldConfigArgument{
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"locale": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchContentByTagAndType(p)
			},
		},

		"searchhelparticles": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"searchtext": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"size": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchHelpArticles(p)
			},
		},

		"gettopicspaged": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"group": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sort": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"quickDate": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"categoryFilter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetTopicsPaged(p)
			},
		},

		"getcoretopicspaged": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sort": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"quickDate": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"categoryFilter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCoreTopicsPaged(p)
			},
		},

		"getpoststhreadedpaged": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"parentPostId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"replySize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"getParentPost": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rootThreadMode": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortMode": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPostsThreadedPaged(p)
			},
		},

		"getpoststhreadedpagedfromchild": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"childPostId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageSize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"replySize": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rootThreadMode": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortMode": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPostsThreadedPagedFromChild(p)
			},
		},

		"getpostandparent": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"childPostId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPostAndParent(p)
			},
		},

		"getpostandparentawaitingapproval": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"childPostId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPostAndParentAwaitingApproval(p)
			},
		},

		"gettopicforcontent": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"contentId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetTopicForContent(p)
			},
		},

		"getforumtagsuggestions": &graphql.Field{
			Type: graphql.NewList(schema.TagResponse),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetForumTagSuggestions(p)
			},
		},

		"getpoll": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"topicId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPoll(p)
			},
		},

		"getrecruitmentthreadsummaries": &graphql.Field{
			Type: graphql.NewList(schema.ForumRecruitmentDetail),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetRecruitmentThreadSummaries(p)
			},
		},

		"getavailableavatars": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetAvailableAvatars(p)
			},
		},

		// "getavailablethemes": &graphql.Field{
		//     Type: graphql.NewList(schema.GroupTheme),
		//     Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//         return resolvers.GetAvailableThemes(p)
		//     },
		// },

		"getuserclaninvitesetting": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"mType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetUserClanInviteSetting(p)
			},
		},

		"getrecommendedgroups": &graphql.Field{
			Type: graphql.NewList(schema.GroupV2Card),
			Args: graphql.FieldConfigArgument{
				"groupType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"createDateRange": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetRecommendedGroups(p)
			},
		},

		"groupsearch": &graphql.Field{
			Type: schema.GroupSearchResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GroupSearch(p)
			},
		},

		"getgroup": &graphql.Field{
			Type: schema.GroupResponse,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGroup(p)
			},
		},

		"getgroupbyname": &graphql.Field{
			Type: schema.GroupResponse,
			Args: graphql.FieldConfigArgument{
				"groupName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"groupType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGroupByName(p)
			},
		},

		"getgroupbynamev2": &graphql.Field{
			Type: schema.GroupResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGroupByNameV2(p)
			},
		},

		"getgroupoptionalconversations": &graphql.Field{
			Type: graphql.NewList(schema.GroupOptionalConversation),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGroupOptionalConversations(p)
			},
		},

		"editgroup": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EditGroup(p)
			},
		},

		"editclanbanner": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EditClanBanner(p)
			},
		},

		"editfounderoptions": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EditFounderOptions(p)
			},
		},

		"addoptionalconversation": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AddOptionalConversation(p)
			},
		},

		"editoptionalconversation": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"conversationId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EditOptionalConversation(p)
			},
		},

		"getmembersofgroup": &graphql.Field{
			Type: schema.SearchResultOfGroupMember,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetMembersOfGroup(p)
			},
		},

		"getadminsandfounderofgroup": &graphql.Field{
			Type: schema.SearchResultOfGroupMember,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetAdminsAndFounderOfGroup(p)
			},
		},

		"editgroupmembership": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"memberType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EditGroupMembership(p)
			},
		},

		"kickmember": &graphql.Field{
			Type: schema.GroupMemberLeaveResult,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.KickMember(p)
			},
		},

		"banmember": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.BanMember(p)
			},
		},

		"unbanmember": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.UnbanMember(p)
			},
		},

		"getbannedmembersofgroup": &graphql.Field{
			Type: schema.SearchResultOfGroupBan,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetBannedMembersOfGroup(p)
			},
		},

		"abdicatefoundership": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"founderIdNew": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AbdicateFoundership(p)
			},
		},

		"getpendingmemberships": &graphql.Field{
			Type: schema.SearchResultOfGroupMemberApplication,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPendingMemberships(p)
			},
		},

		"getinvitedindividuals": &graphql.Field{
			Type: schema.SearchResultOfGroupMemberApplication,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetInvitedIndividuals(p)
			},
		},

		"approveallpending": &graphql.Field{
			Type: graphql.NewList(schema.EntityActionResult),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ApproveAllPending(p)
			},
		},

		"denyallpending": &graphql.Field{
			Type: graphql.NewList(schema.EntityActionResult),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.DenyAllPending(p)
			},
		},

		"approvependingforlist": &graphql.Field{
			Type: graphql.NewList(schema.EntityActionResult),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ApprovePendingForList(p)
			},
		},

		"approvepending": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ApprovePending(p)
			},
		},

		"denypendingforlist": &graphql.Field{
			Type: graphql.NewList(schema.EntityActionResult),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.DenyPendingForList(p)
			},
		},

		"getgroupsformember": &graphql.Field{
			Type: schema.GetGroupsForMemberResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"groupType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGroupsForMember(p)
			},
		},

		"recovergroupforfounder": &graphql.Field{
			Type: schema.GroupMembershipSearchResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"groupType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.RecoverGroupForFounder(p)
			},
		},

		"getpotentialgroupsformember": &graphql.Field{
			Type: schema.GroupPotentialMembershipSearchResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"filter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"groupType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPotentialGroupsForMember(p)
			},
		},

		"individualgroupinvite": &graphql.Field{
			Type: schema.GroupApplicationResponse,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.IndividualGroupInvite(p)
			},
		},

		"individualgroupinvitecancel": &graphql.Field{
			Type: schema.GroupApplicationResponse,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.IndividualGroupInviteCancel(p)
			},
		},

		"claimpartneroffer": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ClaimPartnerOffer(p)
			},
		},

		"applymissingpartnerofferswithoutclaim": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"partnerApplicationId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"targetBnetMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ApplyMissingPartnerOffersWithoutClaim(p)
			},
		},

		"getpartnerofferskuhistory": &graphql.Field{
			Type: graphql.NewList(schema.PartnerOfferSkuHistoryResponse),
			Args: graphql.FieldConfigArgument{
				"partnerApplicationId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"targetBnetMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPartnerOfferSkuHistory(p)
			},
		},

		"getdestinymanifest": &graphql.Field{
			Type: schema.DestinyManifest,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetDestinyManifest(p)
			},
		},

		"getdestinyentitydefinition": &graphql.Field{
			Type: schema.DestinyDefinition,
			Args: graphql.FieldConfigArgument{
				"entityType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"hashIdentifier": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetDestinyEntityDefinition(p)
			},
		},

		"searchdestinyplayer": &graphql.Field{
			Type: graphql.NewList(schema.UserInfoCard),
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"displayName": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchDestinyPlayer(p)
			},
		},

		"getlinkedprofiles": &graphql.Field{
			Type: schema.DestinyLinkedProfilesResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetLinkedProfiles(p)
			},
		},

		"getprofile": &graphql.Field{
			Type: schema.DestinyProfileResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetProfile(p)
			},
		},

		"getcharacter": &graphql.Field{
			Type: schema.DestinyCharacterResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCharacter(p)
			},
		},

		"getclanweeklyrewardstate": &graphql.Field{
			Type: schema.DestinyMilestone,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetClanWeeklyRewardState(p)
			},
		},

		"getclanbannersource": &graphql.Field{
			Type: schema.ClanBannerSource,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetClanBannerSource(p)
			},
		},

		"getitem": &graphql.Field{
			Type: schema.DestinyItemResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"itemInstanceId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetItem(p)
			},
		},

		"getvendors": &graphql.Field{
			Type: schema.DestinyVendorsResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetVendors(p)
			},
		},

		"getvendor": &graphql.Field{
			Type: schema.DestinyVendorResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"vendorHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetVendor(p)
			},
		},

		"getpublicvendors": &graphql.Field{
			Type: schema.DestinyPublicVendorsResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPublicVendors(p)
			},
		},

		"getcollectiblenodedetails": &graphql.Field{
			Type: schema.DestinyCollectibleNodeDetailResponse,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"collectiblePresentationNodeHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCollectibleNodeDetails(p)
			},
		},

		"transferitem": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.TransferItem(p)
			},
		},

		"pullfrompostmaster": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.PullFromPostmaster(p)
			},
		},

		"equipitem": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EquipItem(p)
			},
		},

		"equipitems": &graphql.Field{
			Type: schema.DestinyEquipItemResults,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.EquipItems(p)
			},
		},

		"setitemlockstate": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SetItemLockState(p)
			},
		},

		"setquesttrackedstate": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SetQuestTrackedState(p)
			},
		},

		"insertsocketplug": &graphql.Field{
			Type: schema.DestinyItemChangeResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.InsertSocketPlug(p)
			},
		},

		"getpostgamecarnagereport": &graphql.Field{
			Type: schema.DestinyPostGameCarnageReportData,
			Args: graphql.FieldConfigArgument{
				"activityId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPostGameCarnageReport(p)
			},
		},

		"reportoffensivepostgamecarnagereportplayer": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"activityId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.ReportOffensivePostGameCarnageReportPlayer(p)
			},
		},

		"gethistoricalstatsdefinition": &graphql.Field{
			Type: schema.DestinyHistoricalStatsDefinition,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetHistoricalStatsDefinition(p)
			},
		},

		"getclanleaderboards": &graphql.Field{
			Type: schema.DestinyLeaderboard,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetClanLeaderboards(p)
			},
		},

		"getclanaggregatestats": &graphql.Field{
			Type: graphql.NewList(schema.DestinyClanAggregateStat),
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetClanAggregateStats(p)
			},
		},

		"getleaderboards": &graphql.Field{
			Type: schema.DestinyLeaderboard,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetLeaderboards(p)
			},
		},

		"getleaderboardsforcharacter": &graphql.Field{
			Type: schema.DestinyLeaderboard,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetLeaderboardsForCharacter(p)
			},
		},

		"searchdestinyentities": &graphql.Field{
			Type: schema.DestinyEntitySearchResult,
			Args: graphql.FieldConfigArgument{
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"searchTerm": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchDestinyEntities(p)
			},
		},

		"gethistoricalstats": &graphql.Field{
			Type: schema.DestinyHistoricalStatsByPeriod,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetHistoricalStats(p)
			},
		},

		"gethistoricalstatsforaccount": &graphql.Field{
			Type: schema.DestinyHistoricalStatsAccountResult,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetHistoricalStatsForAccount(p)
			},
		},

		"getactivityhistory": &graphql.Field{
			Type: schema.DestinyActivityHistoryResults,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetActivityHistory(p)
			},
		},

		"getuniqueweaponhistory": &graphql.Field{
			Type: schema.DestinyHistoricalWeaponStatsData,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetUniqueWeaponHistory(p)
			},
		},

		"getdestinyaggregateactivitystats": &graphql.Field{
			Type: schema.DestinyAggregateActivityResults,
			Args: graphql.FieldConfigArgument{
				"membershipType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"destinyMembershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"characterId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetDestinyAggregateActivityStats(p)
			},
		},

		"getpublicmilestonecontent": &graphql.Field{
			Type: schema.DestinyMilestoneContent,
			Args: graphql.FieldConfigArgument{
				"milestoneHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPublicMilestoneContent(p)
			},
		},

		"getpublicmilestones": &graphql.Field{
			Type: schema.DestinyPublicMilestone,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPublicMilestones(p)
			},
		},

		"awainitializerequest": &graphql.Field{
			Type: schema.AwaInitializeResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AwaInitializeRequest(p)
			},
		},

		"awaprovideauthorizationresult": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AwaProvideAuthorizationResult(p)
			},
		},

		"awagetactiontoken": &graphql.Field{
			Type: schema.AwaAuthorizationResult,
			Args: graphql.FieldConfigArgument{
				"correlationId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AwaGetActionToken(p)
			},
		},

		"getcommunitycontent": &graphql.Field{
			Type: schema.PostSearchResponse,
			Args: graphql.FieldConfigArgument{
				"sort": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"mediaFilter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCommunityContent(p)
			},
		},

		"gettrendingcategories": &graphql.Field{
			Type: schema.TrendingCategories,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetTrendingCategories(p)
			},
		},

		"gettrendingcategory": &graphql.Field{
			Type: schema.SearchResultOfTrendingEntry,
			Args: graphql.FieldConfigArgument{
				"categoryId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pageNumber": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetTrendingCategory(p)
			},
		},

		"gettrendingentrydetail": &graphql.Field{
			Type: schema.TrendingDetail,
			Args: graphql.FieldConfigArgument{
				"trendingEntryType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"identifier": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetTrendingEntryDetail(p)
			},
		},

		"getactiveprivateclanfireteamcount": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetActivePrivateClanFireteamCount(p)
			},
		},

		"getavailableclanfireteams": &graphql.Field{
			Type: schema.SearchResultOfFireteamSummary,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"platform": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"activityType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"dateRange": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"slotFilter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"publicOnly": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetAvailableClanFireteams(p)
			},
		},

		"searchpublicavailableclanfireteams": &graphql.Field{
			Type: schema.SearchResultOfFireteamSummary,
			Args: graphql.FieldConfigArgument{
				"platform": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"activityType": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"dateRange": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"slotFilter": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.SearchPublicAvailableClanFireteams(p)
			},
		},

		"getmyclanfireteams": &graphql.Field{
			Type: schema.SearchResultOfFireteamResponse,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"platform": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"includeClosed": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetMyClanFireteams(p)
			},
		},

		"getclanfireteam": &graphql.Field{
			Type: schema.FireteamResponse,
			Args: graphql.FieldConfigArgument{
				"groupId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"fireteamId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetClanFireteam(p)
			},
		},

		"getfriendlist": &graphql.Field{
			Type: schema.BungieFriendListResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetFriendList(p)
			},
		},

		"getfriendrequestlist": &graphql.Field{
			Type: schema.BungieFriendRequestListResponse,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetFriendRequestList(p)
			},
		},

		"issuefriendrequest": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.IssueFriendRequest(p)
			},
		},

		"acceptfriendrequest": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.AcceptFriendRequest(p)
			},
		},

		"declinefriendrequest": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.DeclineFriendRequest(p)
			},
		},

		"removefriend": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.RemoveFriend(p)
			},
		},

		"removefriendrequest": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"membershipId": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.RemoveFriendRequest(p)
			},
		},

		"getplatformfriendlist": &graphql.Field{
			Type: schema.PlatformFriendResponse,
			Args: graphql.FieldConfigArgument{
				"friendPlatform": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"page": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetPlatformFriendList(p)
			},
		},

		"getavailablelocales": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetAvailableLocales(p)
			},
		},

		"getcommonsettings": &graphql.Field{
			Type: schema.CoreSettingsConfiguration,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetCommonSettings(p)
			},
		},

		"getusersystemoverrides": &graphql.Field{
			Type: schema.CoreSystem,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetUserSystemOverrides(p)
			},
		},

		"getglobalalerts": &graphql.Field{
			Type: graphql.NewList(schema.GlobalAlert),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return resolvers.GetGlobalAlerts(p)
			},
		},
	}
)
