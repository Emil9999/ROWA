# syntax=docker/dockerfile:1

FROM python:3.8-slim-buster as build-stage
WORKDIR /app
RUN apt-get update && apt-get install build-essential -y
RUN apt-get -y install git
COPY . .
RUN pip install wheel
RUN python3 setup.py bdist_wheel


FROM python:3.8-slim-buster as production-stage
COPY --from=build-stage /app/dist /app
WORKDIR /app
RUN ls
RUN pip install backend-0.0.0-py3-none-any.whl
RUN pip install waitress
EXPOSE 8080
CMD [ "waitress-serve", "--call", "'main:create_app'"]
