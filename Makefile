build::
	GOOS=linux GOARCH=amd64 go build -o ./lambda ./lambda/handler.go
	zip -j ./lambda/lambda.zip ./lambda/handler
	aws s3 cp ./lambda/lambda.zip s3://kubets/lambda.zip