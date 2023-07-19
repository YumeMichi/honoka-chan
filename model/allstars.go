package model

// AsLoginRes ...
type AsLoginRes struct {
	SessionKey              string             `json:"session_key"`
	UserModel               UserModel          `json:"user_model"`
	IsPlatformServiceLinked bool               `json:"is_platform_service_linked"`
	LastTimestamp           int64              `json:"last_timestamp"`
	Cautions                []any              `json:"cautions"`
	ShowHomeCaution         bool               `json:"show_home_caution"`
	LiveResume              any                `json:"live_resume"`
	FromEea                 bool               `json:"from_eea"`
	GdprConsentedInfo       GdprConsentedInfo  `json:"gdpr_consented_info"`
	UserID                  int                `json:"user_id"`
	IsUnderAge              int                `json:"is_under_age"`
	AreaID                  int                `json:"area_id"`
	AuthCount               int                `json:"auth_count"`
	MemberLovePanels        []MemberLovePanels `json:"member_love_panels"`
	CheckMaintenance        bool               `json:"check_maintenance"`
	ReproInfo               ReproInfo          `json:"repro_info"`
}

// UserModel ...
type UserModel struct {
	UserStatus                                                                 AsUserStatus `json:"user_status"`
	UserMemberByMemberID                                                       any          `json:"user_member_by_member_id"`
	UserCardByCardID                                                           []any        `json:"user_card_by_card_id"`
	UserSuitBySuitID                                                           []any        `json:"user_suit_by_suit_id"`
	UserLiveDeckByID                                                           []any        `json:"user_live_deck_by_id"`
	UserLivePartyByID                                                          []any        `json:"user_live_party_by_id"`
	UserLessonDeckByID                                                         []any        `json:"user_lesson_deck_by_id"`
	UserLiveMvDeckByID                                                         []any        `json:"user_live_mv_deck_by_id"`
	UserLiveMvDeckCustomByID                                                   any          `json:"user_live_mv_deck_custom_by_id"`
	UserLiveDifficultyByDifficultyID                                           []any        `json:"user_live_difficulty_by_difficulty_id"`
	UserStoryMainByStoryMainID                                                 []any        `json:"user_story_main_by_story_main_id"`
	UserStoryMainSelectedByStoryMainCellID                                     []any        `json:"user_story_main_selected_by_story_main_cell_id"`
	UserVoiceByVoiceID                                                         []any        `json:"user_voice_by_voice_id"`
	UserEmblemByEmblemID                                                       []any        `json:"user_emblem_by_emblem_id"`
	UserGachaTicketByTicketID                                                  []any        `json:"user_gacha_ticket_by_ticket_id"`
	UserGachaPointByPointID                                                    []any        `json:"user_gacha_point_by_point_id"`
	UserLessonEnhancingItemByItemID                                            []any        `json:"user_lesson_enhancing_item_by_item_id"`
	UserTrainingMaterialByItemID                                               []any        `json:"user_training_material_by_item_id"`
	UserGradeUpItemByItemID                                                    []any        `json:"user_grade_up_item_by_item_id"`
	UserCustomBackgroundByID                                                   []any        `json:"user_custom_background_by_id"`
	UserStorySideByID                                                          []any        `json:"user_story_side_by_id"`
	UserStoryMemberByID                                                        []any        `json:"user_story_member_by_id"`
	UserCommunicationMemberDetailBadgeByID                                     []any        `json:"user_communication_member_detail_badge_by_id"`
	UserStoryEventHistoryByID                                                  []any        `json:"user_story_event_history_by_id"`
	UserRecoveryLpByID                                                         []any        `json:"user_recovery_lp_by_id"`
	UserRecoveryApByID                                                         []any        `json:"user_recovery_ap_by_id"`
	UserMissionByMissionID                                                     []any        `json:"user_mission_by_mission_id"`
	UserDailyMissionByMissionID                                                []any        `json:"user_daily_mission_by_mission_id"`
	UserWeeklyMissionByMissionID                                               []any        `json:"user_weekly_mission_by_mission_id"`
	UserInfoTriggerBasicByTriggerID                                            []any        `json:"user_info_trigger_basic_by_trigger_id"`
	UserInfoTriggerCardGradeUpByTriggerID                                      []any        `json:"user_info_trigger_card_grade_up_by_trigger_id"`
	UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID                    []any        `json:"user_info_trigger_member_guild_support_item_expired_by_trigger_id"`
	UserInfoTriggerMemberLoveLevelUpByTriggerID                                []any        `json:"user_info_trigger_member_love_level_up_by_trigger_id"`
	UserAccessoryByUserAccessoryID                                             []any        `json:"user_accessory_by_user_accessory_id"`
	UserAccessoryLevelUpItemByID                                               []any        `json:"user_accessory_level_up_item_by_id"`
	UserAccessoryRarityUpItemByID                                              []any        `json:"user_accessory_rarity_up_item_by_id"`
	UserUnlockScenesByEnum                                                     []any        `json:"user_unlock_scenes_by_enum"`
	UserSceneTipsByEnum                                                        []any        `json:"user_scene_tips_by_enum"`
	UserRuleDescriptionByID                                                    []any        `json:"user_rule_description_by_id"`
	UserExchangeEventPointByID                                                 []any        `json:"user_exchange_event_point_by_id"`
	UserSchoolIdolFestivalIDRewardMissionByID                                  []any        `json:"user_school_idol_festival_id_reward_mission_by_id"`
	UserGpsPresentReceivedByID                                                 []any        `json:"user_gps_present_received_by_id"`
	UserEventMarathonByEventMasterID                                           []any        `json:"user_event_marathon_by_event_master_id"`
	UserEventMiningByEventMasterID                                             []any        `json:"user_event_mining_by_event_master_id"`
	UserEventCoopByEventMasterID                                               []any        `json:"user_event_coop_by_event_master_id"`
	UserLiveSkipTicketByID                                                     []any        `json:"user_live_skip_ticket_by_id"`
	UserStoryEventUnlockItemByID                                               []any        `json:"user_story_event_unlock_item_by_id"`
	UserEventMarathonBoosterByID                                               []any        `json:"user_event_marathon_booster_by_id"`
	UserReferenceBookByID                                                      []any        `json:"user_reference_book_by_id"`
	UserReviewRequestProcessFlowByID                                           []any        `json:"user_review_request_process_flow_by_id"`
	UserRankExpByID                                                            []any        `json:"user_rank_exp_by_id"`
	UserShareByID                                                              []any        `json:"user_share_by_id"`
	UserTowerByTowerID                                                         []any        `json:"user_tower_by_tower_id"`
	UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID []any        `json:"user_recovery_tower_card_used_count_item_by_recovery_tower_card_used_count_item_master_id"`
	UserStoryLinkageByID                                                       []any        `json:"user_story_linkage_by_id"`
	UserSubscriptionStatusByID                                                 []any        `json:"user_subscription_status_by_id"`
	UserStoryMainPartDigestMovieByID                                           []any        `json:"user_story_main_part_digest_movie_by_id"`
	UserMemberGuildByID                                                        []any        `json:"user_member_guild_by_id"`
	UserMemberGuildSupportItemByID                                             []any        `json:"user_member_guild_support_item_by_id"`
	UserDailyTheaterByDailyTheaterID                                           []any        `json:"user_daily_theater_by_daily_theater_id"`
	UserPlayListByID                                                           []any        `json:"user_play_list_by_id"`
}

