package canvas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const ISO8601DateFormat = "2006-01-02"

type API struct {
	client    http.Client
	host      string
	authToken string
}

func NewAPI(host string, authToken string) *API {
	return &API{
		host:      host,
		authToken: authToken,
	}
}

func (api *API) setAuthHeader(h *http.Header) {
	h.Set("Authorization", fmt.Sprintf("Bearer %s", api.authToken))
}

func (api *API) doWithAuth(req *http.Request) (*http.Response, error) {
	api.setAuthHeader(&req.Header)
	return api.client.Do(req)
}

func (api *API) getJSON(target interface{}, endpoint string,
	parameters url.Values) error {

	u := url.URL{
		Scheme:   "https",
		Host:     api.host,
		Path:     endpoint,
		RawQuery: parameters.Encode(),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return errors.Wrap(err, "could not create request")
	}

	resp, err := api.doWithAuth(req)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve data from Canvas API")
	}

	defer resp.Body.Close() // nolint: errcheck gosec

	if resp.StatusCode != http.StatusOK {
		return errors.New("received non-OK status from Canvas API")
	}

	d := json.NewDecoder(resp.Body)
	if err = d.Decode(target); err != nil {
		return errors.Wrap(err, "failed to decode response from Canvas API")
	}

	return nil
}

func (api *API) put(endpoint string, payload []byte) error {
	u := url.URL{
		Scheme: "https",
		Host:   api.host,
		Path:   endpoint,
	}

	req, err := http.NewRequest("PUT", u.String(), bytes.NewReader(payload))
	if err != nil {
		return errors.Wrap(err, "could not create request")
	}

	resp, err := api.doWithAuth(req)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve data from Canvas API")
	}

	defer resp.Body.Close() // nolint: errcheck gosec

	if resp.StatusCode != http.StatusNoContent &&
		resp.StatusCode != http.StatusOK {

		return errors.New("received non-OK status from Canvas API")
	}
	return nil
}

func (api *API) GetActiveCourses() ([]Course, error) {
	var courses []Course
	err := api.getJSON(&courses, "/api/v1/courses", url.Values{
		"enrollment_state": {"active"},
	})
	return courses, errors.Wrap(err, "could not retrieve course listing")
}

func (api *API) GetAnnouncementsForCourse(courseID int) ([]Discussion, error) {
	params := make(url.Values)
	params.Set("context_codes[]",
		fmt.Sprintf("course_%d", courseID))
	params.Set("start_date",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Format(ISO8601DateFormat))
	// time.Now().AddDate(-1, 0, 0).Format(ISO8601DateFormat))
	params.Set("end_date",
		time.Now().Format(ISO8601DateFormat))
	params.Set("per_page", strconv.Itoa(500))

	var announcements []Discussion
	err := api.getJSON(&announcements, "/api/v1/announcements", params)
	return announcements, errors.Wrap(err, "could not retrieve course listing")
}

func (api *API) MarkAnnouncementAsRead(courseID, discussionID int) error {
	endpoint := fmt.Sprintf("/api/v1/courses/%d/discussion_topics/%d/read",
		courseID, discussionID)

	err := api.put(endpoint, nil)
	return errors.Wrap(err, "failed to mark announcement as read")
}
