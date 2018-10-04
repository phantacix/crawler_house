package parser

import (
	"fmt"
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
	fp, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[3]/td[4]")
	if err != nil{
		panic(err)
	}
	pr, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[3]/td[6]")
	if err != nil{
		panic(err)
	}
	ht, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[2]")
	if err != nil{
		panic(err)
	}
	psa, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[4]")
	if err != nil{
		panic(err)
	}
	puca, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[5]/td[6]")
	if err != nil{
		panic(err)
	}
	pspa, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[6]")
	if err != nil{
		panic(err)
	}
	csa, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[2]")
	if err != nil{
		panic(err)
	}
	cuca, _ := a.Text()

	a, err = webDriver.FindElement(selenium.ByXPATH, "/html/body/center/center/table/tbody/tr/td/table/tbody/tr/td/table/tbody/tr[7]/td[4]")
	if err != nil{
		panic(err)
	}
	cspa, _ := a.Text()

	profile := model.Profile{}
	profile.ProjectBuilding = pb
	profile.Block = b
	profile.BlockType = bt
	profile.Contract = c
	profile.FilingPrice = fp
	profile.Room = pr
	profile.HouseType = ht
	profile.PreSalesArea = psa
	profile.PreUnitCstArea = puca
	profile.PreSharedPublicArea = pspa
	profile.CSalesArea = csa
	profile.CUnitCstArea = cuca
	profile.CSharedPublicArea = cspa

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
