# Unit Test Example With Generating Mocks using gomock

## Gomock
For more information on package & installation:
https://github.com/golang/mock


## Generate mocks from user.go
    mockgen -source=./services/user.go -destination=./services/mocks/mocks.go

## Generate all mocks in this repo
    make gen
