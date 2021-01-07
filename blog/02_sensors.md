# Plants vs DevOps Consultants: Part 2, Sensors

Welcome to our blog series in which a group of DevOps Consultants from Polar Squad try their hardest at applying DevOps principles in an effort to keep the plants alive.

By the last consensus calculation, we still have 31 plants at the office. A few of them look a bit dry, so we must get this system up and running as soon as possible!

## Plant parameters

There's a few plant parameters that we are particularly interested in.

### Amount of light

Since we're located in Tampere, somewhere around 61 degrees north the amount of natural light is highly dependend on the time of year. At the time of writing, the Sun goes down at 15:26.

Our office lighting is a bit fiddly and turns off automatically. We're not entirely sure about the suitability of generic office lights for plants either. They predominantly use 400 - 700 nm for photosynthesis, but a natural white light would be best overall.

Either way, illuminance is the parameter we are interested in now. It is measured in lux (lumens per square meter) and is a measure of intensity, as perceived by the human eye. The actual needed amount is dependent on the plant, we'll discuss this later.

### Humidity

Plants need water to survive. We should actually measure the moisture of the soil as well as the humidity of the surrounding air.

We're currently in more control of the amount of water the plants get from the soil, so we're going to concentrate on that.

Again, the amount depends largely on the plant and also the time of year. During winter time plants usually need less water than during seasons of heavy growth.

### Fertility of the soil

In addition to light and water, plants need nutrients such as nitrate, phosphorus and potassium. These three are denoted by the N-P-K values in boxes and bottles of plant fertilizer.

The actual amount and ratio between the three is yet again dependent on the plant. Naturally, the time of year has a large effect too.

Measuring this is a bit harder than that of the others, but can be achieved by checking the electrical conductivity (S/m, siemens / meter) of the soil. pH of the soil is another indicator of nutrient availability.

### Temperature

At this point, it should come as no surprise that the desired temperature depends on the plant too. Since our office plants aren't exactly native to Finland, they need a bit more warmth than the ones outside.

Tampere is the Sauna Capital of Finland, but our office temperature is somewhere around 21 degrees Celsius. This is just a guess before we actually measure it.

We're in some control of this, since we can poke the air control unit, but we're not sure how happy too much fiddling around will make our landlord. There's also people at the office in addition to plants, so we can't forget the consultants either.

## Sensors and measuring

Since our plants can't speak we need sensors to let us know what they are feeling in terms of the KPIs, or Key Plant Indicators, described above.

As mentioned, we're mostly in control of soil moisture, and the nutrients. If, and probably when, we get some plant LED lights with a nice and even white light we can control that too. Temperature-wise we and the plants need to take what we are given.

### Measurements

Illuminance ...

Soil moisture ...

Fertility ...

Temperature ...

### Sensors


## Conclusion

In true DevOps fashion, the ultimate goal is to automate everything. For now, we're happy if the plants are happy and are prepared to do some manual labor in terms of watering and fertilizing.

So, the next step is to get some sensors in the plant pots and the measurements flowing to the platform described in part 1. Some coding involved too, so let's get cracking.
