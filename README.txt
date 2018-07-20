
Project Details :

Golang directory 	|
					|--> src(github.com, main.go, site_handlers.go, and views directory containing index.html

---------------------------------------------------
Download the code and run below command from terminal.

go run main.go site_handlers.go


----------------------------------------------------
Checking the ouput.
a) Run below commands from web browser.
http://localhost:8080/views/

Known Bugs :

1) When we enter some invalid url, i am just displays message on browser.
The app should display the message somewhere in the web page in label some exception.

2) Delete button is activated when the website is down.On clicking delete button, 
it removes entires for that perticular website and refresh the page.
At present , we have to do manual refresh.I will see why this is not redirecting.

3) Test handling not done from application .

