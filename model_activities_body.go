/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package strava

import (
	"time"
)

type ActivitiesBody struct {
	// The name of the activity.
	Name string `json:"name,omitempty"`
	// Type of activity. For example - Run, Ride etc.
	Type_ string `json:"type,omitempty"`
	// Sport type of activity. For example - Run, MountainBikeRide, Ride, etc.
	SportType string `json:"sport_type,omitempty"`
	// ISO 8601 formatted date time.
	StartDateLocal time.Time `json:"start_date_local"`
	// In seconds.
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// Description of the activity.
	Description string `json:"description,omitempty"`
	// In meters.
	Distance float32 `json:"distance,omitempty"`
	// Set to 1 to mark as a trainer activity.
	Trainer int32 `json:"trainer,omitempty"`
	// Set to 1 to mark as commute.
	Commute int32 `json:"commute,omitempty"`
}
