# crawler_house
## 爬取深圳预售房源信息
### 信息来源：深圳市规划和国资源委员会公示网站（http://ris.szpl.gov.cn/bol/）

```
go get github.com/go-sql-driver/mysql
go get github.com/tebeka/selenium
go get github.com/tebeka/selenium/chrome
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/inflection
```

Linux chromedriver
```
yum -y install unzip
yum -y install https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
wget https://chromedriver.storage.googleapis.com/2.35/chromedriver_linux64.zip
unzip chromedriver_linux64.zip
mv chromedriver /usr/bin/

```

### 效果：
<br>
<div align="center">
    <img src="pic/house.png" width="1024px">
    <br>
</div>