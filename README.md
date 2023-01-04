# Udactiy GO Nano-degree final project

## Project features

The project represents the backend of a customer relationship management (CRM) web application. As users interact with the app via some user interface, your server will support all of the functionalities:

- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

## Development Strategy

You are welcome to use this overview and the Rubric specifications to create this project. Feel free to design and implement your own workflow, but if you are stuck or could use some inspiration, we've included the following walkthrough the help you get up and running:

1. **Create a representation of a Customer** as a collection of typed fields.
2. **Store your customer data** (including "seed" data) in a data structure that allows for CRUD operations. You'll want to be able to add new Customer entries to this data structure, as well as read from it, update data within it, and delete any items from it.
3. **Define the server endpoints** as well as the logical HTTP methods associated with them. For example, a user making a GET request the /customers route would probably expect a list of customers to be returned. On the other hand, POST to the same route would have different expectations. Think about how you could implement a router.
4. **Assign handlers** for each such request (e.g., a dedicated handler for GET /customers)
5. **Build each handler** to handle incoming request bodies and return logical responses. Incoming data may need to be interpreted, parsed, and/or converted to different media. How would you accomplish that?
6. **Ensure that each handler performs the necessary operations on the data structure** you choose to represent your database.

Throughout the development process, be sure to serve your application (e.g., through localhost) and test requests and responses via Postman to see if your inputs and outputs are as you'd expect.
