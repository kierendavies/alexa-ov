package alexaskill

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kierendavies/alexa-ov/pkg/ov"
)

const timingPointCode = "30008225"
const stopAreaName = `<phoneme alphabet="ipa" ph="maɪˈaŋ.xʀaxt">Majanggracht</phoneme>`

func HandleRequest(ctx context.Context) (*ResponseBody, error) {
	arrival, err := ov.NextArrivalAt(timingPointCode)
	if err != nil {
		switch err.(type) {
		case *ov.NoArrivalsError:
			return NewSSMLResponse(
				fmt.Sprintf(
					`<speak>There are no upcoming arrivales at %s</speak>`,
					stopAreaName,
				),
			), nil
		default:
			return nil, err
		}
	}

	ssml := fmt.Sprintf(
		`<speak>The bus is in %s, arriving at %s at %s</speak>`,
		durationInWords(time.Until(*arrival)),
		arrival.Format("15:04"),
		stopAreaName,
	)
	return NewSSMLResponse(ssml), nil
}

func durationInWords(duration time.Duration) string {
	if duration.Minutes() < 1 {
		return "less than a minute"
	}

	minutes := int(duration.Round(time.Minute).Minutes())
	var buffer bytes.Buffer

	hours := minutes / 60
	minutes = minutes % 60
	if hours > 0 {
		buffer.WriteString(strconv.Itoa(hours))
		if hours == 1 {
			buffer.WriteString(" hour")
		} else {
			buffer.WriteString(" hours")
		}

		if minutes == 0 {
			return buffer.String()
		} else {
			buffer.WriteString(" and ")
		}
	}

	buffer.WriteString(strconv.Itoa(minutes))
	if minutes == 1 {
		buffer.WriteString(" minute")
	} else {
		buffer.WriteString(" minutes")
	}

	return buffer.String()
}
