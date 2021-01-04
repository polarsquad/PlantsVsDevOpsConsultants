# Plants vs DevOps Consultants: Part 1, Introduction

Welcome to our new blog series in which a group of DevOps Consultants from Polar Squad try their hardest at applying DevOps principles in an effort to keep the plants alive.

During our ever-growing expanse in an attempt to conquer the world with Grade A DevOps Consultants, we have stumbled into the Tampere Branch office.
This site has recently shifted offices, same building, different room, and during this we have upgraded the interiors.

This meant that we are now welcomed by a total of 31 plants.

It has become our responsibility to keep these green scenery items alive and merry (we're already starting to see failures on this part).

![alt_text](images/plants_tampere.png "Plants Tampere")


# Prelude

Polar Squad is a, slightly older than, 3-year-old consulting company. We specialize in DevOps consulting and the related fields, such as: DataOps, SRE, Cloud Migrations, Agile methodologies, …

We are a company with about 50 employees with our headquarters situated in Helsinki, Finland. Branch offices can be found in Berlin, Germany, and Tampere, Finland.

As mentioned before, the Tampere office has recently upgraded their premises, as such, we have now obtained some great looking greenery.

Our goal is to keep this alive and happy, so what better way to stretch our programming muscles and design a little system to automate this.

This series of blog posts will be accompanied by code, free for all to inspect. We will be sharing the best DevOps practices, architecture choices, and evaluations of the platform which we make along with this endeavour.

Readers of this series can expect to see a walkthrough in the life of a DevOps consultant, how decisions are made, what matters to us, and mostly: how to simplify a platform creation.

During this, we encourage everyone to ask questions about the steps which we're taking on this path.
The target audience for this ranges from people looking to kickstart their own side-project, developers looking at insights with DevOps practices, architects looking at effects of their design choices, to system operators peaking at who designs the systems which they operate on a daily basis. Or anything in between and obviously: fellow DevOps Experts.

The code will be freely available, please note: we are **not** frontend devs, automation engineers, data analysts, … as such, the aspects touching on these fields will be *"best effort by DevOps consultants".*

As to answer the question about a time-schedule: this will be done as a side-gig with our main focus being our customers. This does mean that there will be quiet periods in this process from time to time.

Welcome on our road to greener pastures, and in one fell swoop we'll also announce the name of our platform: **PlantSquad**


# Platform architecture

Let's start out with designing the generic architecture of the platform. To do so, the ideal start is to define your goal.

In our case this is rather simple: keep the plants happy & alive.

Alright, now we ask ourselves: how can this be achieved? Again, a rather straightforward answer, provide them with enough light, soil nutrients, and water.

Since we're not automation robot designers we'll have to do some manual work like watering the plants and replacing the soil if that's ever needed, but we'll happily do those necessities.

So our platform will attempt to help us with obtaining that information and displaying it in a coherent manner.

To add all of this together, the platform looks at obtaining information, processing it, and presenting it back to the user.

At that point, the user can just enjoy the information or act on it and provide the needed items for the plants. This can be done in such a way that the platform even suggests what the plant is missing.

In the next section we'll take a look at the exact data models but the overarching architecture of our application should look along the following lines:

![alt_text](images/architecture_plantsquad.png "PlantSquad Architecture")


In this architecture, we can denote 3 distinct zones of operation: plant, platform, consumer.

*   Plant: A simple IoT device gathers the data from one or more sensors, does initial processing on the values, and sends this data to the platform
*   Platform: Receives and stores the data from sensors, interacts with the data possible before, during, and after storage, presents the data to the user
*   Consumer: Consumer of the data, this will mostly be done in form of a human being using a Frontend application to interpret the information from the platform, could also be a system that turns on the lights, waters the plants, etc

This already clearly shows 2 connections being in use under normal operations: `Sensor ==> Platform` and `Platform ==> Consumer`.

To make sure that all the parts of this fit neatly together, we'll have to use a standard way of communicating over these links, but more about that later on.

From now on, we can simply earmark any of the work we do in the correct category, as such our Github repository will have the 3 folders sitting neatly at the top-level directory.

**_Expert Note:_** _Plan your architecture as flued & generic as possible, you'll revisit this with each iteration. Notice how we omit any technologies from the architecture's initial draft, aiming instead for a broad design that explains our goals in the simplest manner known_


# Platform Data Models

Now that we have the generic architecture down, let's dive a bit deeper into the heart of our setup: the platform.

The goal here is to define the different sets of data models that we aim to handle with this platform.

Designing the platform will be done by interpreting these data models and their use cases.

So let's take a look at all the things we'll be measuring, and what info we'll provide to our consumers, again we aim to go as broad as possible.

**_Expert Note: Aiming for a broad spectrum prevents you from locking in decisions too early and having to deal with being stuck in later stages_**

We have 2 sets of data models here: measurements and generic information such as plant info, location, sensor.


# Measurements

This data will be provided by the IoT sensors and transported directly to the platform. Generally, these will look as such:


    Measurement


    	- id


    	- timestamp


    	- plant(s)


    	- value

Measurements that we're aiming for: brightness, temperature, moisture, conductivity.


# Generic data objects

These data objects describe items which the measurements are gathering information about.

These can are:



*   location
*   plant
*   sensor

These can be augmented later on down the road with additional models, which we will add at later stages.

**_Expert Note:_** _Stick with a few basic models which provide a minimum viable operation of your platform, no point in trying to get down every single object straight from the start_


# Data Modeling

Whichever data modeling language you use, just make sure that everybody is on the same line and keep these files always up to date.

These are your building blocks upon which your whole platform is built.

This is the template which the connectors will use to exchange information.

What this means: a person developing a sensor only needs to know the template/model in which the data has to be presented.

In our case, we have opted for Schema Defined Language ([graphql.org/learn/schema/](https://graphql.org/learn/schema/)) which is part of GraphQL.

Sadly enough there are far too few Data Modeling Languages that are not pure XML-based. We denote that using these schemas does not mean we will only use GraphQL for this implementation.

They are simply the easiest forms of clarifying our data models.

**_Expert Note: Stick with a simple data model, the actual implementations of this can be a lot more precise: protobuf, JSON, XML, ..._**


# Closing remarks

Hopefully, our path towards healthy plants will entertain our readers and give them insight into how a group of DevOps consultants would tackle the issue of creating a platform.

In the next blog posts, you can expect us to start working on the backend, selecting & designing the data storage, some hardware tidbits with the sensors, and much more.

For now, we'll have to manually judge the health of our green friends and water them adequately.

For all the code enthusiasts, the code is freely available here: [PlantsVsDevOpsConsultants](https://github.com/polarsquad/PlantsVsDevOpsConsultants)

**_Post Scriptum: The choice of data modeling language/interface description language has already sparked quite some discussion in our company, feel free to join in on the comments about what would be your go-to language for this, and why!_**
