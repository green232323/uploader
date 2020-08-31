package parse

import (
	"bytes"
	"github.com/dnahurnyi/uploader/clientAPI/app/contracts"
	"github.com/dnahurnyi/uploader/clientAPI/app/models"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name               string
		parser             contracts.Parser
		stream             io.ReadCloser
		expected           []models.Port
		expectedErrMsg     string
		expectedPrepErrMsg string
	}{
		{
			name:   "Success, 2 ports",
			parser: LargeJsonParser(),
			stream: ioutil.NopCloser(bytes.NewReader([]byte("{\n  \"AEAJM\": {\n    \"name\": \"Ajman\",\n    \"city\": \"Ajman\",\n    \"country\": \"United Arab Emirates\",\n    \"alias\": [],\n    \"regions\": [],\n    \"coordinates\": [\n      55.5136433,\n      25.4052165\n    ],\n    \"province\": \"Ajman\",\n    \"timezone\": \"Asia/Dubai\",\n    \"unlocs\": [\n      \"AEAJM\"\n    ],\n    \"code\": \"52000\"\n  },\n  \"AEAUH\": {\n    \"name\": \"Abu Dhabi\",\n    \"coordinates\": [\n      54.37,\n      24.47\n    ],\n    \"city\": \"Abu Dhabi\",\n    \"province\": \"Abu Z¸aby [Abu Dhabi]\",\n    \"country\": \"United Arab Emirates\",\n    \"alias\": [],\n    \"regions\": [],\n    \"timezone\": \"Asia/Dubai\",\n    \"unlocs\": [\n      \"AEAUH\"\n    ],\n    \"code\": \"52001\"\n  }"))),
			expected: []models.Port{
				{
					Key:         "AEAJM",
					Name:        "Ajman",
					City:        "Ajman",
					Country:     "United Arab Emirates",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{55.5136433, 25.4052165},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAJM"},
					Code:        "52000",
				}, {
					Key:         "AEAUH",
					Name:        "Abu Dhabi",
					City:        "Abu Dhabi",
					Country:     "United Arab Emirates",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{54.37, 24.47},
					Province:    "Abu Z¸aby [Abu Dhabi]",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAUH"},
					Code:        "52001",
				},
			},
		}, {
			name:     "Failure, interruption of the stream",
			parser:   LargeJsonParser(),
			stream:   ioutil.NopCloser(bytes.NewReader([]byte("{\n  \"AEAJM\": {\n    \"name\": \"Ajman\",\n    \"city\": \"Ajman\",\n    \"country\": \"United Arab Emirates\",\n    \"alias\": [],\n    \"regions\": [],\n    \"coordinates\": [\n      55.5136433,\n      25.4052165\n    ],"))),
			expected: []models.Port{},
		}, {
			name:               "Failure, bad data",
			parser:             LargeJsonParser(),
			stream:             ioutil.NopCloser(bytes.NewReader([]byte("bd23d9oh81ph[j9d01291x2kd1kd]120=="))),
			expected:           []models.Port{},
			expectedPrepErrMsg: "invalid character 'b' looking for beginning of value",
		}, {
			name:               "Failure, no data",
			parser:             LargeJsonParser(),
			stream:             ioutil.NopCloser(bytes.NewReader([]byte(""))),
			expected:           []models.Port{},
			expectedPrepErrMsg: "EOF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			parseFunc, prepError := tt.parser.Parse(tt.stream)
			if prepError != nil {
				assert.Equal(t, tt.expectedPrepErrMsg, prepError.Error())
				return
			}
			for {
				value, err := parseFunc()
				if value == nil {
					return
				}
				if err != nil {
					assert.Equal(t, tt.expectedErrMsg, err.Error())
					continue
				}
				assert.Equal(t, tt.expected[i].Present(), value.Present())
				i++
			}
		})
	}
}
