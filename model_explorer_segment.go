/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package strava

type ExplorerSegment struct {
	// The unique identifier of this segment
	Id int64 `json:"id,omitempty"`
	// The name of this segment
	Name string `json:"name,omitempty"`
	// The category of the climb [0, 5]. Higher is harder ie. 5 is Hors catégorie, 0 is uncategorized in climb_category. If climb_category = 5, climb_category_desc = HC. If climb_category = 2, climb_category_desc = 3.
	ClimbCategory int32 `json:"climb_category,omitempty"`
	// The description for the category of the climb
	ClimbCategoryDesc string `json:"climb_category_desc,omitempty"`
	// The segment's average grade, in percents
	AvgGrade float32 `json:"avg_grade,omitempty"`
	StartLatlng *[]float32 `json:"start_latlng,omitempty"`
	EndLatlng *[]float32 `json:"end_latlng,omitempty"`
	// The segments's evelation difference, in meters
	ElevDifference float32 `json:"elev_difference,omitempty"`
	// The segment's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The polyline of the segment
	Points string `json:"points,omitempty"`
}
