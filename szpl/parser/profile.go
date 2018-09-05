package parser

import (
	"fmt"
	"github.com/tebeka/selenium"
	"../../engine"
	"../../model"
)

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

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