// Name ...
type Name struct {
	DotUnderText string `json:"dot_under_text"`
}

// Nickname ...
type Nickname struct {
	DotUnderText string `json:"dot_under_text"`
}

// Message ...
type Message struct {
	DotUnderText string `json:"dot_under_text"`
}

// AsUserStatus ...
type AsUserStatus struct {
	Name                                      Name     `json:"name"`
	Nickname                                  Nickname `json:"nickname"`
	LastLoginAt                               int64    `json:"last_login_at"`
	Rank                                      int      `json:"rank"`
	Exp                                       int      `json:"exp"`
	Message                                   Message  `json:"message"`
	RecommendCardMasterID                     int      `json:"recommend_card_master_id"`
	MaxFriendNum                              int      `json:"max_friend_num"`
	LivePointFullAt                           int      `json:"live_point_full_at"`
	LivePointBroken                           int      `json:"live_point_broken"`
	LivePointSubscriptionRecoveryDailyCount   int      `json:"live_point_subscription_recovery_daily_count"`
	LivePointSubscriptionRecoveryDailyResetAt int      `json:"live_point_subscription_recovery_daily_reset_at"`
	ActivityPointCount                        int      `json:"activity_point_count"`
	ActivityPointResetAt                      int      `json:"activity_point_reset_at"`
	ActivityPointPaymentRecoveryDailyCount    int      `json:"activity_point_payment_recovery_daily_count"`
	ActivityPointPaymentRecoveryDailyResetAt  int      `json:"activity_point_payment_recovery_daily_reset_at"`
	GameMoney                                 int      `json:"game_money"`
	CardExp                                   int      `json:"card_exp"`
	FreeSnsCoin                               int      `json:"free_sns_coin"`
	AppleSnsCoin                              int      `json:"apple_sns_coin"`
	GoogleSnsCoin                             int      `json:"google_sns_coin"`
	Cash                                      int      `json:"cash"`
	SubscriptionCoin                          int      `json:"subscription_coin"`
	BirthDate                                 any      `json:"birth_date"`
	BirthMonth                                int      `json:"birth_month"`
	BirthDay                                  int      `json:"birth_day"`
	LatestLiveDeckID                          int      `json:"latest_live_deck_id"`
	MainLessonDeckID                          int      `json:"main_lesson_deck_id"`
	FavoriteMemberID                          int      `json:"favorite_member_id"`
	LastLiveDifficultyID                      int      `json:"last_live_difficulty_id"`
	LpMagnification                           int      `json:"lp_magnification"`
	EmblemID                                  int      `json:"emblem_id"`
	DeviceToken                               string   `json:"device_token"`
	TutorialPhase                             int      `json:"tutorial_phase"`
	TutorialEndAt                             int      `json:"tutorial_end_at"`
	LoginDays                                 int      `json:"login_days"`
	NaviTapCount                              int      `json:"navi_tap_count"`
	NaviTapRecoverAt                          int      `json:"navi_tap_recover_at"`
	IsAutoMode                                bool     `json:"is_auto_mode"`
	MaxScoreLiveDifficultyMasterID            int      `json:"max_score_live_difficulty_master_id"`
	LiveMaxScore                              int      `json:"live_max_score"`
	MaxComboLiveDifficultyMasterID            int      `json:"max_combo_live_difficulty_master_id"`
	LiveMaxCombo                              int      `json:"live_max_combo"`
	LessonResumeStatus                        int      `json:"lesson_resume_status"`
	AccessoryBoxAdditional                    int      `json:"accessory_box_additional"`
	TermsOfUseVersion                         int      `json:"terms_of_use_version"`
	BootstrapSifidCheckAt                     int      `json:"bootstrap_sifid_check_at"`
	GdprVersion                               int      `json:"gdpr_version"`
	MemberGuildMemberMasterID                 int      `json:"member_guild_member_master_id"`
	MemberGuildLastUpdatedAt                  int      `json:"member_guild_last_updated_at"`
}

