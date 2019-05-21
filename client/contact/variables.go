package contact

const (
	motor = "Motor"
)

/*MotorControlMessageJSON is
{
  "subject": "motor",
  "type": "wheel",
  "motors": [
    {
      "id": 1,
      "direction": 1,
      "pwm": 50
    },
    {
      "id": 2,
      "direction": 1,
      "pwm": 50
    },
    {
      "id": 3,
      "direction": 0,
      "pwm": 80
    },
    {
      "id": 4,
      "direction": 0,
      "pwm": 80
    }
  ]
}
*/
type MotorControlMessageJSON struct {
	Subject string `json:"subject"`
	Type    string `json:"type"`
	Motors  []struct {
		ID        int `json:"id"`
		Direction int `json:"direction"`
		Pwm       int `json:"pwm"`
	} `json:"motors"`
}
