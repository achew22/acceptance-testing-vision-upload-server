package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

type Monocular int

const (
	OD = iota
	OS
	BINOCULAR
)

func (m Monocular) String() string {
	switch m {
	case OD:
		return "OD"
	case OS:
		return "OS"
	case BINOCULAR:
		return "BINOCULAR"
	default:
		return "NOT KNOWN"
	}
}

type ScreeningResult int

const (
	PASS = iota
	REFER
	REFER_OR_TRY_AGAIN
)

func (s ScreeningResult) String() string {
	switch s {
	case PASS:
		return "PASS"
	case REFER:
		return "REFER"
	case REFER_OR_TRY_AGAIN:
		return "REFER_OR_TRY_AGAIN"
	default:
		return "NOT KNOWN"
	}
}

type DataRow struct {
	dateOfMeasurement *time.Time
	primaryKey        string
	familyName        string
	firstName         string
	dateOfBirth       *Date
	id                string
	location          string
	contact           string
	sphereOd          float64
	cylinderOd        float64
	axisOd            float64
	pupilSizeOd       float64
	sphereOs          float64
	cylinderOs        float64
	axisOs            float64
	pupilSizeOs       float64
	gazeAsymmetryOs   float64
	pupilDistance     float64
	monocular         Monocular
	screeningResult   ScreeningResult
}

func (d DataRow) String() string {
	return fmt.Sprintf("%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v;%v",
		d.dateOfMeasurement.Format(time.RFC3339),
		d.primaryKey,
		d.familyName,
		d.firstName,
		d.dateOfBirth,
		d.id,
		d.location,
		d.contact,
		d.sphereOd,
		d.cylinderOd,
		d.axisOd,
		d.pupilSizeOd,
		d.sphereOs,
		d.cylinderOs,
		d.axisOs,
		d.pupilSizeOs,
		d.gazeAsymmetryOs,
		d.pupilDistance,
		d.monocular,
		d.screeningResult)
}

/*
ENDLINE   = CR / CRLF
SEMICOLON = ";"

HEADER = "Date of measurement;Primary key;Family name;First name;Date of birth;ID;Location;Contact;Sphere [dpt] OD;Cylinder [dpt] OD;Axis [°] OD;Pupil size [mm] OD;Sphere [dpt] OS;Cylinder[dpt] OS;Axis [°] OS;Pupil size [mm] OS;Gaze asymmetry [°] OS;Pupil distance [mm];Monocular {1=OD, 2=OS, 3=Binocular};Screening result {0=Pass, 1=Refer, 2=Refer or try again};"
HEADER_ROW = HEADER ENDLINE

date-time = RFC3339 date-time (defined in section 5.6)
full-date = RFC3339 full-date (defined in section 5.6)

INTEGER = 1*DIGIT
NEGATION_OPERATOR = "-"
NUMBER = [NEGATION_OPERATOR]INTEGER[. INTEGER]

UTF_STRING = 1*UTF8_RUNE

DATE_OF_MEASUREMENT = date-time
PRIMARY_KEY = [INTEGER]
FAMILY_NAME = [UTF_STRING]
FIRST_NAME = [UTF_STRING]
DATE_OF_BIRTH = full-date
ID = [UTF_STRING]
LOCATION = [UTF_STRING]
CONTACT = [UTF_STRING]
SPHERE_OD = NUMBER
CYLINDER_OD = NUMBER
AXIS_OD = NUMBER
PUPIL_SIZE_OD = NUMBER
SPHERE_OS = NUMBER
CYLINDER_OS = NUMBER
AXIS_OS = NUMBER
PUPIL_SIZE_OS = NUMBER
GAZE_ASYMMETRY_OS = NUMBER
PUPIL_DISTANCE = NUMBER
MONOCULAR = 1 / 2 / 3
SCREENING_RESULT = 0 / 1 / 2

DATA_ROW = DATE_OF_MEASUREMENT SEMICOLON PRIMARY_KEY SEMICOLON FAMILY_NAME SEMICOLON FIRST_NAME SEMICOLON DATE_OF_BIRTH SEMICOLON ID SEMICOLON LOCATION SEMICOLON CONTACT SEMICOLON SPHERE_OD SEMICOLON CYLINDER_OD SEMICOLON AXIS_OD SEMICOLON PUPIL_SIZE_OD SEMICOLON SPHERE_OS SEMICOLON CYLINDER_OS SEMICOLON AXIS_OS SEMICOLON PUPIL_SIZE_OS SEMICOLON GAZE_ASYMMETRY_OS SEMICOLON PUPIL_DISTANCE SEMICOLON MONOCULAR SEMICOLON SCREENING_RESULT ENDLINE

DOCUMENT = HEADER_ROW 1*DATA_ROW
*/