// GdprConsentedInfo ...
type GdprConsentedInfo struct {
	HasConsentedAdPurposeOfUse bool `json:"has_consented_ad_purpose_of_use"`
	HasConsentedCrashReport    bool `json:"has_consented_crash_report"`
}

// MemberLovePanels ...
type MemberLovePanels struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelCellIds []int `json:"member_love_panel_cell_ids"`
}

// ReproInfo ...
type ReproInfo struct {
	GroupNo int `json:"group_no"`
}

// AsCardInfo ...
type AsCardInfo struct {
	CardMasterID               int   `json:"card_master_id"`
	Level                      int   `json:"level"`
	Exp                        int   `json:"exp"`
	LovePoint                  int   `json:"love_point"`
	IsFavorite                 bool  `json:"is_favorite"`
	IsAwakening                bool  `json:"is_awakening"`
	IsAwakeningImage           bool  `json:"is_awakening_image"`
	IsAllTrainingActivated     bool  `json:"is_all_training_activated"`
	TrainingActivatedCellCount int   `json:"training_activated_cell_count"`
	MaxFreePassiveSkill        int   `json:"max_free_passive_skill"`
	Grade                      int   `json:"grade"`
	TrainingLife               int   `json:"training_life"`
	TrainingAttack             int   `json:"training_attack"`
	TrainingDexterity          int   `json:"training_dexterity"`
	ActiveSkillLevel           int   `json:"active_skill_level"`
	PassiveSkillALevel         int   `json:"passive_skill_a_level"`
	PassiveSkillBLevel         int   `json:"passive_skill_b_level"`
	PassiveSkillCLevel         int   `json:"passive_skill_c_level"`
	AdditionalPassiveSkill1ID  int   `json:"additional_passive_skill_1_id"`
	AdditionalPassiveSkill2ID  int   `json:"additional_passive_skill_2_id"`
	AdditionalPassiveSkill3ID  int   `json:"additional_passive_skill_3_id"`
	AdditionalPassiveSkill4ID  int   `json:"additional_passive_skill_4_id"`
	AcquiredAt                 int64 `json:"acquired_at"`
	IsNew                      bool  `json:"is_new"`
}

