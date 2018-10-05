package parser

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tebeka/selenium"
	"../../engine"
	"../../model"
	//"../../persist"
	"reflect"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ParseProfile(driver selenium.WebDriver) engine.ParseResult {

	webDriver := driver
	defer webDriver.Quit()


	a, err := webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[1]/td[2]")
	if err != nil{
		panic(err)
	}
	pb, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[1]/td[4]")
	if err != nil{
		panic(err)
	}
	b, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[1]/td[6]")
	if err != nil{
		panic(err)
	}
	bt, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[2]/td[2]")
	if err != nil{
		panic(err)
	}
	c, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[2]/td[4]")
	if err != nil{
		panic(err)
	}
	fpt, _ := a.Text()
	fpt = strings.Replace(fpt, "元/平方米(按建筑面积计)", "", -1)
	fpt = strings.Replace(fpt, "--", "", -1)
	fpt = strings.Replace(fpt, " ", "", -1)
	fpt = strings.Replace(fpt, "\n", "", -1)
	fp, _ := strconv.ParseFloat(fpt, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[3]/td[4]")
	if err != nil{
		panic(err)
	}
	prt, _ := a.Text()
	prt = strings.Replace(prt, " ", "", -1)
	prt = strings.Replace(prt, "\n", "", -1)
	pr, _ := strconv.ParseInt(prt, 10, 64)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[3]/td[6]")
	if err != nil{
		panic(err)
	}
	ht, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[2]")
	if err != nil{
		panic(err)
	}
	psat, _ := a.Text()
	psat = strings.Replace(psat, "平方米", "", -1)
	psat = strings.Replace(psat, " ", "", -1)
	psat = strings.Replace(psat, "\n", "", -1)
	psa, _ := strconv.ParseFloat(psat, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[4]")
	if err != nil{
		panic(err)
	}
	pucat, _ := a.Text()
	pucat = strings.Replace(pucat, "平方米", "", -1)
	pucat = strings.Replace(pucat, " ", "", -1)
	pucat = strings.Replace(pucat, "\n", "", -1)
	puca, _ := strconv.ParseFloat(pucat, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[6]")
	if err != nil{
		panic(err)
	}
	pspat, _ := a.Text()
	pspat = strings.Replace(pspat, "平方米", "", -1)
	pspat = strings.Replace(pspat, " ", "", -1)
	pspat = strings.Replace(pspat, "\n", "", -1)
	pspa, _ := strconv.ParseFloat(pspat, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[6]")
	if err != nil{
		panic(err)
	}
	csat, _ := a.Text()
	csat = strings.Replace(csat, "平方米", "", -1)
	csat = strings.Replace(csat, "--", "", -1)
	csat = strings.Replace(csat, " ", "", -1)
	csat = strings.Replace(csat, "\n", "", -1)
	csa, _ := strconv.ParseFloat(csat, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[2]")
	if err != nil{
		panic(err)
	}
	cucat, _ := a.Text()
	cucat = strings.Replace(cucat, "平方米", "", -1)
	cucat = strings.Replace(cucat, "--", "", -1)
	cucat = strings.Replace(cucat, " ", "", -1)
	cucat = strings.Replace(cucat, "\n", "", -1)
	cuca, _ := strconv.ParseFloat(cucat, 32)

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[4]")
	if err != nil{
		panic(err)
	}
	cspat, _ := a.Text()
	cspat = strings.Replace(cspat, "平方米", "", -1)
	cspat = strings.Replace(cspat, "--", "", -1)
	cspat = strings.Replace(cspat, " ", "", -1)
	cspat = strings.Replace(cspat, "\n", "", -1)
	cspa, _ := strconv.ParseFloat(cspat, 32)

	profile := model.Profile{}
	profile.ProjectBuilding = pb
	profile.Block = b
	profile.BlockType = bt
	profile.Contract = c
	profile.FilingPrice = float32(fp)
	profile.Room = pr
	profile.HouseType = ht
	profile.PreSalesArea = float32(psa)
	profile.PreUnitCstArea = float32(puca)
	profile.PreSharedPublicArea = float32(pspa)
	profile.CSalesArea = float32(csa)
	profile.CUnitCstArea = float32(cuca)
	profile.CSharedPublicArea = float32(cspa)

	fmt.Println(pb, b, bt, c, fp, pr, ht, psa, puca, pspa, csa, cuca, cspa)
	data := Struct2Map(profile)
	fmt.Println(data)

	result := engine.ParseResult{
		Items: []interface{}{data},
	}

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/szrem?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&model.Profile{})
	db.NewRecord(profile)
	db.Create(&profile)
	//persist.ItemSaver(result)  // 持久化
	return result
}
