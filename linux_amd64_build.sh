gox -osarch="linux/amd64"
rm -rf blog
mkdir blog
mv blog_linux_amd64 blog/
cp config.ini blog/
cp -r static blog/
cp -r templates blog/