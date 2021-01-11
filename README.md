# NightBot Go
> A simplified and scalable way to connect NightBot and custom widgets.
# About This Project
### How it came to be
As a drummer who occasionally streams on [Twitch](https://www.twitch.tv/the_stoic_sudo), I wanted a way to allow users
in chat to request songs for me to play. This led me to discover [NightBot](https://nightbot.tv/), a great integration
platform that integrates with your Twitch channel to allow users to request songs in chat and add them to a queue.

The challenge then was to find a way to display this queue, along with other NightBot data, in a way that was both simplistic
and modular/scalable. I decided to write an application that integrates with NightBot's API and shows things like a streamer's
queue and current song in a widget, which could then be rendered on a streaming tool such as OBS.
### Goals of the project
- Build an application that balances ease of use with modularity and customization.
- Provide as many integrations as possible with NightBot, enabling the creation of widgets to present the data of these integrations via a web view(s).
- Make the application portable, with the ability to use it in a Docker container, on a local machine, on a hosted web server, or a Raspberry Pi. Any device
that can run web technologies can serve up the application.
### Example High-Level Usage Overview
An example of how you may use this application would be:
1. Clone and build the project
2. Add a custom widget to pull your NightBot queue data and render it in a React component.
3. Start the app and connect it to your NightBot account.
4. Add a Browser source to OBS and point it to the endpoint that exposes your custom widget.
5. Now you have a rendered view of your NightBot data in your stream!
# Installation
### Prerequisites
The following tools/dependencies are required in order to use this app:
- [Nodejs & NPM](https://nodejs.org/en/)
- [Go 1.15.2+](https://golang.org/)

This app assumes that you have correctly configured both of these tools to run from the
terminal on your machine.

### Download
```shell
$ git clone https://github.com/danielfullerton/nightbot-go.git
```

### Installing dependencies
This will install all of the JavaScript/TypeScript dependencies for the UI, as well as the Go packages for the server.
```shell
$ npm install && go get all
```

# Starting The Application
### Running on your local machine
After installing all dependencies, run the following command:
```shell
$ npm start
```
This command, which is defined in *package.json*, will build the UI application and start the server. This command
is only intended to be used on your local machine. It will assume you are running the application at the domain
name "localhost", and will start the application on port **5775**.
<br><br>
You can now [navigate to the application on in your browser](http://localhost:5775) to begin setup, which is explained
in the [Configuring Credentials](#configuring-credentials) section.

### Running with custom arguments
If you want to run the app on a different port or host the app on an external server, you can change the BASE_URL
and PORT properties via environment variables. For example:
```shell
# Build and run the project
$ npm run build:all
$ BASE_URL=http://myexampleserver.com PORT=8081 ./main

# OR alternatively, run without building Go binary
$ BASE_URL=http://myexampleserver.com PORT=8081 go run cmd/main.go
``` 

### Running with UI hot-reload capabilities
If you are creating new widgets, you may want to run the React app with hot reload abilities. This uses [webpack-dev-server](https://www.npmjs.com/package//webpack-dev-server),
and will setup a dedicated server solely for hosting the UI and NOT the Go API endpoints. However, the included **webpack.config.js** configuration
will proxy all calls to */api* to **http://localhost:5775**. This means that you can run the Go server and the webpack development server at the same
time like so:
```shell
# In one shell, run:
$ webpack-dev-server
```
```shell
# In another shell, run:
$ npm start
```

This will start the Go server on **http://localhost:5775** and the UI application on **http://localhost:4200**. You can now edit
your React widgets and watch them change in the browser in real time, and API calls will be proxied to the server in case you
want to fetch real data.

# Using The Application
First, start the application. This section will assume you ran the following command:
```shell
$ npm start
```

## Credential Configuration
For an app to integrate with NightBot, it must use a pair of keys known as a "client id" and "client secret". You can obtain
these by going to the [NightBot applications page](https://nightbot.tv/account/applications) and adding a new app.
Then, click the *pencil icon* next to that app name and click **New Secret**. Write down the Client ID and Client Secret values
somewhere secure. Finally, you will need to add this app's Redirect URI in the **Redirect URIs** box. Enter your base URL and port
for NightBot Go, followed by the path **/api/config/token**. An sample configuration looks like this:
<br><br>
[NightBotExample](app/assets/NightBotExample.png)

### Configuring Credentials via the Browser
Navigate to the [home page](http://localhost:5775) of the app. Follow the instructions on this page to set up the
app to integrate with your NightBot account. Once you have entered your credentials, click the blue **Submit** button.
<br><br>
***NOTE: Everytime you restart the application, you will also need to perform this step. It is recommended that you run
the app as a background process or on a separate server so that you can start it once and leave it alone.***
<br><br>

### Configuring Credentials via Environment Variables
This app also supports setting your NightBot *client id* and *client secret* credentials via environment variables, like so:
```shell
$ CLIENT_ID=******* CLIENT_SECRET=******* npm start
```
This is convenient if you are starting the app from another build or container process, such as [Docker](https://www.docker.com/). If you
use this method of setting your credentials, you will simply need to navigate to the app's [home page](http://localhost:5775) and click the
green **Use Environment Variables** button.

## Using Widgets
This application is written so that you can write your own widgets using React. These can either be static widgets,
such as an image or some simple markup, or they can be data-driven, such as a queue widget that pulls NightBot queue
data and populates a widget with each song title.
<br><br>
In theory, you could this application without ever integrating with NightBot just for the purpose of hosting a server
with many static widgets, or widgets that pull data from external APIs. However, that is not the intended scope of this
project.
### How a Widget Works
There are two potential ways to use widgets.
#### Single Overlay
This is the recommended usage. Using this pattern, there is a single React component; let's call it **overlay.ts**. Because this
app fetches all data and exposes it globally within the UI, you can use a single React component to consume and render all of it.
This widget will be added to a route within the React app, and will serve as the entire overlay. Your streaming app, for example
OBS, can point to this URL and render the overlay over the entire stream window.
<br><br>
For example, we may have a widget called **Overlay.tsx**, which we place on a route called **/overlay**. You could then
create a Browser source in OBS that points to **http://localhost:5775/overlay**, and this overlay widget would be rendered in
your stream window.
#### One Overlay Per Widget
Using this pattern, all widgets would belong to the direction **/app/components/**, and each is a .tsx file representing a React component. 
The standard practice is for each of these components to map to a different route, defined in **/app/app.component.tsx**. Then,
your streaming tool (for example, [OBS](https://obsproject.com/)) will point to each of these widget URLs separately to
render them.
<br><br>
For example, we may have a widget called **HelloWorld.tsx**, which we place on a route called **/hello-world**. You could then
create a Browser source in OBS that points to **http://localhost:5775/hello-world**, and this widget would be rendered in
your stream window.
##### Creating a Custom Widget (Only for One Overlay per Widget Pattern)
In order to create a custom widget, you will need some basic experience with the JavaScript framework [React](https://reactjs.org/).
First, create a new React component in the **/app/components** folder, and add this component to the routes list in **/app/app.component.tsx**
(See the existing placeholder widgets *NowPlaying* and *Queue* for examples of how to do this). From there, you can
define any custom functionality within your widget class using JavaScript or TypeScript. This includes making http calls using
the package [Axios](https://www.npmjs.com/package/axios), which you can use to call external APIs or any of the [endpoints](#api-capabilities) provided
by this application.

## API Capabilities
This section provides the API endpoints that the application exposes, which you can then use from your widgets to interact
with NightBot. This list should keep growing as time goes on and more development work is done.

### Fetch NightBot Queue
Used to fetch your current queue data from NightBot.
> **GET** **/api/queue**

```json
{
  "_total": 1,
  "_currentSong": {
    "_id":  "string",
    "createdAt": "string",
    "updatedAt":  "string",
    "track": {
      "providerId":  "string",
      "provider": "string",
      "duration":  100,
      "title": "string",
      "artist":  "string",
      "url": "string"
    },
    "user":  {
      "name":  "string",
      "displayName": "string",
      "providerId":  "string",
      "provider": "string"
    },
    "_position": 0
  },
  "_requestsEnabled": true,
  "_providers": ["strings"],
  "_searchProvider": "string",
  "status": 0,
  "queue": [{
    "_id":  "string",
    "createdAt": "string",
    "updatedAt":  "string",
    "track": {
      "providerId":  "string",
      "provider": "string",
      "duration":  100,
      "title": "string",
      "artist":  "string",
      "url": "string"
    },
    "user":  {
      "name":  "string",
      "displayName": "string",
      "providerId":  "string",
      "provider": "string"
    },
    "_position": 0
  }]
}
```

*More capabilities are in development and coming soon!*
# Fund The Project
This project is developed completely without funding and is free to use. Please consider donating to support
further development efforts and the ongoing maintenance of this project.

Venmo: [@Daniel-Fullerton-368](https://account.venmo.com/u/Daniel-Fullerton-368)

# Contributors
- [Daniel Fullerton](https://github.com/danielfullerton)

# License

### Apache License 2.0
This project is licensed under the Apache License 2.0 license. Please read [here](https://www.apache.org/licenses/LICENSE-2.0) for more information.
