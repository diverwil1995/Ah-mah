# 網路爬蟲筆記
### 目標網頁
[首頁|拾貨選物](https://www.tengoods.com.tw/)
![image](./img/HomePage.jpg)

### 定義預計抓取的物件結構
```json
{
    "url",    //網站連結
    "image",  //圖片  
    "name",   //名稱
    "store",  //店面  
    "price",  //價錢
    "left"    //庫存
}
```

### 使用curl指令抓取原始碼
```bash
curl -o output.html -L -c cookies.txt -b cookies.txt -A "User-Agent-String" "https://www.tengoods.com.tw/"
```

### 檢視一下是不是靜態網頁
> 能否直接找到我需要爬得資料

![image](./img/css_selector.png)

### 使用工具
Go爬蟲套件 [Colly](https://github.com/gocolly/colly)
Go標準庫 "encoding/csv", "os"

### 寫入.csv檔案並匯出至當前目錄
![image](./img/csv.png)

> 參考文章：
https://www.zenrows.com/blog/web-scraping-golang#visit-target-html-page
https://medium.com/@bob800530/selenium-1-%E9%96%8B%E5%95%9Fchrome%E7%80%8F%E8%A6%BD%E5%99%A8-21448980dff9
https://ithelp.ithome.com.tw/articles/10302836
https://blog.csdn.net/weixin_47533648/article/details/131454751
https://juejin.cn/s/golang%E7%88%AC%E8%99%AB%E6%95%99%E5%AD%A6
https://pptr.dev/
[Golang, Python優缺點對比](https://juejin.cn/post/7222168962774843451)