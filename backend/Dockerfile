# syntax=docker/dockerfile:1

FROM python:3.8-slim-buster
WORKDIR /app
RUN apt-get update && apt-get install build-essential -y
RUN apt-get -y install git
COPY . .
RUN pip3 install -r requirements.txt
RUN python3 setup.py install
EXPOSE 5000

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]