#!/bin/bash
sudo date -s "$(curl -s --head http://google.com | grep ^Date: | sed 's/Date: //g') +0100"
echo Starting backend
sudo docker run --publish 3000:3000 --name backend_x --rm --device=/dev/ttyACM0 -e AWS_ACCESS_KEY_ID=AKIAUJHOO43UQH4ZR7E2 -e AWS_SECRET_ACCESS_KEY=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 backend 
echo Starting frontend
sudo docker run --publish 80:80 --name frontend_x --rm frontend