# syntax=docker/dockerfile:1

FROM python:3.8-slim-buster as build-stage
WORKDIR /app
RUN apt-get update && apt-get install build-essential -y
RUN apt-get -y install git
COPY . .
RUN pip install wheel
RUN python3 setup.py bdist_wheel


FROM python:3.8-slim-buster as production-stage
RUN apt-get update && apt-get install build-essential -y
COPY --from=build-stage /app/dist /app
COPY --from=build-stage /app/requirements.txt /app
COPY --from=build-stage /app/data /app/data
WORKDIR /app
RUN pip install backend-0.0.0-py3-none-any.whl
RUN pip install -r requirements.txt
RUN pip install waitress
EXPOSE 8080
CMD [ "waitress-serve", "--call", "main:create_app"]
