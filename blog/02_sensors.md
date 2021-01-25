# Plants vs DevOps Consultants: Part 2, Sensors

Welcome to our blog series in which a group of DevOps Consultants from Polar Squad try their hardest at applying DevOps principles in an effort to keep the plants alive.

By the last consensus calculation, we still have 31 plants at the office. A few of them look a bit dry, so we must get this system up and running as soon as possible!

There's a few plant parameters that we are particularly interested in.

## Amount of light

Since we're located in Tampere, somewhere around 61 degrees north the amount of natural light is highly dependend on the time of year. At the time of writing, the Sun goes down at 15:26.

Our office lighting is a bit fiddly and turns off automatically. We're not entirely sure about the suitability of generic office lights for plants either. They predominantly use 400 - 700 nm for photosynthesis, but a natural white light would be best overall.

## Humidity

Plants need water to survive. We should actually measure the moisture of the soil as well as the humidity of the surrounding air.

We're currently in more control of the amount of water the plants get from the soil, so we're going to concentrate on that.

## Fertility of the soil

In addition to light and water, plants need nutrients such as nitrate, phosphorus and potassium. These three are denoted by the N-P-K values in boxes and bottles of plant fertilizer.

The actual amount and ratio between the three is yet again dependent on the plant. Naturally, the time of year has a large effect too.

## Temperature

Tampere is the Sauna Capital of Finland, but our office temperature is somewhere around 21 degrees Celsius. This is just a guess before we actually measure it.

We're in some control of this, since we can poke the air control unit, but we're not sure how happy too much fiddling around will make our landlord. There's also people at the office in addition to plants, so we can't forget the consultants either.

## Gathering the data

Since our plants can't speak we need sensors to let us know what they are feeling in terms of the KPIs, or Key Plant Indicators, described above.

As mentioned, we're mostly in control of soil moisture, and the nutrients. If, and probably when, we get some plant LED lights with a nice and even white light we can control that too. Temperature-wise we and the plants need to take what we are given.

### Measurements

Illuminance is the parameter we are interested in related to the amount of light. It is measured in lux (lumens per square meter) and is a measure of intensity, as perceived by the human eye. The actual needed amount is dependent on the plant.

There's most likely more than one way to measure soil moisture. But some sensors use capacitance to measure the dielectric permittivity of the surrounding soil. It doesn't directly measure moisture, but the volumetric water content can be determined from the changes in capacitance between the two pads of the sensor.

Again, the amount depends largely on the plant and also the time of year. During winter time plants usually need less water than during seasons of heavy growth.

Fertility is a bit harder than that of the others, but similarly to moisture measurements can be deduced from other factors. In this case, the electrical conductivity (S/m, siemens / meter) of the soil. For this particular instance, hopefully it's in the microsiemens range though. pH of the soil is another indicator of nutrient availability.

At this point, it should come as no surprise that the desired temperature depends on the plant too. Since our office plants aren't exactly native to Finland, they need a bit more warmth than the ones outside. As for measuring temperature, there's may ways to do it from looking at mercury expanding in a glass tube to sniffing infrared. Or an electrical component called thermistor whose resistance depends on the temperature.

There's other factors involved, such as soil texture and the position of the moon, but we're feeling confident that by measuring the above we can keep our plants alive.

### Sensors

Luckily, the lovely people at Xiaomi have built a device for the occasion. There seem to be many varieties, some under the VegTrug brand. Seems to be pretty much the same device though.

All of these sticks are put in the pots next to the plants and gather all of the data needed:

- Moisture (%)
- Temperature (ºC)
- Conductivity (µS/cm)
- Illuminance (lux)
- Battery (%)

They work as Bluetooth GATT servers, so you're able to connect and read the data. GATT, or Generic Attribute Profile defines how two Bluetooth Low Energy devices can talk to each other. The clients, called Centrals, connect to the servers or Peripherals and write and read Characteristics.

Another idea that we had, would be to go a little deeper into the DIY sector and use some more generic sensors connected to an Arduino or a Raspberry Pi. But we have decided to go a different route this time, since time is of the essence!

## Conclusion

In true DevOps fashion, the ultimate goal is to automate everything. For now, we're happy if the plants are happy and are prepared to do some manual labor in terms of watering and fertilizing. Think of it as a moment of zen in the hectic world of debugging clusters, building pipelines and deploying microservices.

So, the next step is to get some sensors in the plant pots and the measurements flowing to the platform described in part 1. Some coding involved too, so let's get cracking.
