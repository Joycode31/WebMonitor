
WebMonitor
-------------------------------------------------------------------------------

Getting Started :

The application monitors the list of urls stored in the application map.
When a new URL is added,application checks if the url is running and update
the status to U as "Active".If the Url is not reachable then it displays a
delete button in UI so that the Url is marked for deletion.Apart from this
application constantly checks the status of added Urls at every 5 minutes.

Prerequisites :
------------------------------------------------------------------------------
1) Should be familiar with linux,windows,mac
2) Basic knowledge of web stuff.

External Dependency :
------------------------------------------------------------------------------
We have used gorilla mux router in the project.Currently i have added it as a
part of our code.

Installing : 
-------------------------------------------------------------------------------
1) Clone and download the code from my github repository mentioned below.
https://github.com/Joycode31/WebMonitor
2) Copy the project to location anywhere in the system where GOPATH is set
already or otherwise use below step to set the GOPATH.
3) Set the GOPATH in the directory where you want to run the project.
	for eg : export GOPATH="/home/atp/My_Golang_Project"
4) Go to src directory inside "WebMonitor-master" project and run below command
to execute the project.
5) go run main.go site_handlers.go

Eg : 
atp@atp-VM:~/My_Golang_Project$ cd src/
atp@atp-VM:~/My_Golang_Project/src$ go run main.go site_handlers.go



Running the tests :
--------------------------------------------------------------------------------

The test case have been added to support below functionality.
1) Check connectivity with server.
2) For redirection using mux router handling.

For running the testcase.Kindly run below command from your project directory.
Command : go test ./...

atp@atp-VM:~/My_Golang_Project$ ls -lrt
total 12
-rw-rw-r-- 1 atp atp 2237 Jul 20 11:43 README.md
drwxrwxr-x 3 atp atp 4096 Jul 20 12:15 src
drwxrwxr-x 3 atp atp 4096 Jul 20 12:39 pkg
atp@atp-VM:~/My_Golang_Project$ go test ./...
ok      _/home/atp/My_Golang_Project/pkg/mux    0.012s
ok      _/home/atp/My_Golang_Project/src        0.005s


Testing on Web :
------------------------------------------------------------------------------
http://localhost:8080/views/

Following functionality is supported.
1) Adding new url entry.
2) If the Url is down,a "delete" button will be displayed.User can delete Url 
for dormant users.
3) If the url is pingable with 800 ms time it is marked as "Active" and displayed
"Active" on browser.

Built With :
-------------------------------------------------------------------------------
Go compiler,HTML,CSS,Javascript,Jquery.


Known Bugs :
---------------------------------------------------------------------------------

1) When we enter some invalid url, i am just displays message on browser.
The app should display the message somewhere in the web page in label some 
exception.

2) Delete button is activated when the website is down.On clicking delete button, 
it removes entires for that perticular website and refresh the page.
At present , we have to do manual refresh.I will see why this is not redirecting.

3) Test handling not done from application .

