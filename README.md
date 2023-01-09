# Udactiy GO Nano-degree final project

## Project description

The project represents the backend of a customer relationship management (CRM) web application. Users interact with the app via Postman to solicit responses from the server. The server supports all of the following functionalities:

- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

## Installation, Launch, and Usage

# Installation

There is some initial setup we will need to get in place before you are able to make request to our CRM API. Below I have listed the applications and extension you will need to install along with list to their respective installation documentation:

- [Visual Studio Code](https://code.visualstudio.com/) - the preferred code editor
- [Go](https://code.visualstudio.com/docs/languages/go) - a Visual Studio Code extension to help with Go development
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode) - a code formatter for Visual Studio Code

Along with using a code editor with the above extensions, we'll also make use of the [Postman](https://www.postman.com/) app to make API requests.

# Launch

Once you are through with the installation, open the VS Code terminal and run `git clone https://github.com/robcrock/udacity-go-crm` to pull the CRM code down to your local machine.

Then cd into the `udacity-go-crm` repo and run `go run .`. That will start the server on port 3000.

# Usage

With the server up and running you can read all about the API by opening up your browser and navigating to `localhost:3000`.