func trim(t string) string {
	return strings.Trim(t, " \r\n")
}

func parseHeader(s *bufio.Scanner) error {
	headerTokens := []string{
		"", // The file comes with an empty header due to the scanner.
		"Date of measurement",
		"Primary key",
		"Family name",
		"First name",
		"Date of birth",
		"ID",
		"Location",
		"Contact",
		"Sphere [dpt] OD",
		"Cylinder [dpt] OD",
		"Axis [°] OD",
		"Pupil size [mm] OD",
		"Sphere [dpt] OS",
		"Cylinder[dpt] OS",
		"Axis [°] OS",
		"Pupil size [mm] OS",
		"Gaze asymmetry [°] OS",
		"Pupil distance [mm]",
		"Monocular {1=OD, 2=OS, 3=Binocular}",
		"Screening result {0=Pass, 1=Refer, 2=Refer or try again}",
	}
	for _, header := range headerTokens {

		if val := s.Text(); val != header {
			// Error condition
			return fmt.Errorf("Screening data header doesn't match v1 definition. Expected [%s] to equal [%s]", val, header)
		}
		if !s.Scan() {
			return fmt.Errorf("Failed to continue parsing the file")
		}
	}
	return nil
}

func parseDateTime(s *bufio.Scanner) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, trim(s.Text()))
	s.Scan()
	return &t, err
}

func parseDate(s *bufio.Scanner) (*Date, error) {
	t, err := time.Parse("2006-01-02", trim(s.Text()))
	if err != nil {
		return nil, err
	}

	s.Scan()
	return &Date{
		Year:  t.Year(),
		Month: t.Month(),
		Day:   t.Day(),
	}, nil
}

func parseNumber(s *bufio.Scanner) (float64, error) {
	t := s.Text()
	i, err := strconv.ParseFloat(trim(t), 64)
	if err != nil {
		return 0, err
	}
	s.Scan()
	return i, nil
}

func parseUTF8(s *bufio.Scanner, name string) (string, error) {
	t := s.Text()
	if trim(t) == "" {
		return "", fmt.Errorf("Expected a UTF value for %v but got no value", name)
	}
	s.Scan()
	return t, nil
}

func parseOptionalUTF8(s *bufio.Scanner) (string, error) {
	t := s.Text()
	s.Scan()
	return t, nil
}

func parseMonocular(s *bufio.Scanner) (Monocular, error) {
	t := s.Text()
	s.Scan()
	switch trim(t) {
	case "1":
		return OD, nil
	case "2":
		return OS, nil
	case "3":
		return BINOCULAR, nil
	default:
		return 0, fmt.Errorf("Expected OD, OS, or Binocular enum number. Got: %s", t)
	}
}

func parseScreeningResult(s *bufio.Scanner) (ScreeningResult, error) {
	t := s.Text()
	s.Scan()
	switch trim(t) {
	case "0":
		return PASS, nil
	case "1":
		return REFER, nil
	case "2":
		return REFER_OR_TRY_AGAIN, nil
	default:
		return 0, fmt.Errorf("Expected PASS, REFER, or REFER_OR_TRY_AGAIN enum number. Got: %s", t)
	}
}

