package data

type ApplicationScopes struct {
	ApplicationScopes int `json:"applicationscopes"`
}
type ApiUsage struct {
	ApiCalls          []Series `json:"apicalls"`
	ThrottledRequests []Series `json:"throttledrequests"`
}
type Series struct {
	Datapoints []Datapoint `json:"datapoints"`
	Target     string      `json:"target"`
}
type Datapoint struct {
	Time  string  `json:"time"`
	Count float64 `json:"count"`
}
type Application struct {
	ApplicationId             int                    `json:"applicationid"`
	Name                      string                 `json:"name"`
	RedirectUrl               string                 `json:"redirecturl"`
	Link                      string                 `json:"link"`
	Scope                     int                    `json:"scope"`
	Origin                    string                 `json:"origin"`
	Status                    int                    `json:"status"`
	CreationDate              string                 `json:"creationdate"`
	StatusChanged             string                 `json:"statuschanged"`
	FirstPublished            string                 `json:"firstpublished"`
	Team                      []ApplicationDeveloper `json:"team"`
	OverrideAuthorizeViewName string                 `json:"overrideauthorizeviewname"`
}
type ApplicationStatus struct {
	ApplicationStatus int `json:"applicationstatus"`
}
type ApplicationDeveloper struct {
	Role           int          `json:"role"`
	ApiEulaVersion int          `json:"apieulaversion"`
	User           UserInfoCard `json:"user"`
}
type DeveloperRole struct {
	DeveloperRole int `json:"developerrole"`
}
type UserMembership struct {
	MembershipType              int    `json:"membershiptype"`
	MembershipId                int    `json:"membershipid"`
	DisplayName                 string `json:"displayname"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type BungieMembershipType struct {
	BungieMembershipType int `json:"bungiemembershiptype"`
}
type CrossSaveUserMembership struct {
	CrossSaveOverride           int    `json:"crosssaveoverride"`
	ApplicableMembershipTypes   []int  `json:"applicablemembershiptypes"`
	IsPublic                    bool   `json:"ispublic"`
	MembershipType              int    `json:"membershiptype"`
	MembershipId                int    `json:"membershipid"`
	DisplayName                 string `json:"displayname"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type UserInfoCard struct {
	SupplementalDisplayName     string `json:"supplementaldisplayname"`
	IconPath                    string `json:"iconpath"`
	CrossSaveOverride           int    `json:"crosssaveoverride"`
	ApplicableMembershipTypes   []int  `json:"applicablemembershiptypes"`
	IsPublic                    bool   `json:"ispublic"`
	MembershipType              int    `json:"membershiptype"`
	MembershipId                int    `json:"membershipid"`
	DisplayName                 string `json:"displayname"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type GeneralUser struct {
	MembershipId                      string            `json:"membershipid"`
	UniqueName                        string            `json:"uniquename"`
	NormalizedName                    string            `json:"normalizedname"`
	DisplayName                       string            `json:"displayname"`
	ProfilePicture                    int               `json:"profilepicture"`
	ProfileTheme                      int               `json:"profiletheme"`
	UserTitle                         int               `json:"usertitle"`
	SuccessMessageFlags               string            `json:"successmessageflags"`
	IsDeleted                         bool              `json:"isdeleted"`
	About                             string            `json:"about"`
	FirstAccess                       string            `json:"firstaccess"`
	LastUpdate                        string            `json:"lastupdate"`
	LegacyPortalUID                   int               `json:"legacyportaluid"`
	Context                           UserToUserContext `json:"context"`
	PsnDisplayName                    string            `json:"psndisplayname"`
	XboxDisplayName                   string            `json:"xboxdisplayname"`
	FbDisplayName                     string            `json:"fbdisplayname"`
	ShowActivity                      bool              `json:"showactivity"`
	Locale                            string            `json:"locale"`
	LocaleInheritDefault              bool              `json:"localeinheritdefault"`
	LastBanReportId                   int               `json:"lastbanreportid"`
	ShowGroupMessaging                bool              `json:"showgroupmessaging"`
	ProfilePicturePath                string            `json:"profilepicturepath"`
	ProfilePictureWidePath            string            `json:"profilepicturewidepath"`
	ProfileThemeName                  string            `json:"profilethemename"`
	UserTitleDisplay                  string            `json:"usertitledisplay"`
	StatusText                        string            `json:"statustext"`
	StatusDate                        string            `json:"statusdate"`
	ProfileBanExpire                  string            `json:"profilebanexpire"`
	BlizzardDisplayName               string            `json:"blizzarddisplayname"`
	SteamDisplayName                  string            `json:"steamdisplayname"`
	StadiaDisplayName                 string            `json:"stadiadisplayname"`
	TwitchDisplayName                 string            `json:"twitchdisplayname"`
	CachedBungieGlobalDisplayName     string            `json:"cachedbungieglobaldisplayname"`
	CachedBungieGlobalDisplayNameCode int               `json:"cachedbungieglobaldisplaynamecode"`
}
type UserToUserContext struct {
	IsFollowing         bool           `json:"isfollowing"`
	IgnoreStatus        IgnoreResponse `json:"ignorestatus"`
	GlobalIgnoreEndDate string         `json:"globalignoreenddate"`
}
type IgnoreResponse struct {
	IsIgnored   bool `json:"isignored"`
	IgnoreFlags int  `json:"ignoreflags"`
}
type IgnoreStatus struct {
	IgnoreStatus int `json:"ignorestatus"`
}
type GetCredentialTypesForAccountResponse struct {
	CredentialType        int    `json:"credentialtype"`
	CredentialDisplayName string `json:"credentialdisplayname"`
	IsPublic              bool   `json:"ispublic"`
	CredentialAsString    string `json:"credentialasstring"`
}
type BungieCredentialType struct {
	BungieCredentialType int `json:"bungiecredentialtype"`
}
type UserTheme struct {
	UserThemeId          int    `json:"userthemeid"`
	UserThemeName        string `json:"userthemename"`
	UserThemeDescription string `json:"userthemedescription"`
}
type UserMembershipData struct {
	DestinyMemberships  []GroupUserInfoCard `json:"destinymemberships"`
	PrimaryMembershipId int                 `json:"primarymembershipid"`
	BungieNetUser       GeneralUser         `json:"bungienetuser"`
}
type GroupUserInfoCard struct {
	LastSeenDisplayName         string `json:"lastseendisplayname"`
	LastSeenDisplayNameType     int    `json:"lastseendisplaynametype"`
	SupplementalDisplayName     string `json:"supplementaldisplayname"`
	IconPath                    string `json:"iconpath"`
	CrossSaveOverride           int    `json:"crosssaveoverride"`
	ApplicableMembershipTypes   []int  `json:"applicablemembershiptypes"`
	IsPublic                    bool   `json:"ispublic"`
	MembershipType              int    `json:"membershiptype"`
	MembershipId                int    `json:"membershipid"`
	DisplayName                 string `json:"displayname"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type HardLinkedUserMembership struct {
	MembershipType                  int `json:"membershiptype"`
	MembershipId                    int `json:"membershipid"`
	CrossSaveOverriddenType         int `json:"crosssaveoverriddentype"`
	CrossSaveOverriddenMembershipId int `json:"crosssaveoverriddenmembershipid"`
}
type UserSearchResponse struct {
	SearchResults []UserSearchResponseDetail `json:"searchresults"`
	Page          int                        `json:"page"`
	HasMore       bool                       `json:"hasmore"`
}
type UserSearchResponseDetail struct {
	BungieGlobalDisplayName     string         `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int            `json:"bungieglobaldisplaynamecode"`
	BungieNetMembershipId       int            `json:"bungienetmembershipid"`
	DestinyMemberships          []UserInfoCard `json:"destinymemberships"`
}
type ContentTypeDescription struct {
	CType                       string                       `json:"ctype"`
	Name                        string                       `json:"name"`
	ContentDescription          string                       `json:"contentdescription"`
	PreviewImage                string                       `json:"previewimage"`
	Priority                    int                          `json:"priority"`
	Reminder                    string                       `json:"reminder"`
	Properties                  []ContentTypeProperty        `json:"properties"`
	TagMetadata                 []TagMetadataDefinition      `json:"tagmetadata"`
	TagMetadataItems            TagMetadataItem              `json:"tagmetadataitems"`
	UsageExamples               []string                     `json:"usageexamples"`
	ShowInContentEditor         bool                         `json:"showincontenteditor"`
	TypeOf                      string                       `json:"typeof"`
	BindIdentifierToProperty    string                       `json:"bindidentifiertoproperty"`
	BoundRegex                  string                       `json:"boundregex"`
	ForceIdentifierBinding      bool                         `json:"forceidentifierbinding"`
	AllowComments               bool                         `json:"allowcomments"`
	AutoEnglishPropertyFallback bool                         `json:"autoenglishpropertyfallback"`
	BulkUploadable              bool                         `json:"bulkuploadable"`
	Previews                    []ContentPreview             `json:"previews"`
	SuppressCmsPath             bool                         `json:"suppresscmspath"`
	PropertySections            []ContentTypePropertySection `json:"propertysections"`
}
type ContentTypeProperty struct {
	Name                           string                    `json:"name"`
	RootPropertyName               string                    `json:"rootpropertyname"`
	ReadableName                   string                    `json:"readablename"`
	Value                          string                    `json:"value"`
	PropertyDescription            string                    `json:"propertydescription"`
	Localizable                    bool                      `json:"localizable"`
	Fallback                       bool                      `json:"fallback"`
	Enabled                        bool                      `json:"enabled"`
	Order                          int                       `json:"order"`
	Visible                        bool                      `json:"visible"`
	IsTitle                        bool                      `json:"istitle"`
	Required                       bool                      `json:"required"`
	MaxLength                      int                       `json:"maxlength"`
	MaxByteLength                  int                       `json:"maxbytelength"`
	MaxFileSize                    int                       `json:"maxfilesize"`
	Regexp                         string                    `json:"regexp"`
	ValidateAs                     string                    `json:"validateas"`
	RssAttribute                   string                    `json:"rssattribute"`
	VisibleDependency              string                    `json:"visibledependency"`
	VisibleOn                      string                    `json:"visibleon"`
	Datatype                       int                       `json:"datatype"`
	Attributes                     string                    `json:"attributes"`
	ChildProperties                []ContentTypeProperty     `json:"childproperties"`
	ContentTypeAllowed             string                    `json:"contenttypeallowed"`
	BindToProperty                 string                    `json:"bindtoproperty"`
	BoundRegex                     string                    `json:"boundregex"`
	RepresentationSelection        string                    `json:"representationselection"`
	DefaultValues                  []ContentTypeDefaultValue `json:"defaultvalues"`
	IsExternalAllowed              bool                      `json:"isexternalallowed"`
	PropertySection                string                    `json:"propertysection"`
	Weight                         int                       `json:"weight"`
	Entitytype                     string                    `json:"entitytype"`
	IsCombo                        bool                      `json:"iscombo"`
	SuppressProperty               bool                      `json:"suppressproperty"`
	LegalContentTypes              []string                  `json:"legalcontenttypes"`
	RepresentationValidationString string                    `json:"representationvalidationstring"`
	MinWidth                       int                       `json:"minwidth"`
	MaxWidth                       int                       `json:"maxwidth"`
	MinHeight                      int                       `json:"minheight"`
	MaxHeight                      int                       `json:"maxheight"`
	IsVideo                        bool                      `json:"isvideo"`
	IsImage                        bool                      `json:"isimage"`
}
type ContentPropertyDataTypeEnum struct {
	ContentPropertyDataTypeEnum int `json:"contentpropertydatatypeenum"`
}
type ContentTypeDefaultValue struct {
	WhenClause   string `json:"whenclause"`
	WhenValue    string `json:"whenvalue"`
	DefaultValue string `json:"defaultvalue"`
}
type TagMetadataDefinition struct {
	Description string            `json:"description"`
	Order       int               `json:"order"`
	Items       []TagMetadataItem `json:"items"`
	Datatype    string            `json:"datatype"`
	Name        string            `json:"name"`
	IsRequired  bool              `json:"isrequired"`
}
type TagMetadataItem struct {
	Description string   `json:"description"`
	TagText     string   `json:"tagtext"`
	Groups      []string `json:"groups"`
	IsDefault   bool     `json:"isdefault"`
	Name        string   `json:"name"`
}
type ContentPreview struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	ItemInSet  bool   `json:"iteminset"`
	SetTag     string `json:"settag"`
	SetNesting int    `json:"setnesting"`
	UseSetId   int    `json:"usesetid"`
}
type ContentTypePropertySection struct {
	Name         string `json:"name"`
	ReadableName string `json:"readablename"`
	Collapsed    bool   `json:"collapsed"`
}
type ContentItemPublicContract struct {
	ContentId                   int                     `json:"contentid"`
	CType                       string                  `json:"ctype"`
	CmsPath                     string                  `json:"cmspath"`
	CreationDate                string                  `json:"creationdate"`
	ModifyDate                  string                  `json:"modifydate"`
	AllowComments               bool                    `json:"allowcomments"`
	HasAgeGate                  bool                    `json:"hasagegate"`
	MinimumAge                  int                     `json:"minimumage"`
	RatingImagePath             string                  `json:"ratingimagepath"`
	Author                      GeneralUser             `json:"author"`
	AutoEnglishPropertyFallback bool                    `json:"autoenglishpropertyfallback"`
	Properties                  interface{}             `json:"properties"`
	Representations             []ContentRepresentation `json:"representations"`
	Tags                        []string                `json:"tags"`
	CommentSummary              CommentSummary          `json:"commentsummary"`
}
type ContentRepresentation struct {
	Name             string `json:"name"`
	Path             string `json:"path"`
	ValidationString string `json:"validationstring"`
}
type CommentSummary struct {
	TopicId      int `json:"topicid"`
	CommentCount int `json:"commentcount"`
}
type SearchResult struct {
	TotalResults                 int        `json:"totalresults"`
	HasMore                      bool       `json:"hasmore"`
	Query                        PagedQuery `json:"query"`
	ReplacementContinuationToken string     `json:"replacementcontinuationtoken"`
	UseTotalResults              bool       `json:"usetotalresults"`
}
type PagedQuery struct {
	ItemsPerPage             int    `json:"itemsperpage"`
	CurrentPage              int    `json:"currentpage"`
	RequestContinuationToken string `json:"requestcontinuationtoken"`
}
type SearchResultOfContentItemPublicContract struct {
	Results                      []ContentItemPublicContract `json:"results"`
	TotalResults                 int                         `json:"totalresults"`
	HasMore                      bool                        `json:"hasmore"`
	Query                        PagedQuery                  `json:"query"`
	ReplacementContinuationToken string                      `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                        `json:"usetotalresults"`
}
type ForumTopicsCategoryFiltersEnum struct {
	ForumTopicsCategoryFiltersEnum int `json:"forumtopicscategoryfiltersenum"`
}
type ForumTopicsQuickDateEnum struct {
	ForumTopicsQuickDateEnum int `json:"forumtopicsquickdateenum"`
}
type ForumTopicsSortEnum struct {
	ForumTopicsSortEnum int `json:"forumtopicssortenum"`
}
type PostResponse struct {
	LastReplyTimestamp  string         `json:"lastreplytimestamp"`
	IsPinned            bool           `json:"ispinned"`
	UrlMediaType        int            `json:"urlmediatype"`
	Thumbnail           string         `json:"thumbnail"`
	Popularity          int            `json:"popularity"`
	IsActive            bool           `json:"isactive"`
	IsAnnouncement      bool           `json:"isannouncement"`
	UserRating          int            `json:"userrating"`
	UserHasRated        bool           `json:"userhasrated"`
	UserHasMutedPost    bool           `json:"userhasmutedpost"`
	LatestReplyPostId   int            `json:"latestreplypostid"`
	LatestReplyAuthorId int            `json:"latestreplyauthorid"`
	IgnoreStatus        IgnoreResponse `json:"ignorestatus"`
	Locale              string         `json:"locale"`
}
type ForumMediaType struct {
	ForumMediaType int `json:"forummediatype"`
}
type ForumPostPopularity struct {
	ForumPostPopularity int `json:"forumpostpopularity"`
}
type ForumPostCategoryEnums struct {
	ForumPostCategoryEnums int `json:"forumpostcategoryenums"`
}
type ForumFlagsEnum struct {
	ForumFlagsEnum int `json:"forumflagsenum"`
}
type SearchResultOfPostResponse struct {
	Results                      []PostResponse `json:"results"`
	TotalResults                 int            `json:"totalresults"`
	HasMore                      bool           `json:"hasmore"`
	Query                        PagedQuery     `json:"query"`
	ReplacementContinuationToken string         `json:"replacementcontinuationtoken"`
	UseTotalResults              bool           `json:"usetotalresults"`
}
type PostSearchResponse struct {
	RelatedPosts                 []PostResponse           `json:"relatedposts"`
	Authors                      []GeneralUser            `json:"authors"`
	Groups                       []GroupResponse          `json:"groups"`
	SearchedTags                 []TagResponse            `json:"searchedtags"`
	Polls                        []PollResponse           `json:"polls"`
	RecruitmentDetails           []ForumRecruitmentDetail `json:"recruitmentdetails"`
	AvailablePages               int                      `json:"availablepages"`
	Results                      []PostResponse           `json:"results"`
	TotalResults                 int                      `json:"totalresults"`
	HasMore                      bool                     `json:"hasmore"`
	Query                        PagedQuery               `json:"query"`
	ReplacementContinuationToken string                   `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                     `json:"usetotalresults"`
}
type GroupResponse struct {
	Detail                                   GroupV2              `json:"detail"`
	Founder                                  GroupMember          `json:"founder"`
	AlliedIds                                []int                `json:"alliedids"`
	ParentGroup                              GroupV2              `json:"parentgroup"`
	AllianceStatus                           int                  `json:"alliancestatus"`
	GroupJoinInviteCount                     int                  `json:"groupjoininvitecount"`
	CurrentUserMembershipsInactiveForDestiny bool                 `json:"currentusermembershipsinactivefordestiny"`
	CurrentUserMemberMap                     GroupMember          `json:"currentusermembermap"`
	CurrentUserPotentialMemberMap            GroupPotentialMember `json:"currentuserpotentialmembermap"`
}
type GroupV2 struct {
	GroupId                            int                          `json:"groupid"`
	Name                               string                       `json:"name"`
	GroupType                          int                          `json:"grouptype"`
	MembershipIdCreated                int                          `json:"membershipidcreated"`
	CreationDate                       string                       `json:"creationdate"`
	ModificationDate                   string                       `json:"modificationdate"`
	About                              string                       `json:"about"`
	Tags                               []string                     `json:"tags"`
	MemberCount                        int                          `json:"membercount"`
	IsPublic                           bool                         `json:"ispublic"`
	IsPublicTopicAdminOnly             bool                         `json:"ispublictopicadminonly"`
	Motto                              string                       `json:"motto"`
	AllowChat                          bool                         `json:"allowchat"`
	IsDefaultPostPublic                bool                         `json:"isdefaultpostpublic"`
	ChatSecurity                       int                          `json:"chatsecurity"`
	Locale                             string                       `json:"locale"`
	AvatarImageIndex                   int                          `json:"avatarimageindex"`
	Homepage                           int                          `json:"homepage"`
	MembershipOption                   int                          `json:"membershipoption"`
	DefaultPublicity                   int                          `json:"defaultpublicity"`
	Theme                              string                       `json:"theme"`
	BannerPath                         string                       `json:"bannerpath"`
	AvatarPath                         string                       `json:"avatarpath"`
	ConversationId                     int                          `json:"conversationid"`
	EnableInvitationMessagingForAdmins bool                         `json:"enableinvitationmessagingforadmins"`
	BanExpireDate                      string                       `json:"banexpiredate"`
	Features                           GroupFeatures                `json:"features"`
	ClanInfo                           GroupV2ClanInfoAndInvestment `json:"claninfo"`
}
type GroupType struct {
	GroupType int `json:"grouptype"`
}
type ChatSecuritySetting struct {
	ChatSecuritySetting int `json:"chatsecuritysetting"`
}
type GroupHomepage struct {
	GroupHomepage int `json:"grouphomepage"`
}
type MembershipOption struct {
	MembershipOption int `json:"membershipoption"`
}
type GroupPostPublicity struct {
	GroupPostPublicity int `json:"grouppostpublicity"`
}
type GroupFeatures struct {
	MaximumMembers                   int   `json:"maximummembers"`
	MaximumMembershipsOfGroupType    int   `json:"maximummembershipsofgrouptype"`
	Capabilities                     int   `json:"capabilities"`
	MembershipTypes                  []int `json:"membershiptypes"`
	InvitePermissionOverride         bool  `json:"invitepermissionoverride"`
	UpdateCulturePermissionOverride  bool  `json:"updateculturepermissionoverride"`
	HostGuidedGamePermissionOverride int   `json:"hostguidedgamepermissionoverride"`
	UpdateBannerPermissionOverride   bool  `json:"updatebannerpermissionoverride"`
	JoinLevel                        int   `json:"joinlevel"`
}
type Capabilities struct {
	Capabilities int `json:"capabilities"`
}
type BungieMembershipTypeList struct {
	BungieMembershipTypeList int `json:"bungiemembershiptypelist"`
}
type HostGuidedGamesPermissionLevel struct {
	HostGuidedGamesPermissionLevel int `json:"hostguidedgamespermissionlevel"`
}
type RuntimeGroupMemberType struct {
	RuntimeGroupMemberType int `json:"runtimegroupmembertype"`
}
type GroupV2ClanInfo struct {
	ClanCallsign   string     `json:"clancallsign"`
	ClanBannerData ClanBanner `json:"clanbannerdata"`
}
type ClanBanner struct {
	DecalId                int `json:"decalid"`
	DecalColorId           int `json:"decalcolorid"`
	DecalBackgroundColorId int `json:"decalbackgroundcolorid"`
	GonfalonId             int `json:"gonfalonid"`
	GonfalonColorId        int `json:"gonfaloncolorid"`
	GonfalonDetailId       int `json:"gonfalondetailid"`
	GonfalonDetailColorId  int `json:"gonfalondetailcolorid"`
}
type GroupV2ClanInfoAndInvestment struct {
	D2ClanProgressions DestinyProgression `json:"d2clanprogressions"`
	ClanCallsign       string             `json:"clancallsign"`
	ClanBannerData     ClanBanner         `json:"clanbannerdata"`
}
type DestinyProgression struct {
	ProgressionHash     int                            `json:"progressionhash"`
	DailyProgress       int                            `json:"dailyprogress"`
	DailyLimit          int                            `json:"dailylimit"`
	WeeklyProgress      int                            `json:"weeklyprogress"`
	WeeklyLimit         int                            `json:"weeklylimit"`
	CurrentProgress     int                            `json:"currentprogress"`
	Level               int                            `json:"level"`
	LevelCap            int                            `json:"levelcap"`
	StepIndex           int                            `json:"stepindex"`
	ProgressToNextLevel int                            `json:"progresstonextlevel"`
	NextLevelAt         int                            `json:"nextlevelat"`
	CurrentResetCount   int                            `json:"currentresetcount"`
	SeasonResets        []DestinyProgressionResetEntry `json:"seasonresets"`
	RewardItemStates    []int                          `json:"rewarditemstates"`
}
type DestinyProgressionResetEntry struct {
	Season int `json:"season"`
	Resets int `json:"resets"`
}
type DestinyProgressionRewardItemState struct {
	DestinyProgressionRewardItemState int `json:"destinyprogressionrewarditemstate"`
}
type DestinyDefinition struct {
	Hash     int  `json:"hash"`
	Index    int  `json:"index"`
	Redacted bool `json:"redacted"`
}
type DestinyProgressionDefinition struct {
	DisplayProperties DestinyProgressionDisplayPropertiesDefinition `json:"displayproperties"`
	Scope             int                                           `json:"scope"`
	RepeatLastStep    bool                                          `json:"repeatlaststep"`
	Source            string                                        `json:"source"`
	Steps             []DestinyProgressionStepDefinition            `json:"steps"`
	Visible           bool                                          `json:"visible"`
	FactionHash       int                                           `json:"factionhash"`
	Color             DestinyColor                                  `json:"color"`
	RankIcon          string                                        `json:"rankicon"`
	RewardItems       []DestinyProgressionRewardItemQuantity        `json:"rewarditems"`
	Hash              int                                           `json:"hash"`
	Index             int                                           `json:"index"`
	Redacted          bool                                          `json:"redacted"`
}
type DestinyDisplayPropertiesDefinition struct {
	Description   string                          `json:"description"`
	Name          string                          `json:"name"`
	Icon          string                          `json:"icon"`
	IconSequences []DestinyIconSequenceDefinition `json:"iconsequences"`
	HighResIcon   string                          `json:"highresicon"`
	HasIcon       bool                            `json:"hasicon"`
}
type DestinyIconSequenceDefinition struct {
	Frames []string `json:"frames"`
}
type DestinyProgressionDisplayPropertiesDefinition struct {
	DisplayUnitsName string                          `json:"displayunitsname"`
	Description      string                          `json:"description"`
	Name             string                          `json:"name"`
	Icon             string                          `json:"icon"`
	IconSequences    []DestinyIconSequenceDefinition `json:"iconsequences"`
	HighResIcon      string                          `json:"highresicon"`
	HasIcon          bool                            `json:"hasicon"`
}
type DestinyProgressionScope struct {
	DestinyProgressionScope int `json:"destinyprogressionscope"`
}
type DestinyProgressionStepDefinition struct {
	StepName          string                `json:"stepname"`
	DisplayEffectType int                   `json:"displayeffecttype"`
	ProgressTotal     int                   `json:"progresstotal"`
	RewardItems       []DestinyItemQuantity `json:"rewarditems"`
	Icon              string                `json:"icon"`
}
type DestinyProgressionStepDisplayEffect struct {
	DestinyProgressionStepDisplayEffect int `json:"destinyprogressionstepdisplayeffect"`
}
type DestinyItemQuantity struct {
	ItemHash                 int  `json:"itemhash"`
	ItemInstanceId           int  `json:"iteminstanceid"`
	Quantity                 int  `json:"quantity"`
	HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyInventoryItemDefinition struct {
	DisplayProperties                 DestinyDisplayPropertiesDefinition    `json:"displayproperties"`
	TooltipNotifications              []DestinyItemTooltipNotification      `json:"tooltipnotifications"`
	CollectibleHash                   int                                   `json:"collectiblehash"`
	IconWatermark                     string                                `json:"iconwatermark"`
	IconWatermarkShelved              string                                `json:"iconwatermarkshelved"`
	SecondaryIcon                     string                                `json:"secondaryicon"`
	SecondaryOverlay                  string                                `json:"secondaryoverlay"`
	SecondarySpecial                  string                                `json:"secondaryspecial"`
	BackgroundColor                   DestinyColor                          `json:"backgroundcolor"`
	Screenshot                        string                                `json:"screenshot"`
	ItemTypeDisplayName               string                                `json:"itemtypedisplayname"`
	FlavorText                        string                                `json:"flavortext"`
	UiItemDisplayStyle                string                                `json:"uiitemdisplaystyle"`
	ItemTypeAndTierDisplayName        string                                `json:"itemtypeandtierdisplayname"`
	DisplaySource                     string                                `json:"displaysource"`
	TooltipStyle                      string                                `json:"tooltipstyle"`
	Action                            DestinyItemActionBlockDefinition      `json:"action"`
	Inventory                         DestinyItemInventoryBlockDefinition   `json:"inventory"`
	SetData                           DestinyItemSetBlockDefinition         `json:"setdata"`
	Stats                             DestinyItemStatBlockDefinition        `json:"stats"`
	EmblemObjectiveHash               int                                   `json:"emblemobjectivehash"`
	EquippingBlock                    DestinyEquippingBlockDefinition       `json:"equippingblock"`
	TranslationBlock                  DestinyItemTranslationBlockDefinition `json:"translationblock"`
	Preview                           DestinyItemPreviewBlockDefinition     `json:"preview"`
	Quality                           DestinyItemQualityBlockDefinition     `json:"quality"`
	Value                             DestinyItemValueBlockDefinition       `json:"value"`
	SourceData                        DestinyItemSourceBlockDefinition      `json:"sourcedata"`
	Objectives                        DestinyItemObjectiveBlockDefinition   `json:"objectives"`
	Metrics                           DestinyItemMetricBlockDefinition      `json:"metrics"`
	Plug                              DestinyItemPlugDefinition             `json:"plug"`
	Gearset                           DestinyItemGearsetBlockDefinition     `json:"gearset"`
	Sack                              DestinyItemSackBlockDefinition        `json:"sack"`
	Sockets                           DestinyItemSocketBlockDefinition      `json:"sockets"`
	Summary                           DestinyItemSummaryBlockDefinition     `json:"summary"`
	TalentGrid                        DestinyItemTalentGridBlockDefinition  `json:"talentgrid"`
	InvestmentStats                   []DestinyItemInvestmentStatDefinition `json:"investmentstats"`
	Perks                             []DestinyItemPerkEntryDefinition      `json:"perks"`
	LoreHash                          int                                   `json:"lorehash"`
	SummaryItemHash                   int                                   `json:"summaryitemhash"`
	Animations                        []DestinyAnimationReference           `json:"animations"`
	AllowActions                      bool                                  `json:"allowactions"`
	Links                             []HyperlinkReference                  `json:"links"`
	DoesPostmasterPullHaveSideEffects bool                                  `json:"doespostmasterpullhavesideeffects"`
	NonTransferrable                  bool                                  `json:"nontransferrable"`
	ItemCategoryHashes                []int                                 `json:"itemcategoryhashes"`
	SpecialItemType                   int                                   `json:"specialitemtype"`
	ItemType                          int                                   `json:"itemtype"`
	ItemSubType                       int                                   `json:"itemsubtype"`
	ClassType                         int                                   `json:"classtype"`
	BreakerType                       int                                   `json:"breakertype"`
	BreakerTypeHash                   int                                   `json:"breakertypehash"`
	Equippable                        bool                                  `json:"equippable"`
	DamageTypeHashes                  []int                                 `json:"damagetypehashes"`
	DamageTypes                       []int                                 `json:"damagetypes"`
	DefaultDamageType                 int                                   `json:"defaultdamagetype"`
	DefaultDamageTypeHash             int                                   `json:"defaultdamagetypehash"`
	SeasonHash                        int                                   `json:"seasonhash"`
	IsWrapper                         bool                                  `json:"iswrapper"`
	TraitIds                          []string                              `json:"traitids"`
	TraitHashes                       []int                                 `json:"traithashes"`
	Hash                              int                                   `json:"hash"`
	Index                             int                                   `json:"index"`
	Redacted                          bool                                  `json:"redacted"`
}
type DestinyItemTooltipNotification struct {
	DisplayString string `json:"displaystring"`
	DisplayStyle  string `json:"displaystyle"`
}
type DestinyColor struct {
	Red   string `json:"red"`
	Green string `json:"green"`
	Blue  string `json:"blue"`
	Alpha string `json:"alpha"`
}
type DestinyItemActionBlockDefinition struct {
	VerbName                string                                    `json:"verbname"`
	VerbDescription         string                                    `json:"verbdescription"`
	IsPositive              bool                                      `json:"ispositive"`
	OverlayScreenName       string                                    `json:"overlayscreenname"`
	OverlayIcon             string                                    `json:"overlayicon"`
	RequiredCooldownSeconds int                                       `json:"requiredcooldownseconds"`
	RequiredItems           []DestinyItemActionRequiredItemDefinition `json:"requireditems"`
	ProgressionRewards      []DestinyProgressionRewardDefinition      `json:"progressionrewards"`
	ActionTypeLabel         string                                    `json:"actiontypelabel"`
	RequiredLocation        string                                    `json:"requiredlocation"`
	RequiredCooldownHash    int                                       `json:"requiredcooldownhash"`
	DeleteOnAction          bool                                      `json:"deleteonaction"`
	ConsumeEntireStack      bool                                      `json:"consumeentirestack"`
	UseOnAcquire            bool                                      `json:"useonacquire"`
}
type DestinyItemActionRequiredItemDefinition struct {
	Count          int  `json:"count"`
	ItemHash       int  `json:"itemhash"`
	DeleteOnAction bool `json:"deleteonaction"`
}
type DestinyProgressionRewardDefinition struct {
	ProgressionMappingHash int  `json:"progressionmappinghash"`
	Amount                 int  `json:"amount"`
	ApplyThrottles         bool `json:"applythrottles"`
}
type DestinyProgressionMappingDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	DisplayUnits      string                             `json:"displayunits"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyItemInventoryBlockDefinition struct {
	StackUniqueLabel                         string `json:"stackuniquelabel"`
	MaxStackSize                             int    `json:"maxstacksize"`
	BucketTypeHash                           int    `json:"buckettypehash"`
	RecoveryBucketTypeHash                   int    `json:"recoverybuckettypehash"`
	TierTypeHash                             int    `json:"tiertypehash"`
	IsInstanceItem                           bool   `json:"isinstanceitem"`
	TierTypeName                             string `json:"tiertypename"`
	TierType                                 int    `json:"tiertype"`
	ExpirationTooltip                        string `json:"expirationtooltip"`
	ExpiredInActivityMessage                 string `json:"expiredinactivitymessage"`
	ExpiredInOrbitMessage                    string `json:"expiredinorbitmessage"`
	SuppressExpirationWhenObjectivesComplete bool   `json:"suppressexpirationwhenobjectivescomplete"`
}
type TierType struct {
	TierType int `json:"tiertype"`
}
type DestinyInventoryBucketDefinition struct {
	DisplayProperties      DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Scope                  int                                `json:"scope"`
	Category               int                                `json:"category"`
	BucketOrder            int                                `json:"bucketorder"`
	ItemCount              int                                `json:"itemcount"`
	Location               int                                `json:"location"`
	HasTransferDestination bool                               `json:"hastransferdestination"`
	Enabled                bool                               `json:"enabled"`
	Fifo                   bool                               `json:"fifo"`
	Hash                   int                                `json:"hash"`
	Index                  int                                `json:"index"`
	Redacted               bool                               `json:"redacted"`
}
type BucketScope struct {
	BucketScope int `json:"bucketscope"`
}
type BucketCategory struct {
	BucketCategory int `json:"bucketcategory"`
}
type ItemLocation struct {
	ItemLocation int `json:"itemlocation"`
}
type DestinyItemTierTypeDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	InfusionProcess   DestinyItemTierTypeInfusionBlock   `json:"infusionprocess"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyItemTierTypeInfusionBlock struct {
	BaseQualityTransferRatio float64 `json:"basequalitytransferratio"`
	MinimumQualityIncrement  int     `json:"minimumqualityincrement"`
}
type DestinyItemSetBlockDefinition struct {
	ItemList                 []DestinyItemSetBlockEntryDefinition `json:"itemlist"`
	RequireOrderedSetItemAdd bool                                 `json:"requireorderedsetitemadd"`
	SetIsFeatured            bool                                 `json:"setisfeatured"`
	SetType                  string                               `json:"settype"`
	QuestLineName            string                               `json:"questlinename"`
	QuestLineDescription     string                               `json:"questlinedescription"`
	QuestStepSummary         string                               `json:"queststepsummary"`
}
type DestinyItemSetBlockEntryDefinition struct {
	TrackingValue int `json:"trackingvalue"`
	ItemHash      int `json:"itemhash"`
}
type DestinyItemStatBlockDefinition struct {
	DisablePrimaryStatDisplay bool                               `json:"disableprimarystatdisplay"`
	StatGroupHash             int                                `json:"statgrouphash"`
	Stats                     DestinyInventoryItemStatDefinition `json:"stats"`
	HasDisplayableStats       bool                               `json:"hasdisplayablestats"`
	PrimaryBaseStatHash       int                                `json:"primarybasestathash"`
}
type DestinyInventoryItemStatDefinition struct {
	StatHash       int `json:"stathash"`
	Value          int `json:"value"`
	Minimum        int `json:"minimum"`
	Maximum        int `json:"maximum"`
	DisplayMaximum int `json:"displaymaximum"`
}
type DestinyStatDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	AggregationType   int                                `json:"aggregationtype"`
	HasComputedBlock  bool                               `json:"hascomputedblock"`
	StatCategory      int                                `json:"statcategory"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyStatAggregationType struct {
	DestinyStatAggregationType int `json:"destinystataggregationtype"`
}
type DestinyStatCategory struct {
	DestinyStatCategory int `json:"destinystatcategory"`
}
type DestinyStatGroupDefinition struct {
	MaximumValue int                            `json:"maximumvalue"`
	UiPosition   int                            `json:"uiposition"`
	ScaledStats  []DestinyStatDisplayDefinition `json:"scaledstats"`
	Overrides    DestinyStatOverrideDefinition  `json:"overrides"`
	Hash         int                            `json:"hash"`
	Index        int                            `json:"index"`
	Redacted     bool                           `json:"redacted"`
}
type DestinyStatDisplayDefinition struct {
	StatHash             int                  `json:"stathash"`
	MaximumValue         int                  `json:"maximumvalue"`
	DisplayAsNumeric     bool                 `json:"displayasnumeric"`
	DisplayInterpolation []InterpolationPoint `json:"displayinterpolation"`
}
type InterpolationPoint struct {
	Value  int `json:"value"`
	Weight int `json:"weight"`
}
type DestinyStatOverrideDefinition struct {
	StatHash          int                                `json:"stathash"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyEquippingBlockDefinition struct {
	GearsetItemHash       int      `json:"gearsetitemhash"`
	UniqueLabel           string   `json:"uniquelabel"`
	UniqueLabelHash       int      `json:"uniquelabelhash"`
	EquipmentSlotTypeHash int      `json:"equipmentslottypehash"`
	Attributes            int      `json:"attributes"`
	AmmoType              int      `json:"ammotype"`
	DisplayStrings        []string `json:"displaystrings"`
}
type EquippingItemBlockAttributes struct {
	EquippingItemBlockAttributes int `json:"equippingitemblockattributes"`
}
type DestinyAmmunitionType struct {
	DestinyAmmunitionType int `json:"destinyammunitiontype"`
}
type DestinyEquipmentSlotDefinition struct {
	DisplayProperties     DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	EquipmentCategoryHash int                                `json:"equipmentcategoryhash"`
	BucketTypeHash        int                                `json:"buckettypehash"`
	ApplyCustomArtDyes    bool                               `json:"applycustomartdyes"`
	ArtDyeChannels        []DestinyArtDyeReference           `json:"artdyechannels"`
	Hash                  int                                `json:"hash"`
	Index                 int                                `json:"index"`
	Redacted              bool                               `json:"redacted"`
}
type DestinyArtDyeReference struct {
	ArtDyeChannelHash int `json:"artdyechannelhash"`
}
type DestinyItemTranslationBlockDefinition struct {
	WeaponPatternIdentifier string                               `json:"weaponpatternidentifier"`
	WeaponPatternHash       int                                  `json:"weaponpatternhash"`
	DefaultDyes             []DyeReference                       `json:"defaultdyes"`
	LockedDyes              []DyeReference                       `json:"lockeddyes"`
	CustomDyes              []DyeReference                       `json:"customdyes"`
	Arrangements            []DestinyGearArtArrangementReference `json:"arrangements"`
	HasGeometry             bool                                 `json:"hasgeometry"`
}
type DyeReference struct {
	ChannelHash int `json:"channelhash"`
	DyeHash     int `json:"dyehash"`
}
type DestinyGearArtArrangementReference struct {
	ClassHash          int `json:"classhash"`
	ArtArrangementHash int `json:"artarrangementhash"`
}
type DestinyItemPreviewBlockDefinition struct {
	ScreenStyle           string                                 `json:"screenstyle"`
	PreviewVendorHash     int                                    `json:"previewvendorhash"`
	ArtifactHash          int                                    `json:"artifacthash"`
	PreviewActionString   string                                 `json:"previewactionstring"`
	DerivedItemCategories []DestinyDerivedItemCategoryDefinition `json:"deriveditemcategories"`
}
type DestinyDerivedItemCategoryDefinition struct {
	CategoryDescription string                         `json:"categorydescription"`
	Items               []DestinyDerivedItemDefinition `json:"items"`
}
type DestinyDerivedItemDefinition struct {
	ItemHash        int    `json:"itemhash"`
	ItemName        string `json:"itemname"`
	ItemDetail      string `json:"itemdetail"`
	ItemDescription string `json:"itemdescription"`
	IconPath        string `json:"iconpath"`
	VendorItemIndex int    `json:"vendoritemindex"`
}
type DestinyVendorDefinition struct {
	DisplayProperties           DestinyVendorDisplayPropertiesDefinition `json:"displayproperties"`
	VendorProgressionType       int                                      `json:"vendorprogressiontype"`
	BuyString                   string                                   `json:"buystring"`
	SellString                  string                                   `json:"sellstring"`
	DisplayItemHash             int                                      `json:"displayitemhash"`
	InhibitBuying               bool                                     `json:"inhibitbuying"`
	InhibitSelling              bool                                     `json:"inhibitselling"`
	FactionHash                 int                                      `json:"factionhash"`
	ResetIntervalMinutes        int                                      `json:"resetintervalminutes"`
	ResetOffsetMinutes          int                                      `json:"resetoffsetminutes"`
	FailureStrings              []string                                 `json:"failurestrings"`
	UnlockRanges                []DateRange                              `json:"unlockranges"`
	VendorIdentifier            string                                   `json:"vendoridentifier"`
	VendorPortrait              string                                   `json:"vendorportrait"`
	VendorBanner                string                                   `json:"vendorbanner"`
	Enabled                     bool                                     `json:"enabled"`
	Visible                     bool                                     `json:"visible"`
	VendorSubcategoryIdentifier string                                   `json:"vendorsubcategoryidentifier"`
	ConsolidateCategories       bool                                     `json:"consolidatecategories"`
	Actions                     []DestinyVendorActionDefinition          `json:"actions"`
	Categories                  []DestinyVendorCategoryEntryDefinition   `json:"categories"`
	OriginalCategories          []DestinyVendorCategoryEntryDefinition   `json:"originalcategories"`
	DisplayCategories           []DestinyDisplayCategoryDefinition       `json:"displaycategories"`
	Interactions                []DestinyVendorInteractionDefinition     `json:"interactions"`
	InventoryFlyouts            []DestinyVendorInventoryFlyoutDefinition `json:"inventoryflyouts"`
	ItemList                    []DestinyVendorItemDefinition            `json:"itemlist"`
	Services                    []DestinyVendorServiceDefinition         `json:"services"`
	AcceptedItems               []DestinyVendorAcceptedItemDefinition    `json:"accepteditems"`
	ReturnWithVendorRequest     bool                                     `json:"returnwithvendorrequest"`
	Locations                   []DestinyVendorLocationDefinition        `json:"locations"`
	Groups                      []DestinyVendorGroupReference            `json:"groups"`
	IgnoreSaleItemHashes        []int                                    `json:"ignoresaleitemhashes"`
	Hash                        int                                      `json:"hash"`
	Index                       int                                      `json:"index"`
	Redacted                    bool                                     `json:"redacted"`
}
type DestinyVendorDisplayPropertiesDefinition struct {
	LargeIcon            string                                           `json:"largeicon"`
	Subtitle             string                                           `json:"subtitle"`
	OriginalIcon         string                                           `json:"originalicon"`
	RequirementsDisplay  []DestinyVendorRequirementDisplayEntryDefinition `json:"requirementsdisplay"`
	SmallTransparentIcon string                                           `json:"smalltransparenticon"`
	MapIcon              string                                           `json:"mapicon"`
	LargeTransparentIcon string                                           `json:"largetransparenticon"`
	Description          string                                           `json:"description"`
	Name                 string                                           `json:"name"`
	Icon                 string                                           `json:"icon"`
	IconSequences        []DestinyIconSequenceDefinition                  `json:"iconsequences"`
	HighResIcon          string                                           `json:"highresicon"`
	HasIcon              bool                                             `json:"hasicon"`
}
type DestinyVendorRequirementDisplayEntryDefinition struct {
	Icon   string `json:"icon"`
	Name   string `json:"name"`
	Source string `json:"source"`
	Type   string `json:"type"`
}
type DestinyVendorProgressionType struct {
	DestinyVendorProgressionType int `json:"destinyvendorprogressiontype"`
}
type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
type DestinyVendorActionDefinition struct {
	Description       string `json:"description"`
	ExecuteSeconds    int    `json:"executeseconds"`
	Icon              string `json:"icon"`
	Name              string `json:"name"`
	Verb              string `json:"verb"`
	IsPositive        bool   `json:"ispositive"`
	ActionId          string `json:"actionid"`
	ActionHash        int    `json:"actionhash"`
	AutoPerformAction bool   `json:"autoperformaction"`
}
type DestinyVendorCategoryEntryDefinition struct {
	CategoryIndex                int                                    `json:"categoryindex"`
	SortValue                    int                                    `json:"sortvalue"`
	CategoryHash                 int                                    `json:"categoryhash"`
	QuantityAvailable            int                                    `json:"quantityavailable"`
	ShowUnavailableItems         bool                                   `json:"showunavailableitems"`
	HideIfNoCurrency             bool                                   `json:"hideifnocurrency"`
	HideFromRegularPurchase      bool                                   `json:"hidefromregularpurchase"`
	BuyStringOverride            string                                 `json:"buystringoverride"`
	DisabledDescription          string                                 `json:"disableddescription"`
	DisplayTitle                 string                                 `json:"displaytitle"`
	Overlay                      DestinyVendorCategoryOverlayDefinition `json:"overlay"`
	VendorItemIndexes            []int                                  `json:"vendoritemindexes"`
	IsPreview                    bool                                   `json:"ispreview"`
	IsDisplayOnly                bool                                   `json:"isdisplayonly"`
	ResetIntervalMinutesOverride int                                    `json:"resetintervalminutesoverride"`
	ResetOffsetMinutesOverride   int                                    `json:"resetoffsetminutesoverride"`
}
type DestinyVendorCategoryOverlayDefinition struct {
	ChoiceDescription string `json:"choicedescription"`
	Description       string `json:"description"`
	Icon              string `json:"icon"`
	Title             string `json:"title"`
	CurrencyItemHash  int    `json:"currencyitemhash"`
}
type DestinyDisplayCategoryDefinition struct {
	Index                  int                                `json:"index"`
	Identifier             string                             `json:"identifier"`
	DisplayCategoryHash    int                                `json:"displaycategoryhash"`
	DisplayProperties      DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	DisplayInBanner        bool                               `json:"displayinbanner"`
	ProgressionHash        int                                `json:"progressionhash"`
	SortOrder              int                                `json:"sortorder"`
	DisplayStyleHash       int                                `json:"displaystylehash"`
	DisplayStyleIdentifier string                             `json:"displaystyleidentifier"`
}
type VendorDisplayCategorySortOrder struct {
	VendorDisplayCategorySortOrder int `json:"vendordisplaycategorysortorder"`
}
type DestinyVendorInteractionDefinition struct {
	InteractionIndex          int                                           `json:"interactionindex"`
	Replies                   []DestinyVendorInteractionReplyDefinition     `json:"replies"`
	VendorCategoryIndex       int                                           `json:"vendorcategoryindex"`
	QuestlineItemHash         int                                           `json:"questlineitemhash"`
	SackInteractionList       []DestinyVendorInteractionSackEntryDefinition `json:"sackinteractionlist"`
	UiInteractionType         int                                           `json:"uiinteractiontype"`
	InteractionType           int                                           `json:"interactiontype"`
	RewardBlockLabel          string                                        `json:"rewardblocklabel"`
	RewardVendorCategoryIndex int                                           `json:"rewardvendorcategoryindex"`
	FlavorLineOne             string                                        `json:"flavorlineone"`
	FlavorLineTwo             string                                        `json:"flavorlinetwo"`
	HeaderDisplayProperties   DestinyDisplayPropertiesDefinition            `json:"headerdisplayproperties"`
	Instructions              string                                        `json:"instructions"`
}
type DestinyVendorInteractionReplyDefinition struct {
	ItemRewardsSelection int    `json:"itemrewardsselection"`
	Reply                string `json:"reply"`
	ReplyType            int    `json:"replytype"`
}
type DestinyVendorInteractionRewardSelection struct {
	DestinyVendorInteractionRewardSelection int `json:"destinyvendorinteractionrewardselection"`
}
type DestinyVendorReplyType struct {
	DestinyVendorReplyType int `json:"destinyvendorreplytype"`
}
type DestinyVendorInteractionSackEntryDefinition struct {
	SackType int `json:"sacktype"`
}
type VendorInteractionType struct {
	VendorInteractionType int `json:"vendorinteractiontype"`
}
type DestinyVendorInventoryFlyoutDefinition struct {
	LockedDescription string                                         `json:"lockeddescription"`
	DisplayProperties DestinyDisplayPropertiesDefinition             `json:"displayproperties"`
	Buckets           []DestinyVendorInventoryFlyoutBucketDefinition `json:"buckets"`
	FlyoutId          int                                            `json:"flyoutid"`
	SuppressNewness   bool                                           `json:"suppressnewness"`
	EquipmentSlotHash int                                            `json:"equipmentslothash"`
}
type DestinyVendorInventoryFlyoutBucketDefinition struct {
	Collapsible         bool `json:"collapsible"`
	InventoryBucketHash int  `json:"inventorybuckethash"`
	SortItemsBy         int  `json:"sortitemsby"`
}
type DestinyItemSortType struct {
	DestinyItemSortType int `json:"destinyitemsorttype"`
}
type DestinyVendorItemDefinition struct {
	VendorItemIndex       int                                        `json:"vendoritemindex"`
	ItemHash              int                                        `json:"itemhash"`
	Quantity              int                                        `json:"quantity"`
	FailureIndexes        []int                                      `json:"failureindexes"`
	Currencies            []DestinyVendorItemQuantity                `json:"currencies"`
	RefundPolicy          int                                        `json:"refundpolicy"`
	RefundTimeLimit       int                                        `json:"refundtimelimit"`
	CreationLevels        []DestinyItemCreationEntryLevelDefinition  `json:"creationlevels"`
	DisplayCategoryIndex  int                                        `json:"displaycategoryindex"`
	CategoryIndex         int                                        `json:"categoryindex"`
	OriginalCategoryIndex int                                        `json:"originalcategoryindex"`
	MinimumLevel          int                                        `json:"minimumlevel"`
	MaximumLevel          int                                        `json:"maximumlevel"`
	Action                DestinyVendorSaleItemActionBlockDefinition `json:"action"`
	DisplayCategory       string                                     `json:"displaycategory"`
	InventoryBucketHash   int                                        `json:"inventorybuckethash"`
	VisibilityScope       int                                        `json:"visibilityscope"`
	PurchasableScope      int                                        `json:"purchasablescope"`
	Exclusivity           int                                        `json:"exclusivity"`
	IsOffer               bool                                       `json:"isoffer"`
	IsCrm                 bool                                       `json:"iscrm"`
	SortValue             int                                        `json:"sortvalue"`
	ExpirationTooltip     string                                     `json:"expirationtooltip"`
	RedirectToSaleIndexes []int                                      `json:"redirecttosaleindexes"`
	SocketOverrides       []DestinyVendorItemSocketOverride          `json:"socketoverrides"`
	Unpurchasable         bool                                       `json:"unpurchasable"`
}
type DestinyVendorItemQuantity struct {
	ItemHash                 int  `json:"itemhash"`
	ItemInstanceId           int  `json:"iteminstanceid"`
	Quantity                 int  `json:"quantity"`
	HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyVendorItemRefundPolicy struct {
	DestinyVendorItemRefundPolicy int `json:"destinyvendoritemrefundpolicy"`
}
type DestinyItemCreationEntryLevelDefinition struct {
	Level int `json:"level"`
}
type DestinyVendorSaleItemActionBlockDefinition struct {
	ExecuteSeconds float64 `json:"executeseconds"`
	IsPositive     bool    `json:"ispositive"`
}
type DestinyGatingScope struct {
	DestinyGatingScope int `json:"destinygatingscope"`
}
type DestinyVendorItemSocketOverride struct {
	SingleItemHash         int `json:"singleitemhash"`
	RandomizedOptionsCount int `json:"randomizedoptionscount"`
	SocketTypeHash         int `json:"sockettypehash"`
}
type DestinySocketTypeDefinition struct {
	DisplayProperties               DestinyDisplayPropertiesDefinition                `json:"displayproperties"`
	InsertAction                    DestinyInsertPlugActionDefinition                 `json:"insertaction"`
	PlugWhitelist                   []DestinyPlugWhitelistEntryDefinition             `json:"plugwhitelist"`
	SocketCategoryHash              int                                               `json:"socketcategoryhash"`
	Visibility                      int                                               `json:"visibility"`
	AlwaysRandomizeSockets          bool                                              `json:"alwaysrandomizesockets"`
	IsPreviewEnabled                bool                                              `json:"ispreviewenabled"`
	HideDuplicateReusablePlugs      bool                                              `json:"hideduplicatereusableplugs"`
	OverridesUiAppearance           bool                                              `json:"overridesuiappearance"`
	AvoidDuplicatesOnInitialization bool                                              `json:"avoidduplicatesoninitialization"`
	CurrencyScalars                 []DestinySocketTypeScalarMaterialRequirementEntry `json:"currencyscalars"`
	Hash                            int                                               `json:"hash"`
	Index                           int                                               `json:"index"`
	Redacted                        bool                                              `json:"redacted"`
}
type DestinyInsertPlugActionDefinition struct {
	ActionExecuteSeconds int `json:"actionexecuteseconds"`
	ActionType           int `json:"actiontype"`
}
type SocketTypeActionType struct {
	SocketTypeActionType int `json:"sockettypeactiontype"`
}
type DestinyPlugWhitelistEntryDefinition struct {
	CategoryHash                       int    `json:"categoryhash"`
	CategoryIdentifier                 string `json:"categoryidentifier"`
	ReinitializationPossiblePlugHashes []int  `json:"reinitializationpossibleplughashes"`
}
type DestinySocketVisibility struct {
	DestinySocketVisibility int `json:"destinysocketvisibility"`
}
type DestinySocketTypeScalarMaterialRequirementEntry struct {
	CurrencyItemHash int `json:"currencyitemhash"`
	ScalarValue      int `json:"scalarvalue"`
}
type DestinySocketCategoryDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	UiCategoryStyle   int                                `json:"uicategorystyle"`
	CategoryStyle     int                                `json:"categorystyle"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinySocketCategoryStyle struct {
	DestinySocketCategoryStyle int `json:"destinysocketcategorystyle"`
}
type DestinyVendorServiceDefinition struct {
	Name string `json:"name"`
}
type DestinyVendorAcceptedItemDefinition struct {
	AcceptedInventoryBucketHash    int `json:"acceptedinventorybuckethash"`
	DestinationInventoryBucketHash int `json:"destinationinventorybuckethash"`
}
type DestinyVendorLocationDefinition struct {
	DestinationHash     int    `json:"destinationhash"`
	BackgroundImagePath string `json:"backgroundimagepath"`
}
type DestinyDestinationDefinition struct {
	DisplayProperties           DestinyDisplayPropertiesDefinition          `json:"displayproperties"`
	PlaceHash                   int                                         `json:"placehash"`
	DefaultFreeroamActivityHash int                                         `json:"defaultfreeroamactivityhash"`
	ActivityGraphEntries        []DestinyActivityGraphListEntryDefinition   `json:"activitygraphentries"`
	BubbleSettings              []DestinyDestinationBubbleSettingDefinition `json:"bubblesettings"`
	Bubbles                     []DestinyBubbleDefinition                   `json:"bubbles"`
	Hash                        int                                         `json:"hash"`
	Index                       int                                         `json:"index"`
	Redacted                    bool                                        `json:"redacted"`
}
type DestinyActivityGraphListEntryDefinition struct {
	ActivityGraphHash int `json:"activitygraphhash"`
}
type DestinyActivityGraphDefinition struct {
	Nodes               []DestinyActivityGraphNodeDefinition               `json:"nodes"`
	ArtElements         []DestinyActivityGraphArtElementDefinition         `json:"artelements"`
	Connections         []DestinyActivityGraphConnectionDefinition         `json:"connections"`
	DisplayObjectives   []DestinyActivityGraphDisplayObjectiveDefinition   `json:"displayobjectives"`
	DisplayProgressions []DestinyActivityGraphDisplayProgressionDefinition `json:"displayprogressions"`
	LinkedGraphs        []DestinyLinkedGraphDefinition                     `json:"linkedgraphs"`
	Hash                int                                                `json:"hash"`
	Index               int                                                `json:"index"`
	Redacted            bool                                               `json:"redacted"`
}
type DestinyActivityGraphNodeDefinition struct {
	NodeId          int                                                `json:"nodeid"`
	OverrideDisplay DestinyDisplayPropertiesDefinition                 `json:"overridedisplay"`
	Position        DestinyPositionDefinition                          `json:"position"`
	FeaturingStates []DestinyActivityGraphNodeFeaturingStateDefinition `json:"featuringstates"`
	Activities      []DestinyActivityGraphNodeActivityDefinition       `json:"activities"`
	States          []DestinyActivityGraphNodeStateEntry               `json:"states"`
}
type DestinyPositionDefinition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}
type DestinyActivityGraphNodeFeaturingStateDefinition struct {
	HighlightType int `json:"highlighttype"`
}
type ActivityGraphNodeHighlightType struct {
	ActivityGraphNodeHighlightType int `json:"activitygraphnodehighlighttype"`
}
type DestinyActivityGraphNodeActivityDefinition struct {
	NodeActivityId int `json:"nodeactivityid"`
	ActivityHash   int `json:"activityhash"`
}
type DestinyActivityDefinition struct {
	DisplayProperties                DestinyDisplayPropertiesDefinition           `json:"displayproperties"`
	OriginalDisplayProperties        DestinyDisplayPropertiesDefinition           `json:"originaldisplayproperties"`
	SelectionScreenDisplayProperties DestinyDisplayPropertiesDefinition           `json:"selectionscreendisplayproperties"`
	ReleaseIcon                      string                                       `json:"releaseicon"`
	ReleaseTime                      int                                          `json:"releasetime"`
	ActivityLightLevel               int                                          `json:"activitylightlevel"`
	DestinationHash                  int                                          `json:"destinationhash"`
	PlaceHash                        int                                          `json:"placehash"`
	ActivityTypeHash                 int                                          `json:"activitytypehash"`
	Tier                             int                                          `json:"tier"`
	PgcrImage                        string                                       `json:"pgcrimage"`
	Rewards                          []DestinyActivityRewardDefinition            `json:"rewards"`
	Modifiers                        []DestinyActivityModifierReferenceDefinition `json:"modifiers"`
	IsPlaylist                       bool                                         `json:"isplaylist"`
	Challenges                       []DestinyActivityChallengeDefinition         `json:"challenges"`
	OptionalUnlockStrings            []DestinyActivityUnlockStringDefinition      `json:"optionalunlockstrings"`
	PlaylistItems                    []DestinyActivityPlaylistItemDefinition      `json:"playlistitems"`
	ActivityGraphList                []DestinyActivityGraphListEntryDefinition    `json:"activitygraphlist"`
	Matchmaking                      DestinyActivityMatchmakingBlockDefinition    `json:"matchmaking"`
	GuidedGame                       DestinyActivityGuidedBlockDefinition         `json:"guidedgame"`
	DirectActivityModeHash           int                                          `json:"directactivitymodehash"`
	DirectActivityModeType           int                                          `json:"directactivitymodetype"`
	Loadouts                         []DestinyActivityLoadoutRequirementSet       `json:"loadouts"`
	ActivityModeHashes               []int                                        `json:"activitymodehashes"`
	ActivityModeTypes                []int                                        `json:"activitymodetypes"`
	IsPvP                            bool                                         `json:"ispvp"`
	InsertionPoints                  []DestinyActivityInsertionPointDefinition    `json:"insertionpoints"`
	ActivityLocationMappings         []DestinyEnvironmentLocationMapping          `json:"activitylocationmappings"`
	Hash                             int                                          `json:"hash"`
	Index                            int                                          `json:"index"`
	Redacted                         bool                                         `json:"redacted"`
}
type DestinyActivityRewardDefinition struct {
	RewardText  string                `json:"rewardtext"`
	RewardItems []DestinyItemQuantity `json:"rewarditems"`
}
type DestinyActivityModifierReferenceDefinition struct {
	ActivityModifierHash int `json:"activitymodifierhash"`
}
type DestinyActivityModifierDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyActivityChallengeDefinition struct {
	ObjectiveHash int                   `json:"objectivehash"`
	DummyRewards  []DestinyItemQuantity `json:"dummyrewards"`
}
type DestinyObjectiveDefinition struct {
	DisplayProperties             DestinyDisplayPropertiesDefinition  `json:"displayproperties"`
	CompletionValue               int                                 `json:"completionvalue"`
	Scope                         int                                 `json:"scope"`
	LocationHash                  int                                 `json:"locationhash"`
	AllowNegativeValue            bool                                `json:"allownegativevalue"`
	AllowValueChangeWhenCompleted bool                                `json:"allowvaluechangewhencompleted"`
	IsCountingDownward            bool                                `json:"iscountingdownward"`
	ValueStyle                    int                                 `json:"valuestyle"`
	ProgressDescription           string                              `json:"progressdescription"`
	Perks                         DestinyObjectivePerkEntryDefinition `json:"perks"`
	Stats                         DestinyObjectiveStatEntryDefinition `json:"stats"`
	MinimumVisibilityThreshold    int                                 `json:"minimumvisibilitythreshold"`
	AllowOvercompletion           bool                                `json:"allowovercompletion"`
	ShowValueOnComplete           bool                                `json:"showvalueoncomplete"`
	CompletedValueStyle           int                                 `json:"completedvaluestyle"`
	InProgressValueStyle          int                                 `json:"inprogressvaluestyle"`
	Hash                          int                                 `json:"hash"`
	Index                         int                                 `json:"index"`
	Redacted                      bool                                `json:"redacted"`
}
type DestinyUnlockValueUIStyle struct {
	DestinyUnlockValueUIStyle int `json:"destinyunlockvalueuistyle"`
}
type DestinyObjectivePerkEntryDefinition struct {
	PerkHash int `json:"perkhash"`
	Style    int `json:"style"`
}
type DestinyObjectiveGrantStyle struct {
	DestinyObjectiveGrantStyle int `json:"destinyobjectivegrantstyle"`
}
type DestinySandboxPerkDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	PerkIdentifier    string                             `json:"perkidentifier"`
	IsDisplayable     bool                               `json:"isdisplayable"`
	DamageType        int                                `json:"damagetype"`
	DamageTypeHash    int                                `json:"damagetypehash"`
	PerkGroups        DestinyTalentNodeStepGroups        `json:"perkgroups"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DamageType struct {
	DamageType int `json:"damagetype"`
}
type DestinyTalentNodeStepGroups struct {
	WeaponPerformance  int `json:"weaponperformance"`
	ImpactEffects      int `json:"impacteffects"`
	GuardianAttributes int `json:"guardianattributes"`
	LightAbilities     int `json:"lightabilities"`
	DamageTypes        int `json:"damagetypes"`
}
type DestinyTalentNodeStepWeaponPerformances struct {
	DestinyTalentNodeStepWeaponPerformances int `json:"destinytalentnodestepweaponperformances"`
}
type DestinyTalentNodeStepImpactEffects struct {
	DestinyTalentNodeStepImpactEffects int `json:"destinytalentnodestepimpacteffects"`
}
type DestinyTalentNodeStepGuardianAttributes struct {
	DestinyTalentNodeStepGuardianAttributes int `json:"destinytalentnodestepguardianattributes"`
}
type DestinyTalentNodeStepLightAbilities struct {
	DestinyTalentNodeStepLightAbilities int `json:"destinytalentnodesteplightabilities"`
}
type DestinyTalentNodeStepDamageTypes struct {
	DestinyTalentNodeStepDamageTypes int `json:"destinytalentnodestepdamagetypes"`
}
type DestinyObjectiveStatEntryDefinition struct {
	Stat  DestinyItemInvestmentStatDefinition `json:"stat"`
	Style int                                 `json:"style"`
}
type DestinyItemInvestmentStatDefinition struct {
	StatTypeHash          int  `json:"stattypehash"`
	Value                 int  `json:"value"`
	IsConditionallyActive bool `json:"isconditionallyactive"`
}
type DestinyLocationDefinition struct {
	VendorHash       int                                `json:"vendorhash"`
	LocationReleases []DestinyLocationReleaseDefinition `json:"locationreleases"`
	Hash             int                                `json:"hash"`
	Index            int                                `json:"index"`
	Redacted         bool                               `json:"redacted"`
}
type DestinyLocationReleaseDefinition struct {
	DisplayProperties       DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	SmallTransparentIcon    string                             `json:"smalltransparenticon"`
	MapIcon                 string                             `json:"mapicon"`
	LargeTransparentIcon    string                             `json:"largetransparenticon"`
	SpawnPoint              int                                `json:"spawnpoint"`
	DestinationHash         int                                `json:"destinationhash"`
	ActivityHash            int                                `json:"activityhash"`
	ActivityGraphHash       int                                `json:"activitygraphhash"`
	ActivityGraphNodeHash   int                                `json:"activitygraphnodehash"`
	ActivityBubbleName      int                                `json:"activitybubblename"`
	ActivityPathBundle      int                                `json:"activitypathbundle"`
	ActivityPathDestination int                                `json:"activitypathdestination"`
	NavPointType            int                                `json:"navpointtype"`
	WorldPosition           []int                              `json:"worldposition"`
}
type DestinyActivityNavPointType struct {
	DestinyActivityNavPointType int `json:"destinyactivitynavpointtype"`
}
type DestinyActivityUnlockStringDefinition struct {
	DisplayString string `json:"displaystring"`
}
type DestinyActivityPlaylistItemDefinition struct {
	ActivityHash           int   `json:"activityhash"`
	DirectActivityModeHash int   `json:"directactivitymodehash"`
	DirectActivityModeType int   `json:"directactivitymodetype"`
	ActivityModeHashes     []int `json:"activitymodehashes"`
	ActivityModeTypes      []int `json:"activitymodetypes"`
}
type DestinyActivityModeType struct {
	DestinyActivityModeType int `json:"destinyactivitymodetype"`
}
type DestinyActivityModeDefinition struct {
	DisplayProperties    DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	PgcrImage            string                             `json:"pgcrimage"`
	ModeType             int                                `json:"modetype"`
	ActivityModeCategory int                                `json:"activitymodecategory"`
	IsTeamBased          bool                               `json:"isteambased"`
	IsAggregateMode      bool                               `json:"isaggregatemode"`
	ParentHashes         []int                              `json:"parenthashes"`
	FriendlyName         string                             `json:"friendlyname"`
	ActivityModeMappings int                                `json:"activitymodemappings"`
	Display              bool                               `json:"display"`
	Order                int                                `json:"order"`
	Hash                 int                                `json:"hash"`
	Index                int                                `json:"index"`
	Redacted             bool                               `json:"redacted"`
}
type DestinyActivityModeCategory struct {
	DestinyActivityModeCategory int `json:"destinyactivitymodecategory"`
}
type DestinyActivityMatchmakingBlockDefinition struct {
	IsMatchmade          bool `json:"ismatchmade"`
	MinParty             int  `json:"minparty"`
	MaxParty             int  `json:"maxparty"`
	MaxPlayers           int  `json:"maxplayers"`
	RequiresGuardianOath bool `json:"requiresguardianoath"`
}
type DestinyActivityGuidedBlockDefinition struct {
	GuidedMaxLobbySize int `json:"guidedmaxlobbysize"`
	GuidedMinLobbySize int `json:"guidedminlobbysize"`
	GuidedDisbandCount int `json:"guideddisbandcount"`
}
type DestinyActivityLoadoutRequirementSet struct {
	Requirements []DestinyActivityLoadoutRequirement `json:"requirements"`
}
type DestinyActivityLoadoutRequirement struct {
	EquipmentSlotHash         int   `json:"equipmentslothash"`
	AllowedEquippedItemHashes []int `json:"allowedequippeditemhashes"`
	AllowedWeaponSubTypes     []int `json:"allowedweaponsubtypes"`
}
type DestinyItemSubType struct {
	DestinyItemSubType int `json:"destinyitemsubtype"`
}
type DestinyActivityInsertionPointDefinition struct {
	PhaseHash int `json:"phasehash"`
}
type DestinyEnvironmentLocationMapping struct {
	LocationHash     int    `json:"locationhash"`
	ActivationSource string `json:"activationsource"`
	ItemHash         int    `json:"itemhash"`
	ObjectiveHash    int    `json:"objectivehash"`
	ActivityHash     int    `json:"activityhash"`
}
type DestinyPlaceDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyActivityTypeDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyActivityGraphNodeStateEntry struct {
	State int `json:"state"`
}
type DestinyGraphNodeState struct {
	DestinyGraphNodeState int `json:"destinygraphnodestate"`
}
type DestinyActivityGraphArtElementDefinition struct {
	Position DestinyPositionDefinition `json:"position"`
}
type DestinyActivityGraphConnectionDefinition struct {
	SourceNodeHash int `json:"sourcenodehash"`
	DestNodeHash   int `json:"destnodehash"`
}
type DestinyActivityGraphDisplayObjectiveDefinition struct {
	Id            int `json:"id"`
	ObjectiveHash int `json:"objectivehash"`
}
type DestinyActivityGraphDisplayProgressionDefinition struct {
	Id              int `json:"id"`
	ProgressionHash int `json:"progressionhash"`
}
type DestinyLinkedGraphDefinition struct {
	Description      string                              `json:"description"`
	Name             string                              `json:"name"`
	UnlockExpression DestinyUnlockExpressionDefinition   `json:"unlockexpression"`
	LinkedGraphId    int                                 `json:"linkedgraphid"`
	LinkedGraphs     []DestinyLinkedGraphEntryDefinition `json:"linkedgraphs"`
	Overview         string                              `json:"overview"`
}
type DestinyUnlockExpressionDefinition struct {
	Scope int `json:"scope"`
}
type DestinyLinkedGraphEntryDefinition struct {
	ActivityGraphHash int `json:"activitygraphhash"`
}
type DestinyDestinationBubbleSettingDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyBubbleDefinition struct {
	Hash              int                                `json:"hash"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyVendorGroupReference struct {
	VendorGroupHash int `json:"vendorgrouphash"`
}
type DestinyVendorGroupDefinition struct {
	Order        int    `json:"order"`
	CategoryName string `json:"categoryname"`
	Hash         int    `json:"hash"`
	Index        int    `json:"index"`
	Redacted     bool   `json:"redacted"`
}
type DestinyFactionDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	ProgressionHash   int                                `json:"progressionhash"`
	TokenValues       int                                `json:"tokenvalues"`
	RewardItemHash    int                                `json:"rewarditemhash"`
	RewardVendorHash  int                                `json:"rewardvendorhash"`
	Vendors           []DestinyFactionVendorDefinition   `json:"vendors"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyFactionVendorDefinition struct {
	VendorHash          int    `json:"vendorhash"`
	DestinationHash     int    `json:"destinationhash"`
	BackgroundImagePath string `json:"backgroundimagepath"`
}
type DestinyArtifactDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition    `json:"displayproperties"`
	TranslationBlock  DestinyItemTranslationBlockDefinition `json:"translationblock"`
	Tiers             []DestinyArtifactTierDefinition       `json:"tiers"`
	Hash              int                                   `json:"hash"`
	Index             int                                   `json:"index"`
	Redacted          bool                                  `json:"redacted"`
}
type DestinyArtifactTierDefinition struct {
	TierHash                           int                                 `json:"tierhash"`
	DisplayTitle                       string                              `json:"displaytitle"`
	ProgressRequirementMessage         string                              `json:"progressrequirementmessage"`
	Items                              []DestinyArtifactTierItemDefinition `json:"items"`
	MinimumUnlockPointsUsedRequirement int                                 `json:"minimumunlockpointsusedrequirement"`
}
type DestinyArtifactTierItemDefinition struct {
	ItemHash int `json:"itemhash"`
}
type DestinyItemQualityBlockDefinition struct {
	ItemLevels                      []int                          `json:"itemlevels"`
	QualityLevel                    int                            `json:"qualitylevel"`
	InfusionCategoryName            string                         `json:"infusioncategoryname"`
	InfusionCategoryHash            int                            `json:"infusioncategoryhash"`
	InfusionCategoryHashes          []int                          `json:"infusioncategoryhashes"`
	ProgressionLevelRequirementHash int                            `json:"progressionlevelrequirementhash"`
	CurrentVersion                  int                            `json:"currentversion"`
	Versions                        []DestinyItemVersionDefinition `json:"versions"`
	DisplayVersionWatermarkIcons    []string                       `json:"displayversionwatermarkicons"`
}
type DestinyItemVersionDefinition struct {
	PowerCapHash int `json:"powercaphash"`
}
type DestinyPowerCapDefinition struct {
	PowerCap int  `json:"powercap"`
	Hash     int  `json:"hash"`
	Index    int  `json:"index"`
	Redacted bool `json:"redacted"`
}
type DestinyProgressionLevelRequirementDefinition struct {
	RequirementCurve []InterpolationPointFloat `json:"requirementcurve"`
	ProgressionHash  int                       `json:"progressionhash"`
	Hash             int                       `json:"hash"`
	Index            int                       `json:"index"`
	Redacted         bool                      `json:"redacted"`
}
type InterpolationPointFloat struct {
	Value  float64 `json:"value"`
	Weight float64 `json:"weight"`
}
type DestinyItemValueBlockDefinition struct {
	ItemValue        []DestinyItemQuantity `json:"itemvalue"`
	ValueDescription string                `json:"valuedescription"`
}
type DestinyItemSourceBlockDefinition struct {
	SourceHashes  []int                              `json:"sourcehashes"`
	Sources       []DestinyItemSourceDefinition      `json:"sources"`
	Exclusive     int                                `json:"exclusive"`
	VendorSources []DestinyItemVendorSourceReference `json:"vendorsources"`
}
type DestinyItemSourceDefinition struct {
	Level            int                                `json:"level"`
	MinQuality       int                                `json:"minquality"`
	MaxQuality       int                                `json:"maxquality"`
	MinLevelRequired int                                `json:"minlevelrequired"`
	MaxLevelRequired int                                `json:"maxlevelrequired"`
	ComputedStats    DestinyInventoryItemStatDefinition `json:"computedstats"`
	SourceHashes     []int                              `json:"sourcehashes"`
}
type DestinyRewardSourceDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Category          int                                `json:"category"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyRewardSourceCategory struct {
	DestinyRewardSourceCategory int `json:"destinyrewardsourcecategory"`
}
type DestinyItemVendorSourceReference struct {
	VendorHash        int   `json:"vendorhash"`
	VendorItemIndexes []int `json:"vendoritemindexes"`
}
type DestinyItemObjectiveBlockDefinition struct {
	ObjectiveHashes                []int                               `json:"objectivehashes"`
	DisplayActivityHashes          []int                               `json:"displayactivityhashes"`
	RequireFullObjectiveCompletion bool                                `json:"requirefullobjectivecompletion"`
	QuestlineItemHash              int                                 `json:"questlineitemhash"`
	Narrative                      string                              `json:"narrative"`
	ObjectiveVerbName              string                              `json:"objectiveverbname"`
	QuestTypeIdentifier            string                              `json:"questtypeidentifier"`
	QuestTypeHash                  int                                 `json:"questtypehash"`
	PerObjectiveDisplayProperties  []DestinyObjectiveDisplayProperties `json:"perobjectivedisplayproperties"`
	DisplayAsStatTracker           bool                                `json:"displayasstattracker"`
}
type DestinyObjectiveDisplayProperties struct {
	ActivityHash               int  `json:"activityhash"`
	DisplayOnItemPreviewScreen bool `json:"displayonitempreviewscreen"`
}
type DestinyItemMetricBlockDefinition struct {
	AvailableMetricCategoryNodeHashes []int `json:"availablemetriccategorynodehashes"`
}
type DestinyPresentationNodeBaseDefinition struct {
	PresentationNodeType int      `json:"presentationnodetype"`
	TraitIds             []string `json:"traitids"`
	TraitHashes          []int    `json:"traithashes"`
	ParentNodeHashes     []int    `json:"parentnodehashes"`
	Hash                 int      `json:"hash"`
	Index                int      `json:"index"`
	Redacted             bool     `json:"redacted"`
}
type DestinyPresentationNodeType struct {
	DestinyPresentationNodeType int `json:"destinypresentationnodetype"`
}
type DestinyTraitDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	TraitCategoryId   string                             `json:"traitcategoryid"`
	TraitCategoryHash int                                `json:"traitcategoryhash"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyTraitCategoryDefinition struct {
	TraitCategoryId string   `json:"traitcategoryid"`
	TraitHashes     []int    `json:"traithashes"`
	TraitIds        []string `json:"traitids"`
	Hash            int      `json:"hash"`
	Index           int      `json:"index"`
	Redacted        bool     `json:"redacted"`
}
type DestinyScoredPresentationNodeBaseDefinition struct {
	MaxCategoryRecordScore int      `json:"maxcategoryrecordscore"`
	PresentationNodeType   int      `json:"presentationnodetype"`
	TraitIds               []string `json:"traitids"`
	TraitHashes            []int    `json:"traithashes"`
	ParentNodeHashes       []int    `json:"parentnodehashes"`
	Hash                   int      `json:"hash"`
	Index                  int      `json:"index"`
	Redacted               bool     `json:"redacted"`
}
type DestinyPresentationNodeDefinition struct {
	DisplayProperties               DestinyDisplayPropertiesDefinition       `json:"displayproperties"`
	OriginalIcon                    string                                   `json:"originalicon"`
	RootViewIcon                    string                                   `json:"rootviewicon"`
	NodeType                        int                                      `json:"nodetype"`
	Scope                           int                                      `json:"scope"`
	ObjectiveHash                   int                                      `json:"objectivehash"`
	CompletionRecordHash            int                                      `json:"completionrecordhash"`
	Children                        DestinyPresentationNodeChildrenBlock     `json:"children"`
	DisplayStyle                    int                                      `json:"displaystyle"`
	ScreenStyle                     int                                      `json:"screenstyle"`
	Requirements                    DestinyPresentationNodeRequirementsBlock `json:"requirements"`
	DisableChildSubscreenNavigation bool                                     `json:"disablechildsubscreennavigation"`
	MaxCategoryRecordScore          int                                      `json:"maxcategoryrecordscore"`
	PresentationNodeType            int                                      `json:"presentationnodetype"`
	TraitIds                        []string                                 `json:"traitids"`
	TraitHashes                     []int                                    `json:"traithashes"`
	ParentNodeHashes                []int                                    `json:"parentnodehashes"`
	Hash                            int                                      `json:"hash"`
	Index                           int                                      `json:"index"`
	Redacted                        bool                                     `json:"redacted"`
}
type DestinyScope struct {
	DestinyScope int `json:"destinyscope"`
}
type DestinyPresentationNodeChildrenBlock struct {
	PresentationNodes []DestinyPresentationNodeChildEntry            `json:"presentationnodes"`
	Collectibles      []DestinyPresentationNodeCollectibleChildEntry `json:"collectibles"`
	Records           []DestinyPresentationNodeRecordChildEntry      `json:"records"`
	Metrics           []DestinyPresentationNodeMetricChildEntry      `json:"metrics"`
}
type DestinyPresentationNodeChildEntry struct {
	PresentationNodeHash int `json:"presentationnodehash"`
}
type DestinyPresentationNodeCollectibleChildEntry struct {
	CollectibleHash int `json:"collectiblehash"`
}
type DestinyCollectibleDefinition struct {
	DisplayProperties    DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Scope                int                                `json:"scope"`
	SourceString         string                             `json:"sourcestring"`
	SourceHash           int                                `json:"sourcehash"`
	ItemHash             int                                `json:"itemhash"`
	AcquisitionInfo      DestinyCollectibleAcquisitionBlock `json:"acquisitioninfo"`
	StateInfo            DestinyCollectibleStateBlock       `json:"stateinfo"`
	PresentationInfo     DestinyPresentationChildBlock      `json:"presentationinfo"`
	PresentationNodeType int                                `json:"presentationnodetype"`
	TraitIds             []string                           `json:"traitids"`
	TraitHashes          []int                              `json:"traithashes"`
	ParentNodeHashes     []int                              `json:"parentnodehashes"`
	Hash                 int                                `json:"hash"`
	Index                int                                `json:"index"`
	Redacted             bool                               `json:"redacted"`
}
type DestinyCollectibleAcquisitionBlock struct {
	AcquireMaterialRequirementHash  int `json:"acquirematerialrequirementhash"`
	AcquireTimestampUnlockValueHash int `json:"acquiretimestampunlockvaluehash"`
}
type DestinyMaterialRequirementSetDefinition struct {
	Materials []DestinyMaterialRequirement `json:"materials"`
	Hash      int                          `json:"hash"`
	Index     int                          `json:"index"`
	Redacted  bool                         `json:"redacted"`
}
type DestinyMaterialRequirement struct {
	ItemHash             int  `json:"itemhash"`
	DeleteOnAction       bool `json:"deleteonaction"`
	Count                int  `json:"count"`
	OmitFromRequirements bool `json:"omitfromrequirements"`
}
type DestinyUnlockValueDefinition struct {
	Hash     int  `json:"hash"`
	Index    int  `json:"index"`
	Redacted bool `json:"redacted"`
}
type DestinyCollectibleStateBlock struct {
	ObscuredOverrideItemHash int                                      `json:"obscuredoverrideitemhash"`
	Requirements             DestinyPresentationNodeRequirementsBlock `json:"requirements"`
}
type DestinyPresentationNodeRequirementsBlock struct {
	EntitlementUnavailableMessage string `json:"entitlementunavailablemessage"`
}
type DestinyPresentationChildBlock struct {
	PresentationNodeType         int   `json:"presentationnodetype"`
	ParentPresentationNodeHashes []int `json:"parentpresentationnodehashes"`
	DisplayStyle                 int   `json:"displaystyle"`
}
type DestinyPresentationDisplayStyle struct {
	DestinyPresentationDisplayStyle int `json:"destinypresentationdisplaystyle"`
}
type DestinyPresentationNodeRecordChildEntry struct {
	RecordHash int `json:"recordhash"`
}
type DestinyRecordDefinition struct {
	DisplayProperties    DestinyDisplayPropertiesDefinition       `json:"displayproperties"`
	Scope                int                                      `json:"scope"`
	PresentationInfo     DestinyPresentationChildBlock            `json:"presentationinfo"`
	LoreHash             int                                      `json:"lorehash"`
	ObjectiveHashes      []int                                    `json:"objectivehashes"`
	RecordValueStyle     int                                      `json:"recordvaluestyle"`
	ForTitleGilding      bool                                     `json:"fortitlegilding"`
	TitleInfo            DestinyRecordTitleBlock                  `json:"titleinfo"`
	CompletionInfo       DestinyRecordCompletionBlock             `json:"completioninfo"`
	StateInfo            SchemaRecordStateBlock                   `json:"stateinfo"`
	Requirements         DestinyPresentationNodeRequirementsBlock `json:"requirements"`
	ExpirationInfo       DestinyRecordExpirationBlock             `json:"expirationinfo"`
	IntervalInfo         DestinyRecordIntervalBlock               `json:"intervalinfo"`
	RewardItems          []DestinyItemQuantity                    `json:"rewarditems"`
	PresentationNodeType int                                      `json:"presentationnodetype"`
	TraitIds             []string                                 `json:"traitids"`
	TraitHashes          []int                                    `json:"traithashes"`
	ParentNodeHashes     []int                                    `json:"parentnodehashes"`
	Hash                 int                                      `json:"hash"`
	Index                int                                      `json:"index"`
	Redacted             bool                                     `json:"redacted"`
}
type DestinyRecordValueStyle struct {
	DestinyRecordValueStyle int `json:"destinyrecordvaluestyle"`
}
type DestinyRecordTitleBlock struct {
	HasTitle                  bool   `json:"hastitle"`
	TitlesByGender            string `json:"titlesbygender"`
	TitlesByGenderHash        string `json:"titlesbygenderhash"`
	GildingTrackingRecordHash int    `json:"gildingtrackingrecordhash"`
}
type DestinyGender struct {
	DestinyGender int `json:"destinygender"`
}
type DestinyGenderDefinition struct {
	GenderType        int                                `json:"gendertype"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyRecordCompletionBlock struct {
	PartialCompletionObjectiveCountThreshold int  `json:"partialcompletionobjectivecountthreshold"`
	ScoreValue                               int  `json:"scorevalue"`
	ShouldFireToast                          bool `json:"shouldfiretoast"`
	ToastStyle                               int  `json:"toaststyle"`
}
type DestinyRecordToastStyle struct {
	DestinyRecordToastStyle int `json:"destinyrecordtoaststyle"`
}
type SchemaRecordStateBlock struct {
	FeaturedPriority int    `json:"featuredpriority"`
	ObscuredString   string `json:"obscuredstring"`
}
type DestinyRecordExpirationBlock struct {
	HasExpiration bool   `json:"hasexpiration"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
}
type DestinyRecordIntervalBlock struct {
	IntervalObjectives                   []DestinyRecordIntervalObjective `json:"intervalobjectives"`
	IntervalRewards                      []DestinyRecordIntervalRewards   `json:"intervalrewards"`
	OriginalObjectiveArrayInsertionIndex int                              `json:"originalobjectivearrayinsertionindex"`
}
type DestinyRecordIntervalObjective struct {
	IntervalObjectiveHash int `json:"intervalobjectivehash"`
	IntervalScoreValue    int `json:"intervalscorevalue"`
}
type DestinyRecordIntervalRewards struct {
	IntervalRewardItems []DestinyItemQuantity `json:"intervalrewarditems"`
}
type DestinyLoreDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Subtitle          string                             `json:"subtitle"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyPresentationNodeMetricChildEntry struct {
	MetricHash int `json:"metrichash"`
}
type DestinyMetricDefinition struct {
	DisplayProperties     DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	TrackingObjectiveHash int                                `json:"trackingobjectivehash"`
	LowerValueIsBetter    bool                               `json:"lowervalueisbetter"`
	PresentationNodeType  int                                `json:"presentationnodetype"`
	TraitIds              []string                           `json:"traitids"`
	TraitHashes           []int                              `json:"traithashes"`
	ParentNodeHashes      []int                              `json:"parentnodehashes"`
	Hash                  int                                `json:"hash"`
	Index                 int                                `json:"index"`
	Redacted              bool                               `json:"redacted"`
}
type DestinyPresentationScreenStyle struct {
	DestinyPresentationScreenStyle int `json:"destinypresentationscreenstyle"`
}
type DestinyItemPlugDefinition struct {
	InsertionRules                   []DestinyPlugRuleDefinition `json:"insertionrules"`
	PlugCategoryIdentifier           string                      `json:"plugcategoryidentifier"`
	PlugCategoryHash                 int                         `json:"plugcategoryhash"`
	OnActionRecreateSelf             bool                        `json:"onactionrecreateself"`
	InsertionMaterialRequirementHash int                         `json:"insertionmaterialrequirementhash"`
	PreviewItemOverrideHash          int                         `json:"previewitemoverridehash"`
	EnabledMaterialRequirementHash   int                         `json:"enabledmaterialrequirementhash"`
	EnabledRules                     []DestinyPlugRuleDefinition `json:"enabledrules"`
	UiPlugLabel                      string                      `json:"uipluglabel"`
	PlugStyle                        int                         `json:"plugstyle"`
	PlugAvailability                 int                         `json:"plugavailability"`
	AlternateUiPlugLabel             string                      `json:"alternateuipluglabel"`
	AlternatePlugStyle               int                         `json:"alternateplugstyle"`
	IsDummyPlug                      bool                        `json:"isdummyplug"`
	ParentItemOverride               DestinyParentItemOverride   `json:"parentitemoverride"`
	EnergyCapacity                   DestinyEnergyCapacityEntry  `json:"energycapacity"`
	EnergyCost                       DestinyEnergyCostEntry      `json:"energycost"`
}
type DestinyPlugRuleDefinition struct {
	FailureMessage string `json:"failuremessage"`
}
type PlugUiStyles struct {
	PlugUiStyles int `json:"pluguistyles"`
}
type PlugAvailabilityMode struct {
	PlugAvailabilityMode int `json:"plugavailabilitymode"`
}
type DestinyParentItemOverride struct {
	AdditionalEquipRequirementsDisplayStrings []string `json:"additionalequiprequirementsdisplaystrings"`
	PipIcon                                   string   `json:"pipicon"`
}
type DestinyEnergyCapacityEntry struct {
	CapacityValue  int `json:"capacityvalue"`
	EnergyTypeHash int `json:"energytypehash"`
	EnergyType     int `json:"energytype"`
}
type DestinyEnergyType struct {
	DestinyEnergyType int `json:"destinyenergytype"`
}
type DestinyEnergyTypeDefinition struct {
	DisplayProperties   DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	TransparentIconPath string                             `json:"transparenticonpath"`
	ShowIcon            bool                               `json:"showicon"`
	EnumValue           int                                `json:"enumvalue"`
	CapacityStatHash    int                                `json:"capacitystathash"`
	CostStatHash        int                                `json:"coststathash"`
	Hash                int                                `json:"hash"`
	Index               int                                `json:"index"`
	Redacted            bool                               `json:"redacted"`
}
type DestinyEnergyCostEntry struct {
	EnergyCost     int `json:"energycost"`
	EnergyTypeHash int `json:"energytypehash"`
	EnergyType     int `json:"energytype"`
}
type DestinyItemGearsetBlockDefinition struct {
	TrackingValueMax int   `json:"trackingvaluemax"`
	ItemList         []int `json:"itemlist"`
}
type DestinyItemSackBlockDefinition struct {
	DetailAction    string `json:"detailaction"`
	OpenAction      string `json:"openaction"`
	SelectItemCount int    `json:"selectitemcount"`
	VendorSackType  string `json:"vendorsacktype"`
	OpenOnAcquire   bool   `json:"openonacquire"`
}
type DestinyItemSocketBlockDefinition struct {
	Detail           string                                      `json:"detail"`
	SocketEntries    []DestinyItemSocketEntryDefinition          `json:"socketentries"`
	IntrinsicSockets []DestinyItemIntrinsicSocketEntryDefinition `json:"intrinsicsockets"`
	SocketCategories []DestinyItemSocketCategoryDefinition       `json:"socketcategories"`
}
type DestinyItemSocketEntryDefinition struct {
	SocketTypeHash                        int                                        `json:"sockettypehash"`
	SingleInitialItemHash                 int                                        `json:"singleinitialitemhash"`
	ReusablePlugItems                     []DestinyItemSocketEntryPlugItemDefinition `json:"reusableplugitems"`
	PreventInitializationOnVendorPurchase bool                                       `json:"preventinitializationonvendorpurchase"`
	HidePerksInItemTooltip                bool                                       `json:"hideperksinitemtooltip"`
	PlugSources                           int                                        `json:"plugsources"`
	ReusablePlugSetHash                   int                                        `json:"reusableplugsethash"`
	RandomizedPlugSetHash                 int                                        `json:"randomizedplugsethash"`
	DefaultVisible                        bool                                       `json:"defaultvisible"`
}
type DestinyItemSocketEntryPlugItemDefinition struct {
	PlugItemHash int `json:"plugitemhash"`
}
type SocketPlugSources struct {
	SocketPlugSources int `json:"socketplugsources"`
}
type DestinyPlugSetDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition                   `json:"displayproperties"`
	ReusablePlugItems []DestinyItemSocketEntryPlugItemRandomizedDefinition `json:"reusableplugitems"`
	IsFakePlugSet     bool                                                 `json:"isfakeplugset"`
	Hash              int                                                  `json:"hash"`
	Index             int                                                  `json:"index"`
	Redacted          bool                                                 `json:"redacted"`
}
type DestinyItemSocketEntryPlugItemRandomizedDefinition struct {
	CurrentlyCanRoll bool `json:"currentlycanroll"`
	PlugItemHash     int  `json:"plugitemhash"`
}
type DestinyItemIntrinsicSocketEntryDefinition struct {
	PlugItemHash   int  `json:"plugitemhash"`
	SocketTypeHash int  `json:"sockettypehash"`
	DefaultVisible bool `json:"defaultvisible"`
}
type DestinyItemSocketCategoryDefinition struct {
	SocketCategoryHash int   `json:"socketcategoryhash"`
	SocketIndexes      []int `json:"socketindexes"`
}
type DestinyItemSummaryBlockDefinition struct {
	SortPriority int `json:"sortpriority"`
}
type DestinyItemTalentGridBlockDefinition struct {
	TalentGridHash   int    `json:"talentgridhash"`
	ItemDetailString string `json:"itemdetailstring"`
	BuildName        string `json:"buildname"`
	HudDamageType    int    `json:"huddamagetype"`
	HudIcon          string `json:"hudicon"`
}
type DestinyTalentGridDefinition struct {
	MaxGridLevel           int                                       `json:"maxgridlevel"`
	GridLevelPerColumn     int                                       `json:"gridlevelpercolumn"`
	ProgressionHash        int                                       `json:"progressionhash"`
	Nodes                  []DestinyTalentNodeDefinition             `json:"nodes"`
	ExclusiveSets          []DestinyTalentNodeExclusiveSetDefinition `json:"exclusivesets"`
	IndependentNodeIndexes []int                                     `json:"independentnodeindexes"`
	Groups                 DestinyTalentExclusiveGroup               `json:"groups"`
	NodeCategories         []DestinyTalentNodeCategory               `json:"nodecategories"`
	Hash                   int                                       `json:"hash"`
	Index                  int                                       `json:"index"`
	Redacted               bool                                      `json:"redacted"`
}
type DestinyTalentNodeDefinition struct {
	NodeIndex                              int                              `json:"nodeindex"`
	NodeHash                               int                              `json:"nodehash"`
	Row                                    int                              `json:"row"`
	Column                                 int                              `json:"column"`
	PrerequisiteNodeIndexes                []int                            `json:"prerequisitenodeindexes"`
	BinaryPairNodeIndex                    int                              `json:"binarypairnodeindex"`
	AutoUnlocks                            bool                             `json:"autounlocks"`
	LastStepRepeats                        bool                             `json:"laststeprepeats"`
	IsRandom                               bool                             `json:"israndom"`
	RandomActivationRequirement            DestinyNodeActivationRequirement `json:"randomactivationrequirement"`
	IsRandomRepurchasable                  bool                             `json:"israndomrepurchasable"`
	Steps                                  []DestinyNodeStepDefinition      `json:"steps"`
	ExclusiveWithNodeHashes                []int                            `json:"exclusivewithnodehashes"`
	RandomStartProgressionBarAtProgression int                              `json:"randomstartprogressionbaratprogression"`
	LayoutIdentifier                       string                           `json:"layoutidentifier"`
	GroupHash                              int                              `json:"grouphash"`
	LoreHash                               int                              `json:"lorehash"`
	NodeStyleIdentifier                    string                           `json:"nodestyleidentifier"`
	IgnoreForCompletion                    bool                             `json:"ignoreforcompletion"`
}
type DestinyNodeActivationRequirement struct {
	GridLevel                 int   `json:"gridlevel"`
	MaterialRequirementHashes []int `json:"materialrequirementhashes"`
}
type DestinyNodeStepDefinition struct {
	DisplayProperties             DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	StepIndex                     int                                `json:"stepindex"`
	NodeStepHash                  int                                `json:"nodestephash"`
	InteractionDescription        string                             `json:"interactiondescription"`
	DamageType                    int                                `json:"damagetype"`
	DamageTypeHash                int                                `json:"damagetypehash"`
	ActivationRequirement         DestinyNodeActivationRequirement   `json:"activationrequirement"`
	CanActivateNextStep           bool                               `json:"canactivatenextstep"`
	NextStepIndex                 int                                `json:"nextstepindex"`
	IsNextStepRandom              bool                               `json:"isnextsteprandom"`
	PerkHashes                    []int                              `json:"perkhashes"`
	StartProgressionBarAtProgress int                                `json:"startprogressionbaratprogress"`
	StatHashes                    []int                              `json:"stathashes"`
	AffectsQuality                bool                               `json:"affectsquality"`
	StepGroups                    DestinyTalentNodeStepGroups        `json:"stepgroups"`
	AffectsLevel                  bool                               `json:"affectslevel"`
	SocketReplacements            []DestinyNodeSocketReplaceResponse `json:"socketreplacements"`
}
type DestinyNodeSocketReplaceResponse struct {
	SocketTypeHash int `json:"sockettypehash"`
	PlugItemHash   int `json:"plugitemhash"`
}
type DestinyDamageTypeDefinition struct {
	DisplayProperties   DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	TransparentIconPath string                             `json:"transparenticonpath"`
	ShowIcon            bool                               `json:"showicon"`
	EnumValue           int                                `json:"enumvalue"`
	Hash                int                                `json:"hash"`
	Index               int                                `json:"index"`
	Redacted            bool                               `json:"redacted"`
}
type DestinyTalentNodeExclusiveSetDefinition struct {
	NodeIndexes []int `json:"nodeindexes"`
}
type DestinyTalentExclusiveGroup struct {
	GroupHash           int   `json:"grouphash"`
	LoreHash            int   `json:"lorehash"`
	NodeHashes          []int `json:"nodehashes"`
	OpposingGroupHashes []int `json:"opposinggrouphashes"`
	OpposingNodeHashes  []int `json:"opposingnodehashes"`
}
type DestinyTalentNodeCategory struct {
	Identifier        string                             `json:"identifier"`
	IsLoreDriven      bool                               `json:"isloredriven"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	NodeHashes        []int                              `json:"nodehashes"`
}
type DestinyItemPerkEntryDefinition struct {
	RequirementDisplayString string `json:"requirementdisplaystring"`
	PerkHash                 int    `json:"perkhash"`
	PerkVisibility           int    `json:"perkvisibility"`
}
type ItemPerkVisibility struct {
	ItemPerkVisibility int `json:"itemperkvisibility"`
}
type DestinyAnimationReference struct {
	AnimName       string `json:"animname"`
	AnimIdentifier string `json:"animidentifier"`
	Path           string `json:"path"`
}
type HyperlinkReference struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}
type SpecialItemType struct {
	SpecialItemType int `json:"specialitemtype"`
}
type DestinyItemType struct {
	DestinyItemType int `json:"destinyitemtype"`
}
type DestinyClass struct {
	DestinyClass int `json:"destinyclass"`
}
type DestinyBreakerType struct {
	DestinyBreakerType int `json:"destinybreakertype"`
}
type DestinyItemCategoryDefinition struct {
	DisplayProperties       DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Visible                 bool                               `json:"visible"`
	Deprecated              bool                               `json:"deprecated"`
	ShortTitle              string                             `json:"shorttitle"`
	ItemTypeRegex           string                             `json:"itemtyperegex"`
	GrantDestinyBreakerType int                                `json:"grantdestinybreakertype"`
	PlugCategoryIdentifier  string                             `json:"plugcategoryidentifier"`
	ItemTypeRegexNot        string                             `json:"itemtyperegexnot"`
	OriginBucketIdentifier  string                             `json:"originbucketidentifier"`
	GrantDestinyItemType    int                                `json:"grantdestinyitemtype"`
	GrantDestinySubType     int                                `json:"grantdestinysubtype"`
	GrantDestinyClass       int                                `json:"grantdestinyclass"`
	TraitId                 string                             `json:"traitid"`
	GroupedCategoryHashes   []int                              `json:"groupedcategoryhashes"`
	ParentCategoryHashes    []int                              `json:"parentcategoryhashes"`
	GroupCategoryOnly       bool                               `json:"groupcategoryonly"`
	Hash                    int                                `json:"hash"`
	Index                   int                                `json:"index"`
	Redacted                bool                               `json:"redacted"`
}
type DestinyBreakerTypeDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	EnumValue         int                                `json:"enumvalue"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinySeasonDefinition struct {
	DisplayProperties                      DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	BackgroundImagePath                    string                             `json:"backgroundimagepath"`
	SeasonNumber                           int                                `json:"seasonnumber"`
	StartDate                              string                             `json:"startdate"`
	EndDate                                string                             `json:"enddate"`
	SeasonPassHash                         int                                `json:"seasonpasshash"`
	SeasonPassProgressionHash              int                                `json:"seasonpassprogressionhash"`
	ArtifactItemHash                       int                                `json:"artifactitemhash"`
	SealPresentationNodeHash               int                                `json:"sealpresentationnodehash"`
	SeasonalChallengesPresentationNodeHash int                                `json:"seasonalchallengespresentationnodehash"`
	Preview                                DestinySeasonPreviewDefinition     `json:"preview"`
	Hash                                   int                                `json:"hash"`
	Index                                  int                                `json:"index"`
	Redacted                               bool                               `json:"redacted"`
}
type DestinySeasonPreviewDefinition struct {
	Description string                                `json:"description"`
	LinkPath    string                                `json:"linkpath"`
	VideoLink   string                                `json:"videolink"`
	Images      []DestinySeasonPreviewImageDefinition `json:"images"`
}
type DestinySeasonPreviewImageDefinition struct {
	ThumbnailImage string `json:"thumbnailimage"`
	HighResImage   string `json:"highresimage"`
}
type DestinySeasonPassDefinition struct {
	DisplayProperties       DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	RewardProgressionHash   int                                `json:"rewardprogressionhash"`
	PrestigeProgressionHash int                                `json:"prestigeprogressionhash"`
	Hash                    int                                `json:"hash"`
	Index                   int                                `json:"index"`
	Redacted                bool                               `json:"redacted"`
}
type DestinyProgressionRewardItemQuantity struct {
	RewardedAtProgressionLevel int      `json:"rewardedatprogressionlevel"`
	AcquisitionBehavior        int      `json:"acquisitionbehavior"`
	UiDisplayStyle             string   `json:"uidisplaystyle"`
	ClaimUnlockDisplayStrings  []string `json:"claimunlockdisplaystrings"`
	ItemHash                   int      `json:"itemhash"`
	ItemInstanceId             int      `json:"iteminstanceid"`
	Quantity                   int      `json:"quantity"`
	HasConditionalVisibility   bool     `json:"hasconditionalvisibility"`
}
type DestinyProgressionRewardItemAcquisitionBehavior struct {
	DestinyProgressionRewardItemAcquisitionBehavior int `json:"destinyprogressionrewarditemacquisitionbehavior"`
}
type GroupUserBase struct {
	GroupId           int               `json:"groupid"`
	DestinyUserInfo   GroupUserInfoCard `json:"destinyuserinfo"`
	BungieNetUserInfo UserInfoCard      `json:"bungienetuserinfo"`
	JoinDate          string            `json:"joindate"`
}
type GroupMember struct {
	MemberType             int               `json:"membertype"`
	IsOnline               bool              `json:"isonline"`
	LastOnlineStatusChange int               `json:"lastonlinestatuschange"`
	GroupId                int               `json:"groupid"`
	DestinyUserInfo        GroupUserInfoCard `json:"destinyuserinfo"`
	BungieNetUserInfo      UserInfoCard      `json:"bungienetuserinfo"`
	JoinDate               string            `json:"joindate"`
}
type GroupAllianceStatus struct {
	GroupAllianceStatus int `json:"groupalliancestatus"`
}
type GroupPotentialMember struct {
	PotentialStatus   int               `json:"potentialstatus"`
	GroupId           int               `json:"groupid"`
	DestinyUserInfo   GroupUserInfoCard `json:"destinyuserinfo"`
	BungieNetUserInfo UserInfoCard      `json:"bungienetuserinfo"`
	JoinDate          string            `json:"joindate"`
}
type GroupPotentialMemberStatus struct {
	GroupPotentialMemberStatus int `json:"grouppotentialmemberstatus"`
}
type TagResponse struct {
	TagText      string         `json:"tagtext"`
	IgnoreStatus IgnoreResponse `json:"ignorestatus"`
}
type PollResponse struct {
	TopicId    int          `json:"topicid"`
	Results    []PollResult `json:"results"`
	TotalVotes int          `json:"totalvotes"`
}
type PollResult struct {
	AnswerText          string `json:"answertext"`
	AnswerSlot          int    `json:"answerslot"`
	LastVoteDate        string `json:"lastvotedate"`
	Votes               int    `json:"votes"`
	RequestingUserVoted bool   `json:"requestinguservoted"`
}
type ForumRecruitmentDetail struct {
	TopicId              int           `json:"topicid"`
	MicrophoneRequired   bool          `json:"microphonerequired"`
	Intensity            int           `json:"intensity"`
	Tone                 int           `json:"tone"`
	Approved             bool          `json:"approved"`
	ConversationId       int           `json:"conversationid"`
	PlayerSlotsTotal     int           `json:"playerslotstotal"`
	PlayerSlotsRemaining int           `json:"playerslotsremaining"`
	Fireteam             []GeneralUser `json:"fireteam"`
	KickedPlayerIds      []int         `json:"kickedplayerids"`
}
type ForumRecruitmentIntensityLabel struct {
	ForumRecruitmentIntensityLabel int `json:"forumrecruitmentintensitylabel"`
}
type ForumRecruitmentToneLabel struct {
	ForumRecruitmentToneLabel int `json:"forumrecruitmenttonelabel"`
}
type ForumPostSortEnum struct {
	ForumPostSortEnum int `json:"forumpostsortenum"`
}
type GroupTheme struct {
	Name        string `json:"name"`
	Folder      string `json:"folder"`
	Description string `json:"description"`
}
type GroupDateRange struct {
	GroupDateRange int `json:"groupdaterange"`
}
type GroupV2Card struct {
	GroupId          int             `json:"groupid"`
	Name             string          `json:"name"`
	GroupType        int             `json:"grouptype"`
	CreationDate     string          `json:"creationdate"`
	About            string          `json:"about"`
	Motto            string          `json:"motto"`
	MemberCount      int             `json:"membercount"`
	Locale           string          `json:"locale"`
	MembershipOption int             `json:"membershipoption"`
	Capabilities     int             `json:"capabilities"`
	ClanInfo         GroupV2ClanInfo `json:"claninfo"`
	AvatarPath       string          `json:"avatarpath"`
	Theme            string          `json:"theme"`
}
type SearchResultOfGroupV2Card struct {
	Results                      []GroupV2Card `json:"results"`
	TotalResults                 int           `json:"totalresults"`
	HasMore                      bool          `json:"hasmore"`
	Query                        PagedQuery    `json:"query"`
	ReplacementContinuationToken string        `json:"replacementcontinuationtoken"`
	UseTotalResults              bool          `json:"usetotalresults"`
}
type GroupSearchResponse struct {
	Results                      []GroupV2Card `json:"results"`
	TotalResults                 int           `json:"totalresults"`
	HasMore                      bool          `json:"hasmore"`
	Query                        PagedQuery    `json:"query"`
	ReplacementContinuationToken string        `json:"replacementcontinuationtoken"`
	UseTotalResults              bool          `json:"usetotalresults"`
}
type GroupQuery struct {
	Name                     string `json:"name"`
	GroupType                int    `json:"grouptype"`
	CreationDate             int    `json:"creationdate"`
	SortBy                   int    `json:"sortby"`
	GroupMemberCountFilter   int    `json:"groupmembercountfilter"`
	LocaleFilter             string `json:"localefilter"`
	TagText                  string `json:"tagtext"`
	ItemsPerPage             int    `json:"itemsperpage"`
	CurrentPage              int    `json:"currentpage"`
	RequestContinuationToken string `json:"requestcontinuationtoken"`
}
type GroupSortBy struct {
	GroupSortBy int `json:"groupsortby"`
}
type GroupMemberCountFilter struct {
	GroupMemberCountFilter int `json:"groupmembercountfilter"`
}
type GroupNameSearchRequest struct {
	GroupName string `json:"groupname"`
	GroupType int    `json:"grouptype"`
}
type GroupOptionalConversation struct {
	GroupId        int    `json:"groupid"`
	ConversationId int    `json:"conversationid"`
	ChatEnabled    bool   `json:"chatenabled"`
	ChatName       string `json:"chatname"`
	ChatSecurity   int    `json:"chatsecurity"`
}
type GroupEditAction struct {
	Name                               string `json:"name"`
	About                              string `json:"about"`
	Motto                              string `json:"motto"`
	Theme                              string `json:"theme"`
	AvatarImageIndex                   int    `json:"avatarimageindex"`
	Tags                               string `json:"tags"`
	IsPublic                           bool   `json:"ispublic"`
	MembershipOption                   int    `json:"membershipoption"`
	IsPublicTopicAdminOnly             bool   `json:"ispublictopicadminonly"`
	AllowChat                          bool   `json:"allowchat"`
	ChatSecurity                       int    `json:"chatsecurity"`
	Callsign                           string `json:"callsign"`
	Locale                             string `json:"locale"`
	Homepage                           int    `json:"homepage"`
	EnableInvitationMessagingForAdmins bool   `json:"enableinvitationmessagingforadmins"`
	DefaultPublicity                   int    `json:"defaultpublicity"`
}
type GroupOptionsEditAction struct {
	InvitePermissionOverride         bool `json:"invitepermissionoverride"`
	UpdateCulturePermissionOverride  bool `json:"updateculturepermissionoverride"`
	HostGuidedGamePermissionOverride int  `json:"hostguidedgamepermissionoverride"`
	UpdateBannerPermissionOverride   bool `json:"updatebannerpermissionoverride"`
	JoinLevel                        int  `json:"joinlevel"`
}
type GroupOptionalConversationAddRequest struct {
	ChatName     string `json:"chatname"`
	ChatSecurity int    `json:"chatsecurity"`
}
type GroupOptionalConversationEditRequest struct {
	ChatEnabled  bool   `json:"chatenabled"`
	ChatName     string `json:"chatname"`
	ChatSecurity int    `json:"chatsecurity"`
}
type SearchResultOfGroupMember struct {
	Results                      []GroupMember `json:"results"`
	TotalResults                 int           `json:"totalresults"`
	HasMore                      bool          `json:"hasmore"`
	Query                        PagedQuery    `json:"query"`
	ReplacementContinuationToken string        `json:"replacementcontinuationtoken"`
	UseTotalResults              bool          `json:"usetotalresults"`
}
type GroupMemberLeaveResult struct {
	Group        GroupV2 `json:"group"`
	GroupDeleted bool    `json:"groupdeleted"`
}
type GroupBanRequest struct {
	Comment string `json:"comment"`
	Length  int    `json:"length"`
}
type IgnoreLength struct {
	IgnoreLength int `json:"ignorelength"`
}
type GroupBan struct {
	GroupId           int               `json:"groupid"`
	LastModifiedBy    UserInfoCard      `json:"lastmodifiedby"`
	CreatedBy         UserInfoCard      `json:"createdby"`
	DateBanned        string            `json:"datebanned"`
	DateExpires       string            `json:"dateexpires"`
	Comment           string            `json:"comment"`
	BungieNetUserInfo UserInfoCard      `json:"bungienetuserinfo"`
	DestinyUserInfo   GroupUserInfoCard `json:"destinyuserinfo"`
}
type SearchResultOfGroupBan struct {
	Results                      []GroupBan `json:"results"`
	TotalResults                 int        `json:"totalresults"`
	HasMore                      bool       `json:"hasmore"`
	Query                        PagedQuery `json:"query"`
	ReplacementContinuationToken string     `json:"replacementcontinuationtoken"`
	UseTotalResults              bool       `json:"usetotalresults"`
}
type GroupMemberApplication struct {
	GroupId                int               `json:"groupid"`
	CreationDate           string            `json:"creationdate"`
	ResolveState           int               `json:"resolvestate"`
	ResolveDate            string            `json:"resolvedate"`
	ResolvedByMembershipId int               `json:"resolvedbymembershipid"`
	RequestMessage         string            `json:"requestmessage"`
	ResolveMessage         string            `json:"resolvemessage"`
	DestinyUserInfo        GroupUserInfoCard `json:"destinyuserinfo"`
	BungieNetUserInfo      UserInfoCard      `json:"bungienetuserinfo"`
}
type GroupApplicationResolveState struct {
	GroupApplicationResolveState int `json:"groupapplicationresolvestate"`
}
type SearchResultOfGroupMemberApplication struct {
	Results                      []GroupMemberApplication `json:"results"`
	TotalResults                 int                      `json:"totalresults"`
	HasMore                      bool                     `json:"hasmore"`
	Query                        PagedQuery               `json:"query"`
	ReplacementContinuationToken string                   `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                     `json:"usetotalresults"`
}
type EntityActionResult struct {
	EntityId int `json:"entityid"`
	Result   int `json:"result"`
}
type PlatformErrorCodes struct {
	PlatformErrorCodes int `json:"platformerrorcodes"`
}
type GroupApplicationRequest struct {
	Message string `json:"message"`
}
type GroupApplicationListRequest struct {
	Memberships []UserMembership `json:"memberships"`
	Message     string           `json:"message"`
}
type GroupsForMemberFilter struct {
	GroupsForMemberFilter int `json:"groupsformemberfilter"`
}
type GroupMembershipBase struct {
	Group GroupV2 `json:"group"`
}
type GroupMembership struct {
	Member GroupMember `json:"member"`
	Group  GroupV2     `json:"group"`
}
type SearchResultOfGroupMembership struct {
	Results                      []GroupMembership `json:"results"`
	TotalResults                 int               `json:"totalresults"`
	HasMore                      bool              `json:"hasmore"`
	Query                        PagedQuery        `json:"query"`
	ReplacementContinuationToken string            `json:"replacementcontinuationtoken"`
	UseTotalResults              bool              `json:"usetotalresults"`
}
type GroupMembershipSearchResponse struct {
	Results                      []GroupMembership `json:"results"`
	TotalResults                 int               `json:"totalresults"`
	HasMore                      bool              `json:"hasmore"`
	Query                        PagedQuery        `json:"query"`
	ReplacementContinuationToken string            `json:"replacementcontinuationtoken"`
	UseTotalResults              bool              `json:"usetotalresults"`
}
type GetGroupsForMemberResponse struct {
	AreAllMembershipsInactive    bool              `json:"areallmembershipsinactive"`
	Results                      []GroupMembership `json:"results"`
	TotalResults                 int               `json:"totalresults"`
	HasMore                      bool              `json:"hasmore"`
	Query                        PagedQuery        `json:"query"`
	ReplacementContinuationToken string            `json:"replacementcontinuationtoken"`
	UseTotalResults              bool              `json:"usetotalresults"`
}
type GroupPotentialMembership struct {
	Member GroupPotentialMember `json:"member"`
	Group  GroupV2              `json:"group"`
}
type SearchResultOfGroupPotentialMembership struct {
	Results                      []GroupPotentialMembership `json:"results"`
	TotalResults                 int                        `json:"totalresults"`
	HasMore                      bool                       `json:"hasmore"`
	Query                        PagedQuery                 `json:"query"`
	ReplacementContinuationToken string                     `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                       `json:"usetotalresults"`
}
type GroupPotentialMembershipSearchResponse struct {
	Results                      []GroupPotentialMembership `json:"results"`
	TotalResults                 int                        `json:"totalresults"`
	HasMore                      bool                       `json:"hasmore"`
	Query                        PagedQuery                 `json:"query"`
	ReplacementContinuationToken string                     `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                       `json:"usetotalresults"`
}
type GroupApplicationResponse struct {
	Resolution int `json:"resolution"`
}
type PartnerOfferClaimRequest struct {
	PartnerOfferId        string `json:"partnerofferid"`
	BungieNetMembershipId int    `json:"bungienetmembershipid"`
	TransactionId         string `json:"transactionid"`
}
type PartnerOfferSkuHistoryResponse struct {
	SkuIdentifier        string                        `json:"skuidentifier"`
	LocalizedName        string                        `json:"localizedname"`
	LocalizedDescription string                        `json:"localizeddescription"`
	ClaimDate            string                        `json:"claimdate"`
	AllOffersApplied     bool                          `json:"alloffersapplied"`
	TransactionId        string                        `json:"transactionid"`
	SkuOffers            []PartnerOfferHistoryResponse `json:"skuoffers"`
}
type PartnerOfferHistoryResponse struct {
	PartnerOfferKey      string `json:"partnerofferkey"`
	MembershipId         int    `json:"membershipid"`
	MembershipType       int    `json:"membershiptype"`
	LocalizedName        string `json:"localizedname"`
	LocalizedDescription string `json:"localizeddescription"`
	IsConsumable         bool   `json:"isconsumable"`
	QuantityApplied      int    `json:"quantityapplied"`
	ApplyDate            string `json:"applydate"`
}
type DestinyManifest struct {
	Version                        string                        `json:"version"`
	MobileAssetContentPath         string                        `json:"mobileassetcontentpath"`
	MobileGearAssetDataBases       []GearAssetDataBaseDefinition `json:"mobilegearassetdatabases"`
	MobileWorldContentPaths        string                        `json:"mobileworldcontentpaths"`
	JsonWorldContentPaths          string                        `json:"jsonworldcontentpaths"`
	JsonWorldComponentContentPaths interface{}                   `json:"jsonworldcomponentcontentpaths"`
	MobileClanBannerDatabasePath   string                        `json:"mobileclanbannerdatabasepath"`
	MobileGearCDN                  string                        `json:"mobilegearcdn"`
	IconImagePyramidInfo           []ImagePyramidEntry           `json:"iconimagepyramidinfo"`
}
type GearAssetDataBaseDefinition struct {
	Version int    `json:"version"`
	Path    string `json:"path"`
}
type ImagePyramidEntry struct {
	Name   string  `json:"name"`
	Factor float64 `json:"factor"`
}
type DestinyLinkedProfilesResponse struct {
	Profiles           []DestinyProfileUserInfoCard `json:"profiles"`
	BnetMembership     UserInfoCard                 `json:"bnetmembership"`
	ProfilesWithErrors []DestinyErrorProfile        `json:"profileswitherrors"`
}
type DestinyProfileUserInfoCard struct {
	DateLastPlayed              string                         `json:"datelastplayed"`
	IsOverridden                bool                           `json:"isoverridden"`
	IsCrossSavePrimary          bool                           `json:"iscrosssaveprimary"`
	PlatformSilver              DestinyPlatformSilverComponent `json:"platformsilver"`
	UnpairedGameVersions        int                            `json:"unpairedgameversions"`
	SupplementalDisplayName     string                         `json:"supplementaldisplayname"`
	IconPath                    string                         `json:"iconpath"`
	CrossSaveOverride           int                            `json:"crosssaveoverride"`
	ApplicableMembershipTypes   []int                          `json:"applicablemembershiptypes"`
	IsPublic                    bool                           `json:"ispublic"`
	MembershipType              int                            `json:"membershiptype"`
	MembershipId                int                            `json:"membershipid"`
	DisplayName                 string                         `json:"displayname"`
	BungieGlobalDisplayName     string                         `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int                            `json:"bungieglobaldisplaynamecode"`
}
type DestinyPlatformSilverComponent struct {
	PlatformSilver DestinyItemComponent `json:"platformsilver"`
}
type DestinyItemComponent struct {
	ItemHash                   int                      `json:"itemhash"`
	ItemInstanceId             int                      `json:"iteminstanceid"`
	Quantity                   int                      `json:"quantity"`
	BindStatus                 int                      `json:"bindstatus"`
	Location                   int                      `json:"location"`
	BucketHash                 int                      `json:"buckethash"`
	TransferStatus             int                      `json:"transferstatus"`
	Lockable                   bool                     `json:"lockable"`
	State                      int                      `json:"state"`
	OverrideStyleItemHash      int                      `json:"overridestyleitemhash"`
	ExpirationDate             string                   `json:"expirationdate"`
	IsWrapper                  bool                     `json:"iswrapper"`
	TooltipNotificationIndexes []int                    `json:"tooltipnotificationindexes"`
	MetricHash                 int                      `json:"metrichash"`
	MetricObjective            DestinyObjectiveProgress `json:"metricobjective"`
	VersionNumber              int                      `json:"versionnumber"`
	ItemValueVisibility        []bool                   `json:"itemvaluevisibility"`
}
type ItemBindStatus struct {
	ItemBindStatus int `json:"itembindstatus"`
}
type TransferStatuses struct {
	TransferStatuses int `json:"transferstatuses"`
}
type ItemState struct {
	ItemState int `json:"itemstate"`
}
type DestinyObjectiveProgress struct {
	ObjectiveHash   int  `json:"objectivehash"`
	DestinationHash int  `json:"destinationhash"`
	ActivityHash    int  `json:"activityhash"`
	Progress        int  `json:"progress"`
	CompletionValue int  `json:"completionvalue"`
	Complete        bool `json:"complete"`
	Visible         bool `json:"visible"`
}
type DestinyGameVersions struct {
	DestinyGameVersions int `json:"destinygameversions"`
}
type DestinyErrorProfile struct {
	ErrorCode int          `json:"errorcode"`
	InfoCard  UserInfoCard `json:"infocard"`
}
type DestinyComponentType struct {
	DestinyComponentType int `json:"destinycomponenttype"`
}
type DestinyProfileResponse struct {
	VendorReceipts                     SingleComponentResponseOfDestinyVendorReceiptsComponent                   `json:"vendorreceipts"`
	ProfileInventory                   SingleComponentResponseOfDestinyInventoryComponent                        `json:"profileinventory"`
	ProfileCurrencies                  SingleComponentResponseOfDestinyInventoryComponent                        `json:"profilecurrencies"`
	Profile                            SingleComponentResponseOfDestinyProfileComponent                          `json:"profile"`
	PlatformSilver                     SingleComponentResponseOfDestinyPlatformSilverComponent                   `json:"platformsilver"`
	ProfileKiosks                      SingleComponentResponseOfDestinyKiosksComponent                           `json:"profilekiosks"`
	ProfilePlugSets                    SingleComponentResponseOfDestinyPlugSetsComponent                         `json:"profileplugsets"`
	ProfileProgression                 SingleComponentResponseOfDestinyProfileProgressionComponent               `json:"profileprogression"`
	ProfilePresentationNodes           SingleComponentResponseOfDestinyPresentationNodesComponent                `json:"profilepresentationnodes"`
	ProfileRecords                     SingleComponentResponseOfDestinyProfileRecordsComponent                   `json:"profilerecords"`
	ProfileCollectibles                SingleComponentResponseOfDestinyProfileCollectiblesComponent              `json:"profilecollectibles"`
	ProfileTransitoryData              SingleComponentResponseOfDestinyProfileTransitoryComponent                `json:"profiletransitorydata"`
	Metrics                            SingleComponentResponseOfDestinyMetricsComponent                          `json:"metrics"`
	ProfileStringVariables             SingleComponentResponseOfDestinyStringVariablesComponent                  `json:"profilestringvariables"`
	Characters                         DictionaryComponentResponseOfint64AndDestinyCharacterComponent            `json:"characters"`
	CharacterInventories               DictionaryComponentResponseOfint64AndDestinyInventoryComponent            `json:"characterinventories"`
	CharacterProgressions              DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent `json:"characterprogressions"`
	CharacterRenderData                DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent      `json:"characterrenderdata"`
	CharacterActivities                DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent  `json:"characteractivities"`
	CharacterEquipment                 DictionaryComponentResponseOfint64AndDestinyInventoryComponent            `json:"characterequipment"`
	CharacterKiosks                    DictionaryComponentResponseOfint64AndDestinyKiosksComponent               `json:"characterkiosks"`
	CharacterPlugSets                  DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent             `json:"characterplugsets"`
	CharacterUninstancedItemComponents DestinyBaseItemComponentSetOfuint32                                       `json:"characteruninstanceditemcomponents"`
	CharacterPresentationNodes         DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent    `json:"characterpresentationnodes"`
	CharacterRecords                   DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent     `json:"characterrecords"`
	CharacterCollectibles              DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent         `json:"charactercollectibles"`
	CharacterStringVariables           DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent      `json:"characterstringvariables"`
	ItemComponents                     DestinyItemComponentSetOfint64                                            `json:"itemcomponents"`
	CharacterCurrencyLookups           DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent           `json:"charactercurrencylookups"`
}
type DestinyVendorReceiptsComponent struct {
	Receipts []DestinyVendorReceipt `json:"receipts"`
}
type DestinyVendorReceipt struct {
	CurrencyPaid           []DestinyItemQuantity `json:"currencypaid"`
	ItemReceived           DestinyItemQuantity   `json:"itemreceived"`
	LicenseUnlockHash      int                   `json:"licenseunlockhash"`
	PurchasedByCharacterId int                   `json:"purchasedbycharacterid"`
	RefundPolicy           int                   `json:"refundpolicy"`
	SequenceNumber         int                   `json:"sequencenumber"`
	TimeToExpiration       int                   `json:"timetoexpiration"`
	ExpiresOn              string                `json:"expireson"`
}
type ComponentResponse struct {
	Privacy  int  `json:"privacy"`
	Disabled bool `json:"disabled"`
}
type ComponentPrivacySetting struct {
	ComponentPrivacySetting int `json:"componentprivacysetting"`
}
type SingleComponentResponseOfDestinyVendorReceiptsComponent struct {
	Data     DestinyVendorReceiptsComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyInventoryComponent struct {
	Items []DestinyItemComponent `json:"items"`
}
type SingleComponentResponseOfDestinyInventoryComponent struct {
	Data     DestinyInventoryComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DestinyProfileComponent struct {
	UserInfo                    UserInfoCard `json:"userinfo"`
	DateLastPlayed              string       `json:"datelastplayed"`
	VersionsOwned               int          `json:"versionsowned"`
	CharacterIds                []int        `json:"characterids"`
	SeasonHashes                []int        `json:"seasonhashes"`
	CurrentSeasonHash           int          `json:"currentseasonhash"`
	CurrentSeasonRewardPowerCap int          `json:"currentseasonrewardpowercap"`
}
type SingleComponentResponseOfDestinyProfileComponent struct {
	Data     DestinyProfileComponent `json:"data"`
	Privacy  int                     `json:"privacy"`
	Disabled bool                    `json:"disabled"`
}
type SingleComponentResponseOfDestinyPlatformSilverComponent struct {
	Data     DestinyPlatformSilverComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyKiosksComponent struct {
	KioskItems []DestinyKioskItem `json:"kioskitems"`
}
type DestinyKioskItem struct {
	Index           int                      `json:"index"`
	CanAcquire      bool                     `json:"canacquire"`
	FailureIndexes  []int                    `json:"failureindexes"`
	FlavorObjective DestinyObjectiveProgress `json:"flavorobjective"`
}
type SingleComponentResponseOfDestinyKiosksComponent struct {
	Data     DestinyKiosksComponent `json:"data"`
	Privacy  int                    `json:"privacy"`
	Disabled bool                   `json:"disabled"`
}
type DestinyPlugSetsComponent struct {
	Plugs []DestinyItemPlug `json:"plugs"`
}
type DestinyItemPlugBase struct {
	PlugItemHash      int   `json:"plugitemhash"`
	CanInsert         bool  `json:"caninsert"`
	Enabled           bool  `json:"enabled"`
	InsertFailIndexes []int `json:"insertfailindexes"`
	EnableFailIndexes []int `json:"enablefailindexes"`
}
type DestinyItemPlug struct {
	PlugObjectives    []DestinyObjectiveProgress `json:"plugobjectives"`
	PlugItemHash      int                        `json:"plugitemhash"`
	CanInsert         bool                       `json:"caninsert"`
	Enabled           bool                       `json:"enabled"`
	InsertFailIndexes []int                      `json:"insertfailindexes"`
	EnableFailIndexes []int                      `json:"enablefailindexes"`
}
type SingleComponentResponseOfDestinyPlugSetsComponent struct {
	Data     DestinyPlugSetsComponent `json:"data"`
	Privacy  int                      `json:"privacy"`
	Disabled bool                     `json:"disabled"`
}
type DestinyProfileProgressionComponent struct {
	Checklists       interface{}                  `json:"checklists"`
	SeasonalArtifact DestinyArtifactProfileScoped `json:"seasonalartifact"`
}
type DestinyArtifactProfileScoped struct {
	ArtifactHash          int                `json:"artifacthash"`
	PointProgression      DestinyProgression `json:"pointprogression"`
	PointsAcquired        int                `json:"pointsacquired"`
	PowerBonusProgression DestinyProgression `json:"powerbonusprogression"`
	PowerBonus            int                `json:"powerbonus"`
}
type DestinyChecklistDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	ViewActionString  string                             `json:"viewactionstring"`
	Scope             int                                `json:"scope"`
	Entries           []DestinyChecklistEntryDefinition  `json:"entries"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyChecklistEntryDefinition struct {
	Hash                   int                                `json:"hash"`
	DisplayProperties      DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	DestinationHash        int                                `json:"destinationhash"`
	LocationHash           int                                `json:"locationhash"`
	BubbleHash             int                                `json:"bubblehash"`
	ActivityHash           int                                `json:"activityhash"`
	ItemHash               int                                `json:"itemhash"`
	VendorHash             int                                `json:"vendorhash"`
	VendorInteractionIndex int                                `json:"vendorinteractionindex"`
	Scope                  int                                `json:"scope"`
}
type SingleComponentResponseOfDestinyProfileProgressionComponent struct {
	Data     DestinyProfileProgressionComponent `json:"data"`
	Privacy  int                                `json:"privacy"`
	Disabled bool                               `json:"disabled"`
}
type DestinyPresentationNodesComponent struct {
	Nodes DestinyPresentationNodeComponent `json:"nodes"`
}
type DestinyPresentationNodeComponent struct {
	State               int                      `json:"state"`
	Objective           DestinyObjectiveProgress `json:"objective"`
	ProgressValue       int                      `json:"progressvalue"`
	CompletionValue     int                      `json:"completionvalue"`
	RecordCategoryScore int                      `json:"recordcategoryscore"`
}
type DestinyPresentationNodeState struct {
	DestinyPresentationNodeState int `json:"destinypresentationnodestate"`
}
type SingleComponentResponseOfDestinyPresentationNodesComponent struct {
	Data     DestinyPresentationNodesComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DestinyRecordsComponent struct {
	Records                      DestinyRecordComponent `json:"records"`
	RecordCategoriesRootNodeHash int                    `json:"recordcategoriesrootnodehash"`
	RecordSealsRootNodeHash      int                    `json:"recordsealsrootnodehash"`
}
type DestinyRecordComponent struct {
	State                  int                        `json:"state"`
	Objectives             []DestinyObjectiveProgress `json:"objectives"`
	IntervalObjectives     []DestinyObjectiveProgress `json:"intervalobjectives"`
	IntervalsRedeemedCount int                        `json:"intervalsredeemedcount"`
	CompletedCount         int                        `json:"completedcount"`
	RewardVisibilty        []bool                     `json:"rewardvisibilty"`
}
type DestinyRecordState struct {
	DestinyRecordState int `json:"destinyrecordstate"`
}
type DestinyProfileRecordsComponent struct {
	Score                        int                    `json:"score"`
	ActiveScore                  int                    `json:"activescore"`
	LegacyScore                  int                    `json:"legacyscore"`
	LifetimeScore                int                    `json:"lifetimescore"`
	TrackedRecordHash            int                    `json:"trackedrecordhash"`
	Records                      DestinyRecordComponent `json:"records"`
	RecordCategoriesRootNodeHash int                    `json:"recordcategoriesrootnodehash"`
	RecordSealsRootNodeHash      int                    `json:"recordsealsrootnodehash"`
}
type SingleComponentResponseOfDestinyProfileRecordsComponent struct {
	Data     DestinyProfileRecordsComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyCollectiblesComponent struct {
	Collectibles                     DestinyCollectibleComponent `json:"collectibles"`
	CollectionCategoriesRootNodeHash int                         `json:"collectioncategoriesrootnodehash"`
	CollectionBadgesRootNodeHash     int                         `json:"collectionbadgesrootnodehash"`
}
type DestinyCollectibleComponent struct {
	State int `json:"state"`
}
type DestinyCollectibleState struct {
	DestinyCollectibleState int `json:"destinycollectiblestate"`
}
type DestinyProfileCollectiblesComponent struct {
	RecentCollectibleHashes          []int                       `json:"recentcollectiblehashes"`
	NewnessFlaggedCollectibleHashes  []int                       `json:"newnessflaggedcollectiblehashes"`
	Collectibles                     DestinyCollectibleComponent `json:"collectibles"`
	CollectionCategoriesRootNodeHash int                         `json:"collectioncategoriesrootnodehash"`
	CollectionBadgesRootNodeHash     int                         `json:"collectionbadgesrootnodehash"`
}
type SingleComponentResponseOfDestinyProfileCollectiblesComponent struct {
	Data     DestinyProfileCollectiblesComponent `json:"data"`
	Privacy  int                                 `json:"privacy"`
	Disabled bool                                `json:"disabled"`
}
type DestinyProfileTransitoryComponent struct {
	PartyMembers               []DestinyProfileTransitoryPartyMember   `json:"partymembers"`
	CurrentActivity            DestinyProfileTransitoryCurrentActivity `json:"currentactivity"`
	Joinability                DestinyProfileTransitoryJoinability     `json:"joinability"`
	Tracking                   []DestinyProfileTransitoryTrackingEntry `json:"tracking"`
	LastOrbitedDestinationHash int                                     `json:"lastorbiteddestinationhash"`
}
type DestinyProfileTransitoryPartyMember struct {
	MembershipId int    `json:"membershipid"`
	EmblemHash   int    `json:"emblemhash"`
	DisplayName  string `json:"displayname"`
	Status       int    `json:"status"`
}
type DestinyPartyMemberStates struct {
	DestinyPartyMemberStates int `json:"destinypartymemberstates"`
}
type DestinyProfileTransitoryCurrentActivity struct {
	StartTime                   string  `json:"starttime"`
	EndTime                     string  `json:"endtime"`
	Score                       float64 `json:"score"`
	HighestOpposingFactionScore float64 `json:"highestopposingfactionscore"`
	NumberOfOpponents           int     `json:"numberofopponents"`
	NumberOfPlayers             int     `json:"numberofplayers"`
}
type DestinyProfileTransitoryJoinability struct {
	OpenSlots      int `json:"openslots"`
	PrivacySetting int `json:"privacysetting"`
	ClosedReasons  int `json:"closedreasons"`
}
type DestinyGamePrivacySetting struct {
	DestinyGamePrivacySetting int `json:"destinygameprivacysetting"`
}
type DestinyJoinClosedReasons struct {
	DestinyJoinClosedReasons int `json:"destinyjoinclosedreasons"`
}
type DestinyProfileTransitoryTrackingEntry struct {
	LocationHash      int    `json:"locationhash"`
	ItemHash          int    `json:"itemhash"`
	ObjectiveHash     int    `json:"objectivehash"`
	ActivityHash      int    `json:"activityhash"`
	QuestlineItemHash int    `json:"questlineitemhash"`
	TrackedDate       string `json:"trackeddate"`
}
type SingleComponentResponseOfDestinyProfileTransitoryComponent struct {
	Data     DestinyProfileTransitoryComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DestinyMetricsComponent struct {
	Metrics             DestinyMetricComponent `json:"metrics"`
	MetricsRootNodeHash int                    `json:"metricsrootnodehash"`
}
type DestinyMetricComponent struct {
	Invisible         bool                     `json:"invisible"`
	ObjectiveProgress DestinyObjectiveProgress `json:"objectiveprogress"`
}
type SingleComponentResponseOfDestinyMetricsComponent struct {
	Data     DestinyMetricsComponent `json:"data"`
	Privacy  int                     `json:"privacy"`
	Disabled bool                    `json:"disabled"`
}
type DestinyStringVariablesComponent struct {
	IntegerValuesByHash int `json:"integervaluesbyhash"`
}
type SingleComponentResponseOfDestinyStringVariablesComponent struct {
	Data     DestinyStringVariablesComponent `json:"data"`
	Privacy  int                             `json:"privacy"`
	Disabled bool                            `json:"disabled"`
}
type DestinyCharacterComponent struct {
	MembershipId             int                `json:"membershipid"`
	MembershipType           int                `json:"membershiptype"`
	CharacterId              int                `json:"characterid"`
	DateLastPlayed           string             `json:"datelastplayed"`
	MinutesPlayedThisSession int                `json:"minutesplayedthissession"`
	MinutesPlayedTotal       int                `json:"minutesplayedtotal"`
	Light                    int                `json:"light"`
	Stats                    int                `json:"stats"`
	RaceHash                 int                `json:"racehash"`
	GenderHash               int                `json:"genderhash"`
	ClassHash                int                `json:"classhash"`
	RaceType                 int                `json:"racetype"`
	ClassType                int                `json:"classtype"`
	GenderType               int                `json:"gendertype"`
	EmblemPath               string             `json:"emblempath"`
	EmblemBackgroundPath     string             `json:"emblembackgroundpath"`
	EmblemHash               int                `json:"emblemhash"`
	EmblemColor              DestinyColor       `json:"emblemcolor"`
	LevelProgression         DestinyProgression `json:"levelprogression"`
	BaseCharacterLevel       int                `json:"basecharacterlevel"`
	PercentToNextLevel       float64            `json:"percenttonextlevel"`
	TitleRecordHash          int                `json:"titlerecordhash"`
}
type DestinyRace struct {
	DestinyRace int `json:"destinyrace"`
}
type DestinyRaceDefinition struct {
	DisplayProperties             DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	RaceType                      int                                `json:"racetype"`
	GenderedRaceNames             string                             `json:"genderedracenames"`
	GenderedRaceNamesByGenderHash string                             `json:"genderedracenamesbygenderhash"`
	Hash                          int                                `json:"hash"`
	Index                         int                                `json:"index"`
	Redacted                      bool                               `json:"redacted"`
}
type DestinyClassDefinition struct {
	ClassType                      int                                `json:"classtype"`
	DisplayProperties              DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	GenderedClassNames             string                             `json:"genderedclassnames"`
	GenderedClassNamesByGenderHash string                             `json:"genderedclassnamesbygenderhash"`
	MentorVendorHash               int                                `json:"mentorvendorhash"`
	Hash                           int                                `json:"hash"`
	Index                          int                                `json:"index"`
	Redacted                       bool                               `json:"redacted"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterComponent struct {
	Data     DestinyCharacterComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyInventoryComponent struct {
	Data     DestinyInventoryComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DestinyCharacterProgressionComponent struct {
	Progressions              DestinyProgression             `json:"progressions"`
	Factions                  DestinyFactionProgression      `json:"factions"`
	Milestones                DestinyMilestone               `json:"milestones"`
	Quests                    []DestinyQuestStatus           `json:"quests"`
	UninstancedItemObjectives []DestinyObjectiveProgress     `json:"uninstanceditemobjectives"`
	UninstancedItemPerks      DestinyItemPerksComponent      `json:"uninstanceditemperks"`
	Checklists                interface{}                    `json:"checklists"`
	SeasonalArtifact          DestinyArtifactCharacterScoped `json:"seasonalartifact"`
}
type DestinyFactionProgression struct {
	FactionHash         int                            `json:"factionhash"`
	FactionVendorIndex  int                            `json:"factionvendorindex"`
	ProgressionHash     int                            `json:"progressionhash"`
	DailyProgress       int                            `json:"dailyprogress"`
	DailyLimit          int                            `json:"dailylimit"`
	WeeklyProgress      int                            `json:"weeklyprogress"`
	WeeklyLimit         int                            `json:"weeklylimit"`
	CurrentProgress     int                            `json:"currentprogress"`
	Level               int                            `json:"level"`
	LevelCap            int                            `json:"levelcap"`
	StepIndex           int                            `json:"stepindex"`
	ProgressToNextLevel int                            `json:"progresstonextlevel"`
	NextLevelAt         int                            `json:"nextlevelat"`
	CurrentResetCount   int                            `json:"currentresetcount"`
	SeasonResets        []DestinyProgressionResetEntry `json:"seasonresets"`
	RewardItemStates    []int                          `json:"rewarditemstates"`
}
type DestinyMilestone struct {
	MilestoneHash   int                                 `json:"milestonehash"`
	AvailableQuests []DestinyMilestoneQuest             `json:"availablequests"`
	Activities      []DestinyMilestoneChallengeActivity `json:"activities"`
	Values          float64                             `json:"values"`
	VendorHashes    []int                               `json:"vendorhashes"`
	Vendors         []DestinyMilestoneVendor            `json:"vendors"`
	Rewards         []DestinyMilestoneRewardCategory    `json:"rewards"`
	StartDate       string                              `json:"startdate"`
	EndDate         string                              `json:"enddate"`
	Order           int                                 `json:"order"`
}
type DestinyMilestoneQuest struct {
	QuestItemHash int                      `json:"questitemhash"`
	Status        DestinyQuestStatus       `json:"status"`
	Activity      DestinyMilestoneActivity `json:"activity"`
	Challenges    []DestinyChallengeStatus `json:"challenges"`
}
type DestinyQuestStatus struct {
	QuestHash      int                        `json:"questhash"`
	StepHash       int                        `json:"stephash"`
	StepObjectives []DestinyObjectiveProgress `json:"stepobjectives"`
	Tracked        bool                       `json:"tracked"`
	ItemInstanceId int                        `json:"iteminstanceid"`
	Completed      bool                       `json:"completed"`
	Redeemed       bool                       `json:"redeemed"`
	Started        bool                       `json:"started"`
	VendorHash     int                        `json:"vendorhash"`
}
type DestinyMilestoneActivity struct {
	ActivityHash     int                               `json:"activityhash"`
	ActivityModeHash int                               `json:"activitymodehash"`
	ActivityModeType int                               `json:"activitymodetype"`
	ModifierHashes   []int                             `json:"modifierhashes"`
	Variants         []DestinyMilestoneActivityVariant `json:"variants"`
}
type DestinyMilestoneActivityVariant struct {
	ActivityHash     int                                      `json:"activityhash"`
	CompletionStatus DestinyMilestoneActivityCompletionStatus `json:"completionstatus"`
	ActivityModeHash int                                      `json:"activitymodehash"`
	ActivityModeType int                                      `json:"activitymodetype"`
}
type DestinyMilestoneActivityCompletionStatus struct {
	Completed bool                            `json:"completed"`
	Phases    []DestinyMilestoneActivityPhase `json:"phases"`
}
type DestinyMilestoneActivityPhase struct {
	Complete  bool `json:"complete"`
	PhaseHash int  `json:"phasehash"`
}
type DestinyChallengeStatus struct {
	Objective DestinyObjectiveProgress `json:"objective"`
}
type DestinyMilestoneChallengeActivity struct {
	ActivityHash            int                             `json:"activityhash"`
	Challenges              []DestinyChallengeStatus        `json:"challenges"`
	ModifierHashes          []int                           `json:"modifierhashes"`
	BooleanActivityOptions  bool                            `json:"booleanactivityoptions"`
	LoadoutRequirementIndex int                             `json:"loadoutrequirementindex"`
	Phases                  []DestinyMilestoneActivityPhase `json:"phases"`
}
type DestinyMilestoneVendor struct {
	VendorHash      int `json:"vendorhash"`
	PreviewItemHash int `json:"previewitemhash"`
}
type DestinyMilestoneRewardCategory struct {
	RewardCategoryHash int                           `json:"rewardcategoryhash"`
	Entries            []DestinyMilestoneRewardEntry `json:"entries"`
}
type DestinyMilestoneRewardEntry struct {
	RewardEntryHash int  `json:"rewardentryhash"`
	Earned          bool `json:"earned"`
	Redeemed        bool `json:"redeemed"`
}
type DestinyMilestoneDefinition struct {
	DisplayProperties               DestinyDisplayPropertiesDefinition            `json:"displayproperties"`
	DisplayPreference               int                                           `json:"displaypreference"`
	Image                           string                                        `json:"image"`
	MilestoneType                   int                                           `json:"milestonetype"`
	Recruitable                     bool                                          `json:"recruitable"`
	FriendlyName                    string                                        `json:"friendlyname"`
	ShowInExplorer                  bool                                          `json:"showinexplorer"`
	ShowInMilestones                bool                                          `json:"showinmilestones"`
	ExplorePrioritizesActivityImage bool                                          `json:"exploreprioritizesactivityimage"`
	HasPredictableDates             bool                                          `json:"haspredictabledates"`
	Quests                          DestinyMilestoneQuestDefinition               `json:"quests"`
	Rewards                         DestinyMilestoneRewardCategoryDefinition      `json:"rewards"`
	VendorsDisplayTitle             string                                        `json:"vendorsdisplaytitle"`
	Vendors                         []DestinyMilestoneVendorDefinition            `json:"vendors"`
	Values                          DestinyMilestoneValueDefinition               `json:"values"`
	IsInGameMilestone               bool                                          `json:"isingamemilestone"`
	Activities                      []DestinyMilestoneChallengeActivityDefinition `json:"activities"`
	DefaultOrder                    int                                           `json:"defaultorder"`
	Hash                            int                                           `json:"hash"`
	Index                           int                                           `json:"index"`
	Redacted                        bool                                          `json:"redacted"`
}
type DestinyMilestoneDisplayPreference struct {
	DestinyMilestoneDisplayPreference int `json:"destinymilestonedisplaypreference"`
}
type DestinyMilestoneType struct {
	DestinyMilestoneType int `json:"destinymilestonetype"`
}
type DestinyMilestoneQuestDefinition struct {
	QuestItemHash     int                                    `json:"questitemhash"`
	DisplayProperties DestinyDisplayPropertiesDefinition     `json:"displayproperties"`
	OverrideImage     string                                 `json:"overrideimage"`
	QuestRewards      DestinyMilestoneQuestRewardsDefinition `json:"questrewards"`
	Activities        DestinyMilestoneActivityDefinition     `json:"activities"`
	DestinationHash   int                                    `json:"destinationhash"`
}
type DestinyMilestoneQuestRewardsDefinition struct {
	Items []DestinyMilestoneQuestRewardItem `json:"items"`
}
type DestinyMilestoneQuestRewardItem struct {
	VendorHash               int  `json:"vendorhash"`
	VendorItemIndex          int  `json:"vendoritemindex"`
	ItemHash                 int  `json:"itemhash"`
	ItemInstanceId           int  `json:"iteminstanceid"`
	Quantity                 int  `json:"quantity"`
	HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyMilestoneActivityDefinition struct {
	ConceptualActivityHash int                                       `json:"conceptualactivityhash"`
	Variants               DestinyMilestoneActivityVariantDefinition `json:"variants"`
}
type DestinyMilestoneActivityVariantDefinition struct {
	ActivityHash int `json:"activityhash"`
	Order        int `json:"order"`
}
type DestinyMilestoneRewardCategoryDefinition struct {
	CategoryHash       int                                   `json:"categoryhash"`
	CategoryIdentifier string                                `json:"categoryidentifier"`
	DisplayProperties  DestinyDisplayPropertiesDefinition    `json:"displayproperties"`
	RewardEntries      DestinyMilestoneRewardEntryDefinition `json:"rewardentries"`
	Order              int                                   `json:"order"`
}
type DestinyMilestoneRewardEntryDefinition struct {
	RewardEntryHash       int                                `json:"rewardentryhash"`
	RewardEntryIdentifier string                             `json:"rewardentryidentifier"`
	Items                 []DestinyItemQuantity              `json:"items"`
	VendorHash            int                                `json:"vendorhash"`
	DisplayProperties     DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Order                 int                                `json:"order"`
}
type DestinyMilestoneVendorDefinition struct {
	VendorHash int `json:"vendorhash"`
}
type DestinyMilestoneValueDefinition struct {
	Key               string                             `json:"key"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyMilestoneChallengeActivityDefinition struct {
	ActivityHash       int                                               `json:"activityhash"`
	Challenges         []DestinyMilestoneChallengeDefinition             `json:"challenges"`
	ActivityGraphNodes []DestinyMilestoneChallengeActivityGraphNodeEntry `json:"activitygraphnodes"`
	Phases             []DestinyMilestoneChallengeActivityPhase          `json:"phases"`
}
type DestinyMilestoneChallengeDefinition struct {
	ChallengeObjectiveHash int `json:"challengeobjectivehash"`
}
type DestinyMilestoneChallengeActivityGraphNodeEntry struct {
	ActivityGraphHash     int `json:"activitygraphhash"`
	ActivityGraphNodeHash int `json:"activitygraphnodehash"`
}
type DestinyMilestoneChallengeActivityPhase struct {
	PhaseHash int `json:"phasehash"`
}
type DestinyItemPerksComponent struct {
	Perks []DestinyPerkReference `json:"perks"`
}
type DestinyPerkReference struct {
	PerkHash int    `json:"perkhash"`
	IconPath string `json:"iconpath"`
	IsActive bool   `json:"isactive"`
	Visible  bool   `json:"visible"`
}
type DestinyArtifactCharacterScoped struct {
	ArtifactHash int                   `json:"artifacthash"`
	PointsUsed   int                   `json:"pointsused"`
	ResetCount   int                   `json:"resetcount"`
	Tiers        []DestinyArtifactTier `json:"tiers"`
}
type DestinyArtifactTier struct {
	TierHash       int                       `json:"tierhash"`
	IsUnlocked     bool                      `json:"isunlocked"`
	PointsToUnlock int                       `json:"pointstounlock"`
	Items          []DestinyArtifactTierItem `json:"items"`
}
type DestinyArtifactTierItem struct {
	ItemHash int  `json:"itemhash"`
	IsActive bool `json:"isactive"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent struct {
	Data     DestinyCharacterProgressionComponent `json:"data"`
	Privacy  int                                  `json:"privacy"`
	Disabled bool                                 `json:"disabled"`
}
type DestinyCharacterRenderComponent struct {
	CustomDyes    []DyeReference                `json:"customdyes"`
	Customization DestinyCharacterCustomization `json:"customization"`
	PeerView      DestinyCharacterPeerView      `json:"peerview"`
}
type DestinyCharacterCustomization struct {
	Personality   int   `json:"personality"`
	Face          int   `json:"face"`
	SkinColor     int   `json:"skincolor"`
	LipColor      int   `json:"lipcolor"`
	EyeColor      int   `json:"eyecolor"`
	HairColors    []int `json:"haircolors"`
	FeatureColors []int `json:"featurecolors"`
	DecalColor    int   `json:"decalcolor"`
	WearHelmet    bool  `json:"wearhelmet"`
	HairIndex     int   `json:"hairindex"`
	FeatureIndex  int   `json:"featureindex"`
	DecalIndex    int   `json:"decalindex"`
}
type DestinyCharacterPeerView struct {
	Equipment []DestinyItemPeerView `json:"equipment"`
}
type DestinyItemPeerView struct {
	ItemHash int            `json:"itemhash"`
	Dyes     []DyeReference `json:"dyes"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent struct {
	Data     DestinyCharacterRenderComponent `json:"data"`
	Privacy  int                             `json:"privacy"`
	Disabled bool                            `json:"disabled"`
}
type DestinyCharacterActivitiesComponent struct {
	DateActivityStarted         string            `json:"dateactivitystarted"`
	AvailableActivities         []DestinyActivity `json:"availableactivities"`
	CurrentActivityHash         int               `json:"currentactivityhash"`
	CurrentActivityModeHash     int               `json:"currentactivitymodehash"`
	CurrentActivityModeType     int               `json:"currentactivitymodetype"`
	CurrentActivityModeHashes   []int             `json:"currentactivitymodehashes"`
	CurrentActivityModeTypes    []int             `json:"currentactivitymodetypes"`
	CurrentPlaylistActivityHash int               `json:"currentplaylistactivityhash"`
	LastCompletedStoryHash      int               `json:"lastcompletedstoryhash"`
}
type DestinyActivity struct {
	ActivityHash            int                      `json:"activityhash"`
	IsNew                   bool                     `json:"isnew"`
	CanLead                 bool                     `json:"canlead"`
	CanJoin                 bool                     `json:"canjoin"`
	IsCompleted             bool                     `json:"iscompleted"`
	IsVisible               bool                     `json:"isvisible"`
	DisplayLevel            int                      `json:"displaylevel"`
	RecommendedLight        int                      `json:"recommendedlight"`
	DifficultyTier          int                      `json:"difficultytier"`
	Challenges              []DestinyChallengeStatus `json:"challenges"`
	ModifierHashes          []int                    `json:"modifierhashes"`
	BooleanActivityOptions  bool                     `json:"booleanactivityoptions"`
	LoadoutRequirementIndex int                      `json:"loadoutrequirementindex"`
}
type DestinyActivityDifficultyTier struct {
	DestinyActivityDifficultyTier int `json:"destinyactivitydifficultytier"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent struct {
	Data     DestinyCharacterActivitiesComponent `json:"data"`
	Privacy  int                                 `json:"privacy"`
	Disabled bool                                `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyKiosksComponent struct {
	Data     DestinyKiosksComponent `json:"data"`
	Privacy  int                    `json:"privacy"`
	Disabled bool                   `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent struct {
	Data     DestinyPlugSetsComponent `json:"data"`
	Privacy  int                      `json:"privacy"`
	Disabled bool                     `json:"disabled"`
}
type DestinyBaseItemComponentSetOfuint32 struct {
	Objectives DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent      `json:"perks"`
}
type DestinyItemObjectivesComponent struct {
	Objectives      []DestinyObjectiveProgress `json:"objectives"`
	FlavorObjective DestinyObjectiveProgress   `json:"flavorobjective"`
	DateCompleted   string                     `json:"datecompleted"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent struct {
	Data     DestinyItemObjectivesComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent struct {
	Data     DestinyItemPerksComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent struct {
	Data     DestinyPresentationNodesComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DestinyCharacterRecordsComponent struct {
	FeaturedRecordHashes         []int                  `json:"featuredrecordhashes"`
	Records                      DestinyRecordComponent `json:"records"`
	RecordCategoriesRootNodeHash int                    `json:"recordcategoriesrootnodehash"`
	RecordSealsRootNodeHash      int                    `json:"recordsealsrootnodehash"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent struct {
	Data     DestinyCharacterRecordsComponent `json:"data"`
	Privacy  int                              `json:"privacy"`
	Disabled bool                             `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent struct {
	Data     DestinyCollectiblesComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent struct {
	Data     DestinyStringVariablesComponent `json:"data"`
	Privacy  int                             `json:"privacy"`
	Disabled bool                            `json:"disabled"`
}
type DestinyBaseItemComponentSetOfint64 struct {
	Objectives DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfint64AndDestinyItemPerksComponent      `json:"perks"`
}
type DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent struct {
	Data     DestinyItemObjectivesComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyItemPerksComponent struct {
	Data     DestinyItemPerksComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DestinyItemComponentSetOfint64 struct {
	Instances      DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent       `json:"instances"`
	RenderData     DictionaryComponentResponseOfint64AndDestinyItemRenderComponent         `json:"renderdata"`
	Stats          DictionaryComponentResponseOfint64AndDestinyItemStatsComponent          `json:"stats"`
	Sockets        DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent        `json:"sockets"`
	ReusablePlugs  DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent  `json:"reusableplugs"`
	PlugObjectives DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
	TalentGrids    DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent     `json:"talentgrids"`
	PlugStates     DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent          `json:"plugstates"`
	Objectives     DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent     `json:"objectives"`
	Perks          DictionaryComponentResponseOfint64AndDestinyItemPerksComponent          `json:"perks"`
}
type DestinyItemInstanceComponent struct {
	DamageType                  int                       `json:"damagetype"`
	DamageTypeHash              int                       `json:"damagetypehash"`
	PrimaryStat                 DestinyStat               `json:"primarystat"`
	ItemLevel                   int                       `json:"itemlevel"`
	Quality                     int                       `json:"quality"`
	IsEquipped                  bool                      `json:"isequipped"`
	CanEquip                    bool                      `json:"canequip"`
	EquipRequiredLevel          int                       `json:"equiprequiredlevel"`
	UnlockHashesRequiredToEquip []int                     `json:"unlockhashesrequiredtoequip"`
	CannotEquipReason           int                       `json:"cannotequipreason"`
	BreakerType                 int                       `json:"breakertype"`
	BreakerTypeHash             int                       `json:"breakertypehash"`
	Energy                      DestinyItemInstanceEnergy `json:"energy"`
}
type DestinyStat struct {
	StatHash int `json:"stathash"`
	Value    int `json:"value"`
}
type EquipFailureReason struct {
	EquipFailureReason int `json:"equipfailurereason"`
}
type DestinyItemInstanceEnergy struct {
	EnergyTypeHash int `json:"energytypehash"`
	EnergyType     int `json:"energytype"`
	EnergyCapacity int `json:"energycapacity"`
	EnergyUsed     int `json:"energyused"`
	EnergyUnused   int `json:"energyunused"`
}
type DestinyUnlockDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent struct {
	Data     DestinyItemInstanceComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type DestinyItemRenderComponent struct {
	UseCustomDyes bool `json:"usecustomdyes"`
	ArtRegions    int  `json:"artregions"`
}
type DictionaryComponentResponseOfint64AndDestinyItemRenderComponent struct {
	Data     DestinyItemRenderComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type DestinyItemStatsComponent struct {
	Stats DestinyStat `json:"stats"`
}
type DictionaryComponentResponseOfint64AndDestinyItemStatsComponent struct {
	Data     DestinyItemStatsComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DestinyItemSocketsComponent struct {
	Sockets []DestinyItemSocketState `json:"sockets"`
}
type DestinyItemSocketState struct {
	PlugHash          int   `json:"plughash"`
	IsEnabled         bool  `json:"isenabled"`
	IsVisible         bool  `json:"isvisible"`
	EnableFailIndexes []int `json:"enablefailindexes"`
}
type DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent struct {
	Data     DestinyItemSocketsComponent `json:"data"`
	Privacy  int                         `json:"privacy"`
	Disabled bool                        `json:"disabled"`
}
type DestinyItemReusablePlugsComponent struct {
	Plugs []DestinyItemPlugBase `json:"plugs"`
}
type DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent struct {
	Data     DestinyItemReusablePlugsComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DestinyItemPlugObjectivesComponent struct {
	ObjectivesPerPlug []DestinyObjectiveProgress `json:"objectivesperplug"`
}
type DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent struct {
	Data     DestinyItemPlugObjectivesComponent `json:"data"`
	Privacy  int                                `json:"privacy"`
	Disabled bool                               `json:"disabled"`
}
type DestinyItemTalentGridComponent struct {
	TalentGridHash  int                 `json:"talentgridhash"`
	Nodes           []DestinyTalentNode `json:"nodes"`
	IsGridComplete  bool                `json:"isgridcomplete"`
	GridProgression DestinyProgression  `json:"gridprogression"`
}
type DestinyTalentNode struct {
	NodeIndex           int                          `json:"nodeindex"`
	NodeHash            int                          `json:"nodehash"`
	State               int                          `json:"state"`
	IsActivated         bool                         `json:"isactivated"`
	StepIndex           int                          `json:"stepindex"`
	MaterialsToUpgrade  []DestinyMaterialRequirement `json:"materialstoupgrade"`
	ActivationGridLevel int                          `json:"activationgridlevel"`
	ProgressPercent     float64                      `json:"progresspercent"`
	Hidden              bool                         `json:"hidden"`
	NodeStatsBlock      DestinyTalentNodeStatBlock   `json:"nodestatsblock"`
}
type DestinyTalentNodeState struct {
	DestinyTalentNodeState int `json:"destinytalentnodestate"`
}
type DestinyTalentNodeStatBlock struct {
	CurrentStepStats []DestinyStat `json:"currentstepstats"`
	NextStepStats    []DestinyStat `json:"nextstepstats"`
}
type DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent struct {
	Data     DestinyItemTalentGridComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyItemPlugComponent struct {
	PlugObjectives    []DestinyObjectiveProgress `json:"plugobjectives"`
	PlugItemHash      int                        `json:"plugitemhash"`
	CanInsert         bool                       `json:"caninsert"`
	Enabled           bool                       `json:"enabled"`
	InsertFailIndexes []int                      `json:"insertfailindexes"`
	EnableFailIndexes []int                      `json:"enablefailindexes"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent struct {
	Data     DestinyItemPlugComponent `json:"data"`
	Privacy  int                      `json:"privacy"`
	Disabled bool                     `json:"disabled"`
}
type DestinyCurrenciesComponent struct {
	ItemQuantities int `json:"itemquantities"`
}
type DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent struct {
	Data     DestinyCurrenciesComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type DestinyCharacterResponse struct {
	Inventory                 SingleComponentResponseOfDestinyInventoryComponent            `json:"inventory"`
	Character                 SingleComponentResponseOfDestinyCharacterComponent            `json:"character"`
	Progressions              SingleComponentResponseOfDestinyCharacterProgressionComponent `json:"progressions"`
	RenderData                SingleComponentResponseOfDestinyCharacterRenderComponent      `json:"renderdata"`
	Activities                SingleComponentResponseOfDestinyCharacterActivitiesComponent  `json:"activities"`
	Equipment                 SingleComponentResponseOfDestinyInventoryComponent            `json:"equipment"`
	Kiosks                    SingleComponentResponseOfDestinyKiosksComponent               `json:"kiosks"`
	PlugSets                  SingleComponentResponseOfDestinyPlugSetsComponent             `json:"plugsets"`
	PresentationNodes         SingleComponentResponseOfDestinyPresentationNodesComponent    `json:"presentationnodes"`
	Records                   SingleComponentResponseOfDestinyCharacterRecordsComponent     `json:"records"`
	Collectibles              SingleComponentResponseOfDestinyCollectiblesComponent         `json:"collectibles"`
	ItemComponents            DestinyItemComponentSetOfint64                                `json:"itemcomponents"`
	UninstancedItemComponents DestinyBaseItemComponentSetOfuint32                           `json:"uninstanceditemcomponents"`
	CurrencyLookups           SingleComponentResponseOfDestinyCurrenciesComponent           `json:"currencylookups"`
}
type SingleComponentResponseOfDestinyCharacterComponent struct {
	Data     DestinyCharacterComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterProgressionComponent struct {
	Data     DestinyCharacterProgressionComponent `json:"data"`
	Privacy  int                                  `json:"privacy"`
	Disabled bool                                 `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterRenderComponent struct {
	Data     DestinyCharacterRenderComponent `json:"data"`
	Privacy  int                             `json:"privacy"`
	Disabled bool                            `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterActivitiesComponent struct {
	Data     DestinyCharacterActivitiesComponent `json:"data"`
	Privacy  int                                 `json:"privacy"`
	Disabled bool                                `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterRecordsComponent struct {
	Data     DestinyCharacterRecordsComponent `json:"data"`
	Privacy  int                              `json:"privacy"`
	Disabled bool                             `json:"disabled"`
}
type SingleComponentResponseOfDestinyCollectiblesComponent struct {
	Data     DestinyCollectiblesComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type SingleComponentResponseOfDestinyCurrenciesComponent struct {
	Data     DestinyCurrenciesComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type ClanBannerSource struct {
	ClanBannerSource interface{} `json:"clanbannersource"`
}
type ClanBannerDecal struct {
	Identifier     string `json:"identifier"`
	ForegroundPath string `json:"foregroundpath"`
	BackgroundPath string `json:"backgroundpath"`
}
type DestinyItemResponse struct {
	CharacterId    int                                                         `json:"characterid"`
	Item           SingleComponentResponseOfDestinyItemComponent               `json:"item"`
	Instance       SingleComponentResponseOfDestinyItemInstanceComponent       `json:"instance"`
	Objectives     SingleComponentResponseOfDestinyItemObjectivesComponent     `json:"objectives"`
	Perks          SingleComponentResponseOfDestinyItemPerksComponent          `json:"perks"`
	RenderData     SingleComponentResponseOfDestinyItemRenderComponent         `json:"renderdata"`
	Stats          SingleComponentResponseOfDestinyItemStatsComponent          `json:"stats"`
	TalentGrid     SingleComponentResponseOfDestinyItemTalentGridComponent     `json:"talentgrid"`
	Sockets        SingleComponentResponseOfDestinyItemSocketsComponent        `json:"sockets"`
	ReusablePlugs  SingleComponentResponseOfDestinyItemReusablePlugsComponent  `json:"reusableplugs"`
	PlugObjectives SingleComponentResponseOfDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
}
type SingleComponentResponseOfDestinyItemComponent struct {
	Data     DestinyItemComponent `json:"data"`
	Privacy  int                  `json:"privacy"`
	Disabled bool                 `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemInstanceComponent struct {
	Data     DestinyItemInstanceComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemObjectivesComponent struct {
	Data     DestinyItemObjectivesComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemPerksComponent struct {
	Data     DestinyItemPerksComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemRenderComponent struct {
	Data     DestinyItemRenderComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemStatsComponent struct {
	Data     DestinyItemStatsComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemTalentGridComponent struct {
	Data     DestinyItemTalentGridComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemSocketsComponent struct {
	Data     DestinyItemSocketsComponent `json:"data"`
	Privacy  int                         `json:"privacy"`
	Disabled bool                        `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemReusablePlugsComponent struct {
	Data     DestinyItemReusablePlugsComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemPlugObjectivesComponent struct {
	Data     DestinyItemPlugObjectivesComponent `json:"data"`
	Privacy  int                                `json:"privacy"`
	Disabled bool                               `json:"disabled"`
}
type DestinyVendorFilter struct {
	DestinyVendorFilter int `json:"destinyvendorfilter"`
}
type DestinyVendorsResponse struct {
	VendorGroups    SingleComponentResponseOfDestinyVendorGroupComponent                            `json:"vendorgroups"`
	Vendors         DictionaryComponentResponseOfuint32AndDestinyVendorComponent                    `json:"vendors"`
	Categories      DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent          `json:"categories"`
	Sales           DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent `json:"sales"`
	ItemComponents  DestinyItemComponentSetOfint32                                                  `json:"itemcomponents"`
	CurrencyLookups SingleComponentResponseOfDestinyCurrenciesComponent                             `json:"currencylookups"`
	StringVariables SingleComponentResponseOfDestinyStringVariablesComponent                        `json:"stringvariables"`
}
type DestinyVendorGroupComponent struct {
	Groups []DestinyVendorGroup `json:"groups"`
}
type DestinyVendorGroup struct {
	VendorGroupHash int   `json:"vendorgrouphash"`
	VendorHashes    []int `json:"vendorhashes"`
}
type SingleComponentResponseOfDestinyVendorGroupComponent struct {
	Data     DestinyVendorGroupComponent `json:"data"`
	Privacy  int                         `json:"privacy"`
	Disabled bool                        `json:"disabled"`
}
type DestinyVendorBaseComponent struct {
	VendorHash      int    `json:"vendorhash"`
	NextRefreshDate string `json:"nextrefreshdate"`
	Enabled         bool   `json:"enabled"`
}
type DestinyVendorComponent struct {
	CanPurchase         bool               `json:"canpurchase"`
	Progression         DestinyProgression `json:"progression"`
	VendorLocationIndex int                `json:"vendorlocationindex"`
	SeasonalRank        int                `json:"seasonalrank"`
	VendorHash          int                `json:"vendorhash"`
	NextRefreshDate     string             `json:"nextrefreshdate"`
	Enabled             bool               `json:"enabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyVendorComponent struct {
	Data     DestinyVendorComponent `json:"data"`
	Privacy  int                    `json:"privacy"`
	Disabled bool                   `json:"disabled"`
}
type DestinyVendorCategoriesComponent struct {
	Categories []DestinyVendorCategory `json:"categories"`
}
type DestinyVendorCategory struct {
	DisplayCategoryIndex int   `json:"displaycategoryindex"`
	ItemIndexes          []int `json:"itemindexes"`
}
type DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent struct {
	Data     DestinyVendorCategoriesComponent `json:"data"`
	Privacy  int                              `json:"privacy"`
	Disabled bool                             `json:"disabled"`
}
type DestinyVendorSaleItemBaseComponent struct {
	VendorItemIndex         int                   `json:"vendoritemindex"`
	ItemHash                int                   `json:"itemhash"`
	OverrideStyleItemHash   int                   `json:"overridestyleitemhash"`
	Quantity                int                   `json:"quantity"`
	Costs                   []DestinyItemQuantity `json:"costs"`
	OverrideNextRefreshDate string                `json:"overridenextrefreshdate"`
	ApiPurchasable          bool                  `json:"apipurchasable"`
}
type DestinyVendorSaleItemComponent struct {
	SaleStatus              int                   `json:"salestatus"`
	RequiredUnlocks         []int                 `json:"requiredunlocks"`
	UnlockStatuses          []DestinyUnlockStatus `json:"unlockstatuses"`
	FailureIndexes          []int                 `json:"failureindexes"`
	Augments                int                   `json:"augments"`
	ItemValueVisibility     []bool                `json:"itemvaluevisibility"`
	VendorItemIndex         int                   `json:"vendoritemindex"`
	ItemHash                int                   `json:"itemhash"`
	OverrideStyleItemHash   int                   `json:"overridestyleitemhash"`
	Quantity                int                   `json:"quantity"`
	Costs                   []DestinyItemQuantity `json:"costs"`
	OverrideNextRefreshDate string                `json:"overridenextrefreshdate"`
	ApiPurchasable          bool                  `json:"apipurchasable"`
}
type VendorItemStatus struct {
	VendorItemStatus int `json:"vendoritemstatus"`
}
type DestinyUnlockStatus struct {
	UnlockHash int  `json:"unlockhash"`
	IsSet      bool `json:"isset"`
}
type DestinyVendorItemState struct {
	DestinyVendorItemState int `json:"destinyvendoritemstate"`
}
type DestinyVendorSaleItemSetComponentOfDestinyVendorSaleItemComponent struct {
	SaleItems DestinyVendorSaleItemComponent `json:"saleitems"`
}
type PersonalDestinyVendorSaleItemSetComponent struct {
	SaleItems DestinyVendorSaleItemComponent `json:"saleitems"`
}
type DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent struct {
	Data     PersonalDestinyVendorSaleItemSetComponent `json:"data"`
	Privacy  int                                       `json:"privacy"`
	Disabled bool                                      `json:"disabled"`
}
type DestinyBaseItemComponentSetOfint32 struct {
	Objectives DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent `json:"objectives"`
	Perks      DictionaryComponentResponseOfint32AndDestinyItemPerksComponent      `json:"perks"`
}
type DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent struct {
	Data     DestinyItemObjectivesComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemPerksComponent struct {
	Data     DestinyItemPerksComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DestinyItemComponentSetOfint32 struct {
	Instances      DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent       `json:"instances"`
	RenderData     DictionaryComponentResponseOfint32AndDestinyItemRenderComponent         `json:"renderdata"`
	Stats          DictionaryComponentResponseOfint32AndDestinyItemStatsComponent          `json:"stats"`
	Sockets        DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent        `json:"sockets"`
	ReusablePlugs  DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent  `json:"reusableplugs"`
	PlugObjectives DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
	TalentGrids    DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent     `json:"talentgrids"`
	PlugStates     DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent          `json:"plugstates"`
	Objectives     DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent     `json:"objectives"`
	Perks          DictionaryComponentResponseOfint32AndDestinyItemPerksComponent          `json:"perks"`
}
type DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent struct {
	Data     DestinyItemInstanceComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemRenderComponent struct {
	Data     DestinyItemRenderComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemStatsComponent struct {
	Data     DestinyItemStatsComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent struct {
	Data     DestinyItemSocketsComponent `json:"data"`
	Privacy  int                         `json:"privacy"`
	Disabled bool                        `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent struct {
	Data     DestinyItemReusablePlugsComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent struct {
	Data     DestinyItemPlugObjectivesComponent `json:"data"`
	Privacy  int                                `json:"privacy"`
	Disabled bool                               `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent struct {
	Data     DestinyItemTalentGridComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyVendorResponse struct {
	Vendor          SingleComponentResponseOfDestinyVendorComponent                     `json:"vendor"`
	Categories      SingleComponentResponseOfDestinyVendorCategoriesComponent           `json:"categories"`
	Sales           DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent `json:"sales"`
	ItemComponents  DestinyItemComponentSetOfint32                                      `json:"itemcomponents"`
	CurrencyLookups SingleComponentResponseOfDestinyCurrenciesComponent                 `json:"currencylookups"`
	StringVariables SingleComponentResponseOfDestinyStringVariablesComponent            `json:"stringvariables"`
}
type SingleComponentResponseOfDestinyVendorComponent struct {
	Data     DestinyVendorComponent `json:"data"`
	Privacy  int                    `json:"privacy"`
	Disabled bool                   `json:"disabled"`
}
type SingleComponentResponseOfDestinyVendorCategoriesComponent struct {
	Data     DestinyVendorCategoriesComponent `json:"data"`
	Privacy  int                              `json:"privacy"`
	Disabled bool                             `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent struct {
	Data     DestinyVendorSaleItemComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyPublicVendorsResponse struct {
	VendorGroups    SingleComponentResponseOfDestinyVendorGroupComponent                          `json:"vendorgroups"`
	Vendors         DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent            `json:"vendors"`
	Categories      DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent        `json:"categories"`
	Sales           DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent `json:"sales"`
	StringVariables SingleComponentResponseOfDestinyStringVariablesComponent                      `json:"stringvariables"`
}
type DestinyPublicVendorComponent struct {
	VendorHash      int    `json:"vendorhash"`
	NextRefreshDate string `json:"nextrefreshdate"`
	Enabled         bool   `json:"enabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent struct {
	Data     DestinyPublicVendorComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type DestinyPublicVendorSaleItemComponent struct {
	VendorItemIndex         int                   `json:"vendoritemindex"`
	ItemHash                int                   `json:"itemhash"`
	OverrideStyleItemHash   int                   `json:"overridestyleitemhash"`
	Quantity                int                   `json:"quantity"`
	Costs                   []DestinyItemQuantity `json:"costs"`
	OverrideNextRefreshDate string                `json:"overridenextrefreshdate"`
	ApiPurchasable          bool                  `json:"apipurchasable"`
}
type DestinyVendorSaleItemSetComponentOfDestinyPublicVendorSaleItemComponent struct {
	SaleItems DestinyPublicVendorSaleItemComponent `json:"saleitems"`
}
type PublicDestinyVendorSaleItemSetComponent struct {
	SaleItems DestinyPublicVendorSaleItemComponent `json:"saleitems"`
}
type DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent struct {
	Data     PublicDestinyVendorSaleItemSetComponent `json:"data"`
	Privacy  int                                     `json:"privacy"`
	Disabled bool                                    `json:"disabled"`
}
type DestinyCollectibleNodeDetailResponse struct {
	Collectibles              SingleComponentResponseOfDestinyCollectiblesComponent `json:"collectibles"`
	CollectibleItemComponents DestinyItemComponentSetOfuint32                       `json:"collectibleitemcomponents"`
}
type DestinyItemComponentSetOfuint32 struct {
	Instances      DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent       `json:"instances"`
	RenderData     DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent         `json:"renderdata"`
	Stats          DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent          `json:"stats"`
	Sockets        DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent        `json:"sockets"`
	ReusablePlugs  DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent  `json:"reusableplugs"`
	PlugObjectives DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
	TalentGrids    DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent     `json:"talentgrids"`
	PlugStates     DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent           `json:"plugstates"`
	Objectives     DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent     `json:"objectives"`
	Perks          DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent          `json:"perks"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent struct {
	Data     DestinyItemInstanceComponent `json:"data"`
	Privacy  int                          `json:"privacy"`
	Disabled bool                         `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent struct {
	Data     DestinyItemRenderComponent `json:"data"`
	Privacy  int                        `json:"privacy"`
	Disabled bool                       `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent struct {
	Data     DestinyItemStatsComponent `json:"data"`
	Privacy  int                       `json:"privacy"`
	Disabled bool                      `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent struct {
	Data     DestinyItemSocketsComponent `json:"data"`
	Privacy  int                         `json:"privacy"`
	Disabled bool                        `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent struct {
	Data     DestinyItemReusablePlugsComponent `json:"data"`
	Privacy  int                               `json:"privacy"`
	Disabled bool                              `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent struct {
	Data     DestinyItemPlugObjectivesComponent `json:"data"`
	Privacy  int                                `json:"privacy"`
	Disabled bool                               `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent struct {
	Data     DestinyItemTalentGridComponent `json:"data"`
	Privacy  int                            `json:"privacy"`
	Disabled bool                           `json:"disabled"`
}
type DestinyActionRequest struct {
	MembershipType int `json:"membershiptype"`
}
type DestinyCharacterActionRequest struct {
	CharacterId    int `json:"characterid"`
	MembershipType int `json:"membershiptype"`
}
type DestinyItemActionRequest struct {
	ItemId         int `json:"itemid"`
	CharacterId    int `json:"characterid"`
	MembershipType int `json:"membershiptype"`
}
type DestinyItemTransferRequest struct {
	ItemReferenceHash int  `json:"itemreferencehash"`
	StackSize         int  `json:"stacksize"`
	TransferToVault   bool `json:"transfertovault"`
	ItemId            int  `json:"itemid"`
	CharacterId       int  `json:"characterid"`
	MembershipType    int  `json:"membershiptype"`
}
type DestinyPostmasterTransferRequest struct {
	ItemReferenceHash int `json:"itemreferencehash"`
	StackSize         int `json:"stacksize"`
	ItemId            int `json:"itemid"`
	CharacterId       int `json:"characterid"`
	MembershipType    int `json:"membershiptype"`
}
type DestinyEquipItemResults struct {
	EquipResults []DestinyEquipItemResult `json:"equipresults"`
}
type DestinyEquipItemResult struct {
	ItemInstanceId int `json:"iteminstanceid"`
	EquipStatus    int `json:"equipstatus"`
}
type DestinyItemSetActionRequest struct {
	ItemIds        []int `json:"itemids"`
	CharacterId    int   `json:"characterid"`
	MembershipType int   `json:"membershiptype"`
}
type DestinyItemStateRequest struct {
	State          bool `json:"state"`
	ItemId         int  `json:"itemid"`
	CharacterId    int  `json:"characterid"`
	MembershipType int  `json:"membershiptype"`
}
type InventoryChangedResponse struct {
	AddedInventoryItems   []DestinyItemComponent `json:"addedinventoryitems"`
	RemovedInventoryItems []DestinyItemComponent `json:"removedinventoryitems"`
}
type DestinyItemChangeResponse struct {
	Item                  DestinyItemResponse    `json:"item"`
	AddedInventoryItems   []DestinyItemComponent `json:"addedinventoryitems"`
	RemovedInventoryItems []DestinyItemComponent `json:"removedinventoryitems"`
}
type DestinyInsertPlugsActionRequest struct {
	ActionToken    string                         `json:"actiontoken"`
	ItemInstanceId int                            `json:"iteminstanceid"`
	Plug           DestinyInsertPlugsRequestEntry `json:"plug"`
	CharacterId    int                            `json:"characterid"`
	MembershipType int                            `json:"membershiptype"`
}
type DestinyInsertPlugsRequestEntry struct {
	SocketIndex     int `json:"socketindex"`
	SocketArrayType int `json:"socketarraytype"`
	PlugItemHash    int `json:"plugitemhash"`
}
type DestinySocketArrayType struct {
	DestinySocketArrayType int `json:"destinysocketarraytype"`
}
type DestinyPostGameCarnageReportData struct {
	Period             string                                  `json:"period"`
	StartingPhaseIndex int                                     `json:"startingphaseindex"`
	ActivityDetails    DestinyHistoricalStatsActivity          `json:"activitydetails"`
	Entries            []DestinyPostGameCarnageReportEntry     `json:"entries"`
	Teams              []DestinyPostGameCarnageReportTeamEntry `json:"teams"`
}
type DestinyHistoricalStatsActivity struct {
	ReferenceId          int   `json:"referenceid"`
	DirectorActivityHash int   `json:"directoractivityhash"`
	InstanceId           int   `json:"instanceid"`
	Mode                 int   `json:"mode"`
	Modes                []int `json:"modes"`
	IsPrivate            bool  `json:"isprivate"`
	MembershipType       int   `json:"membershiptype"`
}
type DestinyPostGameCarnageReportEntry struct {
	Standing    int                                      `json:"standing"`
	Score       DestinyHistoricalStatsValue              `json:"score"`
	Player      DestinyPlayer                            `json:"player"`
	CharacterId int                                      `json:"characterid"`
	Values      DestinyHistoricalStatsValue              `json:"values"`
	Extended    DestinyPostGameCarnageReportExtendedData `json:"extended"`
}
type DestinyHistoricalStatsValue struct {
	StatId     string                          `json:"statid"`
	Basic      DestinyHistoricalStatsValuePair `json:"basic"`
	Pga        DestinyHistoricalStatsValuePair `json:"pga"`
	Weighted   DestinyHistoricalStatsValuePair `json:"weighted"`
	ActivityId int                             `json:"activityid"`
}
type DestinyHistoricalStatsValuePair struct {
	Value        float64 `json:"value"`
	DisplayValue string  `json:"displayvalue"`
}
type DestinyPlayer struct {
	DestinyUserInfo   UserInfoCard `json:"destinyuserinfo"`
	CharacterClass    string       `json:"characterclass"`
	ClassHash         int          `json:"classhash"`
	RaceHash          int          `json:"racehash"`
	GenderHash        int          `json:"genderhash"`
	CharacterLevel    int          `json:"characterlevel"`
	LightLevel        int          `json:"lightlevel"`
	BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
	ClanName          string       `json:"clanname"`
	ClanTag           string       `json:"clantag"`
	EmblemHash        int          `json:"emblemhash"`
}
type DestinyPostGameCarnageReportExtendedData struct {
	Weapons []DestinyHistoricalWeaponStats `json:"weapons"`
	Values  DestinyHistoricalStatsValue    `json:"values"`
}
type DestinyHistoricalWeaponStats struct {
	ReferenceId int                         `json:"referenceid"`
	Values      DestinyHistoricalStatsValue `json:"values"`
}
type DestinyPostGameCarnageReportTeamEntry struct {
	TeamId   int                         `json:"teamid"`
	Standing DestinyHistoricalStatsValue `json:"standing"`
	Score    DestinyHistoricalStatsValue `json:"score"`
	TeamName string                      `json:"teamname"`
}
type DestinyReportOffensePgcrRequest struct {
	ReasonCategoryHashes []int `json:"reasoncategoryhashes"`
	ReasonHashes         []int `json:"reasonhashes"`
	OffendingCharacterId int   `json:"offendingcharacterid"`
}
type DestinyReportReasonCategoryDefinition struct {
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Reasons           DestinyReportReasonDefinition      `json:"reasons"`
	Hash              int                                `json:"hash"`
	Index             int                                `json:"index"`
	Redacted          bool                               `json:"redacted"`
}
type DestinyReportReasonDefinition struct {
	ReasonHash        int                                `json:"reasonhash"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyHistoricalStatsDefinition struct {
	StatId          string `json:"statid"`
	Group           int    `json:"group"`
	PeriodTypes     []int  `json:"periodtypes"`
	Modes           []int  `json:"modes"`
	Category        int    `json:"category"`
	StatName        string `json:"statname"`
	StatNameAbbr    string `json:"statnameabbr"`
	StatDescription string `json:"statdescription"`
	UnitType        int    `json:"unittype"`
	IconImage       string `json:"iconimage"`
	MergeMethod     int    `json:"mergemethod"`
	UnitLabel       string `json:"unitlabel"`
	Weight          int    `json:"weight"`
	MedalTierHash   int    `json:"medaltierhash"`
}
type DestinyStatsGroupType struct {
	DestinyStatsGroupType int `json:"destinystatsgrouptype"`
}
type PeriodTypeList struct {
	PeriodTypeList int `json:"periodtypelist"`
}
type DestinyActivityModeTypeList struct {
	DestinyActivityModeTypeList int `json:"destinyactivitymodetypelist"`
}
type DestinyStatsCategoryType struct {
	DestinyStatsCategoryType int `json:"destinystatscategorytype"`
}
type UnitType struct {
	UnitType int `json:"unittype"`
}
type DestinyStatsMergeMethod struct {
	DestinyStatsMergeMethod int `json:"destinystatsmergemethod"`
}
type DestinyLeaderboard struct {
	StatId  string                    `json:"statid"`
	Entries []DestinyLeaderboardEntry `json:"entries"`
}
type DestinyLeaderboardEntry struct {
	Rank        int                         `json:"rank"`
	Player      DestinyPlayer               `json:"player"`
	CharacterId int                         `json:"characterid"`
	Value       DestinyHistoricalStatsValue `json:"value"`
}
type DestinyLeaderboardResults struct {
	FocusMembershipId int `json:"focusmembershipid"`
	FocusCharacterId  int `json:"focuscharacterid"`
}
type DestinyClanAggregateStat struct {
	Mode   int                         `json:"mode"`
	StatId string                      `json:"statid"`
	Value  DestinyHistoricalStatsValue `json:"value"`
}
type DestinyEntitySearchResult struct {
	SuggestedWords []string                                    `json:"suggestedwords"`
	Results        SearchResultOfDestinyEntitySearchResultItem `json:"results"`
}
type DestinyEntitySearchResultItem struct {
	Hash              int                                `json:"hash"`
	EntityType        string                             `json:"entitytype"`
	DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
	Weight            float64                            `json:"weight"`
}
type SearchResultOfDestinyEntitySearchResultItem struct {
	Results                      []DestinyEntitySearchResultItem `json:"results"`
	TotalResults                 int                             `json:"totalresults"`
	HasMore                      bool                            `json:"hasmore"`
	Query                        PagedQuery                      `json:"query"`
	ReplacementContinuationToken string                          `json:"replacementcontinuationtoken"`
	UseTotalResults              bool                            `json:"usetotalresults"`
}
type PeriodType struct {
	PeriodType int `json:"periodtype"`
}
type DestinyHistoricalStatsByPeriod struct {
	AllTime      DestinyHistoricalStatsValue         `json:"alltime"`
	AllTimeTier1 DestinyHistoricalStatsValue         `json:"alltimetier1"`
	AllTimeTier2 DestinyHistoricalStatsValue         `json:"alltimetier2"`
	AllTimeTier3 DestinyHistoricalStatsValue         `json:"alltimetier3"`
	Daily        []DestinyHistoricalStatsPeriodGroup `json:"daily"`
	Monthly      []DestinyHistoricalStatsPeriodGroup `json:"monthly"`
}
type DestinyHistoricalStatsPeriodGroup struct {
	Period          string                         `json:"period"`
	ActivityDetails DestinyHistoricalStatsActivity `json:"activitydetails"`
	Values          DestinyHistoricalStatsValue    `json:"values"`
}
type DestinyHistoricalStatsResults struct {
	DestinyHistoricalStatsResults interface{} `json:"destinyhistoricalstatsresults"`
}
type DestinyHistoricalStatsAccountResult struct {
	MergedDeletedCharacters DestinyHistoricalStatsWithMerged     `json:"mergeddeletedcharacters"`
	MergedAllCharacters     DestinyHistoricalStatsWithMerged     `json:"mergedallcharacters"`
	Characters              []DestinyHistoricalStatsPerCharacter `json:"characters"`
}
type DestinyHistoricalStatsWithMerged struct {
	Results DestinyHistoricalStatsByPeriod `json:"results"`
	Merged  DestinyHistoricalStatsByPeriod `json:"merged"`
}
type DestinyHistoricalStatsPerCharacter struct {
	CharacterId int                            `json:"characterid"`
	Deleted     bool                           `json:"deleted"`
	Results     DestinyHistoricalStatsByPeriod `json:"results"`
	Merged      DestinyHistoricalStatsByPeriod `json:"merged"`
}
type DestinyActivityHistoryResults struct {
	Activities []DestinyHistoricalStatsPeriodGroup `json:"activities"`
}
type DestinyHistoricalWeaponStatsData struct {
	Weapons []DestinyHistoricalWeaponStats `json:"weapons"`
}
type DestinyAggregateActivityResults struct {
	Activities []DestinyAggregateActivityStats `json:"activities"`
}
type DestinyAggregateActivityStats struct {
	ActivityHash int                         `json:"activityhash"`
	Values       DestinyHistoricalStatsValue `json:"values"`
}
type DestinyMilestoneContent struct {
	About          string                                `json:"about"`
	Status         string                                `json:"status"`
	Tips           []string                              `json:"tips"`
	ItemCategories []DestinyMilestoneContentItemCategory `json:"itemcategories"`
}
type DestinyMilestoneContentItemCategory struct {
	Title      string `json:"title"`
	ItemHashes []int  `json:"itemhashes"`
}
type DestinyPublicMilestone struct {
	MilestoneHash   int                                       `json:"milestonehash"`
	AvailableQuests []DestinyPublicMilestoneQuest             `json:"availablequests"`
	Activities      []DestinyPublicMilestoneChallengeActivity `json:"activities"`
	VendorHashes    []int                                     `json:"vendorhashes"`
	Vendors         []DestinyPublicMilestoneVendor            `json:"vendors"`
	StartDate       string                                    `json:"startdate"`
	EndDate         string                                    `json:"enddate"`
	Order           int                                       `json:"order"`
}
type DestinyPublicMilestoneQuest struct {
	QuestItemHash int                               `json:"questitemhash"`
	Activity      DestinyPublicMilestoneActivity    `json:"activity"`
	Challenges    []DestinyPublicMilestoneChallenge `json:"challenges"`
}
type DestinyPublicMilestoneActivity struct {
	ActivityHash     int                                     `json:"activityhash"`
	ModifierHashes   []int                                   `json:"modifierhashes"`
	Variants         []DestinyPublicMilestoneActivityVariant `json:"variants"`
	ActivityModeHash int                                     `json:"activitymodehash"`
	ActivityModeType int                                     `json:"activitymodetype"`
}
type DestinyPublicMilestoneActivityVariant struct {
	ActivityHash     int `json:"activityhash"`
	ActivityModeHash int `json:"activitymodehash"`
	ActivityModeType int `json:"activitymodetype"`
}
type DestinyPublicMilestoneChallenge struct {
	ObjectiveHash int `json:"objectivehash"`
	ActivityHash  int `json:"activityhash"`
}
type DestinyPublicMilestoneChallengeActivity struct {
	ActivityHash             int   `json:"activityhash"`
	ChallengeObjectiveHashes []int `json:"challengeobjectivehashes"`
	ModifierHashes           []int `json:"modifierhashes"`
	LoadoutRequirementIndex  int   `json:"loadoutrequirementindex"`
	PhaseHashes              []int `json:"phasehashes"`
	BooleanActivityOptions   bool  `json:"booleanactivityoptions"`
}
type DestinyPublicMilestoneVendor struct {
	VendorHash      int `json:"vendorhash"`
	PreviewItemHash int `json:"previewitemhash"`
}
type AwaInitializeResponse struct {
	CorrelationId string `json:"correlationid"`
	SentToSelf    bool   `json:"senttoself"`
}
type AwaPermissionRequested struct {
	Type           int `json:"type"`
	AffectedItemId int `json:"affecteditemid"`
	MembershipType int `json:"membershiptype"`
	CharacterId    int `json:"characterid"`
}
type AwaType struct {
	AwaType int `json:"awatype"`
}
type AwaUserResponse struct {
	Selection     int      `json:"selection"`
	CorrelationId string   `json:"correlationid"`
	Nonce         []string `json:"nonce"`
}
type AwaUserSelection struct {
	AwaUserSelection int `json:"awauserselection"`
}
type AwaAuthorizationResult struct {
	UserSelection       int    `json:"userselection"`
	ResponseReason      int    `json:"responsereason"`
	DeveloperNote       string `json:"developernote"`
	ActionToken         string `json:"actiontoken"`
	MaximumNumberOfUses int    `json:"maximumnumberofuses"`
	ValidUntil          string `json:"validuntil"`
	Type                int    `json:"type"`
	MembershipType      int    `json:"membershiptype"`
}
type AwaResponseReason struct {
	AwaResponseReason int `json:"awaresponsereason"`
}
type CommunityContentSortMode struct {
	CommunityContentSortMode int `json:"communitycontentsortmode"`
}
type TrendingCategories struct {
	Categories []TrendingCategory `json:"categories"`
}
type TrendingCategory struct {
	CategoryName string                      `json:"categoryname"`
	Entries      SearchResultOfTrendingEntry `json:"entries"`
	CategoryId   string                      `json:"categoryid"`
}
type TrendingEntry struct {
	Weight       float64         `json:"weight"`
	IsFeatured   bool            `json:"isfeatured"`
	Identifier   string          `json:"identifier"`
	EntityType   int             `json:"entitytype"`
	DisplayName  string          `json:"displayname"`
	Tagline      string          `json:"tagline"`
	Image        string          `json:"image"`
	StartDate    string          `json:"startdate"`
	EndDate      string          `json:"enddate"`
	Link         string          `json:"link"`
	WebmVideo    string          `json:"webmvideo"`
	Mp4Video     string          `json:"mp4video"`
	FeatureImage string          `json:"featureimage"`
	Items        []TrendingEntry `json:"items"`
	CreationDate string          `json:"creationdate"`
}
type TrendingEntryType struct {
	TrendingEntryType int `json:"trendingentrytype"`
}
type SearchResultOfTrendingEntry struct {
	Results                      []TrendingEntry `json:"results"`
	TotalResults                 int             `json:"totalresults"`
	HasMore                      bool            `json:"hasmore"`
	Query                        PagedQuery      `json:"query"`
	ReplacementContinuationToken string          `json:"replacementcontinuationtoken"`
	UseTotalResults              bool            `json:"usetotalresults"`
}
type TrendingDetail struct {
	Identifier      string                         `json:"identifier"`
	EntityType      int                            `json:"entitytype"`
	News            TrendingEntryNews              `json:"news"`
	Support         TrendingEntrySupportArticle    `json:"support"`
	DestinyItem     TrendingEntryDestinyItem       `json:"destinyitem"`
	DestinyActivity TrendingEntryDestinyActivity   `json:"destinyactivity"`
	DestinyRitual   TrendingEntryDestinyRitual     `json:"destinyritual"`
	Creation        TrendingEntryCommunityCreation `json:"creation"`
}
type TrendingEntryNews struct {
	Article ContentItemPublicContract `json:"article"`
}
type TrendingEntrySupportArticle struct {
	Article ContentItemPublicContract `json:"article"`
}
type TrendingEntryDestinyItem struct {
	ItemHash int `json:"itemhash"`
}
type TrendingEntryDestinyActivity struct {
	ActivityHash int                         `json:"activityhash"`
	Status       DestinyPublicActivityStatus `json:"status"`
}
type DestinyPublicActivityStatus struct {
	ChallengeObjectiveHashes []int                 `json:"challengeobjectivehashes"`
	ModifierHashes           []int                 `json:"modifierhashes"`
	RewardTooltipItems       []DestinyItemQuantity `json:"rewardtooltipitems"`
}
type TrendingEntryDestinyRitual struct {
	Image            string                  `json:"image"`
	Icon             string                  `json:"icon"`
	Title            string                  `json:"title"`
	Subtitle         string                  `json:"subtitle"`
	DateStart        string                  `json:"datestart"`
	DateEnd          string                  `json:"dateend"`
	MilestoneDetails DestinyPublicMilestone  `json:"milestonedetails"`
	EventContent     DestinyMilestoneContent `json:"eventcontent"`
}
type TrendingEntryCommunityCreation struct {
	Media              string `json:"media"`
	Title              string `json:"title"`
	Author             string `json:"author"`
	AuthorMembershipId int    `json:"authormembershipid"`
	PostId             int    `json:"postid"`
	Body               string `json:"body"`
	Upvotes            int    `json:"upvotes"`
}
type FireteamDateRange struct {
	FireteamDateRange int `json:"fireteamdaterange"`
}
type FireteamPlatform struct {
	FireteamPlatform int `json:"fireteamplatform"`
}
type FireteamPublicSearchOption struct {
	FireteamPublicSearchOption int `json:"fireteampublicsearchoption"`
}
type FireteamSlotSearch struct {
	FireteamSlotSearch int `json:"fireteamslotsearch"`
}
type FireteamSummary struct {
	FireteamId                  int    `json:"fireteamid"`
	GroupId                     int    `json:"groupid"`
	Platform                    int    `json:"platform"`
	ActivityType                int    `json:"activitytype"`
	IsImmediate                 bool   `json:"isimmediate"`
	ScheduledTime               string `json:"scheduledtime"`
	OwnerMembershipId           int    `json:"ownermembershipid"`
	PlayerSlotCount             int    `json:"playerslotcount"`
	AlternateSlotCount          int    `json:"alternateslotcount"`
	AvailablePlayerSlotCount    int    `json:"availableplayerslotcount"`
	AvailableAlternateSlotCount int    `json:"availablealternateslotcount"`
	Title                       string `json:"title"`
	DateCreated                 string `json:"datecreated"`
	DateModified                string `json:"datemodified"`
	IsPublic                    bool   `json:"ispublic"`
	Locale                      string `json:"locale"`
	IsValid                     bool   `json:"isvalid"`
	DatePlayerModified          string `json:"dateplayermodified"`
	TitleBeforeModeration       string `json:"titlebeforemoderation"`
}
type SearchResultOfFireteamSummary struct {
	Results                      []FireteamSummary `json:"results"`
	TotalResults                 int               `json:"totalresults"`
	HasMore                      bool              `json:"hasmore"`
	Query                        PagedQuery        `json:"query"`
	ReplacementContinuationToken string            `json:"replacementcontinuationtoken"`
	UseTotalResults              bool              `json:"usetotalresults"`
}
type FireteamResponse struct {
	Summary    FireteamSummary  `json:"summary"`
	Members    []FireteamMember `json:"members"`
	Alternates []FireteamMember `json:"alternates"`
}
type FireteamMember struct {
	DestinyUserInfo                 FireteamUserInfoCard `json:"destinyuserinfo"`
	BungieNetUserInfo               UserInfoCard         `json:"bungienetuserinfo"`
	CharacterId                     int                  `json:"characterid"`
	DateJoined                      string               `json:"datejoined"`
	HasMicrophone                   bool                 `json:"hasmicrophone"`
	LastPlatformInviteAttemptDate   string               `json:"lastplatforminviteattemptdate"`
	LastPlatformInviteAttemptResult int                  `json:"lastplatforminviteattemptresult"`
}
type FireteamUserInfoCard struct {
	FireteamDisplayName         string `json:"fireteamdisplayname"`
	FireteamMembershipType      int    `json:"fireteammembershiptype"`
	SupplementalDisplayName     string `json:"supplementaldisplayname"`
	IconPath                    string `json:"iconpath"`
	CrossSaveOverride           int    `json:"crosssaveoverride"`
	ApplicableMembershipTypes   []int  `json:"applicablemembershiptypes"`
	IsPublic                    bool   `json:"ispublic"`
	MembershipType              int    `json:"membershiptype"`
	MembershipId                int    `json:"membershipid"`
	DisplayName                 string `json:"displayname"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type FireteamPlatformInviteResult struct {
	FireteamPlatformInviteResult int `json:"fireteamplatforminviteresult"`
}
type SearchResultOfFireteamResponse struct {
	Results                      []FireteamResponse `json:"results"`
	TotalResults                 int                `json:"totalresults"`
	HasMore                      bool               `json:"hasmore"`
	Query                        PagedQuery         `json:"query"`
	ReplacementContinuationToken string             `json:"replacementcontinuationtoken"`
	UseTotalResults              bool               `json:"usetotalresults"`
}
type BungieFriendListResponse struct {
	Friends []BungieFriend `json:"friends"`
}
type BungieFriend struct {
	LastSeenAsMembershipId         int         `json:"lastseenasmembershipid"`
	LastSeenAsBungieMembershipType int         `json:"lastseenasbungiemembershiptype"`
	BungieGlobalDisplayName        string      `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode    int         `json:"bungieglobaldisplaynamecode"`
	OnlineStatus                   int         `json:"onlinestatus"`
	OnlineTitle                    int         `json:"onlinetitle"`
	Relationship                   int         `json:"relationship"`
	BungieNetUser                  GeneralUser `json:"bungienetuser"`
}
type PresenceStatus struct {
	PresenceStatus int `json:"presencestatus"`
}
type PresenceOnlineStateFlags struct {
	PresenceOnlineStateFlags int `json:"presenceonlinestateflags"`
}
type FriendRelationshipState struct {
	FriendRelationshipState int `json:"friendrelationshipstate"`
}
type BungieFriendRequestListResponse struct {
	IncomingRequests []BungieFriend `json:"incomingrequests"`
	OutgoingRequests []BungieFriend `json:"outgoingrequests"`
}
type PlatformFriendType struct {
	PlatformFriendType int `json:"platformfriendtype"`
}
type PlatformFriendResponse struct {
	ItemsPerPage    int              `json:"itemsperpage"`
	CurrentPage     int              `json:"currentpage"`
	HasMore         bool             `json:"hasmore"`
	PlatformFriends []PlatformFriend `json:"platformfriends"`
}
type PlatformFriend struct {
	PlatformDisplayName         string `json:"platformdisplayname"`
	FriendPlatform              int    `json:"friendplatform"`
	DestinyMembershipId         int    `json:"destinymembershipid"`
	DestinyMembershipType       int    `json:"destinymembershiptype"`
	BungieNetMembershipId       int    `json:"bungienetmembershipid"`
	BungieGlobalDisplayName     string `json:"bungieglobaldisplayname"`
	BungieGlobalDisplayNameCode int    `json:"bungieglobaldisplaynamecode"`
}
type CoreSettingsConfiguration struct {
	Environment                    string               `json:"environment"`
	Systems                        CoreSystem           `json:"systems"`
	IgnoreReasons                  []CoreSetting        `json:"ignorereasons"`
	ForumCategories                []CoreSetting        `json:"forumcategories"`
	GroupAvatars                   []CoreSetting        `json:"groupavatars"`
	DestinyMembershipTypes         []CoreSetting        `json:"destinymembershiptypes"`
	RecruitmentPlatformTags        []CoreSetting        `json:"recruitmentplatformtags"`
	RecruitmentMiscTags            []CoreSetting        `json:"recruitmentmisctags"`
	RecruitmentActivities          []CoreSetting        `json:"recruitmentactivities"`
	UserContentLocales             []CoreSetting        `json:"usercontentlocales"`
	SystemContentLocales           []CoreSetting        `json:"systemcontentlocales"`
	ClanBannerDecals               []CoreSetting        `json:"clanbannerdecals"`
	ClanBannerDecalColors          []CoreSetting        `json:"clanbannerdecalcolors"`
	ClanBannerGonfalons            []CoreSetting        `json:"clanbannergonfalons"`
	ClanBannerGonfalonColors       []CoreSetting        `json:"clanbannergonfaloncolors"`
	ClanBannerGonfalonDetails      []CoreSetting        `json:"clanbannergonfalondetails"`
	ClanBannerGonfalonDetailColors []CoreSetting        `json:"clanbannergonfalondetailcolors"`
	ClanBannerStandards            []CoreSetting        `json:"clanbannerstandards"`
	Destiny2CoreSettings           Destiny2CoreSettings `json:"destiny2coresettings"`
	EmailSettings                  EmailSettings        `json:"emailsettings"`
	FireteamActivities             []CoreSetting        `json:"fireteamactivities"`
}
type CoreSystem struct {
	Enabled    bool   `json:"enabled"`
	Parameters string `json:"parameters"`
}
type CoreSetting struct {
	Identifier    string        `json:"identifier"`
	IsDefault     bool          `json:"isdefault"`
	DisplayName   string        `json:"displayname"`
	Summary       string        `json:"summary"`
	ImagePath     string        `json:"imagepath"`
	ChildSettings []CoreSetting `json:"childsettings"`
}
type Destiny2CoreSettings struct {
	CollectionRootNode                     int    `json:"collectionrootnode"`
	BadgesRootNode                         int    `json:"badgesrootnode"`
	RecordsRootNode                        int    `json:"recordsrootnode"`
	MedalsRootNode                         int    `json:"medalsrootnode"`
	MetricsRootNode                        int    `json:"metricsrootnode"`
	ActiveTriumphsRootNodeHash             int    `json:"activetriumphsrootnodehash"`
	ActiveSealsRootNodeHash                int    `json:"activesealsrootnodehash"`
	LegacyTriumphsRootNodeHash             int    `json:"legacytriumphsrootnodehash"`
	LegacySealsRootNodeHash                int    `json:"legacysealsrootnodehash"`
	MedalsRootNodeHash                     int    `json:"medalsrootnodehash"`
	ExoticCatalystsRootNodeHash            int    `json:"exoticcatalystsrootnodehash"`
	LoreRootNodeHash                       int    `json:"lorerootnodehash"`
	CurrentRankProgressionHashes           []int  `json:"currentrankprogressionhashes"`
	UndiscoveredCollectibleImage           string `json:"undiscoveredcollectibleimage"`
	AmmoTypeHeavyIcon                      string `json:"ammotypeheavyicon"`
	AmmoTypeSpecialIcon                    string `json:"ammotypespecialicon"`
	AmmoTypePrimaryIcon                    string `json:"ammotypeprimaryicon"`
	CurrentSeasonalArtifactHash            int    `json:"currentseasonalartifacthash"`
	CurrentSeasonHash                      int    `json:"currentseasonhash"`
	SeasonalChallengesPresentationNodeHash int    `json:"seasonalchallengespresentationnodehash"`
	FutureSeasonHashes                     []int  `json:"futureseasonhashes"`
	PastSeasonHashes                       []int  `json:"pastseasonhashes"`
}
type EmailSettings struct {
	OptInDefinitions        EmailOptInDefinition        `json:"optindefinitions"`
	SubscriptionDefinitions EmailSubscriptionDefinition `json:"subscriptiondefinitions"`
	Views                   EmailViewDefinition         `json:"views"`
}
type EmailOptInDefinition struct {
	Name                   string                        `json:"name"`
	Value                  int                           `json:"value"`
	SetByDefault           bool                          `json:"setbydefault"`
	DependentSubscriptions []EmailSubscriptionDefinition `json:"dependentsubscriptions"`
}
type OptInFlags struct {
	OptInFlags int `json:"optinflags"`
}
type EmailSubscriptionDefinition struct {
	Name         string                               `json:"name"`
	Localization EMailSettingSubscriptionLocalization `json:"localization"`
	Value        int                                  `json:"value"`
}
type EMailSettingLocalization struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type EMailSettingSubscriptionLocalization struct {
	UnknownUserDescription      string `json:"unknownuserdescription"`
	RegisteredUserDescription   string `json:"registereduserdescription"`
	UnregisteredUserDescription string `json:"unregistereduserdescription"`
	UnknownUserActionText       string `json:"unknownuseractiontext"`
	KnownUserActionText         string `json:"knownuseractiontext"`
	Title                       string `json:"title"`
	Description                 string `json:"description"`
}
type EmailViewDefinition struct {
	Name         string                       `json:"name"`
	ViewSettings []EmailViewDefinitionSetting `json:"viewsettings"`
}
type EmailViewDefinitionSetting struct {
	Name                string                        `json:"name"`
	Localization        EMailSettingLocalization      `json:"localization"`
	SetByDefault        bool                          `json:"setbydefault"`
	OptInAggregateValue int                           `json:"optinaggregatevalue"`
	Subscriptions       []EmailSubscriptionDefinition `json:"subscriptions"`
}
type GlobalAlert struct {
	AlertKey       string     `json:"alertkey"`
	AlertHtml      string     `json:"alerthtml"`
	AlertTimestamp string     `json:"alerttimestamp"`
	AlertLink      string     `json:"alertlink"`
	AlertLevel     int        `json:"alertlevel"`
	AlertType      int        `json:"alerttype"`
	StreamInfo     StreamInfo `json:"streaminfo"`
}
type GlobalAlertLevel struct {
	GlobalAlertLevel int `json:"globalalertlevel"`
}
type GlobalAlertType struct {
	GlobalAlertType int `json:"globalalerttype"`
}
type StreamInfo struct {
	ChannelName string `json:"channelname"`
}
