package schema

import (
	"github.com/graphql-go/graphql"
)

var ApplicationScopes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ApplicationScopes",
		Fields: graphql.Fields{
			"ApplicationScopes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ApiUsage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ApiUsage",
		Fields: graphql.Fields{
			"ApiCalls": &graphql.Field{
				Type: graphql.NewList(Series),
			},
			"ThrottledRequests": &graphql.Field{
				Type: graphql.NewList(Series),
			},
		},
	},
)

var Series = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Series",
		Fields: graphql.Fields{
			"Datapoints": &graphql.Field{
				Type: graphql.NewList(Datapoint),
			},
			"Target": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var Datapoint = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Datapoint",
		Fields: graphql.Fields{
			"Time": &graphql.Field{
				Type: graphql.String,
			},
			"Count": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var Application = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Application",
		Fields: graphql.Fields{
			"ApplicationId": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"RedirectUrl": &graphql.Field{
				Type: graphql.String,
			},
			"Link": &graphql.Field{
				Type: graphql.String,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"Origin": &graphql.Field{
				Type: graphql.String,
			},
			"Status": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
			"StatusChanged": &graphql.Field{
				Type: graphql.String,
			},
			"FirstPublished": &graphql.Field{
				Type: graphql.String,
			},
			"Team": &graphql.Field{
				Type: graphql.NewList(ApplicationDeveloper),
			},
			"OverrideAuthorizeViewName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ApplicationStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ApplicationStatus",
		Fields: graphql.Fields{
			"ApplicationStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ApplicationDeveloper = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ApplicationDeveloper",
		Fields: graphql.Fields{
			"Role": &graphql.Field{
				Type: graphql.Int,
			},
			"ApiEulaVersion": &graphql.Field{
				Type: graphql.Int,
			},
			"User": &graphql.Field{
				Type: UserInfoCard,
			},
		},
	},
)

var DeveloperRole = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DeveloperRole",
		Fields: graphql.Fields{
			"DeveloperRole": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UserMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserMembership",
		Fields: graphql.Fields{
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BungieMembershipType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieMembershipType",
		Fields: graphql.Fields{
			"BungieMembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CrossSaveUserMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CrossSaveUserMembership",
		Fields: graphql.Fields{
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UserInfoCard = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserInfoCard",
		Fields: graphql.Fields{
			"SupplementalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GeneralUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GeneralUser",
		Fields: graphql.Fields{
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"UniqueName": &graphql.Field{
				Type: graphql.String,
			},
			"NormalizedName": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"ProfilePicture": &graphql.Field{
				Type: graphql.Int,
			},
			"ProfileTheme": &graphql.Field{
				Type: graphql.Int,
			},
			"UserTitle": &graphql.Field{
				Type: graphql.Int,
			},
			"SuccessMessageFlags": &graphql.Field{
				Type: graphql.Int,
			},
			"IsDeleted": &graphql.Field{
				Type: graphql.Boolean,
			},
			"About": &graphql.Field{
				Type: graphql.String,
			},
			"FirstAccess": &graphql.Field{
				Type: graphql.String,
			},
			"LastUpdate": &graphql.Field{
				Type: graphql.String,
			},
			"LegacyPortalUID": &graphql.Field{
				Type: graphql.Int,
			},
			"Context": &graphql.Field{
				Type: UserToUserContext,
			},
			"PsnDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"XboxDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"FbDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"ShowActivity": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
			"LocaleInheritDefault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LastBanReportId": &graphql.Field{
				Type: graphql.Int,
			},
			"ShowGroupMessaging": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ProfilePicturePath": &graphql.Field{
				Type: graphql.String,
			},
			"ProfilePictureWidePath": &graphql.Field{
				Type: graphql.String,
			},
			"ProfileThemeName": &graphql.Field{
				Type: graphql.String,
			},
			"UserTitleDisplay": &graphql.Field{
				Type: graphql.String,
			},
			"StatusText": &graphql.Field{
				Type: graphql.String,
			},
			"StatusDate": &graphql.Field{
				Type: graphql.String,
			},
			"ProfileBanExpire": &graphql.Field{
				Type: graphql.String,
			},
			"BlizzardDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"SteamDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"StadiaDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"TwitchDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"CachedBungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"CachedBungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UserToUserContext = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserToUserContext",
		Fields: graphql.Fields{
			"IsFollowing": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IgnoreStatus": &graphql.Field{
				Type: IgnoreResponse,
			},
			"GlobalIgnoreEndDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var IgnoreResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IgnoreResponse",
		Fields: graphql.Fields{
			"IsIgnored": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IgnoreFlags": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var IgnoreStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IgnoreStatus",
		Fields: graphql.Fields{
			"IgnoreStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GetCredentialTypesForAccountResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GetCredentialTypesForAccountResponse",
		Fields: graphql.Fields{
			"CredentialType": &graphql.Field{
				Type: graphql.Int,
			},
			"CredentialDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CredentialAsString": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var BungieCredentialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieCredentialType",
		Fields: graphql.Fields{
			"BungieCredentialType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UserTheme = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserTheme",
		Fields: graphql.Fields{
			"UserThemeId": &graphql.Field{
				Type: graphql.Int,
			},
			"UserThemeName": &graphql.Field{
				Type: graphql.String,
			},
			"UserThemeDescription": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var UserMembershipData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserMembershipData",
		Fields: graphql.Fields{
			"DestinyMemberships": &graphql.Field{
				Type: graphql.NewList(GroupUserInfoCard),
			},
			"PrimaryMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetUser": &graphql.Field{
				Type: GeneralUser,
			},
		},
	},
)

var GroupUserInfoCard = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupUserInfoCard",
		Fields: graphql.Fields{
			"LastSeenDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"LastSeenDisplayNameType": &graphql.Field{
				Type: graphql.Int,
			},
			"SupplementalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var HardLinkedUserMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "HardLinkedUserMembership",
		Fields: graphql.Fields{
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"CrossSaveOverriddenType": &graphql.Field{
				Type: graphql.Int,
			},
			"CrossSaveOverriddenMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UserSearchResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserSearchResponse",
		Fields: graphql.Fields{
			"SearchResults": &graphql.Field{
				Type: graphql.NewList(UserSearchResponseDetail),
			},
			"Page": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var UserSearchResponseDetail = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserSearchResponseDetail",
		Fields: graphql.Fields{
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyMemberships": &graphql.Field{
				Type: graphql.NewList(UserInfoCard),
			},
		},
	},
)

var ContentTypeDescription = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentTypeDescription",
		Fields: graphql.Fields{
			"CType": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"ContentDescription": &graphql.Field{
				Type: graphql.String,
			},
			"PreviewImage": &graphql.Field{
				Type: graphql.String,
			},
			"Priority": &graphql.Field{
				Type: graphql.Int,
			},
			"Reminder": &graphql.Field{
				Type: graphql.String,
			},
			"Properties": &graphql.Field{
				Type: graphql.NewList(ContentTypeProperty),
			},
			"TagMetadata": &graphql.Field{
				Type: graphql.NewList(TagMetadataDefinition),
			},
			"TagMetadataItems": &graphql.Field{
				Type: graphql.String,
			},
			"UsageExamples": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"ShowInContentEditor": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TypeOf": &graphql.Field{
				Type: graphql.String,
			},
			"BindIdentifierToProperty": &graphql.Field{
				Type: graphql.String,
			},
			"BoundRegex": &graphql.Field{
				Type: graphql.String,
			},
			"ForceIdentifierBinding": &graphql.Field{
				Type: graphql.Boolean,
			},
			"AllowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
			"AutoEnglishPropertyFallback": &graphql.Field{
				Type: graphql.Boolean,
			},
			"BulkUploadable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Previews": &graphql.Field{
				Type: graphql.NewList(ContentPreview),
			},
			"SuppressCmsPath": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PropertySections": &graphql.Field{
				Type: graphql.NewList(ContentTypePropertySection),
			},
		},
	},
)

var ContentTypeProperty = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentTypeProperty",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"RootPropertyName": &graphql.Field{
				Type: graphql.String,
			},
			"ReadableName": &graphql.Field{
				Type: graphql.String,
			},
			"Value": &graphql.Field{
				Type: graphql.String,
			},
			"PropertyDescription": &graphql.Field{
				Type: graphql.String,
			},
			"Localizable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Fallback": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsTitle": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Required": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MaxLength": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxByteLength": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxFileSize": &graphql.Field{
				Type: graphql.Int,
			},
			"Regexp": &graphql.Field{
				Type: graphql.String,
			},
			"ValidateAs": &graphql.Field{
				Type: graphql.String,
			},
			"RssAttribute": &graphql.Field{
				Type: graphql.String,
			},
			"VisibleDependency": &graphql.Field{
				Type: graphql.String,
			},
			"VisibleOn": &graphql.Field{
				Type: graphql.String,
			},
			"Datatype": &graphql.Field{
				Type: graphql.Int,
			},
			"Attributes": &graphql.Field{
				Type: graphql.String,
			},
			"ChildProperties": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"ContentTypeAllowed": &graphql.Field{
				Type: graphql.String,
			},
			"BindToProperty": &graphql.Field{
				Type: graphql.String,
			},
			"BoundRegex": &graphql.Field{
				Type: graphql.String,
			},
			"RepresentationSelection": &graphql.Field{
				Type: graphql.String,
			},
			"DefaultValues": &graphql.Field{
				Type: graphql.NewList(ContentTypeDefaultValue),
			},
			"IsExternalAllowed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PropertySection": &graphql.Field{
				Type: graphql.String,
			},
			"Weight": &graphql.Field{
				Type: graphql.Int,
			},
			"Entitytype": &graphql.Field{
				Type: graphql.String,
			},
			"IsCombo": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SuppressProperty": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LegalContentTypes": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"RepresentationValidationString": &graphql.Field{
				Type: graphql.String,
			},
			"MinWidth": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxWidth": &graphql.Field{
				Type: graphql.Int,
			},
			"MinHeight": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxHeight": &graphql.Field{
				Type: graphql.Int,
			},
			"IsVideo": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsImage": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ContentPropertyDataTypeEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentPropertyDataTypeEnum",
		Fields: graphql.Fields{
			"ContentPropertyDataTypeEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ContentTypeDefaultValue = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentTypeDefaultValue",
		Fields: graphql.Fields{
			"WhenClause": &graphql.Field{
				Type: graphql.String,
			},
			"WhenValue": &graphql.Field{
				Type: graphql.String,
			},
			"DefaultValue": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TagMetadataDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TagMetadataDefinition",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(TagMetadataItem),
			},
			"Datatype": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"IsRequired": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var TagMetadataItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TagMetadataItem",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"TagText": &graphql.Field{
				Type: graphql.String,
			},
			"Groups": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"IsDefault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ContentPreview = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentPreview",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Path": &graphql.Field{
				Type: graphql.String,
			},
			"ItemInSet": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SetTag": &graphql.Field{
				Type: graphql.String,
			},
			"SetNesting": &graphql.Field{
				Type: graphql.Int,
			},
			"UseSetId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ContentTypePropertySection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentTypePropertySection",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"ReadableName": &graphql.Field{
				Type: graphql.String,
			},
			"Collapsed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ContentItemPublicContract = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentItemPublicContract",
		Fields: graphql.Fields{
			"ContentId": &graphql.Field{
				Type: graphql.Int,
			},
			"CType": &graphql.Field{
				Type: graphql.String,
			},
			"CmsPath": &graphql.Field{
				Type: graphql.String,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
			"ModifyDate": &graphql.Field{
				Type: graphql.String,
			},
			"AllowComments": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HasAgeGate": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MinimumAge": &graphql.Field{
				Type: graphql.Int,
			},
			"RatingImagePath": &graphql.Field{
				Type: graphql.String,
			},
			"Author": &graphql.Field{
				Type: GeneralUser,
			},
			"AutoEnglishPropertyFallback": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Properties": &graphql.Field{
				Type: graphql.String,
			},
			"Representations": &graphql.Field{
				Type: graphql.NewList(ContentRepresentation),
			},
			"Tags": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"CommentSummary": &graphql.Field{
				Type: CommentSummary,
			},
		},
	},
)

