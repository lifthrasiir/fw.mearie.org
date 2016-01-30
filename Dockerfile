FROM tianon/true
EXPOSE 80

ADD main index.html robots.txt favicon.ico logo.png script.js redirect.html /

CMD ["/main"]
