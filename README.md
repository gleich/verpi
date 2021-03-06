<!-- DO NOT REMOVE - contributor_list:data:start:["gleich"]:end -->

# verpi

🚥 See the status of your vercel deployments on the pimoroni blinkt!

![build](https://github.com/gleich/verpi/workflows/build/badge.svg)
![test](https://github.com/gleich/verpi/workflows/test/badge.svg)
![lint](https://github.com/gleich/verpi/workflows/lint/badge.svg)
![release](https://github.com/gleich/verpi/workflows/release/badge.svg)

- [verpi](#verpi)
  - [🎥 Demo](#-demo)
  - [🚥 Setup your own version](#-setup-your-own-version)
    - [💵 Getting the parts](#-getting-the-parts)
    - [🚥 Install the pimoroni blinkt](#-install-the-pimoroni-blinkt)
    - [🖼️ Flash an image](#️-flash-an-image)
    - [🥾 Headless boot](#-headless-boot)
    - [🚀 Installing the needed deps](#-installing-the-needed-deps)
    - [🔑 Creating a token](#-creating-a-token)
    - [🚀 Installing verpi](#-installing-verpi)
    - [👋 Uninstalling verpi](#-uninstalling-verpi)
  - [⚙️ Configuring verpi](#️-configuring-verpi)
  - [🙌 Contributing](#-contributing)
  - [👥 Contributors](#-contributors)

## 🎥 Demo

https://user-images.githubusercontent.com/43759105/132117267-ec147769-a7af-4f61-bdc8-269a1fd8a466.mp4

Note: Normally all 8 LEDs would be lit up but I only have 6 projects on vercel so only 6 are lit up.

## 🚥 Setup your own version

Setting up verpi for yourself is very simple! Just follow the instructions below. If you have any problems please make an issue on this repo.

### 💵 Getting the parts

- [Case, Card, Heat sink, and other tools](https://www.amazon.com/iUniker-Raspberry-Starter-Acrylic-Clear/dp/B075FLGWJL/ref=sr_1_19?dchild=1&keywords=raspberry%2Bpi%2Bzero%2Bw&qid=1630780013&sr=8-19&th=1) - _$8.99_
- [Raspberry Pi Zero WH (Zero W with Headers)](https://www.amazon.com/Raspberry-Pi-Zero-WH-Headers/dp/B07BHMRTTY/ref=sr_1_9?crid=2VW24AF5F0854&dchild=1&keywords=raspberry+pi+zero+w&qid=1630853169&sprefix=raspberry+pi+zero+%2Caps%2C189&sr=8-9) - _$39.95_
- [Power supply](https://www.amazon.com/CanaKit-Raspberry-Supply-Adapter-Listed/dp/B00MARDJZ4/ref=sr_1_3?crid=113RLXDYJ9KMZ&dchild=1&keywords=raspberry+pi+charger&qid=1630853230&sprefix=raspberry+pi+charger%2Caps%2C176&sr=8-3) - _$9.95_
- [Pimoroni blinkt](https://shop.pimoroni.com/products/blinkt) - _~$8.32_

### 🚥 Install the pimoroni blinkt

To install the pimoroni blinkt simply set it on the GPIO headers. The correct way round is where it has curves on the top that match the corners of your Raspberry Pi.

### 🖼️ Flash an image

To flash an operating system to the micro sd card please use [the Raspberry Pi Imager program](https://www.raspberrypi.org/software/). You only need Raspberry Pi OS Lite for verpi to operate (with some dependencies installed later).

### 🥾 Headless boot

So you can ssh to the pi on boot please follow this tutorial for setting up the raspberry pi headless. It should only take a few seconds to do as you only need to make two small files.

[Headless Raspberry Pi Tutorial](https://pimylifeup.com/headless-raspberry-pi-setup/).

### 🚀 Installing the needed deps

Before actually installing verpi you need a few deps installed. Please run the following terminal command on your raspberry pi:

```sh
sudo apt -yq update && sudo apt -yq upgrade && sudo apt install -yq wiringpi git wget
```

### 🔑 Creating a token

1. Create a token on [vercel's token page](https://vercel.com/account/tokens) with a name of verpi.
2. Copy the token to your clipboard.
3. On the raspberry pi add a file to `~/.config/verpi/` called `conf.toml`.
4. Add the following to that file, replacing `<TOKEN>` with your token.

```toml
token = "<TOKEN>"
```

### 🚀 Installing verpi

To install verpi just run the following command on your raspberry pi:

```sh
wget -q -O - https://raw.githubusercontent.com/gleich/verpi/master/setup.py | sudo python3 - install
```

This script will install a temporary version of golang to produce a binary and install a systemd service. It will not mess with whatever version of go you might already have installed on the pi.

### 👋 Uninstalling verpi

To uninstall verpi just run the following command on your raspberry pi:

```sh
wget -q -O - https://raw.githubusercontent.com/gleich/verpi/master/setup.py | sudo python3 - uninstall
```

## ⚙️ Configuring verpi

Using the configuration file located at `~/.config/verpi/conf.toml` you can change the brightness of the lights. By default, the brightness is set to 0.1. To turn the lights off you can set the brightness to 0.0 as seen below:

```toml
brightness = 0.0
```

Below is an example configuration file:

```toml
token = "<YOUR VERCEL TOKEN>"
brightness = 0.0
```

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/gleich/verpi/blob/master/CONTRIBUTING.md).

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@gleich](https://github.com/gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
