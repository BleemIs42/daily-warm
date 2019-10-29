package main

// HTML for email template
const HTML = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>每日一暖, 温情一生</title>
</head>
<body>
  <div style="max-width: 375px; margin: 20px auto;color:#444; font-size: 16px;">
    <h3 >{{one.Date}}</h3>
    <h3 style="text-align: center">{{weather.City}}</h3>
    <div style="text-align: center;font-size: 30px;">❣️</div>
    <br>
    <div style="padding: 0;width: 100%;">
      <div><span style="color: #6e6e6e">天气：</span>{{weather.Weather}}</div>
      <div><span style="color: #6e6e6e">温度：</span>{{weather.Temp}}</div>
      <div><span style="color: #6e6e6e">湿度：</span>{{weather.Humidity}}</div>
      <div><span style="color: #6e6e6e">风向：</span>{{weather.Wind}}</div>
      <div><span style="color: #6e6e6e">空气：</span>{{weather.Air}}</div>
      <div><span style="color: #6e6e6e">限行：</span>{{weather.Limit}}</div>
      <div><span style="color: #6e6e6e">提示：</span>{{weather.Note}}</div>
    </div>
    <br>
    <div style="text-align: center;font-size: 30px;">📝</div>
    <br>
    <div> 
      <div><img width="100%" src="{{english.ImgURL}}"></div>
      <div style="margin-top: 10px;line-height: 1.5">&emsp;&emsp;{{english.Sentence}}</div>
    </div>
    <br>
    <div style="text-align: center;font-size: 30px;">📖</div>
    <br>
    <div style="text-align: center">
      <div>{{poem.Title}}</div>
      <div style="font-size: 12px">{{poem.Dynasty}} {{poem.Author}}</div>
      <br>
      <div>{{poem.Content}}</div>
    </div>
    <br>
    <div style="text-align: center;font-size: 30px;">🔔</div>
    <br>
    <div>
      <div><img width="100%" src="{{one.ImgURL}}"></div>
      <div style="margin-top: 10px;line-height: 1.5">&emsp;&emsp;{{one.Sentence}}</div>
    </div>
    <br>
    <div style="text-align: center;font-size: 30px;">🏞</div>
    <br>
    <div>
      <div><img width="100%" src="{{wallpaper.ImgURL}}"></div>
      <div style="margin-top: 10px;line-height: 1.5;text-align: center;">{{wallpaper.Title}}</div>
    </div>
    <br>
    <div style="text-align: center;font-size: 30px;">📚</div>
    <br>
    <div>
      <div><img width="100%" src="{{trivia.ImgURL}}"></div>
      <div style="margin-top: 10px;line-height: 1.5">&emsp;&emsp;{{trivia.Description}}</div>
    </div>
  </div>
  <br><br>
</body>
</html>
`
