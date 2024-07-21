# web-page-analyzer

This service implements a simple form that is used to submit a valid url,
which is analysed, and a result is then shown beneath the form.

## Objective
The objective is to build a web application that does an analysis of a web-page/URL.

The application should show a form with a text field in which users can type in the URL of the
webpage to be analysed. Additionally, to the form, it should contain a button to send a request to the
server.

After processing the results should be shown to the user.

Results should contain next information:

* What HTML version has the document?
* What is the page title?
* How many headings of what level are in the document?
* How many internal and external links are in the document? Are there any inaccessible links and
* how many?
* Does the page contain a login form?

In case the URL given by the user is not reachable an error message should be presented to a user.
The message should contain the HTTP status code and a useful error description.

## Solution
The solution is using the standard http library to run a server on `localhost:8080`,
where a main page is opened with a form to submit a URL.

The page is implemented as a simple `HTML` template that is rendered using the `html/template` library.

After the URL is submitted, the service requests the URL and then uses the `golang.org/x/net/html` to parse the `HTML` response document,
which is then iterated over to get the analysis results that are rendered again using the same `HTML` template.

### Assumptions

For simplicity of the solution the following assumptions/decisions were made:
* I used the input of type `url` in the form to simplify the validation,
but it's expecting a full valid URL (incl. `scheme` and `host`)
  * `https://google.com` or `https://google.com` would work, but `google.com` is considered an invalid URL
* I skipped the unit tests and only tested it manually. As writing manual tests for an HTTP server with different scenarios
that include valid HTML pages is complicated and time-consuming.
* No authentication/authorization is implemented, as it's also not required

### How to run
I expect the environment to have Make and Docker installed.

The application can be built using the command `make build` which is building the application
and containerize it in an `ubuntu` docker container.
Which you then can run using the command `make run`, which will run the container and bind the 8080 port.
After that, it will be available at `localhost:8080`.

So in summery, you can use the following commands to build and run the application:
* `make build`
* `make run`

### Improvements
The following improvements would be needed for a production service:
* Add authentication/authorization
* Having a better Frontend
* Implement unit/integration tests