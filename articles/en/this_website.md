# This Website
DESCRIPTION:The creation of this website had multiple phases in life, just like a child. First my experiments with css and html where quite terrible than I switched to a website with a neobrutalistic style but had on both no real content and ...
TAGS:wellbeing
DATE:08.09.2026 - Updated 09.09.2026 10:43 
HEADIMAGE:picture_zed_programming_this_website.png
The creation of this website had multiple phases in life, just like a child. First my experiments with css and html where quite terrible. Then I switched to a website with a neobrutalistic style but had no real content on it and finally I decided that old style is good style. With that I mean that you just have a website with articles and that is it. No beautiful startpage, just content.
## Compilation of the Site
So you might think, what do you mean by compilation of a Site, but hear me out. I am using golang to convert my md files into html. This helps me to create a index of articles, without righting to a JSON file, but rather just right all metadata at the top of the file. When compiling the index generator just picks every file and reads out what is written in it and puts it together.
To read the file in the correctly, it utilises the idea of new lines and looks a certain trademarks. These trademarks can be hastags or also selfmade trademarks for metadata. Then every line that for example has an hastag in it, will be made a h1 element. Selfmade trademarks also follow the same idea and data after HEADIMAGE gets convertet into the top picture.
As I do not use any Database this webpage is protected against typical hacking. Furthermore this systen is not complete, at least at the moment of writing as I need coding fields and other elements.
## The Future
When further developing this Website I hope to make it as perfect as I want to and find some users on the Internet that want to read it. It will definetly contain an about section as well as a projects page. As soon as I continue my projects(I have a habit of forgetting them).
