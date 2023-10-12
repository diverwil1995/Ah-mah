# 網路爬蟲基礎練習
使用Golang爬蟲套件 [Colly](github.com/gocolly/colly)
[目標頁面](https://scrapeme.live/shop)
![](https://hackmd.io/_uploads/SJP_VZr-p.png)

定義要抓取的資料結構
![](https://hackmd.io/_uploads/Sk9sL-H-p.png)

找出想要的資料html tag，透過colly抓取資料後丟進上述物件
![](https://hackmd.io/_uploads/HyAtU-SWT.png)

使用Golang原生套件包os 建立csv檔案
定義這個csv檔案的欄位內容
再將抓取資料依序丟進csv
最後再匯出csv檔案