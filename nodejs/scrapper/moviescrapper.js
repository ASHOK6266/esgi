const resquest_promise = require('request-promise');

const $ = require('cheerio'); // parse the HTML 

const url = 'https://www.merriam-webster.com/dictionary/movie';



resquest_promise(url)
.then(function(html){
    console.log($('.mw-list',html).text());
})
.catch(function(err){
    //errors
});


//module.exports = movieCrawler;