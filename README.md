# recognition plate number and face
It is system for recognition plate number, faces and save information in database.

# Install GoCV 
go get -u -d gocv.io/x/gocv
# Install gosseract
go get -t github.com/otiai10/gosseract
# Install Go-Face
go get github.com/Kagami/go-face
# Install dlib
### Ubuntu
sudo apt-get install libdlib-dev libblas-dev liblapack-dev libjpeg-turbo8-dev
### Debian
sudo apt-get install libdlib-dev libblas-dev liblapack-dev libjpeg62-turbo-dev
# Start
go run main.go