var ContentRepresentation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ContentRepresentation",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Path": &graphql.Field{
				Type: graphql.String,
			},
			"ValidationString": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CommentSummary = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CommentSummary",
		Fields: graphql.Fields{
			"TopicId": &graphql.Field{
				Type: graphql.Int,
			},
			"CommentCount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResult",
		Fields: graphql.Fields{
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PagedQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PagedQuery",
		Fields: graphql.Fields{
			"ItemsPerPage": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentPage": &graphql.Field{
				Type: graphql.Int,
			},
			"RequestContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SearchResultOfContentItemPublicContract = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfContentItemPublicContract",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(ContentItemPublicContract),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ForumTopicsCategoryFiltersEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumTopicsCategoryFiltersEnum",
		Fields: graphql.Fields{
			"ForumTopicsCategoryFiltersEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumTopicsQuickDateEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumTopicsQuickDateEnum",
		Fields: graphql.Fields{
			"ForumTopicsQuickDateEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumTopicsSortEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumTopicsSortEnum",
		Fields: graphql.Fields{
			"ForumTopicsSortEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PostResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostResponse",
		Fields: graphql.Fields{
			"LastReplyTimestamp": &graphql.Field{
				Type: graphql.String,
			},
			"IsPinned": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UrlMediaType": &graphql.Field{
				Type: graphql.Int,
			},
			"Thumbnail": &graphql.Field{
				Type: graphql.String,
			},
			"Popularity": &graphql.Field{
				Type: graphql.Int,
			},
			"IsActive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsAnnouncement": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UserRating": &graphql.Field{
				Type: graphql.Int,
			},
			"UserHasRated": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UserHasMutedPost": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LatestReplyPostId": &graphql.Field{
				Type: graphql.Int,
			},
			"LatestReplyAuthorId": &graphql.Field{
				Type: graphql.Int,
			},
			"IgnoreStatus": &graphql.Field{
				Type: IgnoreResponse,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ForumMediaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumMediaType",
		Fields: graphql.Fields{
			"ForumMediaType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumPostPopularity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumPostPopularity",
		Fields: graphql.Fields{
			"ForumPostPopularity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumPostCategoryEnums = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumPostCategoryEnums",
		Fields: graphql.Fields{
			"ForumPostCategoryEnums": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumFlagsEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumFlagsEnum",
		Fields: graphql.Fields{
			"ForumFlagsEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResultOfPostResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfPostResponse",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(PostResponse),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PostSearchResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostSearchResponse",
		Fields: graphql.Fields{
			"RelatedPosts": &graphql.Field{
				Type: graphql.NewList(PostResponse),
			},
			"Authors": &graphql.Field{
				Type: graphql.NewList(GeneralUser),
			},
			"Groups": &graphql.Field{
				Type: graphql.NewList(GroupResponse),
			},
			"SearchedTags": &graphql.Field{
				Type: graphql.NewList(TagResponse),
			},
			"Polls": &graphql.Field{
				Type: graphql.NewList(PollResponse),
			},
			"RecruitmentDetails": &graphql.Field{
				Type: graphql.NewList(ForumRecruitmentDetail),
			},
			"AvailablePages": &graphql.Field{
				Type: graphql.Int,
			},
			"Results": &graphql.Field{
				Type: graphql.NewList(PostResponse),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupResponse",
		Fields: graphql.Fields{
			"Detail": &graphql.Field{
				Type: GroupV2,
			},
			"Founder": &graphql.Field{
				Type: GroupMember,
			},
			"AlliedIds": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentGroup": &graphql.Field{
				Type: GroupV2,
			},
			"AllianceStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"GroupJoinInviteCount": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentUserMembershipsInactiveForDestiny": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CurrentUserMemberMap": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentUserPotentialMemberMap": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupV2 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupV2",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"GroupType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipIdCreated": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
			"ModificationDate": &graphql.Field{
				Type: graphql.String,
			},
			"About": &graphql.Field{
				Type: graphql.String,
			},
			"Tags": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"MemberCount": &graphql.Field{
				Type: graphql.Int,
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsPublicTopicAdminOnly": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Motto": &graphql.Field{
				Type: graphql.String,
			},
			"AllowChat": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsDefaultPostPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ChatSecurity": &graphql.Field{
				Type: graphql.Int,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
			"AvatarImageIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"Homepage": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipOption": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultPublicity": &graphql.Field{
				Type: graphql.Int,
			},
			"Theme": &graphql.Field{
				Type: graphql.String,
			},
			"BannerPath": &graphql.Field{
				Type: graphql.String,
			},
			"AvatarPath": &graphql.Field{
				Type: graphql.String,
			},
			"ConversationId": &graphql.Field{
				Type: graphql.Int,
			},
			"EnableInvitationMessagingForAdmins": &graphql.Field{
				Type: graphql.Boolean,
			},
			"BanExpireDate": &graphql.Field{
				Type: graphql.String,
			},
			"Features": &graphql.Field{
				Type: GroupFeatures,
			},
			"ClanInfo": &graphql.Field{
				Type: GroupV2ClanInfoAndInvestment,
			},
		},
	},
)

var GroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupType",
		Fields: graphql.Fields{
			"GroupType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ChatSecuritySetting = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ChatSecuritySetting",
		Fields: graphql.Fields{
			"ChatSecuritySetting": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupHomepage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupHomepage",
		Fields: graphql.Fields{
			"GroupHomepage": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var MembershipOption = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MembershipOption",
		Fields: graphql.Fields{
			"MembershipOption": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupPostPublicity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupPostPublicity",
		Fields: graphql.Fields{
			"GroupPostPublicity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupFeatures = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupFeatures",
		Fields: graphql.Fields{
			"MaximumMembers": &graphql.Field{
				Type: graphql.Int,
			},
			"MaximumMembershipsOfGroupType": &graphql.Field{
				Type: graphql.Int,
			},
			"Capabilities": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"InvitePermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UpdateCulturePermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HostGuidedGamePermissionOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"UpdateBannerPermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"JoinLevel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var Capabilities = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Capabilities",
		Fields: graphql.Fields{
			"Capabilities": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BungieMembershipTypeList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieMembershipTypeList",
		Fields: graphql.Fields{
			"BungieMembershipType": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var HostGuidedGamesPermissionLevel = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "HostGuidedGamesPermissionLevel",
		Fields: graphql.Fields{
			"HostGuidedGamesPermissionLevel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var RuntimeGroupMemberType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RuntimeGroupMemberType",
		Fields: graphql.Fields{
			"RuntimeGroupMemberType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupV2ClanInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupV2ClanInfo",
		Fields: graphql.Fields{
			"ClanCallsign": &graphql.Field{
				Type: graphql.String,
			},
			"ClanBannerData": &graphql.Field{
				Type: ClanBanner,
			},
		},
	},
)

var ClanBanner = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClanBanner",
		Fields: graphql.Fields{
			"DecalId": &graphql.Field{
				Type: graphql.Int,
			},
			"DecalColorId": &graphql.Field{
				Type: graphql.Int,
			},
			"DecalBackgroundColorId": &graphql.Field{
				Type: graphql.Int,
			},
			"GonfalonId": &graphql.Field{
				Type: graphql.Int,
			},
			"GonfalonColorId": &graphql.Field{
				Type: graphql.Int,
			},
			"GonfalonDetailId": &graphql.Field{
				Type: graphql.Int,
			},
			"GonfalonDetailColorId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupV2ClanInfoAndInvestment = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupV2ClanInfoAndInvestment",
		Fields: graphql.Fields{
			"D2ClanProgressions": &graphql.Field{
				Type: graphql.Int,
			},
			"ClanCallsign": &graphql.Field{
				Type: graphql.String,
			},
			"ClanBannerData": &graphql.Field{
				Type: ClanBanner,
			},
		},
	},
)

var DestinyProgression = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgression",
		Fields: graphql.Fields{
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DailyProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"DailyLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"WeeklyProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"WeeklyLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"Level": &graphql.Field{
				Type: graphql.Int,
			},
			"LevelCap": &graphql.Field{
				Type: graphql.Int,
			},
			"StepIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressToNextLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"NextLevelAt": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentResetCount": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonResets": &graphql.Field{
				Type: graphql.NewList(DestinyProgressionResetEntry),
			},
			"RewardItemStates": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyProgressionResetEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionResetEntry",
		Fields: graphql.Fields{
			"Season": &graphql.Field{
				Type: graphql.Int,
			},
			"Resets": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProgressionRewardItemState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionRewardItemState",
		Fields: graphql.Fields{
			"DestinyProgressionRewardItemState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDefinition",
		Fields: graphql.Fields{
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyProgressionDisplayPropertiesDefinition,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"RepeatLastStep": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Source": &graphql.Field{
				Type: graphql.String,
			},
			"Steps": &graphql.Field{
				Type: graphql.NewList(DestinyProgressionStepDefinition),
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"FactionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Color": &graphql.Field{
				Type: DestinyColor,
			},
			"RankIcon": &graphql.Field{
				Type: graphql.String,
			},
			"RewardItems": &graphql.Field{
				Type: graphql.NewList(DestinyProgressionRewardItemQuantity),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyDisplayPropertiesDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDisplayPropertiesDefinition",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"IconSequences": &graphql.Field{
				Type: graphql.NewList(DestinyIconSequenceDefinition),
			},
			"HighResIcon": &graphql.Field{
				Type: graphql.String,
			},
			"HasIcon": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyIconSequenceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyIconSequenceDefinition",
		Fields: graphql.Fields{
			"Frames": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var DestinyProgressionDisplayPropertiesDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionDisplayPropertiesDefinition",
		Fields: graphql.Fields{
			"DisplayUnitsName": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"IconSequences": &graphql.Field{
				Type: graphql.NewList(DestinyIconSequenceDefinition),
			},
			"HighResIcon": &graphql.Field{
				Type: graphql.String,
			},
			"HasIcon": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionScope = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionScope",
		Fields: graphql.Fields{
			"DestinyProgressionScope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProgressionStepDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionStepDefinition",
		Fields: graphql.Fields{
			"StepName": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayEffectType": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressTotal": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyProgressionStepDisplayEffect = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionStepDisplayEffect",
		Fields: graphql.Fields{
			"DestinyProgressionStepDisplayEffect": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemQuantity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemQuantity",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"HasConditionalVisibility": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyInventoryItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInventoryItemDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TooltipNotifications": &graphql.Field{
				Type: graphql.NewList(DestinyItemTooltipNotification),
			},
			"CollectibleHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IconWatermark": &graphql.Field{
				Type: graphql.String,
			},
			"IconWatermarkShelved": &graphql.Field{
				Type: graphql.String,
			},
			"SecondaryIcon": &graphql.Field{
				Type: graphql.String,
			},
			"SecondaryOverlay": &graphql.Field{
				Type: graphql.String,
			},
			"SecondarySpecial": &graphql.Field{
				Type: graphql.String,
			},
			"BackgroundColor": &graphql.Field{
				Type: DestinyColor,
			},
			"Screenshot": &graphql.Field{
				Type: graphql.String,
			},
			"ItemTypeDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"FlavorText": &graphql.Field{
				Type: graphql.String,
			},
			"UiItemDisplayStyle": &graphql.Field{
				Type: graphql.String,
			},
			"ItemTypeAndTierDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"DisplaySource": &graphql.Field{
				Type: graphql.String,
			},
			"TooltipStyle": &graphql.Field{
				Type: graphql.String,
			},
			"Action": &graphql.Field{
				Type: DestinyItemActionBlockDefinition,
			},
			"Inventory": &graphql.Field{
				Type: DestinyItemInventoryBlockDefinition,
			},
			"SetData": &graphql.Field{
				Type: DestinyItemSetBlockDefinition,
			},
			"Stats": &graphql.Field{
				Type: DestinyItemStatBlockDefinition,
			},
			"EmblemObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EquippingBlock": &graphql.Field{
				Type: DestinyEquippingBlockDefinition,
			},
			"TranslationBlock": &graphql.Field{
				Type: DestinyItemTranslationBlockDefinition,
			},
			"Preview": &graphql.Field{
				Type: DestinyItemPreviewBlockDefinition,
			},
			"Quality": &graphql.Field{
				Type: DestinyItemQualityBlockDefinition,
			},
			"Value": &graphql.Field{
				Type: DestinyItemValueBlockDefinition,
			},
			"SourceData": &graphql.Field{
				Type: DestinyItemSourceBlockDefinition,
			},
			"Objectives": &graphql.Field{
				Type: DestinyItemObjectiveBlockDefinition,
			},
			"Metrics": &graphql.Field{
				Type: DestinyItemMetricBlockDefinition,
			},
			"Plug": &graphql.Field{
				Type: DestinyItemPlugDefinition,
			},
			"Gearset": &graphql.Field{
				Type: DestinyItemGearsetBlockDefinition,
			},
			"Sack": &graphql.Field{
				Type: DestinyItemSackBlockDefinition,
			},
			"Sockets": &graphql.Field{
				Type: DestinyItemSocketBlockDefinition,
			},
			"Summary": &graphql.Field{
				Type: DestinyItemSummaryBlockDefinition,
			},
			"TalentGrid": &graphql.Field{
				Type: DestinyItemTalentGridBlockDefinition,
			},
			"InvestmentStats": &graphql.Field{
				Type: graphql.NewList(DestinyItemInvestmentStatDefinition),
			},
			"Perks": &graphql.Field{
				Type: graphql.NewList(DestinyItemPerkEntryDefinition),
			},
			"LoreHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SummaryItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Animations": &graphql.Field{
				Type: graphql.NewList(DestinyAnimationReference),
			},
			"AllowActions": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Links": &graphql.Field{
				Type: graphql.NewList(HyperlinkReference),
			},
			"DoesPostmasterPullHaveSideEffects": &graphql.Field{
				Type: graphql.Boolean,
			},
			"NonTransferrable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ItemCategoryHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"SpecialItemType": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemType": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemSubType": &graphql.Field{
				Type: graphql.Int,
			},
			"ClassType": &graphql.Field{
				Type: graphql.Int,
			},
			"BreakerType": &graphql.Field{
				Type: graphql.Int,
			},
			"BreakerTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Equippable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DamageTypeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"DamageTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"DefaultDamageType": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultDamageTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsWrapper": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemTooltipNotification = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTooltipNotification",
		Fields: graphql.Fields{
			"DisplayString": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayStyle": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyColor = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyColor",
		Fields: graphql.Fields{
			"Red": &graphql.Field{
				Type: graphql.String,
			},
			"Green": &graphql.Field{
				Type: graphql.String,
			},
			"Blue": &graphql.Field{
				Type: graphql.String,
			},
			"Alpha": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyItemActionBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemActionBlockDefinition",
		Fields: graphql.Fields{
			"VerbName": &graphql.Field{
				Type: graphql.String,
			},
			"VerbDescription": &graphql.Field{
				Type: graphql.String,
			},
			"IsPositive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"OverlayScreenName": &graphql.Field{
				Type: graphql.String,
			},
			"OverlayIcon": &graphql.Field{
				Type: graphql.String,
			},
			"RequiredCooldownSeconds": &graphql.Field{
				Type: graphql.Int,
			},
			"RequiredItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemActionRequiredItemDefinition),
			},
			"ProgressionRewards": &graphql.Field{
				Type: graphql.NewList(DestinyProgressionRewardDefinition),
			},
			"ActionTypeLabel": &graphql.Field{
				Type: graphql.String,
			},
			"RequiredLocation": &graphql.Field{
				Type: graphql.String,
			},
			"RequiredCooldownHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DeleteOnAction": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ConsumeEntireStack": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UseOnAcquire": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemActionRequiredItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemActionRequiredItemDefinition",
		Fields: graphql.Fields{
			"Count": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DeleteOnAction": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionRewardDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionRewardDefinition",
		Fields: graphql.Fields{
			"ProgressionMappingHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Amount": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplyThrottles": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionMappingDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionMappingDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"DisplayUnits": &graphql.Field{
				Type: graphql.String,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemInventoryBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemInventoryBlockDefinition",
		Fields: graphql.Fields{
			"StackUniqueLabel": &graphql.Field{
				Type: graphql.String,
			},
			"MaxStackSize": &graphql.Field{
				Type: graphql.Int,
			},
			"BucketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RecoveryBucketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"TierTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsInstanceItem": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TierTypeName": &graphql.Field{
				Type: graphql.String,
			},
			"TierType": &graphql.Field{
				Type: graphql.Int,
			},
			"ExpirationTooltip": &graphql.Field{
				Type: graphql.String,
			},
			"ExpiredInActivityMessage": &graphql.Field{
				Type: graphql.String,
			},
			"ExpiredInOrbitMessage": &graphql.Field{
				Type: graphql.String,
			},
			"SuppressExpirationWhenObjectivesComplete": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var TierType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TierType",
		Fields: graphql.Fields{
			"TierType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyInventoryBucketDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInventoryBucketDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"Category": &graphql.Field{
				Type: graphql.Int,
			},
			"BucketOrder": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemCount": &graphql.Field{
				Type: graphql.Int,
			},
			"Location": &graphql.Field{
				Type: graphql.Int,
			},
			"HasTransferDestination": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Fifo": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var BucketScope = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BucketScope",
		Fields: graphql.Fields{
			"BucketScope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BucketCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BucketCategory",
		Fields: graphql.Fields{
			"BucketCategory": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ItemLocation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemLocation",
		Fields: graphql.Fields{
			"ItemLocation": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemTierTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTierTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"InfusionProcess": &graphql.Field{
				Type: DestinyItemTierTypeInfusionBlock,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemTierTypeInfusionBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTierTypeInfusionBlock",
		Fields: graphql.Fields{
			"BaseQualityTransferRatio": &graphql.Field{
				Type: graphql.Float,
			},
			"MinimumQualityIncrement": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemSetBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSetBlockDefinition",
		Fields: graphql.Fields{
			"ItemList": &graphql.Field{
				Type: graphql.NewList(DestinyItemSetBlockEntryDefinition),
			},
			"RequireOrderedSetItemAdd": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SetIsFeatured": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SetType": &graphql.Field{
				Type: graphql.String,
			},
			"QuestLineName": &graphql.Field{
				Type: graphql.String,
			},
			"QuestLineDescription": &graphql.Field{
				Type: graphql.String,
			},
			"QuestStepSummary": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyItemSetBlockEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSetBlockEntryDefinition",
		Fields: graphql.Fields{
			"TrackingValue": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemStatBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemStatBlockDefinition",
		Fields: graphql.Fields{
			"DisablePrimaryStatDisplay": &graphql.Field{
				Type: graphql.Boolean,
			},
			"StatGroupHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Stats": &graphql.Field{
				Type: graphql.Int,
			},
			"HasDisplayableStats": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PrimaryBaseStatHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyInventoryItemStatDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInventoryItemStatDefinition",
		Fields: graphql.Fields{
			"StatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
			"Minimum": &graphql.Field{
				Type: graphql.Int,
			},
			"Maximum": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayMaximum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"AggregationType": &graphql.Field{
				Type: graphql.Int,
			},
			"HasComputedBlock": &graphql.Field{
				Type: graphql.Boolean,
			},
			"StatCategory": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyStatAggregationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatAggregationType",
		Fields: graphql.Fields{
			"DestinyStatAggregationType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatCategory",
		Fields: graphql.Fields{
			"DestinyStatCategory": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatGroupDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatGroupDefinition",
		Fields: graphql.Fields{
			"MaximumValue": &graphql.Field{
				Type: graphql.Int,
			},
			"UiPosition": &graphql.Field{
				Type: graphql.Int,
			},
			"ScaledStats": &graphql.Field{
				Type: graphql.NewList(DestinyStatDisplayDefinition),
			},
			"Overrides": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyStatDisplayDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatDisplayDefinition",
		Fields: graphql.Fields{
			"StatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"MaximumValue": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayAsNumeric": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DisplayInterpolation": &graphql.Field{
				Type: graphql.NewList(InterpolationPoint),
			},
		},
	},
)

var InterpolationPoint = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InterpolationPoint",
		Fields: graphql.Fields{
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
			"Weight": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatOverrideDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatOverrideDefinition",
		Fields: graphql.Fields{
			"StatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
		},
	},
)

var DestinyEquippingBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEquippingBlockDefinition",
		Fields: graphql.Fields{
			"GearsetItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"UniqueLabel": &graphql.Field{
				Type: graphql.String,
			},
			"UniqueLabelHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EquipmentSlotTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Attributes": &graphql.Field{
				Type: graphql.Int,
			},
			"AmmoType": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayStrings": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var EquippingItemBlockAttributes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EquippingItemBlockAttributes",
		Fields: graphql.Fields{
			"EquippingItemBlockAttributes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyAmmunitionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyAmmunitionType",
		Fields: graphql.Fields{
			"DestinyAmmunitionType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyEquipmentSlotDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEquipmentSlotDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"EquipmentCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"BucketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplyCustomArtDyes": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ArtDyeChannels": &graphql.Field{
				Type: graphql.NewList(DestinyArtDyeReference),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyArtDyeReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtDyeReference",
		Fields: graphql.Fields{
			"ArtDyeChannelHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemTranslationBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTranslationBlockDefinition",
		Fields: graphql.Fields{
			"WeaponPatternIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"WeaponPatternHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultDyes": &graphql.Field{
				Type: graphql.NewList(DyeReference),
			},
			"LockedDyes": &graphql.Field{
				Type: graphql.NewList(DyeReference),
			},
			"CustomDyes": &graphql.Field{
				Type: graphql.NewList(DyeReference),
			},
			"Arrangements": &graphql.Field{
				Type: graphql.NewList(DestinyGearArtArrangementReference),
			},
			"HasGeometry": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DyeReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DyeReference",
		Fields: graphql.Fields{
			"ChannelHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DyeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyGearArtArrangementReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGearArtArrangementReference",
		Fields: graphql.Fields{
			"ClassHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ArtArrangementHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemPreviewBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPreviewBlockDefinition",
		Fields: graphql.Fields{
			"ScreenStyle": &graphql.Field{
				Type: graphql.String,
			},
			"PreviewVendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ArtifactHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PreviewActionString": &graphql.Field{
				Type: graphql.String,
			},
			"DerivedItemCategories": &graphql.Field{
				Type: graphql.NewList(DestinyDerivedItemCategoryDefinition),
			},
		},
	},
)

var DestinyDerivedItemCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDerivedItemCategoryDefinition",
		Fields: graphql.Fields{
			"CategoryDescription": &graphql.Field{
				Type: graphql.String,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyDerivedItemDefinition),
			},
		},
	},
)

var DestinyDerivedItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDerivedItemDefinition",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemName": &graphql.Field{
				Type: graphql.String,
			},
			"ItemDetail": &graphql.Field{
				Type: graphql.String,
			},
			"ItemDescription": &graphql.Field{
				Type: graphql.String,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyVendorDisplayPropertiesDefinition,
			},
			"VendorProgressionType": &graphql.Field{
				Type: graphql.Int,
			},
			"BuyString": &graphql.Field{
				Type: graphql.String,
			},
			"SellString": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"InhibitBuying": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InhibitSelling": &graphql.Field{
				Type: graphql.Boolean,
			},
			"FactionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ResetIntervalMinutes": &graphql.Field{
				Type: graphql.Int,
			},
			"ResetOffsetMinutes": &graphql.Field{
				Type: graphql.Int,
			},
			"FailureStrings": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"UnlockRanges": &graphql.Field{
				Type: graphql.NewList(DateRange),
			},
			"VendorIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"VendorPortrait": &graphql.Field{
				Type: graphql.String,
			},
			"VendorBanner": &graphql.Field{
				Type: graphql.String,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"VendorSubcategoryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"ConsolidateCategories": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Actions": &graphql.Field{
				Type: graphql.NewList(DestinyVendorActionDefinition),
			},
			"Categories": &graphql.Field{
				Type: graphql.NewList(DestinyVendorCategoryEntryDefinition),
			},
			"OriginalCategories": &graphql.Field{
				Type: graphql.NewList(DestinyVendorCategoryEntryDefinition),
			},
			"DisplayCategories": &graphql.Field{
				Type: graphql.NewList(DestinyDisplayCategoryDefinition),
			},
			"Interactions": &graphql.Field{
				Type: graphql.NewList(DestinyVendorInteractionDefinition),
			},
			"InventoryFlyouts": &graphql.Field{
				Type: graphql.NewList(DestinyVendorInventoryFlyoutDefinition),
			},
			"ItemList": &graphql.Field{
				Type: graphql.NewList(DestinyVendorItemDefinition),
			},
			"Services": &graphql.Field{
				Type: graphql.NewList(DestinyVendorServiceDefinition),
			},
			"AcceptedItems": &graphql.Field{
				Type: graphql.NewList(DestinyVendorAcceptedItemDefinition),
			},
			"ReturnWithVendorRequest": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Locations": &graphql.Field{
				Type: graphql.NewList(DestinyVendorLocationDefinition),
			},
			"Groups": &graphql.Field{
				Type: graphql.NewList(DestinyVendorGroupReference),
			},
			"IgnoreSaleItemHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorDisplayPropertiesDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorDisplayPropertiesDefinition",
		Fields: graphql.Fields{
			"LargeIcon": &graphql.Field{
				Type: graphql.String,
			},
			"Subtitle": &graphql.Field{
				Type: graphql.String,
			},
			"OriginalIcon": &graphql.Field{
				Type: graphql.String,
			},
			"RequirementsDisplay": &graphql.Field{
				Type: graphql.NewList(DestinyVendorRequirementDisplayEntryDefinition),
			},
			"SmallTransparentIcon": &graphql.Field{
				Type: graphql.String,
			},
			"MapIcon": &graphql.Field{
				Type: graphql.String,
			},
			"LargeTransparentIcon": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"IconSequences": &graphql.Field{
				Type: graphql.NewList(DestinyIconSequenceDefinition),
			},
			"HighResIcon": &graphql.Field{
				Type: graphql.String,
			},
			"HasIcon": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorRequirementDisplayEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorRequirementDisplayEntryDefinition",
		Fields: graphql.Fields{
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Source": &graphql.Field{
				Type: graphql.String,
			},
			"Type": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyVendorProgressionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorProgressionType",
		Fields: graphql.Fields{
			"DestinyVendorProgressionType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DateRange = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DateRange",
		Fields: graphql.Fields{
			"Start": &graphql.Field{
				Type: graphql.String,
			},
			"End": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyVendorActionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorActionDefinition",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"ExecuteSeconds": &graphql.Field{
				Type: graphql.Int,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Verb": &graphql.Field{
				Type: graphql.String,
			},
			"IsPositive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ActionId": &graphql.Field{
				Type: graphql.String,
			},
			"ActionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AutoPerformAction": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorCategoryEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorCategoryEntryDefinition",
		Fields: graphql.Fields{
			"CategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"SortValue": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"QuantityAvailable": &graphql.Field{
				Type: graphql.Int,
			},
			"ShowUnavailableItems": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HideIfNoCurrency": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HideFromRegularPurchase": &graphql.Field{
				Type: graphql.Boolean,
			},
			"BuyStringOverride": &graphql.Field{
				Type: graphql.String,
			},
			"DisabledDescription": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayTitle": &graphql.Field{
				Type: graphql.String,
			},
			"Overlay": &graphql.Field{
				Type: DestinyVendorCategoryOverlayDefinition,
			},
			"VendorItemIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPreview": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsDisplayOnly": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ResetIntervalMinutesOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ResetOffsetMinutesOverride": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorCategoryOverlayDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorCategoryOverlayDefinition",
		Fields: graphql.Fields{
			"ChoiceDescription": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"CurrencyItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyDisplayCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDisplayCategoryDefinition",
		Fields: graphql.Fields{
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"DisplayInBanner": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SortOrder": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayStyleHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayStyleIdentifier": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var VendorDisplayCategorySortOrder = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VendorDisplayCategorySortOrder",
		Fields: graphql.Fields{
			"VendorDisplayCategorySortOrder": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorInteractionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInteractionDefinition",
		Fields: graphql.Fields{
			"InteractionIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"Replies": &graphql.Field{
				Type: graphql.NewList(DestinyVendorInteractionReplyDefinition),
			},
			"VendorCategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"QuestlineItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SackInteractionList": &graphql.Field{
				Type: graphql.NewList(DestinyVendorInteractionSackEntryDefinition),
			},
			"UiInteractionType": &graphql.Field{
				Type: graphql.Int,
			},
			"InteractionType": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardBlockLabel": &graphql.Field{
				Type: graphql.String,
			},
			"RewardVendorCategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"FlavorLineOne": &graphql.Field{
				Type: graphql.String,
			},
			"FlavorLineTwo": &graphql.Field{
				Type: graphql.String,
			},
			"HeaderDisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Instructions": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyVendorInteractionReplyDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInteractionReplyDefinition",
		Fields: graphql.Fields{
			"ItemRewardsSelection": &graphql.Field{
				Type: graphql.Int,
			},
			"Reply": &graphql.Field{
				Type: graphql.String,
			},
			"ReplyType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorInteractionRewardSelection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInteractionRewardSelection",
		Fields: graphql.Fields{
			"DestinyVendorInteractionRewardSelection": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorReplyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorReplyType",
		Fields: graphql.Fields{
			"DestinyVendorReplyType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorInteractionSackEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInteractionSackEntryDefinition",
		Fields: graphql.Fields{
			"SackType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var VendorInteractionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VendorInteractionType",
		Fields: graphql.Fields{
			"VendorInteractionType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorInventoryFlyoutDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInventoryFlyoutDefinition",
		Fields: graphql.Fields{
			"LockedDescription": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Buckets": &graphql.Field{
				Type: graphql.NewList(DestinyVendorInventoryFlyoutBucketDefinition),
			},
			"FlyoutId": &graphql.Field{
				Type: graphql.Int,
			},
			"SuppressNewness": &graphql.Field{
				Type: graphql.Boolean,
			},
			"EquipmentSlotHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorInventoryFlyoutBucketDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorInventoryFlyoutBucketDefinition",
		Fields: graphql.Fields{
			"Collapsible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InventoryBucketHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SortItemsBy": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemSortType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSortType",
		Fields: graphql.Fields{
			"DestinyItemSortType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorItemDefinition",
		Fields: graphql.Fields{
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"FailureIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Currencies": &graphql.Field{
				Type: graphql.NewList(DestinyVendorItemQuantity),
			},
			"RefundPolicy": &graphql.Field{
				Type: graphql.Int,
			},
			"RefundTimeLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationLevels": &graphql.Field{
				Type: graphql.NewList(DestinyItemCreationEntryLevelDefinition),
			},
			"DisplayCategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"OriginalCategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"MinimumLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"MaximumLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"Action": &graphql.Field{
				Type: DestinyVendorSaleItemActionBlockDefinition,
			},
			"DisplayCategory": &graphql.Field{
				Type: graphql.String,
			},
			"InventoryBucketHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VisibilityScope": &graphql.Field{
				Type: graphql.Int,
			},
			"PurchasableScope": &graphql.Field{
				Type: graphql.Int,
			},
			"Exclusivity": &graphql.Field{
				Type: graphql.Int,
			},
			"IsOffer": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsCrm": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SortValue": &graphql.Field{
				Type: graphql.Int,
			},
			"ExpirationTooltip": &graphql.Field{
				Type: graphql.String,
			},
			"RedirectToSaleIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"SocketOverrides": &graphql.Field{
				Type: graphql.NewList(DestinyVendorItemSocketOverride),
			},
			"Unpurchasable": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorItemQuantity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorItemQuantity",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"HasConditionalVisibility": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorItemRefundPolicy = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorItemRefundPolicy",
		Fields: graphql.Fields{
			"DestinyVendorItemRefundPolicy": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemCreationEntryLevelDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemCreationEntryLevelDefinition",
		Fields: graphql.Fields{
			"Level": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorSaleItemActionBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorSaleItemActionBlockDefinition",
		Fields: graphql.Fields{
			"ExecuteSeconds": &graphql.Field{
				Type: graphql.Float,
			},
			"IsPositive": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyGatingScope = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGatingScope",
		Fields: graphql.Fields{
			"DestinyGatingScope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorItemSocketOverride = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorItemSocketOverride",
		Fields: graphql.Fields{
			"SingleItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RandomizedOptionsCount": &graphql.Field{
				Type: graphql.Int,
			},
			"SocketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinySocketTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"InsertAction": &graphql.Field{
				Type: DestinyInsertPlugActionDefinition,
			},
			"PlugWhitelist": &graphql.Field{
				Type: graphql.NewList(DestinyPlugWhitelistEntryDefinition),
			},
			"SocketCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Visibility": &graphql.Field{
				Type: graphql.Int,
			},
			"AlwaysRandomizeSockets": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsPreviewEnabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HideDuplicateReusablePlugs": &graphql.Field{
				Type: graphql.Boolean,
			},
			"OverridesUiAppearance": &graphql.Field{
				Type: graphql.Boolean,
			},
			"AvoidDuplicatesOnInitialization": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CurrencyScalars": &graphql.Field{
				Type: graphql.NewList(DestinySocketTypeScalarMaterialRequirementEntry),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyInsertPlugActionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInsertPlugActionDefinition",
		Fields: graphql.Fields{
			"ActionExecuteSeconds": &graphql.Field{
				Type: graphql.Int,
			},
			"ActionType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SocketTypeActionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SocketTypeActionType",
		Fields: graphql.Fields{
			"SocketTypeActionType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPlugWhitelistEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlugWhitelistEntryDefinition",
		Fields: graphql.Fields{
			"CategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"ReinitializationPossiblePlugHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinySocketVisibility = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketVisibility",
		Fields: graphql.Fields{
			"DestinySocketVisibility": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinySocketTypeScalarMaterialRequirementEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketTypeScalarMaterialRequirementEntry",
		Fields: graphql.Fields{
			"CurrencyItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ScalarValue": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinySocketCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketCategoryDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"UiCategoryStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinySocketCategoryStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketCategoryStyle",
		Fields: graphql.Fields{
			"DestinySocketCategoryStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorServiceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorServiceDefinition",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyVendorAcceptedItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorAcceptedItemDefinition",
		Fields: graphql.Fields{
			"AcceptedInventoryBucketHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationInventoryBucketHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorLocationDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorLocationDefinition",
		Fields: graphql.Fields{
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"BackgroundImagePath": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyDestinationDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDestinationDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"PlaceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultFreeroamActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityGraphEntries": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphListEntryDefinition),
			},
			"BubbleSettings": &graphql.Field{
				Type: graphql.NewList(DestinyDestinationBubbleSettingDefinition),
			},
			"Bubbles": &graphql.Field{
				Type: graphql.NewList(DestinyBubbleDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityGraphListEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphListEntryDefinition",
		Fields: graphql.Fields{
			"ActivityGraphHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphDefinition",
		Fields: graphql.Fields{
			"Nodes": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphNodeDefinition),
			},
			"ArtElements": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphArtElementDefinition),
			},
			"Connections": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphConnectionDefinition),
			},
			"DisplayObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphDisplayObjectiveDefinition),
			},
			"DisplayProgressions": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphDisplayProgressionDefinition),
			},
			"LinkedGraphs": &graphql.Field{
				Type: graphql.NewList(DestinyLinkedGraphDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityGraphNodeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphNodeDefinition",
		Fields: graphql.Fields{
			"NodeId": &graphql.Field{
				Type: graphql.Int,
			},
			"OverrideDisplay": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Position": &graphql.Field{
				Type: DestinyPositionDefinition,
			},
			"FeaturingStates": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphNodeFeaturingStateDefinition),
			},
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphNodeActivityDefinition),
			},
			"States": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphNodeStateEntry),
			},
		},
	},
)

var DestinyPositionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPositionDefinition",
		Fields: graphql.Fields{
			"X": &graphql.Field{
				Type: graphql.Int,
			},
			"Y": &graphql.Field{
				Type: graphql.Int,
			},
			"Z": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphNodeFeaturingStateDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphNodeFeaturingStateDefinition",
		Fields: graphql.Fields{
			"HighlightType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ActivityGraphNodeHighlightType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ActivityGraphNodeHighlightType",
		Fields: graphql.Fields{
			"ActivityGraphNodeHighlightType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphNodeActivityDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphNodeActivityDefinition",
		Fields: graphql.Fields{
			"NodeActivityId": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"OriginalDisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"SelectionScreenDisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"ReleaseIcon": &graphql.Field{
				Type: graphql.String,
			},
			"ReleaseTime": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityLightLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PlaceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Tier": &graphql.Field{
				Type: graphql.Int,
			},
			"PgcrImage": &graphql.Field{
				Type: graphql.String,
			},
			"Rewards": &graphql.Field{
				Type: graphql.NewList(DestinyActivityRewardDefinition),
			},
			"Modifiers": &graphql.Field{
				Type: graphql.NewList(DestinyActivityModifierReferenceDefinition),
			},
			"IsPlaylist": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyActivityChallengeDefinition),
			},
			"OptionalUnlockStrings": &graphql.Field{
				Type: graphql.NewList(DestinyActivityUnlockStringDefinition),
			},
			"PlaylistItems": &graphql.Field{
				Type: graphql.NewList(DestinyActivityPlaylistItemDefinition),
			},
			"ActivityGraphList": &graphql.Field{
				Type: graphql.NewList(DestinyActivityGraphListEntryDefinition),
			},
			"Matchmaking": &graphql.Field{
				Type: DestinyActivityMatchmakingBlockDefinition,
			},
			"GuidedGame": &graphql.Field{
				Type: DestinyActivityGuidedBlockDefinition,
			},
			"DirectActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DirectActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
			"Loadouts": &graphql.Field{
				Type: graphql.NewList(DestinyActivityLoadoutRequirementSet),
			},
			"ActivityModeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ActivityModeTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPvP": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InsertionPoints": &graphql.Field{
				Type: graphql.NewList(DestinyActivityInsertionPointDefinition),
			},
			"ActivityLocationMappings": &graphql.Field{
				Type: graphql.NewList(DestinyEnvironmentLocationMapping),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityRewardDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityRewardDefinition",
		Fields: graphql.Fields{
			"RewardText": &graphql.Field{
				Type: graphql.String,
			},
			"RewardItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
		},
	},
)

var DestinyActivityModifierReferenceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModifierReferenceDefinition",
		Fields: graphql.Fields{
			"ActivityModifierHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityModifierDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModifierDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityChallengeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityChallengeDefinition",
		Fields: graphql.Fields{
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DummyRewards": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
		},
	},
)

var DestinyObjectiveDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectiveDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"CompletionValue": &graphql.Field{
				Type: graphql.Int,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"LocationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AllowNegativeValue": &graphql.Field{
				Type: graphql.Boolean,
			},
			"AllowValueChangeWhenCompleted": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsCountingDownward": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ValueStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressDescription": &graphql.Field{
				Type: graphql.String,
			},
			"Perks": &graphql.Field{
				Type: DestinyObjectivePerkEntryDefinition,
			},
			"Stats": &graphql.Field{
				Type: DestinyObjectiveStatEntryDefinition,
			},
			"MinimumVisibilityThreshold": &graphql.Field{
				Type: graphql.Int,
			},
			"AllowOvercompletion": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ShowValueOnComplete": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CompletedValueStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"InProgressValueStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyUnlockValueUIStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyUnlockValueUIStyle",
		Fields: graphql.Fields{
			"DestinyUnlockValueUIStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyObjectivePerkEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectivePerkEntryDefinition",
		Fields: graphql.Fields{
			"PerkHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Style": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyObjectiveGrantStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectiveGrantStyle",
		Fields: graphql.Fields{
			"DestinyObjectiveGrantStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinySandboxPerkDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySandboxPerkDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"PerkIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"IsDisplayable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DamageType": &graphql.Field{
				Type: graphql.Int,
			},
			"DamageTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PerkGroups": &graphql.Field{
				Type: DestinyTalentNodeStepGroups,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DamageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DamageType",
		Fields: graphql.Fields{
			"DamageType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepGroups = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepGroups",
		Fields: graphql.Fields{
			"WeaponPerformance": &graphql.Field{
				Type: graphql.Int,
			},
			"ImpactEffects": &graphql.Field{
				Type: graphql.Int,
			},
			"GuardianAttributes": &graphql.Field{
				Type: graphql.Int,
			},
			"LightAbilities": &graphql.Field{
				Type: graphql.Int,
			},
			"DamageTypes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepWeaponPerformances = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepWeaponPerformances",
		Fields: graphql.Fields{
			"DestinyTalentNodeStepWeaponPerformances": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepImpactEffects = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepImpactEffects",
		Fields: graphql.Fields{
			"DestinyTalentNodeStepImpactEffects": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepGuardianAttributes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepGuardianAttributes",
		Fields: graphql.Fields{
			"DestinyTalentNodeStepGuardianAttributes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepLightAbilities = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepLightAbilities",
		Fields: graphql.Fields{
			"DestinyTalentNodeStepLightAbilities": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStepDamageTypes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStepDamageTypes",
		Fields: graphql.Fields{
			"DestinyTalentNodeStepDamageTypes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyObjectiveStatEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectiveStatEntryDefinition",
		Fields: graphql.Fields{
			"Stat": &graphql.Field{
				Type: DestinyItemInvestmentStatDefinition,
			},
			"Style": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemInvestmentStatDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemInvestmentStatDefinition",
		Fields: graphql.Fields{
			"StatTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
			"IsConditionallyActive": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyLocationDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLocationDefinition",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LocationReleases": &graphql.Field{
				Type: graphql.NewList(DestinyLocationReleaseDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyLocationReleaseDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLocationReleaseDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"SmallTransparentIcon": &graphql.Field{
				Type: graphql.String,
			},
			"MapIcon": &graphql.Field{
				Type: graphql.String,
			},
			"LargeTransparentIcon": &graphql.Field{
				Type: graphql.String,
			},
			"SpawnPoint": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityGraphHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityGraphNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityBubbleName": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityPathBundle": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityPathDestination": &graphql.Field{
				Type: graphql.Int,
			},
			"NavPointType": &graphql.Field{
				Type: graphql.Int,
			},
			"WorldPosition": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyActivityNavPointType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityNavPointType",
		Fields: graphql.Fields{
			"DestinyActivityNavPointType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityUnlockStringDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityUnlockStringDefinition",
		Fields: graphql.Fields{
			"DisplayString": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyActivityPlaylistItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityPlaylistItemDefinition",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DirectActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DirectActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ActivityModeTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyActivityModeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModeType",
		Fields: graphql.Fields{
			"DestinyActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityModeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"PgcrImage": &graphql.Field{
				Type: graphql.String,
			},
			"ModeType": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeCategory": &graphql.Field{
				Type: graphql.Int,
			},
			"IsTeamBased": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsAggregateMode": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ParentHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"FriendlyName": &graphql.Field{
				Type: graphql.String,
			},
			"ActivityModeMappings": &graphql.Field{
				Type: graphql.Int,
			},
			"Display": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityModeCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModeCategory",
		Fields: graphql.Fields{
			"DestinyActivityModeCategory": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityMatchmakingBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityMatchmakingBlockDefinition",
		Fields: graphql.Fields{
			"IsMatchmade": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MinParty": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxParty": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxPlayers": &graphql.Field{
				Type: graphql.Int,
			},
			"RequiresGuardianOath": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityGuidedBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGuidedBlockDefinition",
		Fields: graphql.Fields{
			"GuidedMaxLobbySize": &graphql.Field{
				Type: graphql.Int,
			},
			"GuidedMinLobbySize": &graphql.Field{
				Type: graphql.Int,
			},
			"GuidedDisbandCount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityLoadoutRequirementSet = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityLoadoutRequirementSet",
		Fields: graphql.Fields{
			"Requirements": &graphql.Field{
				Type: graphql.NewList(DestinyActivityLoadoutRequirement),
			},
		},
	},
)

var DestinyActivityLoadoutRequirement = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityLoadoutRequirement",
		Fields: graphql.Fields{
			"EquipmentSlotHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AllowedEquippedItemHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"AllowedWeaponSubTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemSubType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSubType",
		Fields: graphql.Fields{
			"DestinyItemSubType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityInsertionPointDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityInsertionPointDefinition",
		Fields: graphql.Fields{
			"PhaseHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyEnvironmentLocationMapping = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEnvironmentLocationMapping",
		Fields: graphql.Fields{
			"LocationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivationSource": &graphql.Field{
				Type: graphql.String,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPlaceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlaceDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActivityGraphNodeStateEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphNodeStateEntry",
		Fields: graphql.Fields{
			"State": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyGraphNodeState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGraphNodeState",
		Fields: graphql.Fields{
			"DestinyGraphNodeState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphArtElementDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphArtElementDefinition",
		Fields: graphql.Fields{
			"Position": &graphql.Field{
				Type: DestinyPositionDefinition,
			},
		},
	},
)

var DestinyActivityGraphConnectionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphConnectionDefinition",
		Fields: graphql.Fields{
			"SourceNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DestNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphDisplayObjectiveDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphDisplayObjectiveDefinition",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityGraphDisplayProgressionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityGraphDisplayProgressionDefinition",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyLinkedGraphDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLinkedGraphDefinition",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"UnlockExpression": &graphql.Field{
				Type: DestinyUnlockExpressionDefinition,
			},
			"LinkedGraphId": &graphql.Field{
				Type: graphql.Int,
			},
			"LinkedGraphs": &graphql.Field{
				Type: graphql.NewList(DestinyLinkedGraphEntryDefinition),
			},
			"Overview": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyUnlockExpressionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyUnlockExpressionDefinition",
		Fields: graphql.Fields{
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyLinkedGraphEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLinkedGraphEntryDefinition",
		Fields: graphql.Fields{
			"ActivityGraphHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyDestinationBubbleSettingDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDestinationBubbleSettingDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
		},
	},
)

var DestinyBubbleDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBubbleDefinition",
		Fields: graphql.Fields{
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
		},
	},
)

var DestinyVendorGroupReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorGroupReference",
		Fields: graphql.Fields{
			"VendorGroupHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorGroupDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorGroupDefinition",
		Fields: graphql.Fields{
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryName": &graphql.Field{
				Type: graphql.String,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyFactionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyFactionDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"TokenValues": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardVendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Vendors": &graphql.Field{
				Type: graphql.NewList(DestinyFactionVendorDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyFactionVendorDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyFactionVendorDefinition",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"BackgroundImagePath": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyArtifactDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TranslationBlock": &graphql.Field{
				Type: DestinyItemTranslationBlockDefinition,
			},
			"Tiers": &graphql.Field{
				Type: graphql.NewList(DestinyArtifactTierDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyArtifactTierDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactTierDefinition",
		Fields: graphql.Fields{
			"TierHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayTitle": &graphql.Field{
				Type: graphql.String,
			},
			"ProgressRequirementMessage": &graphql.Field{
				Type: graphql.String,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyArtifactTierItemDefinition),
			},
			"MinimumUnlockPointsUsedRequirement": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyArtifactTierItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactTierItemDefinition",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemQualityBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemQualityBlockDefinition",
		Fields: graphql.Fields{
			"ItemLevels": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"QualityLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"InfusionCategoryName": &graphql.Field{
				Type: graphql.String,
			},
			"InfusionCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"InfusionCategoryHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ProgressionLevelRequirementHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentVersion": &graphql.Field{
				Type: graphql.Int,
			},
			"Versions": &graphql.Field{
				Type: graphql.NewList(DestinyItemVersionDefinition),
			},
			"DisplayVersionWatermarkIcons": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var DestinyItemVersionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemVersionDefinition",
		Fields: graphql.Fields{
			"PowerCapHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPowerCapDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPowerCapDefinition",
		Fields: graphql.Fields{
			"PowerCap": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionLevelRequirementDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionLevelRequirementDefinition",
		Fields: graphql.Fields{
			"RequirementCurve": &graphql.Field{
				Type: graphql.NewList(InterpolationPointFloat),
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var InterpolationPointFloat = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InterpolationPointFloat",
		Fields: graphql.Fields{
			"Value": &graphql.Field{
				Type: graphql.Float,
			},
			"Weight": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DestinyItemValueBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemValueBlockDefinition",
		Fields: graphql.Fields{
			"ItemValue": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"ValueDescription": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyItemSourceBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSourceBlockDefinition",
		Fields: graphql.Fields{
			"SourceHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Sources": &graphql.Field{
				Type: graphql.NewList(DestinyItemSourceDefinition),
			},
			"Exclusive": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorSources": &graphql.Field{
				Type: graphql.NewList(DestinyItemVendorSourceReference),
			},
		},
	},
)

var DestinyItemSourceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSourceDefinition",
		Fields: graphql.Fields{
			"Level": &graphql.Field{
				Type: graphql.Int,
			},
			"MinQuality": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxQuality": &graphql.Field{
				Type: graphql.Int,
			},
			"MinLevelRequired": &graphql.Field{
				Type: graphql.Int,
			},
			"MaxLevelRequired": &graphql.Field{
				Type: graphql.Int,
			},
			"ComputedStats": &graphql.Field{
				Type: graphql.Int,
			},
			"SourceHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyRewardSourceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRewardSourceDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Category": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyRewardSourceCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRewardSourceCategory",
		Fields: graphql.Fields{
			"DestinyRewardSourceCategory": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemVendorSourceReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemVendorSourceReference",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorItemIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemObjectiveBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemObjectiveBlockDefinition",
		Fields: graphql.Fields{
			"ObjectiveHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"DisplayActivityHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"RequireFullObjectiveCompletion": &graphql.Field{
				Type: graphql.Boolean,
			},
			"QuestlineItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Narrative": &graphql.Field{
				Type: graphql.String,
			},
			"ObjectiveVerbName": &graphql.Field{
				Type: graphql.String,
			},
			"QuestTypeIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"QuestTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PerObjectiveDisplayProperties": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveDisplayProperties),
			},
			"DisplayAsStatTracker": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyObjectiveDisplayProperties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectiveDisplayProperties",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayOnItemPreviewScreen": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemMetricBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemMetricBlockDefinition",
		Fields: graphql.Fields{
			"AvailableMetricCategoryNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyPresentationNodeBaseDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeBaseDefinition",
		Fields: graphql.Fields{
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPresentationNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeType",
		Fields: graphql.Fields{
			"DestinyPresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTraitDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTraitDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TraitCategoryId": &graphql.Field{
				Type: graphql.String,
			},
			"TraitCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyTraitCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTraitCategoryDefinition",
		Fields: graphql.Fields{
			"TraitCategoryId": &graphql.Field{
				Type: graphql.String,
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyScoredPresentationNodeBaseDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyScoredPresentationNodeBaseDefinition",
		Fields: graphql.Fields{
			"MaxCategoryRecordScore": &graphql.Field{
				Type: graphql.Int,
			},
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPresentationNodeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"OriginalIcon": &graphql.Field{
				Type: graphql.String,
			},
			"RootViewIcon": &graphql.Field{
				Type: graphql.String,
			},
			"NodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CompletionRecordHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Children": &graphql.Field{
				Type: DestinyPresentationNodeChildrenBlock,
			},
			"DisplayStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"ScreenStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"Requirements": &graphql.Field{
				Type: DestinyPresentationNodeRequirementsBlock,
			},
			"DisableChildSubscreenNavigation": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MaxCategoryRecordScore": &graphql.Field{
				Type: graphql.Int,
			},
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyScope = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyScope",
		Fields: graphql.Fields{
			"DestinyScope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationNodeChildrenBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeChildrenBlock",
		Fields: graphql.Fields{
			"PresentationNodes": &graphql.Field{
				Type: graphql.NewList(DestinyPresentationNodeChildEntry),
			},
			"Collectibles": &graphql.Field{
				Type: graphql.NewList(DestinyPresentationNodeCollectibleChildEntry),
			},
			"Records": &graphql.Field{
				Type: graphql.NewList(DestinyPresentationNodeRecordChildEntry),
			},
			"Metrics": &graphql.Field{
				Type: graphql.NewList(DestinyPresentationNodeMetricChildEntry),
			},
		},
	},
)

var DestinyPresentationNodeChildEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeChildEntry",
		Fields: graphql.Fields{
			"PresentationNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationNodeCollectibleChildEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeCollectibleChildEntry",
		Fields: graphql.Fields{
			"CollectibleHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyCollectibleDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"SourceString": &graphql.Field{
				Type: graphql.String,
			},
			"SourceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AcquisitionInfo": &graphql.Field{
				Type: DestinyCollectibleAcquisitionBlock,
			},
			"StateInfo": &graphql.Field{
				Type: DestinyCollectibleStateBlock,
			},
			"PresentationInfo": &graphql.Field{
				Type: DestinyPresentationChildBlock,
			},
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCollectibleAcquisitionBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleAcquisitionBlock",
		Fields: graphql.Fields{
			"AcquireMaterialRequirementHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AcquireTimestampUnlockValueHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMaterialRequirementSetDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMaterialRequirementSetDefinition",
		Fields: graphql.Fields{
			"Materials": &graphql.Field{
				Type: graphql.NewList(DestinyMaterialRequirement),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyMaterialRequirement = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMaterialRequirement",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DeleteOnAction": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Count": &graphql.Field{
				Type: graphql.Int,
			},
			"OmitFromRequirements": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyUnlockValueDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyUnlockValueDefinition",
		Fields: graphql.Fields{
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCollectibleStateBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleStateBlock",
		Fields: graphql.Fields{
			"ObscuredOverrideItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Requirements": &graphql.Field{
				Type: DestinyPresentationNodeRequirementsBlock,
			},
		},
	},
)

var DestinyPresentationNodeRequirementsBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeRequirementsBlock",
		Fields: graphql.Fields{
			"EntitlementUnavailableMessage": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyPresentationChildBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationChildBlock",
		Fields: graphql.Fields{
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"ParentPresentationNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"DisplayStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationDisplayStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationDisplayStyle",
		Fields: graphql.Fields{
			"DestinyPresentationDisplayStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationNodeRecordChildEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeRecordChildEntry",
		Fields: graphql.Fields{
			"RecordHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"PresentationInfo": &graphql.Field{
				Type: DestinyPresentationChildBlock,
			},
			"LoreHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ObjectiveHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"RecordValueStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"ForTitleGilding": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TitleInfo": &graphql.Field{
				Type: DestinyRecordTitleBlock,
			},
			"CompletionInfo": &graphql.Field{
				Type: DestinyRecordCompletionBlock,
			},
			"StateInfo": &graphql.Field{
				Type: SchemaRecordStateBlock,
			},
			"Requirements": &graphql.Field{
				Type: DestinyPresentationNodeRequirementsBlock,
			},
			"ExpirationInfo": &graphql.Field{
				Type: DestinyRecordExpirationBlock,
			},
			"IntervalInfo": &graphql.Field{
				Type: DestinyRecordIntervalBlock,
			},
			"RewardItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyRecordValueStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordValueStyle",
		Fields: graphql.Fields{
			"DestinyRecordValueStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordTitleBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordTitleBlock",
		Fields: graphql.Fields{
			"HasTitle": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TitlesByGender": &graphql.Field{
				Type: graphql.Int,
			},
			"TitlesByGenderHash": &graphql.Field{
				Type: graphql.Int,
			},
			"GildingTrackingRecordHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyGender = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGender",
		Fields: graphql.Fields{
			"DestinyGender": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyGenderDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGenderDefinition",
		Fields: graphql.Fields{
			"GenderType": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyRecordCompletionBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordCompletionBlock",
		Fields: graphql.Fields{
			"PartialCompletionObjectiveCountThreshold": &graphql.Field{
				Type: graphql.Int,
			},
			"ScoreValue": &graphql.Field{
				Type: graphql.Int,
			},
			"ShouldFireToast": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ToastStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordToastStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordToastStyle",
		Fields: graphql.Fields{
			"DestinyRecordToastStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SchemaRecordStateBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SchemaRecordStateBlock",
		Fields: graphql.Fields{
			"FeaturedPriority": &graphql.Field{
				Type: graphql.Int,
			},
			"ObscuredString": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyRecordExpirationBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordExpirationBlock",
		Fields: graphql.Fields{
			"HasExpiration": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyRecordIntervalBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordIntervalBlock",
		Fields: graphql.Fields{
			"IntervalObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyRecordIntervalObjective),
			},
			"IntervalRewards": &graphql.Field{
				Type: graphql.NewList(DestinyRecordIntervalRewards),
			},
			"OriginalObjectiveArrayInsertionIndex": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordIntervalObjective = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordIntervalObjective",
		Fields: graphql.Fields{
			"IntervalObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IntervalScoreValue": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordIntervalRewards = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordIntervalRewards",
		Fields: graphql.Fields{
			"IntervalRewardItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
		},
	},
)

var DestinyLoreDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLoreDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Subtitle": &graphql.Field{
				Type: graphql.String,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPresentationNodeMetricChildEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeMetricChildEntry",
		Fields: graphql.Fields{
			"MetricHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMetricDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMetricDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TrackingObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LowerValueIsBetter": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PresentationNodeType": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitIds": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"TraitHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPresentationScreenStyle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationScreenStyle",
		Fields: graphql.Fields{
			"DestinyPresentationScreenStyle": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemPlugDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPlugDefinition",
		Fields: graphql.Fields{
			"InsertionRules": &graphql.Field{
				Type: graphql.NewList(DestinyPlugRuleDefinition),
			},
			"PlugCategoryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"PlugCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"OnActionRecreateSelf": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InsertionMaterialRequirementHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PreviewItemOverrideHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EnabledMaterialRequirementHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EnabledRules": &graphql.Field{
				Type: graphql.NewList(DestinyPlugRuleDefinition),
			},
			"UiPlugLabel": &graphql.Field{
				Type: graphql.String,
			},
			"PlugStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"PlugAvailability": &graphql.Field{
				Type: graphql.Int,
			},
			"AlternateUiPlugLabel": &graphql.Field{
				Type: graphql.String,
			},
			"AlternatePlugStyle": &graphql.Field{
				Type: graphql.Int,
			},
			"IsDummyPlug": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ParentItemOverride": &graphql.Field{
				Type: DestinyParentItemOverride,
			},
			"EnergyCapacity": &graphql.Field{
				Type: DestinyEnergyCapacityEntry,
			},
			"EnergyCost": &graphql.Field{
				Type: DestinyEnergyCostEntry,
			},
		},
	},
)

var DestinyPlugRuleDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlugRuleDefinition",
		Fields: graphql.Fields{
			"FailureMessage": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PlugUiStyles = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlugUiStyles",
		Fields: graphql.Fields{
			"PlugUiStyles": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PlugAvailabilityMode = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlugAvailabilityMode",
		Fields: graphql.Fields{
			"PlugAvailabilityMode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyParentItemOverride = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyParentItemOverride",
		Fields: graphql.Fields{
			"AdditionalEquipRequirementsDisplayStrings": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"PipIcon": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyEnergyCapacityEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEnergyCapacityEntry",
		Fields: graphql.Fields{
			"CapacityValue": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyEnergyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEnergyType",
		Fields: graphql.Fields{
			"DestinyEnergyType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyEnergyTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEnergyTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TransparentIconPath": &graphql.Field{
				Type: graphql.String,
			},
			"ShowIcon": &graphql.Field{
				Type: graphql.Boolean,
			},
			"EnumValue": &graphql.Field{
				Type: graphql.Int,
			},
			"CapacityStatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CostStatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyEnergyCostEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEnergyCostEntry",
		Fields: graphql.Fields{
			"EnergyCost": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemGearsetBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemGearsetBlockDefinition",
		Fields: graphql.Fields{
			"TrackingValueMax": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemList": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemSackBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSackBlockDefinition",
		Fields: graphql.Fields{
			"DetailAction": &graphql.Field{
				Type: graphql.String,
			},
			"OpenAction": &graphql.Field{
				Type: graphql.String,
			},
			"SelectItemCount": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorSackType": &graphql.Field{
				Type: graphql.String,
			},
			"OpenOnAcquire": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemSocketBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketBlockDefinition",
		Fields: graphql.Fields{
			"Detail": &graphql.Field{
				Type: graphql.String,
			},
			"SocketEntries": &graphql.Field{
				Type: graphql.NewList(DestinyItemSocketEntryDefinition),
			},
			"IntrinsicSockets": &graphql.Field{
				Type: graphql.NewList(DestinyItemIntrinsicSocketEntryDefinition),
			},
			"SocketCategories": &graphql.Field{
				Type: graphql.NewList(DestinyItemSocketCategoryDefinition),
			},
		},
	},
)

var DestinyItemSocketEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketEntryDefinition",
		Fields: graphql.Fields{
			"SocketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SingleInitialItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ReusablePlugItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemSocketEntryPlugItemDefinition),
			},
			"PreventInitializationOnVendorPurchase": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HidePerksInItemTooltip": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PlugSources": &graphql.Field{
				Type: graphql.Int,
			},
			"ReusablePlugSetHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RandomizedPlugSetHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultVisible": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemSocketEntryPlugItemDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketEntryPlugItemDefinition",
		Fields: graphql.Fields{
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SocketPlugSources = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SocketPlugSources",
		Fields: graphql.Fields{
			"SocketPlugSources": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPlugSetDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlugSetDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"ReusablePlugItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemSocketEntryPlugItemRandomizedDefinition),
			},
			"IsFakePlugSet": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemSocketEntryPlugItemRandomizedDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketEntryPlugItemRandomizedDefinition",
		Fields: graphql.Fields{
			"CurrentlyCanRoll": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemIntrinsicSocketEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemIntrinsicSocketEntryDefinition",
		Fields: graphql.Fields{
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SocketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DefaultVisible": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemSocketCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketCategoryDefinition",
		Fields: graphql.Fields{
			"SocketCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SocketIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemSummaryBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSummaryBlockDefinition",
		Fields: graphql.Fields{
			"SortPriority": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemTalentGridBlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTalentGridBlockDefinition",
		Fields: graphql.Fields{
			"TalentGridHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemDetailString": &graphql.Field{
				Type: graphql.String,
			},
			"BuildName": &graphql.Field{
				Type: graphql.String,
			},
			"HudDamageType": &graphql.Field{
				Type: graphql.Int,
			},
			"HudIcon": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyTalentGridDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentGridDefinition",
		Fields: graphql.Fields{
			"MaxGridLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"GridLevelPerColumn": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Nodes": &graphql.Field{
				Type: graphql.NewList(DestinyTalentNodeDefinition),
			},
			"ExclusiveSets": &graphql.Field{
				Type: graphql.NewList(DestinyTalentNodeExclusiveSetDefinition),
			},
			"IndependentNodeIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Groups": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeCategories": &graphql.Field{
				Type: graphql.NewList(DestinyTalentNodeCategory),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyTalentNodeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeDefinition",
		Fields: graphql.Fields{
			"NodeIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Row": &graphql.Field{
				Type: graphql.Int,
			},
			"Column": &graphql.Field{
				Type: graphql.Int,
			},
			"PrerequisiteNodeIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"BinaryPairNodeIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"AutoUnlocks": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LastStepRepeats": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsRandom": &graphql.Field{
				Type: graphql.Boolean,
			},
			"RandomActivationRequirement": &graphql.Field{
				Type: DestinyNodeActivationRequirement,
			},
			"IsRandomRepurchasable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Steps": &graphql.Field{
				Type: graphql.NewList(DestinyNodeStepDefinition),
			},
			"ExclusiveWithNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"RandomStartProgressionBarAtProgression": &graphql.Field{
				Type: graphql.Int,
			},
			"LayoutIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"GroupHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LoreHash": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeStyleIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"IgnoreForCompletion": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyNodeActivationRequirement = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyNodeActivationRequirement",
		Fields: graphql.Fields{
			"GridLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"MaterialRequirementHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyNodeStepDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyNodeStepDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"StepIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeStepHash": &graphql.Field{
				Type: graphql.Int,
			},
			"InteractionDescription": &graphql.Field{
				Type: graphql.String,
			},
			"DamageType": &graphql.Field{
				Type: graphql.Int,
			},
			"DamageTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivationRequirement": &graphql.Field{
				Type: DestinyNodeActivationRequirement,
			},
			"CanActivateNextStep": &graphql.Field{
				Type: graphql.Boolean,
			},
			"NextStepIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"IsNextStepRandom": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PerkHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"StartProgressionBarAtProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"StatHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"AffectsQuality": &graphql.Field{
				Type: graphql.Boolean,
			},
			"StepGroups": &graphql.Field{
				Type: DestinyTalentNodeStepGroups,
			},
			"AffectsLevel": &graphql.Field{
				Type: graphql.Boolean,
			},
			"SocketReplacements": &graphql.Field{
				Type: graphql.NewList(DestinyNodeSocketReplaceResponse),
			},
		},
	},
)

var DestinyNodeSocketReplaceResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyNodeSocketReplaceResponse",
		Fields: graphql.Fields{
			"SocketTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyDamageTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyDamageTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"TransparentIconPath": &graphql.Field{
				Type: graphql.String,
			},
			"ShowIcon": &graphql.Field{
				Type: graphql.Boolean,
			},
			"EnumValue": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyTalentNodeExclusiveSetDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeExclusiveSetDefinition",
		Fields: graphql.Fields{
			"NodeIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyTalentExclusiveGroup = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentExclusiveGroup",
		Fields: graphql.Fields{
			"GroupHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LoreHash": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"OpposingGroupHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"OpposingNodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyTalentNodeCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeCategory",
		Fields: graphql.Fields{
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"IsLoreDriven": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"NodeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemPerkEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPerkEntryDefinition",
		Fields: graphql.Fields{
			"RequirementDisplayString": &graphql.Field{
				Type: graphql.String,
			},
			"PerkHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PerkVisibility": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ItemPerkVisibility = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemPerkVisibility",
		Fields: graphql.Fields{
			"ItemPerkVisibility": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyAnimationReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyAnimationReference",
		Fields: graphql.Fields{
			"AnimName": &graphql.Field{
				Type: graphql.String,
			},
			"AnimIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"Path": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var HyperlinkReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "HyperlinkReference",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SpecialItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SpecialItemType",
		Fields: graphql.Fields{
			"SpecialItemType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemType",
		Fields: graphql.Fields{
			"DestinyItemType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyClass = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyClass",
		Fields: graphql.Fields{
			"DestinyClass": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyBreakerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBreakerType",
		Fields: graphql.Fields{
			"DestinyBreakerType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemCategoryDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Deprecated": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ShortTitle": &graphql.Field{
				Type: graphql.String,
			},
			"ItemTypeRegex": &graphql.Field{
				Type: graphql.String,
			},
			"GrantDestinyBreakerType": &graphql.Field{
				Type: graphql.Int,
			},
			"PlugCategoryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"ItemTypeRegexNot": &graphql.Field{
				Type: graphql.String,
			},
			"OriginBucketIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"GrantDestinyItemType": &graphql.Field{
				Type: graphql.Int,
			},
			"GrantDestinySubType": &graphql.Field{
				Type: graphql.Int,
			},
			"GrantDestinyClass": &graphql.Field{
				Type: graphql.Int,
			},
			"TraitId": &graphql.Field{
				Type: graphql.String,
			},
			"GroupedCategoryHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ParentCategoryHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"GroupCategoryOnly": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyBreakerTypeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBreakerTypeDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"EnumValue": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinySeasonDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySeasonDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"BackgroundImagePath": &graphql.Field{
				Type: graphql.String,
			},
			"SeasonNumber": &graphql.Field{
				Type: graphql.Int,
			},
			"StartDate": &graphql.Field{
				Type: graphql.String,
			},
			"EndDate": &graphql.Field{
				Type: graphql.String,
			},
			"SeasonPassHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonPassProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ArtifactItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SealPresentationNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonalChallengesPresentationNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Preview": &graphql.Field{
				Type: DestinySeasonPreviewDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinySeasonPreviewDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySeasonPreviewDefinition",
		Fields: graphql.Fields{
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"LinkPath": &graphql.Field{
				Type: graphql.String,
			},
			"VideoLink": &graphql.Field{
				Type: graphql.String,
			},
			"Images": &graphql.Field{
				Type: graphql.NewList(DestinySeasonPreviewImageDefinition),
			},
		},
	},
)

var DestinySeasonPreviewImageDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySeasonPreviewImageDefinition",
		Fields: graphql.Fields{
			"ThumbnailImage": &graphql.Field{
				Type: graphql.String,
			},
			"HighResImage": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinySeasonPassDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySeasonPassDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"RewardProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PrestigeProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionRewardItemQuantity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionRewardItemQuantity",
		Fields: graphql.Fields{
			"RewardedAtProgressionLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"AcquisitionBehavior": &graphql.Field{
				Type: graphql.Int,
			},
			"UiDisplayStyle": &graphql.Field{
				Type: graphql.String,
			},
			"ClaimUnlockDisplayStrings": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"HasConditionalVisibility": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProgressionRewardItemAcquisitionBehavior = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProgressionRewardItemAcquisitionBehavior",
		Fields: graphql.Fields{
			"DestinyProgressionRewardItemAcquisitionBehavior": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupUserBase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupUserBase",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyUserInfo": &graphql.Field{
				Type: GroupUserInfoCard,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"JoinDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupMember = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMember",
		Fields: graphql.Fields{
			"MemberType": &graphql.Field{
				Type: graphql.Int,
			},
			"IsOnline": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LastOnlineStatusChange": &graphql.Field{
				Type: graphql.Int,
			},
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyUserInfo": &graphql.Field{
				Type: GroupUserInfoCard,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"JoinDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupAllianceStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupAllianceStatus",
		Fields: graphql.Fields{
			"GroupAllianceStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupPotentialMember = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupPotentialMember",
		Fields: graphql.Fields{
			"PotentialStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyUserInfo": &graphql.Field{
				Type: GroupUserInfoCard,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"JoinDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupPotentialMemberStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupPotentialMemberStatus",
		Fields: graphql.Fields{
			"GroupPotentialMemberStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var TagResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TagResponse",
		Fields: graphql.Fields{
			"TagText": &graphql.Field{
				Type: graphql.String,
			},
			"IgnoreStatus": &graphql.Field{
				Type: IgnoreResponse,
			},
		},
	},
)

var PollResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PollResponse",
		Fields: graphql.Fields{
			"TopicId": &graphql.Field{
				Type: graphql.Int,
			},
			"Results": &graphql.Field{
				Type: graphql.NewList(PollResult),
			},
			"TotalVotes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PollResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PollResult",
		Fields: graphql.Fields{
			"AnswerText": &graphql.Field{
				Type: graphql.String,
			},
			"AnswerSlot": &graphql.Field{
				Type: graphql.Int,
			},
			"LastVoteDate": &graphql.Field{
				Type: graphql.String,
			},
			"Votes": &graphql.Field{
				Type: graphql.Int,
			},
			"RequestingUserVoted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ForumRecruitmentDetail = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumRecruitmentDetail",
		Fields: graphql.Fields{
			"TopicId": &graphql.Field{
				Type: graphql.Int,
			},
			"MicrophoneRequired": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Intensity": &graphql.Field{
				Type: graphql.Int,
			},
			"Tone": &graphql.Field{
				Type: graphql.Int,
			},
			"Approved": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ConversationId": &graphql.Field{
				Type: graphql.Int,
			},
			"PlayerSlotsTotal": &graphql.Field{
				Type: graphql.Int,
			},
			"PlayerSlotsRemaining": &graphql.Field{
				Type: graphql.Int,
			},
			"Fireteam": &graphql.Field{
				Type: graphql.NewList(GeneralUser),
			},
			"KickedPlayerIds": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var ForumRecruitmentIntensityLabel = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumRecruitmentIntensityLabel",
		Fields: graphql.Fields{
			"ForumRecruitmentIntensityLabel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumRecruitmentToneLabel = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumRecruitmentToneLabel",
		Fields: graphql.Fields{
			"ForumRecruitmentToneLabel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ForumPostSortEnum = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ForumPostSortEnum",
		Fields: graphql.Fields{
			"ForumPostSortEnum": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupTheme = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupTheme",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Folder": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupDateRange = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupDateRange",
		Fields: graphql.Fields{
			"GroupDateRange": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupV2Card = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupV2Card",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"GroupType": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
			"About": &graphql.Field{
				Type: graphql.String,
			},
			"Motto": &graphql.Field{
				Type: graphql.String,
			},
			"MemberCount": &graphql.Field{
				Type: graphql.Int,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
			"MembershipOption": &graphql.Field{
				Type: graphql.Int,
			},
			"Capabilities": &graphql.Field{
				Type: graphql.Int,
			},
			"ClanInfo": &graphql.Field{
				Type: GroupV2ClanInfo,
			},
			"AvatarPath": &graphql.Field{
				Type: graphql.String,
			},
			"Theme": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SearchResultOfGroupV2Card = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupV2Card",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupV2Card),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupSearchResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupSearchResponse",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupV2Card),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupQuery",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"GroupType": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.Int,
			},
			"SortBy": &graphql.Field{
				Type: graphql.Int,
			},
			"GroupMemberCountFilter": &graphql.Field{
				Type: graphql.Int,
			},
			"LocaleFilter": &graphql.Field{
				Type: graphql.String,
			},
			"TagText": &graphql.Field{
				Type: graphql.String,
			},
			"ItemsPerPage": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentPage": &graphql.Field{
				Type: graphql.Int,
			},
			"RequestContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupSortBy = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupSortBy",
		Fields: graphql.Fields{
			"GroupSortBy": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupMemberCountFilter = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMemberCountFilter",
		Fields: graphql.Fields{
			"GroupMemberCountFilter": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupNameSearchRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupNameSearchRequest",
		Fields: graphql.Fields{
			"GroupName": &graphql.Field{
				Type: graphql.String,
			},
			"GroupType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupOptionalConversation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupOptionalConversation",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"ConversationId": &graphql.Field{
				Type: graphql.Int,
			},
			"ChatEnabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ChatName": &graphql.Field{
				Type: graphql.String,
			},
			"ChatSecurity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupEditAction = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupEditAction",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"About": &graphql.Field{
				Type: graphql.String,
			},
			"Motto": &graphql.Field{
				Type: graphql.String,
			},
			"Theme": &graphql.Field{
				Type: graphql.String,
			},
			"AvatarImageIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"Tags": &graphql.Field{
				Type: graphql.String,
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipOption": &graphql.Field{
				Type: graphql.Int,
			},
			"IsPublicTopicAdminOnly": &graphql.Field{
				Type: graphql.Boolean,
			},
			"AllowChat": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ChatSecurity": &graphql.Field{
				Type: graphql.Int,
			},
			"Callsign": &graphql.Field{
				Type: graphql.String,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
			"Homepage": &graphql.Field{
				Type: graphql.Int,
			},
			"EnableInvitationMessagingForAdmins": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DefaultPublicity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupOptionsEditAction = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupOptionsEditAction",
		Fields: graphql.Fields{
			"InvitePermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"UpdateCulturePermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HostGuidedGamePermissionOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"UpdateBannerPermissionOverride": &graphql.Field{
				Type: graphql.Boolean,
			},
			"JoinLevel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupOptionalConversationAddRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupOptionalConversationAddRequest",
		Fields: graphql.Fields{
			"ChatName": &graphql.Field{
				Type: graphql.String,
			},
			"ChatSecurity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupOptionalConversationEditRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupOptionalConversationEditRequest",
		Fields: graphql.Fields{
			"ChatEnabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ChatName": &graphql.Field{
				Type: graphql.String,
			},
			"ChatSecurity": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResultOfGroupMember = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupMember",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupMember),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupMemberLeaveResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMemberLeaveResult",
		Fields: graphql.Fields{
			"Group": &graphql.Field{
				Type: GroupV2,
			},
			"GroupDeleted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupBanRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupBanRequest",
		Fields: graphql.Fields{
			"Comment": &graphql.Field{
				Type: graphql.String,
			},
			"Length": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var IgnoreLength = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "IgnoreLength",
		Fields: graphql.Fields{
			"IgnoreLength": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupBan = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupBan",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"LastModifiedBy": &graphql.Field{
				Type: UserInfoCard,
			},
			"CreatedBy": &graphql.Field{
				Type: UserInfoCard,
			},
			"DateBanned": &graphql.Field{
				Type: graphql.String,
			},
			"DateExpires": &graphql.Field{
				Type: graphql.String,
			},
			"Comment": &graphql.Field{
				Type: graphql.String,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"DestinyUserInfo": &graphql.Field{
				Type: GroupUserInfoCard,
			},
		},
	},
)

var SearchResultOfGroupBan = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupBan",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupBan),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupMemberApplication = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMemberApplication",
		Fields: graphql.Fields{
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
			"ResolveState": &graphql.Field{
				Type: graphql.Int,
			},
			"ResolveDate": &graphql.Field{
				Type: graphql.String,
			},
			"ResolvedByMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"RequestMessage": &graphql.Field{
				Type: graphql.String,
			},
			"ResolveMessage": &graphql.Field{
				Type: graphql.String,
			},
			"DestinyUserInfo": &graphql.Field{
				Type: GroupUserInfoCard,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
		},
	},
)

var GroupApplicationResolveState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupApplicationResolveState",
		Fields: graphql.Fields{
			"GroupApplicationResolveState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResultOfGroupMemberApplication = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupMemberApplication",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupMemberApplication),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var EntityActionResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EntityActionResult",
		Fields: graphql.Fields{
			"EntityId": &graphql.Field{
				Type: graphql.Int,
			},
			"Result": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PlatformErrorCodes = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlatformErrorCodes",
		Fields: graphql.Fields{
			"PlatformErrorCodes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupApplicationRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupApplicationRequest",
		Fields: graphql.Fields{
			"Message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupApplicationListRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupApplicationListRequest",
		Fields: graphql.Fields{
			"Memberships": &graphql.Field{
				Type: graphql.NewList(UserMembership),
			},
			"Message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var GroupsForMemberFilter = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupsForMemberFilter",
		Fields: graphql.Fields{
			"GroupsForMemberFilter": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GroupMembershipBase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMembershipBase",
		Fields: graphql.Fields{
			"Group": &graphql.Field{
				Type: GroupV2,
			},
		},
	},
)

var GroupMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMembership",
		Fields: graphql.Fields{
			"Member": &graphql.Field{
				Type: GroupMember,
			},
			"Group": &graphql.Field{
				Type: GroupV2,
			},
		},
	},
)

var SearchResultOfGroupMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupMembership",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupMembership),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupMembershipSearchResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupMembershipSearchResponse",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupMembership),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GetGroupsForMemberResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GetGroupsForMemberResponse",
		Fields: graphql.Fields{
			"AreAllMembershipsInactive": &graphql.Field{
				Type: graphql.Int,
			},
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupMembership),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupPotentialMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupPotentialMembership",
		Fields: graphql.Fields{
			"Member": &graphql.Field{
				Type: GroupPotentialMember,
			},
			"Group": &graphql.Field{
				Type: GroupV2,
			},
		},
	},
)

var SearchResultOfGroupPotentialMembership = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfGroupPotentialMembership",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupPotentialMembership),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupPotentialMembershipSearchResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupPotentialMembershipSearchResponse",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(GroupPotentialMembership),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var GroupApplicationResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GroupApplicationResponse",
		Fields: graphql.Fields{
			"Resolution": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PartnerOfferClaimRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PartnerOfferClaimRequest",
		Fields: graphql.Fields{
			"PartnerOfferId": &graphql.Field{
				Type: graphql.String,
			},
			"BungieNetMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"TransactionId": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PartnerOfferSkuHistoryResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PartnerOfferSkuHistoryResponse",
		Fields: graphql.Fields{
			"SkuIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"LocalizedName": &graphql.Field{
				Type: graphql.String,
			},
			"LocalizedDescription": &graphql.Field{
				Type: graphql.String,
			},
			"ClaimDate": &graphql.Field{
				Type: graphql.String,
			},
			"AllOffersApplied": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TransactionId": &graphql.Field{
				Type: graphql.String,
			},
			"SkuOffers": &graphql.Field{
				Type: graphql.NewList(PartnerOfferHistoryResponse),
			},
		},
	},
)

var PartnerOfferHistoryResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PartnerOfferHistoryResponse",
		Fields: graphql.Fields{
			"PartnerOfferKey": &graphql.Field{
				Type: graphql.String,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"LocalizedName": &graphql.Field{
				Type: graphql.String,
			},
			"LocalizedDescription": &graphql.Field{
				Type: graphql.String,
			},
			"IsConsumable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"QuantityApplied": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplyDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyManifest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyManifest",
		Fields: graphql.Fields{
			"Version": &graphql.Field{
				Type: graphql.String,
			},
			"MobileAssetContentPath": &graphql.Field{
				Type: graphql.String,
			},
			"MobileGearAssetDataBases": &graphql.Field{
				Type: graphql.NewList(GearAssetDataBaseDefinition),
			},
			"MobileWorldContentPaths": &graphql.Field{
				Type: graphql.String,
			},
			"JsonWorldContentPaths": &graphql.Field{
				Type: graphql.String,
			},
			"JsonWorldComponentContentPaths": &graphql.Field{
				Type: graphql.String,
			},
			"MobileClanBannerDatabasePath": &graphql.Field{
				Type: graphql.String,
			},
			"MobileGearCDN": &graphql.Field{
				Type: graphql.String,
			},
			"IconImagePyramidInfo": &graphql.Field{
				Type: graphql.NewList(ImagePyramidEntry),
			},
		},
	},
)

var GearAssetDataBaseDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GearAssetDataBaseDefinition",
		Fields: graphql.Fields{
			"Version": &graphql.Field{
				Type: graphql.Int,
			},
			"Path": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ImagePyramidEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ImagePyramidEntry",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Factor": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var DestinyLinkedProfilesResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLinkedProfilesResponse",
		Fields: graphql.Fields{
			"Profiles": &graphql.Field{
				Type: graphql.NewList(DestinyProfileUserInfoCard),
			},
			"BnetMembership": &graphql.Field{
				Type: UserInfoCard,
			},
			"ProfilesWithErrors": &graphql.Field{
				Type: graphql.NewList(DestinyErrorProfile),
			},
		},
	},
)

var DestinyProfileUserInfoCard = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileUserInfoCard",
		Fields: graphql.Fields{
			"DateLastPlayed": &graphql.Field{
				Type: graphql.String,
			},
			"IsOverridden": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsCrossSavePrimary": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PlatformSilver": &graphql.Field{
				Type: DestinyPlatformSilverComponent,
			},
			"UnpairedGameVersions": &graphql.Field{
				Type: graphql.Int,
			},
			"SupplementalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPlatformSilverComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlatformSilverComponent",
		Fields: graphql.Fields{
			"PlatformSilver": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemComponent",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"BindStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"Location": &graphql.Field{
				Type: graphql.Int,
			},
			"BucketHash": &graphql.Field{
				Type: graphql.Int,
			},
			"TransferStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"Lockable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"State": &graphql.Field{
				Type: graphql.Int,
			},
			"OverrideStyleItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ExpirationDate": &graphql.Field{
				Type: graphql.String,
			},
			"IsWrapper": &graphql.Field{
				Type: graphql.Boolean,
			},
			"TooltipNotificationIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"MetricHash": &graphql.Field{
				Type: graphql.Int,
			},
			"MetricObjective": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
			"VersionNumber": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemValueVisibility": &graphql.Field{
				Type: graphql.NewList(graphql.Boolean),
			},
		},
	},
)

var ItemBindStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemBindStatus",
		Fields: graphql.Fields{
			"ItemBindStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var TransferStatuses = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TransferStatuses",
		Fields: graphql.Fields{
			"TransferStatuses": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var ItemState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ItemState",
		Fields: graphql.Fields{
			"ItemState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyObjectiveProgress = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyObjectiveProgress",
		Fields: graphql.Fields{
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Progress": &graphql.Field{
				Type: graphql.Int,
			},
			"CompletionValue": &graphql.Field{
				Type: graphql.Int,
			},
			"Complete": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyGameVersions = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGameVersions",
		Fields: graphql.Fields{
			"DestinyGameVersions": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyErrorProfile = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyErrorProfile",
		Fields: graphql.Fields{
			"ErrorCode": &graphql.Field{
				Type: graphql.Int,
			},
			"InfoCard": &graphql.Field{
				Type: UserInfoCard,
			},
		},
	},
)

var DestinyComponentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyComponentType",
		Fields: graphql.Fields{
			"DestinyComponentType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileResponse",
		Fields: graphql.Fields{
			"VendorReceipts": &graphql.Field{
				Type: SingleComponentResponseOfDestinyVendorReceiptsComponent,
			},
			"ProfileInventory": &graphql.Field{
				Type: SingleComponentResponseOfDestinyInventoryComponent,
			},
			"ProfileCurrencies": &graphql.Field{
				Type: SingleComponentResponseOfDestinyInventoryComponent,
			},
			"Profile": &graphql.Field{
				Type: SingleComponentResponseOfDestinyProfileComponent,
			},
			"PlatformSilver": &graphql.Field{
				Type: SingleComponentResponseOfDestinyPlatformSilverComponent,
			},
			"ProfileKiosks": &graphql.Field{
				Type: SingleComponentResponseOfDestinyKiosksComponent,
			},
			"ProfilePlugSets": &graphql.Field{
				Type: SingleComponentResponseOfDestinyPlugSetsComponent,
			},
			"ProfileProgression": &graphql.Field{
				Type: SingleComponentResponseOfDestinyProfileProgressionComponent,
			},
			"ProfilePresentationNodes": &graphql.Field{
				Type: SingleComponentResponseOfDestinyPresentationNodesComponent,
			},
			"ProfileRecords": &graphql.Field{
				Type: SingleComponentResponseOfDestinyProfileRecordsComponent,
			},
			"ProfileCollectibles": &graphql.Field{
				Type: SingleComponentResponseOfDestinyProfileCollectiblesComponent,
			},
			"ProfileTransitoryData": &graphql.Field{
				Type: SingleComponentResponseOfDestinyProfileTransitoryComponent,
			},
			"Metrics": &graphql.Field{
				Type: SingleComponentResponseOfDestinyMetricsComponent,
			},
			"ProfileStringVariables": &graphql.Field{
				Type: SingleComponentResponseOfDestinyStringVariablesComponent,
			},
			"Characters": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCharacterComponent,
			},
			"CharacterInventories": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyInventoryComponent,
			},
			"CharacterProgressions": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent,
			},
			"CharacterRenderData": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent,
			},
			"CharacterActivities": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent,
			},
			"CharacterEquipment": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyInventoryComponent,
			},
			"CharacterKiosks": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyKiosksComponent,
			},
			"CharacterPlugSets": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent,
			},
			"CharacterUninstancedItemComponents": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterPresentationNodes": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent,
			},
			"CharacterRecords": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent,
			},
			"CharacterCollectibles": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent,
			},
			"CharacterStringVariables": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent,
			},
			"ItemComponents": &graphql.Field{
				Type: DestinyItemComponentSetOfint64,
			},
			"CharacterCurrencyLookups": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent,
			},
		},
	},
)

var DestinyVendorReceiptsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorReceiptsComponent",
		Fields: graphql.Fields{
			"Receipts": &graphql.Field{
				Type: graphql.NewList(DestinyVendorReceipt),
			},
		},
	},
)

var DestinyVendorReceipt = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorReceipt",
		Fields: graphql.Fields{
			"CurrencyPaid": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"ItemReceived": &graphql.Field{
				Type: DestinyItemQuantity,
			},
			"LicenseUnlockHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PurchasedByCharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"RefundPolicy": &graphql.Field{
				Type: graphql.Int,
			},
			"SequenceNumber": &graphql.Field{
				Type: graphql.Int,
			},
			"TimeToExpiration": &graphql.Field{
				Type: graphql.Int,
			},
			"ExpiresOn": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ComponentResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ComponentResponse",
		Fields: graphql.Fields{
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ComponentPrivacySetting = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ComponentPrivacySetting",
		Fields: graphql.Fields{
			"ComponentPrivacySetting": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyVendorReceiptsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyVendorReceiptsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyVendorReceiptsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyInventoryComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInventoryComponent",
		Fields: graphql.Fields{
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyItemComponent),
			},
		},
	},
)

var SingleComponentResponseOfDestinyInventoryComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyInventoryComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyInventoryComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProfileComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileComponent",
		Fields: graphql.Fields{
			"UserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"DateLastPlayed": &graphql.Field{
				Type: graphql.String,
			},
			"VersionsOwned": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterIds": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"SeasonHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"CurrentSeasonHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentSeasonRewardPowerCap": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyProfileComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyProfileComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyProfileComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyPlatformSilverComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyPlatformSilverComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyPlatformSilverComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyKiosksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyKiosksComponent",
		Fields: graphql.Fields{
			"KioskItems": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyKioskItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyKioskItem",
		Fields: graphql.Fields{
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"CanAcquire": &graphql.Field{
				Type: graphql.Boolean,
			},
			"FailureIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"FlavorObjective": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
		},
	},
)

var SingleComponentResponseOfDestinyKiosksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyKiosksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyKiosksComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPlugSetsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlugSetsComponent",
		Fields: graphql.Fields{
			"Plugs": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemPlugBase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPlugBase",
		Fields: graphql.Fields{
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CanInsert": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InsertFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"EnableFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyItemPlug = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPlug",
		Fields: graphql.Fields{
			"PlugObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CanInsert": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InsertFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"EnableFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var SingleComponentResponseOfDestinyPlugSetsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyPlugSetsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyPlugSetsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProfileProgressionComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileProgressionComponent",
		Fields: graphql.Fields{
			"Checklists": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonalArtifact": &graphql.Field{
				Type: DestinyArtifactProfileScoped,
			},
		},
	},
)

var DestinyArtifactProfileScoped = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactProfileScoped",
		Fields: graphql.Fields{
			"ArtifactHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PointProgression": &graphql.Field{
				Type: DestinyProgression,
			},
			"PointsAcquired": &graphql.Field{
				Type: graphql.Int,
			},
			"PowerBonusProgression": &graphql.Field{
				Type: DestinyProgression,
			},
			"PowerBonus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyChecklistDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyChecklistDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"ViewActionString": &graphql.Field{
				Type: graphql.String,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
			"Entries": &graphql.Field{
				Type: graphql.NewList(DestinyChecklistEntryDefinition),
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyChecklistEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyChecklistEntryDefinition",
		Fields: graphql.Fields{
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LocationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"BubbleHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorInteractionIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"Scope": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyProfileProgressionComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyProfileProgressionComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyProfileProgressionComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPresentationNodesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodesComponent",
		Fields: graphql.Fields{
			"Nodes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationNodeComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeComponent",
		Fields: graphql.Fields{
			"State": &graphql.Field{
				Type: graphql.Int,
			},
			"Objective": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
			"ProgressValue": &graphql.Field{
				Type: graphql.Int,
			},
			"CompletionValue": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordCategoryScore": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPresentationNodeState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPresentationNodeState",
		Fields: graphql.Fields{
			"DestinyPresentationNodeState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyPresentationNodesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyPresentationNodesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyPresentationNodesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordsComponent",
		Fields: graphql.Fields{
			"Records": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordCategoriesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordSealsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRecordComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordComponent",
		Fields: graphql.Fields{
			"State": &graphql.Field{
				Type: graphql.Int,
			},
			"Objectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"IntervalObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"IntervalsRedeemedCount": &graphql.Field{
				Type: graphql.Int,
			},
			"CompletedCount": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardVisibilty": &graphql.Field{
				Type: graphql.NewList(graphql.Boolean),
			},
		},
	},
)

var DestinyRecordState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRecordState",
		Fields: graphql.Fields{
			"DestinyRecordState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileRecordsComponent",
		Fields: graphql.Fields{
			"Score": &graphql.Field{
				Type: graphql.Int,
			},
			"ActiveScore": &graphql.Field{
				Type: graphql.Int,
			},
			"LegacyScore": &graphql.Field{
				Type: graphql.Int,
			},
			"LifetimeScore": &graphql.Field{
				Type: graphql.Int,
			},
			"TrackedRecordHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Records": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordCategoriesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordSealsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyProfileRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyProfileRecordsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyProfileRecordsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCollectiblesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectiblesComponent",
		Fields: graphql.Fields{
			"Collectibles": &graphql.Field{
				Type: graphql.Int,
			},
			"CollectionCategoriesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CollectionBadgesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyCollectibleComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleComponent",
		Fields: graphql.Fields{
			"State": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyCollectibleState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleState",
		Fields: graphql.Fields{
			"DestinyCollectibleState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileCollectiblesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileCollectiblesComponent",
		Fields: graphql.Fields{
			"RecentCollectibleHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"NewnessFlaggedCollectibleHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Collectibles": &graphql.Field{
				Type: graphql.Int,
			},
			"CollectionCategoriesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CollectionBadgesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyProfileCollectiblesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyProfileCollectiblesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyProfileCollectiblesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyProfileTransitoryComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileTransitoryComponent",
		Fields: graphql.Fields{
			"PartyMembers": &graphql.Field{
				Type: graphql.NewList(DestinyProfileTransitoryPartyMember),
			},
			"CurrentActivity": &graphql.Field{
				Type: DestinyProfileTransitoryCurrentActivity,
			},
			"Joinability": &graphql.Field{
				Type: DestinyProfileTransitoryJoinability,
			},
			"Tracking": &graphql.Field{
				Type: graphql.NewList(DestinyProfileTransitoryTrackingEntry),
			},
			"LastOrbitedDestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileTransitoryPartyMember = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileTransitoryPartyMember",
		Fields: graphql.Fields{
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"EmblemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"Status": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPartyMemberStates = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPartyMemberStates",
		Fields: graphql.Fields{
			"DestinyPartyMemberStates": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileTransitoryCurrentActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileTransitoryCurrentActivity",
		Fields: graphql.Fields{
			"StartTime": &graphql.Field{
				Type: graphql.String,
			},
			"EndTime": &graphql.Field{
				Type: graphql.String,
			},
			"Score": &graphql.Field{
				Type: graphql.Float,
			},
			"HighestOpposingFactionScore": &graphql.Field{
				Type: graphql.Float,
			},
			"NumberOfOpponents": &graphql.Field{
				Type: graphql.Int,
			},
			"NumberOfPlayers": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileTransitoryJoinability = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileTransitoryJoinability",
		Fields: graphql.Fields{
			"OpenSlots": &graphql.Field{
				Type: graphql.Int,
			},
			"PrivacySetting": &graphql.Field{
				Type: graphql.Int,
			},
			"ClosedReasons": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyGamePrivacySetting = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyGamePrivacySetting",
		Fields: graphql.Fields{
			"DestinyGamePrivacySetting": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyJoinClosedReasons = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyJoinClosedReasons",
		Fields: graphql.Fields{
			"DestinyJoinClosedReasons": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyProfileTransitoryTrackingEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyProfileTransitoryTrackingEntry",
		Fields: graphql.Fields{
			"LocationHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"QuestlineItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"TrackedDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SingleComponentResponseOfDestinyProfileTransitoryComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyProfileTransitoryComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyProfileTransitoryComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyMetricsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMetricsComponent",
		Fields: graphql.Fields{
			"Metrics": &graphql.Field{
				Type: graphql.Int,
			},
			"MetricsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMetricComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMetricComponent",
		Fields: graphql.Fields{
			"Invisible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ObjectiveProgress": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
		},
	},
)

var SingleComponentResponseOfDestinyMetricsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyMetricsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyMetricsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyStringVariablesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStringVariablesComponent",
		Fields: graphql.Fields{
			"IntegerValuesByHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SingleComponentResponseOfDestinyStringVariablesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyStringVariablesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyStringVariablesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterComponent",
		Fields: graphql.Fields{
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"DateLastPlayed": &graphql.Field{
				Type: graphql.String,
			},
			"MinutesPlayedThisSession": &graphql.Field{
				Type: graphql.Int,
			},
			"MinutesPlayedTotal": &graphql.Field{
				Type: graphql.Int,
			},
			"Light": &graphql.Field{
				Type: graphql.Int,
			},
			"Stats": &graphql.Field{
				Type: graphql.Int,
			},
			"RaceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ClassHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RaceType": &graphql.Field{
				Type: graphql.Int,
			},
			"ClassType": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderType": &graphql.Field{
				Type: graphql.Int,
			},
			"EmblemPath": &graphql.Field{
				Type: graphql.String,
			},
			"EmblemBackgroundPath": &graphql.Field{
				Type: graphql.String,
			},
			"EmblemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EmblemColor": &graphql.Field{
				Type: DestinyColor,
			},
			"LevelProgression": &graphql.Field{
				Type: DestinyProgression,
			},
			"BaseCharacterLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"PercentToNextLevel": &graphql.Field{
				Type: graphql.Float,
			},
			"TitleRecordHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRace = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRace",
		Fields: graphql.Fields{
			"DestinyRace": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyRaceDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyRaceDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"RaceType": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderedRaceNames": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderedRaceNamesByGenderHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyClassDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyClassDefinition",
		Fields: graphql.Fields{
			"ClassType": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"GenderedClassNames": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderedClassNamesByGenderHash": &graphql.Field{
				Type: graphql.Int,
			},
			"MentorVendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCharacterComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCharacterComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyInventoryComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyInventoryComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterProgressionComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterProgressionComponent",
		Fields: graphql.Fields{
			"Progressions": &graphql.Field{
				Type: graphql.Int,
			},
			"Factions": &graphql.Field{
				Type: graphql.Int,
			},
			"Milestones": &graphql.Field{
				Type: graphql.Int,
			},
			"Quests": &graphql.Field{
				Type: graphql.NewList(DestinyQuestStatus),
			},
			"UninstancedItemObjectives": &graphql.Field{
				Type: graphql.Int,
			},
			"UninstancedItemPerks": &graphql.Field{
				Type: graphql.Int,
			},
			"Checklists": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonalArtifact": &graphql.Field{
				Type: DestinyArtifactCharacterScoped,
			},
		},
	},
)

var DestinyFactionProgression = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyFactionProgression",
		Fields: graphql.Fields{
			"FactionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"FactionVendorIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressionHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DailyProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"DailyLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"WeeklyProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"WeeklyLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentProgress": &graphql.Field{
				Type: graphql.Int,
			},
			"Level": &graphql.Field{
				Type: graphql.Int,
			},
			"LevelCap": &graphql.Field{
				Type: graphql.Int,
			},
			"StepIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressToNextLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"NextLevelAt": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentResetCount": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonResets": &graphql.Field{
				Type: graphql.NewList(DestinyProgressionResetEntry),
			},
			"RewardItemStates": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyMilestone = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestone",
		Fields: graphql.Fields{
			"MilestoneHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AvailableQuests": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneQuest),
			},
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneChallengeActivity),
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
			"VendorHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Vendors": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneVendor),
			},
			"Rewards": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneRewardCategory),
			},
			"StartDate": &graphql.Field{
				Type: graphql.String,
			},
			"EndDate": &graphql.Field{
				Type: graphql.String,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneQuest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneQuest",
		Fields: graphql.Fields{
			"QuestItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Status": &graphql.Field{
				Type: DestinyQuestStatus,
			},
			"Activity": &graphql.Field{
				Type: DestinyMilestoneActivity,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyChallengeStatus),
			},
		},
	},
)

var DestinyQuestStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyQuestStatus",
		Fields: graphql.Fields{
			"QuestHash": &graphql.Field{
				Type: graphql.Int,
			},
			"StepHash": &graphql.Field{
				Type: graphql.Int,
			},
			"StepObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"Tracked": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Completed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Redeemed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Started": &graphql.Field{
				Type: graphql.Boolean,
			},
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Variants": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneActivityVariant),
			},
		},
	},
)

var DestinyMilestoneActivityVariant = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivityVariant",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CompletionStatus": &graphql.Field{
				Type: DestinyMilestoneActivityCompletionStatus,
			},
			"ActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneActivityCompletionStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivityCompletionStatus",
		Fields: graphql.Fields{
			"Completed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Phases": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneActivityPhase),
			},
		},
	},
)

var DestinyMilestoneActivityPhase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivityPhase",
		Fields: graphql.Fields{
			"Complete": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PhaseHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyChallengeStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyChallengeStatus",
		Fields: graphql.Fields{
			"Objective": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
		},
	},
)

var DestinyMilestoneChallengeActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneChallengeActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyChallengeStatus),
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"BooleanActivityOptions": &graphql.Field{
				Type: graphql.Int,
			},
			"LoadoutRequirementIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"Phases": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneActivityPhase),
			},
		},
	},
)

var DestinyMilestoneVendor = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneVendor",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PreviewItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneRewardCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneRewardCategory",
		Fields: graphql.Fields{
			"RewardCategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Entries": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneRewardEntry),
			},
		},
	},
)

var DestinyMilestoneRewardEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneRewardEntry",
		Fields: graphql.Fields{
			"RewardEntryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Earned": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Redeemed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyMilestoneDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"DisplayPreference": &graphql.Field{
				Type: graphql.Int,
			},
			"Image": &graphql.Field{
				Type: graphql.String,
			},
			"MilestoneType": &graphql.Field{
				Type: graphql.Int,
			},
			"Recruitable": &graphql.Field{
				Type: graphql.Boolean,
			},
			"FriendlyName": &graphql.Field{
				Type: graphql.String,
			},
			"ShowInExplorer": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ShowInMilestones": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ExplorePrioritizesActivityImage": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HasPredictableDates": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Quests": &graphql.Field{
				Type: graphql.Int,
			},
			"Rewards": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorsDisplayTitle": &graphql.Field{
				Type: graphql.String,
			},
			"Vendors": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneVendorDefinition),
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
			"IsInGameMilestone": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneChallengeActivityDefinition),
			},
			"DefaultOrder": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyMilestoneDisplayPreference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneDisplayPreference",
		Fields: graphql.Fields{
			"DestinyMilestoneDisplayPreference": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneType",
		Fields: graphql.Fields{
			"DestinyMilestoneType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneQuestDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneQuestDefinition",
		Fields: graphql.Fields{
			"QuestItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"OverrideImage": &graphql.Field{
				Type: graphql.String,
			},
			"QuestRewards": &graphql.Field{
				Type: DestinyMilestoneQuestRewardsDefinition,
			},
			"Activities": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinationHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneQuestRewardsDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneQuestRewardsDefinition",
		Fields: graphql.Fields{
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneQuestRewardItem),
			},
		},
	},
)

var DestinyMilestoneQuestRewardItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneQuestRewardItem",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"HasConditionalVisibility": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyMilestoneActivityDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivityDefinition",
		Fields: graphql.Fields{
			"ConceptualActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Variants": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneActivityVariantDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneActivityVariantDefinition",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneRewardCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneRewardCategoryDefinition",
		Fields: graphql.Fields{
			"CategoryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CategoryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"RewardEntries": &graphql.Field{
				Type: graphql.Int,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneRewardEntryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneRewardEntryDefinition",
		Fields: graphql.Fields{
			"RewardEntryHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RewardEntryIdentifier": &graphql.Field{
				Type: graphql.String,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneVendorDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneVendorDefinition",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneValueDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneValueDefinition",
		Fields: graphql.Fields{
			"Key": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
		},
	},
)

var DestinyMilestoneChallengeActivityDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneChallengeActivityDefinition",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneChallengeDefinition),
			},
			"ActivityGraphNodes": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneChallengeActivityGraphNodeEntry),
			},
			"Phases": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneChallengeActivityPhase),
			},
		},
	},
)

var DestinyMilestoneChallengeDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneChallengeDefinition",
		Fields: graphql.Fields{
			"ChallengeObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneChallengeActivityGraphNodeEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneChallengeActivityGraphNodeEntry",
		Fields: graphql.Fields{
			"ActivityGraphHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityGraphNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyMilestoneChallengeActivityPhase = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneChallengeActivityPhase",
		Fields: graphql.Fields{
			"PhaseHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemPerksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPerksComponent",
		Fields: graphql.Fields{
			"Perks": &graphql.Field{
				Type: graphql.NewList(DestinyPerkReference),
			},
		},
	},
)

var DestinyPerkReference = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPerkReference",
		Fields: graphql.Fields{
			"PerkHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"IsActive": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Visible": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyArtifactCharacterScoped = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactCharacterScoped",
		Fields: graphql.Fields{
			"ArtifactHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PointsUsed": &graphql.Field{
				Type: graphql.Int,
			},
			"ResetCount": &graphql.Field{
				Type: graphql.Int,
			},
			"Tiers": &graphql.Field{
				Type: graphql.NewList(DestinyArtifactTier),
			},
		},
	},
)

var DestinyArtifactTier = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactTier",
		Fields: graphql.Fields{
			"TierHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsUnlocked": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PointsToUnlock": &graphql.Field{
				Type: graphql.Int,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(DestinyArtifactTierItem),
			},
		},
	},
)

var DestinyArtifactTierItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyArtifactTierItem",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsActive": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterRenderComponent",
		Fields: graphql.Fields{
			"CustomDyes": &graphql.Field{
				Type: graphql.NewList(DyeReference),
			},
			"Customization": &graphql.Field{
				Type: DestinyCharacterCustomization,
			},
			"PeerView": &graphql.Field{
				Type: DestinyCharacterPeerView,
			},
		},
	},
)

var DestinyCharacterCustomization = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterCustomization",
		Fields: graphql.Fields{
			"Personality": &graphql.Field{
				Type: graphql.Int,
			},
			"Face": &graphql.Field{
				Type: graphql.Int,
			},
			"SkinColor": &graphql.Field{
				Type: graphql.Int,
			},
			"LipColor": &graphql.Field{
				Type: graphql.Int,
			},
			"EyeColor": &graphql.Field{
				Type: graphql.Int,
			},
			"HairColors": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"FeatureColors": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"DecalColor": &graphql.Field{
				Type: graphql.Int,
			},
			"WearHelmet": &graphql.Field{
				Type: graphql.Boolean,
			},
			"HairIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"FeatureIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"DecalIndex": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyCharacterPeerView = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterPeerView",
		Fields: graphql.Fields{
			"Equipment": &graphql.Field{
				Type: graphql.NewList(DestinyItemPeerView),
			},
		},
	},
)

var DestinyItemPeerView = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPeerView",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Dyes": &graphql.Field{
				Type: graphql.NewList(DyeReference),
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterActivitiesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterActivitiesComponent",
		Fields: graphql.Fields{
			"DateActivityStarted": &graphql.Field{
				Type: graphql.String,
			},
			"AvailableActivities": &graphql.Field{
				Type: graphql.NewList(DestinyActivity),
			},
			"CurrentActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentActivityModeHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"CurrentActivityModeTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"CurrentPlaylistActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LastCompletedStoryHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsNew": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CanLead": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CanJoin": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsCompleted": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsVisible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DisplayLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"RecommendedLight": &graphql.Field{
				Type: graphql.Int,
			},
			"DifficultyTier": &graphql.Field{
				Type: graphql.Int,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyChallengeStatus),
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"BooleanActivityOptions": &graphql.Field{
				Type: graphql.Int,
			},
			"LoadoutRequirementIndex": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyActivityDifficultyTier = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityDifficultyTier",
		Fields: graphql.Fields{
			"DestinyActivityDifficultyTier": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyKiosksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyKiosksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyBaseItemComponentSetOfuint32 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBaseItemComponentSetOfuint32",
		Fields: graphql.Fields{
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent,
			},
		},
	},
)

var DestinyItemObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemObjectivesComponent",
		Fields: graphql.Fields{
			"Objectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"FlavorObjective": &graphql.Field{
				Type: DestinyObjectiveProgress,
			},
			"DateCompleted": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterRecordsComponent",
		Fields: graphql.Fields{
			"FeaturedRecordHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Records": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordCategoriesRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordSealsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyBaseItemComponentSetOfint64 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBaseItemComponentSetOfint64",
		Fields: graphql.Fields{
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemPerksComponent,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemPerksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemPerksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemComponentSetOfint64 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemComponentSetOfint64",
		Fields: graphql.Fields{
			"Instances": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent,
			},
			"RenderData": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemRenderComponent,
			},
			"Stats": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemStatsComponent,
			},
			"Sockets": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent,
			},
			"ReusablePlugs": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent,
			},
			"PlugObjectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent,
			},
			"TalentGrids": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent,
			},
			"PlugStates": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent,
			},
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfint64AndDestinyItemPerksComponent,
			},
		},
	},
)

var DestinyItemInstanceComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemInstanceComponent",
		Fields: graphql.Fields{
			"DamageType": &graphql.Field{
				Type: graphql.Int,
			},
			"DamageTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PrimaryStat": &graphql.Field{
				Type: DestinyStat,
			},
			"ItemLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"Quality": &graphql.Field{
				Type: graphql.Int,
			},
			"IsEquipped": &graphql.Field{
				Type: graphql.Boolean,
			},
			"CanEquip": &graphql.Field{
				Type: graphql.Boolean,
			},
			"EquipRequiredLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"UnlockHashesRequiredToEquip": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"CannotEquipReason": &graphql.Field{
				Type: graphql.Int,
			},
			"BreakerType": &graphql.Field{
				Type: graphql.Int,
			},
			"BreakerTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Energy": &graphql.Field{
				Type: DestinyItemInstanceEnergy,
			},
		},
	},
)

var DestinyStat = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStat",
		Fields: graphql.Fields{
			"StatHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var EquipFailureReason = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EquipFailureReason",
		Fields: graphql.Fields{
			"EquipFailureReason": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemInstanceEnergy = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemInstanceEnergy",
		Fields: graphql.Fields{
			"EnergyTypeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyType": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyCapacity": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyUsed": &graphql.Field{
				Type: graphql.Int,
			},
			"EnergyUnused": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyUnlockDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyUnlockDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemRenderComponent",
		Fields: graphql.Fields{
			"UseCustomDyes": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ArtRegions": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemStatsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemStatsComponent",
		Fields: graphql.Fields{
			"Stats": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemStatsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemStatsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemSocketsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketsComponent",
		Fields: graphql.Fields{
			"Sockets": &graphql.Field{
				Type: graphql.NewList(DestinyItemSocketState),
			},
		},
	},
)

var DestinyItemSocketState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSocketState",
		Fields: graphql.Fields{
			"PlugHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsEnabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"IsVisible": &graphql.Field{
				Type: graphql.Boolean,
			},
			"EnableFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemReusablePlugsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemReusablePlugsComponent",
		Fields: graphql.Fields{
			"Plugs": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemPlugObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPlugObjectivesComponent",
		Fields: graphql.Fields{
			"ObjectivesPerPlug": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemTalentGridComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTalentGridComponent",
		Fields: graphql.Fields{
			"TalentGridHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Nodes": &graphql.Field{
				Type: graphql.NewList(DestinyTalentNode),
			},
			"IsGridComplete": &graphql.Field{
				Type: graphql.Boolean,
			},
			"GridProgression": &graphql.Field{
				Type: DestinyProgression,
			},
		},
	},
)

var DestinyTalentNode = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNode",
		Fields: graphql.Fields{
			"NodeIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"NodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"State": &graphql.Field{
				Type: graphql.Int,
			},
			"IsActivated": &graphql.Field{
				Type: graphql.Boolean,
			},
			"StepIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"MaterialsToUpgrade": &graphql.Field{
				Type: graphql.NewList(DestinyMaterialRequirement),
			},
			"ActivationGridLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"ProgressPercent": &graphql.Field{
				Type: graphql.Float,
			},
			"Hidden": &graphql.Field{
				Type: graphql.Boolean,
			},
			"NodeStatsBlock": &graphql.Field{
				Type: DestinyTalentNodeStatBlock,
			},
		},
	},
)

var DestinyTalentNodeState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeState",
		Fields: graphql.Fields{
			"DestinyTalentNodeState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyTalentNodeStatBlock = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyTalentNodeStatBlock",
		Fields: graphql.Fields{
			"CurrentStepStats": &graphql.Field{
				Type: graphql.NewList(DestinyStat),
			},
			"NextStepStats": &graphql.Field{
				Type: graphql.NewList(DestinyStat),
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemPlugComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemPlugComponent",
		Fields: graphql.Fields{
			"PlugObjectives": &graphql.Field{
				Type: graphql.NewList(DestinyObjectiveProgress),
			},
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CanInsert": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"InsertFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"EnableFailIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCurrenciesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCurrenciesComponent",
		Fields: graphql.Fields{
			"ItemQuantities": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCharacterResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterResponse",
		Fields: graphql.Fields{
			"Inventory": &graphql.Field{
				Type: SingleComponentResponseOfDestinyInventoryComponent,
			},
			"Character": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCharacterComponent,
			},
			"Progressions": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCharacterProgressionComponent,
			},
			"RenderData": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCharacterRenderComponent,
			},
			"Activities": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCharacterActivitiesComponent,
			},
			"Equipment": &graphql.Field{
				Type: SingleComponentResponseOfDestinyInventoryComponent,
			},
			"Kiosks": &graphql.Field{
				Type: SingleComponentResponseOfDestinyKiosksComponent,
			},
			"PlugSets": &graphql.Field{
				Type: SingleComponentResponseOfDestinyPlugSetsComponent,
			},
			"PresentationNodes": &graphql.Field{
				Type: SingleComponentResponseOfDestinyPresentationNodesComponent,
			},
			"Records": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCharacterRecordsComponent,
			},
			"Collectibles": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCollectiblesComponent,
			},
			"ItemComponents": &graphql.Field{
				Type: DestinyItemComponentSetOfint64,
			},
			"UninstancedItemComponents": &graphql.Field{
				Type: DestinyBaseItemComponentSetOfuint32,
			},
			"CurrencyLookups": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCurrenciesComponent,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCharacterComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCharacterComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCharacterComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCharacterProgressionComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCharacterProgressionComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCharacterProgressionComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCharacterRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCharacterRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCharacterRenderComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCharacterActivitiesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCharacterActivitiesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCharacterActivitiesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCharacterRecordsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCharacterRecordsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCharacterRecordsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCollectiblesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCollectiblesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCollectiblesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyCurrenciesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyCurrenciesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyCurrenciesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var ClanBannerSource = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClanBannerSource",
		Fields: graphql.Fields{
			"ClanBannerSource": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ClanBannerDecal = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ClanBannerDecal",
		Fields: graphql.Fields{
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"ForegroundPath": &graphql.Field{
				Type: graphql.String,
			},
			"BackgroundPath": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyItemResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemResponse",
		Fields: graphql.Fields{
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"Item": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemComponent,
			},
			"Instance": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemInstanceComponent,
			},
			"Objectives": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemPerksComponent,
			},
			"RenderData": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemRenderComponent,
			},
			"Stats": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemStatsComponent,
			},
			"TalentGrid": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemTalentGridComponent,
			},
			"Sockets": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemSocketsComponent,
			},
			"ReusablePlugs": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemReusablePlugsComponent,
			},
			"PlugObjectives": &graphql.Field{
				Type: SingleComponentResponseOfDestinyItemPlugObjectivesComponent,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemInstanceComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemInstanceComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemInstanceComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemObjectivesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemPerksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemPerksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemPerksComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemRenderComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemStatsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemStatsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemStatsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemTalentGridComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemTalentGridComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemTalentGridComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemSocketsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemSocketsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemSocketsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemReusablePlugsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemReusablePlugsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemReusablePlugsComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyItemPlugObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyItemPlugObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyItemPlugObjectivesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorFilter = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorFilter",
		Fields: graphql.Fields{
			"DestinyVendorFilter": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorsResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorsResponse",
		Fields: graphql.Fields{
			"VendorGroups": &graphql.Field{
				Type: SingleComponentResponseOfDestinyVendorGroupComponent,
			},
			"Vendors": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyVendorComponent,
			},
			"Categories": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent,
			},
			"Sales": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent,
			},
			"ItemComponents": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrencyLookups": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCurrenciesComponent,
			},
			"StringVariables": &graphql.Field{
				Type: SingleComponentResponseOfDestinyStringVariablesComponent,
			},
		},
	},
)

var DestinyVendorGroupComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorGroupComponent",
		Fields: graphql.Fields{
			"Groups": &graphql.Field{
				Type: graphql.NewList(DestinyVendorGroup),
			},
		},
	},
)

var DestinyVendorGroup = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorGroup",
		Fields: graphql.Fields{
			"VendorGroupHash": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var SingleComponentResponseOfDestinyVendorGroupComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyVendorGroupComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyVendorGroupComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorBaseComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorBaseComponent",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"NextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorComponent",
		Fields: graphql.Fields{
			"CanPurchase": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Progression": &graphql.Field{
				Type: DestinyProgression,
			},
			"VendorLocationIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonalRank": &graphql.Field{
				Type: graphql.Int,
			},
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"NextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyVendorComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyVendorComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorCategoriesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorCategoriesComponent",
		Fields: graphql.Fields{
			"Categories": &graphql.Field{
				Type: graphql.NewList(DestinyVendorCategory),
			},
		},
	},
)

var DestinyVendorCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorCategory",
		Fields: graphql.Fields{
			"DisplayCategoryIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorSaleItemBaseComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorSaleItemBaseComponent",
		Fields: graphql.Fields{
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"OverrideStyleItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"Costs": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"OverrideNextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"ApiPurchasable": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorSaleItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorSaleItemComponent",
		Fields: graphql.Fields{
			"SaleStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"RequiredUnlocks": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"UnlockStatuses": &graphql.Field{
				Type: graphql.NewList(DestinyUnlockStatus),
			},
			"FailureIndexes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Augments": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemValueVisibility": &graphql.Field{
				Type: graphql.NewList(graphql.Boolean),
			},
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"OverrideStyleItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"Costs": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"OverrideNextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"ApiPurchasable": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var VendorItemStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VendorItemStatus",
		Fields: graphql.Fields{
			"VendorItemStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyUnlockStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyUnlockStatus",
		Fields: graphql.Fields{
			"UnlockHash": &graphql.Field{
				Type: graphql.Int,
			},
			"IsSet": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorItemState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorItemState",
		Fields: graphql.Fields{
			"DestinyVendorItemState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyVendorSaleItemSetComponentOfDestinyVendorSaleItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorSaleItemSetComponentOfDestinyVendorSaleItemComponent",
		Fields: graphql.Fields{
			"SaleItems": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PersonalDestinyVendorSaleItemSetComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PersonalDestinyVendorSaleItemSetComponent",
		Fields: graphql.Fields{
			"SaleItems": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyBaseItemComponentSetOfint32 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyBaseItemComponentSetOfint32",
		Fields: graphql.Fields{
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemPerksComponent,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemPerksComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemPerksComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyItemComponentSetOfint32 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemComponentSetOfint32",
		Fields: graphql.Fields{
			"Instances": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent,
			},
			"RenderData": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemRenderComponent,
			},
			"Stats": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemStatsComponent,
			},
			"Sockets": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent,
			},
			"ReusablePlugs": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent,
			},
			"PlugObjectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent,
			},
			"TalentGrids": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent,
			},
			"PlugStates": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent,
			},
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyItemPerksComponent,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemStatsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemStatsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorResponse",
		Fields: graphql.Fields{
			"Vendor": &graphql.Field{
				Type: SingleComponentResponseOfDestinyVendorComponent,
			},
			"Categories": &graphql.Field{
				Type: SingleComponentResponseOfDestinyVendorCategoriesComponent,
			},
			"Sales": &graphql.Field{
				Type: DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent,
			},
			"ItemComponents": &graphql.Field{
				Type: DestinyItemComponentSetOfint32,
			},
			"CurrencyLookups": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCurrenciesComponent,
			},
			"StringVariables": &graphql.Field{
				Type: SingleComponentResponseOfDestinyStringVariablesComponent,
			},
		},
	},
)

var SingleComponentResponseOfDestinyVendorComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyVendorComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyVendorComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var SingleComponentResponseOfDestinyVendorCategoriesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SingleComponentResponseOfDestinyVendorCategoriesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: DestinyVendorCategoriesComponent,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPublicVendorsResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicVendorsResponse",
		Fields: graphql.Fields{
			"VendorGroups": &graphql.Field{
				Type: SingleComponentResponseOfDestinyVendorGroupComponent,
			},
			"Vendors": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent,
			},
			"Categories": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent,
			},
			"Sales": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent,
			},
			"StringVariables": &graphql.Field{
				Type: SingleComponentResponseOfDestinyStringVariablesComponent,
			},
		},
	},
)

var DestinyPublicVendorComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicVendorComponent",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"NextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyPublicVendorSaleItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicVendorSaleItemComponent",
		Fields: graphql.Fields{
			"VendorItemIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"OverrideStyleItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"Costs": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
			"OverrideNextRefreshDate": &graphql.Field{
				Type: graphql.String,
			},
			"ApiPurchasable": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyVendorSaleItemSetComponentOfDestinyPublicVendorSaleItemComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyVendorSaleItemSetComponentOfDestinyPublicVendorSaleItemComponent",
		Fields: graphql.Fields{
			"SaleItems": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PublicDestinyVendorSaleItemSetComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PublicDestinyVendorSaleItemSetComponent",
		Fields: graphql.Fields{
			"SaleItems": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyCollectibleNodeDetailResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCollectibleNodeDetailResponse",
		Fields: graphql.Fields{
			"Collectibles": &graphql.Field{
				Type: SingleComponentResponseOfDestinyCollectiblesComponent,
			},
			"CollectibleItemComponents": &graphql.Field{
				Type: DestinyItemComponentSetOfuint32,
			},
		},
	},
)

var DestinyItemComponentSetOfuint32 = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemComponentSetOfuint32",
		Fields: graphql.Fields{
			"Instances": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent,
			},
			"RenderData": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent,
			},
			"Stats": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent,
			},
			"Sockets": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent,
			},
			"ReusablePlugs": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent,
			},
			"PlugObjectives": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent,
			},
			"TalentGrids": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent,
			},
			"PlugStates": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent,
			},
			"Objectives": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent,
			},
			"Perks": &graphql.Field{
				Type: DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent",
		Fields: graphql.Fields{
			"Data": &graphql.Field{
				Type: graphql.Int,
			},
			"Privacy": &graphql.Field{
				Type: graphql.Int,
			},
			"Disabled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyActionRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActionRequest",
		Fields: graphql.Fields{
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyCharacterActionRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyCharacterActionRequest",
		Fields: graphql.Fields{
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemActionRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemActionRequest",
		Fields: graphql.Fields{
			"ItemId": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemTransferRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemTransferRequest",
		Fields: graphql.Fields{
			"ItemReferenceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"StackSize": &graphql.Field{
				Type: graphql.Int,
			},
			"TransferToVault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ItemId": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPostmasterTransferRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPostmasterTransferRequest",
		Fields: graphql.Fields{
			"ItemReferenceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"StackSize": &graphql.Field{
				Type: graphql.Int,
			},
			"ItemId": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyEquipItemResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEquipItemResults",
		Fields: graphql.Fields{
			"EquipResults": &graphql.Field{
				Type: graphql.NewList(DestinyEquipItemResult),
			},
		},
	},
)

var DestinyEquipItemResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEquipItemResult",
		Fields: graphql.Fields{
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"EquipStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemSetActionRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemSetActionRequest",
		Fields: graphql.Fields{
			"ItemIds": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyItemStateRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemStateRequest",
		Fields: graphql.Fields{
			"State": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ItemId": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var InventoryChangedResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InventoryChangedResponse",
		Fields: graphql.Fields{
			"AddedInventoryItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemComponent),
			},
			"RemovedInventoryItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemComponent),
			},
		},
	},
)

var DestinyItemChangeResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyItemChangeResponse",
		Fields: graphql.Fields{
			"Item": &graphql.Field{
				Type: DestinyItemResponse,
			},
			"AddedInventoryItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemComponent),
			},
			"RemovedInventoryItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemComponent),
			},
		},
	},
)

var DestinyInsertPlugsActionRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInsertPlugsActionRequest",
		Fields: graphql.Fields{
			"ActionToken": &graphql.Field{
				Type: graphql.String,
			},
			"ItemInstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Plug": &graphql.Field{
				Type: DestinyInsertPlugsRequestEntry,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyInsertPlugsRequestEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyInsertPlugsRequestEntry",
		Fields: graphql.Fields{
			"SocketIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"SocketArrayType": &graphql.Field{
				Type: graphql.Int,
			},
			"PlugItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinySocketArrayType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinySocketArrayType",
		Fields: graphql.Fields{
			"DestinySocketArrayType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPostGameCarnageReportData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPostGameCarnageReportData",
		Fields: graphql.Fields{
			"Period": &graphql.Field{
				Type: graphql.String,
			},
			"StartingPhaseIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityDetails": &graphql.Field{
				Type: DestinyHistoricalStatsActivity,
			},
			"Entries": &graphql.Field{
				Type: graphql.NewList(DestinyPostGameCarnageReportEntry),
			},
			"Teams": &graphql.Field{
				Type: graphql.NewList(DestinyPostGameCarnageReportTeamEntry),
			},
		},
	},
)

var DestinyHistoricalStatsActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsActivity",
		Fields: graphql.Fields{
			"ReferenceId": &graphql.Field{
				Type: graphql.Int,
			},
			"DirectorActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"InstanceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Mode": &graphql.Field{
				Type: graphql.Int,
			},
			"Modes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPrivate": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPostGameCarnageReportEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPostGameCarnageReportEntry",
		Fields: graphql.Fields{
			"Standing": &graphql.Field{
				Type: graphql.Int,
			},
			"Score": &graphql.Field{
				Type: DestinyHistoricalStatsValue,
			},
			"Player": &graphql.Field{
				Type: DestinyPlayer,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
			"Extended": &graphql.Field{
				Type: DestinyPostGameCarnageReportExtendedData,
			},
		},
	},
)

