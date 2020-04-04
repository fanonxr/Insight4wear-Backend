package Models

// SensorDataSet{
// id:id,
// startTime='11:28:37 AM',
// endTime='11:30:18 AM',
// calorieData=CalorieData{
// calories=2.0427833},
// }

// representing the shared values that each sensor will have
type CommonSensorData struct {
	id string `json:"id"`
	starttime string `json:"starttime:"`
	endtime string `json:"endtime"`
}

// structure of activity sensor data
type ActivitySensorData struct {
	CommonSensorData
	activity int `json:"activity"`
	duration float32 `json:"duration"`
}

