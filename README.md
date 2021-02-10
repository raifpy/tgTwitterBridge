# tgTwitterBridge
Telegram Channel - Twitter Tweet Bridge 

# Preview
![gif](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/tg.gif)


# Setup
* git clone [raifpy/tgTwitterBridge](https://github.com/raifpy/tgTwitterBridge)
* Install Tamper Dev for [Chrome](https://chrome.google.com/webstore/detail/tamper-dev/mdemppnhjflbejfbnlddahjbpdbeejnn) *(Like Burp Suite)*
* Visit [twitter](https://twitter.com/home) and open Tamper Dev extention
* ![setupImage1](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_1.png)
* Tweet any text
* Find /update.json on Tamper Dev
* ![setupImage2](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_2.png)
* Focus HTTP Request layout
* ![setupImage3](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_3.png)
* Copy **Cookie's Header Value** and paste in cookie.txt **[!]**
* Copy **authorization's Header Value** and paste in header.json's authorization value **[!]**
* ![setupImage4](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_4.png)
* Open [telegram/botfather](https://t.me/botfather)
* Create new bot 
* ![setupImage5](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_5.png)
* Start your bot & Copy bot token and paste telegram.json's Token value
* Forward any message your channel to [jsonBot](https://t.me/JsonBot)
* ![setupImage6](https://github.com/raifpy/tgTwitterBridge/blob/main/resources/setup_6.png)
* Edit telegram.json with this informations
* Add your bot to your channel and don't forget to make an administrator

# On Background
```terminal
      
tgTwitterBridge[.exe] -background
      
```

# Compile 
```bash
git clone https://github.com/raifpy/tgTwitterBridge
cd tgTwitterBridge
go build -x -v .
```

You can [join my telegram channel](https://t.me/raifBlog)
