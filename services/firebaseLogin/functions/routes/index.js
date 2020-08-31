
const firebase = require("firebase");
const firebaseConf = require('../../config/firebaseConf')
firebase.initializeApp(firebaseConf);



module.exports = (app) => {
    app.post("/login", (req, res)=> {
        firebase.auth().signInWithEmailAndPassword(req.body.email, req.body.password)
        .then((user)=>{ 
            res.json(user)
        })
        .catch((err)=> console.log(err))
    });

    app.post("/signup", (req, res)=> {
        firebase.auth().createUserWithEmailAndPassword(req.body.email, req.body.password)
        .then(user=> {
            res.json(user)
        })
        .catch(err=> console.log(err))
    });

    app.get("/", (req, res)=>{
        res.send("Hello from Firebase!");
    });
}