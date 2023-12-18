package smappee

// {
//   "appName": "JimmyLemmens_API",
//   "serviceLocations": [
//     {
//       "serviceLocationId": 50778,
//       "serviceLocationUuid": "e4df2db6-7b17-48e2-b508-716c6548903a",
//       "name": "Lemmens - Vanhoren",
//       "deviceSerialNumber": "5010003904"
//     },

type ServiceLocations struct {
	ServiceLocations []ServiceLocation
}

type ServiceLocation struct {
	ServiceLocationId   int
	ServiceLocationUuid string
	Name                string
	DeviceSerialNumber  string
}
