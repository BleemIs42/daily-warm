package main

// HTML for email template
const HTML = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>每日一暖</title>
</head>
<body>
  <div style="max-width: 375px; margin: 0 auto;">
    <h3 style="color:#1b2224">{{one.Date}}</h3>
    <h3 style="color:#1b2224; text-align: center">{{weather.City}}</h3>
    <ul style="padding: 0;width: 100%;">
      <li style="list-style: none"><img width="100%" src="{{one.ImgURL}}"></li>
      <li style="list-style: none">&emsp;&emsp;{{one.Word}}</li>
      <br><br>
      <li style="list-style: none">天气：<span style="color: #6e6e6e">{{weather.Weather}}</span></li>
      <li style="list-style: none">温度：<span style="color: #6e6e6e">{{weather.Temp}}</span></li>
      <li style="list-style: none">湿度：<span style="color: #6e6e6e">{{weather.Humidity}}</span></li>
      <li style="list-style: none">风向：<span style="color: #6e6e6e">{{weather.Wind}}</span></li>
      <li style="list-style: none">空气：<span style="color: #6e6e6e">{{weather.Air}}</span></li>
      <li style="list-style: none">限行：<span style="color: #6e6e6e">{{weather.Limit}}</span></li>
      <li style="list-style: none">提示：<span style="color: #6e6e6e">{{weather.Note}}</span></li>
    </ul>
  </div>
</body>
</html>
`