// AsSuitInfo ...
type AsSuitInfo struct {
	SuitMasterID int  `json:"suit_master_id"`
	IsNew        bool `json:"is_new"`
}

// AsEmblemInfo ...
type AsEmblemInfo struct {
	EmblemMID   int   `json:"emblem_m_id"`
	IsNew       bool  `json:"is_new"`
	EmblemParam any   `json:"emblem_param"`
	AcquiredAt  int64 `json:"acquired_at"`
}

// AsEmblemId ...
type AsEmblemId struct {
	EmblemMasterID int  `json:"emblem_master_id"`
	IsNew          bool `json:"is_new"`
}

// AsLiveDifficultyInfo ...
type AsLiveDifficultyInfo struct {
	LiveDifficultyID              int  `json:"live_difficulty_id"`
	MaxScore                      int  `json:"max_score"`
	MaxCombo                      int  `json:"max_combo"`
	PlayCount                     int  `json:"play_count"`
	ClearCount                    int  `json:"clear_count"`
	CancelCount                   int  `json:"cancel_count"`
	NotClearedCount               int  `json:"not_cleared_count"`
	IsFullCombo                   bool `json:"is_full_combo"`
	ClearedDifficultyAchievement1 any  `json:"cleared_difficulty_achievement_1"`
	ClearedDifficultyAchievement2 any  `json:"cleared_difficulty_achievement_2"`
	ClearedDifficultyAchievement3 any  `json:"cleared_difficulty_achievement_3"`
	EnableAutoplay                bool `json:"enable_autoplay"`
	IsAutoplay                    bool `json:"is_autoplay"`
	IsNew                         bool `json:"is_new"`
}

// AsMainStoryInfo ...
type AsMainStoryInfo struct {
	StoryMainMasterID int `json:"story_main_master_id"`
}

// AsNaviVoiceInfo ...
type AsNaviVoiceInfo struct {
	NaviVoiceMasterID int  `json:"navi_voice_master_id"`
	IsNew             bool `json:"is_new"`
}

// AsCustomBackgroundInfo ...
type AsCustomBackgroundInfo struct {
	CustomBackgroundMasterID int  `json:"custom_background_master_id"`
	IsNew                    bool `json:"is_new"`
}

// AsStorySideInfo ...
type AsStorySideInfo struct {
	StorySideMasterID int   `json:"story_side_master_id"`
	IsNew             bool  `json:"is_new"`
	AcquiredAt        int64 `json:"acquired_at"`
}

// AsStoryMemberInfo ...
type AsStoryMemberInfo struct {
	StoryMemberMasterID int   `json:"story_member_master_id"`
	IsNew               bool  `json:"is_new"`
	AcquiredAt          int64 `json:"acquired_at"`
}

// AsStoryEventInfo ...
type AsStoryEventInfo struct {
	StoryEventID int `json:"story_event_id"`
}

// AsMissionRes ...
type AsMissionRes struct {
	ID    int `xorm:"id"`
	Count int `xorm:"mission_clear_condition_count"`
}

// AsFreeMissionInfo ...
type AsFreeMissionInfo struct {
	MissionMID       int   `json:"mission_m_id"`
	IsNew            bool  `json:"is_new"`
	MissionCount     int   `json:"mission_count"`
	IsCleared        bool  `json:"is_cleared"`
	IsReceivedReward bool  `json:"is_received_reward"`
	NewExpiredAt     int64 `json:"new_expired_at"`
}

// AsDailyMissionInfo ...
type AsDailyMissionInfo struct {
	MissionMID        int   `json:"mission_m_id"`
	IsNew             bool  `json:"is_new"`
	MissionStartCount int   `json:"mission_start_count"`
	MissionCount      int   `json:"mission_count"`
	IsCleared         bool  `json:"is_cleared"`
	IsReceivedReward  bool  `json:"is_received_reward"`
	ClearedExpiredAt  int64 `json:"cleared_expired_at"`
}

