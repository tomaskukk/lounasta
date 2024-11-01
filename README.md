### Running the application

Build the objective c targets:
`clang -c -arch arm64 -o location_manager_darwin.o location_manager_darwin.m`
`ar rcs liblocation.a location_manager_darwin.o`

Build the go binary:
`go build -gcflags "all=-N -l" -ldflags='-extldflags "-sectcreate __TEXT __info_plist /Users/tomaskukk/hox/location-v2/Info.plist" -linkmode=external' -o lounasta -v -x main.go`

Sign the binary:
`codesign -s - lounasta`

Voila!
