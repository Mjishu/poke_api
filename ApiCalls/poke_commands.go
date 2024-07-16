package apicalls

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
)

type APIresponse struct{
    Results []struct {
        Name string `json::"name"`
    } `json:"results"`
}

type config struct{
    prevUrl string
    nextUrl string
}

func pokelocationurl(offset int) string{
    url := "https://pokeapi.co/api/v2/location-area?limit=20"
    if offset>0{
        return url + "&offset=" + strconv.Itoa(offset)
    }
    return url
}

func CallPoke() ([]string, error){
    var sliceZero []string

    url := pokelocationurl(offset)
    
    resp,err := http.Get(url)
    if err != nil{
        return sliceZero,err
    }
    defer resp.Body.Close()

    body,err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return sliceZero,err
    }
    
    var apiResponse APIresponse
    err = json.Unmarshal(body,&apiResponse)
    if err != nil{
        return nil, fmt.Errorf("error parsing json %w",err)
    }

    var locationNames []string
    for _,result := range apiResponse.Results{
        locationNames = append(locationNames,result.Name)
    }
    return locationNames,nil
}

