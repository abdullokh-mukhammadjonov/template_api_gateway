package util

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/Shopify/sarama"
)

func FormatTimeStandart(time string) string {
	return time[0:10] + "T" + time[11:19] + ".000Z"
}

func ArrayIncludes(array []string, search string) bool {
	for _, element := range array {
		if element == search {
			return true
		}
	}
	return false
}

func ParsePriceAsFloat(priceString string) (float64, error) {
	if strings.Contains(priceString, ".") {
		// if string includes a ".", parse as float
		entityPrice, err := strconv.ParseFloat(priceString, 64)
		return entityPrice, err
	} else {
		// else, parse as int
		entityPrice, err := strconv.Atoi(priceString)

		if err != nil {
			return 0, err
		}
		return float64(entityPrice), nil
	}
}

func GetFileFromMinio(cfg config.Config, fileName string, folderName string) error {
	minioClient, err := minio.New(cfg.MinioDomain, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretAccesKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println("Error minio client create: --> ", err)
		return err
	}

	exists, _ := minioClient.BucketExists(context.Background(), cfg.FilesBucketName)
	fmt.Println("BucketExists: ", exists, "bucketname: ", cfg.FilesBucketName, "filename: ", fileName)
	if !exists {
		err = minioClient.MakeBucket(context.Background(), cfg.FilesBucketName, minio.MakeBucketOptions{Region: cfg.MinioDomain})
		fmt.Println(err)
		if err != nil {
			return err
		}
	}

	err = minioClient.FGetObject(context.Background(), cfg.FilesBucketName, fileName, folderName+fileName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("Error: --> ", err, "bucketname: ", cfg.FilesBucketName, "filename: ", fileName)
		return err
	}

	return nil
}

func StringToBase64(value string) string {
	encodedText := base64.StdEncoding.EncodeToString([]byte(value))
	return encodedText
}

func CoordinatesToArray(coordinates string) []string {
	coorArray := strings.Split(strings.Trim(coordinates, " "), ",")
	newArray := []string{}
	for _, coor := range coorArray {
		if strings.HasPrefix(coor, "[[[") {
			coor = coor[3:]
		} else if strings.HasSuffix(coor, "]]]") {
			coor = coor[:len(coor)-3]
		} else if strings.HasPrefix(coor, "[") {
			coor = coor[1:]
		} else if strings.HasSuffix(coor, "]") {
			coor = coor[:len(coor)-1]
		}
		newArray = append(newArray, coor)
	}
	return newArray
}

func CoordinatesToArray3(coordinates string) [][][]string {
	coorArray := strings.Split(strings.Trim(coordinates, " "), "],")
	oldArray := [][][]string{}
	newArray := [][]string{}
	nestNewArr := []string{}
	for _, coor := range coorArray {
		nestArr := strings.Split(strings.Trim(coor, " "), ",")
		for _, v := range nestArr {
			if strings.HasPrefix(v, "[[[") {
				v = v[3:]
			} else if strings.HasSuffix(v, "]]]") {
				v = v[:len(v)-3]
			} else if strings.HasPrefix(v, "[") {
				v = v[1:]
			} else if strings.HasSuffix(v, "]") {
				v = v[:len(v)-1]
			}
			nestNewArr = append(nestNewArr, v)
		}
		newArray = append(newArray, nestNewArr)
		nestNewArr = []string{}
	}
	oldArray = append(oldArray, newArray)
	return oldArray
}

func IsFile(fileName string) (string, bool) {
	fileTypes := []string{".xlsx", ".xls", ".doc", ".docx", ".jpeg", ".jpg", ".svg", ".png", ".pdf", ".zip"}
	for _, t := range fileTypes {
		if strings.HasSuffix(strings.ToLower(fileName), t) {
			return t, true
		}
	}
	return "", false
}

func IsPhoto(fileName string) (string, bool) {
	fileTypes := []string{".jpeg", ".jpg", ".svg", ".png", ".pdf"}
	for _, t := range fileTypes {
		if strings.HasSuffix(strings.ToLower(fileName), t) {
			return t, true
		}
	}
	return "", false
}

func ClearCoordinates(prop string) string {
	prop = strings.Replace(prop, " ", "", -1)
	parts := strings.Split(prop, ",")
	rejoinParts := []string{}
	for _, part := range parts {
		cleared := strings.Replace(strings.Replace(part, "[", "", -1), "]", "", -1)
		if cleared == "0.0" {
			part = strings.Replace(part, "0.0", "0", -1)
		}
		rejoinParts = append(rejoinParts, part)
	}
	return strings.Join(rejoinParts, ",")
}

func MessageToEvent(message *sarama.ConsumerMessage) cloudevents.Event {
	event := cloudevents.NewEvent()

	for _, header := range message.Headers {
		if x := string(header.Key); x == "ce_id" {
			event.SetID(string(header.Value))
		} else if x == "ce_source" {
			event.SetSource(string(header.Value))
		} else if x == "ce_type" {
			event.SetType(string(header.Value))
		} else if x == "ce_time" {
			t, _ := time.Parse("2006-01-02T15:04:05.999999999Z", string(header.Value))
			event.SetTime(t)
		} else {
			fmt.Println("not equal: ", x)
		}
	}

	var m map[string]interface{}
	json.Unmarshal(message.Value, &m)

	event.SetData(cloudevents.ApplicationJSON, m)

	return event
}

