package canvas

import "time"

type Course struct {
	ID                          int         `json:"id"`
	SisCourseID                 interface{} `json:"sis_course_id"`
	UUID                        string      `json:"uuid"`
	IntegrationID               interface{} `json:"integration_id"`
	SisImportID                 int         `json:"sis_import_id"`
	Name                        string      `json:"name"`
	CourseCode                  string      `json:"course_code"`
	WorkflowState               string      `json:"workflow_state"`
	AccountID                   int         `json:"account_id"`
	RootAccountID               int         `json:"root_account_id"`
	EnrollmentTermID            int         `json:"enrollment_term_id"`
	GradingStandardID           int         `json:"grading_standard_id"`
	GradePassbackSetting        string      `json:"grade_passback_setting"`
	CreatedAt                   string      `json:"created_at"`
	StartAt                     string      `json:"start_at"`
	EndAt                       string      `json:"end_at"`
	Locale                      string      `json:"locale"`
	Enrollments                 interface{} `json:"enrollments"`
	TotalStudents               int         `json:"total_students"`
	Calendar                    interface{} `json:"calendar"`
	DefaultView                 string      `json:"default_view"`
	SyllabusBody                string      `json:"syllabus_body"`
	NeedsGradingCount           int         `json:"needs_grading_count"`
	Term                        interface{} `json:"term"`
	CourseProgress              interface{} `json:"course_progress"`
	ApplyAssignmentGroupWeights bool        `json:"apply_assignment_group_weights"`
	Permissions                 struct {
		CreateDiscussionTopic bool `json:"create_discussion_topic"`
		CreateAnnouncement    bool `json:"create_announcement"`
	} `json:"permissions"`
	IsPublic                         bool   `json:"is_public"`
	IsPublicToAuthUsers              bool   `json:"is_public_to_auth_users"`
	PublicSyllabus                   bool   `json:"public_syllabus"`
	PublicSyllabusToAuth             bool   `json:"public_syllabus_to_auth"`
	PublicDescription                string `json:"public_description"`
	StorageQuotaMb                   int    `json:"storage_quota_mb"`
	StorageQuotaUsedMb               int    `json:"storage_quota_used_mb"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	License                          string `json:"license"`
	AllowStudentAssignmentEdits      bool   `json:"allow_student_assignment_edits"`
	AllowWikiComments                bool   `json:"allow_wiki_comments"`
	AllowStudentForumAttachments     bool   `json:"allow_student_forum_attachments"`
	OpenEnrollment                   bool   `json:"open_enrollment"`
	SelfEnrollment                   bool   `json:"self_enrollment"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
	CourseFormat                     string `json:"course_format"`
	AccessRestrictedByDate           bool   `json:"access_restricted_by_date"`
	TimeZone                         string `json:"time_zone"`
	Blueprint                        bool   `json:"blueprint"`
	BlueprintRestrictions            struct {
		Content           bool `json:"content"`
		Points            bool `json:"points"`
		DueDates          bool `json:"due_dates"`
		AvailabilityDates bool `json:"availability_dates"`
	} `json:"blueprint_restrictions"`
	BlueprintRestrictionsByObjectType struct {
		Assignment struct {
			Content bool `json:"content"`
			Points  bool `json:"points"`
		} `json:"assignment"`
		WikiPage struct {
			Content bool `json:"content"`
		} `json:"wiki_page"`
	} `json:"blueprint_restrictions_by_object_type"`
}

type Discussion struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	Message                 string      `json:"message"`
	PostedAt                time.Time   `json:"posted_at"`
	LastReplyAt             time.Time   `json:"last_reply_at"`
	RequireInitialPost      bool        `json:"require_initial_post"`
	UserCanSeePosts         bool        `json:"user_can_see_posts"`
	DiscussionSubentryCount int         `json:"discussion_subentry_count"`
	ReadState               string      `json:"read_state"`
	UnreadCount             int         `json:"unread_count"`
	Subscribed              bool        `json:"subscribed"`
	SubscriptionHold        string      `json:"subscription_hold"`
	AssignmentID            interface{} `json:"assignment_id"`
	DelayedPostAt           interface{} `json:"delayed_post_at"`
	Published               bool        `json:"published"`
	LockAt                  interface{} `json:"lock_at"`
	Locked                  bool        `json:"locked"`
	Pinned                  bool        `json:"pinned"`
	LockedForUser           bool        `json:"locked_for_user"`
	LockInfo                interface{} `json:"lock_info"`
	LockExplanation         string      `json:"lock_explanation"`
	UserName                string      `json:"user_name"`
	TopicChildren           []int       `json:"topic_children"`
	GroupTopicChildren      []struct {
		ID      int `json:"id"`
		GroupID int `json:"group_id"`
	} `json:"group_topic_children"`
	RootTopicID     interface{} `json:"root_topic_id"`
	PodcastURL      string      `json:"podcast_url"`
	DiscussionType  string      `json:"discussion_type"`
	GroupCategoryID interface{} `json:"group_category_id"`
	Attachments     interface{} `json:"attachments"`
	Permissions     struct {
		Attach bool `json:"attach"`
	} `json:"permissions"`
	AllowRating        bool `json:"allow_rating"`
	OnlyGradersCanRate bool `json:"only_graders_can_rate"`
	SortByRating       bool `json:"sort_by_rating"`
}