// AsWeeklyMissionInfo ...
type AsWeeklyMissionInfo struct {
	MissionMID        int   `json:"mission_m_id"`
	IsNew             bool  `json:"is_new"`
	MissionStartCount int   `json:"mission_start_count"`
	MissionCount      int   `json:"mission_count"`
	IsCleared         bool  `json:"is_cleared"`
	IsReceivedReward  bool  `json:"is_received_reward"`
	ClearedExpiredAt  int64 `json:"cleared_expired_at"`
	NewExpiredAt      int64 `json:"new_expired_at"`
}

// AsMemberLovePanelInfo ...
type AsMemberLovePanelInfo struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelCellIds []int `json:"member_love_panel_cell_ids"`
}

// AsCardRes ...
type AsCardRes struct {
	ID                  int `xorm:"id"`
	CardRarityType      int `xorm:"card_rarity_type"`
	MaxPassiveSkillSlot int `xorm:"max_passive_skill_slot"`
}

// AsSaveDeckReq ...
type AsSaveDeckReq struct {
	DeckID       int   `json:"deck_id"`
	CardWithSuit []int `json:"card_with_suit"`
	SquadDict    []any `json:"squad_dict"`
}

// AsDeckSquadDict ...
type AsDeckSquadDict struct {
	CardMasterIds    []int `json:"card_master_ids"`
	UserAccessoryIds []any `json:"user_accessory_ids"`
}

// AsDeckInfo ...
type AsDeckInfo struct {
	UserLiveDeckID int        `json:"user_live_deck_id"`
	Name           AsDeckName `json:"name"`
	CardMasterID1  int        `json:"card_master_id_1"`
	CardMasterID2  int        `json:"card_master_id_2"`
	CardMasterID3  int        `json:"card_master_id_3"`
	CardMasterID4  int        `json:"card_master_id_4"`
	CardMasterID5  int        `json:"card_master_id_5"`
	CardMasterID6  int        `json:"card_master_id_6"`
	CardMasterID7  int        `json:"card_master_id_7"`
	CardMasterID8  int        `json:"card_master_id_8"`
	CardMasterID9  int        `json:"card_master_id_9"`
	SuitMasterID1  int        `json:"suit_master_id_1"`
	SuitMasterID2  int        `json:"suit_master_id_2"`
	SuitMasterID3  int        `json:"suit_master_id_3"`
	SuitMasterID4  int        `json:"suit_master_id_4"`
	SuitMasterID5  int        `json:"suit_master_id_5"`
	SuitMasterID6  int        `json:"suit_master_id_6"`
	SuitMasterID7  int        `json:"suit_master_id_7"`
	SuitMasterID8  int        `json:"suit_master_id_8"`
	SuitMasterID9  int        `json:"suit_master_id_9"`
}

// AsDeckName ...
type AsDeckName struct {
	DotUnderText string `json:"dot_under_text"`
}

// AsPartyInfo ...
type AsPartyInfo struct {
	PartyID          int         `json:"party_id"`
	UserLiveDeckID   int         `json:"user_live_deck_id"`
	Name             AsPartyName `json:"name"`
	IconMasterID     int         `json:"icon_master_id"`
	CardMasterID1    int         `json:"card_master_id_1"`
	CardMasterID2    int         `json:"card_master_id_2"`
	CardMasterID3    int         `json:"card_master_id_3"`
	UserAccessoryID1 any         `json:"user_accessory_id_1"`
	UserAccessoryID2 any         `json:"user_accessory_id_2"`
	UserAccessoryID3 any         `json:"user_accessory_id_3"`
}

// AsPartyName ...
type AsPartyName struct {
	DotUnderText string `json:"dot_under_text"`
}

// AsLiveStartReq ...
type AsLiveStartReq struct {
	LiveDifficultyID    int  `json:"live_difficulty_id"`
	DeckID              int  `json:"deck_id"`
	PartnerUserID       int  `json:"partner_user_id"`
	PartnerCardMasterID int  `json:"partner_card_master_id"`
	LpMagnification     int  `json:"lp_magnification"`
	IsAutoPlay          bool `json:"is_auto_play"`
	IsReferenceBook     bool `json:"is_reference_book"`
}

