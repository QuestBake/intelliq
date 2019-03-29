package common

import (
	"bufio"
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"intelliq/app/dto"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"intelliq/app/enums"
)

//GetErrorResponse prepares error response with msg
func GetErrorResponse(msg string) *dto.AppResponseDto {
	status := enums.Status.ERROR
	if msg == MSG_NO_RECORD {
		status = enums.Status.SUCCESS
	}
	return &dto.AppResponseDto{
		Status: status,
		Msg:    msg,
	}
}

//GetSuccessResponse prepares success response with body
func GetSuccessResponse(body interface{}) *dto.AppResponseDto {
	return &dto.AppResponseDto{
		Status: enums.Status.SUCCESS,
		Body:   body,
	}
}

//IsPrimaryIDValid checks if bson id is vaild
func IsPrimaryIDValid(_id bson.ObjectId) bool {
	return _id.Valid()
}

//IsStringIDValid checks if string id is vaild bson id
func IsStringIDValid(_id string) bool {
	return bson.IsObjectIdHex(_id)
}

//GetErrorMsg specific db errors
func GetErrorMsg(err error) string {
	fmt.Printf("error type: %T", err)
	if err == mgo.ErrNotFound {
		return MSG_NO_RECORD
	}
	switch err.(type) {
	case *mgo.LastError:
		errorCode := (err.(*mgo.LastError)).Code
		switch errorCode {
		case ERR_CODE_DUPLICATE:
			return MSG_DUPLICATE_RECORD
		default:
			return ""
		}
	case *mgo.QueryError:
		errorCode := (err.(*mgo.QueryError)).Code
		fmt.Println("QUERY ERROR CODE => ", errorCode)
		return err.Error()
	default:
		return ""
	}
}

//FormatDateToString formats date to readable string
func FormatDateToString(date time.Time) string {
	return date.Format(DATE_TIME_FORMAT)
}

//FormatStringToDate formats string to time
func FormatStringToDate(date string) time.Time {
	t, _ := time.Parse(DATE_TIME_FORMAT, date)
	return t
}

//EncryptData data encryption
func EncryptData(pwd string) string {
	pwdHash := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdHash, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//ComparePasswords compare hashed and plain text
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainHash := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainHash)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//IsValidMobile checks mobile number format or not
func IsValidMobile(mobile string) bool {
	_, err := strconv.ParseUint(mobile, 10, 64)
	return err == nil && len(mobile) == MOBILE_LENGTH
}

//GenerateUserName generates username from name,mobile e.g. user@FIR_MOB
func GenerateUserName(name string, mobile string) string {
	return USERNAME_PREFIX + strings.ToLower(
		name[0:USERNAME_MIN_LENGTH]) + "_" +
		mobile[MOBILE_LENGTH-USERNAME_MIN_LENGTH:MOBILE_LENGTH]
}

//IsValidGroupCode checks for groupPrefix
func IsValidGroupCode(groupCode string) bool {
	return strings.HasPrefix(groupCode,
		GROUP_CODE_PREFIX) && len(GROUP_CODE_PREFIX) < len(groupCode)
}

//GenerateRandom generated random number between give 0 & upperlimit excluding upperlimit
func GenerateRandom(lowerLimit, upperLimit int) int {
	return rand.Intn(upperLimit-lowerLimit) + lowerLimit
}

//GetMin return min of 2 int args
func GetMin(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

//PrintJSON converts obj to json string
func PrintJSON(obj interface{}) {
	tagFilter, _ := json.Marshal(obj)
	fmt.Println(string(tagFilter))
}

//GenerateHash computes hash of given obje
func GenerateHash(obj interface{}) string {
	sha512 := sha512.New()
	json := ObjectToJSON(obj)
	if json == nil {
		return ""
	}
	sha512.Write(json)
	return fmt.Sprintf("%x", sha512.Sum(nil))
}

//ObjectToJSON converts obj to json string
func ObjectToJSON(obj interface{}) []byte {
	json, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return json
}

//GenerateUUID generates random uuid
func GenerateUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	return uuid.String()
}

//ReadFile reads file content
func ReadFile(filePath string) []byte {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Read File failed: %v\n", err)
		return nil
	}
	reader := bufio.NewReader(file)
	var buffer bytes.Buffer
	for {
		var line []byte
		for {
			line, err = reader.ReadBytes('\n')
			buffer.Write(line)
			if err != nil {
				break
			}
		}
		if err == io.EOF {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf("Read File failed: %v\n", err)
		return nil
	}
	return buffer.Bytes()
}
