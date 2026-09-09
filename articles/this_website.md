# This Website
DESCRIPTION:The creation of this website had multiple phases in life, just like a child. First my experiments with css and html where quite terrible than I switched to a website with a neobrutalistic style but had on both no real content and ...
TAGS:wellbeing
DATE:08.09.2026 - Updated 09.09.2026
HEADIMAGE:picture_zed_programming_this_website.png
The creation of this website had multiple phases in life, just like a child. First my experiments with css and html where quite terrible than I switched to a website with a neobrutalistic style but had on both no real content and finally I decided that old style is good style. With that I mean that you just have a website with articles and that is it. No beautiful startpage, nothing, just content.
## Compilation of the Site
So you might think, what do you mean by compilation of a Site, but hear me out. I am using golang to convert my md files into html. This helps me to create a index of articles, without righting to a JSON file, but rather just right all metadata at the top of the file. When compiling the index generator just picks every file and reads out what is written in it and puts it together.
To read the file in the correct it utilises the idea of new lines and puts looks if the contain certain trademarks. These trademarks can be hastags or also selfmade trademarks for metadata. Then every line that for example has an hastag in it, will be made a h1 element. Selfmade trademarks also follow the same idea and data after HEADIMAGE gets convertet into the top picture.
So I use no DB, therefore no backend and can just be sure that I am not being hacked. This systen is not complete, at least yet at the moment of righting as I need coding fields and other elements to it.
## The Future
When further developing this Website I hope to make it as perfect as I want to and find some users on the internet that want to read it. It will definetly contain an about be section as well as projects. As soon as I continue them(I have a habit of forgetting them).