// AsLivePartnerInfo ...
type AsLivePartnerInfo struct {
	UserID                              int                 `json:"user_id"`
	Name                                PartnerName         `json:"name"`
	Rank                                int                 `json:"rank"`
	LastPlayedAt                        int64               `json:"last_played_at"`
	RecommendCardMasterID               int                 `json:"recommend_card_master_id"`
	RecommendCardLevel                  int                 `json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool                `json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool                `json:"is_recommend_card_all_training_activated"`
	EmblemID                            int                 `json:"emblem_id"`
	IsNew                               bool                `json:"is_new"`
	IntroductionMessage                 IntroductionMessage `json:"introduction_message"`
	FriendApprovedAt                    any                 `json:"friend_approved_at"`
	RequestStatus                       int                 `json:"request_status"`
	IsRequestPending                    bool                `json:"is_request_pending"`
}

// PartnerName ...
type PartnerName struct {
	DotUnderText string `json:"dot_under_text"`
}

// IntroductionMessage ...
type IntroductionMessage struct {
	DotUnderText string `json:"dot_under_text"`
}

// AsLiveResultAchievementStatus ...
type AsLiveResultAchievementStatus struct {
	ClearCount       int64 `json:"clear_count"`
	GotVoltage       int64 `json:"got_voltage"`
	RemainingStamina int64 `json:"remaining_stamina"`
}

// AsMvpInfo ...
type AsMvpInfo struct {
	CardMasterID        int64 `json:"card_master_id"`
	GetVoltage          int64 `json:"get_voltage"`
	SkillTriggeredCount int64 `json:"skill_triggered_count"`
	AppealCount         int64 `json:"appeal_count"`
}

// AsUserCardReq ...
type AsUserCardReq struct {
	UserID       int64 `json:"user_id"`
	CardMasterID int64 `json:"card_master_id"`
}

// AsCardAwakeningReq ...
type AsCardAwakeningReq struct {
	CardMasterID     int  `json:"card_master_id"`
	IsAwakeningImage bool `json:"is_awakening_image"`
}

// AsReq ...
type AsReq any

// AsResp ...
type AsResp any

// PackUrlReqBody ...
type PackUrlReqBody struct {
	PackNames []string `json:"pack_names"`
}

// PackUrlRespBody ...
type PackUrlRespBody struct {
	UrlList []string `json:"url_list"`
}

// LiveSaveDeckReq ...
type LiveSaveDeckReq struct {
	LiveMasterID        int   `json:"live_master_id"`
	LiveMvDeckType      int   `json:"live_mv_deck_type"`
	MemberMasterIDByPos []int `json:"member_master_id_by_pos"`
	SuitMasterIDByPos   []int `json:"suit_master_id_by_pos"`
	ViewStatusByPos     []int `json:"view_status_by_pos"`
}

// LiveSaveDeckResp ...
type LiveSaveDeckResp struct {
	UserModel UserModel `json:"user_model"`
}

// UserMemberInfo ...
type UserMemberInfo struct {
	MemberMasterID           int  `json:"member_master_id"`
	CustomBackgroundMasterID int  `json:"custom_background_master_id"`
	SuitMasterID             int  `json:"suit_master_id"`
	LovePoint                int  `json:"love_point"`
	LovePointLimit           int  `json:"love_point_limit"`
	LoveLevel                int  `json:"love_level"`
	ViewStatus               int  `json:"view_status"`
	IsNew                    bool `json:"is_new"`
}

// UserMemberByMemberID ...
type UserMemberByMemberID any

// UserLiveMvDeckInfo ...
type UserLiveMvDeckInfo struct {
	LiveMasterID     any `json:"live_master_id"`
	MemberMasterID1  any `json:"member_master_id_1"`
	MemberMasterID2  any `json:"member_master_id_2"`
	MemberMasterID3  any `json:"member_master_id_3"`
	MemberMasterID4  any `json:"member_master_id_4"`
	MemberMasterID5  any `json:"member_master_id_5"`
	MemberMasterID6  any `json:"member_master_id_6"`
	MemberMasterID7  any `json:"member_master_id_7"`
	MemberMasterID8  any `json:"member_master_id_8"`
	MemberMasterID9  any `json:"member_master_id_9"`
	MemberMasterID10 any `json:"member_master_id_10"`
	MemberMasterID11 any `json:"member_master_id_11"`
	MemberMasterID12 any `json:"member_master_id_12"`
	SuitMasterID1    any `json:"suit_master_id_1"`
	SuitMasterID2    any `json:"suit_master_id_2"`
	SuitMasterID3    any `json:"suit_master_id_3"`
	SuitMasterID4    any `json:"suit_master_id_4"`
	SuitMasterID5    any `json:"suit_master_id_5"`
	SuitMasterID6    any `json:"suit_master_id_6"`
	SuitMasterID7    any `json:"suit_master_id_7"`
	SuitMasterID8    any `json:"suit_master_id_8"`
	SuitMasterID9    any `json:"suit_master_id_9"`
	SuitMasterID10   any `json:"suit_master_id_10"`
	SuitMasterID11   any `json:"suit_master_id_11"`
	SuitMasterID12   any `json:"suit_master_id_12"`
}

