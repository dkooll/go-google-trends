# Google Trends Parser

This Go program allows you to parse Google Trends data and print out relevant information such as the top trending search terms, their links, and associated news headlines.

To use the program, simply run the main function and it will retrieve the daily Google Trends data for the Netherlands and parse it into a struct for easy access. The program then iterates through the top trending search terms and prints out their rank, title, link, and associated news headline.

You can also customize the program by editing the URL passed to the http.Get() function to retrieve data for a different location or time frame. Additionally, you can modify the structs and print statements to access and display different pieces of information from the Google Trends data.

Note: you can change the geo code in the url to get the relevant google trends data.

![MarineGEO circle logo](/images/structs.png "MarineGEO logo")

## Details

This Go program is designed to parse and display information from Google Trends data. The program uses the Go standard library's net/http package to retrieve the RSS feed of daily trending search terms from Google Trends, and the encoding/xml package to parse the XML data into a struct for easy access.

The program starts by defining several structs to represent the different elements of the Google Trends data. The RSS struct represents the root element of the RSS feed, which contains a single Channel element. The Channel struct contains the title of the channel and a list of Item elements, which represent each individual trending search term. Each Item contains a title, link, and approximate traffic and also has an embedded struct NewItems which contains Headline and HeadlineLink.

The main function is the entry point of the program and is where the majority of the logic takes place. In the main function, the program first initializes an empty RSS struct and calls the readGoogleTrends() function to retrieve the RSS feed data as a byte slice.

Then, the program uses the xml.Unmarshal() function to parse the byte slice into the RSS struct. If there is an error during this process, the program will print the error and exit with a status of 1.

The program then iterates through the ItemList field of the Channel struct and prints out the rank, title, link, and associated news headline of each trending search term.

The getGoogleTrends() function is used to make a GET request to the Google Trends RSS feed URL and returns the http.Response struct. The readGoogleTrends() function reads the response body and returns it as a byte slice.

The program also has comments for further usage like to print the headline links, this program is a basic version of parsing the google trends data, you can customize the program by editing the URL passed to the http.Get() function to retrieve data for a different location or time frame. Additionally, you can modify the structs and print statements to access and display different pieces of information from the Google Trends data.