const functions = require('firebase-functions');

// Create and Deploy Your First Cloud Functions
// https://firebase.google.com/docs/functions/write-firebase-functions

const express = require("express");
const app = express();
const bodyParser = require("body-parser")
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

require('./routes/index')(app);

// API Endpoint : http://localhost:5000/blog-goland/asia-northeast3/api 
// firebase function serving express app.
exports.api = functions.region('asia-northeast3').https.onRequest(app);