// UserLiveMvDeckCustomByID ...
type UserLiveMvDeckCustomByID any

// SetThemeRes ...
type SetThemeRes any

// UserCommunicationMemberDetailBadgeByID ...
type UserCommunicationMemberDetailBadgeByID struct {
	MemberMasterID     int  `json:"member_master_id"`
	IsStoryMemberBadge bool `json:"is_story_member_badge"`
	IsStorySideBadge   bool `json:"is_story_side_badge"`
	IsVoiceBadge       bool `json:"is_voice_badge"`
	IsThemeBadge       bool `json:"is_theme_badge"`
	IsCardBadge        bool `json:"is_card_badge"`
	IsMusicBadge       bool `json:"is_music_badge"`
}

// LiveDaily ...
type LiveDaily struct {
	LiveDailyMasterID      int `json:"live_daily_master_id" xorm:"id"`
	LiveMasterID           int `json:"live_master_id" xorm:"live_id"`
	EndAt                  int `json:"end_at"`
	RemainingPlayCount     int `json:"remaining_play_count"`
	RemainingRecoveryCount int `json:"remaining_recovery_count"`
}

// AsLiveStageInfo ...
type AsLiveStageInfo struct {
	LiveDifficultyID int                  `json:"live_difficulty_id"`
	LiveNotes        []AsLiveNotes        `json:"live_notes"`
	LiveWaveSettings []AsLiveWaveSettings `json:"live_wave_settings"`
	NoteGimmicks     []AsNoteGimmicks     `json:"note_gimmicks"`
	StageGimmickDict []any                `json:"stage_gimmick_dict"`
}

// AsLiveNotes ...
type AsLiveNotes struct {
	ID                  int `json:"id"`
	CallTime            int `json:"call_time"`
	NoteType            int `json:"note_type"`
	NotePosition        int `json:"note_position"`
	GimmickID           int `json:"gimmick_id"`
	NoteAction          int `json:"note_action"`
	WaveID              int `json:"wave_id"`
	NoteRandomDropColor int `json:"note_random_drop_color"`
	AutoJudgeType       int `json:"auto_judge_type"`
}

// AsLiveWaveSettings ...
type AsLiveWaveSettings struct {
	ID            int `json:"id"`
	WaveDamage    int `json:"wave_damage"`
	MissionType   int `json:"mission_type"`
	Arg1          int `json:"arg_1"`
	Arg2          int `json:"arg_2"`
	RewardVoltage int `json:"reward_voltage"`
}

// AsNoteGimmicks ...
type AsNoteGimmicks struct {
	UniqID          int `json:"uniq_id"`
	ID              int `json:"id"`
	NoteGimmickType int `json:"note_gimmick_type"`
	Arg1            int `json:"arg_1"`
	Arg2            int `json:"arg_2"`
	EffectMID       int `json:"effect_m_id"`
	IconType        int `json:"icon_type"`
}

// AsLessonMenuAction ...
type AsLessonMenuAction struct {
	CardMasterID                  int64 `json:"card_master_id"`
	Position                      int   `json:"position"`
	IsAddedPassiveSkill           bool  `json:"is_added_passive_skill"`
	IsAddedSpecialPassiveSkill    bool  `json:"is_added_special_passive_skill"`
	IsRankupedPassiveSkill        bool  `json:"is_rankuped_passive_skill"`
	IsRankupedSpecialPassiveSkill bool  `json:"is_rankuped_special_passive_skill"`
	IsPromotedSkill               bool  `json:"is_promoted_skill"`
	MaxRarity                     any   `json:"max_rarity"`
	UpCount                       int   `json:"up_count"`
}
