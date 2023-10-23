package richdata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/test-task/internal/entity"

	"github.com/test-task/pkg/logger"
)


func RichAPI(ctx context.Context,l logger.Interface,data *entity.User)(*entity.RichData){
	//context timeout 5sec

	age :=richAge(ctx,l,data)
	gender:= richGender(ctx,l,data)
	natio:=richNatoinality(ctx,l,data)

	richData := entity.RichData{Age: age,Gender: gender,Nationality: natio}

	return &richData
}

func richAge(ctx context.Context,l logger.Interface, data *entity.User)(int){
	type Agify struct {
		Age int `json:"age"`
	}
	var agify Agify

	req := fmt.Sprintf("https://api.agify.io/?name=%s", data.Name)

	res, err := http.Get(req)
	if err != nil {
		l.Error("error making http request: %s\n", err)
		return 0 
	}

	payload,err:= io.ReadAll(res.Body)
	if err!=nil{
		l.Error("can't read api-body request agify", err)
		return 0 
	}
	
	if err = json.Unmarshal(payload,&agify);err!=nil{
		l.Error("error unmarshall agify", err)
		return 0 
	}

	return agify.Age 
}

func richGender(ctx context.Context,l logger.Interface, data *entity.User)(string){
	type Genderize struct{
			Gender string `json:gender`
	}
	var gender Genderize

	req := fmt.Sprintf("https://api.genderize.io/?name=%s", data.Name)

	res, err := http.Get(req)
	if err != nil {
		l.Error("error making http request: %s\n", err)
		return "" 
	}

	payload,err:= io.ReadAll(res.Body)
	if err!=nil{
		l.Error("can't read api-body request genderize", err)
		return "" 
	}
	
	if err = json.Unmarshal(payload, &gender);err!=nil{
		l.Error("error unmarshall genderizy", err)
		return "" 
	}

	return gender.Gender 
}

func richNatoinality(ctx context.Context,l logger.Interface, data *entity.User)(string){
	type Nationalize struct{
		Country []struct{
			CountryID string `json:"country_id"`
		} `json:country`
	}
	
	var nationality Nationalize

	req := fmt.Sprintf("https://api.nationalize.io/?name=%s", data.Name)

	res, err := http.Get(req)
	if err != nil {
		l.Error("error making http request: %s\n", err)
		return "" 
	}

	payload,err:= io.ReadAll(res.Body)
	if err!=nil{
		l.Error("can't read api-body request nationalizy", err)
		return "" 
	}
	
	if err = json.Unmarshal(payload,&nationality);err!=nil{
		l.Error("error unmarshall nationalizy", err)
		return "" 
	}

	return nationality.Country[0].CountryID 
}