var DestinyHistoricalStatsValue = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsValue",
		Fields: graphql.Fields{
			"StatId": &graphql.Field{
				Type: graphql.String,
			},
			"Basic": &graphql.Field{
				Type: DestinyHistoricalStatsValuePair,
			},
			"Pga": &graphql.Field{
				Type: DestinyHistoricalStatsValuePair,
			},
			"Weighted": &graphql.Field{
				Type: DestinyHistoricalStatsValuePair,
			},
			"ActivityId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyHistoricalStatsValuePair = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsValuePair",
		Fields: graphql.Fields{
			"Value": &graphql.Field{
				Type: graphql.Float,
			},
			"DisplayValue": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyPlayer = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPlayer",
		Fields: graphql.Fields{
			"DestinyUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"CharacterClass": &graphql.Field{
				Type: graphql.String,
			},
			"ClassHash": &graphql.Field{
				Type: graphql.Int,
			},
			"RaceHash": &graphql.Field{
				Type: graphql.Int,
			},
			"GenderHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"LightLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"ClanName": &graphql.Field{
				Type: graphql.String,
			},
			"ClanTag": &graphql.Field{
				Type: graphql.String,
			},
			"EmblemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPostGameCarnageReportExtendedData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPostGameCarnageReportExtendedData",
		Fields: graphql.Fields{
			"Weapons": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalWeaponStats),
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyHistoricalWeaponStats = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalWeaponStats",
		Fields: graphql.Fields{
			"ReferenceId": &graphql.Field{
				Type: graphql.Int,
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyPostGameCarnageReportTeamEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPostGameCarnageReportTeamEntry",
		Fields: graphql.Fields{
			"TeamId": &graphql.Field{
				Type: graphql.Int,
			},
			"Standing": &graphql.Field{
				Type: DestinyHistoricalStatsValue,
			},
			"Score": &graphql.Field{
				Type: DestinyHistoricalStatsValue,
			},
			"TeamName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyReportOffensePgcrRequest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyReportOffensePgcrRequest",
		Fields: graphql.Fields{
			"ReasonCategoryHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ReasonHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"OffendingCharacterId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyReportReasonCategoryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyReportReasonCategoryDefinition",
		Fields: graphql.Fields{
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Reasons": &graphql.Field{
				Type: graphql.Int,
			},
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"Index": &graphql.Field{
				Type: graphql.Int,
			},
			"Redacted": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var DestinyReportReasonDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyReportReasonDefinition",
		Fields: graphql.Fields{
			"ReasonHash": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
		},
	},
)

var DestinyHistoricalStatsDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsDefinition",
		Fields: graphql.Fields{
			"StatId": &graphql.Field{
				Type: graphql.String,
			},
			"Group": &graphql.Field{
				Type: graphql.Int,
			},
			"PeriodTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Modes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Category": &graphql.Field{
				Type: graphql.Int,
			},
			"StatName": &graphql.Field{
				Type: graphql.String,
			},
			"StatNameAbbr": &graphql.Field{
				Type: graphql.String,
			},
			"StatDescription": &graphql.Field{
				Type: graphql.String,
			},
			"UnitType": &graphql.Field{
				Type: graphql.Int,
			},
			"IconImage": &graphql.Field{
				Type: graphql.String,
			},
			"MergeMethod": &graphql.Field{
				Type: graphql.Int,
			},
			"UnitLabel": &graphql.Field{
				Type: graphql.String,
			},
			"Weight": &graphql.Field{
				Type: graphql.Int,
			},
			"MedalTierHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatsGroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatsGroupType",
		Fields: graphql.Fields{
			"DestinyStatsGroupType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PeriodTypeList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PeriodTypeList",
		Fields: graphql.Fields{
			"Destiny.HistoricalStats.Definitions.PeriodType": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyActivityModeTypeList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityModeTypeList",
		Fields: graphql.Fields{
			"Destiny.HistoricalStats.Definitions.DestinyActivityModeType": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyStatsCategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatsCategoryType",
		Fields: graphql.Fields{
			"DestinyStatsCategoryType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var UnitType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UnitType",
		Fields: graphql.Fields{
			"UnitType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyStatsMergeMethod = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyStatsMergeMethod",
		Fields: graphql.Fields{
			"DestinyStatsMergeMethod": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyLeaderboard = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLeaderboard",
		Fields: graphql.Fields{
			"StatId": &graphql.Field{
				Type: graphql.String,
			},
			"Entries": &graphql.Field{
				Type: graphql.NewList(DestinyLeaderboardEntry),
			},
		},
	},
)

var DestinyLeaderboardEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLeaderboardEntry",
		Fields: graphql.Fields{
			"Rank": &graphql.Field{
				Type: graphql.Int,
			},
			"Player": &graphql.Field{
				Type: DestinyPlayer,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"Value": &graphql.Field{
				Type: DestinyHistoricalStatsValue,
			},
		},
	},
)

var DestinyLeaderboardResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyLeaderboardResults",
		Fields: graphql.Fields{
			"FocusMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"FocusCharacterId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyClanAggregateStat = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyClanAggregateStat",
		Fields: graphql.Fields{
			"Mode": &graphql.Field{
				Type: graphql.Int,
			},
			"StatId": &graphql.Field{
				Type: graphql.String,
			},
			"Value": &graphql.Field{
				Type: DestinyHistoricalStatsValue,
			},
		},
	},
)

var DestinyEntitySearchResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEntitySearchResult",
		Fields: graphql.Fields{
			"SuggestedWords": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Results": &graphql.Field{
				Type: SearchResultOfDestinyEntitySearchResultItem,
			},
		},
	},
)

var DestinyEntitySearchResultItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyEntitySearchResultItem",
		Fields: graphql.Fields{
			"Hash": &graphql.Field{
				Type: graphql.Int,
			},
			"EntityType": &graphql.Field{
				Type: graphql.String,
			},
			"DisplayProperties": &graphql.Field{
				Type: DestinyDisplayPropertiesDefinition,
			},
			"Weight": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var SearchResultOfDestinyEntitySearchResultItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfDestinyEntitySearchResultItem",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(DestinyEntitySearchResultItem),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PeriodType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PeriodType",
		Fields: graphql.Fields{
			"PeriodType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyHistoricalStatsByPeriod = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsByPeriod",
		Fields: graphql.Fields{
			"AllTime": &graphql.Field{
				Type: graphql.String,
			},
			"AllTimeTier1": &graphql.Field{
				Type: graphql.String,
			},
			"AllTimeTier2": &graphql.Field{
				Type: graphql.String,
			},
			"AllTimeTier3": &graphql.Field{
				Type: graphql.String,
			},
			"Daily": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalStatsPeriodGroup),
			},
			"Monthly": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalStatsPeriodGroup),
			},
		},
	},
)

