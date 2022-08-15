#!/bin/bash
wails build -upx -platform=windows
wails build -upx
chmod +x build/bin/ajou

/home/seok/Limelighter/Limelighter -I build/bin/ajou.exe -O build/bin/signed_ajou.exe -Domain www.naver.com
rm build/bin/ajou.exe
mv build/bin/signed_ajou.exe build/bin/ajou.exe
