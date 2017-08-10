package internal

import (
	"fmt"
	"math"
	"time"

	"github.com/pkg/errors"
)

type char byte
type unsigned_int uint
type size_t unsigned_int

/*
unsigned int uint4(unsigned const char *p) {
    return ((p[0] << 24) + (p[1] << 16) + (p[2] << 8) + p[3]);
}
*/
func uint4(p []byte) unsigned_int {
	return ((unsigned_int(p[0]) << 24) + (unsigned_int(p[1]) << 16) + (unsigned_int(p[2]) << 8) + unsigned_int(p[3]))
}

/*
int int4(unsigned const char *p) {
	int i;
	if (p[0] & 0x80) {
		i = -(((p[0] & 0x7f) << 24) + (p[1] << 16) + (p[2] << 8) + p[3]);
	}
	else {
		i = (p[0] << 24) + (p[1] << 16) + (p[2] << 8) + p[3];
	}
	return i;
}
*/
func int4(p []byte) int {
	var i int
	if (p[0] & 0x80) != 0 {
		i = -(((int(p[0]) & 0x7f) << 24) + (int(p[1]) << 16) + (int(p[2]) << 8) + int(p[3]))
	} else {
		i = (int(p[0]) << 24) + (int(p[1]) << 16) + (int(p[2]) << 8) + int(p[3])
	}
	return i
}

func fabs(v float64) float64 {
	return float64(math.Abs(float64(v)))
}

func ldexp(frac float64, exp int) float64 {
	return float64(math.Ldexp(float64(frac), exp))
}

func pow(x, y float64) float64 {
	return float64(math.Pow(float64(x), float64(y)))
}

func sqrt(x float64) float64 {
	return float64(math.Sqrt(float64(x)))
}

func Int_Power(x float64, y int) float64 {
	return float64(math.Pow(float64(x), float64(y)))
}

func exp(v float64) float64 {
	return float64(math.Exp(float64(v)))
}

func sin(v float64) float64 {
	return float64(math.Sin(float64(v)))
}

func cos(v float64) float64 {
	return float64(math.Cos(float64(v)))
}

func asin(v float64) float64 {
	return float64(math.Asin(float64(v)))
}

func atan2(x, y float64) float64 {
	return float64(math.Atan2(float64(x), float64(y)))
}

func log(v float64) float64 {
	return float64(math.Log(float64(v)))
}

func tan(v float64) float64 {
	return float64(math.Tan(float64(v)))
}

func atan(v float64) float64 {
	return float64(math.Atan(float64(v)))
}

func fatal_error(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_ii(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_i(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

func fatal_error_wrap(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args)
}

func fprintf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func LatLon(sec [][]byte, lon *[]float64, lat *[]float64) error {
	return get_latlon(sec, lon, lat)
}

func UnpackData(sec [][]byte) ([]float32, error) {
	ndata := GB2_Sec3_npts(sec)
	data := make([]float32, ndata)
	err := unpk_grib(sec, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RefTime(sec [][]byte) time.Time {
	var year, month, day, hour, minute, second int

	reftime(sec, &year, &month, &day, &hour, &minute, &second)

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}

func VerfTime(sec [][]byte) (time.Time, error) {
	var year, month, day, hour, minute, second int

	err := verftime(sec, &year, &month, &day, &hour, &minute, &second)
	if err != nil {
		return time.Now(), errors.Wrapf(err, "Failed to run verftime")
	}

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC), nil
}

func GetInfo(sec [][]byte) (name string, desc string, unit string, err error) {
	err = getName(sec, &name, &desc, &unit)
	if err != nil {
		return "", "", "", errors.Wrapf(err, "Failed to execute getName")
	}
	return name, desc, unit, nil
}

func GetLevel(sec [][]byte) (level string, err error) {
	err = f_lev(sec, &level)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to execute f_lev")
	}
	return level, nil
}