var DestinyHistoricalStatsPeriodGroup = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsPeriodGroup",
		Fields: graphql.Fields{
			"Period": &graphql.Field{
				Type: graphql.String,
			},
			"ActivityDetails": &graphql.Field{
				Type: DestinyHistoricalStatsActivity,
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyHistoricalStatsResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "DestinyHistoricalStatsResults",
		Fields: graphql.Fields{},
	},
)

var DestinyHistoricalStatsAccountResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsAccountResult",
		Fields: graphql.Fields{
			"MergedDeletedCharacters": &graphql.Field{
				Type: DestinyHistoricalStatsWithMerged,
			},
			"MergedAllCharacters": &graphql.Field{
				Type: DestinyHistoricalStatsWithMerged,
			},
			"Characters": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalStatsPerCharacter),
			},
		},
	},
)

var DestinyHistoricalStatsWithMerged = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsWithMerged",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.String,
			},
			"Merged": &graphql.Field{
				Type: DestinyHistoricalStatsByPeriod,
			},
		},
	},
)

var DestinyHistoricalStatsPerCharacter = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalStatsPerCharacter",
		Fields: graphql.Fields{
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"Deleted": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Results": &graphql.Field{
				Type: graphql.String,
			},
			"Merged": &graphql.Field{
				Type: DestinyHistoricalStatsByPeriod,
			},
		},
	},
)

var DestinyActivityHistoryResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyActivityHistoryResults",
		Fields: graphql.Fields{
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalStatsPeriodGroup),
			},
		},
	},
)

var DestinyHistoricalWeaponStatsData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyHistoricalWeaponStatsData",
		Fields: graphql.Fields{
			"Weapons": &graphql.Field{
				Type: graphql.NewList(DestinyHistoricalWeaponStats),
			},
		},
	},
)

var DestinyAggregateActivityResults = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyAggregateActivityResults",
		Fields: graphql.Fields{
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyAggregateActivityStats),
			},
		},
	},
)

var DestinyAggregateActivityStats = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyAggregateActivityStats",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Values": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var DestinyMilestoneContent = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneContent",
		Fields: graphql.Fields{
			"About": &graphql.Field{
				Type: graphql.String,
			},
			"Status": &graphql.Field{
				Type: graphql.String,
			},
			"Tips": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"ItemCategories": &graphql.Field{
				Type: graphql.NewList(DestinyMilestoneContentItemCategory),
			},
		},
	},
)

var DestinyMilestoneContentItemCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyMilestoneContentItemCategory",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"ItemHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var DestinyPublicMilestone = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestone",
		Fields: graphql.Fields{
			"MilestoneHash": &graphql.Field{
				Type: graphql.Int,
			},
			"AvailableQuests": &graphql.Field{
				Type: graphql.NewList(DestinyPublicMilestoneQuest),
			},
			"Activities": &graphql.Field{
				Type: graphql.NewList(DestinyPublicMilestoneChallengeActivity),
			},
			"VendorHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Vendors": &graphql.Field{
				Type: graphql.NewList(DestinyPublicMilestoneVendor),
			},
			"StartDate": &graphql.Field{
				Type: graphql.String,
			},
			"EndDate": &graphql.Field{
				Type: graphql.String,
			},
			"Order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPublicMilestoneQuest = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneQuest",
		Fields: graphql.Fields{
			"QuestItemHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Activity": &graphql.Field{
				Type: DestinyPublicMilestoneActivity,
			},
			"Challenges": &graphql.Field{
				Type: graphql.NewList(DestinyPublicMilestoneChallenge),
			},
		},
	},
)

var DestinyPublicMilestoneActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"Variants": &graphql.Field{
				Type: graphql.NewList(DestinyPublicMilestoneActivityVariant),
			},
			"ActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPublicMilestoneActivityVariant = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneActivityVariant",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityModeType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPublicMilestoneChallenge = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneChallenge",
		Fields: graphql.Fields{
			"ObjectiveHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPublicMilestoneChallengeActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneChallengeActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ChallengeObjectiveHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"LoadoutRequirementIndex": &graphql.Field{
				Type: graphql.Int,
			},
			"PhaseHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"BooleanActivityOptions": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var DestinyPublicMilestoneVendor = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicMilestoneVendor",
		Fields: graphql.Fields{
			"VendorHash": &graphql.Field{
				Type: graphql.Int,
			},
			"PreviewItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AwaInitializeResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaInitializeResponse",
		Fields: graphql.Fields{
			"CorrelationId": &graphql.Field{
				Type: graphql.String,
			},
			"SentToSelf": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var AwaPermissionRequested = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaPermissionRequested",
		Fields: graphql.Fields{
			"Type": &graphql.Field{
				Type: graphql.Int,
			},
			"AffectedItemId": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AwaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaType",
		Fields: graphql.Fields{
			"AwaType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AwaUserResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaUserResponse",
		Fields: graphql.Fields{
			"Selection": &graphql.Field{
				Type: graphql.Int,
			},
			"CorrelationId": &graphql.Field{
				Type: graphql.String,
			},
			"Nonce": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var AwaUserSelection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaUserSelection",
		Fields: graphql.Fields{
			"AwaUserSelection": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AwaAuthorizationResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaAuthorizationResult",
		Fields: graphql.Fields{
			"UserSelection": &graphql.Field{
				Type: graphql.Int,
			},
			"ResponseReason": &graphql.Field{
				Type: graphql.Int,
			},
			"DeveloperNote": &graphql.Field{
				Type: graphql.String,
			},
			"ActionToken": &graphql.Field{
				Type: graphql.String,
			},
			"MaximumNumberOfUses": &graphql.Field{
				Type: graphql.Int,
			},
			"ValidUntil": &graphql.Field{
				Type: graphql.String,
			},
			"Type": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AwaResponseReason = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AwaResponseReason",
		Fields: graphql.Fields{
			"AwaResponseReason": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CommunityContentSortMode = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CommunityContentSortMode",
		Fields: graphql.Fields{
			"CommunityContentSortMode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var TrendingCategories = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingCategories",
		Fields: graphql.Fields{
			"Categories": &graphql.Field{
				Type: graphql.NewList(TrendingCategory),
			},
		},
	},
)

var TrendingCategory = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingCategory",
		Fields: graphql.Fields{
			"CategoryName": &graphql.Field{
				Type: graphql.String,
			},
			"Entries": &graphql.Field{
				Type: SearchResultOfTrendingEntry,
			},
			"CategoryId": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TrendingEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntry",
		Fields: graphql.Fields{
			"Weight": &graphql.Field{
				Type: graphql.Float,
			},
			"IsFeatured": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"EntityType": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"Tagline": &graphql.Field{
				Type: graphql.String,
			},
			"Image": &graphql.Field{
				Type: graphql.String,
			},
			"StartDate": &graphql.Field{
				Type: graphql.String,
			},
			"EndDate": &graphql.Field{
				Type: graphql.String,
			},
			"Link": &graphql.Field{
				Type: graphql.String,
			},
			"WebmVideo": &graphql.Field{
				Type: graphql.String,
			},
			"Mp4Video": &graphql.Field{
				Type: graphql.String,
			},
			"FeatureImage": &graphql.Field{
				Type: graphql.String,
			},
			"Items": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"CreationDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TrendingEntryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryType",
		Fields: graphql.Fields{
			"TrendingEntryType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResultOfTrendingEntry = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfTrendingEntry",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(TrendingEntry),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var TrendingDetail = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingDetail",
		Fields: graphql.Fields{
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"EntityType": &graphql.Field{
				Type: graphql.Int,
			},
			"News": &graphql.Field{
				Type: TrendingEntryNews,
			},
			"Support": &graphql.Field{
				Type: TrendingEntrySupportArticle,
			},
			"DestinyItem": &graphql.Field{
				Type: TrendingEntryDestinyItem,
			},
			"DestinyActivity": &graphql.Field{
				Type: TrendingEntryDestinyActivity,
			},
			"DestinyRitual": &graphql.Field{
				Type: TrendingEntryDestinyRitual,
			},
			"Creation": &graphql.Field{
				Type: TrendingEntryCommunityCreation,
			},
		},
	},
)

var TrendingEntryNews = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryNews",
		Fields: graphql.Fields{
			"Article": &graphql.Field{
				Type: ContentItemPublicContract,
			},
		},
	},
)

var TrendingEntrySupportArticle = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntrySupportArticle",
		Fields: graphql.Fields{
			"Article": &graphql.Field{
				Type: ContentItemPublicContract,
			},
		},
	},
)

var TrendingEntryDestinyItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryDestinyItem",
		Fields: graphql.Fields{
			"ItemHash": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var TrendingEntryDestinyActivity = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryDestinyActivity",
		Fields: graphql.Fields{
			"ActivityHash": &graphql.Field{
				Type: graphql.Int,
			},
			"Status": &graphql.Field{
				Type: DestinyPublicActivityStatus,
			},
		},
	},
)

var DestinyPublicActivityStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DestinyPublicActivityStatus",
		Fields: graphql.Fields{
			"ChallengeObjectiveHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"ModifierHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"RewardTooltipItems": &graphql.Field{
				Type: graphql.NewList(DestinyItemQuantity),
			},
		},
	},
)

var TrendingEntryDestinyRitual = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryDestinyRitual",
		Fields: graphql.Fields{
			"Image": &graphql.Field{
				Type: graphql.String,
			},
			"Icon": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Subtitle": &graphql.Field{
				Type: graphql.String,
			},
			"DateStart": &graphql.Field{
				Type: graphql.String,
			},
			"DateEnd": &graphql.Field{
				Type: graphql.String,
			},
			"MilestoneDetails": &graphql.Field{
				Type: DestinyPublicMilestone,
			},
			"EventContent": &graphql.Field{
				Type: DestinyMilestoneContent,
			},
		},
	},
)

