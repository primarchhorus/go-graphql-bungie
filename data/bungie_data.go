package data

import (
"encoding/json"
)

type ApplicationScopes struct {
    ApplicationScopes json.Number `json:"applicationscopes"`
}
type ApiUsage struct {
    ApiCalls []Series `json:"apicalls"`
    ThrottledRequests []Series `json:"throttledrequests"`
}
type Series struct {
    Datapoints []Datapoint `json:"datapoints"`
    Target string `json:"target"`
}
type Datapoint struct {
    Time string `json:"time"`
    Count float64 `json:"count"`
}
type Application struct {
    ApplicationId json.Number `json:"applicationid"`
    Name string `json:"name"`
    RedirectUrl string `json:"redirecturl"`
    Link string `json:"link"`
    Scope json.Number `json:"scope"`
    Origin string `json:"origin"`
    Status json.Number `json:"status"`
    CreationDate string `json:"creationdate"`
    StatusChanged string `json:"statuschanged"`
    FirstPublished string `json:"firstpublished"`
    Team []ApplicationDeveloper `json:"team"`
    OverrideAuthorizeViewName string `json:"overrideauthorizeviewname"`
}
type ApplicationStatus struct {
    ApplicationStatus json.Number `json:"applicationstatus"`
}
type ApplicationDeveloper struct {
    Role json.Number `json:"role"`
    ApiEulaVersion json.Number `json:"apieulaversion"`
    User UserInfoCard `json:"user"`
}
type DeveloperRole struct {
    DeveloperRole json.Number `json:"developerrole"`
}
type UserMembership struct {
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type BungieMembershipType struct {
    BungieMembershipType json.Number `json:"bungiemembershiptype"`
}
type CrossSaveUserMembership struct {
    CrossSaveOverride json.Number `json:"crosssaveoverride"`
    ApplicableMembershipTypes []int `json:"applicablemembershiptypes"`
    IsPublic bool `json:"ispublic"`
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type UserInfoCard struct {
    SupplementalDisplayName string `json:"supplementaldisplayname"`
    IconPath string `json:"iconpath"`
    CrossSaveOverride json.Number `json:"crosssaveoverride"`
    ApplicableMembershipTypes []int `json:"applicablemembershiptypes"`
    IsPublic bool `json:"ispublic"`
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type GeneralUser struct {
    MembershipId json.Number `json:"membershipid"`
    UniqueName string `json:"uniquename"`
    NormalizedName string `json:"normalizedname"`
    DisplayName string `json:"displayname"`
    ProfilePicture json.Number `json:"profilepicture"`
    ProfileTheme json.Number `json:"profiletheme"`
    UserTitle json.Number `json:"usertitle"`
    SuccessMessageFlags json.Number `json:"successmessageflags"`
    IsDeleted bool `json:"isdeleted"`
    About string `json:"about"`
    FirstAccess string `json:"firstaccess"`
    LastUpdate string `json:"lastupdate"`
    LegacyPortalUID json.Number `json:"legacyportaluid"`
    Context UserToUserContext `json:"context"`
    PsnDisplayName string `json:"psndisplayname"`
    XboxDisplayName string `json:"xboxdisplayname"`
    FbDisplayName string `json:"fbdisplayname"`
    ShowActivity bool `json:"showactivity"`
    Locale string `json:"locale"`
    LocaleInheritDefault bool `json:"localeinheritdefault"`
    LastBanReportId json.Number `json:"lastbanreportid"`
    ShowGroupMessaging bool `json:"showgroupmessaging"`
    ProfilePicturePath string `json:"profilepicturepath"`
    ProfilePictureWidePath string `json:"profilepicturewidepath"`
    ProfileThemeName string `json:"profilethemename"`
    UserTitleDisplay string `json:"usertitledisplay"`
    StatusText string `json:"statustext"`
    StatusDate string `json:"statusdate"`
    ProfileBanExpire string `json:"profilebanexpire"`
    BlizzardDisplayName string `json:"blizzarddisplayname"`
    SteamDisplayName string `json:"steamdisplayname"`
    StadiaDisplayName string `json:"stadiadisplayname"`
    TwitchDisplayName string `json:"twitchdisplayname"`
    CachedBungieGlobalDisplayName string `json:"cachedbungieglobaldisplayname"`
    CachedBungieGlobalDisplayNameCode json.Number `json:"cachedbungieglobaldisplaynamecode"`
}
type UserToUserContext struct {
    IsFollowing bool `json:"isfollowing"`
    IgnoreStatus IgnoreResponse `json:"ignorestatus"`
    GlobalIgnoreEndDate string `json:"globalignoreenddate"`
}
type IgnoreResponse struct {
    IsIgnored bool `json:"isignored"`
    IgnoreFlags json.Number `json:"ignoreflags"`
}
type IgnoreStatus struct {
    IgnoreStatus json.Number `json:"ignorestatus"`
}
type GetCredentialTypesForAccountResponse struct {
    CredentialType json.Number `json:"credentialtype"`
    CredentialDisplayName string `json:"credentialdisplayname"`
    IsPublic bool `json:"ispublic"`
    CredentialAsString string `json:"credentialasstring"`
}
type BungieCredentialType struct {
    BungieCredentialType json.Number `json:"bungiecredentialtype"`
}
type UserTheme struct {
    UserThemeId json.Number `json:"userthemeid"`
    UserThemeName string `json:"userthemename"`
    UserThemeDescription string `json:"userthemedescription"`
}
type UserMembershipData struct {
    DestinyMemberships []GroupUserInfoCard `json:"destinymemberships"`
    PrimaryMembershipId json.Number `json:"primarymembershipid"`
    BungieNetUser GeneralUser `json:"bungienetuser"`
}
type GroupUserInfoCard struct {
    LastSeenDisplayName string `json:"lastseendisplayname"`
    LastSeenDisplayNameType json.Number `json:"lastseendisplaynametype"`
    SupplementalDisplayName string `json:"supplementaldisplayname"`
    IconPath string `json:"iconpath"`
    CrossSaveOverride json.Number `json:"crosssaveoverride"`
    ApplicableMembershipTypes []int `json:"applicablemembershiptypes"`
    IsPublic bool `json:"ispublic"`
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type HardLinkedUserMembership struct {
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    CrossSaveOverriddenType json.Number `json:"crosssaveoverriddentype"`
    CrossSaveOverriddenMembershipId json.Number `json:"crosssaveoverriddenmembershipid"`
}
type UserSearchResponse struct {
    SearchResults []UserSearchResponseDetail `json:"searchresults"`
    Page json.Number `json:"page"`
    HasMore bool `json:"hasmore"`
}
type UserSearchResponseDetail struct {
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
    BungieNetMembershipId json.Number `json:"bungienetmembershipid"`
    DestinyMemberships []UserInfoCard `json:"destinymemberships"`
}
type ContentTypeDescription struct {
    CType string `json:"ctype"`
    Name string `json:"name"`
    ContentDescription string `json:"contentdescription"`
    PreviewImage string `json:"previewimage"`
    Priority json.Number `json:"priority"`
    Reminder string `json:"reminder"`
    Properties []ContentTypeProperty `json:"properties"`
    TagMetadata []TagMetadataDefinition `json:"tagmetadata"`
    TagMetadataItems TagMetadataItem `json:"tagmetadataitems"`
    UsageExamples []string `json:"usageexamples"`
    ShowInContentEditor bool `json:"showincontenteditor"`
    TypeOf string `json:"typeof"`
    BindIdentifierToProperty string `json:"bindidentifiertoproperty"`
    BoundRegex string `json:"boundregex"`
    ForceIdentifierBinding bool `json:"forceidentifierbinding"`
    AllowComments bool `json:"allowcomments"`
    AutoEnglishPropertyFallback bool `json:"autoenglishpropertyfallback"`
    BulkUploadable bool `json:"bulkuploadable"`
    Previews []ContentPreview `json:"previews"`
    SuppressCmsPath bool `json:"suppresscmspath"`
    PropertySections []ContentTypePropertySection `json:"propertysections"`
}
type ContentTypeProperty struct {
    Name string `json:"name"`
    RootPropertyName string `json:"rootpropertyname"`
    ReadableName string `json:"readablename"`
    Value string `json:"value"`
    PropertyDescription string `json:"propertydescription"`
    Localizable bool `json:"localizable"`
    Fallback bool `json:"fallback"`
    Enabled bool `json:"enabled"`
    Order json.Number `json:"order"`
    Visible bool `json:"visible"`
    IsTitle bool `json:"istitle"`
    Required bool `json:"required"`
    MaxLength json.Number `json:"maxlength"`
    MaxByteLength json.Number `json:"maxbytelength"`
    MaxFileSize json.Number `json:"maxfilesize"`
    Regexp string `json:"regexp"`
    ValidateAs string `json:"validateas"`
    RssAttribute string `json:"rssattribute"`
    VisibleDependency string `json:"visibledependency"`
    VisibleOn string `json:"visibleon"`
    Datatype json.Number `json:"datatype"`
    Attributes string `json:"attributes"`
    ChildProperties []ContentTypeProperty `json:"childproperties"`
    ContentTypeAllowed string `json:"contenttypeallowed"`
    BindToProperty string `json:"bindtoproperty"`
    BoundRegex string `json:"boundregex"`
    RepresentationSelection string `json:"representationselection"`
    DefaultValues []ContentTypeDefaultValue `json:"defaultvalues"`
    IsExternalAllowed bool `json:"isexternalallowed"`
    PropertySection string `json:"propertysection"`
    Weight json.Number `json:"weight"`
    Entitytype string `json:"entitytype"`
    IsCombo bool `json:"iscombo"`
    SuppressProperty bool `json:"suppressproperty"`
    LegalContentTypes []string `json:"legalcontenttypes"`
    RepresentationValidationString string `json:"representationvalidationstring"`
    MinWidth json.Number `json:"minwidth"`
    MaxWidth json.Number `json:"maxwidth"`
    MinHeight json.Number `json:"minheight"`
    MaxHeight json.Number `json:"maxheight"`
    IsVideo bool `json:"isvideo"`
    IsImage bool `json:"isimage"`
}
type ContentPropertyDataTypeEnum struct {
    ContentPropertyDataTypeEnum json.Number `json:"contentpropertydatatypeenum"`
}
type ContentTypeDefaultValue struct {
    WhenClause string `json:"whenclause"`
    WhenValue string `json:"whenvalue"`
    DefaultValue string `json:"defaultvalue"`
}
type TagMetadataDefinition struct {
    Description string `json:"description"`
    Order json.Number `json:"order"`
    Items []TagMetadataItem `json:"items"`
    Datatype string `json:"datatype"`
    Name string `json:"name"`
    IsRequired bool `json:"isrequired"`
}
type TagMetadataItem struct {
    Description string `json:"description"`
    TagText string `json:"tagtext"`
    Groups []string `json:"groups"`
    IsDefault bool `json:"isdefault"`
    Name string `json:"name"`
}
type ContentPreview struct {
    Name string `json:"name"`
    Path string `json:"path"`
    ItemInSet bool `json:"iteminset"`
    SetTag string `json:"settag"`
    SetNesting json.Number `json:"setnesting"`
    UseSetId json.Number `json:"usesetid"`
}
type ContentTypePropertySection struct {
    Name string `json:"name"`
    ReadableName string `json:"readablename"`
    Collapsed bool `json:"collapsed"`
}
type ContentItemPublicContract struct {
    ContentId json.Number `json:"contentid"`
    CType string `json:"ctype"`
    CmsPath string `json:"cmspath"`
    CreationDate string `json:"creationdate"`
    ModifyDate string `json:"modifydate"`
    AllowComments bool `json:"allowcomments"`
    HasAgeGate bool `json:"hasagegate"`
    MinimumAge json.Number `json:"minimumage"`
    RatingImagePath string `json:"ratingimagepath"`
    Author GeneralUser `json:"author"`
    AutoEnglishPropertyFallback bool `json:"autoenglishpropertyfallback"`
    Properties interface{} `json:"properties"`
    Representations []ContentRepresentation `json:"representations"`
    Tags []string `json:"tags"`
    CommentSummary CommentSummary `json:"commentsummary"`
}
type ContentRepresentation struct {
    Name string `json:"name"`
    Path string `json:"path"`
    ValidationString string `json:"validationstring"`
}
type CommentSummary struct {
    TopicId json.Number `json:"topicid"`
    CommentCount json.Number `json:"commentcount"`
}
type SearchResult struct {
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type PagedQuery struct {
    ItemsPerPage json.Number `json:"itemsperpage"`
    CurrentPage json.Number `json:"currentpage"`
    RequestContinuationToken string `json:"requestcontinuationtoken"`
}
type SearchResultOfContentItemPublicContract struct {
    Results []ContentItemPublicContract `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type ForumTopicsCategoryFiltersEnum struct {
    ForumTopicsCategoryFiltersEnum json.Number `json:"forumtopicscategoryfiltersenum"`
}
type ForumTopicsQuickDateEnum struct {
    ForumTopicsQuickDateEnum json.Number `json:"forumtopicsquickdateenum"`
}
type ForumTopicsSortEnum struct {
    ForumTopicsSortEnum json.Number `json:"forumtopicssortenum"`
}
type PostResponse struct {
    LastReplyTimestamp string `json:"lastreplytimestamp"`
    IsPinned bool `json:"ispinned"`
    UrlMediaType json.Number `json:"urlmediatype"`
    Thumbnail string `json:"thumbnail"`
    Popularity json.Number `json:"popularity"`
    IsActive bool `json:"isactive"`
    IsAnnouncement bool `json:"isannouncement"`
    UserRating json.Number `json:"userrating"`
    UserHasRated bool `json:"userhasrated"`
    UserHasMutedPost bool `json:"userhasmutedpost"`
    LatestReplyPostId json.Number `json:"latestreplypostid"`
    LatestReplyAuthorId json.Number `json:"latestreplyauthorid"`
    IgnoreStatus IgnoreResponse `json:"ignorestatus"`
    Locale string `json:"locale"`
}
type ForumMediaType struct {
    ForumMediaType json.Number `json:"forummediatype"`
}
type ForumPostPopularity struct {
    ForumPostPopularity json.Number `json:"forumpostpopularity"`
}
type ForumPostCategoryEnums struct {
    ForumPostCategoryEnums json.Number `json:"forumpostcategoryenums"`
}
type ForumFlagsEnum struct {
    ForumFlagsEnum json.Number `json:"forumflagsenum"`
}
type SearchResultOfPostResponse struct {
    Results []PostResponse `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type PostSearchResponse struct {
    RelatedPosts []PostResponse `json:"relatedposts"`
    Authors []GeneralUser `json:"authors"`
    Groups []GroupResponse `json:"groups"`
    SearchedTags []TagResponse `json:"searchedtags"`
    Polls []PollResponse `json:"polls"`
    RecruitmentDetails []ForumRecruitmentDetail `json:"recruitmentdetails"`
    AvailablePages json.Number `json:"availablepages"`
    Results []PostResponse `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupResponse struct {
    Detail GroupV2 `json:"detail"`
    Founder GroupMember `json:"founder"`
    AlliedIds []int `json:"alliedids"`
    ParentGroup GroupV2 `json:"parentgroup"`
    AllianceStatus json.Number `json:"alliancestatus"`
    GroupJoinInviteCount json.Number `json:"groupjoininvitecount"`
    CurrentUserMembershipsInactiveForDestiny bool `json:"currentusermembershipsinactivefordestiny"`
    CurrentUserMemberMap GroupMember `json:"currentusermembermap"`
    CurrentUserPotentialMemberMap GroupPotentialMember `json:"currentuserpotentialmembermap"`
}
type GroupV2 struct {
    GroupId json.Number `json:"groupid"`
    Name string `json:"name"`
    GroupType json.Number `json:"grouptype"`
    MembershipIdCreated json.Number `json:"membershipidcreated"`
    CreationDate string `json:"creationdate"`
    ModificationDate string `json:"modificationdate"`
    About string `json:"about"`
    Tags []string `json:"tags"`
    MemberCount json.Number `json:"membercount"`
    IsPublic bool `json:"ispublic"`
    IsPublicTopicAdminOnly bool `json:"ispublictopicadminonly"`
    Motto string `json:"motto"`
    AllowChat bool `json:"allowchat"`
    IsDefaultPostPublic bool `json:"isdefaultpostpublic"`
    ChatSecurity json.Number `json:"chatsecurity"`
    Locale string `json:"locale"`
    AvatarImageIndex json.Number `json:"avatarimageindex"`
    Homepage json.Number `json:"homepage"`
    MembershipOption json.Number `json:"membershipoption"`
    DefaultPublicity json.Number `json:"defaultpublicity"`
    Theme string `json:"theme"`
    BannerPath string `json:"bannerpath"`
    AvatarPath string `json:"avatarpath"`
    ConversationId json.Number `json:"conversationid"`
    EnableInvitationMessagingForAdmins bool `json:"enableinvitationmessagingforadmins"`
    BanExpireDate string `json:"banexpiredate"`
    Features GroupFeatures `json:"features"`
    ClanInfo GroupV2ClanInfoAndInvestment `json:"claninfo"`
}
type GroupType struct {
    GroupType json.Number `json:"grouptype"`
}
type ChatSecuritySetting struct {
    ChatSecuritySetting json.Number `json:"chatsecuritysetting"`
}
type GroupHomepage struct {
    GroupHomepage json.Number `json:"grouphomepage"`
}
type MembershipOption struct {
    MembershipOption json.Number `json:"membershipoption"`
}
type GroupPostPublicity struct {
    GroupPostPublicity json.Number `json:"grouppostpublicity"`
}
type GroupFeatures struct {
    MaximumMembers json.Number `json:"maximummembers"`
    MaximumMembershipsOfGroupType json.Number `json:"maximummembershipsofgrouptype"`
    Capabilities json.Number `json:"capabilities"`
    MembershipTypes []int `json:"membershiptypes"`
    InvitePermissionOverride bool `json:"invitepermissionoverride"`
    UpdateCulturePermissionOverride bool `json:"updateculturepermissionoverride"`
    HostGuidedGamePermissionOverride json.Number `json:"hostguidedgamepermissionoverride"`
    UpdateBannerPermissionOverride bool `json:"updatebannerpermissionoverride"`
    JoinLevel json.Number `json:"joinlevel"`
}
type Capabilities struct {
    Capabilities json.Number `json:"capabilities"`
}
type BungieMembershipTypeList struct {
    BungieMembershipTypeList json.Number `json:"bungiemembershiptypelist"`
}
type HostGuidedGamesPermissionLevel struct {
    HostGuidedGamesPermissionLevel json.Number `json:"hostguidedgamespermissionlevel"`
}
type RuntimeGroupMemberType struct {
    RuntimeGroupMemberType json.Number `json:"runtimegroupmembertype"`
}
type GroupV2ClanInfo struct {
    ClanCallsign string `json:"clancallsign"`
    ClanBannerData ClanBanner `json:"clanbannerdata"`
}
type ClanBanner struct {
    DecalId json.Number `json:"decalid"`
    DecalColorId json.Number `json:"decalcolorid"`
    DecalBackgroundColorId json.Number `json:"decalbackgroundcolorid"`
    GonfalonId json.Number `json:"gonfalonid"`
    GonfalonColorId json.Number `json:"gonfaloncolorid"`
    GonfalonDetailId json.Number `json:"gonfalondetailid"`
    GonfalonDetailColorId json.Number `json:"gonfalondetailcolorid"`
}
type GroupV2ClanInfoAndInvestment struct {
    D2ClanProgressions DestinyProgression `json:"d2clanprogressions"`
    ClanCallsign string `json:"clancallsign"`
    ClanBannerData ClanBanner `json:"clanbannerdata"`
}
type DestinyProgression struct {
    ProgressionHash json.Number `json:"progressionhash"`
    DailyProgress json.Number `json:"dailyprogress"`
    DailyLimit json.Number `json:"dailylimit"`
    WeeklyProgress json.Number `json:"weeklyprogress"`
    WeeklyLimit json.Number `json:"weeklylimit"`
    CurrentProgress json.Number `json:"currentprogress"`
    Level json.Number `json:"level"`
    LevelCap json.Number `json:"levelcap"`
    StepIndex json.Number `json:"stepindex"`
    ProgressToNextLevel json.Number `json:"progresstonextlevel"`
    NextLevelAt json.Number `json:"nextlevelat"`
    CurrentResetCount json.Number `json:"currentresetcount"`
    SeasonResets []DestinyProgressionResetEntry `json:"seasonresets"`
    RewardItemStates []int `json:"rewarditemstates"`
}
type DestinyProgressionResetEntry struct {
    Season json.Number `json:"season"`
    Resets json.Number `json:"resets"`
}
type DestinyProgressionRewardItemState struct {
    DestinyProgressionRewardItemState json.Number `json:"destinyprogressionrewarditemstate"`
}
type DestinyDefinition struct {
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyProgressionDefinition struct {
    DisplayProperties DestinyProgressionDisplayPropertiesDefinition `json:"displayproperties"`
    Scope json.Number `json:"scope"`
    RepeatLastStep bool `json:"repeatlaststep"`
    Source string `json:"source"`
    Steps []DestinyProgressionStepDefinition `json:"steps"`
    Visible bool `json:"visible"`
    FactionHash json.Number `json:"factionhash"`
    Color DestinyColor `json:"color"`
    RankIcon string `json:"rankicon"`
    RewardItems []DestinyProgressionRewardItemQuantity `json:"rewarditems"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyDisplayPropertiesDefinition struct {
    Description string `json:"description"`
    Name string `json:"name"`
    Icon string `json:"icon"`
    IconSequences []DestinyIconSequenceDefinition `json:"iconsequences"`
    HighResIcon string `json:"highresicon"`
    HasIcon bool `json:"hasicon"`
}
type DestinyIconSequenceDefinition struct {
    Frames []string `json:"frames"`
}
type DestinyProgressionDisplayPropertiesDefinition struct {
    DisplayUnitsName string `json:"displayunitsname"`
    Description string `json:"description"`
    Name string `json:"name"`
    Icon string `json:"icon"`
    IconSequences []DestinyIconSequenceDefinition `json:"iconsequences"`
    HighResIcon string `json:"highresicon"`
    HasIcon bool `json:"hasicon"`
}
type DestinyProgressionScope struct {
    DestinyProgressionScope json.Number `json:"destinyprogressionscope"`
}
type DestinyProgressionStepDefinition struct {
    StepName string `json:"stepname"`
    DisplayEffectType json.Number `json:"displayeffecttype"`
    ProgressTotal json.Number `json:"progresstotal"`
    RewardItems []DestinyItemQuantity `json:"rewarditems"`
    Icon string `json:"icon"`
}
type DestinyProgressionStepDisplayEffect struct {
    DestinyProgressionStepDisplayEffect json.Number `json:"destinyprogressionstepdisplayeffect"`
}
type DestinyItemQuantity struct {
    ItemHash json.Number `json:"itemhash"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Quantity json.Number `json:"quantity"`
    HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyInventoryItemDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TooltipNotifications []DestinyItemTooltipNotification `json:"tooltipnotifications"`
    CollectibleHash json.Number `json:"collectiblehash"`
    IconWatermark string `json:"iconwatermark"`
    IconWatermarkShelved string `json:"iconwatermarkshelved"`
    SecondaryIcon string `json:"secondaryicon"`
    SecondaryOverlay string `json:"secondaryoverlay"`
    SecondarySpecial string `json:"secondaryspecial"`
    BackgroundColor DestinyColor `json:"backgroundcolor"`
    Screenshot string `json:"screenshot"`
    ItemTypeDisplayName string `json:"itemtypedisplayname"`
    FlavorText string `json:"flavortext"`
    UiItemDisplayStyle string `json:"uiitemdisplaystyle"`
    ItemTypeAndTierDisplayName string `json:"itemtypeandtierdisplayname"`
    DisplaySource string `json:"displaysource"`
    TooltipStyle string `json:"tooltipstyle"`
    Action DestinyItemActionBlockDefinition `json:"action"`
    Inventory DestinyItemInventoryBlockDefinition `json:"inventory"`
    SetData DestinyItemSetBlockDefinition `json:"setdata"`
    Stats DestinyItemStatBlockDefinition `json:"stats"`
    EmblemObjectiveHash json.Number `json:"emblemobjectivehash"`
    EquippingBlock DestinyEquippingBlockDefinition `json:"equippingblock"`
    TranslationBlock DestinyItemTranslationBlockDefinition `json:"translationblock"`
    Preview DestinyItemPreviewBlockDefinition `json:"preview"`
    Quality DestinyItemQualityBlockDefinition `json:"quality"`
    Value DestinyItemValueBlockDefinition `json:"value"`
    SourceData DestinyItemSourceBlockDefinition `json:"sourcedata"`
    Objectives DestinyItemObjectiveBlockDefinition `json:"objectives"`
    Metrics DestinyItemMetricBlockDefinition `json:"metrics"`
    Plug DestinyItemPlugDefinition `json:"plug"`
    Gearset DestinyItemGearsetBlockDefinition `json:"gearset"`
    Sack DestinyItemSackBlockDefinition `json:"sack"`
    Sockets DestinyItemSocketBlockDefinition `json:"sockets"`
    Summary DestinyItemSummaryBlockDefinition `json:"summary"`
    TalentGrid DestinyItemTalentGridBlockDefinition `json:"talentgrid"`
    InvestmentStats []DestinyItemInvestmentStatDefinition `json:"investmentstats"`
    Perks []DestinyItemPerkEntryDefinition `json:"perks"`
    LoreHash json.Number `json:"lorehash"`
    SummaryItemHash json.Number `json:"summaryitemhash"`
    Animations []DestinyAnimationReference `json:"animations"`
    AllowActions bool `json:"allowactions"`
    Links []HyperlinkReference `json:"links"`
    DoesPostmasterPullHaveSideEffects bool `json:"doespostmasterpullhavesideeffects"`
    NonTransferrable bool `json:"nontransferrable"`
    ItemCategoryHashes []int `json:"itemcategoryhashes"`
    SpecialItemType json.Number `json:"specialitemtype"`
    ItemType json.Number `json:"itemtype"`
    ItemSubType json.Number `json:"itemsubtype"`
    ClassType json.Number `json:"classtype"`
    BreakerType json.Number `json:"breakertype"`
    BreakerTypeHash json.Number `json:"breakertypehash"`
    Equippable bool `json:"equippable"`
    DamageTypeHashes []int `json:"damagetypehashes"`
    DamageTypes []int `json:"damagetypes"`
    DefaultDamageType json.Number `json:"defaultdamagetype"`
    DefaultDamageTypeHash json.Number `json:"defaultdamagetypehash"`
    SeasonHash json.Number `json:"seasonhash"`
    IsWrapper bool `json:"iswrapper"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyItemTooltipNotification struct {
    DisplayString string `json:"displaystring"`
    DisplayStyle string `json:"displaystyle"`
}
type DestinyColor struct {
    Red string `json:"red"`
    Green string `json:"green"`
    Blue string `json:"blue"`
    Alpha string `json:"alpha"`
}
type DestinyItemActionBlockDefinition struct {
    VerbName string `json:"verbname"`
    VerbDescription string `json:"verbdescription"`
    IsPositive bool `json:"ispositive"`
    OverlayScreenName string `json:"overlayscreenname"`
    OverlayIcon string `json:"overlayicon"`
    RequiredCooldownSeconds json.Number `json:"requiredcooldownseconds"`
    RequiredItems []DestinyItemActionRequiredItemDefinition `json:"requireditems"`
    ProgressionRewards []DestinyProgressionRewardDefinition `json:"progressionrewards"`
    ActionTypeLabel string `json:"actiontypelabel"`
    RequiredLocation string `json:"requiredlocation"`
    RequiredCooldownHash json.Number `json:"requiredcooldownhash"`
    DeleteOnAction bool `json:"deleteonaction"`
    ConsumeEntireStack bool `json:"consumeentirestack"`
    UseOnAcquire bool `json:"useonacquire"`
}
type DestinyItemActionRequiredItemDefinition struct {
    Count json.Number `json:"count"`
    ItemHash json.Number `json:"itemhash"`
    DeleteOnAction bool `json:"deleteonaction"`
}
type DestinyProgressionRewardDefinition struct {
    ProgressionMappingHash json.Number `json:"progressionmappinghash"`
    Amount json.Number `json:"amount"`
    ApplyThrottles bool `json:"applythrottles"`
}
type DestinyProgressionMappingDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    DisplayUnits string `json:"displayunits"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyItemInventoryBlockDefinition struct {
    StackUniqueLabel string `json:"stackuniquelabel"`
    MaxStackSize json.Number `json:"maxstacksize"`
    BucketTypeHash json.Number `json:"buckettypehash"`
    RecoveryBucketTypeHash json.Number `json:"recoverybuckettypehash"`
    TierTypeHash json.Number `json:"tiertypehash"`
    IsInstanceItem bool `json:"isinstanceitem"`
    TierTypeName string `json:"tiertypename"`
    TierType json.Number `json:"tiertype"`
    ExpirationTooltip string `json:"expirationtooltip"`
    ExpiredInActivityMessage string `json:"expiredinactivitymessage"`
    ExpiredInOrbitMessage string `json:"expiredinorbitmessage"`
    SuppressExpirationWhenObjectivesComplete bool `json:"suppressexpirationwhenobjectivescomplete"`
}
type TierType struct {
    TierType json.Number `json:"tiertype"`
}
type DestinyInventoryBucketDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Scope json.Number `json:"scope"`
    Category json.Number `json:"category"`
    BucketOrder json.Number `json:"bucketorder"`
    ItemCount json.Number `json:"itemcount"`
    Location json.Number `json:"location"`
    HasTransferDestination bool `json:"hastransferdestination"`
    Enabled bool `json:"enabled"`
    Fifo bool `json:"fifo"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type BucketScope struct {
    BucketScope json.Number `json:"bucketscope"`
}
type BucketCategory struct {
    BucketCategory json.Number `json:"bucketcategory"`
}
type ItemLocation struct {
    ItemLocation json.Number `json:"itemlocation"`
}
type DestinyItemTierTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    InfusionProcess DestinyItemTierTypeInfusionBlock `json:"infusionprocess"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyItemTierTypeInfusionBlock struct {
    BaseQualityTransferRatio float64 `json:"basequalitytransferratio"`
    MinimumQualityIncrement json.Number `json:"minimumqualityincrement"`
}
type DestinyItemSetBlockDefinition struct {
    ItemList []DestinyItemSetBlockEntryDefinition `json:"itemlist"`
    RequireOrderedSetItemAdd bool `json:"requireorderedsetitemadd"`
    SetIsFeatured bool `json:"setisfeatured"`
    SetType string `json:"settype"`
    QuestLineName string `json:"questlinename"`
    QuestLineDescription string `json:"questlinedescription"`
    QuestStepSummary string `json:"queststepsummary"`
}
type DestinyItemSetBlockEntryDefinition struct {
    TrackingValue json.Number `json:"trackingvalue"`
    ItemHash json.Number `json:"itemhash"`
}
type DestinyItemStatBlockDefinition struct {
    DisablePrimaryStatDisplay bool `json:"disableprimarystatdisplay"`
    StatGroupHash json.Number `json:"statgrouphash"`
    Stats DestinyInventoryItemStatDefinition `json:"stats"`
    HasDisplayableStats bool `json:"hasdisplayablestats"`
    PrimaryBaseStatHash json.Number `json:"primarybasestathash"`
}
type DestinyInventoryItemStatDefinition struct {
    StatHash json.Number `json:"stathash"`
    Value json.Number `json:"value"`
    Minimum json.Number `json:"minimum"`
    Maximum json.Number `json:"maximum"`
    DisplayMaximum json.Number `json:"displaymaximum"`
}
type DestinyStatDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    AggregationType json.Number `json:"aggregationtype"`
    HasComputedBlock bool `json:"hascomputedblock"`
    StatCategory json.Number `json:"statcategory"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyStatAggregationType struct {
    DestinyStatAggregationType json.Number `json:"destinystataggregationtype"`
}
type DestinyStatCategory struct {
    DestinyStatCategory json.Number `json:"destinystatcategory"`
}
type DestinyStatGroupDefinition struct {
    MaximumValue json.Number `json:"maximumvalue"`
    UiPosition json.Number `json:"uiposition"`
    ScaledStats []DestinyStatDisplayDefinition `json:"scaledstats"`
    Overrides DestinyStatOverrideDefinition `json:"overrides"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyStatDisplayDefinition struct {
    StatHash json.Number `json:"stathash"`
    MaximumValue json.Number `json:"maximumvalue"`
    DisplayAsNumeric bool `json:"displayasnumeric"`
    DisplayInterpolation []InterpolationPoint `json:"displayinterpolation"`
}
type InterpolationPoint struct {
    Value json.Number `json:"value"`
    Weight json.Number `json:"weight"`
}
type DestinyStatOverrideDefinition struct {
    StatHash json.Number `json:"stathash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyEquippingBlockDefinition struct {
    GearsetItemHash json.Number `json:"gearsetitemhash"`
    UniqueLabel string `json:"uniquelabel"`
    UniqueLabelHash json.Number `json:"uniquelabelhash"`
    EquipmentSlotTypeHash json.Number `json:"equipmentslottypehash"`
    Attributes json.Number `json:"attributes"`
    AmmoType json.Number `json:"ammotype"`
    DisplayStrings []string `json:"displaystrings"`
}
type EquippingItemBlockAttributes struct {
    EquippingItemBlockAttributes json.Number `json:"equippingitemblockattributes"`
}
type DestinyAmmunitionType struct {
    DestinyAmmunitionType json.Number `json:"destinyammunitiontype"`
}
type DestinyEquipmentSlotDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    EquipmentCategoryHash json.Number `json:"equipmentcategoryhash"`
    BucketTypeHash json.Number `json:"buckettypehash"`
    ApplyCustomArtDyes bool `json:"applycustomartdyes"`
    ArtDyeChannels []DestinyArtDyeReference `json:"artdyechannels"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyArtDyeReference struct {
    ArtDyeChannelHash json.Number `json:"artdyechannelhash"`
}
type DestinyItemTranslationBlockDefinition struct {
    WeaponPatternIdentifier string `json:"weaponpatternidentifier"`
    WeaponPatternHash json.Number `json:"weaponpatternhash"`
    DefaultDyes []DyeReference `json:"defaultdyes"`
    LockedDyes []DyeReference `json:"lockeddyes"`
    CustomDyes []DyeReference `json:"customdyes"`
    Arrangements []DestinyGearArtArrangementReference `json:"arrangements"`
    HasGeometry bool `json:"hasgeometry"`
}
type DyeReference struct {
    ChannelHash json.Number `json:"channelhash"`
    DyeHash json.Number `json:"dyehash"`
}
type DestinyGearArtArrangementReference struct {
    ClassHash json.Number `json:"classhash"`
    ArtArrangementHash json.Number `json:"artarrangementhash"`
}
type DestinyItemPreviewBlockDefinition struct {
    ScreenStyle string `json:"screenstyle"`
    PreviewVendorHash json.Number `json:"previewvendorhash"`
    ArtifactHash json.Number `json:"artifacthash"`
    PreviewActionString string `json:"previewactionstring"`
    DerivedItemCategories []DestinyDerivedItemCategoryDefinition `json:"deriveditemcategories"`
}
type DestinyDerivedItemCategoryDefinition struct {
    CategoryDescription string `json:"categorydescription"`
    Items []DestinyDerivedItemDefinition `json:"items"`
}
type DestinyDerivedItemDefinition struct {
    ItemHash json.Number `json:"itemhash"`
    ItemName string `json:"itemname"`
    ItemDetail string `json:"itemdetail"`
    ItemDescription string `json:"itemdescription"`
    IconPath string `json:"iconpath"`
    VendorItemIndex json.Number `json:"vendoritemindex"`
}
type DestinyVendorDefinition struct {
    DisplayProperties DestinyVendorDisplayPropertiesDefinition `json:"displayproperties"`
    VendorProgressionType json.Number `json:"vendorprogressiontype"`
    BuyString string `json:"buystring"`
    SellString string `json:"sellstring"`
    DisplayItemHash json.Number `json:"displayitemhash"`
    InhibitBuying bool `json:"inhibitbuying"`
    InhibitSelling bool `json:"inhibitselling"`
    FactionHash json.Number `json:"factionhash"`
    ResetIntervalMinutes json.Number `json:"resetintervalminutes"`
    ResetOffsetMinutes json.Number `json:"resetoffsetminutes"`
    FailureStrings []string `json:"failurestrings"`
    UnlockRanges []DateRange `json:"unlockranges"`
    VendorIdentifier string `json:"vendoridentifier"`
    VendorPortrait string `json:"vendorportrait"`
    VendorBanner string `json:"vendorbanner"`
    Enabled bool `json:"enabled"`
    Visible bool `json:"visible"`
    VendorSubcategoryIdentifier string `json:"vendorsubcategoryidentifier"`
    ConsolidateCategories bool `json:"consolidatecategories"`
    Actions []DestinyVendorActionDefinition `json:"actions"`
    Categories []DestinyVendorCategoryEntryDefinition `json:"categories"`
    OriginalCategories []DestinyVendorCategoryEntryDefinition `json:"originalcategories"`
    DisplayCategories []DestinyDisplayCategoryDefinition `json:"displaycategories"`
    Interactions []DestinyVendorInteractionDefinition `json:"interactions"`
    InventoryFlyouts []DestinyVendorInventoryFlyoutDefinition `json:"inventoryflyouts"`
    ItemList []DestinyVendorItemDefinition `json:"itemlist"`
    Services []DestinyVendorServiceDefinition `json:"services"`
    AcceptedItems []DestinyVendorAcceptedItemDefinition `json:"accepteditems"`
    ReturnWithVendorRequest bool `json:"returnwithvendorrequest"`
    Locations []DestinyVendorLocationDefinition `json:"locations"`
    Groups []DestinyVendorGroupReference `json:"groups"`
    IgnoreSaleItemHashes []int `json:"ignoresaleitemhashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyVendorDisplayPropertiesDefinition struct {
    LargeIcon string `json:"largeicon"`
    Subtitle string `json:"subtitle"`
    OriginalIcon string `json:"originalicon"`
    RequirementsDisplay []DestinyVendorRequirementDisplayEntryDefinition `json:"requirementsdisplay"`
    SmallTransparentIcon string `json:"smalltransparenticon"`
    MapIcon string `json:"mapicon"`
    LargeTransparentIcon string `json:"largetransparenticon"`
    Description string `json:"description"`
    Name string `json:"name"`
    Icon string `json:"icon"`
    IconSequences []DestinyIconSequenceDefinition `json:"iconsequences"`
    HighResIcon string `json:"highresicon"`
    HasIcon bool `json:"hasicon"`
}
type DestinyVendorRequirementDisplayEntryDefinition struct {
    Icon string `json:"icon"`
    Name string `json:"name"`
    Source string `json:"source"`
    Type string `json:"type"`
}
type DestinyVendorProgressionType struct {
    DestinyVendorProgressionType json.Number `json:"destinyvendorprogressiontype"`
}
type DateRange struct {
    Start string `json:"start"`
    End string `json:"end"`
}
type DestinyVendorActionDefinition struct {
    Description string `json:"description"`
    ExecuteSeconds json.Number `json:"executeseconds"`
    Icon string `json:"icon"`
    Name string `json:"name"`
    Verb string `json:"verb"`
    IsPositive bool `json:"ispositive"`
    ActionId string `json:"actionid"`
    ActionHash json.Number `json:"actionhash"`
    AutoPerformAction bool `json:"autoperformaction"`
}
type DestinyVendorCategoryEntryDefinition struct {
    CategoryIndex json.Number `json:"categoryindex"`
    SortValue json.Number `json:"sortvalue"`
    CategoryHash json.Number `json:"categoryhash"`
    QuantityAvailable json.Number `json:"quantityavailable"`
    ShowUnavailableItems bool `json:"showunavailableitems"`
    HideIfNoCurrency bool `json:"hideifnocurrency"`
    HideFromRegularPurchase bool `json:"hidefromregularpurchase"`
    BuyStringOverride string `json:"buystringoverride"`
    DisabledDescription string `json:"disableddescription"`
    DisplayTitle string `json:"displaytitle"`
    Overlay DestinyVendorCategoryOverlayDefinition `json:"overlay"`
    VendorItemIndexes []int `json:"vendoritemindexes"`
    IsPreview bool `json:"ispreview"`
    IsDisplayOnly bool `json:"isdisplayonly"`
    ResetIntervalMinutesOverride json.Number `json:"resetintervalminutesoverride"`
    ResetOffsetMinutesOverride json.Number `json:"resetoffsetminutesoverride"`
}
type DestinyVendorCategoryOverlayDefinition struct {
    ChoiceDescription string `json:"choicedescription"`
    Description string `json:"description"`
    Icon string `json:"icon"`
    Title string `json:"title"`
    CurrencyItemHash json.Number `json:"currencyitemhash"`
}
type DestinyDisplayCategoryDefinition struct {
    Index json.Number `json:"index"`
    Identifier string `json:"identifier"`
    DisplayCategoryHash json.Number `json:"displaycategoryhash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    DisplayInBanner bool `json:"displayinbanner"`
    ProgressionHash json.Number `json:"progressionhash"`
    SortOrder json.Number `json:"sortorder"`
    DisplayStyleHash json.Number `json:"displaystylehash"`
    DisplayStyleIdentifier string `json:"displaystyleidentifier"`
}
type VendorDisplayCategorySortOrder struct {
    VendorDisplayCategorySortOrder json.Number `json:"vendordisplaycategorysortorder"`
}
type DestinyVendorInteractionDefinition struct {
    InteractionIndex json.Number `json:"interactionindex"`
    Replies []DestinyVendorInteractionReplyDefinition `json:"replies"`
    VendorCategoryIndex json.Number `json:"vendorcategoryindex"`
    QuestlineItemHash json.Number `json:"questlineitemhash"`
    SackInteractionList []DestinyVendorInteractionSackEntryDefinition `json:"sackinteractionlist"`
    UiInteractionType json.Number `json:"uiinteractiontype"`
    InteractionType json.Number `json:"interactiontype"`
    RewardBlockLabel string `json:"rewardblocklabel"`
    RewardVendorCategoryIndex json.Number `json:"rewardvendorcategoryindex"`
    FlavorLineOne string `json:"flavorlineone"`
    FlavorLineTwo string `json:"flavorlinetwo"`
    HeaderDisplayProperties DestinyDisplayPropertiesDefinition `json:"headerdisplayproperties"`
    Instructions string `json:"instructions"`
}
type DestinyVendorInteractionReplyDefinition struct {
    ItemRewardsSelection json.Number `json:"itemrewardsselection"`
    Reply string `json:"reply"`
    ReplyType json.Number `json:"replytype"`
}
type DestinyVendorInteractionRewardSelection struct {
    DestinyVendorInteractionRewardSelection json.Number `json:"destinyvendorinteractionrewardselection"`
}
type DestinyVendorReplyType struct {
    DestinyVendorReplyType json.Number `json:"destinyvendorreplytype"`
}
type DestinyVendorInteractionSackEntryDefinition struct {
    SackType json.Number `json:"sacktype"`
}
type VendorInteractionType struct {
    VendorInteractionType json.Number `json:"vendorinteractiontype"`
}
type DestinyVendorInventoryFlyoutDefinition struct {
    LockedDescription string `json:"lockeddescription"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Buckets []DestinyVendorInventoryFlyoutBucketDefinition `json:"buckets"`
    FlyoutId json.Number `json:"flyoutid"`
    SuppressNewness bool `json:"suppressnewness"`
    EquipmentSlotHash json.Number `json:"equipmentslothash"`
}
type DestinyVendorInventoryFlyoutBucketDefinition struct {
    Collapsible bool `json:"collapsible"`
    InventoryBucketHash json.Number `json:"inventorybuckethash"`
    SortItemsBy json.Number `json:"sortitemsby"`
}
type DestinyItemSortType struct {
    DestinyItemSortType json.Number `json:"destinyitemsorttype"`
}
type DestinyVendorItemDefinition struct {
    VendorItemIndex json.Number `json:"vendoritemindex"`
    ItemHash json.Number `json:"itemhash"`
    Quantity json.Number `json:"quantity"`
    FailureIndexes []int `json:"failureindexes"`
    Currencies []DestinyVendorItemQuantity `json:"currencies"`
    RefundPolicy json.Number `json:"refundpolicy"`
    RefundTimeLimit json.Number `json:"refundtimelimit"`
    CreationLevels []DestinyItemCreationEntryLevelDefinition `json:"creationlevels"`
    DisplayCategoryIndex json.Number `json:"displaycategoryindex"`
    CategoryIndex json.Number `json:"categoryindex"`
    OriginalCategoryIndex json.Number `json:"originalcategoryindex"`
    MinimumLevel json.Number `json:"minimumlevel"`
    MaximumLevel json.Number `json:"maximumlevel"`
    Action DestinyVendorSaleItemActionBlockDefinition `json:"action"`
    DisplayCategory string `json:"displaycategory"`
    InventoryBucketHash json.Number `json:"inventorybuckethash"`
    VisibilityScope json.Number `json:"visibilityscope"`
    PurchasableScope json.Number `json:"purchasablescope"`
    Exclusivity json.Number `json:"exclusivity"`
    IsOffer bool `json:"isoffer"`
    IsCrm bool `json:"iscrm"`
    SortValue json.Number `json:"sortvalue"`
    ExpirationTooltip string `json:"expirationtooltip"`
    RedirectToSaleIndexes []int `json:"redirecttosaleindexes"`
    SocketOverrides []DestinyVendorItemSocketOverride `json:"socketoverrides"`
    Unpurchasable bool `json:"unpurchasable"`
}
type DestinyVendorItemQuantity struct {
    ItemHash json.Number `json:"itemhash"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Quantity json.Number `json:"quantity"`
    HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyVendorItemRefundPolicy struct {
    DestinyVendorItemRefundPolicy json.Number `json:"destinyvendoritemrefundpolicy"`
}
type DestinyItemCreationEntryLevelDefinition struct {
    Level json.Number `json:"level"`
}
type DestinyVendorSaleItemActionBlockDefinition struct {
    ExecuteSeconds float64 `json:"executeseconds"`
    IsPositive bool `json:"ispositive"`
}
type DestinyGatingScope struct {
    DestinyGatingScope json.Number `json:"destinygatingscope"`
}
type DestinyVendorItemSocketOverride struct {
    SingleItemHash json.Number `json:"singleitemhash"`
    RandomizedOptionsCount json.Number `json:"randomizedoptionscount"`
    SocketTypeHash json.Number `json:"sockettypehash"`
}
type DestinySocketTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    InsertAction DestinyInsertPlugActionDefinition `json:"insertaction"`
    PlugWhitelist []DestinyPlugWhitelistEntryDefinition `json:"plugwhitelist"`
    SocketCategoryHash json.Number `json:"socketcategoryhash"`
    Visibility json.Number `json:"visibility"`
    AlwaysRandomizeSockets bool `json:"alwaysrandomizesockets"`
    IsPreviewEnabled bool `json:"ispreviewenabled"`
    HideDuplicateReusablePlugs bool `json:"hideduplicatereusableplugs"`
    OverridesUiAppearance bool `json:"overridesuiappearance"`
    AvoidDuplicatesOnInitialization bool `json:"avoidduplicatesoninitialization"`
    CurrencyScalars []DestinySocketTypeScalarMaterialRequirementEntry `json:"currencyscalars"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyInsertPlugActionDefinition struct {
    ActionExecuteSeconds json.Number `json:"actionexecuteseconds"`
    ActionType json.Number `json:"actiontype"`
}
type SocketTypeActionType struct {
    SocketTypeActionType json.Number `json:"sockettypeactiontype"`
}
type DestinyPlugWhitelistEntryDefinition struct {
    CategoryHash json.Number `json:"categoryhash"`
    CategoryIdentifier string `json:"categoryidentifier"`
    ReinitializationPossiblePlugHashes []int `json:"reinitializationpossibleplughashes"`
}
type DestinySocketVisibility struct {
    DestinySocketVisibility json.Number `json:"destinysocketvisibility"`
}
type DestinySocketTypeScalarMaterialRequirementEntry struct {
    CurrencyItemHash json.Number `json:"currencyitemhash"`
    ScalarValue json.Number `json:"scalarvalue"`
}
type DestinySocketCategoryDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    UiCategoryStyle json.Number `json:"uicategorystyle"`
    CategoryStyle json.Number `json:"categorystyle"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinySocketCategoryStyle struct {
    DestinySocketCategoryStyle json.Number `json:"destinysocketcategorystyle"`
}
type DestinyVendorServiceDefinition struct {
    Name string `json:"name"`
}
type DestinyVendorAcceptedItemDefinition struct {
    AcceptedInventoryBucketHash json.Number `json:"acceptedinventorybuckethash"`
    DestinationInventoryBucketHash json.Number `json:"destinationinventorybuckethash"`
}
type DestinyVendorLocationDefinition struct {
    DestinationHash json.Number `json:"destinationhash"`
    BackgroundImagePath string `json:"backgroundimagepath"`
}
type DestinyDestinationDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    PlaceHash json.Number `json:"placehash"`
    DefaultFreeroamActivityHash json.Number `json:"defaultfreeroamactivityhash"`
    ActivityGraphEntries []DestinyActivityGraphListEntryDefinition `json:"activitygraphentries"`
    BubbleSettings []DestinyDestinationBubbleSettingDefinition `json:"bubblesettings"`
    Bubbles []DestinyBubbleDefinition `json:"bubbles"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityGraphListEntryDefinition struct {
    ActivityGraphHash json.Number `json:"activitygraphhash"`
}
type DestinyActivityGraphDefinition struct {
    Nodes []DestinyActivityGraphNodeDefinition `json:"nodes"`
    ArtElements []DestinyActivityGraphArtElementDefinition `json:"artelements"`
    Connections []DestinyActivityGraphConnectionDefinition `json:"connections"`
    DisplayObjectives []DestinyActivityGraphDisplayObjectiveDefinition `json:"displayobjectives"`
    DisplayProgressions []DestinyActivityGraphDisplayProgressionDefinition `json:"displayprogressions"`
    LinkedGraphs []DestinyLinkedGraphDefinition `json:"linkedgraphs"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityGraphNodeDefinition struct {
    NodeId json.Number `json:"nodeid"`
    OverrideDisplay DestinyDisplayPropertiesDefinition `json:"overridedisplay"`
    Position DestinyPositionDefinition `json:"position"`
    FeaturingStates []DestinyActivityGraphNodeFeaturingStateDefinition `json:"featuringstates"`
    Activities []DestinyActivityGraphNodeActivityDefinition `json:"activities"`
    States []DestinyActivityGraphNodeStateEntry `json:"states"`
}
type DestinyPositionDefinition struct {
    X json.Number `json:"x"`
    Y json.Number `json:"y"`
    Z json.Number `json:"z"`
}
type DestinyActivityGraphNodeFeaturingStateDefinition struct {
    HighlightType json.Number `json:"highlighttype"`
}
type ActivityGraphNodeHighlightType struct {
    ActivityGraphNodeHighlightType json.Number `json:"activitygraphnodehighlighttype"`
}
type DestinyActivityGraphNodeActivityDefinition struct {
    NodeActivityId json.Number `json:"nodeactivityid"`
    ActivityHash json.Number `json:"activityhash"`
}
type DestinyActivityDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    OriginalDisplayProperties DestinyDisplayPropertiesDefinition `json:"originaldisplayproperties"`
    SelectionScreenDisplayProperties DestinyDisplayPropertiesDefinition `json:"selectionscreendisplayproperties"`
    ReleaseIcon string `json:"releaseicon"`
    ReleaseTime json.Number `json:"releasetime"`
    ActivityLightLevel json.Number `json:"activitylightlevel"`
    DestinationHash json.Number `json:"destinationhash"`
    PlaceHash json.Number `json:"placehash"`
    ActivityTypeHash json.Number `json:"activitytypehash"`
    Tier json.Number `json:"tier"`
    PgcrImage string `json:"pgcrimage"`
    Rewards []DestinyActivityRewardDefinition `json:"rewards"`
    Modifiers []DestinyActivityModifierReferenceDefinition `json:"modifiers"`
    IsPlaylist bool `json:"isplaylist"`
    Challenges []DestinyActivityChallengeDefinition `json:"challenges"`
    OptionalUnlockStrings []DestinyActivityUnlockStringDefinition `json:"optionalunlockstrings"`
    PlaylistItems []DestinyActivityPlaylistItemDefinition `json:"playlistitems"`
    ActivityGraphList []DestinyActivityGraphListEntryDefinition `json:"activitygraphlist"`
    Matchmaking DestinyActivityMatchmakingBlockDefinition `json:"matchmaking"`
    GuidedGame DestinyActivityGuidedBlockDefinition `json:"guidedgame"`
    DirectActivityModeHash json.Number `json:"directactivitymodehash"`
    DirectActivityModeType json.Number `json:"directactivitymodetype"`
    Loadouts []DestinyActivityLoadoutRequirementSet `json:"loadouts"`
    ActivityModeHashes []int `json:"activitymodehashes"`
    ActivityModeTypes []int `json:"activitymodetypes"`
    IsPvP bool `json:"ispvp"`
    InsertionPoints []DestinyActivityInsertionPointDefinition `json:"insertionpoints"`
    ActivityLocationMappings []DestinyEnvironmentLocationMapping `json:"activitylocationmappings"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityRewardDefinition struct {
    RewardText string `json:"rewardtext"`
    RewardItems []DestinyItemQuantity `json:"rewarditems"`
}
type DestinyActivityModifierReferenceDefinition struct {
    ActivityModifierHash json.Number `json:"activitymodifierhash"`
}
type DestinyActivityModifierDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityChallengeDefinition struct {
    ObjectiveHash json.Number `json:"objectivehash"`
    DummyRewards []DestinyItemQuantity `json:"dummyrewards"`
}
type DestinyObjectiveDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    CompletionValue json.Number `json:"completionvalue"`
    Scope json.Number `json:"scope"`
    LocationHash json.Number `json:"locationhash"`
    AllowNegativeValue bool `json:"allownegativevalue"`
    AllowValueChangeWhenCompleted bool `json:"allowvaluechangewhencompleted"`
    IsCountingDownward bool `json:"iscountingdownward"`
    ValueStyle json.Number `json:"valuestyle"`
    ProgressDescription string `json:"progressdescription"`
    Perks DestinyObjectivePerkEntryDefinition `json:"perks"`
    Stats DestinyObjectiveStatEntryDefinition `json:"stats"`
    MinimumVisibilityThreshold json.Number `json:"minimumvisibilitythreshold"`
    AllowOvercompletion bool `json:"allowovercompletion"`
    ShowValueOnComplete bool `json:"showvalueoncomplete"`
    CompletedValueStyle json.Number `json:"completedvaluestyle"`
    InProgressValueStyle json.Number `json:"inprogressvaluestyle"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyUnlockValueUIStyle struct {
    DestinyUnlockValueUIStyle json.Number `json:"destinyunlockvalueuistyle"`
}
type DestinyObjectivePerkEntryDefinition struct {
    PerkHash json.Number `json:"perkhash"`
    Style json.Number `json:"style"`
}
type DestinyObjectiveGrantStyle struct {
    DestinyObjectiveGrantStyle json.Number `json:"destinyobjectivegrantstyle"`
}
type DestinySandboxPerkDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    PerkIdentifier string `json:"perkidentifier"`
    IsDisplayable bool `json:"isdisplayable"`
    DamageType json.Number `json:"damagetype"`
    DamageTypeHash json.Number `json:"damagetypehash"`
    PerkGroups DestinyTalentNodeStepGroups `json:"perkgroups"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DamageType struct {
    DamageType json.Number `json:"damagetype"`
}
type DestinyTalentNodeStepGroups struct {
    WeaponPerformance json.Number `json:"weaponperformance"`
    ImpactEffects json.Number `json:"impacteffects"`
    GuardianAttributes json.Number `json:"guardianattributes"`
    LightAbilities json.Number `json:"lightabilities"`
    DamageTypes json.Number `json:"damagetypes"`
}
type DestinyTalentNodeStepWeaponPerformances struct {
    DestinyTalentNodeStepWeaponPerformances json.Number `json:"destinytalentnodestepweaponperformances"`
}
type DestinyTalentNodeStepImpactEffects struct {
    DestinyTalentNodeStepImpactEffects json.Number `json:"destinytalentnodestepimpacteffects"`
}
type DestinyTalentNodeStepGuardianAttributes struct {
    DestinyTalentNodeStepGuardianAttributes json.Number `json:"destinytalentnodestepguardianattributes"`
}
type DestinyTalentNodeStepLightAbilities struct {
    DestinyTalentNodeStepLightAbilities json.Number `json:"destinytalentnodesteplightabilities"`
}
type DestinyTalentNodeStepDamageTypes struct {
    DestinyTalentNodeStepDamageTypes json.Number `json:"destinytalentnodestepdamagetypes"`
}
type DestinyObjectiveStatEntryDefinition struct {
    Stat DestinyItemInvestmentStatDefinition `json:"stat"`
    Style json.Number `json:"style"`
}
type DestinyItemInvestmentStatDefinition struct {
    StatTypeHash json.Number `json:"stattypehash"`
    Value json.Number `json:"value"`
    IsConditionallyActive bool `json:"isconditionallyactive"`
}
type DestinyLocationDefinition struct {
    VendorHash json.Number `json:"vendorhash"`
    LocationReleases []DestinyLocationReleaseDefinition `json:"locationreleases"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyLocationReleaseDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    SmallTransparentIcon string `json:"smalltransparenticon"`
    MapIcon string `json:"mapicon"`
    LargeTransparentIcon string `json:"largetransparenticon"`
    SpawnPoint json.Number `json:"spawnpoint"`
    DestinationHash json.Number `json:"destinationhash"`
    ActivityHash json.Number `json:"activityhash"`
    ActivityGraphHash json.Number `json:"activitygraphhash"`
    ActivityGraphNodeHash json.Number `json:"activitygraphnodehash"`
    ActivityBubbleName json.Number `json:"activitybubblename"`
    ActivityPathBundle json.Number `json:"activitypathbundle"`
    ActivityPathDestination json.Number `json:"activitypathdestination"`
    NavPointType json.Number `json:"navpointtype"`
    WorldPosition []int `json:"worldposition"`
}
type DestinyActivityNavPointType struct {
    DestinyActivityNavPointType json.Number `json:"destinyactivitynavpointtype"`
}
type DestinyActivityUnlockStringDefinition struct {
    DisplayString string `json:"displaystring"`
}
type DestinyActivityPlaylistItemDefinition struct {
    ActivityHash json.Number `json:"activityhash"`
    DirectActivityModeHash json.Number `json:"directactivitymodehash"`
    DirectActivityModeType json.Number `json:"directactivitymodetype"`
    ActivityModeHashes []int `json:"activitymodehashes"`
    ActivityModeTypes []int `json:"activitymodetypes"`
}
type DestinyActivityModeType struct {
    DestinyActivityModeType json.Number `json:"destinyactivitymodetype"`
}
type DestinyActivityModeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    PgcrImage string `json:"pgcrimage"`
    ModeType json.Number `json:"modetype"`
    ActivityModeCategory json.Number `json:"activitymodecategory"`
    IsTeamBased bool `json:"isteambased"`
    IsAggregateMode bool `json:"isaggregatemode"`
    ParentHashes []int `json:"parenthashes"`
    FriendlyName string `json:"friendlyname"`
    ActivityModeMappings int `json:"activitymodemappings"`
    Display bool `json:"display"`
    Order json.Number `json:"order"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityModeCategory struct {
    DestinyActivityModeCategory json.Number `json:"destinyactivitymodecategory"`
}
type DestinyActivityMatchmakingBlockDefinition struct {
    IsMatchmade bool `json:"ismatchmade"`
    MinParty json.Number `json:"minparty"`
    MaxParty json.Number `json:"maxparty"`
    MaxPlayers json.Number `json:"maxplayers"`
    RequiresGuardianOath bool `json:"requiresguardianoath"`
}
type DestinyActivityGuidedBlockDefinition struct {
    GuidedMaxLobbySize json.Number `json:"guidedmaxlobbysize"`
    GuidedMinLobbySize json.Number `json:"guidedminlobbysize"`
    GuidedDisbandCount json.Number `json:"guideddisbandcount"`
}
type DestinyActivityLoadoutRequirementSet struct {
    Requirements []DestinyActivityLoadoutRequirement `json:"requirements"`
}
type DestinyActivityLoadoutRequirement struct {
    EquipmentSlotHash json.Number `json:"equipmentslothash"`
    AllowedEquippedItemHashes []int `json:"allowedequippeditemhashes"`
    AllowedWeaponSubTypes []int `json:"allowedweaponsubtypes"`
}
type DestinyItemSubType struct {
    DestinyItemSubType json.Number `json:"destinyitemsubtype"`
}
type DestinyActivityInsertionPointDefinition struct {
    PhaseHash json.Number `json:"phasehash"`
}
type DestinyEnvironmentLocationMapping struct {
    LocationHash json.Number `json:"locationhash"`
    ActivationSource string `json:"activationsource"`
    ItemHash json.Number `json:"itemhash"`
    ObjectiveHash json.Number `json:"objectivehash"`
    ActivityHash json.Number `json:"activityhash"`
}
type DestinyPlaceDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyActivityGraphNodeStateEntry struct {
    State json.Number `json:"state"`
}
type DestinyGraphNodeState struct {
    DestinyGraphNodeState json.Number `json:"destinygraphnodestate"`
}
type DestinyActivityGraphArtElementDefinition struct {
    Position DestinyPositionDefinition `json:"position"`
}
type DestinyActivityGraphConnectionDefinition struct {
    SourceNodeHash json.Number `json:"sourcenodehash"`
    DestNodeHash json.Number `json:"destnodehash"`
}
type DestinyActivityGraphDisplayObjectiveDefinition struct {
    Id json.Number `json:"id"`
    ObjectiveHash json.Number `json:"objectivehash"`
}
type DestinyActivityGraphDisplayProgressionDefinition struct {
    Id json.Number `json:"id"`
    ProgressionHash json.Number `json:"progressionhash"`
}
type DestinyLinkedGraphDefinition struct {
    Description string `json:"description"`
    Name string `json:"name"`
    UnlockExpression DestinyUnlockExpressionDefinition `json:"unlockexpression"`
    LinkedGraphId json.Number `json:"linkedgraphid"`
    LinkedGraphs []DestinyLinkedGraphEntryDefinition `json:"linkedgraphs"`
    Overview string `json:"overview"`
}
type DestinyUnlockExpressionDefinition struct {
    Scope json.Number `json:"scope"`
}
type DestinyLinkedGraphEntryDefinition struct {
    ActivityGraphHash json.Number `json:"activitygraphhash"`
}
type DestinyDestinationBubbleSettingDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyBubbleDefinition struct {
    Hash json.Number `json:"hash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyVendorGroupReference struct {
    VendorGroupHash json.Number `json:"vendorgrouphash"`
}
type DestinyVendorGroupDefinition struct {
    Order json.Number `json:"order"`
    CategoryName string `json:"categoryname"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyFactionDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    ProgressionHash json.Number `json:"progressionhash"`
    TokenValues int `json:"tokenvalues"`
    RewardItemHash json.Number `json:"rewarditemhash"`
    RewardVendorHash json.Number `json:"rewardvendorhash"`
    Vendors []DestinyFactionVendorDefinition `json:"vendors"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyFactionVendorDefinition struct {
    VendorHash json.Number `json:"vendorhash"`
    DestinationHash json.Number `json:"destinationhash"`
    BackgroundImagePath string `json:"backgroundimagepath"`
}
type DestinyArtifactDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TranslationBlock DestinyItemTranslationBlockDefinition `json:"translationblock"`
    Tiers []DestinyArtifactTierDefinition `json:"tiers"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyArtifactTierDefinition struct {
    TierHash json.Number `json:"tierhash"`
    DisplayTitle string `json:"displaytitle"`
    ProgressRequirementMessage string `json:"progressrequirementmessage"`
    Items []DestinyArtifactTierItemDefinition `json:"items"`
    MinimumUnlockPointsUsedRequirement json.Number `json:"minimumunlockpointsusedrequirement"`
}
type DestinyArtifactTierItemDefinition struct {
    ItemHash json.Number `json:"itemhash"`
}
type DestinyItemQualityBlockDefinition struct {
    ItemLevels []int `json:"itemlevels"`
    QualityLevel json.Number `json:"qualitylevel"`
    InfusionCategoryName string `json:"infusioncategoryname"`
    InfusionCategoryHash json.Number `json:"infusioncategoryhash"`
    InfusionCategoryHashes []int `json:"infusioncategoryhashes"`
    ProgressionLevelRequirementHash json.Number `json:"progressionlevelrequirementhash"`
    CurrentVersion json.Number `json:"currentversion"`
    Versions []DestinyItemVersionDefinition `json:"versions"`
    DisplayVersionWatermarkIcons []string `json:"displayversionwatermarkicons"`
}
type DestinyItemVersionDefinition struct {
    PowerCapHash json.Number `json:"powercaphash"`
}
type DestinyPowerCapDefinition struct {
    PowerCap json.Number `json:"powercap"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyProgressionLevelRequirementDefinition struct {
    RequirementCurve []InterpolationPointFloat `json:"requirementcurve"`
    ProgressionHash json.Number `json:"progressionhash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type InterpolationPointFloat struct {
    Value float64 `json:"value"`
    Weight float64 `json:"weight"`
}
type DestinyItemValueBlockDefinition struct {
    ItemValue []DestinyItemQuantity `json:"itemvalue"`
    ValueDescription string `json:"valuedescription"`
}
type DestinyItemSourceBlockDefinition struct {
    SourceHashes []int `json:"sourcehashes"`
    Sources []DestinyItemSourceDefinition `json:"sources"`
    Exclusive json.Number `json:"exclusive"`
    VendorSources []DestinyItemVendorSourceReference `json:"vendorsources"`
}
type DestinyItemSourceDefinition struct {
    Level json.Number `json:"level"`
    MinQuality json.Number `json:"minquality"`
    MaxQuality json.Number `json:"maxquality"`
    MinLevelRequired json.Number `json:"minlevelrequired"`
    MaxLevelRequired json.Number `json:"maxlevelrequired"`
    ComputedStats DestinyInventoryItemStatDefinition `json:"computedstats"`
    SourceHashes []int `json:"sourcehashes"`
}
type DestinyRewardSourceDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Category json.Number `json:"category"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyRewardSourceCategory struct {
    DestinyRewardSourceCategory json.Number `json:"destinyrewardsourcecategory"`
}
type DestinyItemVendorSourceReference struct {
    VendorHash json.Number `json:"vendorhash"`
    VendorItemIndexes []int `json:"vendoritemindexes"`
}
type DestinyItemObjectiveBlockDefinition struct {
    ObjectiveHashes []int `json:"objectivehashes"`
    DisplayActivityHashes []int `json:"displayactivityhashes"`
    RequireFullObjectiveCompletion bool `json:"requirefullobjectivecompletion"`
    QuestlineItemHash json.Number `json:"questlineitemhash"`
    Narrative string `json:"narrative"`
    ObjectiveVerbName string `json:"objectiveverbname"`
    QuestTypeIdentifier string `json:"questtypeidentifier"`
    QuestTypeHash json.Number `json:"questtypehash"`
    PerObjectiveDisplayProperties []DestinyObjectiveDisplayProperties `json:"perobjectivedisplayproperties"`
    DisplayAsStatTracker bool `json:"displayasstattracker"`
}
type DestinyObjectiveDisplayProperties struct {
    ActivityHash json.Number `json:"activityhash"`
    DisplayOnItemPreviewScreen bool `json:"displayonitempreviewscreen"`
}
type DestinyItemMetricBlockDefinition struct {
    AvailableMetricCategoryNodeHashes []int `json:"availablemetriccategorynodehashes"`
}
type DestinyPresentationNodeBaseDefinition struct {
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyPresentationNodeType struct {
    DestinyPresentationNodeType json.Number `json:"destinypresentationnodetype"`
}
type DestinyTraitDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TraitCategoryId string `json:"traitcategoryid"`
    TraitCategoryHash json.Number `json:"traitcategoryhash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyTraitCategoryDefinition struct {
    TraitCategoryId string `json:"traitcategoryid"`
    TraitHashes []int `json:"traithashes"`
    TraitIds []string `json:"traitids"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyScoredPresentationNodeBaseDefinition struct {
    MaxCategoryRecordScore json.Number `json:"maxcategoryrecordscore"`
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyPresentationNodeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    OriginalIcon string `json:"originalicon"`
    RootViewIcon string `json:"rootviewicon"`
    NodeType json.Number `json:"nodetype"`
    Scope json.Number `json:"scope"`
    ObjectiveHash json.Number `json:"objectivehash"`
    CompletionRecordHash json.Number `json:"completionrecordhash"`
    Children DestinyPresentationNodeChildrenBlock `json:"children"`
    DisplayStyle json.Number `json:"displaystyle"`
    ScreenStyle json.Number `json:"screenstyle"`
    Requirements DestinyPresentationNodeRequirementsBlock `json:"requirements"`
    DisableChildSubscreenNavigation bool `json:"disablechildsubscreennavigation"`
    MaxCategoryRecordScore json.Number `json:"maxcategoryrecordscore"`
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyScope struct {
    DestinyScope json.Number `json:"destinyscope"`
}
type DestinyPresentationNodeChildrenBlock struct {
    PresentationNodes []DestinyPresentationNodeChildEntry `json:"presentationnodes"`
    Collectibles []DestinyPresentationNodeCollectibleChildEntry `json:"collectibles"`
    Records []DestinyPresentationNodeRecordChildEntry `json:"records"`
    Metrics []DestinyPresentationNodeMetricChildEntry `json:"metrics"`
}
type DestinyPresentationNodeChildEntry struct {
    PresentationNodeHash json.Number `json:"presentationnodehash"`
}
type DestinyPresentationNodeCollectibleChildEntry struct {
    CollectibleHash json.Number `json:"collectiblehash"`
}
type DestinyCollectibleDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Scope json.Number `json:"scope"`
    SourceString string `json:"sourcestring"`
    SourceHash json.Number `json:"sourcehash"`
    ItemHash json.Number `json:"itemhash"`
    AcquisitionInfo DestinyCollectibleAcquisitionBlock `json:"acquisitioninfo"`
    StateInfo DestinyCollectibleStateBlock `json:"stateinfo"`
    PresentationInfo DestinyPresentationChildBlock `json:"presentationinfo"`
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyCollectibleAcquisitionBlock struct {
    AcquireMaterialRequirementHash json.Number `json:"acquirematerialrequirementhash"`
    AcquireTimestampUnlockValueHash json.Number `json:"acquiretimestampunlockvaluehash"`
}
type DestinyMaterialRequirementSetDefinition struct {
    Materials []DestinyMaterialRequirement `json:"materials"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyMaterialRequirement struct {
    ItemHash json.Number `json:"itemhash"`
    DeleteOnAction bool `json:"deleteonaction"`
    Count json.Number `json:"count"`
    OmitFromRequirements bool `json:"omitfromrequirements"`
}
type DestinyUnlockValueDefinition struct {
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyCollectibleStateBlock struct {
    ObscuredOverrideItemHash json.Number `json:"obscuredoverrideitemhash"`
    Requirements DestinyPresentationNodeRequirementsBlock `json:"requirements"`
}
type DestinyPresentationNodeRequirementsBlock struct {
    EntitlementUnavailableMessage string `json:"entitlementunavailablemessage"`
}
type DestinyPresentationChildBlock struct {
    PresentationNodeType json.Number `json:"presentationnodetype"`
    ParentPresentationNodeHashes []int `json:"parentpresentationnodehashes"`
    DisplayStyle json.Number `json:"displaystyle"`
}
type DestinyPresentationDisplayStyle struct {
    DestinyPresentationDisplayStyle json.Number `json:"destinypresentationdisplaystyle"`
}
type DestinyPresentationNodeRecordChildEntry struct {
    RecordHash json.Number `json:"recordhash"`
}
type DestinyRecordDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Scope json.Number `json:"scope"`
    PresentationInfo DestinyPresentationChildBlock `json:"presentationinfo"`
    LoreHash json.Number `json:"lorehash"`
    ObjectiveHashes []int `json:"objectivehashes"`
    RecordValueStyle json.Number `json:"recordvaluestyle"`
    ForTitleGilding bool `json:"fortitlegilding"`
    TitleInfo DestinyRecordTitleBlock `json:"titleinfo"`
    CompletionInfo DestinyRecordCompletionBlock `json:"completioninfo"`
    StateInfo SchemaRecordStateBlock `json:"stateinfo"`
    Requirements DestinyPresentationNodeRequirementsBlock `json:"requirements"`
    ExpirationInfo DestinyRecordExpirationBlock `json:"expirationinfo"`
    IntervalInfo DestinyRecordIntervalBlock `json:"intervalinfo"`
    RewardItems []DestinyItemQuantity `json:"rewarditems"`
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyRecordValueStyle struct {
    DestinyRecordValueStyle json.Number `json:"destinyrecordvaluestyle"`
}
type DestinyRecordTitleBlock struct {
    HasTitle bool `json:"hastitle"`
    TitlesByGender string `json:"titlesbygender"`
    TitlesByGenderHash string `json:"titlesbygenderhash"`
    GildingTrackingRecordHash json.Number `json:"gildingtrackingrecordhash"`
}
type DestinyGender struct {
    DestinyGender json.Number `json:"destinygender"`
}
type DestinyGenderDefinition struct {
    GenderType json.Number `json:"gendertype"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyRecordCompletionBlock struct {
    PartialCompletionObjectiveCountThreshold json.Number `json:"partialcompletionobjectivecountthreshold"`
    ScoreValue json.Number `json:"scorevalue"`
    ShouldFireToast bool `json:"shouldfiretoast"`
    ToastStyle json.Number `json:"toaststyle"`
}
type DestinyRecordToastStyle struct {
    DestinyRecordToastStyle json.Number `json:"destinyrecordtoaststyle"`
}
type SchemaRecordStateBlock struct {
    FeaturedPriority json.Number `json:"featuredpriority"`
    ObscuredString string `json:"obscuredstring"`
}
type DestinyRecordExpirationBlock struct {
    HasExpiration bool `json:"hasexpiration"`
    Description string `json:"description"`
    Icon string `json:"icon"`
}
type DestinyRecordIntervalBlock struct {
    IntervalObjectives []DestinyRecordIntervalObjective `json:"intervalobjectives"`
    IntervalRewards []DestinyRecordIntervalRewards `json:"intervalrewards"`
    OriginalObjectiveArrayInsertionIndex json.Number `json:"originalobjectivearrayinsertionindex"`
}
type DestinyRecordIntervalObjective struct {
    IntervalObjectiveHash json.Number `json:"intervalobjectivehash"`
    IntervalScoreValue json.Number `json:"intervalscorevalue"`
}
type DestinyRecordIntervalRewards struct {
    IntervalRewardItems []DestinyItemQuantity `json:"intervalrewarditems"`
}
type DestinyLoreDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Subtitle string `json:"subtitle"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyPresentationNodeMetricChildEntry struct {
    MetricHash json.Number `json:"metrichash"`
}
type DestinyMetricDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TrackingObjectiveHash json.Number `json:"trackingobjectivehash"`
    LowerValueIsBetter bool `json:"lowervalueisbetter"`
    PresentationNodeType json.Number `json:"presentationnodetype"`
    TraitIds []string `json:"traitids"`
    TraitHashes []int `json:"traithashes"`
    ParentNodeHashes []int `json:"parentnodehashes"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyPresentationScreenStyle struct {
    DestinyPresentationScreenStyle json.Number `json:"destinypresentationscreenstyle"`
}
type DestinyItemPlugDefinition struct {
    InsertionRules []DestinyPlugRuleDefinition `json:"insertionrules"`
    PlugCategoryIdentifier string `json:"plugcategoryidentifier"`
    PlugCategoryHash json.Number `json:"plugcategoryhash"`
    OnActionRecreateSelf bool `json:"onactionrecreateself"`
    InsertionMaterialRequirementHash json.Number `json:"insertionmaterialrequirementhash"`
    PreviewItemOverrideHash json.Number `json:"previewitemoverridehash"`
    EnabledMaterialRequirementHash json.Number `json:"enabledmaterialrequirementhash"`
    EnabledRules []DestinyPlugRuleDefinition `json:"enabledrules"`
    UiPlugLabel string `json:"uipluglabel"`
    PlugStyle json.Number `json:"plugstyle"`
    PlugAvailability json.Number `json:"plugavailability"`
    AlternateUiPlugLabel string `json:"alternateuipluglabel"`
    AlternatePlugStyle json.Number `json:"alternateplugstyle"`
    IsDummyPlug bool `json:"isdummyplug"`
    ParentItemOverride DestinyParentItemOverride `json:"parentitemoverride"`
    EnergyCapacity DestinyEnergyCapacityEntry `json:"energycapacity"`
    EnergyCost DestinyEnergyCostEntry `json:"energycost"`
}
type DestinyPlugRuleDefinition struct {
    FailureMessage string `json:"failuremessage"`
}
type PlugUiStyles struct {
    PlugUiStyles json.Number `json:"pluguistyles"`
}
type PlugAvailabilityMode struct {
    PlugAvailabilityMode json.Number `json:"plugavailabilitymode"`
}
type DestinyParentItemOverride struct {
    AdditionalEquipRequirementsDisplayStrings []string `json:"additionalequiprequirementsdisplaystrings"`
    PipIcon string `json:"pipicon"`
}
type DestinyEnergyCapacityEntry struct {
    CapacityValue json.Number `json:"capacityvalue"`
    EnergyTypeHash json.Number `json:"energytypehash"`
    EnergyType json.Number `json:"energytype"`
}
type DestinyEnergyType struct {
    DestinyEnergyType json.Number `json:"destinyenergytype"`
}
type DestinyEnergyTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TransparentIconPath string `json:"transparenticonpath"`
    ShowIcon bool `json:"showicon"`
    EnumValue json.Number `json:"enumvalue"`
    CapacityStatHash json.Number `json:"capacitystathash"`
    CostStatHash json.Number `json:"coststathash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyEnergyCostEntry struct {
    EnergyCost json.Number `json:"energycost"`
    EnergyTypeHash json.Number `json:"energytypehash"`
    EnergyType json.Number `json:"energytype"`
}
type DestinyItemGearsetBlockDefinition struct {
    TrackingValueMax json.Number `json:"trackingvaluemax"`
    ItemList []int `json:"itemlist"`
}
type DestinyItemSackBlockDefinition struct {
    DetailAction string `json:"detailaction"`
    OpenAction string `json:"openaction"`
    SelectItemCount json.Number `json:"selectitemcount"`
    VendorSackType string `json:"vendorsacktype"`
    OpenOnAcquire bool `json:"openonacquire"`
}
type DestinyItemSocketBlockDefinition struct {
    Detail string `json:"detail"`
    SocketEntries []DestinyItemSocketEntryDefinition `json:"socketentries"`
    IntrinsicSockets []DestinyItemIntrinsicSocketEntryDefinition `json:"intrinsicsockets"`
    SocketCategories []DestinyItemSocketCategoryDefinition `json:"socketcategories"`
}
type DestinyItemSocketEntryDefinition struct {
    SocketTypeHash json.Number `json:"sockettypehash"`
    SingleInitialItemHash json.Number `json:"singleinitialitemhash"`
    ReusablePlugItems []DestinyItemSocketEntryPlugItemDefinition `json:"reusableplugitems"`
    PreventInitializationOnVendorPurchase bool `json:"preventinitializationonvendorpurchase"`
    HidePerksInItemTooltip bool `json:"hideperksinitemtooltip"`
    PlugSources json.Number `json:"plugsources"`
    ReusablePlugSetHash json.Number `json:"reusableplugsethash"`
    RandomizedPlugSetHash json.Number `json:"randomizedplugsethash"`
    DefaultVisible bool `json:"defaultvisible"`
}
type DestinyItemSocketEntryPlugItemDefinition struct {
    PlugItemHash json.Number `json:"plugitemhash"`
}
type SocketPlugSources struct {
    SocketPlugSources json.Number `json:"socketplugsources"`
}
type DestinyPlugSetDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    ReusablePlugItems []DestinyItemSocketEntryPlugItemRandomizedDefinition `json:"reusableplugitems"`
    IsFakePlugSet bool `json:"isfakeplugset"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyItemSocketEntryPlugItemRandomizedDefinition struct {
    CurrentlyCanRoll bool `json:"currentlycanroll"`
    PlugItemHash json.Number `json:"plugitemhash"`
}
type DestinyItemIntrinsicSocketEntryDefinition struct {
    PlugItemHash json.Number `json:"plugitemhash"`
    SocketTypeHash json.Number `json:"sockettypehash"`
    DefaultVisible bool `json:"defaultvisible"`
}
type DestinyItemSocketCategoryDefinition struct {
    SocketCategoryHash json.Number `json:"socketcategoryhash"`
    SocketIndexes []int `json:"socketindexes"`
}
type DestinyItemSummaryBlockDefinition struct {
    SortPriority json.Number `json:"sortpriority"`
}
type DestinyItemTalentGridBlockDefinition struct {
    TalentGridHash json.Number `json:"talentgridhash"`
    ItemDetailString string `json:"itemdetailstring"`
    BuildName string `json:"buildname"`
    HudDamageType json.Number `json:"huddamagetype"`
    HudIcon string `json:"hudicon"`
}
type DestinyTalentGridDefinition struct {
    MaxGridLevel json.Number `json:"maxgridlevel"`
    GridLevelPerColumn json.Number `json:"gridlevelpercolumn"`
    ProgressionHash json.Number `json:"progressionhash"`
    Nodes []DestinyTalentNodeDefinition `json:"nodes"`
    ExclusiveSets []DestinyTalentNodeExclusiveSetDefinition `json:"exclusivesets"`
    IndependentNodeIndexes []int `json:"independentnodeindexes"`
    Groups DestinyTalentExclusiveGroup `json:"groups"`
    NodeCategories []DestinyTalentNodeCategory `json:"nodecategories"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyTalentNodeDefinition struct {
    NodeIndex json.Number `json:"nodeindex"`
    NodeHash json.Number `json:"nodehash"`
    Row json.Number `json:"row"`
    Column json.Number `json:"column"`
    PrerequisiteNodeIndexes []int `json:"prerequisitenodeindexes"`
    BinaryPairNodeIndex json.Number `json:"binarypairnodeindex"`
    AutoUnlocks bool `json:"autounlocks"`
    LastStepRepeats bool `json:"laststeprepeats"`
    IsRandom bool `json:"israndom"`
    RandomActivationRequirement DestinyNodeActivationRequirement `json:"randomactivationrequirement"`
    IsRandomRepurchasable bool `json:"israndomrepurchasable"`
    Steps []DestinyNodeStepDefinition `json:"steps"`
    ExclusiveWithNodeHashes []int `json:"exclusivewithnodehashes"`
    RandomStartProgressionBarAtProgression json.Number `json:"randomstartprogressionbaratprogression"`
    LayoutIdentifier string `json:"layoutidentifier"`
    GroupHash json.Number `json:"grouphash"`
    LoreHash json.Number `json:"lorehash"`
    NodeStyleIdentifier string `json:"nodestyleidentifier"`
    IgnoreForCompletion bool `json:"ignoreforcompletion"`
}
type DestinyNodeActivationRequirement struct {
    GridLevel json.Number `json:"gridlevel"`
    MaterialRequirementHashes []int `json:"materialrequirementhashes"`
}
type DestinyNodeStepDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    StepIndex json.Number `json:"stepindex"`
    NodeStepHash json.Number `json:"nodestephash"`
    InteractionDescription string `json:"interactiondescription"`
    DamageType json.Number `json:"damagetype"`
    DamageTypeHash json.Number `json:"damagetypehash"`
    ActivationRequirement DestinyNodeActivationRequirement `json:"activationrequirement"`
    CanActivateNextStep bool `json:"canactivatenextstep"`
    NextStepIndex json.Number `json:"nextstepindex"`
    IsNextStepRandom bool `json:"isnextsteprandom"`
    PerkHashes []int `json:"perkhashes"`
    StartProgressionBarAtProgress json.Number `json:"startprogressionbaratprogress"`
    StatHashes []int `json:"stathashes"`
    AffectsQuality bool `json:"affectsquality"`
    StepGroups DestinyTalentNodeStepGroups `json:"stepgroups"`
    AffectsLevel bool `json:"affectslevel"`
    SocketReplacements []DestinyNodeSocketReplaceResponse `json:"socketreplacements"`
}
type DestinyNodeSocketReplaceResponse struct {
    SocketTypeHash json.Number `json:"sockettypehash"`
    PlugItemHash json.Number `json:"plugitemhash"`
}
type DestinyDamageTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    TransparentIconPath string `json:"transparenticonpath"`
    ShowIcon bool `json:"showicon"`
    EnumValue json.Number `json:"enumvalue"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyTalentNodeExclusiveSetDefinition struct {
    NodeIndexes []int `json:"nodeindexes"`
}
type DestinyTalentExclusiveGroup struct {
    GroupHash json.Number `json:"grouphash"`
    LoreHash json.Number `json:"lorehash"`
    NodeHashes []int `json:"nodehashes"`
    OpposingGroupHashes []int `json:"opposinggrouphashes"`
    OpposingNodeHashes []int `json:"opposingnodehashes"`
}
type DestinyTalentNodeCategory struct {
    Identifier string `json:"identifier"`
    IsLoreDriven bool `json:"isloredriven"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    NodeHashes []int `json:"nodehashes"`
}
type DestinyItemPerkEntryDefinition struct {
    RequirementDisplayString string `json:"requirementdisplaystring"`
    PerkHash json.Number `json:"perkhash"`
    PerkVisibility json.Number `json:"perkvisibility"`
}
type ItemPerkVisibility struct {
    ItemPerkVisibility json.Number `json:"itemperkvisibility"`
}
type DestinyAnimationReference struct {
    AnimName string `json:"animname"`
    AnimIdentifier string `json:"animidentifier"`
    Path string `json:"path"`
}
type HyperlinkReference struct {
    Title string `json:"title"`
    Url string `json:"url"`
}
type SpecialItemType struct {
    SpecialItemType json.Number `json:"specialitemtype"`
}
type DestinyItemType struct {
    DestinyItemType json.Number `json:"destinyitemtype"`
}
type DestinyClass struct {
    DestinyClass json.Number `json:"destinyclass"`
}
type DestinyBreakerType struct {
    DestinyBreakerType json.Number `json:"destinybreakertype"`
}
type DestinyItemCategoryDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Visible bool `json:"visible"`
    Deprecated bool `json:"deprecated"`
    ShortTitle string `json:"shorttitle"`
    ItemTypeRegex string `json:"itemtyperegex"`
    GrantDestinyBreakerType json.Number `json:"grantdestinybreakertype"`
    PlugCategoryIdentifier string `json:"plugcategoryidentifier"`
    ItemTypeRegexNot string `json:"itemtyperegexnot"`
    OriginBucketIdentifier string `json:"originbucketidentifier"`
    GrantDestinyItemType json.Number `json:"grantdestinyitemtype"`
    GrantDestinySubType json.Number `json:"grantdestinysubtype"`
    GrantDestinyClass json.Number `json:"grantdestinyclass"`
    TraitId string `json:"traitid"`
    GroupedCategoryHashes []int `json:"groupedcategoryhashes"`
    ParentCategoryHashes []int `json:"parentcategoryhashes"`
    GroupCategoryOnly bool `json:"groupcategoryonly"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyBreakerTypeDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    EnumValue json.Number `json:"enumvalue"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinySeasonDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    BackgroundImagePath string `json:"backgroundimagepath"`
    SeasonNumber json.Number `json:"seasonnumber"`
    StartDate string `json:"startdate"`
    EndDate string `json:"enddate"`
    SeasonPassHash json.Number `json:"seasonpasshash"`
    SeasonPassProgressionHash json.Number `json:"seasonpassprogressionhash"`
    ArtifactItemHash json.Number `json:"artifactitemhash"`
    SealPresentationNodeHash json.Number `json:"sealpresentationnodehash"`
    SeasonalChallengesPresentationNodeHash json.Number `json:"seasonalchallengespresentationnodehash"`
    Preview DestinySeasonPreviewDefinition `json:"preview"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinySeasonPreviewDefinition struct {
    Description string `json:"description"`
    LinkPath string `json:"linkpath"`
    VideoLink string `json:"videolink"`
    Images []DestinySeasonPreviewImageDefinition `json:"images"`
}
type DestinySeasonPreviewImageDefinition struct {
    ThumbnailImage string `json:"thumbnailimage"`
    HighResImage string `json:"highresimage"`
}
type DestinySeasonPassDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    RewardProgressionHash json.Number `json:"rewardprogressionhash"`
    PrestigeProgressionHash json.Number `json:"prestigeprogressionhash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyProgressionRewardItemQuantity struct {
    RewardedAtProgressionLevel json.Number `json:"rewardedatprogressionlevel"`
    AcquisitionBehavior json.Number `json:"acquisitionbehavior"`
    UiDisplayStyle string `json:"uidisplaystyle"`
    ClaimUnlockDisplayStrings []string `json:"claimunlockdisplaystrings"`
    ItemHash json.Number `json:"itemhash"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Quantity json.Number `json:"quantity"`
    HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyProgressionRewardItemAcquisitionBehavior struct {
    DestinyProgressionRewardItemAcquisitionBehavior json.Number `json:"destinyprogressionrewarditemacquisitionbehavior"`
}
type GroupUserBase struct {
    GroupId json.Number `json:"groupid"`
    DestinyUserInfo GroupUserInfoCard `json:"destinyuserinfo"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    JoinDate string `json:"joindate"`
}
type GroupMember struct {
    MemberType json.Number `json:"membertype"`
    IsOnline bool `json:"isonline"`
    LastOnlineStatusChange json.Number `json:"lastonlinestatuschange"`
    GroupId json.Number `json:"groupid"`
    DestinyUserInfo GroupUserInfoCard `json:"destinyuserinfo"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    JoinDate string `json:"joindate"`
}
type GroupAllianceStatus struct {
    GroupAllianceStatus json.Number `json:"groupalliancestatus"`
}
type GroupPotentialMember struct {
    PotentialStatus json.Number `json:"potentialstatus"`
    GroupId json.Number `json:"groupid"`
    DestinyUserInfo GroupUserInfoCard `json:"destinyuserinfo"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    JoinDate string `json:"joindate"`
}
type GroupPotentialMemberStatus struct {
    GroupPotentialMemberStatus json.Number `json:"grouppotentialmemberstatus"`
}
type TagResponse struct {
    TagText string `json:"tagtext"`
    IgnoreStatus IgnoreResponse `json:"ignorestatus"`
}
type PollResponse struct {
    TopicId json.Number `json:"topicid"`
    Results []PollResult `json:"results"`
    TotalVotes json.Number `json:"totalvotes"`
}
type PollResult struct {
    AnswerText string `json:"answertext"`
    AnswerSlot json.Number `json:"answerslot"`
    LastVoteDate string `json:"lastvotedate"`
    Votes json.Number `json:"votes"`
    RequestingUserVoted bool `json:"requestinguservoted"`
}
type ForumRecruitmentDetail struct {
    TopicId json.Number `json:"topicid"`
    MicrophoneRequired bool `json:"microphonerequired"`
    Intensity json.Number `json:"intensity"`
    Tone json.Number `json:"tone"`
    Approved bool `json:"approved"`
    ConversationId json.Number `json:"conversationid"`
    PlayerSlotsTotal json.Number `json:"playerslotstotal"`
    PlayerSlotsRemaining json.Number `json:"playerslotsremaining"`
    Fireteam []GeneralUser `json:"fireteam"`
    KickedPlayerIds []int `json:"kickedplayerids"`
}
type ForumRecruitmentIntensityLabel struct {
    ForumRecruitmentIntensityLabel json.Number `json:"forumrecruitmentintensitylabel"`
}
type ForumRecruitmentToneLabel struct {
    ForumRecruitmentToneLabel json.Number `json:"forumrecruitmenttonelabel"`
}
type ForumPostSortEnum struct {
    ForumPostSortEnum json.Number `json:"forumpostsortenum"`
}
type GroupTheme struct {
    Name string `json:"name"`
    Folder string `json:"folder"`
    Description string `json:"description"`
}
type GroupDateRange struct {
    GroupDateRange json.Number `json:"groupdaterange"`
}
type GroupV2Card struct {
    GroupId json.Number `json:"groupid"`
    Name string `json:"name"`
    GroupType json.Number `json:"grouptype"`
    CreationDate string `json:"creationdate"`
    About string `json:"about"`
    Motto string `json:"motto"`
    MemberCount json.Number `json:"membercount"`
    Locale string `json:"locale"`
    MembershipOption json.Number `json:"membershipoption"`
    Capabilities json.Number `json:"capabilities"`
    ClanInfo GroupV2ClanInfo `json:"claninfo"`
    AvatarPath string `json:"avatarpath"`
    Theme string `json:"theme"`
}
type SearchResultOfGroupV2Card struct {
    Results []GroupV2Card `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupSearchResponse struct {
    Results []GroupV2Card `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupQuery struct {
    Name string `json:"name"`
    GroupType json.Number `json:"grouptype"`
    CreationDate json.Number `json:"creationdate"`
    SortBy json.Number `json:"sortby"`
    GroupMemberCountFilter json.Number `json:"groupmembercountfilter"`
    LocaleFilter string `json:"localefilter"`
    TagText string `json:"tagtext"`
    ItemsPerPage json.Number `json:"itemsperpage"`
    CurrentPage json.Number `json:"currentpage"`
    RequestContinuationToken string `json:"requestcontinuationtoken"`
}
type GroupSortBy struct {
    GroupSortBy json.Number `json:"groupsortby"`
}
type GroupMemberCountFilter struct {
    GroupMemberCountFilter json.Number `json:"groupmembercountfilter"`
}
type GroupNameSearchRequest struct {
    GroupName string `json:"groupname"`
    GroupType json.Number `json:"grouptype"`
}
type GroupOptionalConversation struct {
    GroupId json.Number `json:"groupid"`
    ConversationId json.Number `json:"conversationid"`
    ChatEnabled bool `json:"chatenabled"`
    ChatName string `json:"chatname"`
    ChatSecurity json.Number `json:"chatsecurity"`
}
type GroupEditAction struct {
    Name string `json:"name"`
    About string `json:"about"`
    Motto string `json:"motto"`
    Theme string `json:"theme"`
    AvatarImageIndex json.Number `json:"avatarimageindex"`
    Tags string `json:"tags"`
    IsPublic bool `json:"ispublic"`
    MembershipOption json.Number `json:"membershipoption"`
    IsPublicTopicAdminOnly bool `json:"ispublictopicadminonly"`
    AllowChat bool `json:"allowchat"`
    ChatSecurity json.Number `json:"chatsecurity"`
    Callsign string `json:"callsign"`
    Locale string `json:"locale"`
    Homepage json.Number `json:"homepage"`
    EnableInvitationMessagingForAdmins bool `json:"enableinvitationmessagingforadmins"`
    DefaultPublicity json.Number `json:"defaultpublicity"`
}
type GroupOptionsEditAction struct {
    InvitePermissionOverride bool `json:"invitepermissionoverride"`
    UpdateCulturePermissionOverride bool `json:"updateculturepermissionoverride"`
    HostGuidedGamePermissionOverride json.Number `json:"hostguidedgamepermissionoverride"`
    UpdateBannerPermissionOverride bool `json:"updatebannerpermissionoverride"`
    JoinLevel json.Number `json:"joinlevel"`
}
type GroupOptionalConversationAddRequest struct {
    ChatName string `json:"chatname"`
    ChatSecurity json.Number `json:"chatsecurity"`
}
type GroupOptionalConversationEditRequest struct {
    ChatEnabled bool `json:"chatenabled"`
    ChatName string `json:"chatname"`
    ChatSecurity json.Number `json:"chatsecurity"`
}
type SearchResultOfGroupMember struct {
    Results []GroupMember `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupMemberLeaveResult struct {
    Group GroupV2 `json:"group"`
    GroupDeleted bool `json:"groupdeleted"`
}
type GroupBanRequest struct {
    Comment string `json:"comment"`
    Length json.Number `json:"length"`
}
type IgnoreLength struct {
    IgnoreLength json.Number `json:"ignorelength"`
}
type GroupBan struct {
    GroupId json.Number `json:"groupid"`
    LastModifiedBy UserInfoCard `json:"lastmodifiedby"`
    CreatedBy UserInfoCard `json:"createdby"`
    DateBanned string `json:"datebanned"`
    DateExpires string `json:"dateexpires"`
    Comment string `json:"comment"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    DestinyUserInfo GroupUserInfoCard `json:"destinyuserinfo"`
}
type SearchResultOfGroupBan struct {
    Results []GroupBan `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupMemberApplication struct {
    GroupId json.Number `json:"groupid"`
    CreationDate string `json:"creationdate"`
    ResolveState json.Number `json:"resolvestate"`
    ResolveDate string `json:"resolvedate"`
    ResolvedByMembershipId json.Number `json:"resolvedbymembershipid"`
    RequestMessage string `json:"requestmessage"`
    ResolveMessage string `json:"resolvemessage"`
    DestinyUserInfo GroupUserInfoCard `json:"destinyuserinfo"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
}
type GroupApplicationResolveState struct {
    GroupApplicationResolveState json.Number `json:"groupapplicationresolvestate"`
}
type SearchResultOfGroupMemberApplication struct {
    Results []GroupMemberApplication `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type EntityActionResult struct {
    EntityId json.Number `json:"entityid"`
    Result json.Number `json:"result"`
}
type PlatformErrorCodes struct {
    PlatformErrorCodes json.Number `json:"platformerrorcodes"`
}
type GroupApplicationRequest struct {
    Message string `json:"message"`
}
type GroupApplicationListRequest struct {
    Memberships []UserMembership `json:"memberships"`
    Message string `json:"message"`
}
type GroupsForMemberFilter struct {
    GroupsForMemberFilter json.Number `json:"groupsformemberfilter"`
}
type GroupMembershipBase struct {
    Group GroupV2 `json:"group"`
}
type GroupMembership struct {
    Member GroupMember `json:"member"`
    Group GroupV2 `json:"group"`
}
type SearchResultOfGroupMembership struct {
    Results []GroupMembership `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupMembershipSearchResponse struct {
    Results []GroupMembership `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GetGroupsForMemberResponse struct {
    AreAllMembershipsInactive bool `json:"areallmembershipsinactive"`
    Results []GroupMembership `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupPotentialMembership struct {
    Member GroupPotentialMember `json:"member"`
    Group GroupV2 `json:"group"`
}
type SearchResultOfGroupPotentialMembership struct {
    Results []GroupPotentialMembership `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupPotentialMembershipSearchResponse struct {
    Results []GroupPotentialMembership `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type GroupApplicationResponse struct {
    Resolution json.Number `json:"resolution"`
}
type PartnerOfferClaimRequest struct {
    PartnerOfferId string `json:"partnerofferid"`
    BungieNetMembershipId json.Number `json:"bungienetmembershipid"`
    TransactionId string `json:"transactionid"`
}
type PartnerOfferSkuHistoryResponse struct {
    SkuIdentifier string `json:"skuidentifier"`
    LocalizedName string `json:"localizedname"`
    LocalizedDescription string `json:"localizeddescription"`
    ClaimDate string `json:"claimdate"`
    AllOffersApplied bool `json:"alloffersapplied"`
    TransactionId string `json:"transactionid"`
    SkuOffers []PartnerOfferHistoryResponse `json:"skuoffers"`
}
type PartnerOfferHistoryResponse struct {
    PartnerOfferKey string `json:"partnerofferkey"`
    MembershipId json.Number `json:"membershipid"`
    MembershipType json.Number `json:"membershiptype"`
    LocalizedName string `json:"localizedname"`
    LocalizedDescription string `json:"localizeddescription"`
    IsConsumable bool `json:"isconsumable"`
    QuantityApplied json.Number `json:"quantityapplied"`
    ApplyDate string `json:"applydate"`
}
type DestinyManifest struct {
    Version string `json:"version"`
    MobileAssetContentPath string `json:"mobileassetcontentpath"`
    MobileGearAssetDataBases []GearAssetDataBaseDefinition `json:"mobilegearassetdatabases"`
    MobileWorldContentPaths string `json:"mobileworldcontentpaths"`
    JsonWorldContentPaths string `json:"jsonworldcontentpaths"`
    JsonWorldComponentContentPaths interface{} `json:"jsonworldcomponentcontentpaths"`
    MobileClanBannerDatabasePath string `json:"mobileclanbannerdatabasepath"`
    MobileGearCDN string `json:"mobilegearcdn"`
    IconImagePyramidInfo []ImagePyramidEntry `json:"iconimagepyramidinfo"`
}
type GearAssetDataBaseDefinition struct {
    Version json.Number `json:"version"`
    Path string `json:"path"`
}
type ImagePyramidEntry struct {
    Name string `json:"name"`
    Factor float64 `json:"factor"`
}
type DestinyLinkedProfilesResponse struct {
    Profiles []DestinyProfileUserInfoCard `json:"profiles"`
    BnetMembership UserInfoCard `json:"bnetmembership"`
    ProfilesWithErrors []DestinyErrorProfile `json:"profileswitherrors"`
}
type DestinyProfileUserInfoCard struct {
    DateLastPlayed string `json:"datelastplayed"`
    IsOverridden bool `json:"isoverridden"`
    IsCrossSavePrimary bool `json:"iscrosssaveprimary"`
    PlatformSilver DestinyPlatformSilverComponent `json:"platformsilver"`
    UnpairedGameVersions json.Number `json:"unpairedgameversions"`
    SupplementalDisplayName string `json:"supplementaldisplayname"`
    IconPath string `json:"iconpath"`
    CrossSaveOverride json.Number `json:"crosssaveoverride"`
    ApplicableMembershipTypes []int `json:"applicablemembershiptypes"`
    IsPublic bool `json:"ispublic"`
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type DestinyPlatformSilverComponent struct {
    PlatformSilver DestinyItemComponent `json:"platformsilver"`
}
type DestinyItemComponent struct {
    ItemHash json.Number `json:"itemhash"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Quantity json.Number `json:"quantity"`
    BindStatus json.Number `json:"bindstatus"`
    Location json.Number `json:"location"`
    BucketHash json.Number `json:"buckethash"`
    TransferStatus json.Number `json:"transferstatus"`
    Lockable bool `json:"lockable"`
    State json.Number `json:"state"`
    OverrideStyleItemHash json.Number `json:"overridestyleitemhash"`
    ExpirationDate string `json:"expirationdate"`
    IsWrapper bool `json:"iswrapper"`
    TooltipNotificationIndexes []int `json:"tooltipnotificationindexes"`
    MetricHash json.Number `json:"metrichash"`
    MetricObjective DestinyObjectiveProgress `json:"metricobjective"`
    VersionNumber json.Number `json:"versionnumber"`
    ItemValueVisibility []bool `json:"itemvaluevisibility"`
}
type ItemBindStatus struct {
    ItemBindStatus json.Number `json:"itembindstatus"`
}
type TransferStatuses struct {
    TransferStatuses json.Number `json:"transferstatuses"`
}
type ItemState struct {
    ItemState json.Number `json:"itemstate"`
}
type DestinyObjectiveProgress struct {
    ObjectiveHash json.Number `json:"objectivehash"`
    DestinationHash json.Number `json:"destinationhash"`
    ActivityHash json.Number `json:"activityhash"`
    Progress json.Number `json:"progress"`
    CompletionValue json.Number `json:"completionvalue"`
    Complete bool `json:"complete"`
    Visible bool `json:"visible"`
}
type DestinyGameVersions struct {
    DestinyGameVersions json.Number `json:"destinygameversions"`
}
type DestinyErrorProfile struct {
    ErrorCode json.Number `json:"errorcode"`
    InfoCard UserInfoCard `json:"infocard"`
}
type DestinyComponentType struct {
    DestinyComponentType json.Number `json:"destinycomponenttype"`
}
type DestinyProfileResponse struct {
    VendorReceipts SingleComponentResponseOfDestinyVendorReceiptsComponent `json:"vendorreceipts"`
    ProfileInventory SingleComponentResponseOfDestinyInventoryComponent `json:"profileinventory"`
    ProfileCurrencies SingleComponentResponseOfDestinyInventoryComponent `json:"profilecurrencies"`
    Profile SingleComponentResponseOfDestinyProfileComponent `json:"profile"`
    PlatformSilver SingleComponentResponseOfDestinyPlatformSilverComponent `json:"platformsilver"`
    ProfileKiosks SingleComponentResponseOfDestinyKiosksComponent `json:"profilekiosks"`
    ProfilePlugSets SingleComponentResponseOfDestinyPlugSetsComponent `json:"profileplugsets"`
    ProfileProgression SingleComponentResponseOfDestinyProfileProgressionComponent `json:"profileprogression"`
    ProfilePresentationNodes SingleComponentResponseOfDestinyPresentationNodesComponent `json:"profilepresentationnodes"`
    ProfileRecords SingleComponentResponseOfDestinyProfileRecordsComponent `json:"profilerecords"`
    ProfileCollectibles SingleComponentResponseOfDestinyProfileCollectiblesComponent `json:"profilecollectibles"`
    ProfileTransitoryData SingleComponentResponseOfDestinyProfileTransitoryComponent `json:"profiletransitorydata"`
    Metrics SingleComponentResponseOfDestinyMetricsComponent `json:"metrics"`
    ProfileStringVariables SingleComponentResponseOfDestinyStringVariablesComponent `json:"profilestringvariables"`
    Characters DictionaryComponentResponseOfint64AndDestinyCharacterComponent `json:"characters"`
    CharacterInventories DictionaryComponentResponseOfint64AndDestinyInventoryComponent `json:"characterinventories"`
    CharacterProgressions DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent `json:"characterprogressions"`
    CharacterRenderData DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent `json:"characterrenderdata"`
    CharacterActivities DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent `json:"characteractivities"`
    CharacterEquipment DictionaryComponentResponseOfint64AndDestinyInventoryComponent `json:"characterequipment"`
    CharacterKiosks DictionaryComponentResponseOfint64AndDestinyKiosksComponent `json:"characterkiosks"`
    CharacterPlugSets DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent `json:"characterplugsets"`
    CharacterUninstancedItemComponents DestinyBaseItemComponentSetOfuint32 `json:"characteruninstanceditemcomponents"`
    CharacterPresentationNodes DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent `json:"characterpresentationnodes"`
    CharacterRecords DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent `json:"characterrecords"`
    CharacterCollectibles DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent `json:"charactercollectibles"`
    CharacterStringVariables DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent `json:"characterstringvariables"`
    ItemComponents DestinyItemComponentSetOfint64 `json:"itemcomponents"`
    CharacterCurrencyLookups DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent `json:"charactercurrencylookups"`
}
type DestinyVendorReceiptsComponent struct {
    Receipts []DestinyVendorReceipt `json:"receipts"`
}
type DestinyVendorReceipt struct {
    CurrencyPaid []DestinyItemQuantity `json:"currencypaid"`
    ItemReceived DestinyItemQuantity `json:"itemreceived"`
    LicenseUnlockHash json.Number `json:"licenseunlockhash"`
    PurchasedByCharacterId json.Number `json:"purchasedbycharacterid"`
    RefundPolicy json.Number `json:"refundpolicy"`
    SequenceNumber json.Number `json:"sequencenumber"`
    TimeToExpiration json.Number `json:"timetoexpiration"`
    ExpiresOn string `json:"expireson"`
}
type ComponentResponse struct {
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type ComponentPrivacySetting struct {
    ComponentPrivacySetting json.Number `json:"componentprivacysetting"`
}
type SingleComponentResponseOfDestinyVendorReceiptsComponent struct {
    Data DestinyVendorReceiptsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyInventoryComponent struct {
    Items []DestinyItemComponent `json:"items"`
}
type SingleComponentResponseOfDestinyInventoryComponent struct {
    Data DestinyInventoryComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyProfileComponent struct {
    UserInfo UserInfoCard `json:"userinfo"`
    DateLastPlayed string `json:"datelastplayed"`
    VersionsOwned json.Number `json:"versionsowned"`
    CharacterIds []int `json:"characterids"`
    SeasonHashes []int `json:"seasonhashes"`
    CurrentSeasonHash json.Number `json:"currentseasonhash"`
    CurrentSeasonRewardPowerCap json.Number `json:"currentseasonrewardpowercap"`
}
type SingleComponentResponseOfDestinyProfileComponent struct {
    Data DestinyProfileComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyPlatformSilverComponent struct {
    Data DestinyPlatformSilverComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyKiosksComponent struct {
    KioskItems []DestinyKioskItem `json:"kioskitems"`
}
type DestinyKioskItem struct {
    Index json.Number `json:"index"`
    CanAcquire bool `json:"canacquire"`
    FailureIndexes []int `json:"failureindexes"`
    FlavorObjective DestinyObjectiveProgress `json:"flavorobjective"`
}
type SingleComponentResponseOfDestinyKiosksComponent struct {
    Data DestinyKiosksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyPlugSetsComponent struct {
    Plugs []DestinyItemPlug `json:"plugs"`
}
type DestinyItemPlugBase struct {
    PlugItemHash json.Number `json:"plugitemhash"`
    CanInsert bool `json:"caninsert"`
    Enabled bool `json:"enabled"`
    InsertFailIndexes []int `json:"insertfailindexes"`
    EnableFailIndexes []int `json:"enablefailindexes"`
}
type DestinyItemPlug struct {
    PlugObjectives []DestinyObjectiveProgress `json:"plugobjectives"`
    PlugItemHash json.Number `json:"plugitemhash"`
    CanInsert bool `json:"caninsert"`
    Enabled bool `json:"enabled"`
    InsertFailIndexes []int `json:"insertfailindexes"`
    EnableFailIndexes []int `json:"enablefailindexes"`
}
type SingleComponentResponseOfDestinyPlugSetsComponent struct {
    Data DestinyPlugSetsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyProfileProgressionComponent struct {
    Checklists interface{} `json:"checklists"`
    SeasonalArtifact DestinyArtifactProfileScoped `json:"seasonalartifact"`
}
type DestinyArtifactProfileScoped struct {
    ArtifactHash json.Number `json:"artifacthash"`
    PointProgression DestinyProgression `json:"pointprogression"`
    PointsAcquired json.Number `json:"pointsacquired"`
    PowerBonusProgression DestinyProgression `json:"powerbonusprogression"`
    PowerBonus json.Number `json:"powerbonus"`
}
type DestinyChecklistDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    ViewActionString string `json:"viewactionstring"`
    Scope json.Number `json:"scope"`
    Entries []DestinyChecklistEntryDefinition `json:"entries"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyChecklistEntryDefinition struct {
    Hash json.Number `json:"hash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    DestinationHash json.Number `json:"destinationhash"`
    LocationHash json.Number `json:"locationhash"`
    BubbleHash json.Number `json:"bubblehash"`
    ActivityHash json.Number `json:"activityhash"`
    ItemHash json.Number `json:"itemhash"`
    VendorHash json.Number `json:"vendorhash"`
    VendorInteractionIndex json.Number `json:"vendorinteractionindex"`
    Scope json.Number `json:"scope"`
}
type SingleComponentResponseOfDestinyProfileProgressionComponent struct {
    Data DestinyProfileProgressionComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyPresentationNodesComponent struct {
    Nodes DestinyPresentationNodeComponent `json:"nodes"`
}
type DestinyPresentationNodeComponent struct {
    State json.Number `json:"state"`
    Objective DestinyObjectiveProgress `json:"objective"`
    ProgressValue json.Number `json:"progressvalue"`
    CompletionValue json.Number `json:"completionvalue"`
    RecordCategoryScore json.Number `json:"recordcategoryscore"`
}
type DestinyPresentationNodeState struct {
    DestinyPresentationNodeState json.Number `json:"destinypresentationnodestate"`
}
type SingleComponentResponseOfDestinyPresentationNodesComponent struct {
    Data DestinyPresentationNodesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyRecordsComponent struct {
    Records DestinyRecordComponent `json:"records"`
    RecordCategoriesRootNodeHash json.Number `json:"recordcategoriesrootnodehash"`
    RecordSealsRootNodeHash json.Number `json:"recordsealsrootnodehash"`
}
type DestinyRecordComponent struct {
    State json.Number `json:"state"`
    Objectives []DestinyObjectiveProgress `json:"objectives"`
    IntervalObjectives []DestinyObjectiveProgress `json:"intervalobjectives"`
    IntervalsRedeemedCount json.Number `json:"intervalsredeemedcount"`
    CompletedCount json.Number `json:"completedcount"`
    RewardVisibilty []bool `json:"rewardvisibilty"`
}
type DestinyRecordState struct {
    DestinyRecordState json.Number `json:"destinyrecordstate"`
}
type DestinyProfileRecordsComponent struct {
    Score json.Number `json:"score"`
    ActiveScore json.Number `json:"activescore"`
    LegacyScore json.Number `json:"legacyscore"`
    LifetimeScore json.Number `json:"lifetimescore"`
    TrackedRecordHash json.Number `json:"trackedrecordhash"`
    Records DestinyRecordComponent `json:"records"`
    RecordCategoriesRootNodeHash json.Number `json:"recordcategoriesrootnodehash"`
    RecordSealsRootNodeHash json.Number `json:"recordsealsrootnodehash"`
}
type SingleComponentResponseOfDestinyProfileRecordsComponent struct {
    Data DestinyProfileRecordsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCollectiblesComponent struct {
    Collectibles DestinyCollectibleComponent `json:"collectibles"`
    CollectionCategoriesRootNodeHash json.Number `json:"collectioncategoriesrootnodehash"`
    CollectionBadgesRootNodeHash json.Number `json:"collectionbadgesrootnodehash"`
}
type DestinyCollectibleComponent struct {
    State json.Number `json:"state"`
}
type DestinyCollectibleState struct {
    DestinyCollectibleState json.Number `json:"destinycollectiblestate"`
}
type DestinyProfileCollectiblesComponent struct {
    RecentCollectibleHashes []int `json:"recentcollectiblehashes"`
    NewnessFlaggedCollectibleHashes []int `json:"newnessflaggedcollectiblehashes"`
    Collectibles DestinyCollectibleComponent `json:"collectibles"`
    CollectionCategoriesRootNodeHash json.Number `json:"collectioncategoriesrootnodehash"`
    CollectionBadgesRootNodeHash json.Number `json:"collectionbadgesrootnodehash"`
}
type SingleComponentResponseOfDestinyProfileCollectiblesComponent struct {
    Data DestinyProfileCollectiblesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyProfileTransitoryComponent struct {
    PartyMembers []DestinyProfileTransitoryPartyMember `json:"partymembers"`
    CurrentActivity DestinyProfileTransitoryCurrentActivity `json:"currentactivity"`
    Joinability DestinyProfileTransitoryJoinability `json:"joinability"`
    Tracking []DestinyProfileTransitoryTrackingEntry `json:"tracking"`
    LastOrbitedDestinationHash json.Number `json:"lastorbiteddestinationhash"`
}
type DestinyProfileTransitoryPartyMember struct {
    MembershipId json.Number `json:"membershipid"`
    EmblemHash json.Number `json:"emblemhash"`
    DisplayName string `json:"displayname"`
    Status json.Number `json:"status"`
}
type DestinyPartyMemberStates struct {
    DestinyPartyMemberStates json.Number `json:"destinypartymemberstates"`
}
type DestinyProfileTransitoryCurrentActivity struct {
    StartTime string `json:"starttime"`
    EndTime string `json:"endtime"`
    Score float64 `json:"score"`
    HighestOpposingFactionScore float64 `json:"highestopposingfactionscore"`
    NumberOfOpponents json.Number `json:"numberofopponents"`
    NumberOfPlayers json.Number `json:"numberofplayers"`
}
type DestinyProfileTransitoryJoinability struct {
    OpenSlots json.Number `json:"openslots"`
    PrivacySetting json.Number `json:"privacysetting"`
    ClosedReasons json.Number `json:"closedreasons"`
}
type DestinyGamePrivacySetting struct {
    DestinyGamePrivacySetting json.Number `json:"destinygameprivacysetting"`
}
type DestinyJoinClosedReasons struct {
    DestinyJoinClosedReasons json.Number `json:"destinyjoinclosedreasons"`
}
type DestinyProfileTransitoryTrackingEntry struct {
    LocationHash json.Number `json:"locationhash"`
    ItemHash json.Number `json:"itemhash"`
    ObjectiveHash json.Number `json:"objectivehash"`
    ActivityHash json.Number `json:"activityhash"`
    QuestlineItemHash json.Number `json:"questlineitemhash"`
    TrackedDate string `json:"trackeddate"`
}
type SingleComponentResponseOfDestinyProfileTransitoryComponent struct {
    Data DestinyProfileTransitoryComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyMetricsComponent struct {
    Metrics DestinyMetricComponent `json:"metrics"`
    MetricsRootNodeHash json.Number `json:"metricsrootnodehash"`
}
type DestinyMetricComponent struct {
    Invisible bool `json:"invisible"`
    ObjectiveProgress DestinyObjectiveProgress `json:"objectiveprogress"`
}
type SingleComponentResponseOfDestinyMetricsComponent struct {
    Data DestinyMetricsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyStringVariablesComponent struct {
    IntegerValuesByHash int `json:"integervaluesbyhash"`
}
type SingleComponentResponseOfDestinyStringVariablesComponent struct {
    Data DestinyStringVariablesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterComponent struct {
    MembershipId json.Number `json:"membershipid"`
    MembershipType json.Number `json:"membershiptype"`
    CharacterId json.Number `json:"characterid"`
    DateLastPlayed string `json:"datelastplayed"`
    MinutesPlayedThisSession json.Number `json:"minutesplayedthissession"`
    MinutesPlayedTotal json.Number `json:"minutesplayedtotal"`
    Light json.Number `json:"light"`
    Stats int `json:"stats"`
    RaceHash json.Number `json:"racehash"`
    GenderHash json.Number `json:"genderhash"`
    ClassHash json.Number `json:"classhash"`
    RaceType json.Number `json:"racetype"`
    ClassType json.Number `json:"classtype"`
    GenderType json.Number `json:"gendertype"`
    EmblemPath string `json:"emblempath"`
    EmblemBackgroundPath string `json:"emblembackgroundpath"`
    EmblemHash json.Number `json:"emblemhash"`
    EmblemColor DestinyColor `json:"emblemcolor"`
    LevelProgression DestinyProgression `json:"levelprogression"`
    BaseCharacterLevel json.Number `json:"basecharacterlevel"`
    PercentToNextLevel float64 `json:"percenttonextlevel"`
    TitleRecordHash json.Number `json:"titlerecordhash"`
}
type DestinyRace struct {
    DestinyRace json.Number `json:"destinyrace"`
}
type DestinyRaceDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    RaceType json.Number `json:"racetype"`
    GenderedRaceNames string `json:"genderedracenames"`
    GenderedRaceNamesByGenderHash string `json:"genderedracenamesbygenderhash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyClassDefinition struct {
    ClassType json.Number `json:"classtype"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    GenderedClassNames string `json:"genderedclassnames"`
    GenderedClassNamesByGenderHash string `json:"genderedclassnamesbygenderhash"`
    MentorVendorHash json.Number `json:"mentorvendorhash"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterComponent struct {
    Data DestinyCharacterComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyInventoryComponent struct {
    Data DestinyInventoryComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterProgressionComponent struct {
    Progressions DestinyProgression `json:"progressions"`
    Factions DestinyFactionProgression `json:"factions"`
    Milestones DestinyMilestone `json:"milestones"`
    Quests []DestinyQuestStatus `json:"quests"`
    UninstancedItemObjectives []DestinyObjectiveProgress `json:"uninstanceditemobjectives"`
    UninstancedItemPerks DestinyItemPerksComponent `json:"uninstanceditemperks"`
    Checklists interface{} `json:"checklists"`
    SeasonalArtifact DestinyArtifactCharacterScoped `json:"seasonalartifact"`
}
type DestinyFactionProgression struct {
    FactionHash json.Number `json:"factionhash"`
    FactionVendorIndex json.Number `json:"factionvendorindex"`
    ProgressionHash json.Number `json:"progressionhash"`
    DailyProgress json.Number `json:"dailyprogress"`
    DailyLimit json.Number `json:"dailylimit"`
    WeeklyProgress json.Number `json:"weeklyprogress"`
    WeeklyLimit json.Number `json:"weeklylimit"`
    CurrentProgress json.Number `json:"currentprogress"`
    Level json.Number `json:"level"`
    LevelCap json.Number `json:"levelcap"`
    StepIndex json.Number `json:"stepindex"`
    ProgressToNextLevel json.Number `json:"progresstonextlevel"`
    NextLevelAt json.Number `json:"nextlevelat"`
    CurrentResetCount json.Number `json:"currentresetcount"`
    SeasonResets []DestinyProgressionResetEntry `json:"seasonresets"`
    RewardItemStates []int `json:"rewarditemstates"`
}
type DestinyMilestone struct {
    MilestoneHash json.Number `json:"milestonehash"`
    AvailableQuests []DestinyMilestoneQuest `json:"availablequests"`
    Activities []DestinyMilestoneChallengeActivity `json:"activities"`
    Values float64 `json:"values"`
    VendorHashes []int `json:"vendorhashes"`
    Vendors []DestinyMilestoneVendor `json:"vendors"`
    Rewards []DestinyMilestoneRewardCategory `json:"rewards"`
    StartDate string `json:"startdate"`
    EndDate string `json:"enddate"`
    Order json.Number `json:"order"`
}
type DestinyMilestoneQuest struct {
    QuestItemHash json.Number `json:"questitemhash"`
    Status DestinyQuestStatus `json:"status"`
    Activity DestinyMilestoneActivity `json:"activity"`
    Challenges []DestinyChallengeStatus `json:"challenges"`
}
type DestinyQuestStatus struct {
    QuestHash json.Number `json:"questhash"`
    StepHash json.Number `json:"stephash"`
    StepObjectives []DestinyObjectiveProgress `json:"stepobjectives"`
    Tracked bool `json:"tracked"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Completed bool `json:"completed"`
    Redeemed bool `json:"redeemed"`
    Started bool `json:"started"`
    VendorHash json.Number `json:"vendorhash"`
}
type DestinyMilestoneActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    ActivityModeHash json.Number `json:"activitymodehash"`
    ActivityModeType json.Number `json:"activitymodetype"`
    ModifierHashes []int `json:"modifierhashes"`
    Variants []DestinyMilestoneActivityVariant `json:"variants"`
}
type DestinyMilestoneActivityVariant struct {
    ActivityHash json.Number `json:"activityhash"`
    CompletionStatus DestinyMilestoneActivityCompletionStatus `json:"completionstatus"`
    ActivityModeHash json.Number `json:"activitymodehash"`
    ActivityModeType json.Number `json:"activitymodetype"`
}
type DestinyMilestoneActivityCompletionStatus struct {
    Completed bool `json:"completed"`
    Phases []DestinyMilestoneActivityPhase `json:"phases"`
}
type DestinyMilestoneActivityPhase struct {
    Complete bool `json:"complete"`
    PhaseHash json.Number `json:"phasehash"`
}
type DestinyChallengeStatus struct {
    Objective DestinyObjectiveProgress `json:"objective"`
}
type DestinyMilestoneChallengeActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    Challenges []DestinyChallengeStatus `json:"challenges"`
    ModifierHashes []int `json:"modifierhashes"`
    BooleanActivityOptions bool `json:"booleanactivityoptions"`
    LoadoutRequirementIndex json.Number `json:"loadoutrequirementindex"`
    Phases []DestinyMilestoneActivityPhase `json:"phases"`
}
type DestinyMilestoneVendor struct {
    VendorHash json.Number `json:"vendorhash"`
    PreviewItemHash json.Number `json:"previewitemhash"`
}
type DestinyMilestoneRewardCategory struct {
    RewardCategoryHash json.Number `json:"rewardcategoryhash"`
    Entries []DestinyMilestoneRewardEntry `json:"entries"`
}
type DestinyMilestoneRewardEntry struct {
    RewardEntryHash json.Number `json:"rewardentryhash"`
    Earned bool `json:"earned"`
    Redeemed bool `json:"redeemed"`
}
type DestinyMilestoneDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    DisplayPreference json.Number `json:"displaypreference"`
    Image string `json:"image"`
    MilestoneType json.Number `json:"milestonetype"`
    Recruitable bool `json:"recruitable"`
    FriendlyName string `json:"friendlyname"`
    ShowInExplorer bool `json:"showinexplorer"`
    ShowInMilestones bool `json:"showinmilestones"`
    ExplorePrioritizesActivityImage bool `json:"exploreprioritizesactivityimage"`
    HasPredictableDates bool `json:"haspredictabledates"`
    Quests DestinyMilestoneQuestDefinition `json:"quests"`
    Rewards DestinyMilestoneRewardCategoryDefinition `json:"rewards"`
    VendorsDisplayTitle string `json:"vendorsdisplaytitle"`
    Vendors []DestinyMilestoneVendorDefinition `json:"vendors"`
    Values DestinyMilestoneValueDefinition `json:"values"`
    IsInGameMilestone bool `json:"isingamemilestone"`
    Activities []DestinyMilestoneChallengeActivityDefinition `json:"activities"`
    DefaultOrder json.Number `json:"defaultorder"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyMilestoneDisplayPreference struct {
    DestinyMilestoneDisplayPreference json.Number `json:"destinymilestonedisplaypreference"`
}
type DestinyMilestoneType struct {
    DestinyMilestoneType json.Number `json:"destinymilestonetype"`
}
type DestinyMilestoneQuestDefinition struct {
    QuestItemHash json.Number `json:"questitemhash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    OverrideImage string `json:"overrideimage"`
    QuestRewards DestinyMilestoneQuestRewardsDefinition `json:"questrewards"`
    Activities DestinyMilestoneActivityDefinition `json:"activities"`
    DestinationHash json.Number `json:"destinationhash"`
}
type DestinyMilestoneQuestRewardsDefinition struct {
    Items []DestinyMilestoneQuestRewardItem `json:"items"`
}
type DestinyMilestoneQuestRewardItem struct {
    VendorHash json.Number `json:"vendorhash"`
    VendorItemIndex json.Number `json:"vendoritemindex"`
    ItemHash json.Number `json:"itemhash"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Quantity json.Number `json:"quantity"`
    HasConditionalVisibility bool `json:"hasconditionalvisibility"`
}
type DestinyMilestoneActivityDefinition struct {
    ConceptualActivityHash json.Number `json:"conceptualactivityhash"`
    Variants DestinyMilestoneActivityVariantDefinition `json:"variants"`
}
type DestinyMilestoneActivityVariantDefinition struct {
    ActivityHash json.Number `json:"activityhash"`
    Order json.Number `json:"order"`
}
type DestinyMilestoneRewardCategoryDefinition struct {
    CategoryHash json.Number `json:"categoryhash"`
    CategoryIdentifier string `json:"categoryidentifier"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    RewardEntries DestinyMilestoneRewardEntryDefinition `json:"rewardentries"`
    Order json.Number `json:"order"`
}
type DestinyMilestoneRewardEntryDefinition struct {
    RewardEntryHash json.Number `json:"rewardentryhash"`
    RewardEntryIdentifier string `json:"rewardentryidentifier"`
    Items []DestinyItemQuantity `json:"items"`
    VendorHash json.Number `json:"vendorhash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Order json.Number `json:"order"`
}
type DestinyMilestoneVendorDefinition struct {
    VendorHash json.Number `json:"vendorhash"`
}
type DestinyMilestoneValueDefinition struct {
    Key string `json:"key"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyMilestoneChallengeActivityDefinition struct {
    ActivityHash json.Number `json:"activityhash"`
    Challenges []DestinyMilestoneChallengeDefinition `json:"challenges"`
    ActivityGraphNodes []DestinyMilestoneChallengeActivityGraphNodeEntry `json:"activitygraphnodes"`
    Phases []DestinyMilestoneChallengeActivityPhase `json:"phases"`
}
type DestinyMilestoneChallengeDefinition struct {
    ChallengeObjectiveHash json.Number `json:"challengeobjectivehash"`
}
type DestinyMilestoneChallengeActivityGraphNodeEntry struct {
    ActivityGraphHash json.Number `json:"activitygraphhash"`
    ActivityGraphNodeHash json.Number `json:"activitygraphnodehash"`
}
type DestinyMilestoneChallengeActivityPhase struct {
    PhaseHash json.Number `json:"phasehash"`
}
type DestinyItemPerksComponent struct {
    Perks []DestinyPerkReference `json:"perks"`
}
type DestinyPerkReference struct {
    PerkHash json.Number `json:"perkhash"`
    IconPath string `json:"iconpath"`
    IsActive bool `json:"isactive"`
    Visible bool `json:"visible"`
}
type DestinyArtifactCharacterScoped struct {
    ArtifactHash json.Number `json:"artifacthash"`
    PointsUsed json.Number `json:"pointsused"`
    ResetCount json.Number `json:"resetcount"`
    Tiers []DestinyArtifactTier `json:"tiers"`
}
type DestinyArtifactTier struct {
    TierHash json.Number `json:"tierhash"`
    IsUnlocked bool `json:"isunlocked"`
    PointsToUnlock json.Number `json:"pointstounlock"`
    Items []DestinyArtifactTierItem `json:"items"`
}
type DestinyArtifactTierItem struct {
    ItemHash json.Number `json:"itemhash"`
    IsActive bool `json:"isactive"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterProgressionComponent struct {
    Data DestinyCharacterProgressionComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterRenderComponent struct {
    CustomDyes []DyeReference `json:"customdyes"`
    Customization DestinyCharacterCustomization `json:"customization"`
    PeerView DestinyCharacterPeerView `json:"peerview"`
}
type DestinyCharacterCustomization struct {
    Personality json.Number `json:"personality"`
    Face json.Number `json:"face"`
    SkinColor json.Number `json:"skincolor"`
    LipColor json.Number `json:"lipcolor"`
    EyeColor json.Number `json:"eyecolor"`
    HairColors []int `json:"haircolors"`
    FeatureColors []int `json:"featurecolors"`
    DecalColor json.Number `json:"decalcolor"`
    WearHelmet bool `json:"wearhelmet"`
    HairIndex json.Number `json:"hairindex"`
    FeatureIndex json.Number `json:"featureindex"`
    DecalIndex json.Number `json:"decalindex"`
}
type DestinyCharacterPeerView struct {
    Equipment []DestinyItemPeerView `json:"equipment"`
}
type DestinyItemPeerView struct {
    ItemHash json.Number `json:"itemhash"`
    Dyes []DyeReference `json:"dyes"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterRenderComponent struct {
    Data DestinyCharacterRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterActivitiesComponent struct {
    DateActivityStarted string `json:"dateactivitystarted"`
    AvailableActivities []DestinyActivity `json:"availableactivities"`
    CurrentActivityHash json.Number `json:"currentactivityhash"`
    CurrentActivityModeHash json.Number `json:"currentactivitymodehash"`
    CurrentActivityModeType json.Number `json:"currentactivitymodetype"`
    CurrentActivityModeHashes []int `json:"currentactivitymodehashes"`
    CurrentActivityModeTypes []int `json:"currentactivitymodetypes"`
    CurrentPlaylistActivityHash json.Number `json:"currentplaylistactivityhash"`
    LastCompletedStoryHash json.Number `json:"lastcompletedstoryhash"`
}
type DestinyActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    IsNew bool `json:"isnew"`
    CanLead bool `json:"canlead"`
    CanJoin bool `json:"canjoin"`
    IsCompleted bool `json:"iscompleted"`
    IsVisible bool `json:"isvisible"`
    DisplayLevel json.Number `json:"displaylevel"`
    RecommendedLight json.Number `json:"recommendedlight"`
    DifficultyTier json.Number `json:"difficultytier"`
    Challenges []DestinyChallengeStatus `json:"challenges"`
    ModifierHashes []int `json:"modifierhashes"`
    BooleanActivityOptions bool `json:"booleanactivityoptions"`
    LoadoutRequirementIndex json.Number `json:"loadoutrequirementindex"`
}
type DestinyActivityDifficultyTier struct {
    DestinyActivityDifficultyTier json.Number `json:"destinyactivitydifficultytier"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterActivitiesComponent struct {
    Data DestinyCharacterActivitiesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyKiosksComponent struct {
    Data DestinyKiosksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyPlugSetsComponent struct {
    Data DestinyPlugSetsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyBaseItemComponentSetOfuint32 struct {
    Objectives DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent `json:"perks"`
}
type DestinyItemObjectivesComponent struct {
    Objectives []DestinyObjectiveProgress `json:"objectives"`
    FlavorObjective DestinyObjectiveProgress `json:"flavorobjective"`
    DateCompleted string `json:"datecompleted"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent struct {
    Data DestinyItemObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent struct {
    Data DestinyItemPerksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyPresentationNodesComponent struct {
    Data DestinyPresentationNodesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterRecordsComponent struct {
    FeaturedRecordHashes []int `json:"featuredrecordhashes"`
    Records DestinyRecordComponent `json:"records"`
    RecordCategoriesRootNodeHash json.Number `json:"recordcategoriesrootnodehash"`
    RecordSealsRootNodeHash json.Number `json:"recordsealsrootnodehash"`
}
type DictionaryComponentResponseOfint64AndDestinyCharacterRecordsComponent struct {
    Data DestinyCharacterRecordsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyCollectiblesComponent struct {
    Data DestinyCollectiblesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyStringVariablesComponent struct {
    Data DestinyStringVariablesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyBaseItemComponentSetOfint64 struct {
    Objectives DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfint64AndDestinyItemPerksComponent `json:"perks"`
}
type DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent struct {
    Data DestinyItemObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint64AndDestinyItemPerksComponent struct {
    Data DestinyItemPerksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemComponentSetOfint64 struct {
    Instances DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent `json:"instances"`
    RenderData DictionaryComponentResponseOfint64AndDestinyItemRenderComponent `json:"renderdata"`
    Stats DictionaryComponentResponseOfint64AndDestinyItemStatsComponent `json:"stats"`
    Sockets DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent `json:"sockets"`
    ReusablePlugs DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent `json:"reusableplugs"`
    PlugObjectives DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
    TalentGrids DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent `json:"talentgrids"`
    PlugStates DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent `json:"plugstates"`
    Objectives DictionaryComponentResponseOfint64AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfint64AndDestinyItemPerksComponent `json:"perks"`
}
type DestinyItemInstanceComponent struct {
    DamageType json.Number `json:"damagetype"`
    DamageTypeHash json.Number `json:"damagetypehash"`
    PrimaryStat DestinyStat `json:"primarystat"`
    ItemLevel json.Number `json:"itemlevel"`
    Quality json.Number `json:"quality"`
    IsEquipped bool `json:"isequipped"`
    CanEquip bool `json:"canequip"`
    EquipRequiredLevel json.Number `json:"equiprequiredlevel"`
    UnlockHashesRequiredToEquip []int `json:"unlockhashesrequiredtoequip"`
    CannotEquipReason json.Number `json:"cannotequipreason"`
    BreakerType json.Number `json:"breakertype"`
    BreakerTypeHash json.Number `json:"breakertypehash"`
    Energy DestinyItemInstanceEnergy `json:"energy"`
}
type DestinyStat struct {
    StatHash json.Number `json:"stathash"`
    Value json.Number `json:"value"`
}
type EquipFailureReason struct {
    EquipFailureReason json.Number `json:"equipfailurereason"`
}
type DestinyItemInstanceEnergy struct {
    EnergyTypeHash json.Number `json:"energytypehash"`
    EnergyType json.Number `json:"energytype"`
    EnergyCapacity json.Number `json:"energycapacity"`
    EnergyUsed json.Number `json:"energyused"`
    EnergyUnused json.Number `json:"energyunused"`
}
type DestinyUnlockDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DictionaryComponentResponseOfint64AndDestinyItemInstanceComponent struct {
    Data DestinyItemInstanceComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemRenderComponent struct {
    UseCustomDyes bool `json:"usecustomdyes"`
    ArtRegions int `json:"artregions"`
}
type DictionaryComponentResponseOfint64AndDestinyItemRenderComponent struct {
    Data DestinyItemRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemStatsComponent struct {
    Stats DestinyStat `json:"stats"`
}
type DictionaryComponentResponseOfint64AndDestinyItemStatsComponent struct {
    Data DestinyItemStatsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemSocketsComponent struct {
    Sockets []DestinyItemSocketState `json:"sockets"`
}
type DestinyItemSocketState struct {
    PlugHash json.Number `json:"plughash"`
    IsEnabled bool `json:"isenabled"`
    IsVisible bool `json:"isvisible"`
    EnableFailIndexes []int `json:"enablefailindexes"`
}
type DictionaryComponentResponseOfint64AndDestinyItemSocketsComponent struct {
    Data DestinyItemSocketsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemReusablePlugsComponent struct {
    Plugs []DestinyItemPlugBase `json:"plugs"`
}
type DictionaryComponentResponseOfint64AndDestinyItemReusablePlugsComponent struct {
    Data DestinyItemReusablePlugsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemPlugObjectivesComponent struct {
    ObjectivesPerPlug []DestinyObjectiveProgress `json:"objectivesperplug"`
}
type DictionaryComponentResponseOfint64AndDestinyItemPlugObjectivesComponent struct {
    Data DestinyItemPlugObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemTalentGridComponent struct {
    TalentGridHash json.Number `json:"talentgridhash"`
    Nodes []DestinyTalentNode `json:"nodes"`
    IsGridComplete bool `json:"isgridcomplete"`
    GridProgression DestinyProgression `json:"gridprogression"`
}
type DestinyTalentNode struct {
    NodeIndex json.Number `json:"nodeindex"`
    NodeHash json.Number `json:"nodehash"`
    State json.Number `json:"state"`
    IsActivated bool `json:"isactivated"`
    StepIndex json.Number `json:"stepindex"`
    MaterialsToUpgrade []DestinyMaterialRequirement `json:"materialstoupgrade"`
    ActivationGridLevel json.Number `json:"activationgridlevel"`
    ProgressPercent float64 `json:"progresspercent"`
    Hidden bool `json:"hidden"`
    NodeStatsBlock DestinyTalentNodeStatBlock `json:"nodestatsblock"`
}
type DestinyTalentNodeState struct {
    DestinyTalentNodeState json.Number `json:"destinytalentnodestate"`
}
type DestinyTalentNodeStatBlock struct {
    CurrentStepStats []DestinyStat `json:"currentstepstats"`
    NextStepStats []DestinyStat `json:"nextstepstats"`
}
type DictionaryComponentResponseOfint64AndDestinyItemTalentGridComponent struct {
    Data DestinyItemTalentGridComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemPlugComponent struct {
    PlugObjectives []DestinyObjectiveProgress `json:"plugobjectives"`
    PlugItemHash json.Number `json:"plugitemhash"`
    CanInsert bool `json:"caninsert"`
    Enabled bool `json:"enabled"`
    InsertFailIndexes []int `json:"insertfailindexes"`
    EnableFailIndexes []int `json:"enablefailindexes"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent struct {
    Data DestinyItemPlugComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCurrenciesComponent struct {
    ItemQuantities int `json:"itemquantities"`
}
type DictionaryComponentResponseOfint64AndDestinyCurrenciesComponent struct {
    Data DestinyCurrenciesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCharacterResponse struct {
    Inventory SingleComponentResponseOfDestinyInventoryComponent `json:"inventory"`
    Character SingleComponentResponseOfDestinyCharacterComponent `json:"character"`
    Progressions SingleComponentResponseOfDestinyCharacterProgressionComponent `json:"progressions"`
    RenderData SingleComponentResponseOfDestinyCharacterRenderComponent `json:"renderdata"`
    Activities SingleComponentResponseOfDestinyCharacterActivitiesComponent `json:"activities"`
    Equipment SingleComponentResponseOfDestinyInventoryComponent `json:"equipment"`
    Kiosks SingleComponentResponseOfDestinyKiosksComponent `json:"kiosks"`
    PlugSets SingleComponentResponseOfDestinyPlugSetsComponent `json:"plugsets"`
    PresentationNodes SingleComponentResponseOfDestinyPresentationNodesComponent `json:"presentationnodes"`
    Records SingleComponentResponseOfDestinyCharacterRecordsComponent `json:"records"`
    Collectibles SingleComponentResponseOfDestinyCollectiblesComponent `json:"collectibles"`
    ItemComponents DestinyItemComponentSetOfint64 `json:"itemcomponents"`
    UninstancedItemComponents DestinyBaseItemComponentSetOfuint32 `json:"uninstanceditemcomponents"`
    CurrencyLookups SingleComponentResponseOfDestinyCurrenciesComponent `json:"currencylookups"`
}
type SingleComponentResponseOfDestinyCharacterComponent struct {
    Data DestinyCharacterComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterProgressionComponent struct {
    Data DestinyCharacterProgressionComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterRenderComponent struct {
    Data DestinyCharacterRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterActivitiesComponent struct {
    Data DestinyCharacterActivitiesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCharacterRecordsComponent struct {
    Data DestinyCharacterRecordsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCollectiblesComponent struct {
    Data DestinyCollectiblesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyCurrenciesComponent struct {
    Data DestinyCurrenciesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type ClanBannerSource struct {
    ClanBannerSource interface{} `json:"clanbannersource"`
}
type ClanBannerDecal struct {
    Identifier string `json:"identifier"`
    ForegroundPath string `json:"foregroundpath"`
    BackgroundPath string `json:"backgroundpath"`
}
type DestinyItemResponse struct {
    CharacterId json.Number `json:"characterid"`
    Item SingleComponentResponseOfDestinyItemComponent `json:"item"`
    Instance SingleComponentResponseOfDestinyItemInstanceComponent `json:"instance"`
    Objectives SingleComponentResponseOfDestinyItemObjectivesComponent `json:"objectives"`
    Perks SingleComponentResponseOfDestinyItemPerksComponent `json:"perks"`
    RenderData SingleComponentResponseOfDestinyItemRenderComponent `json:"renderdata"`
    Stats SingleComponentResponseOfDestinyItemStatsComponent `json:"stats"`
    TalentGrid SingleComponentResponseOfDestinyItemTalentGridComponent `json:"talentgrid"`
    Sockets SingleComponentResponseOfDestinyItemSocketsComponent `json:"sockets"`
    ReusablePlugs SingleComponentResponseOfDestinyItemReusablePlugsComponent `json:"reusableplugs"`
    PlugObjectives SingleComponentResponseOfDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
}
type SingleComponentResponseOfDestinyItemComponent struct {
    Data DestinyItemComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemInstanceComponent struct {
    Data DestinyItemInstanceComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemObjectivesComponent struct {
    Data DestinyItemObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemPerksComponent struct {
    Data DestinyItemPerksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemRenderComponent struct {
    Data DestinyItemRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemStatsComponent struct {
    Data DestinyItemStatsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemTalentGridComponent struct {
    Data DestinyItemTalentGridComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemSocketsComponent struct {
    Data DestinyItemSocketsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemReusablePlugsComponent struct {
    Data DestinyItemReusablePlugsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyItemPlugObjectivesComponent struct {
    Data DestinyItemPlugObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyVendorFilter struct {
    DestinyVendorFilter json.Number `json:"destinyvendorfilter"`
}
type DestinyVendorsResponse struct {
    VendorGroups SingleComponentResponseOfDestinyVendorGroupComponent `json:"vendorgroups"`
    Vendors DictionaryComponentResponseOfuint32AndDestinyVendorComponent `json:"vendors"`
    Categories DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent `json:"categories"`
    Sales DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent `json:"sales"`
    ItemComponents DestinyItemComponentSetOfint32 `json:"itemcomponents"`
    CurrencyLookups SingleComponentResponseOfDestinyCurrenciesComponent `json:"currencylookups"`
    StringVariables SingleComponentResponseOfDestinyStringVariablesComponent `json:"stringvariables"`
}
type DestinyVendorGroupComponent struct {
    Groups []DestinyVendorGroup `json:"groups"`
}
type DestinyVendorGroup struct {
    VendorGroupHash json.Number `json:"vendorgrouphash"`
    VendorHashes []int `json:"vendorhashes"`
}
type SingleComponentResponseOfDestinyVendorGroupComponent struct {
    Data DestinyVendorGroupComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyVendorBaseComponent struct {
    VendorHash json.Number `json:"vendorhash"`
    NextRefreshDate string `json:"nextrefreshdate"`
    Enabled bool `json:"enabled"`
}
type DestinyVendorComponent struct {
    CanPurchase bool `json:"canpurchase"`
    Progression DestinyProgression `json:"progression"`
    VendorLocationIndex json.Number `json:"vendorlocationindex"`
    SeasonalRank json.Number `json:"seasonalrank"`
    VendorHash json.Number `json:"vendorhash"`
    NextRefreshDate string `json:"nextrefreshdate"`
    Enabled bool `json:"enabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyVendorComponent struct {
    Data DestinyVendorComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyVendorCategoriesComponent struct {
    Categories []DestinyVendorCategory `json:"categories"`
}
type DestinyVendorCategory struct {
    DisplayCategoryIndex json.Number `json:"displaycategoryindex"`
    ItemIndexes []int `json:"itemindexes"`
}
type DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent struct {
    Data DestinyVendorCategoriesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyVendorSaleItemBaseComponent struct {
    VendorItemIndex json.Number `json:"vendoritemindex"`
    ItemHash json.Number `json:"itemhash"`
    OverrideStyleItemHash json.Number `json:"overridestyleitemhash"`
    Quantity json.Number `json:"quantity"`
    Costs []DestinyItemQuantity `json:"costs"`
    OverrideNextRefreshDate string `json:"overridenextrefreshdate"`
    ApiPurchasable bool `json:"apipurchasable"`
}
type DestinyVendorSaleItemComponent struct {
    SaleStatus json.Number `json:"salestatus"`
    RequiredUnlocks []int `json:"requiredunlocks"`
    UnlockStatuses []DestinyUnlockStatus `json:"unlockstatuses"`
    FailureIndexes []int `json:"failureindexes"`
    Augments json.Number `json:"augments"`
    ItemValueVisibility []bool `json:"itemvaluevisibility"`
    VendorItemIndex json.Number `json:"vendoritemindex"`
    ItemHash json.Number `json:"itemhash"`
    OverrideStyleItemHash json.Number `json:"overridestyleitemhash"`
    Quantity json.Number `json:"quantity"`
    Costs []DestinyItemQuantity `json:"costs"`
    OverrideNextRefreshDate string `json:"overridenextrefreshdate"`
    ApiPurchasable bool `json:"apipurchasable"`
}
type VendorItemStatus struct {
    VendorItemStatus json.Number `json:"vendoritemstatus"`
}
type DestinyUnlockStatus struct {
    UnlockHash json.Number `json:"unlockhash"`
    IsSet bool `json:"isset"`
}
type DestinyVendorItemState struct {
    DestinyVendorItemState json.Number `json:"destinyvendoritemstate"`
}
type DestinyVendorSaleItemSetComponentOfDestinyVendorSaleItemComponent struct {
    SaleItems DestinyVendorSaleItemComponent `json:"saleitems"`
}
type PersonalDestinyVendorSaleItemSetComponent struct {
    SaleItems DestinyVendorSaleItemComponent `json:"saleitems"`
}
type DictionaryComponentResponseOfuint32AndPersonalDestinyVendorSaleItemSetComponent struct {
    Data PersonalDestinyVendorSaleItemSetComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyBaseItemComponentSetOfint32 struct {
    Objectives DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfint32AndDestinyItemPerksComponent `json:"perks"`
}
type DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent struct {
    Data DestinyItemObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemPerksComponent struct {
    Data DestinyItemPerksComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyItemComponentSetOfint32 struct {
    Instances DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent `json:"instances"`
    RenderData DictionaryComponentResponseOfint32AndDestinyItemRenderComponent `json:"renderdata"`
    Stats DictionaryComponentResponseOfint32AndDestinyItemStatsComponent `json:"stats"`
    Sockets DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent `json:"sockets"`
    ReusablePlugs DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent `json:"reusableplugs"`
    PlugObjectives DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
    TalentGrids DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent `json:"talentgrids"`
    PlugStates DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent `json:"plugstates"`
    Objectives DictionaryComponentResponseOfint32AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfint32AndDestinyItemPerksComponent `json:"perks"`
}
type DictionaryComponentResponseOfint32AndDestinyItemInstanceComponent struct {
    Data DestinyItemInstanceComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemRenderComponent struct {
    Data DestinyItemRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemStatsComponent struct {
    Data DestinyItemStatsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent struct {
    Data DestinyItemSocketsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemReusablePlugsComponent struct {
    Data DestinyItemReusablePlugsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemPlugObjectivesComponent struct {
    Data DestinyItemPlugObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyItemTalentGridComponent struct {
    Data DestinyItemTalentGridComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyVendorResponse struct {
    Vendor SingleComponentResponseOfDestinyVendorComponent `json:"vendor"`
    Categories SingleComponentResponseOfDestinyVendorCategoriesComponent `json:"categories"`
    Sales DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent `json:"sales"`
    ItemComponents DestinyItemComponentSetOfint32 `json:"itemcomponents"`
    CurrencyLookups SingleComponentResponseOfDestinyCurrenciesComponent `json:"currencylookups"`
    StringVariables SingleComponentResponseOfDestinyStringVariablesComponent `json:"stringvariables"`
}
type SingleComponentResponseOfDestinyVendorComponent struct {
    Data DestinyVendorComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type SingleComponentResponseOfDestinyVendorCategoriesComponent struct {
    Data DestinyVendorCategoriesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfint32AndDestinyVendorSaleItemComponent struct {
    Data DestinyVendorSaleItemComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyPublicVendorsResponse struct {
    VendorGroups SingleComponentResponseOfDestinyVendorGroupComponent `json:"vendorgroups"`
    Vendors DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent `json:"vendors"`
    Categories DictionaryComponentResponseOfuint32AndDestinyVendorCategoriesComponent `json:"categories"`
    Sales DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent `json:"sales"`
    StringVariables SingleComponentResponseOfDestinyStringVariablesComponent `json:"stringvariables"`
}
type DestinyPublicVendorComponent struct {
    VendorHash json.Number `json:"vendorhash"`
    NextRefreshDate string `json:"nextrefreshdate"`
    Enabled bool `json:"enabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyPublicVendorComponent struct {
    Data DestinyPublicVendorComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyPublicVendorSaleItemComponent struct {
    VendorItemIndex json.Number `json:"vendoritemindex"`
    ItemHash json.Number `json:"itemhash"`
    OverrideStyleItemHash json.Number `json:"overridestyleitemhash"`
    Quantity json.Number `json:"quantity"`
    Costs []DestinyItemQuantity `json:"costs"`
    OverrideNextRefreshDate string `json:"overridenextrefreshdate"`
    ApiPurchasable bool `json:"apipurchasable"`
}
type DestinyVendorSaleItemSetComponentOfDestinyPublicVendorSaleItemComponent struct {
    SaleItems DestinyPublicVendorSaleItemComponent `json:"saleitems"`
}
type PublicDestinyVendorSaleItemSetComponent struct {
    SaleItems DestinyPublicVendorSaleItemComponent `json:"saleitems"`
}
type DictionaryComponentResponseOfuint32AndPublicDestinyVendorSaleItemSetComponent struct {
    Data PublicDestinyVendorSaleItemSetComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyCollectibleNodeDetailResponse struct {
    Collectibles SingleComponentResponseOfDestinyCollectiblesComponent `json:"collectibles"`
    CollectibleItemComponents DestinyItemComponentSetOfuint32 `json:"collectibleitemcomponents"`
}
type DestinyItemComponentSetOfuint32 struct {
    Instances DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent `json:"instances"`
    RenderData DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent `json:"renderdata"`
    Stats DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent `json:"stats"`
    Sockets DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent `json:"sockets"`
    ReusablePlugs DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent `json:"reusableplugs"`
    PlugObjectives DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent `json:"plugobjectives"`
    TalentGrids DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent `json:"talentgrids"`
    PlugStates DictionaryComponentResponseOfuint32AndDestinyItemPlugComponent `json:"plugstates"`
    Objectives DictionaryComponentResponseOfuint32AndDestinyItemObjectivesComponent `json:"objectives"`
    Perks DictionaryComponentResponseOfuint32AndDestinyItemPerksComponent `json:"perks"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemInstanceComponent struct {
    Data DestinyItemInstanceComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemRenderComponent struct {
    Data DestinyItemRenderComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemStatsComponent struct {
    Data DestinyItemStatsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemSocketsComponent struct {
    Data DestinyItemSocketsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemReusablePlugsComponent struct {
    Data DestinyItemReusablePlugsComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemPlugObjectivesComponent struct {
    Data DestinyItemPlugObjectivesComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DictionaryComponentResponseOfuint32AndDestinyItemTalentGridComponent struct {
    Data DestinyItemTalentGridComponent `json:"data"`
    Privacy json.Number `json:"privacy"`
    Disabled bool `json:"disabled"`
}
type DestinyActionRequest struct {
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyCharacterActionRequest struct {
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyItemActionRequest struct {
    ItemId json.Number `json:"itemid"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyItemTransferRequest struct {
    ItemReferenceHash json.Number `json:"itemreferencehash"`
    StackSize json.Number `json:"stacksize"`
    TransferToVault bool `json:"transfertovault"`
    ItemId json.Number `json:"itemid"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyPostmasterTransferRequest struct {
    ItemReferenceHash json.Number `json:"itemreferencehash"`
    StackSize json.Number `json:"stacksize"`
    ItemId json.Number `json:"itemid"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyEquipItemResults struct {
    EquipResults []DestinyEquipItemResult `json:"equipresults"`
}
type DestinyEquipItemResult struct {
    ItemInstanceId json.Number `json:"iteminstanceid"`
    EquipStatus json.Number `json:"equipstatus"`
}
type DestinyItemSetActionRequest struct {
    ItemIds []int `json:"itemids"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyItemStateRequest struct {
    State bool `json:"state"`
    ItemId json.Number `json:"itemid"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type InventoryChangedResponse struct {
    AddedInventoryItems []DestinyItemComponent `json:"addedinventoryitems"`
    RemovedInventoryItems []DestinyItemComponent `json:"removedinventoryitems"`
}
type DestinyItemChangeResponse struct {
    Item DestinyItemResponse `json:"item"`
    AddedInventoryItems []DestinyItemComponent `json:"addedinventoryitems"`
    RemovedInventoryItems []DestinyItemComponent `json:"removedinventoryitems"`
}
type DestinyInsertPlugsActionRequest struct {
    ActionToken string `json:"actiontoken"`
    ItemInstanceId json.Number `json:"iteminstanceid"`
    Plug DestinyInsertPlugsRequestEntry `json:"plug"`
    CharacterId json.Number `json:"characterid"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyInsertPlugsRequestEntry struct {
    SocketIndex json.Number `json:"socketindex"`
    SocketArrayType json.Number `json:"socketarraytype"`
    PlugItemHash json.Number `json:"plugitemhash"`
}
type DestinySocketArrayType struct {
    DestinySocketArrayType json.Number `json:"destinysocketarraytype"`
}
type DestinyPostGameCarnageReportData struct {
    Period string `json:"period"`
    StartingPhaseIndex json.Number `json:"startingphaseindex"`
    ActivityDetails DestinyHistoricalStatsActivity `json:"activitydetails"`
    Entries []DestinyPostGameCarnageReportEntry `json:"entries"`
    Teams []DestinyPostGameCarnageReportTeamEntry `json:"teams"`
}
type DestinyHistoricalStatsActivity struct {
    ReferenceId json.Number `json:"referenceid"`
    DirectorActivityHash json.Number `json:"directoractivityhash"`
    InstanceId json.Number `json:"instanceid"`
    Mode json.Number `json:"mode"`
    Modes []int `json:"modes"`
    IsPrivate bool `json:"isprivate"`
    MembershipType json.Number `json:"membershiptype"`
}
type DestinyPostGameCarnageReportEntry struct {
    Standing json.Number `json:"standing"`
    Score DestinyHistoricalStatsValue `json:"score"`
    Player DestinyPlayer `json:"player"`
    CharacterId json.Number `json:"characterid"`
    Values DestinyHistoricalStatsValue `json:"values"`
    Extended DestinyPostGameCarnageReportExtendedData `json:"extended"`
}
type DestinyHistoricalStatsValue struct {
    StatId string `json:"statid"`
    Basic DestinyHistoricalStatsValuePair `json:"basic"`
    Pga DestinyHistoricalStatsValuePair `json:"pga"`
    Weighted DestinyHistoricalStatsValuePair `json:"weighted"`
    ActivityId json.Number `json:"activityid"`
}
type DestinyHistoricalStatsValuePair struct {
    Value float64 `json:"value"`
    DisplayValue string `json:"displayvalue"`
}
type DestinyPlayer struct {
    DestinyUserInfo UserInfoCard `json:"destinyuserinfo"`
    CharacterClass string `json:"characterclass"`
    ClassHash json.Number `json:"classhash"`
    RaceHash json.Number `json:"racehash"`
    GenderHash json.Number `json:"genderhash"`
    CharacterLevel json.Number `json:"characterlevel"`
    LightLevel json.Number `json:"lightlevel"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    ClanName string `json:"clanname"`
    ClanTag string `json:"clantag"`
    EmblemHash json.Number `json:"emblemhash"`
}
type DestinyPostGameCarnageReportExtendedData struct {
    Weapons []DestinyHistoricalWeaponStats `json:"weapons"`
    Values DestinyHistoricalStatsValue `json:"values"`
}
type DestinyHistoricalWeaponStats struct {
    ReferenceId json.Number `json:"referenceid"`
    Values DestinyHistoricalStatsValue `json:"values"`
}
type DestinyPostGameCarnageReportTeamEntry struct {
    TeamId json.Number `json:"teamid"`
    Standing DestinyHistoricalStatsValue `json:"standing"`
    Score DestinyHistoricalStatsValue `json:"score"`
    TeamName string `json:"teamname"`
}
type DestinyReportOffensePgcrRequest struct {
    ReasonCategoryHashes []int `json:"reasoncategoryhashes"`
    ReasonHashes []int `json:"reasonhashes"`
    OffendingCharacterId json.Number `json:"offendingcharacterid"`
}
type DestinyReportReasonCategoryDefinition struct {
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Reasons DestinyReportReasonDefinition `json:"reasons"`
    Hash json.Number `json:"hash"`
    Index json.Number `json:"index"`
    Redacted bool `json:"redacted"`
}
type DestinyReportReasonDefinition struct {
    ReasonHash json.Number `json:"reasonhash"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
}
type DestinyHistoricalStatsDefinition struct {
    StatId string `json:"statid"`
    Group json.Number `json:"group"`
    PeriodTypes []int `json:"periodtypes"`
    Modes []int `json:"modes"`
    Category json.Number `json:"category"`
    StatName string `json:"statname"`
    StatNameAbbr string `json:"statnameabbr"`
    StatDescription string `json:"statdescription"`
    UnitType json.Number `json:"unittype"`
    IconImage string `json:"iconimage"`
    MergeMethod json.Number `json:"mergemethod"`
    UnitLabel string `json:"unitlabel"`
    Weight json.Number `json:"weight"`
    MedalTierHash json.Number `json:"medaltierhash"`
}
type DestinyStatsGroupType struct {
    DestinyStatsGroupType json.Number `json:"destinystatsgrouptype"`
}
type PeriodTypeList struct {
    PeriodTypeList json.Number `json:"periodtypelist"`
}
type DestinyActivityModeTypeList struct {
    DestinyActivityModeTypeList json.Number `json:"destinyactivitymodetypelist"`
}
type DestinyStatsCategoryType struct {
    DestinyStatsCategoryType json.Number `json:"destinystatscategorytype"`
}
type UnitType struct {
    UnitType json.Number `json:"unittype"`
}
type DestinyStatsMergeMethod struct {
    DestinyStatsMergeMethod json.Number `json:"destinystatsmergemethod"`
}
type DestinyLeaderboard struct {
    StatId string `json:"statid"`
    Entries []DestinyLeaderboardEntry `json:"entries"`
}
type DestinyLeaderboardEntry struct {
    Rank json.Number `json:"rank"`
    Player DestinyPlayer `json:"player"`
    CharacterId json.Number `json:"characterid"`
    Value DestinyHistoricalStatsValue `json:"value"`
}
type DestinyLeaderboardResults struct {
    FocusMembershipId json.Number `json:"focusmembershipid"`
    FocusCharacterId json.Number `json:"focuscharacterid"`
}
type DestinyClanAggregateStat struct {
    Mode json.Number `json:"mode"`
    StatId string `json:"statid"`
    Value DestinyHistoricalStatsValue `json:"value"`
}
type DestinyEntitySearchResult struct {
    SuggestedWords []string `json:"suggestedwords"`
    Results SearchResultOfDestinyEntitySearchResultItem `json:"results"`
}
type DestinyEntitySearchResultItem struct {
    Hash json.Number `json:"hash"`
    EntityType string `json:"entitytype"`
    DisplayProperties DestinyDisplayPropertiesDefinition `json:"displayproperties"`
    Weight float64 `json:"weight"`
}
type SearchResultOfDestinyEntitySearchResultItem struct {
    Results []DestinyEntitySearchResultItem `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type PeriodType struct {
    PeriodType json.Number `json:"periodtype"`
}
type DestinyHistoricalStatsByPeriod struct {
    AllTime DestinyHistoricalStatsValue `json:"alltime"`
    AllTimeTier1 DestinyHistoricalStatsValue `json:"alltimetier1"`
    AllTimeTier2 DestinyHistoricalStatsValue `json:"alltimetier2"`
    AllTimeTier3 DestinyHistoricalStatsValue `json:"alltimetier3"`
    Daily []DestinyHistoricalStatsPeriodGroup `json:"daily"`
    Monthly []DestinyHistoricalStatsPeriodGroup `json:"monthly"`
}
type DestinyHistoricalStatsPeriodGroup struct {
    Period string `json:"period"`
    ActivityDetails DestinyHistoricalStatsActivity `json:"activitydetails"`
    Values DestinyHistoricalStatsValue `json:"values"`
}
type DestinyHistoricalStatsResults struct {
    DestinyHistoricalStatsResults interface{} `json:"destinyhistoricalstatsresults"`
}
type DestinyHistoricalStatsAccountResult struct {
    MergedDeletedCharacters DestinyHistoricalStatsWithMerged `json:"mergeddeletedcharacters"`
    MergedAllCharacters DestinyHistoricalStatsWithMerged `json:"mergedallcharacters"`
    Characters []DestinyHistoricalStatsPerCharacter `json:"characters"`
}
type DestinyHistoricalStatsWithMerged struct {
    Results DestinyHistoricalStatsByPeriod `json:"results"`
    Merged DestinyHistoricalStatsByPeriod `json:"merged"`
}
type DestinyHistoricalStatsPerCharacter struct {
    CharacterId json.Number `json:"characterid"`
    Deleted bool `json:"deleted"`
    Results DestinyHistoricalStatsByPeriod `json:"results"`
    Merged DestinyHistoricalStatsByPeriod `json:"merged"`
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
    ActivityHash json.Number `json:"activityhash"`
    Values DestinyHistoricalStatsValue `json:"values"`
}
type DestinyMilestoneContent struct {
    About string `json:"about"`
    Status string `json:"status"`
    Tips []string `json:"tips"`
    ItemCategories []DestinyMilestoneContentItemCategory `json:"itemcategories"`
}
type DestinyMilestoneContentItemCategory struct {
    Title string `json:"title"`
    ItemHashes []int `json:"itemhashes"`
}
type DestinyPublicMilestone struct {
    MilestoneHash json.Number `json:"milestonehash"`
    AvailableQuests []DestinyPublicMilestoneQuest `json:"availablequests"`
    Activities []DestinyPublicMilestoneChallengeActivity `json:"activities"`
    VendorHashes []int `json:"vendorhashes"`
    Vendors []DestinyPublicMilestoneVendor `json:"vendors"`
    StartDate string `json:"startdate"`
    EndDate string `json:"enddate"`
    Order json.Number `json:"order"`
}
type DestinyPublicMilestoneQuest struct {
    QuestItemHash json.Number `json:"questitemhash"`
    Activity DestinyPublicMilestoneActivity `json:"activity"`
    Challenges []DestinyPublicMilestoneChallenge `json:"challenges"`
}
type DestinyPublicMilestoneActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    ModifierHashes []int `json:"modifierhashes"`
    Variants []DestinyPublicMilestoneActivityVariant `json:"variants"`
    ActivityModeHash json.Number `json:"activitymodehash"`
    ActivityModeType json.Number `json:"activitymodetype"`
}
type DestinyPublicMilestoneActivityVariant struct {
    ActivityHash json.Number `json:"activityhash"`
    ActivityModeHash json.Number `json:"activitymodehash"`
    ActivityModeType json.Number `json:"activitymodetype"`
}
type DestinyPublicMilestoneChallenge struct {
    ObjectiveHash json.Number `json:"objectivehash"`
    ActivityHash json.Number `json:"activityhash"`
}
type DestinyPublicMilestoneChallengeActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    ChallengeObjectiveHashes []int `json:"challengeobjectivehashes"`
    ModifierHashes []int `json:"modifierhashes"`
    LoadoutRequirementIndex json.Number `json:"loadoutrequirementindex"`
    PhaseHashes []int `json:"phasehashes"`
    BooleanActivityOptions bool `json:"booleanactivityoptions"`
}
type DestinyPublicMilestoneVendor struct {
    VendorHash json.Number `json:"vendorhash"`
    PreviewItemHash json.Number `json:"previewitemhash"`
}
type AwaInitializeResponse struct {
    CorrelationId string `json:"correlationid"`
    SentToSelf bool `json:"senttoself"`
}
type AwaPermissionRequested struct {
    Type json.Number `json:"type"`
    AffectedItemId json.Number `json:"affecteditemid"`
    MembershipType json.Number `json:"membershiptype"`
    CharacterId json.Number `json:"characterid"`
}
type AwaType struct {
    AwaType json.Number `json:"awatype"`
}
type AwaUserResponse struct {
    Selection json.Number `json:"selection"`
    CorrelationId string `json:"correlationid"`
    Nonce []string `json:"nonce"`
}
type AwaUserSelection struct {
    AwaUserSelection json.Number `json:"awauserselection"`
}
type AwaAuthorizationResult struct {
    UserSelection json.Number `json:"userselection"`
    ResponseReason json.Number `json:"responsereason"`
    DeveloperNote string `json:"developernote"`
    ActionToken string `json:"actiontoken"`
    MaximumNumberOfUses json.Number `json:"maximumnumberofuses"`
    ValidUntil string `json:"validuntil"`
    Type json.Number `json:"type"`
    MembershipType json.Number `json:"membershiptype"`
}
type AwaResponseReason struct {
    AwaResponseReason json.Number `json:"awaresponsereason"`
}
type CommunityContentSortMode struct {
    CommunityContentSortMode json.Number `json:"communitycontentsortmode"`
}
type TrendingCategories struct {
    Categories []TrendingCategory `json:"categories"`
}
type TrendingCategory struct {
    CategoryName string `json:"categoryname"`
    Entries SearchResultOfTrendingEntry `json:"entries"`
    CategoryId string `json:"categoryid"`
}
type TrendingEntry struct {
    Weight float64 `json:"weight"`
    IsFeatured bool `json:"isfeatured"`
    Identifier string `json:"identifier"`
    EntityType json.Number `json:"entitytype"`
    DisplayName string `json:"displayname"`
    Tagline string `json:"tagline"`
    Image string `json:"image"`
    StartDate string `json:"startdate"`
    EndDate string `json:"enddate"`
    Link string `json:"link"`
    WebmVideo string `json:"webmvideo"`
    Mp4Video string `json:"mp4video"`
    FeatureImage string `json:"featureimage"`
    Items []TrendingEntry `json:"items"`
    CreationDate string `json:"creationdate"`
}
type TrendingEntryType struct {
    TrendingEntryType json.Number `json:"trendingentrytype"`
}
type SearchResultOfTrendingEntry struct {
    Results []TrendingEntry `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type TrendingDetail struct {
    Identifier string `json:"identifier"`
    EntityType json.Number `json:"entitytype"`
    News TrendingEntryNews `json:"news"`
    Support TrendingEntrySupportArticle `json:"support"`
    DestinyItem TrendingEntryDestinyItem `json:"destinyitem"`
    DestinyActivity TrendingEntryDestinyActivity `json:"destinyactivity"`
    DestinyRitual TrendingEntryDestinyRitual `json:"destinyritual"`
    Creation TrendingEntryCommunityCreation `json:"creation"`
}
type TrendingEntryNews struct {
    Article ContentItemPublicContract `json:"article"`
}
type TrendingEntrySupportArticle struct {
    Article ContentItemPublicContract `json:"article"`
}
type TrendingEntryDestinyItem struct {
    ItemHash json.Number `json:"itemhash"`
}
type TrendingEntryDestinyActivity struct {
    ActivityHash json.Number `json:"activityhash"`
    Status DestinyPublicActivityStatus `json:"status"`
}
type DestinyPublicActivityStatus struct {
    ChallengeObjectiveHashes []int `json:"challengeobjectivehashes"`
    ModifierHashes []int `json:"modifierhashes"`
    RewardTooltipItems []DestinyItemQuantity `json:"rewardtooltipitems"`
}
type TrendingEntryDestinyRitual struct {
    Image string `json:"image"`
    Icon string `json:"icon"`
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    DateStart string `json:"datestart"`
    DateEnd string `json:"dateend"`
    MilestoneDetails DestinyPublicMilestone `json:"milestonedetails"`
    EventContent DestinyMilestoneContent `json:"eventcontent"`
}
type TrendingEntryCommunityCreation struct {
    Media string `json:"media"`
    Title string `json:"title"`
    Author string `json:"author"`
    AuthorMembershipId json.Number `json:"authormembershipid"`
    PostId json.Number `json:"postid"`
    Body string `json:"body"`
    Upvotes json.Number `json:"upvotes"`
}
type FireteamDateRange struct {
    FireteamDateRange json.Number `json:"fireteamdaterange"`
}
type FireteamPlatform struct {
    FireteamPlatform json.Number `json:"fireteamplatform"`
}
type FireteamPublicSearchOption struct {
    FireteamPublicSearchOption json.Number `json:"fireteampublicsearchoption"`
}
type FireteamSlotSearch struct {
    FireteamSlotSearch json.Number `json:"fireteamslotsearch"`
}
type FireteamSummary struct {
    FireteamId json.Number `json:"fireteamid"`
    GroupId json.Number `json:"groupid"`
    Platform json.Number `json:"platform"`
    ActivityType json.Number `json:"activitytype"`
    IsImmediate bool `json:"isimmediate"`
    ScheduledTime string `json:"scheduledtime"`
    OwnerMembershipId json.Number `json:"ownermembershipid"`
    PlayerSlotCount json.Number `json:"playerslotcount"`
    AlternateSlotCount json.Number `json:"alternateslotcount"`
    AvailablePlayerSlotCount json.Number `json:"availableplayerslotcount"`
    AvailableAlternateSlotCount json.Number `json:"availablealternateslotcount"`
    Title string `json:"title"`
    DateCreated string `json:"datecreated"`
    DateModified string `json:"datemodified"`
    IsPublic bool `json:"ispublic"`
    Locale string `json:"locale"`
    IsValid bool `json:"isvalid"`
    DatePlayerModified string `json:"dateplayermodified"`
    TitleBeforeModeration string `json:"titlebeforemoderation"`
}
type SearchResultOfFireteamSummary struct {
    Results []FireteamSummary `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type FireteamResponse struct {
    Summary FireteamSummary `json:"summary"`
    Members []FireteamMember `json:"members"`
    Alternates []FireteamMember `json:"alternates"`
}
type FireteamMember struct {
    DestinyUserInfo FireteamUserInfoCard `json:"destinyuserinfo"`
    BungieNetUserInfo UserInfoCard `json:"bungienetuserinfo"`
    CharacterId json.Number `json:"characterid"`
    DateJoined string `json:"datejoined"`
    HasMicrophone bool `json:"hasmicrophone"`
    LastPlatformInviteAttemptDate string `json:"lastplatforminviteattemptdate"`
    LastPlatformInviteAttemptResult json.Number `json:"lastplatforminviteattemptresult"`
}
type FireteamUserInfoCard struct {
    FireteamDisplayName string `json:"fireteamdisplayname"`
    FireteamMembershipType json.Number `json:"fireteammembershiptype"`
    SupplementalDisplayName string `json:"supplementaldisplayname"`
    IconPath string `json:"iconpath"`
    CrossSaveOverride json.Number `json:"crosssaveoverride"`
    ApplicableMembershipTypes []int `json:"applicablemembershiptypes"`
    IsPublic bool `json:"ispublic"`
    MembershipType json.Number `json:"membershiptype"`
    MembershipId json.Number `json:"membershipid"`
    DisplayName string `json:"displayname"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type FireteamPlatformInviteResult struct {
    FireteamPlatformInviteResult json.Number `json:"fireteamplatforminviteresult"`
}
type SearchResultOfFireteamResponse struct {
    Results []FireteamResponse `json:"results"`
    TotalResults json.Number `json:"totalresults"`
    HasMore bool `json:"hasmore"`
    Query PagedQuery `json:"query"`
    ReplacementContinuationToken string `json:"replacementcontinuationtoken"`
    UseTotalResults bool `json:"usetotalresults"`
}
type BungieFriendListResponse struct {
    Friends []BungieFriend `json:"friends"`
}
type BungieFriend struct {
    LastSeenAsMembershipId json.Number `json:"lastseenasmembershipid"`
    LastSeenAsBungieMembershipType json.Number `json:"lastseenasbungiemembershiptype"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
    OnlineStatus json.Number `json:"onlinestatus"`
    OnlineTitle json.Number `json:"onlinetitle"`
    Relationship json.Number `json:"relationship"`
    BungieNetUser GeneralUser `json:"bungienetuser"`
}
type PresenceStatus struct {
    PresenceStatus json.Number `json:"presencestatus"`
}
type PresenceOnlineStateFlags struct {
    PresenceOnlineStateFlags json.Number `json:"presenceonlinestateflags"`
}
type FriendRelationshipState struct {
    FriendRelationshipState json.Number `json:"friendrelationshipstate"`
}
type BungieFriendRequestListResponse struct {
    IncomingRequests []BungieFriend `json:"incomingrequests"`
    OutgoingRequests []BungieFriend `json:"outgoingrequests"`
}
type PlatformFriendType struct {
    PlatformFriendType json.Number `json:"platformfriendtype"`
}
type PlatformFriendResponse struct {
    ItemsPerPage json.Number `json:"itemsperpage"`
    CurrentPage json.Number `json:"currentpage"`
    HasMore bool `json:"hasmore"`
    PlatformFriends []PlatformFriend `json:"platformfriends"`
}
type PlatformFriend struct {
    PlatformDisplayName string `json:"platformdisplayname"`
    FriendPlatform json.Number `json:"friendplatform"`
    DestinyMembershipId json.Number `json:"destinymembershipid"`
    DestinyMembershipType json.Number `json:"destinymembershiptype"`
    BungieNetMembershipId json.Number `json:"bungienetmembershipid"`
    BungieGlobalDisplayName string `json:"bungieglobaldisplayname"`
    BungieGlobalDisplayNameCode json.Number `json:"bungieglobaldisplaynamecode"`
}
type CoreSettingsConfiguration struct {
    Environment string `json:"environment"`
    Systems CoreSystem `json:"systems"`
    IgnoreReasons []CoreSetting `json:"ignorereasons"`
    ForumCategories []CoreSetting `json:"forumcategories"`
    GroupAvatars []CoreSetting `json:"groupavatars"`
    DestinyMembershipTypes []CoreSetting `json:"destinymembershiptypes"`
    RecruitmentPlatformTags []CoreSetting `json:"recruitmentplatformtags"`
    RecruitmentMiscTags []CoreSetting `json:"recruitmentmisctags"`
    RecruitmentActivities []CoreSetting `json:"recruitmentactivities"`
    UserContentLocales []CoreSetting `json:"usercontentlocales"`
    SystemContentLocales []CoreSetting `json:"systemcontentlocales"`
    ClanBannerDecals []CoreSetting `json:"clanbannerdecals"`
    ClanBannerDecalColors []CoreSetting `json:"clanbannerdecalcolors"`
    ClanBannerGonfalons []CoreSetting `json:"clanbannergonfalons"`
    ClanBannerGonfalonColors []CoreSetting `json:"clanbannergonfaloncolors"`
    ClanBannerGonfalonDetails []CoreSetting `json:"clanbannergonfalondetails"`
    ClanBannerGonfalonDetailColors []CoreSetting `json:"clanbannergonfalondetailcolors"`
    ClanBannerStandards []CoreSetting `json:"clanbannerstandards"`
    Destiny2CoreSettings Destiny2CoreSettings `json:"destiny2coresettings"`
    EmailSettings EmailSettings `json:"emailsettings"`
    FireteamActivities []CoreSetting `json:"fireteamactivities"`
}
type CoreSystem struct {
    Enabled bool `json:"enabled"`
    Parameters string `json:"parameters"`
}
type CoreSetting struct {
    Identifier string `json:"identifier"`
    IsDefault bool `json:"isdefault"`
    DisplayName string `json:"displayname"`
    Summary string `json:"summary"`
    ImagePath string `json:"imagepath"`
    ChildSettings []CoreSetting `json:"childsettings"`
}
type Destiny2CoreSettings struct {
    CollectionRootNode json.Number `json:"collectionrootnode"`
    BadgesRootNode json.Number `json:"badgesrootnode"`
    RecordsRootNode json.Number `json:"recordsrootnode"`
    MedalsRootNode json.Number `json:"medalsrootnode"`
    MetricsRootNode json.Number `json:"metricsrootnode"`
    ActiveTriumphsRootNodeHash json.Number `json:"activetriumphsrootnodehash"`
    ActiveSealsRootNodeHash json.Number `json:"activesealsrootnodehash"`
    LegacyTriumphsRootNodeHash json.Number `json:"legacytriumphsrootnodehash"`
    LegacySealsRootNodeHash json.Number `json:"legacysealsrootnodehash"`
    MedalsRootNodeHash json.Number `json:"medalsrootnodehash"`
    ExoticCatalystsRootNodeHash json.Number `json:"exoticcatalystsrootnodehash"`
    LoreRootNodeHash json.Number `json:"lorerootnodehash"`
    CurrentRankProgressionHashes []int `json:"currentrankprogressionhashes"`
    UndiscoveredCollectibleImage string `json:"undiscoveredcollectibleimage"`
    AmmoTypeHeavyIcon string `json:"ammotypeheavyicon"`
    AmmoTypeSpecialIcon string `json:"ammotypespecialicon"`
    AmmoTypePrimaryIcon string `json:"ammotypeprimaryicon"`
    CurrentSeasonalArtifactHash json.Number `json:"currentseasonalartifacthash"`
    CurrentSeasonHash json.Number `json:"currentseasonhash"`
    SeasonalChallengesPresentationNodeHash json.Number `json:"seasonalchallengespresentationnodehash"`
    FutureSeasonHashes []int `json:"futureseasonhashes"`
    PastSeasonHashes []int `json:"pastseasonhashes"`
}
type EmailSettings struct {
    OptInDefinitions EmailOptInDefinition `json:"optindefinitions"`
    SubscriptionDefinitions EmailSubscriptionDefinition `json:"subscriptiondefinitions"`
    Views EmailViewDefinition `json:"views"`
}
type EmailOptInDefinition struct {
    Name string `json:"name"`
    Value json.Number `json:"value"`
    SetByDefault bool `json:"setbydefault"`
    DependentSubscriptions []EmailSubscriptionDefinition `json:"dependentsubscriptions"`
}
type OptInFlags struct {
    OptInFlags json.Number `json:"optinflags"`
}
type EmailSubscriptionDefinition struct {
    Name string `json:"name"`
    Localization EMailSettingSubscriptionLocalization `json:"localization"`
    Value json.Number `json:"value"`
}
type EMailSettingLocalization struct {
    Title string `json:"title"`
    Description string `json:"description"`
}
type EMailSettingSubscriptionLocalization struct {
    UnknownUserDescription string `json:"unknownuserdescription"`
    RegisteredUserDescription string `json:"registereduserdescription"`
    UnregisteredUserDescription string `json:"unregistereduserdescription"`
    UnknownUserActionText string `json:"unknownuseractiontext"`
    KnownUserActionText string `json:"knownuseractiontext"`
    Title string `json:"title"`
    Description string `json:"description"`
}
type EmailViewDefinition struct {
    Name string `json:"name"`
    ViewSettings []EmailViewDefinitionSetting `json:"viewsettings"`
}
type EmailViewDefinitionSetting struct {
    Name string `json:"name"`
    Localization EMailSettingLocalization `json:"localization"`
    SetByDefault bool `json:"setbydefault"`
    OptInAggregateValue json.Number `json:"optinaggregatevalue"`
    Subscriptions []EmailSubscriptionDefinition `json:"subscriptions"`
}
type GlobalAlert struct {
    AlertKey string `json:"alertkey"`
    AlertHtml string `json:"alerthtml"`
    AlertTimestamp string `json:"alerttimestamp"`
    AlertLink string `json:"alertlink"`
    AlertLevel json.Number `json:"alertlevel"`
    AlertType json.Number `json:"alerttype"`
    StreamInfo StreamInfo `json:"streaminfo"`
}
type GlobalAlertLevel struct {
    GlobalAlertLevel json.Number `json:"globalalertlevel"`
}
type GlobalAlertType struct {
    GlobalAlertType json.Number `json:"globalalerttype"`
}
type StreamInfo struct {
    ChannelName string `json:"channelname"`
}
