### Running the application

Build the objective c targets:

#### ARM

Use the `./run.sh` script or:

Create the object file & static library based on that

```bash
-> clang -c -arch arm64 -o location_manager_darwin.o location_manager_darwin.m
-> ar rcs liblocation.a location_manager_darwin.o
```

Build the go binary:

```bash
-> go build -gcflags "all=-N -l" -ldflags="-extldflags \"-sectcreate __TEXT __info_plist $(pwd)/Info.plist\" -linkmode=external" -o lounasta -v -x main.go
```

Sign the binary:

```bash
-> codesign -s - lounasta
```

Voila! You can now run the app by executing the binary:

```bash
-> ./lounasta
```
