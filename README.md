# Photo-uploader

A small utility to read images from an SD card and upload them to cloud storage.
The utility should verify the file contents are properly uploaded, then delete the file contents from the SD card

Files should be named in timestamp format, eg: 20220611-1437.jpg

Non image files or image files with no timestamp should be stored in a separate directory named with their sha1 as the filename.
