var express = require('express');
var app = express();
var bodyParser = require('body-parser');
app.use(bodyParser.raw());

app.get('/*', function (req, res) {
  console.log("REQ:", req.query);
  /*if(req.query.name){
    console.log(req.query.name);
  }*/
  res.status(200).send('Req recieved!');
});

app.listen(3000, function () {
  console.log('Server listening on port 3000!');
});
