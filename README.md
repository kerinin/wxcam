# Weather Camera

wxcam is a Go-based app designed to run on a Raspberry Pi.  While it is built for taking weather oriented-pictures and uploading them to AWS this app could be used to take automated pictures of pretty much anything.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. Please note that currently you cannot fully test and debug this app locally as it requires the underlying raspistill application. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
- Golang
- Glide
```

### Installing

A step by step series of examples that tell you have to get a development env running

Clone the repo

```
$ git clone https://github.com/brentpabst/wxcam.git
```

Restore dependencies

```
$ glide i
```

Build it!

```
$ go build
```

## Running the tests

I'm a slacker and haven't build tests into the app yet

## Deployment & Installation

These deployment steps assume you already have a Raspberry Pi with a Camera Module installed, patched, and up to date.  We also assume you have installed the latest version of golang on the Raspberry Pi.  There are lots of instructions online for how to install golang on a Pi. These instructions assume you are building the binary on a machine other than the destination Pi, if that's not the case you may be able to omit or change a few steps.

1. Make a compatible build for the Pi

```
$ env GOOS=linux GOARCH=arm go build
```

2. Copy the resulting wxcam binary and wxcam.service file to your Pi.
3. Install the binary

```
$ sudo mv wxcam /usr/bin/wxcam
$ sudo chmod +x wxcam
```

4. Update the service file and replace variables as needed, specifically your AWS variables.  You can also override other default variables to change the intervals and scheduling of the pictures and uploads.

```
ExecStart=/usr/bin/wxcam --file-folder=/tmp/wxcam --aws-id=YOURID  --aws-secret=YOURSECRET  --s3-bucket=YOURBUCKET
```

5. Install and enable the service

```
$ sudo mv wxcam.service /lib/systemd/system/wxcam.service
$ sudo systemctl enable wxcam.service
$ sudo systemctl start wxcam.service
```

6. Check for any issues or errors with the app

```
$ tail -f /var/log/syslog
```

## Built With

* [Glide](http://glide.sh/) - Dependency Management

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/brentpabst/wxcam/tags). 

## Authors

* **Brent Pabst** - *Initial work* - [brentpabst](https://github.com/brentpabst)

See also the list of [contributors](https://github.com/brentpabst/wxcam/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
