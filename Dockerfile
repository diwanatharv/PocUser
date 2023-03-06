# search this for everything and add comments
FROM golang:latest

#creating the build directory
RUN mkdir /app
WORKDIR /app



COPY go.mod ./
COPY go.sum ./

COPY .. ./

#so we can pull any version of package from github
RUN export GO111MODULE=on

#dowloading all the dependencies
RUN go mod download

# go build will build the code in the current directory
RUN go build -o main


#same as the port we have define in our application
EXPOSE 7000

#starting point of the program
CMD [ "./main" ]