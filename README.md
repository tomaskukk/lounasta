### Running the application

Build the executable

#### ARM:

```bash
./build.sh
```

Some version of MacOS require signing certificate for the location API popup to show:

To create the certificate:

1. Open keychain access
2. Certificate assistant -> Create certificate
3. Fill in the details, identity type should be self signed root, type codesigning

```bash
./build.sh -c <Cert name>
```
