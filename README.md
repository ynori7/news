# News
This simple web app loads news from sources like Bild without all the ads and clutter.
This project was built in response to massive drains on mobile data.

Here's an example of the request overhead from a typical Bild page:

![Bild Loading](doc/bild-loading.png)

I've even seen it load over 15MB on the initial request when a video was embedded. The 
page continues to load content forever afterward due to tracking scripts, ads, and
other dynamic content.

Here is this application's loading statistics for the exact same page for comparison:

![News Loading](doc/news-loading.png)