### Raspberry Pi
On a Raspberry Pi you might run into an error message like this when running `docker-compose up`:

    ERROR: http://dl-cdn.alpinelinux.org/alpine/v3.15/main: temporary error (try again later)

To fix this, follow these instructions adapted from [here](https://sensorsiot.github.io/IOTstack/Basic_setup/#patch-2-update-libseccomp2):

1. Check OS version


       $ grep "PRETTY_NAME" /etc/os-release
       PRETTY_NAME="Raspbian GNU/Linux 10 (buster)"


    If your output doesn't contain `buster`, you won't have to run the next step.

2. Run the following commands:


       sudo apt-key adv --keyserver hkps://keyserver.ubuntu.com:443 --recv-keys 04EE7237B7D453EC 648ACFD622F3D138
       echo "deb http://httpredir.debian.org/debian buster-backports main contrib non-free" | sudo tee -a "/etc/apt/sources.list.d/debian-backports.list"
       sudo apt update
       sudo apt install libseccomp2 -t buster-backports

    Now try again, it should work now.
