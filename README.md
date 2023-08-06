# BACKOFFICE FROM BUSINESSES

## Description

This project is a backoffice for businesses. It allows to manage the products, the orders and the customers. Initially it is used to manage the [ecommerce](https://yipipet.com) I have. With the time it will be scaleted to other businesses and people who wants to manage the important parts of their business.

## CI/CD

The pipelines are managed with github actions. The application is built in a docker image and pushed to a ecs repository. The deployment is done with a ecs task definition and a ecs service.