const functions = require('firebase-functions');

// Create and Deploy Your First Cloud Functions
// https://firebase.google.com/docs/functions/write-firebase-functions

const firebase = require("firebase");
const express = require("express");
const app = express();
exports.helloWorld = functions.region('asia-northeast3').https.onRequest((request, response) => {
  response.send("Hello from Firebase!");
});


