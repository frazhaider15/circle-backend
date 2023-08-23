package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
	"github.com/rs/xid"
)

var utilObj Utils

type Utils interface {
	GenerateRandomNumber(randNumLowerLimit, randNumUpperLimit int) (string, error)
	GenerateUID() string
	GetCurrentWorkingDirectory() string
	GetRegionFromPhone(phoneNumber string) (string, error)
	GetUnixNanoFromGranularity(granularity int64) int64
	NowTime() time.Time
	SetCurrentWorkingDirectory(cwd string)
	CheckNameMatch(fullName string, firstName string, lastName string) error
	StringToByte(str string) []byte
	HashComparision(hash1 []byte, hash2 []byte) int
}
type utils struct {
	cwd string
}

func NewUtils() Utils {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return &utils{
		cwd: pwd,
	}
}

func GetUtils() Utils {

	if utilObj == nil {
		utilObj = NewUtils()
	}
	return utilObj
}

func (u *utils) GetCurrentWorkingDirectory() string {
	return u.cwd
}

func (u *utils) SetCurrentWorkingDirectory(cwd string) {
	u.cwd = cwd
}

func (u *utils) GenerateUID() string {
	uid := xid.New().String()
	return uid
}

func (u *utils) GetUnixNanoFromGranularity(granularity int64) int64 {
	timestamp := time.Now().Truncate(time.Duration(granularity) * time.Minute).Unix()
	return timestamp
}
func (u *utils) NowTime() time.Time {
	loc, _ := time.LoadLocation("UTC")
	dateTime := time.Now().In(loc)
	return dateTime
}

func (u *utils) GenerateRandomNumber(randNumLowerLimit, randNumUpperLimit int) (string, error) {
	var randomNumber string

	min := u.IntToInt64(randNumLowerLimit)
	max := u.IntToInt64(randNumUpperLimit)
	randNum, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return randomNumber, err
	}

	randomNumber = u.IntToString(u.BigIntToInt64(randNum) + min)

	return randomNumber, nil
}

// BigIntToInt64 converts big int to int64
func (u *utils) BigIntToInt64(bInt *big.Int) int64 {
	return bInt.Int64()
}

// IntToInt64 converts int to int64
func (u *utils) IntToInt64(iVal int) int64 {
	return int64(iVal)
}

// IntToString converts itn64 to
func (u *utils) IntToString(iVal int64) string {
	return fmt.Sprintf("%v", iVal)
}

func (u *utils) GetRegionFromPhone(phoneNumber string) (string, error) {
	num, err := phonenumbers.Parse(phoneNumber, "")

	if err != nil {
		return "", fmt.Errorf("error in converting number")
	}

	regionNumber := phonenumbers.GetRegionCodeForNumber(num)
	return regionNumber, nil
}

func (u *utils) CheckNameMatch(fullName string, firstName string, lastName string) error {
	signupRequestName := strings.ReplaceAll(firstName+lastName, " ", "")
	if strings.EqualFold(signupRequestName, fullName) {
		return nil
	}
	return fmt.Errorf("name does not match")
}

// StringToByte converts string to byte array
func (u *utils) StringToByte(str string) []byte {
	return []byte(str)
}

/*
HashComparision compare hash valuee
@params hash1, hash2
@return int (0 for error), (1 for success)
*/
func (u *utils) HashComparision(hash1 []byte, hash2 []byte) int {
	return subtle.ConstantTimeCompare(hash1, hash2)
}
