
const firebase = require("firebase");
const firebaseConf = require('../../config/firebaseConf')
firebase.initializeApp(firebaseConf);



module.exports = (app) => {
    app.post("/login", (req, res)=> {
        firebase.auth().signInWithEmailAndPassword(req.body.email, req.body.password)
        .then((data)=>{ 
            res.json({
                data
            })
        })
        .catch((err)=>{
            console.log("error ::", err);
            res.json(err);
        })
    });

    app.post("/signup", (req, res)=> {
        firebase.auth().createUserWithEmailAndPassword(req.body.email, req.body.password)
        .then(data=> {
            res.json({ 
                data,
                idToken: data.user.getIdToken()
            })
        })
        .catch(err=> {
            console.log("error ::", err);
            res.json(err);
        })
    });

    app.get("/", (req, res)=>{
        res.send("Hello from Firebase!");
    });
}