func CheckFileExistance(f string) (*os.File, error) {
	r, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type EpiguPushInfoRes struct {
	Url      string `json:"url"`
	Response string `json:"response"`
}

type GovCommissionAnswerReq struct {
	GovCommissionFormLargeInvestmentProject GovCommissionFormLargeInvestmentProject `json:"GovCommissionFormLargeInvestmentProject"`
}

type GovCommissionFormLargeInvestmentProject struct {
	GovCommission        string `json:"gov_commission"`
	GovCommissionComment string `json:"gov_commission_comment"`
}

type CounsilAnswerReq struct {
	AnswerCounsilFormLargeInvestmentProject AnswerCounsilFormLargeInvestmentProject `json:"AnswerCounsilFormLargeInvestmentProject"`
}
type AnswerCounsilFormLargeInvestmentProject struct {
	Counsil        string `json:"answer_counsil"`
	CounsilComment string `json:"answer_counsil_comm"`
}

type EcologyAnswerReq struct {
	AnswerEcologyFormLargeInvestmentProject AnswerEcologyFormLargeInvestmentProject `json:"AnswerEcologyFormLargeInvestmentProject"`
}

type AnswerEcologyFormLargeInvestmentProject struct {
	Ecology        string `json:"answer_ecology"`
	EcologyComment string `json:"answer_ecology_comm"`
}

type CadastrAnswerReq struct {
	AnswerCadastrFormLargeInvestmentProject AnswerCadastrFormLargeInvestmentProject `json:"AnswerCadastrFormLargeInvestmentProject"`
}

type AnswerCadastrFormLargeInvestmentProject struct {
	Cadastr        string `json:"answer_cadastr"`
	CadastrComment string `json:"answer_cadastr_comm"`
}

type WorkingGroupAnswerReq struct {
	WorkingGroupFormLargeInvestmentProject WorkingGroupFormLargeInvestmentProject `json:"WorkingGroupFormLargeInvestmentProject"`
}

type WorkingGroupFormLargeInvestmentProject struct {
	WorkingGroup        string `json:"working_group"`
	WorkingGroupComment string `json:"working_group_comm"`
}

type RedirectReq struct {
	RedirectFormLargeInvestmentProject RedirectFormLargeInvestmentProject `json:"RedirectFormLargeInvestmentProject"`
}

type RedirectFormLargeInvestmentProject struct {
	FinalAnswer string `json:"final_answer"`
}

type ArchitectureAnswerReq struct {
	AnswerArchitectureFormLargeInvestmentProject AnswerArchitectureFormLargeInvestmentProject `json:"AnswerArchitectureFormLargeInvestmentProject"`
}

type AnswerArchitectureFormLargeInvestmentProject struct {
	Architecture        string `json:"answer_architecture"`
	ArchitectureComment string `json:"answer_architecture_comm"`
}
type AnswerOthersReq struct {
	AnswerOthersFormLargeInvestmentProject AnswerOthersFormLargeInvestmentProject `json:"AnswerOthersFormLargeInvestmentProject"`
}
type AnswerOthersFormLargeInvestmentProject struct {
	AnswerSanitation                 string `json:"answer_sanitation"`                   //
	AnswerEnergyComment              string `json:"answer_energy_comment"`               //
	AnswerTourismAndSportComment     string `json:"answer_tourism_and_sport_comment"`    //
	AnswerHighways                   string `json:"answer_highways"`                     //
	AnswerHighwaysComment            string `json:"answer_highways_comment"`             //
	AnswerTourismAndSport            string `json:"answer_tourism_and_sport"`            //
	AnswerWater                      string `json:"answer_water"`                        //
	AnswerInternalAffairsComment     string `json:"answer_internal_affairs_comment"`     //
	AnswerGeologyComment             string `json:"answer_geology_comment"`              //
	AnswerSanitationComment          string `json:"answer_sanitation_comment"`           //
	AnswerEnergy                     string `json:"answer_energy"`                       //
	AnswerInternalAffairs            string `json:"answer_internal_affairs"`             //
	AnswerGeology                    string `json:"answer_geology"`                      //
	AnswerHomeCommunal               string `json:"answer_home_communal"`                //
	AnswerWaterComment               string `json:"answer_water_comment"`                //
	AnswerUztransgas                 string `json:"answer_uztransgas"`                   //
	AnswerCulturalHeritageComment    string `json:"answer_cultural_heritage_comment"`    //
	AnswerHomeCommunalComment        string `json:"answer_home_communal_comment"`        //
	AnswerNationalElectricity        string `json:"answer_national_electricity"`         //
	AnswerNationalElectricityComment string `json:"answer_national_electricity_comment"` //
	AnswerUztransgasComment          string `json:"answer_uztransgas_comment"`           //
	AnswerCulturalHeritage           string `json:"answer_cultural_heritage"`            //
	AnswerLandResource               string `json:"answer_land_resource"`
	AnswerGas                        string `json:"answer_gas"`
	AnswerGasComment                 string `json:"answer_gas_comment"`
	AnswerFireSafety                 string `json:"answer_fire_safety"`
	AnswerFireSafetyComment          string `json:"answer_fire_safety_comment"`
	AnswerLandResourceComment        string `json:"answer_land_resource_comment"`
}