var TrendingEntryCommunityCreation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TrendingEntryCommunityCreation",
		Fields: graphql.Fields{
			"Media": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Author": &graphql.Field{
				Type: graphql.String,
			},
			"AuthorMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"PostId": &graphql.Field{
				Type: graphql.Int,
			},
			"Body": &graphql.Field{
				Type: graphql.String,
			},
			"Upvotes": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamDateRange = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamDateRange",
		Fields: graphql.Fields{
			"FireteamDateRange": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamPlatform = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamPlatform",
		Fields: graphql.Fields{
			"FireteamPlatform": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamPublicSearchOption = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamPublicSearchOption",
		Fields: graphql.Fields{
			"FireteamPublicSearchOption": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamSlotSearch = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamSlotSearch",
		Fields: graphql.Fields{
			"FireteamSlotSearch": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamSummary = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamSummary",
		Fields: graphql.Fields{
			"FireteamId": &graphql.Field{
				Type: graphql.Int,
			},
			"GroupId": &graphql.Field{
				Type: graphql.Int,
			},
			"Platform": &graphql.Field{
				Type: graphql.Int,
			},
			"ActivityType": &graphql.Field{
				Type: graphql.Int,
			},
			"IsImmediate": &graphql.Field{
				Type: graphql.Boolean,
			},
			"ScheduledTime": &graphql.Field{
				Type: graphql.String,
			},
			"OwnerMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"PlayerSlotCount": &graphql.Field{
				Type: graphql.Int,
			},
			"AlternateSlotCount": &graphql.Field{
				Type: graphql.Int,
			},
			"AvailablePlayerSlotCount": &graphql.Field{
				Type: graphql.Int,
			},
			"AvailableAlternateSlotCount": &graphql.Field{
				Type: graphql.Int,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"DateCreated": &graphql.Field{
				Type: graphql.String,
			},
			"DateModified": &graphql.Field{
				Type: graphql.String,
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Locale": &graphql.Field{
				Type: graphql.String,
			},
			"IsValid": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DatePlayerModified": &graphql.Field{
				Type: graphql.String,
			},
			"TitleBeforeModeration": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SearchResultOfFireteamSummary = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfFireteamSummary",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(FireteamSummary),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var FireteamResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamResponse",
		Fields: graphql.Fields{
			"Summary": &graphql.Field{
				Type: FireteamSummary,
			},
			"Members": &graphql.Field{
				Type: graphql.NewList(FireteamMember),
			},
			"Alternates": &graphql.Field{
				Type: graphql.NewList(FireteamMember),
			},
		},
	},
)

var FireteamMember = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamMember",
		Fields: graphql.Fields{
			"DestinyUserInfo": &graphql.Field{
				Type: FireteamUserInfoCard,
			},
			"BungieNetUserInfo": &graphql.Field{
				Type: UserInfoCard,
			},
			"CharacterId": &graphql.Field{
				Type: graphql.Int,
			},
			"DateJoined": &graphql.Field{
				Type: graphql.String,
			},
			"HasMicrophone": &graphql.Field{
				Type: graphql.Boolean,
			},
			"LastPlatformInviteAttemptDate": &graphql.Field{
				Type: graphql.String,
			},
			"LastPlatformInviteAttemptResult": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamUserInfoCard = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamUserInfoCard",
		Fields: graphql.Fields{
			"FireteamDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"FireteamMembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"SupplementalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"IconPath": &graphql.Field{
				Type: graphql.String,
			},
			"CrossSaveOverride": &graphql.Field{
				Type: graphql.Int,
			},
			"ApplicableMembershipTypes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"IsPublic": &graphql.Field{
				Type: graphql.Boolean,
			},
			"MembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"MembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FireteamPlatformInviteResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FireteamPlatformInviteResult",
		Fields: graphql.Fields{
			"FireteamPlatformInviteResult": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var SearchResultOfFireteamResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResultOfFireteamResponse",
		Fields: graphql.Fields{
			"Results": &graphql.Field{
				Type: graphql.NewList(FireteamResponse),
			},
			"TotalResults": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Query": &graphql.Field{
				Type: PagedQuery,
			},
			"ReplacementContinuationToken": &graphql.Field{
				Type: graphql.String,
			},
			"UseTotalResults": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var BungieFriendListResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieFriendListResponse",
		Fields: graphql.Fields{
			"Friends": &graphql.Field{
				Type: graphql.NewList(BungieFriend),
			},
		},
	},
)

var BungieFriend = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieFriend",
		Fields: graphql.Fields{
			"LastSeenAsMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"LastSeenAsBungieMembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
			"OnlineStatus": &graphql.Field{
				Type: graphql.Int,
			},
			"OnlineTitle": &graphql.Field{
				Type: graphql.Int,
			},
			"Relationship": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetUser": &graphql.Field{
				Type: GeneralUser,
			},
		},
	},
)

var PresenceStatus = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PresenceStatus",
		Fields: graphql.Fields{
			"PresenceStatus": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PresenceOnlineStateFlags = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PresenceOnlineStateFlags",
		Fields: graphql.Fields{
			"PresenceOnlineStateFlags": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var FriendRelationshipState = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "FriendRelationshipState",
		Fields: graphql.Fields{
			"FriendRelationshipState": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var BungieFriendRequestListResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BungieFriendRequestListResponse",
		Fields: graphql.Fields{
			"IncomingRequests": &graphql.Field{
				Type: graphql.NewList(BungieFriend),
			},
			"OutgoingRequests": &graphql.Field{
				Type: graphql.NewList(BungieFriend),
			},
		},
	},
)

var PlatformFriendType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlatformFriendType",
		Fields: graphql.Fields{
			"PlatformFriendType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var PlatformFriendResponse = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlatformFriendResponse",
		Fields: graphql.Fields{
			"ItemsPerPage": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentPage": &graphql.Field{
				Type: graphql.Int,
			},
			"HasMore": &graphql.Field{
				Type: graphql.Boolean,
			},
			"PlatformFriends": &graphql.Field{
				Type: graphql.NewList(PlatformFriend),
			},
		},
	},
)

var PlatformFriend = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PlatformFriend",
		Fields: graphql.Fields{
			"PlatformDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"FriendPlatform": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"DestinyMembershipType": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieNetMembershipId": &graphql.Field{
				Type: graphql.Int,
			},
			"BungieGlobalDisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"BungieGlobalDisplayNameCode": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CoreSettingsConfiguration = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreSettingsConfiguration",
		Fields: graphql.Fields{
			"Environment": &graphql.Field{
				Type: graphql.String,
			},
			"Systems": &graphql.Field{
				Type: graphql.String,
			},
			"IgnoreReasons": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ForumCategories": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"GroupAvatars": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"DestinyMembershipTypes": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"RecruitmentPlatformTags": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"RecruitmentMiscTags": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"RecruitmentActivities": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"UserContentLocales": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"SystemContentLocales": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerDecals": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerDecalColors": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerGonfalons": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerGonfalonColors": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerGonfalonDetails": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerGonfalonDetailColors": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"ClanBannerStandards": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
			"Destiny2CoreSettings": &graphql.Field{
				Type: Destiny2CoreSettings,
			},
			"EmailSettings": &graphql.Field{
				Type: EmailSettings,
			},
			"FireteamActivities": &graphql.Field{
				Type: graphql.NewList(CoreSetting),
			},
		},
	},
)

var CoreSystem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreSystem",
		Fields: graphql.Fields{
			"Enabled": &graphql.Field{
				Type: graphql.Boolean,
			},
			"Parameters": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CoreSetting = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreSetting",
		Fields: graphql.Fields{
			"Identifier": &graphql.Field{
				Type: graphql.String,
			},
			"IsDefault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DisplayName": &graphql.Field{
				Type: graphql.String,
			},
			"Summary": &graphql.Field{
				Type: graphql.String,
			},
			"ImagePath": &graphql.Field{
				Type: graphql.String,
			},
			"ChildSettings": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var Destiny2CoreSettings = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Destiny2CoreSettings",
		Fields: graphql.Fields{
			"CollectionRootNode": &graphql.Field{
				Type: graphql.Int,
			},
			"BadgesRootNode": &graphql.Field{
				Type: graphql.Int,
			},
			"RecordsRootNode": &graphql.Field{
				Type: graphql.Int,
			},
			"MedalsRootNode": &graphql.Field{
				Type: graphql.Int,
			},
			"MetricsRootNode": &graphql.Field{
				Type: graphql.Int,
			},
			"ActiveTriumphsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ActiveSealsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LegacyTriumphsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LegacySealsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"MedalsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"ExoticCatalystsRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"LoreRootNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentRankProgressionHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"UndiscoveredCollectibleImage": &graphql.Field{
				Type: graphql.String,
			},
			"AmmoTypeHeavyIcon": &graphql.Field{
				Type: graphql.String,
			},
			"AmmoTypeSpecialIcon": &graphql.Field{
				Type: graphql.String,
			},
			"AmmoTypePrimaryIcon": &graphql.Field{
				Type: graphql.String,
			},
			"CurrentSeasonalArtifactHash": &graphql.Field{
				Type: graphql.Int,
			},
			"CurrentSeasonHash": &graphql.Field{
				Type: graphql.Int,
			},
			"SeasonalChallengesPresentationNodeHash": &graphql.Field{
				Type: graphql.Int,
			},
			"FutureSeasonHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"PastSeasonHashes": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var EmailSettings = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EmailSettings",
		Fields: graphql.Fields{
			"OptInDefinitions": &graphql.Field{
				Type: graphql.String,
			},
			"SubscriptionDefinitions": &graphql.Field{
				Type: graphql.String,
			},
			"Views": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var EmailOptInDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EmailOptInDefinition",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
			"SetByDefault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"DependentSubscriptions": &graphql.Field{
				Type: graphql.NewList(EmailSubscriptionDefinition),
			},
		},
	},
)

var OptInFlags = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OptInFlags",
		Fields: graphql.Fields{
			"OptInFlags": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var EmailSubscriptionDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EmailSubscriptionDefinition",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Localization": &graphql.Field{
				Type: graphql.String,
			},
			"Value": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var EMailSettingLocalization = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EMailSettingLocalization",
		Fields: graphql.Fields{
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var EMailSettingSubscriptionLocalization = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EMailSettingSubscriptionLocalization",
		Fields: graphql.Fields{
			"UnknownUserDescription": &graphql.Field{
				Type: graphql.String,
			},
			"RegisteredUserDescription": &graphql.Field{
				Type: graphql.String,
			},
			"UnregisteredUserDescription": &graphql.Field{
				Type: graphql.String,
			},
			"UnknownUserActionText": &graphql.Field{
				Type: graphql.String,
			},
			"KnownUserActionText": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var EmailViewDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EmailViewDefinition",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"ViewSettings": &graphql.Field{
				Type: graphql.NewList(EmailViewDefinitionSetting),
			},
		},
	},
)

var EmailViewDefinitionSetting = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "EmailViewDefinitionSetting",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Localization": &graphql.Field{
				Type: graphql.String,
			},
			"SetByDefault": &graphql.Field{
				Type: graphql.Boolean,
			},
			"OptInAggregateValue": &graphql.Field{
				Type: graphql.Int,
			},
			"Subscriptions": &graphql.Field{
				Type: graphql.NewList(EmailSubscriptionDefinition),
			},
		},
	},
)

var GlobalAlert = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GlobalAlert",
		Fields: graphql.Fields{
			"AlertKey": &graphql.Field{
				Type: graphql.String,
			},
			"AlertHtml": &graphql.Field{
				Type: graphql.String,
			},
			"AlertTimestamp": &graphql.Field{
				Type: graphql.String,
			},
			"AlertLink": &graphql.Field{
				Type: graphql.String,
			},
			"AlertLevel": &graphql.Field{
				Type: graphql.Int,
			},
			"AlertType": &graphql.Field{
				Type: graphql.Int,
			},
			"StreamInfo": &graphql.Field{
				Type: StreamInfo,
			},
		},
	},
)

var GlobalAlertLevel = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GlobalAlertLevel",
		Fields: graphql.Fields{
			"GlobalAlertLevel": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var GlobalAlertType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GlobalAlertType",
		Fields: graphql.Fields{
			"GlobalAlertType": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var StreamInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StreamInfo",
		Fields: graphql.Fields{
			"ChannelName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