func parseDataRow(s *bufio.Scanner) (*DataRow, error) {
	var err error

	// DATE_OF_MEASUREMENT = date-time
	dateOfMeasurement, err := parseDateTime(s)
	if err != nil {
		// This could be a terminus token. s.Err() returns nil if that was the error.
		if s.Err() == nil {
			return nil, nil
		}
		return nil, err
	}

	// PRIMARY_KEY = UTF_STRING
	primaryKey, err := parseUTF8(s, "primary key")
	if err != nil {
		return nil, err
	}

	// FAMILY_NAME = [UTF_STRING]
	familyName, err := parseOptionalUTF8(s)
	if err != nil {
		return nil, err
	}

	// FIRST_NAME = [UTF_STRING]
	firstName, err := parseOptionalUTF8(s)
	if err != nil {
		return nil, err
	}

	// DATE_OF_BIRTH = full-date
	dateOfBirth, err := parseDate(s)
	if err != nil {
		return nil, err
	}

	// ID = [UTF_STRING]
	id, err := parseOptionalUTF8(s)
	if err != nil {
		return nil, err
	}

	// LOCATION = [UTF_STRING]
	location, err := parseOptionalUTF8(s)
	if err != nil {
		return nil, err
	}

	// CONTACT = UTF_STRING
	contact, err := parseOptionalUTF8(s)
	if err != nil {
		return nil, err
	}

	// SPHERE_OD = NUMBER
	sphereOd, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// CYLINDER_OD = NUMBER
	cylinderOd, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// AXIS_OD = NUMBER
	axisOd, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// PUPIL_SIZE_OD = NUMBER
	pupilSizeOd, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// SPHERE_OS = NUMBER
	sphereOs, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// CYLINDER_OS = NUMBER
	cylinderOs, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// AXIS_OS = NUMBER
	axisOs, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// PUPIL_SIZE_OS = NUMBER
	pupilSizeOs, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// GAZE_ASYMMETRY_OS = NUMBER
	gazeAsymmetryOs, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// PUPIL_DISTANCE = NUMBER
	pupilDistance, err := parseNumber(s)
	if err != nil {
		return nil, err
	}

	// MONOCULAR = 1 / 2 / 3
	monocular, err := parseMonocular(s)
	if err != nil {
		return nil, err
	}

	// SCREENING_RESULT = 0 / 1 / 2
	screeningResult, err := parseScreeningResult(s)
	if err != nil {
		return nil, err
	}

	return &DataRow{
		dateOfMeasurement: dateOfMeasurement,
		primaryKey:        primaryKey,
		familyName:        familyName,
		firstName:         firstName,
		dateOfBirth:       dateOfBirth,
		id:                id,
		location:          location,
		contact:           contact,
		sphereOd:          sphereOd,
		cylinderOd:        cylinderOd,
		axisOd:            axisOd,
		pupilSizeOd:       pupilSizeOd,
		sphereOs:          sphereOs,
		cylinderOs:        cylinderOs,
		axisOs:            axisOs,
		pupilSizeOs:       pupilSizeOs,
		gazeAsymmetryOs:   gazeAsymmetryOs,
		pupilDistance:     pupilDistance,
		monocular:         monocular,
		screeningResult:   screeningResult,
	}, nil
}

func Parse(in io.Reader) ([]*DataRow, error) {
	s := bufio.NewScanner(in)

	line := 1
	columnEnd := 1
	tokenSize := 0
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			columnEnd += 1
			if data[i] == ';' {
				tokenSize = i + 1
				return i + 1, data[:i], nil
			} else if data[i] == '\n' {
				line += 1
				columnEnd = 1
				tokenSize = 0
			}
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	})

	rows := []*DataRow{}

	err := parseHeader(s)
	if err != nil {
		return rows, fmt.Errorf("%s\nLine: %v, Column: %v-%v\n", err, line, columnEnd-tokenSize, columnEnd)
	}

	for {
		row, err := parseDataRow(s)
		if err != nil {
			return []*DataRow{}, fmt.Errorf("Error: %v\nLine: %v, Column: %v-%v, Token: \"%s\"\n", err, line, columnEnd-tokenSize, columnEnd, trim(s.Text()))
		}
		if row == nil {
			break
		}
		rows = append(rows, row)
	}

	if s.Scan() {
		fmt.Printf("[%v]\n", s.Text())
		for s.Scan() {
			fmt.Printf("[%v]\n", s.Text())
		}
		return []*DataRow{}, fmt.Errorf("Tokens still reamin. failing...\nLine: %v-%v, Column: %v-%v, Token: \"%s\"\n", line, columnEnd-tokenSize, columnEnd, s.Text())
	}

	return rows, nil
}
